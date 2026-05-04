---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

고성능과 유연성을 갖춘 Go 로깅 라이브러리로, zap을 기반으로 구축되었으며 풍부한 기능과 간단한 API를 제공합니다.

## 📖 언어

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/)
-   [🇰🇷 한국어](README.md) (현재)
-   [🇵🇹 Português](https://lazygophers.github.io/log/pt/)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/)
-   [🇵🇱 Polski](https://lazygophers.github.io/log/pl/)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/)
-   [🇹🇷 Türkçe](https://lazygophers.github.io/log/tr/)

## ✨ 기능

-   **🚀 고성능**：zap을 기반으로 오브젝트 풀링과 조건부 필드 기록
-   **📊 풍부한 로그 레벨**：Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ 유연한 구성**：
    -   로그 레벨 제어
    -   호출자 정보 기록
    -   추적 정보(goroutine ID 포함)
    -   사용자 정의 접두사 및 접미사
    -   사용자 정의 출력 대상(콘솔, 파일 등)
    -   로그 형식 옵션
-   **🔄 파일 로테이션**：시간별 로그 파일 로테이션 지원
-   **🔌 Zap 호환성**：zap WriteSyncer와 원활한 통합
-   **🎯 간단한 API**：표준 로그 라이브러리와 유사한 간단한 API

## 🚀 빠른 시작

### 설치

:::tip 설치
```bash
go get github.com/lazygophers/log
```
:::

### 기본 사용법

```go title="빠른 시작"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // 기본 전역 logger 사용
    log.Debug("디버그 메시지")
    log.Info("정보 메시지")
    log.Warn("경고 메시지")
    log.Error("오류 메시지")

    // 형식화된 출력 사용
    log.Infof("사용자 %s가 로그인했습니다", "admin")

    // 사용자 정의 구성
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("이것은 사용자 정의 logger의 로그입니다")
}
```

## 📚 고급 사용법

### 파일 출력 포함 사용자 정의 Logger

```go title="파일 출력 구성"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 파일 출력 포함 logger 생성
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("호출자 정보가 포함된 디버그 로그")
    logger.Info("추적 정보가 포함된 정보 로그")
}
```

### 로그 레벨 제어

```go title="로그 레벨 제어"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // warn 및 그 이상만 기록됩니다
    logger.Debug("이것은 기록되지 않습니다")  // 무시됨
    logger.Info("이것은 기록되지 않습니다")   // 무시됨
    logger.Warn("이것은 기록됩니다")    // 기록됨
    logger.Error("이것은 기록됩니다")   // 기록됨
}
```

## 🎯 사용 시나리오

### 적용 가능한 시나리오

-   **웹 서비스 및 API 백엔드**：요청 추적, 구조화된 로깅, 성능 모니터링
-   **마이크로서비스 아키텍처**：분산 추적(TraceID), 통합된 로그 형식, 높은 처리량
-   **명령줄 도구**：레벨 제어, 깔끔한 출력, 오류 보고
-   **일괄 처리 작업**：파일 로테이션, 장기 실행, 리소스 최적화

### 특별한 이점

-   **오브젝트 풀 최적화**：Entry 및 Buffer 오브젝트 재사용, GC 압력 감소
-   **비동기 쓰기**：높은 처리량 시나리오(10000+ 로그/초)에서 차단 없음
-   **TraceID 지원**：분산 시스템 요청 추적, OpenTelemetry와 통합
-   **영설정 시작**：즉시 사용 가능, 점진적 구성

## 🔧 구성 옵션

:::note 구성 옵션
다음 모든 메서드는 체인 호출을 지원하며 사용자 정의 Logger를 구축하기 위해 결합할 수 있습니다.
:::

### Logger 구성

| 메서드                  | 설명                | 기본값      |
| --------------------- | ------------------- | ------------ |
| `SetLevel(level)`       | 최소 로그 레벨 설정     | `DebugLevel` |
| `EnableCaller(enable)`  | 호출자 정보 활성화/비활성화 | `false`      |
| `EnableTrace(enable)`   | 추적 정보 활성화/비활성화  | `false`      |
| `SetCallerDepth(depth)` | 호출자 깊이 설정   | `2`          |
| `SetPrefixMsg(prefix)`  | 로그 접두사 설정  | `""`         |
| `SetSuffixMsg(suffix)`  | 로그 접미사 설정  | `""`         |
| `SetOutput(writers...)` | 출력 대상 설정         | `os.Stdout`  |

### 로그 레벨

| 레벨        | 설명                        |
| ----------- | --------------------------- |
| `TraceLevel` | 가장 상세한, 상세한 추적용        |
| `DebugLevel` | 디버그 정보                  |
| `InfoLevel`  | 일반 정보                    |
| `WarnLevel`  | 경고 메시지                  |
| `ErrorLevel` | 오류 메시지                  |
| `FatalLevel` | 치명적인 오류(os.Exit(1) 호출)|
| `PanicLevel` | 패닉 오류(panic() 호출)      |

## 🏗️ 아키텍처

### 핵심 구성 요소

-   **Logger**：구성 가능한 옵션이 있는 주요 로깅 구조
-   **Entry**：포괄적인 필드 지원이 있는 개별 로그 레코드
-   **Level**：로그 레벨 정의 및 유틸리티 함수
-   **Format**：로그 형식 인터페이스 및 구현

### 성능 최적화

-   **오브젝트 풀링**：메모리 할당을 줄이기 위해 Entry 오브젝트 재사용
-   **조건부 기록**：필요한 경우에만 비용이 많이 드는 필드 기록
-   **빠른 레벨 확인**：가장 바깥쪽 레이어에서 로그 레벨 확인
-   **잠금 없는 설계**：대부분의 작업에 잠금 필요 없음

## 📊 성능 비교

:::info 성능 비교
다음 데이터는 벤치마크를 기반으로 합니다. 실제 성능은 환경과 구성에 따라 다를 수 있습니다.
:::

| 기능          | lazygophers/log | zap    | logrus | 표준 로그 |
| ------------- | --------------- | ------ | ------ | -------- |
| 성능      | 높음              | 높음     | 중간     | 낮음       |
| API 간단성    | 높음              | 중간     | 높음     | 높음       |
| 기능 풍부함    | 중간          | 높음     | 높음     | 낮음       |
| 유연성      | 중간          | 높음     | 높음     | 낮음       |
| 학습 곡선      | 낮음              | 중간     | 중간     | 낮음       |

## ❓ 자주 묻는 질문

### 적절한 로그 레벨을 선택하는 방법은?

-   **개발 환경**：상세한 정보를 얻기 위해 `DebugLevel` 또는 `TraceLevel` 사용
-   **프로덕션 환경**：오버헤드를 줄이기 위해 `InfoLevel` 또는 `WarnLevel` 사용
-   **성능 테스트**：모든 로그를 비활성화하기 위해 `PanicLevel` 사용

### 프로덕션 환경에서 성능을 최적화하는 방법은?

:::warning 참고
높은 처리량 시나리오에서는 성능을 최적화하기 위해 비동기 쓰기와 합리적인 로그 레벨을 결합하는 것이 좋습니다.
:::

1. 비동기 쓰기에 `AsyncWriter` 사용：

```go title="비동기 쓰기 구성"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. 불필요한 로그를 피하기 위해 로그 레벨 조정：

```go title="레벨 최적화"
logger.SetLevel(log.InfoLevel)  // Debug 및 Trace 건너뛰기
```

3. 오버헤드를 줄이기 위해 조건부 로그 사용：

```go title="조건부 로그"
if logger.Level >= log.DebugLevel {
    logger.Debug("상세한 디버그 정보")
}
```

### `Caller`와 `EnableCaller`의 차이점은?

-   **`EnableCaller(enable bool)`**：Logger가 호출자 정보를 수집하는지 제어
    -   `EnableCaller(true)`는 호출자 정보 수집을 활성화
-   **`Caller(disable bool)`**：Formatter가 호출자 정보를 출력하는지 제어
    -   `Caller(true)`는 호출자 정보 출력을 비활성화

전역 제어에는 `EnableCaller`를 사용하는 것이 좋습니다.

### 사용자 정의 포맷터를 구현하는 방법은?

`Format` 인터페이스를 구현합니다：

```go title="사용자 정의 포맷터"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 관련 문서

-   [📚 API 문서](API.md) - 전체 API 참조
-   [🤝 기여 가이드](/ko/CONTRIBUTING) - 기여 방법
-   [📋 변경 로그](/ko/CHANGELOG) - 버전 기록
-   [🔒 보안 정책](/ko/SECURITY) - 보안 가이드라인
-   [📜 행동 강령](/ko/CODE_OF_CONDUCT) - 커뮤니티 가이드라인

## 🚀 도움말 얻기

-   **GitHub Issues**：[버그 보고 또는 기능 요청](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[API 문서](https://pkg.go.dev/github.com/lazygophers/log)
-   **예제**：[사용 예제](https://github.com/lazygophers/log/tree/main/examples)

## 📄 라이선스

이 프로젝트는 MIT 라이선스 하에 라이선스가 부여됩니다 - [LICENSE](/ko/LICENSE) 파일을 참조하세요.

## 🤝 기여

기여를 환영합니다! [기여 가이드](/ko/CONTRIBUTING)를 참조하세요.

---

**lazygophers/log**는 성능과 간결성을 중시하는 Go 개발자를 위한 최고의 로깅 솔루션입니다. 작은 유틸리티를 구축하든 대규모 분산 시스템을 구축하든 이 라이브러리는 기능과 사용성의 균형을 제공합니다.
