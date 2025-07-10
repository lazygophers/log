// Package log 提供日志输出功能，支持设置多个输出目标和日志轮转
package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// SetOutput 设置日志输出目标
// 参数 writes: 一个或多个实现 io.Writer 接口的输出目标
// 返回: 日志记录器实例，支持链式调用
func SetOutput(writes ...io.Writer) *Logger {
	return std.SetOutput(writes...)
}

// GetOutputWriter 创建基于文件的日志输出器
// 参数 filename: 日志文件路径
// 返回: 实现 io.Writer 接口的日志输出器
func GetOutputWriter(filename string) io.Writer {
	if filepath.Dir(filename) != filename && !isDir(filepath.Dir(filename)) {
		err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			Errorf("err:%v", err)
		}
	}

	hook, err := rotatelogs.New(filename)
	if err != nil {
		std.Panicf("err:%v", err)
	}
	return hook
}

// isDir 检查给定路径是否为目录
// 参数 path: 要检查的路径
// 返回: true表示是目录，false表示不是目录或路径不存在
func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Output 接口定义日志输出行为
type Output interface {
	io.Writer
}

var (
	// cleanRotatelogOnce 确保每个日志文件只启动一次清理协程
	cleanRotatelogOnce = make(map[string]bool)
)

// GetOutputWriterHourly 创建按小时轮转的日志输出器
// 参数 filename: 日志文件基础路径
// 返回: 实现 Writer 接口的日志输出器
//
// 特性:
// - 每小时生成新日志文件
// - 保留最近12小时的日志
// - 自动创建日志目录
// - 后台协程定期清理旧日志
func GetOutputWriterHourly(filename string) Writer {
	if filepath.Dir(filename) != filename && !isDir(filepath.Dir(filename)) {
		err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			Errorf("err:%v", err)
		}
	}

	hook, err := rotatelogs.
		New(filename+"%Y%m%d%H.log",
			rotatelogs.WithLinkName(filename+".log"),
			rotatelogs.WithRotationSize(1024*1024*8*100),
			rotatelogs.WithRotationTime(time.Hour),
			rotatelogs.WithRotationCount(12),
		)
	if err != nil {
		std.Panicf("err:%v", err)
	}

	if _, ok := cleanRotatelogOnce[filename]; !ok {
		go func() {
			for {
				files, err := os.ReadDir(filepath.Dir(filename))
				if err != nil {
					fmt.Printf("err:%v\n", err)
					continue
				}

				filesOnly := make([]string, 0, len(files))
				for _, file := range files {
					filesOnly = append(filesOnly, file.Name())
				}

				sort.Slice(filesOnly, func(i, j int) bool {
					return filesOnly[i] > filesOnly[j]
				})

				for i, s := range filesOnly {
					if i < 12 {
						continue
					}
					if s == ".log" {
						continue
					}

					fmt.Printf("remove:%s\n", s)
					err = os.Remove(filepath.Join(filepath.Dir(filename), s))
					if err != nil {
						Errorf("err:%v", err)
					}
				}

				time.Sleep(time.Minute * 10)
			}
		}()
		cleanRotatelogOnce[filename] = true
	}

	return hook
}
