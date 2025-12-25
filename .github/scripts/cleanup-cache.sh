#!/bin/bash

# GitHub Actions 自建 Runner 缓存清理脚本
# 用于定期清理过期的缓存文件，释放磁盘空间

set -euo pipefail

# 默认配置
CACHE_DIR="${CACHE_DIR:-/tmp/action}"
RETENTION_DAYS="${RETENTION_DAYS:-7}"
MAX_FILES="${MAX_FILES:-50}"
MAX_SIZE_GB="${MAX_SIZE_GB:-10}"
DRY_RUN="${DRY_RUN:-false}"
VERBOSE="${VERBOSE:-false}"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 日志函数
log() {
    echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')]${NC} $1"
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

verbose() {
    if [[ "$VERBOSE" == "true" ]]; then
        echo -e "${BLUE}[VERBOSE]${NC} $1"
    fi
}

# 显示帮助信息
show_help() {
    cat << EOF
GitHub Actions Cache Cleanup Script

Usage: $0 [OPTIONS]

Options:
    -d, --cache-dir DIR      Cache directory (default: /tmp/action)
    -r, --retention-days N   Retention days (default: 7)
    -f, --max-files N        Maximum number of cache files (default: 50)
    -s, --max-size-gb N      Maximum cache size in GB (default: 10)
    -n, --dry-run           Dry run mode, don't actually delete files
    -v, --verbose           Verbose output
    -h, --help              Show this help message

Environment Variables:
    CACHE_DIR               Same as --cache-dir
    RETENTION_DAYS          Same as --retention-days
    MAX_FILES              Same as --max-files
    MAX_SIZE_GB            Same as --max-size-gb
    DRY_RUN                Same as --dry-run (true/false)
    VERBOSE                Same as --verbose (true/false)

Examples:
    # Clean cache older than 7 days
    $0

    # Clean cache older than 3 days with verbose output
    $0 -r 3 -v

    # Dry run to see what would be deleted
    $0 -n -v

    # Clean specific cache directory
    $0 -d /opt/github-cache -r 14
EOF
}

# 解析命令行参数
parse_args() {
    while [[ $# -gt 0 ]]; do
        case $1 in
            -d|--cache-dir)
                CACHE_DIR="$2"
                shift 2
                ;;
            -r|--retention-days)
                RETENTION_DAYS="$2"
                shift 2
                ;;
            -f|--max-files)
                MAX_FILES="$2"
                shift 2
                ;;
            -s|--max-size-gb)
                MAX_SIZE_GB="$2"
                shift 2
                ;;
            -n|--dry-run)
                DRY_RUN="true"
                shift
                ;;
            -v|--verbose)
                VERBOSE="true"
                shift
                ;;
            -h|--help)
                show_help
                exit 0
                ;;
            *)
                error "Unknown option: $1"
                show_help
                exit 1
                ;;
        esac
    done
}

# 检查缓存目录
check_cache_dir() {
    if [[ ! -d "$CACHE_DIR" ]]; then
        warn "Cache directory does not exist: $CACHE_DIR"
        return 1
    fi

    verbose "Cache directory: $CACHE_DIR"
    verbose "Retention days: $RETENTION_DAYS"
    verbose "Max files: $MAX_FILES"
    verbose "Max size: ${MAX_SIZE_GB}GB"
    verbose "Dry run: $DRY_RUN"
}

# 获取缓存统计信息
get_cache_stats() {
    local total_files
    local total_size_bytes
    local total_size_gb

    total_files=$(find "$CACHE_DIR" -name "*.tar.gz" -type f | wc -l)
    
    # 兼容 Linux 和 macOS 的 stat 命令
    if stat -c%s /dev/null >/dev/null 2>&1; then
        # Linux
        total_size_bytes=$(find "$CACHE_DIR" -name "*.tar.gz" -type f -exec stat -c%s {} \; 2>/dev/null | awk '{sum+=$1} END {print sum+0}')
    else
        # macOS
        total_size_bytes=$(find "$CACHE_DIR" -name "*.tar.gz" -type f -exec stat -f%z {} \; 2>/dev/null | awk '{sum+=$1} END {print sum+0}')
    fi
    
    total_size_gb=$(echo "scale=2; $total_size_bytes / 1024 / 1024 / 1024" | bc 2>/dev/null || echo "0.00")

    log "Current cache stats:"
    log "  Files: $total_files"
    log "  Size: ${total_size_gb}GB"

    echo "$total_files:$total_size_gb"
}

# 按时间清理缓存
cleanup_by_age() {
    log "Cleaning cache files older than $RETENTION_DAYS days..."

    local old_files
    old_files=$(find "$CACHE_DIR" -name "*.tar.gz" -type f -mtime +$RETENTION_DAYS)

    if [[ -z "$old_files" ]]; then
        log "No cache files older than $RETENTION_DAYS days found"
        return 0
    fi

    local count=0
    while IFS= read -r file; do
        if [[ -n "$file" ]]; then
            verbose "Deleting old file: $file"
            if [[ "$DRY_RUN" == "false" ]]; then
                rm -f "$file"
            fi
            ((count++))
        fi
    done <<< "$old_files"

    success "Cleaned $count old cache files"
}

# 按文件数量清理缓存
cleanup_by_count() {
    local current_files
    current_files=$(find "$CACHE_DIR" -name "*.tar.gz" -type f | wc -l)

    if [[ $current_files -le $MAX_FILES ]]; then
        verbose "File count ($current_files) is within limit ($MAX_FILES)"
        return 0
    fi

    log "Too many cache files ($current_files > $MAX_FILES), cleaning oldest..."

    local excess_count=$((current_files - MAX_FILES))
    local old_files

    # 兼容 Linux 和 macOS 的 find 命令
    if find "$CACHE_DIR" -name "*.tar.gz" -type f -printf '%T@ %p\n' >/dev/null 2>&1; then
        # Linux 支持 -printf
        old_files=$(find "$CACHE_DIR" -name "*.tar.gz" -type f -printf '%T@ %p\n' | sort -n | head -n $excess_count | cut -d' ' -f2-)
    else
        # macOS 不支持 -printf，使用 stat 命令
        old_files=$(find "$CACHE_DIR" -name "*.tar.gz" -type f -exec stat -f "%m %N" {} \; | sort -n | head -n $excess_count | cut -d' ' -f2-)
    fi

    local count=0
    while IFS= read -r file; do
        if [[ -n "$file" ]]; then
            verbose "Deleting excess file: $file"
            if [[ "$DRY_RUN" == "false" ]]; then
                rm -f "$file"
            fi
            ((count++))
        fi
    done <<< "$old_files"

    success "Cleaned $count excess cache files"
}

# 按大小清理缓存
cleanup_by_size() {
    local current_size_gb
    current_size_gb=$(get_cache_stats | cut -d: -f2)

    if (( $(echo "$current_size_gb <= $MAX_SIZE_GB" | bc -l 2>/dev/null || echo "1") )); then
        verbose "Cache size (${current_size_gb}GB) is within limit (${MAX_SIZE_GB}GB)"
        return 0
    fi

    log "Cache size (${current_size_gb}GB) exceeds limit (${MAX_SIZE_GB}GB), cleaning oldest files..."

    # 计算需要删除的大小
    local excess_gb
    excess_gb=$(echo "scale=2; $current_size_gb - $MAX_SIZE_GB" | bc 2>/dev/null || echo "0")

    log "Need to free approximately ${excess_gb}GB"

    # 按时间排序，删除最老的文件直到满足大小要求
    local deleted_size=0
    local count=0

    # 兼容 Linux 和 macOS 的 find 命令
    if find "$CACHE_DIR" -name "*.tar.gz" -type f -printf '%T@ %s %p\n' >/dev/null 2>&1; then
        # Linux 支持 -printf
        find "$CACHE_DIR" -name "*.tar.gz" -type f -printf '%T@ %s %p\n' | sort -n | while read timestamp size file; do
            if (( $(echo "$deleted_size >= $excess_gb * 1024 * 1024 * 1024" | bc -l 2>/dev/null || echo "1") )); then
                break
            fi

            verbose "Deleting large cache file: $file ($(echo "scale=2; $size / 1024 / 1024" | bc 2>/dev/null || echo "0")MB)"
            if [[ "$DRY_RUN" == "false" ]]; then
                rm -f "$file"
            fi

            deleted_size=$((deleted_size + size))
            ((count++))
        done
    else
        # macOS 不支持 -printf，使用 stat 命令
        find "$CACHE_DIR" -name "*.tar.gz" -type f | while read file; do
            local timestamp
            local size
            timestamp=$(stat -f "%m" "$file")
            size=$(stat -f "%z" "$file")
            echo "${timestamp} ${size} ${file}"
        done | sort -n | while read timestamp size file; do
            if (( $(echo "$deleted_size >= $excess_gb * 1024 * 1024 * 1024" | bc -l 2>/dev/null || echo "1") )); then
                break
            fi

            verbose "Deleting large cache file: $file ($(echo "scale=2; $size / 1024 / 1024" | bc 2>/dev/null || echo "0")MB)"
            if [[ "$DRY_RUN" == "false" ]]; then
                rm -f "$file"
            fi

            deleted_size=$((deleted_size + size))
            ((count++))
        done
    fi

    success "Cleaned $count cache files to free space"
}

# 清理空目录
cleanup_empty_dirs() {
    log "Cleaning empty directories..."

    local empty_dirs
    empty_dirs=$(find "$CACHE_DIR" -type d -empty)

    local count=0
    while IFS= read -r dir; do
        if [[ -n "$dir" && "$dir" != "$CACHE_DIR" ]]; then
            verbose "Removing empty directory: $dir"
            if [[ "$DRY_RUN" == "false" ]]; then
                rmdir "$dir" 2>/dev/null || true
            fi
            ((count++))
        fi
    done <<< "$empty_dirs"

    if [[ $count -gt 0 ]]; then
        success "Removed $count empty directories"
    fi
}

# 生成统计报告
generate_stats() {
    local stats_file="$CACHE_DIR/cleanup-stats.json"
    local stats
    stats=$(get_cache_stats)
    local file_count
    local size_gb
    file_count=$(echo "$stats" | cut -d: -f1)
    size_gb=$(echo "$stats" | cut -d: -f2)

    cat > "$stats_file" << EOF
{
  "timestamp": "$(date -u +%Y-%m-%dT%H:%M:%SZ)",
  "cache_dir": "$CACHE_DIR",
  "retention_days": $RETENTION_DAYS,
  "max_files": $MAX_FILES,
  "max_size_gb": $MAX_SIZE_GB,
  "current_files": $file_count,
  "current_size_gb": $size_gb,
  "dry_run": $DRY_RUN
}
EOF

    verbose "Stats saved to: $stats_file"
}

# 主函数
main() {
    parse_args "$@"

    if ! check_cache_dir; then
        exit 1
    fi

    if [[ "$DRY_RUN" == "true" ]]; then
        warn "Running in DRY RUN mode - no files will be deleted"
    fi

    log "Starting cache cleanup..."

    # 显示清理前的统计信息
    get_cache_stats > /dev/null

    # 执行清理
    cleanup_by_age
    cleanup_by_count
    cleanup_by_size
    cleanup_empty_dirs

    # 显示清理后的统计信息
    log "Cache cleanup completed!"
    get_cache_stats > /dev/null

    # 生成统计报告
    generate_stats
}

# 运行主函数
main "$@"