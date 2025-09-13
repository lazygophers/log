package log

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// HourlyRotator 实现按小时轮转的日志写入器
type HourlyRotator struct {
	mu           sync.Mutex
	filename     string
	linkName     string
	currentFile  *os.File
	currentHour  string
	maxSize      int64 // 最大文件大小（字节）
	maxFiles     int   // 最大保留文件数
	cleanupOnce  sync.Once
}

// NewHourlyRotator 创建一个新的按小时轮转的日志写入器
func NewHourlyRotator(filename string, maxSize int64, maxFiles int) *HourlyRotator {
	r := &HourlyRotator{
		filename: filename,
		linkName: filename + ".log",
		maxSize:  maxSize,
		maxFiles: maxFiles,
	}
	
	// 启动清理协程
	r.cleanupOnce.Do(func() {
		go r.cleanup()
	})
	
	return r
}

// Write 实现 io.Writer 接口
func (r *HourlyRotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if err := r.rotate(); err != nil {
		return 0, err
	}
	
	if r.currentFile == nil {
		return 0, fmt.Errorf("no current file")
	}
	
	return r.currentFile.Write(p)
}

// rotate 检查是否需要轮转并执行轮转
func (r *HourlyRotator) rotate() error {
	now := time.Now()
	hour := now.Format("2006010215")
	
	// 检查是否需要轮转（小时变化或文件大小超限）
	needRotate := r.currentHour != hour
	if r.currentFile != nil && !needRotate {
		// 检查文件大小
		if stat, err := r.currentFile.Stat(); err == nil && stat.Size() >= r.maxSize {
			needRotate = true
		}
	}
	
	if needRotate || r.currentFile == nil {
		return r.doRotate(hour)
	}
	
	return nil
}

// doRotate 执行实际的轮转操作
func (r *HourlyRotator) doRotate(hour string) error {
	// 关闭当前文件
	if r.currentFile != nil {
		r.currentFile.Close()
	}
	
	// 确保目录存在
	ensureDir(filepath.Dir(r.filename))
	
	// 生成新的文件名
	newFilename := r.filename + hour + ".log"
	
	// 打开新文件
	file, err := os.OpenFile(newFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	
	r.currentFile = file
	r.currentHour = hour
	
	// 更新软链接
	r.updateLink(newFilename)
	
	return nil
}

// updateLink 更新指向最新日志文件的软链接
func (r *HourlyRotator) updateLink(target string) {
	// 删除旧链接
	os.Remove(r.linkName)
	
	// 创建新链接（忽略错误，软链接创建失败不应影响日志写入）
	os.Symlink(filepath.Base(target), r.linkName)
}

// cleanup 定期清理旧的日志文件
func (r *HourlyRotator) cleanup() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		r.cleanupOldFiles()
	}
}

// cleanupOldFiles 清理过期的日志文件
func (r *HourlyRotator) cleanupOldFiles() {
	dir := filepath.Dir(r.filename)
	base := filepath.Base(r.filename)
	
	// 读取目录中的所有文件
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("读取日志目录 %s 用于清理时出错: %v\n", dir, err)
		return
	}
	
	// 筛选出日志文件
	var logFiles []string
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			// 匹配格式：base + YYYYMMDDHH + .log
			if strings.HasPrefix(name, base) && strings.HasSuffix(name, ".log") && len(name) == len(base)+14 {
				logFiles = append(logFiles, name)
			}
		}
	}
	
	// 按文件名降序排序（最新的在前）
	sort.Slice(logFiles, func(i, j int) bool {
		return logFiles[i] > logFiles[j]
	})
	
	// 删除超过保留数量的旧文件
	for i, filename := range logFiles {
		if i < r.maxFiles {
			continue
		}
		
		fullPath := filepath.Join(dir, filename)
		fmt.Printf("正在删除旧的日志文件: %s\n", fullPath)
		if err := os.Remove(fullPath); err != nil {
			fmt.Printf("删除旧日志文件 %s 失败: %v\n", fullPath, err)
		}
	}
}

// Sync 同步当前文件的内容到磁盘
func (r *HourlyRotator) Sync() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if r.currentFile != nil {
		return r.currentFile.Sync()
	}
	
	return nil
}

// Close 关闭轮转器并清理资源
func (r *HourlyRotator) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if r.currentFile != nil {
		return r.currentFile.Close()
	}
	
	return nil
}