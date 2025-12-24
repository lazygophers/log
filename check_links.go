package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var linkRegex = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

func main() {
	// 检查当前目录下所有.md文件
	var brokenLinks []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过非.md文件和.git目录
		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		// 检查文件中的链接
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 0

		for scanner.Scan() {
			lineNum++
			line := scanner.Text()

			// 查找所有链接
			matches := linkRegex.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				if len(match) < 3 {
					continue
				}

				link := match[2]
				// 跳过外部链接
				if strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://") {
					continue
				}

				// 跳过锚点链接
				if strings.HasPrefix(link, "#") {
					continue
				}

				// 检查相对路径
				if !filepath.IsAbs(link) {
					// 获取文件所在目录
					fileDir := filepath.Dir(path)
					fullPath := filepath.Join(fileDir, link)

					// 处理../路径
					fullPath, err = filepath.Abs(fullPath)
					if err != nil {
						brokenLinks = append(brokenLinks, fmt.Sprintf("%s:%d: %s (无法解析路径)", path, lineNum, link))
						continue
					}

					// 检查文件是否存在
					if _, err := os.Stat(fullPath); os.IsNotExist(err) {
						// 尝试检查是否是目录（如果链接指向目录）
						if _, err := os.Stat(fullPath + "/"); os.IsNotExist(err) {
							brokenLinks = append(brokenLinks, fmt.Sprintf("%s:%d: %s (文件不存在)", path, lineNum, link))
						}
					}
				}
			}
		}

		return scanner.Err()
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		return
	}

	if len(brokenLinks) == 0 {
		fmt.Println("All links are valid!")
	} else {
		fmt.Println("Found broken links:")
		for _, link := range brokenLinks {
			fmt.Println("  -", link)
		}
	}
}
