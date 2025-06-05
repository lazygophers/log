package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/elliotchance/pie/v2"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func SetOutput(writes ...io.Writer) *Logger {
	return std.SetOutput(writes...)
}

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

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

var (
	cleanRotatelogOnce = make(map[string]bool)
)

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

				pie.Each(
					pie.DropTop(
						pie.SortUsing(
							pie.Map(
								files,
								func(file os.DirEntry) string {
									return file.Name()
								},
							),
							func(a, b string) bool {
								return a > b
							},
						),
						12,
					),
					func(s string) {
						if s == ".log" {
							return
						}

						fmt.Printf("remove:%s\n", s)
						err = os.Remove(filepath.Join(filepath.Dir(filename), s))
						if err != nil {
							Errorf("err:%v", err)
						}
					})

				time.Sleep(time.Minute * 10)
			}
		}()
		cleanRotatelogOnce[filename] = true
	}

	return hook
}
