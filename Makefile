# Makefile for testing log package coverage under different build tags
# 用于测试不同构建标签条件下的日志包覆盖率

.PHONY: test test-all coverage coverage-all coverage-debug coverage-release coverage-discard coverage-debug-discard coverage-release-discard coverage-other clean help

# 默认目标：运行所有覆盖率测试
all: coverage-all

# 基本测试命令
test:
	go test ./...

test-verbose:
	go test -v ./...

# 清理覆盖率文件
clean:
	rm -f coverage-*.out coverage-*.html

# 帮助信息
help:
	@echo "Available targets:"
	@echo "  test                    - Run basic tests"
	@echo "  test-verbose           - Run tests with verbose output" 
	@echo "  coverage-all           - Run coverage tests for all build tag combinations"
	@echo "  coverage-other         - Coverage for default (no build tags)"
	@echo "  coverage-debug         - Coverage for debug build tag"
	@echo "  coverage-release       - Coverage for release build tag"
	@echo "  coverage-discard       - Coverage for discard build tag"
	@echo "  coverage-debug-discard - Coverage for debug+discard build tags"
	@echo "  coverage-release-discard - Coverage for release+discard build tags"
	@echo "  coverage-html          - Generate HTML coverage reports for all"
	@echo "  coverage-summary       - Show coverage summary for all build tags"
	@echo "  clean                  - Remove coverage files"
	@echo "  help                   - Show this help"

# 默认情况（无构建标签）的覆盖率测试
coverage-other:
	@echo "=== Testing coverage for default build (no tags) ==="
	go test ./... -cover -coverprofile=coverage-other.out -tags=""
	@echo "Coverage for default build:"
	@go tool cover -func=coverage-other.out | tail -1

# debug 构建标签的覆盖率测试
coverage-debug:
	@echo "=== Testing coverage for debug build tag ==="
	go test ./... -cover -coverprofile=coverage-debug.out -tags="debug"
	@echo "Coverage for debug build:"
	@go tool cover -func=coverage-debug.out | tail -1

# release 构建标签的覆盖率测试
coverage-release:
	@echo "=== Testing coverage for release build tag ==="
	go test ./... -cover -coverprofile=coverage-release.out -tags="release"
	@echo "Coverage for release build:"
	@go tool cover -func=coverage-release.out | tail -1

# discard 构建标签的覆盖率测试
coverage-discard:
	@echo "=== Testing coverage for discard build tag ==="
	go test ./... -cover -coverprofile=coverage-discard.out -tags="discard"
	@echo "Coverage for discard build:"
	@go tool cover -func=coverage-discard.out | tail -1

# debug+discard 构建标签组合的覆盖率测试
coverage-debug-discard:
	@echo "=== Testing coverage for debug+discard build tags ==="
	go test ./... -cover -coverprofile=coverage-debug-discard.out -tags="debug,discard"
	@echo "Coverage for debug+discard build:"
	@go tool cover -func=coverage-debug-discard.out | tail -1

# release+discard 构建标签组合的覆盖率测试
coverage-release-discard:
	@echo "=== Testing coverage for release+discard build tags ==="
	go test ./... -cover -coverprofile=coverage-release-discard.out -tags="release,discard"
	@echo "Coverage for release+discard build:"
	@go tool cover -func=coverage-release-discard.out | tail -1

# 运行所有构建标签组合的覆盖率测试
coverage-all: coverage-other coverage-debug coverage-release coverage-discard coverage-debug-discard coverage-release-discard
	@echo ""
	@echo "=== Coverage Summary for All Build Tags ==="
	@echo "Default (no tags):"
	@go tool cover -func=coverage-other.out | tail -1 | awk '{print "  " $$3}'
	@echo "Debug:"
	@go tool cover -func=coverage-debug.out | tail -1 | awk '{print "  " $$3}'
	@echo "Release:"
	@go tool cover -func=coverage-release.out | tail -1 | awk '{print "  " $$3}'
	@echo "Discard:"
	@go tool cover -func=coverage-discard.out | tail -1 | awk '{print "  " $$3}'
	@echo "Debug + Discard:"
	@go tool cover -func=coverage-debug-discard.out | tail -1 | awk '{print "  " $$3}'
	@echo "Release + Discard:"
	@go tool cover -func=coverage-release-discard.out | tail -1 | awk '{print "  " $$3}'

# 生成HTML覆盖率报告
coverage-html: coverage-all
	@echo "=== Generating HTML coverage reports ==="
	go tool cover -html=coverage-other.out -o coverage-other.html
	go tool cover -html=coverage-debug.out -o coverage-debug.html
	go tool cover -html=coverage-release.out -o coverage-release.html
	go tool cover -html=coverage-discard.out -o coverage-discard.html
	go tool cover -html=coverage-debug-discard.out -o coverage-debug-discard.html
	go tool cover -html=coverage-release-discard.out -o coverage-release-discard.html
	@echo "HTML reports generated:"
	@echo "  coverage-other.html"
	@echo "  coverage-debug.html" 
	@echo "  coverage-release.html"
	@echo "  coverage-discard.html"
	@echo "  coverage-debug-discard.html"
	@echo "  coverage-release-discard.html"

# 显示详细的覆盖率摘要
coverage-summary: coverage-all
	@echo ""
	@echo "=== Detailed Coverage Analysis ==="
	@echo ""
	@echo "Default Build (no tags):"
	@go tool cover -func=coverage-other.out | grep -E "(Total|github.com/lazygophers/log/)" | head -20
	@echo ""
	@echo "Debug Build:"
	@go tool cover -func=coverage-debug.out | grep -E "(Total|github.com/lazygophers/log/)" | head -20
	@echo ""
	@echo "Release Build:"
	@go tool cover -func=coverage-release.out | grep -E "(Total|github.com/lazygophers/log/)" | head -20
	@echo ""
	@echo "Discard Build:"
	@go tool cover -func=coverage-discard.out | grep -E "(Total|github.com/lazygophers/log/)" | head -20

# 快速测试（不生成详细报告）
test-quick:
	@echo "=== Quick test across all build tags ==="
	@echo -n "Default: "
	@go test ./... -tags="" >/dev/null 2>&1 && echo "PASS" || echo "FAIL"
	@echo -n "Debug: "
	@go test ./... -tags="debug" >/dev/null 2>&1 && echo "PASS" || echo "FAIL"
	@echo -n "Release: "
	@go test ./... -tags="release" >/dev/null 2>&1 && echo "PASS" || echo "FAIL"
	@echo -n "Discard: "
	@go test ./... -tags="discard" >/dev/null 2>&1 && echo "PASS" || echo "FAIL"
	@echo -n "Debug+Discard: "
	@go test ./... -tags="debug,discard" >/dev/null 2>&1 && echo "PASS" || echo "FAIL"
	@echo -n "Release+Discard: "
	@go test ./... -tags="release,discard" >/dev/null 2>&1 && echo "PASS" || echo "FAIL"

# 基准测试
benchmark:
	go test -bench=. -benchmem ./...

benchmark-debug:
	go test -bench=. -benchmem -tags="debug" ./...

benchmark-release:
	go test -bench=. -benchmem -tags="release" ./...

benchmark-discard:
	go test -bench=. -benchmem -tags="discard" ./...

# CI/CD 友好的目标
ci: test-quick coverage-all
	@echo "All CI tests completed successfully"

# 开发者常用目标
dev: test coverage-other
	@echo "Development testing completed"