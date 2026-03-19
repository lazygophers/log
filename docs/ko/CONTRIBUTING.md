---
titleSuffix: ' | LazyGophers Log'
---
# 🤝 LazyGophers Log에 기여하기

여러분의 기여를 매우 환영합니다! LazyGophers Log에 기여하는 것을 가능한 한 간단하고 투명하게 만들고자 합니다. 다음과 같은 경우:

-   🐛 버그 보고
-   💬 코드의 현재 상태에 대한 토론
-   ✨ 기능 요청 제출
-   🔧 수정 제안
-   🚀 새로운 기능 구현

## 📋 목차

-   [행동 강령](#-행동-강령)
-   [개발 프로세스](#-개발-프로세스)
-   [시작하기](#-시작하기)
-   [풀 리퀘스트 프로세스](#-풀-리퀘스트-프로세스)
-   [코딩 표준](#-코딩-표준)
-   [테스트 가이드라인](#-테스트-가이드라인)
-   [빌드 태그 요구사항](#️-빌드-태그-요구사항)
-   [문서](#-문서)
-   [이슈 가이드라인](#-이슈-가이드라인)
-   [성능 고려사항](#-성능-고려사항)
-   [보안 가이드라인](#-보안-가이드라인)
-   [커뮤니티](#-커뮤니티)

## 📜 행동 강령

이 프로젝트와 모든 참가자는 [행동 강령](/ko/CODE_OF_CONDUCT)의 적용을 받습니다. 참여함으로써 규칙을 준수하는 데 동의하게 됩니다.

## 🔄 개발 프로세스

우리는 코드를 호스팅하고 이슈와 기능 요청을 추적하며 풀 리퀘스트를 받기 위해 GitHub를 사용합니다.

### 워크플로우

:::note 개발 프로세스 개요
1. 저장소를 **Fork**합니다
2. fork를 로컬로 **Clone**합니다
3. `master` 브랜치에서 **생성**하는 기능 브랜치를 만듭니다
4. **변경**을 수행합니다
5. 모든 빌드 태그에서 **테스트**합니다
6. **풀 리퀘스트**를 제출합니다
:::

## 🚀 시작하기

### 전제 조건

-   **Go 1.21+** - [Go 설치](https://golang.org/doc/install)
-   **Git** - [Git 설치](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
-   **Make** (선택 사항이지만 권장됨)

### 로컬 개발 설정

```bash title="저장소 복제 및 개발 환경 설정"
# 1. GitHub에서 저장소를 Fork합니다
# 2. fork를 Clone합니다
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. 업스트림 원격 저장소를 추가합니다
git remote add upstream https://github.com/lazygophers/log.git

# 4. 의존성을 설치합니다
go mod tidy

# 5. 설치를 확인합니다
make test-quick
```

### 환경 설정

:::info 환경 구성
최상의 개발 경험을 위해 Go 환경 변수가 올바르게 구성되어 있고 권장 개발 도구가 설치되어 있는지 확인하십시오.
:::

```bash title="환경 설정"
# Go 환경 설정(아직 설정하지 않은 경우)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# 선택 사항: 유용한 도구 설치
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 풀 리퀘스트 프로세스

### 제출 전

1. **검색** 중복을 피하기 위해 기존 PR을 검색합니다
2. **테스트** 모든 빌드 구성에서 변경 사항을 테스트합니다
3. **문서화** 모든 breaking change를 문서화합니다
4. **업데이트** 관련 문서를 업데이트합니다
5. **추가** 새로운 기능에 대한 테스트를 추가합니다

### PR 체크리스트

:::warning PR 제출 전 다음 모든 항목을 확인하십시오
체크리스트 요구사항을 충족하지 않는 PR은 병합되지 않습니다.
:::

-   [ ] **코드 품질**

    -   [ ] 코드가 프로젝트 스타일 가이드를 따릅니다
    -   [ ] 새로운 lint 경고가 없습니다
    -   [ ] 올바른 오류 처리
    -   [ ] 효율적인 알고리즘 및 데이터 구조

-   [ ] **테스트**

    -   [ ] 모든 기존 테스트 통과: `make test`
    -   [ ] 새로운 기능을 위한 새 테스트 추가
    -   [ ] 테스트覆盖率 유지 또는 향상
    -   [ ] 모든 빌드 태그 테스트 완료: `make test-all`

-   [ ] **문서**

    -   [ ] 코드에 적절한 주석이 있습니다
    -   [ ] API 문서가 업데이트되었습니다(필요한 경우)
    -   [ ] README가 업데이트되었습니다(필요한 경우)
    -   [ ] 다국어 문서가 업데이트되었습니다(사용자용인 경우)

-   [ ] **빌드 호환성**
    -   [ ] 기본 모드: `go build`
    -   [ ] 디버그 모드: `go build -tags debug`
    -   [ ] 릴리스 모드: `go build -tags release`
    -   [ ] discard 모드: `go build -tags discard`
    -   [ ] 조합 모드 테스트 완료

### PR 템플릿

풀 리퀘스트를 제출할 때 [PR 템플릿](https://github.com/lazygophers/log/blob/main/.github/pull_request_template.md)을 사용하십시오.

## 📏 코딩 표준

### Go 스타일 가이드

:::tip Go 코드 규칙
표준 Go 스타일 가이드를 따르며 몇 가지 추가 사항이 있습니다. `go fmt` 및 `goimports` 검사를 통과하는지 확인하십시오.
:::

```go
// ✅ Good
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ Bad
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### 명명 규칙

-   **패키지**: 짧고 소문자, 가능한 단일 단어 사용
-   **함수**: PascalCase, 설명적
-   **변수**: 로컬 변수는 camelCase, 내보내기 변수는 PascalCase
-   **상수**: 내보내기 상수는 PascalCase, 비내보내기 상수는 camelCase
-   **인터페이스**: 일반적으로 "er"로 끝남(예: `Writer`, `Formatter`)

### 코드 조직

```
project/
├── docs/           # 다국어 문서
├── .github/        # GitHub 템플릿 및 워크플로우
├── logger.go       # 주요 로거 구현
├── entry.go        # 로그 항목 구조
├── level.go        # 로그 수준
├── formatter.go    # 로그 포맷팅
├── output.go       # 출력 관리
└── *_test.go      # 소스 코드와 공존하는 테스트
```

### 오류 처리

:::tip 오류 처리 모범 사례
라이브러리 코드는 panic보다는 오류를 반환해야 하며, 호출자가 예외 상황을 처리하도록 합니다.
:::

```go title="오류 처리 예시"
// ✅ 권장: 오류를 반환하고 panic하지 마십시오
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ 피하세요: 라이브러리 코드에서 panic 사용
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // 이렇게 하지 마십시오
    }
    return &Logger{...}
}
```

## 🧪 테스트 가이드라인

### 테스트 구조

```go title="표 기반 테스트 예시"
func TestLogger_Info(t *testing.T) {
    tests := []struct {
        name     string
        level    Level
        message  string
        expected bool
    }{
        {"info level allows info", InfoLevel, "test", true},
        {"warn level blocks info", WarnLevel, "test", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### 커버리지 요구사항

:::warning 커버리지 강력한 요구사항
새 코드의 커버리지가 90% 미만인 PR은 CI 검사를 통과하지 못합니다.
:::

-   **최소 요구사항**: 새 코드 커버리지 90%
-   **목표**: 전체 커버리지 95%+
-   **모든 빌드 태그**가 커버리지를 유지해야 함
-   `make coverage-all`을 사용하여 확인

### 테스트 명령

```bash title="테스트 실행"
# 모든 빌드 태그에서 빠른 테스트
make test-quick

# 커버리지가 포함된 전체 테스트 스위트
make test-all

# 커버리지 보고서
make coverage-html

# 벤치마크
make benchmark
```

## 🏗️ 빌드 태그 요구사항

:::warning 빌드 호환성
모든 변경사항은 빌드 태그 시스템과 호환되어야 하며, 모든 빌드 태그 테스트를 통과하지 않은 코드는 병합되지 않습니다.
:::

### 지원되는 빌드 태그

-   **기본** (`go build`): 완전한 기능
-   **디버그** (`go build -tags debug`): 향상된 디버깅 기능
-   **릴리스** (`go build -tags release`): 프로덕션 최적화
-   **discard** (`go build -tags discard`): 최대 성능

### 빌드 태그 테스트

:::info 빌드 태그 설명
프로젝트는 빌드 태그를 사용하여 조건부 컴파일을 구현하며, 다른 태그는 다른 실행 모드에 해당합니다. 제출 전 모든 태그에서 테스트하십시오.
:::

```bash title="빌드 태그 테스트"
# 각 빌드 구성 테스트
make test-default
make test-debug
make test-release
make test-discard

# 조합 테스트
make test-debug-discard
make test-release-discard

# 모두 한 번에 테스트
make test-all
```

### 빌드 태그 가이드라인

```go
//go:build debug
// +build debug

package log

// 디버그 특정 구현
```

## 📚 문서

### 코드 문서

-   **모든 내보낸 함수**에는 명확한 주석이 있어야 합니다
-   **복잡한 알고리즘**에는 설명이 필요합니다
-   **예시**는 사소하지 않은 사용법에 필요합니다
-   **스레드 안전성** 참고 사항(해당하는 경우)

```go
// SetLevel은 최소 로깅 수준을 설정합니다.
// 이 수준보다 낮은 로그는 무시됩니다.
// 이 메서드는 스레드 안전합니다.
//
// Example:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // 출력되지 않음
//   logger.Info("visible")   // 출력됨
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### README 업데이트

기능을 추가할 때 다음을 업데이트하십시오:

-   주요 README.md
-   `docs/`의 모든 언어별 README
-   코드 예시
-   기능 목록

## 🐛 이슈 가이드라인

### 버그 보고

[버그 보고 템플릿](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/bug_report.md)을 사용하고 다음을 포함하십시오:

-   **명확한 문제 설명**
-   **재현 단계**
-   **예상된 실제 동작**
-   **환경 세부정보**(운영 체제, Go 버전, 빌드 태그)
-   **최소 코드 예시**

### 기능 요청

[기능 요청 템플릿](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/feature_request.md)을 사용하고 다음을 포함하십시오:

-   **명확한 기능 동기**
-   **제안된 API** 설계
-   **구현 고려사항**
-   **breaking change 분석**

### 질문

[질문 템플릿](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/question.md)을 사용합니다:

-   사용 문제
-   구성 도움말
-   모범 사례
-   통합 지침

## 🚀 성능 고려사항

### 벤치마킹

성능에 민감한 변경에 대해서는 항상 벤치마킹을 수행하십시오:

```bash title="벤치마킹 실행"
# 벤치마킹 실행
go test -bench=. -benchmem

# 전후 성능 비교
go test -bench=. -benchmem > before.txt
# 변경 수행
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### 성능 가이드라인

:::tip 성능 최적화 요점
이것은 성능에 민감한 로깅 라이브러리이며, 모든 변경은 핫 경로에 대한 영향을 고려해야 합니다.
:::

-   **최소화** 핫 경로에서의 메모리 할당
-   **객체 풀 사용** 자주 생성되는 객체의 경우
-   **조기 반환** 비활성화된 로그 수준의 경우
-   **반영 피하기** 성능 중요 코드에서
-   **최적화 전에 프로파일링**

### 메모리 관리

```go
// ✅ 권장: 객체 풀 사용
var entryPool = sync.Pool{
    New: func() interface{} {
        return &Entry{}
    },
}

func getEntry() *Entry {
    return entryPool.Get().(*Entry)
}

func putEntry(e *Entry) {
    e.Reset()
    entryPool.Put(e)
}
```

## 🔒 보안 가이드라인

### 민감한 데이터

:::warning 보안 주의사항
로그에 민감한 데이터가 유출되면 심각한 보안 사고가 발생할 수 있으므로 다음 규칙을 반드시 준수하십시오.
:::

-   **절대 기록하지 마십시오** 비밀번호, 토큰 또는 민감한 데이터
-   **정화** 로그 메시지의 사용자 입력
-   **피하십시오** 전체 요청/응답 본문 기록
-   **사용** 더 나은 제어를 위한 구조화된 로깅

```go
// ✅ 권장
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ 피하세요
logger.Infof("User login: %+v", userRequest) // 비밀번호가 포함될 수 있음
```

### 의존성

-   의존성을 **최신으로** 유지
-   **신중하게 검토** 새로운 의존성
-   **최소화** 외부 의존성
-   **사용** `go mod verify` 무결성 확인

## 👥 커뮤니티

### 도움 얻기

-   📖 [문서](README.md)
-   💬 [GitHub 토론](https://github.com/lazygophers/log/discussions)
-   🐛 [이슈 트래커](https://github.com/lazygophers/log/issues)
-   📧 이메일: support@lazygophers.com

### 커뮤니케이션 가이드라인

-   **존중과** 포용성 유지
-   **검색 먼저** 질문하기 전에
-   **도움 요청 시 컨텍스트 제공**
-   **다른 사람을 도울** 수 있을 때 도움
-   **준수** [행동 강령](/ko/CODE_OF_CONDUCT)

## 🎯 인정

기여자는 다음과 같은 방법으로 인정받습니다:

-   **README 기여자** 섹션
-   **릴리스 노트** 언급
-   **GitHub 기여자** 차트
-   **커뮤니티 감사** 게시물

## 📝 라이선스

기여함으로써 기여가 MIT 라이선스에 따라 라이선스됨에 동의합니다.

---

## 🌍 다국어 문서

이 문서는 여러 언어로 제공됩니다:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CONTRIBUTING.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/CONTRIBUTING.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CONTRIBUTING.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CONTRIBUTING.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CONTRIBUTING.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CONTRIBUTING.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CONTRIBUTING.md)
-   [🇰🇷 한국어](/ko/CONTRIBUTING)（현재）

---

**LazyGophers Log에 기여해 주셔서 감사합니다!🚀**
