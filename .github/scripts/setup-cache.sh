#!/bin/bash

# GitHub Actions 自建 Runner 缓存设置脚本
# 用于快速设置本地缓存环境

set -euo pipefail

# 默认配置
CACHE_DIR="${CACHE_DIR:-/tmp/action}"
RUNNER_USER="${RUNNER_USER:-runner}"
SETUP_CRON="${SETUP_CRON:-true}"
CREATE_MARKER="${CREATE_MARKER:-true}"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log() {
    echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')]${NC} $1"
}

success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 显示帮助信息
show_help() {
    cat << EOF
GitHub Actions Cache Setup Script

Usage: $0 [OPTIONS]

Options:
    -d, --cache-dir DIR      Cache directory (default: /tmp/action)
    -u, --runner-user USER   Runner user (default: runner)
    -c, --setup-cron         Setup cron job for cleanup (default: true)
    -m, --create-marker      Create self-hosted marker file (default: true)
    -h, --help              Show this help message

Examples:
    # Basic setup
    $0

    # Custom cache directory
    $0 -d /opt/github-cache

    # Setup without cron
    $0 --setup-cron false
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
            -u|--runner-user)
                RUNNER_USER="$2"
                shift 2
                ;;
            -c|--setup-cron)
                SETUP_CRON="$2"
                shift 2
                ;;
            -m|--create-marker)
                CREATE_MARKER="$2"
                shift 2
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

# 检查权限
check_permissions() {
    if [[ $EUID -ne 0 ]] && [[ "$CREATE_MARKER" == "true" || "$SETUP_CRON" == "true" ]]; then
        error "This script needs to be run as root for full setup"
        error "Run with sudo or disable marker/cron setup"
        exit 1
    fi
}

# 创建缓存目录
setup_cache_dir() {
    log "Setting up cache directory: $CACHE_DIR"

    # 创建主目录
    mkdir -p "$CACHE_DIR"
    mkdir -p "$CACHE_DIR/go-modules"
    mkdir -p "$CACHE_DIR/go-build"

    # 设置权限
    if id "$RUNNER_USER" &>/dev/null; then
        chown -R "$RUNNER_USER:$RUNNER_USER" "$CACHE_DIR"
        log "Set ownership to $RUNNER_USER"
    else
        warn "User $RUNNER_USER does not exist, skipping ownership change"
    fi

    chmod -R 755 "$CACHE_DIR"

    success "Cache directory setup completed"
}

# 创建自建 runner 标识文件
create_marker_file() {
    if [[ "$CREATE_MARKER" != "true" ]]; then
        log "Skipping marker file creation"
        return 0
    fi

    log "Creating self-hosted runner marker file"

    mkdir -p /etc/github-runner
    touch /etc/github-runner/self-hosted
    chmod 644 /etc/github-runner/self-hosted

    success "Marker file created: /etc/github-runner/self-hosted"
}

# 设置清理任务
setup_cleanup_cron() {
    if [[ "$SETUP_CRON" != "true" ]]; then
        log "Skipping cron setup"
        return 0
    fi

    log "Setting up cache cleanup cron job"

    local script_path
    script_path="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/cleanup-cache.sh"

    if [[ ! -f "$script_path" ]]; then
        error "Cleanup script not found: $script_path"
        return 1
    fi

    # 为 runner 用户添加 cron 任务
    local cron_job="0 2 * * * $script_path -d $CACHE_DIR -r 7 >/dev/null 2>&1"

    if id "$RUNNER_USER" &>/dev/null; then
        # 添加到 runner 用户的 crontab
        (crontab -u "$RUNNER_USER" -l 2>/dev/null; echo "$cron_job") | crontab -u "$RUNNER_USER" -
        success "Cron job added for user $RUNNER_USER"
    else
        # 添加到系统 crontab
        echo "$cron_job" >> /etc/crontab
        success "Cron job added to system crontab"
    fi

    log "Cleanup will run daily at 2:00 AM"
}

# 验证设置
verify_setup() {
    log "Verifying setup..."

    # 检查缓存目录
    if [[ -d "$CACHE_DIR" ]]; then
        success "Cache directory exists: $CACHE_DIR"
    else
        error "Cache directory not found: $CACHE_DIR"
        return 1
    fi

    # 检查权限
    if [[ -w "$CACHE_DIR" ]]; then
        success "Cache directory is writable"
    else
        error "Cache directory is not writable"
        return 1
    fi

    # 检查标识文件
    if [[ "$CREATE_MARKER" == "true" && -f "/etc/github-runner/self-hosted" ]]; then
        success "Self-hosted marker file exists"
    fi

    # 测试缓存功能
    local test_file="$CACHE_DIR/test-$(date +%s).txt"
    if echo "test" > "$test_file" 2>/dev/null; then
        rm -f "$test_file"
        success "Cache write test passed"
    else
        error "Cache write test failed"
        return 1
    fi
}

# 显示设置摘要
show_summary() {
    log "Setup Summary:"
    echo "  Cache Directory: $CACHE_DIR"
    echo "  Runner User: $RUNNER_USER"
    echo "  Cron Setup: $SETUP_CRON"
    echo "  Marker File: $CREATE_MARKER"
    echo ""
    log "Next Steps:"
    echo "  1. Verify your GitHub Actions workflow uses the new cache actions"
    echo "  2. Test the cache functionality with a build"
    echo "  3. Monitor cache usage with: du -sh $CACHE_DIR"
    echo "  4. Check cleanup logs in cron or run manually"
    echo ""
    log "Documentation: .github/CACHE_SETUP.md"
}

# 主函数
main() {
    log "Starting GitHub Actions cache setup..."

    parse_args "$@"
    check_permissions
    setup_cache_dir
    create_marker_file
    setup_cleanup_cron
    verify_setup
    show_summary

    success "GitHub Actions cache setup completed successfully!"
}

# 运行主函数
main "$@"