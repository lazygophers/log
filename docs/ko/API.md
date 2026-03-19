---
titleSuffix: ' | LazyGophers Log'
---
# 📚 API 문서

## 개요

LazyGophers Log는 다중 로그 수준, 사용자 정의 포맷팅, 비동기 쓰기 및 빌드 태그 최적화를 지원하는 포괄적인 로깅 API를 제공합니다. 이 문서는 모든 공개 API, 구성 옵션 및 사용 패턴을 다룹니다.

## 목차

-   [핵심 타입](#핵심-타입)
-   [Logger API](#logger-api)
-   [전역 함수](#전역-함수)
-   [로그 수준](#로그-수준)
-   [포매터](#포매터)
-   [출력 라이터](#출력-라이터)
-   [컨텍스트 로깅](#컨텍스트-로깅)
-   [빌드 태그](#빌드-태그)
-   [성능 최적화](#성능-최적화)
-   [예제](#예제)

## 핵심 타입

### Logger

모든 로깅 기능을 제공하는 주요 로거 구조체입니다.

```go
type Logger struct {
    // 스레드 안전한 작업을 위한 비공개 필드 포함
}
```

#### 생성자

```go
func New() *Logger
```

기본 구성으로 새 로거 인스턴스를 생성합니다:

-   수준: `DebugLevel`
-   출력: `os.Stdout`
-   포매터: 기본 텍스트 포매터
-   호출자 추적: 비활성화

**예제:**

```go title="로거 생성"
logger := log.New()
logger.Info("새 로거가 생성되었습니다")
```

### Entry

모든 연관된 메타데이터가 있는 단일 로그 항목을 나타냅니다.

```go
type Entry struct {
    Time       time.Time     // 항목 생성 시간의 타임스탬프
    Level      Level         // 로그 수준
    Message    string        // 로그 메시지
    Pid        int          // 프로세스 ID
    Gid        uint64       // 고루틴 ID
    TraceID    string       // 분산 추적을 위한 추적 ID
    CallerName string       // 호출자 함수 이름
    CallerFile string       // 호출자 파일 경로
    CallerLine int          // 호출자 라인 번호
}
```

## Logger API

### 구성 메서드

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

최저 로그 수준을 설정합니다. 이 수준보다 낮은 메시지는 무시됩니다.

**매개변수:**

-   `level`: 처리할 최저 로그 수준

**반환값:**

-   `*Logger`: 메서드 체이닝을 지원하기 위해 자신을 반환

**예제:**

```go title="로그 수준 설정"
logger.SetLevel(log.InfoLevel)
logger.Debug("이것은 표시되지 않습니다")  // 무시됨
logger.Info("이것은 표시됩니다")        // 처리됨
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

로그 메시지의 하나 이상의 출력 대상을 설정합니다.

**매개변수:**

-   `writers`: 하나 이상의 `io.Writer` 출력 대상

**반환값:**

-   `*Logger`: 메서드 체이닝을 지원하기 위해 자신을 반환

**예제:**

```go title="출력 대상 설정"
// 단일 출력
logger.SetOutput(os.Stdout)

// 여러 출력
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

로그 출력의 사용자 정의 포매터를 설정합니다.

**매개변수:**

-   `formatter`: `Format` 인터페이스를 구현하는 포매터

**반환값:**

-   `*Logger`: 메서드 체이닝을 지원하기 위해 자신을 반환

**예제:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

로그 항목에서 호출자 정보 기록을 활성화하거나 비활성화합니다.

**매개변수:**

-   `enable`: 호출자 정보를 활성화할지 여부 (`true` 전달 시 활성화)

**반환값:**

-   `*Logger`: 메서드 체이닝을 지원하기 위해 자신을 반환

**예제:**

```go
logger.EnableCaller(true)
logger.Info("이것은 파일:라인 번호 정보를 포함합니다")

logger.EnableCaller(false)
logger.Info("이것은 파일:라인 번호 정보를 포함하지 않습니다")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

포매터에서 호출자 정보를 제어합니다.

**매개변수:**

-   `disable`: 호출자 정보를 비활성화할지 여부 (`true` 전달 시 비활성화)

**반환값:**

-   `*Logger`: 메서드 체이닝을 지원하기 위해 자신을 반환

**예제:**

```go
logger.Caller(false)  // 비활성화하지 않음, 호출자 정보 표시
logger.Info("이것은 파일:라인 번호 정보를 포함합니다")

logger.Caller(true)   // 호출자 정보 비활성화
logger.Info("이것은 파일:라인 번호 정보를 포함하지 않습니다")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

로거를 래핑할 때 호출자 정보의 스택 깊이를 설정합니다.

**매개변수:**

-   `depth`: 건너뛸 스택 프레임 수

**반환값:**

-   `*Logger`: 메서드 체이닝을 지원하기 위해 자신을 반환

**예제:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // 래퍼 함수 건너뛰기
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

모든 로그 메시지의 접두사 또는 접미사 텍스트를 설정합니다.

**매개변수:**

-   `prefix/suffix`: 메시지 앞/뒤에 추가할 텍스트

**반환값:**

-   `*Logger`: 메서드 체이닝을 지원하기 위해 자신을 반환

**예제:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // 출력: [APP] Hello [END]
```

### 로깅 메서드

모든 로깅 메서드는 두 가지 변형이 있습니다: 간단한 버전과 포맷팅된 버전.

#### Trace 수준

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

trace 수준에서 로그를 기록합니다 (가장 상세함).

**예제:**

```go
logger.Trace("상세한 실행 추적")
logger.Tracef("%d 항목 처리 중, 총 %d 항목", i, total)
```

#### Debug 수준

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

debug 수준에서 개발 정보를 기록합니다.

**예제:**

```go
logger.Debug("변수 상태:", variable)
logger.Debugf("사용자 %s 인증 성공", username)
```

#### Info 수준

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

정보성 메시지를 기록합니다.

**예제:**

```go
logger.Info("애플리케이션이 시작되었습니다")
logger.Infof("서버가 포트 %d에서 수신 중입니다", port)
```

#### Warn 수준

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

잠재적인 문제 상황에 대한 경고 메시지를 기록합니다.

**예제:**

```go
logger.Warn("사용되지 않는 함수가 호출되었습니다")
logger.Warnf("메모리 사용량이 높음: %d%%", memoryPercent)
```

#### Error 수준

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

오류 메시지를 기록합니다.

**예제:**

```go
logger.Error("데이터베이스 연결 실패")
logger.Errorf("요청 처리 실패: %v", err)
```

#### Fatal 수준

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

치명적인 오류를 기록하고 `os.Exit(1)`을 호출합니다.

:::danger 파괴적 작업
`Fatal`과 `Fatalf`은 로깅 후 즉시 `os.Exit(1)`을 호출하여 프로세스를 종료합니다. 회복할 수 없는 오류 상황에서만 사용하세요. `defer` 문은 실행되지 **않습니다**.
:::

**예제:**

```go
logger.Fatal("치명적인 시스템 오류")
logger.Fatalf("서버를 시작할 수 없음: %v", err)
```

#### Panic 수준

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

오류 메시지를 기록하고 `panic()`을 호출합니다.

:::danger 파괴적 작업
`Panic`과 `Panicf`는 로깅 후 `panic()`을 호출합니다. `Fatal`과 달리 `panic`은 `recover()`로捕获할 수 있지만,捕获되지 않으면 프로그램을 종료합니다.
:::

**예제:**

```go
logger.Panic("회복할 수 없는 오류가 발생했습니다")
logger.Panicf("잘못된 상태: %v", state)
```

### 유틸리티 메서드

#### Clone

```go
func (l *Logger) Clone() *Logger
```

동일한 구성으로 로거 복사본을 생성합니다.

**반환값:**

-   `*Logger`: 복사된 설정을 가진 새 로거 인스턴스

**예제:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

첫 번째 매개변수로 `context.Context`를 받는 컨텍스트 인식 로거를 생성합니다.

**반환값:**

-   `LoggerWithCtx`: 컨텍스트 인식 로거 인스턴스

**예제:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "컨텍스트 인식 메시지")
```

## 전역 함수

기본 전역 로거를 사용하는 패키지 수준 함수입니다.

```go
func SetLevel(level Level)
func SetOutput(writers ...io.Writer)
func SetFormatter(formatter Format)
func Caller(disable bool)

func Trace(v ...any)
func Tracef(format string, v ...any)
func Debug(v ...any)
func Debugf(format string, v ...any)
func Info(v ...any)
func Infof(format string, v ...any)
func Warn(v ...any)
func Warnf(format string, v ...any)
func Error(v ...any)
func Errorf(format string, v ...any)
func Fatal(v ...any)
func Fatalf(format string, v ...any)
func Panic(v ...any)
func Panicf(format string, v ...any)
```

**예제:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("전역 로거 사용")
```

## 로그 수준

### Level 타입

```go
type Level int8
```

### 사용 가능한 수준

```go
const (
    PanicLevel Level = iota  // 0 - 패닉 후 종료
    FatalLevel              // 1 - 치명적 오류 후 종료
    ErrorLevel              // 2 - 오류 조건
    WarnLevel               // 3 - 경고 조건
    InfoLevel               // 4 - 정보성 메시지
    DebugLevel              // 5 - 디버그 메시지
    TraceLevel              // 6 - 가장 상세한 추적
)
```

### Level 메서드

```go
func (l Level) String() string
```

수준의 문자열 표현을 반환합니다.

**예제:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## 포매터

### Format 인터페이스

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

사용자 정의 포매터는 이 인터페이스를 구현해야 합니다.

### 기본 포매터

사용자 정의 가능한 옵션이 있는 내장 텍스트 포매터입니다.

```go
type Formatter struct {
    // 구성 옵션
}
```

### JSON 포매터 예제

```go
type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "caller":    fmt.Sprintf("%s:%d", entry.CallerFile, entry.CallerLine),
    }
    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }

    jsonData, _ := json.Marshal(data)
    return append(jsonData, '\n')
}

// 사용
logger.SetFormatter(&JSONFormatter{})
```

## 출력 라이터

### 파일 출력 및 로테이션

```go
func GetOutputWriterHourly(filename string) io.Writer
```

매시간 로그 파일을 로테이션하는 라이터를 생성합니다.

**매개변수:**

-   `filename`: 로그 파일의 기본 파일 이름

**반환값:**

-   `io.Writer`: 로테이션 파일 라이터

**예제:**

```go title="시간별 로그 로테이션"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// 다음과 같은 파일 생성: app-2024010115.log, app-2024010116.log 등
```

### 비동기 라이터

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

고성능 로깅을 위한 비동기 라이터를 생성합니다.

**매개변수:**

-   `writer`: 기본 라이터
-   `bufferSize`: 내부 버퍼 크기

**반환값:**

-   `*AsyncWriter`: 비동기 라이터 인스턴스

**메서드:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**예제:**

```go title="비동기 라이터"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## 컨텍스트 로깅

### LoggerWithCtx 인터페이스

```go
type LoggerWithCtx interface {
    Trace(ctx context.Context, v ...any)
    Tracef(ctx context.Context, format string, v ...any)
    Debug(ctx context.Context, v ...any)
    Debugf(ctx context.Context, format string, v ...any)
    Info(ctx context.Context, v ...any)
    Infof(ctx context.Context, format string, v ...any)
    Warn(ctx context.Context, v ...any)
    Warnf(ctx context.Context, format string, v ...any)
    Error(ctx context.Context, v ...any)
    Errorf(ctx context.Context, format string, v ...any)
    Fatal(ctx context.Context, v ...any)
    Fatalf(ctx context.Context, format string, v ...any)
    Panic(ctx context.Context, v ...any)
    Panicf(ctx context.Context, format string, v ...any)
}
```

### 컨텍스트 함수

```go
func SetTrace(traceID string)
func GetTrace() string
```

현재 고루틴의 추적 ID를 설정하고 가져옵니다.

**예제:**

```go
log.SetTrace("trace-123-456")
log.Info("이 메시지는 추적 ID를 포함합니다")

traceID := log.GetTrace()
fmt.Println("현재 추적 ID:", traceID)
```

## 빌드 태그

이 라이브러리는 빌드 태그를 통한 조건부 컴파일을 지원합니다:

:::info 빌드 태그 설명
빌드 태그는 `go build -tags` 매개변수로 지정되며, 다른 태그는 로그 라이브러리의 컴파일 동작과 런타임 특성을 변경합니다. 적절한 태그를 선택하면 개발 편의성과 프로덕션 성능 사이의 균형을 맞출 수 있습니다.
:::

### 기본 모드

```bash
go build
```

-   전체 기능 활성화
-   디버그 메시지 포함
-   표준 성능

### 디버그 모드

```bash
go build -tags debug
```

-   강화된 디버그 정보
-   추가 런타임 검사
-   상세한 호출자 정보

### 릴리스 모드

```bash
go build -tags release
```

-   프로덕션 환경 최적화
-   디버그 메시지 비활성화
-   자동 로그 로테이션 활성화

### 폐기 모드

```bash
go build -tags discard
```

-   최대 성능
-   모든 로깅 작업은 무작업
-   제로 오버헤드

### 조합 모드

```bash
go build -tags "debug,discard"    # 디버그 및 폐기
go build -tags "release,discard"  # 릴리스 및 폐기
```

## 성능 최적화

:::tip 성능 모범 사례
이 라이브러리는 객체 풀, 수준 검사 전치, 비동기 쓰기 등의 메커니즘을 통해 깊게 성능 최적화되어 있습니다. 고처리량 시나리오에서는 비동기 라이터와 적절한 빌드 태그를 조합하여 사용하면 최적의 성능을 얻을 수 있습니다.
:::

### 객체 풀

이 라이브러리는 내부적으로 `sync.Pool`을 사용하여 다음을 관리합니다:

-   로그 항목 객체
-   바이트 버퍼
-   포매터 버퍼

이는 고처리량 시나리오에서 가비지 수집 압력을 줄입니다.

### 수준 검사

로그 수준 검사는 비싼 작업 전에 수행됩니다:

```go
// 효율적 - 수준이 활성화된 경우에만 메시지 포맷팅
logger.Debugf("비싼 작업 결과: %+v", expensiveCall())

// 프로덕션 환경에서 디버그가 비활성화된 경우 덜 효율적
result := expensiveCall()
logger.Debug("결과:", result)
```

### 비동기 쓰기

고처리량 애플리케이션의 경우:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // 큰 버퍼
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### 빌드 태그 최적화

환경에 따라 적절한 빌드 태그를 사용하세요:

-   개발: 기본 또는 디버그 태그
-   프로덕션: 릴리스 태그
-   성능 중요: 폐기 태그

## 예제

### 기본 사용법

```go title="기본 사용법"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("애플리케이션 시작 중")
    log.Warn("이것은 경고입니다")
    log.Error("이것은 오류입니다")
}
```

### 사용자 정의 로거

```go title="사용자 정의 구성"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // 로거 구성
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  // 호출자 정보 비활성화
    logger.SetPrefixMsg("[MyApp] ")

    // 파일로 출력 설정
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("사용자 정의 로거가 구성되었습니다")
    logger.Debug("호출자가 있는 디버그 정보")
}
```

### 고성능 로깅

```go title="고성능 로깅"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // 로테이션 파일 라이터 생성
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // 비동기 라이터로 성능 향상
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // 프로덕션 환경에서 디버그 로그 건너뛰기

    // 고처리량 로깅
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### 컨텍스트 인식 로깅

```go title="컨텍스트 인식 로깅"
package main

import (
    "context"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()
    ctxLogger := logger.CloneToCtx()

    ctx := context.Background()
    log.SetTrace("trace-123-456")

    ctxLogger.Info(ctx, "사용자 요청 처리")
    ctxLogger.Debug(ctx, "검증 완료")
}
```

### 사용자 정의 JSON 포매터

```go title="사용자 정의 JSON 포매터"
package main

import (
    "encoding/json"
    "os"
    "time"
    "github.com/lazygophers/log"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *log.Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339Nano),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "pid":       entry.Pid,
        "gid":       entry.Gid,
    }

    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }

    if entry.CallerName != "" {
        data["caller"] = map[string]interface{}{
            "function": entry.CallerName,
            "file":     entry.CallerFile,
            "line":     entry.CallerLine,
        }
    }

    jsonData, _ := json.MarshalIndent(data, "", "  ")
    return append(jsonData, '\n')
}

func main() {
    logger := log.New()
    logger.SetFormatter(&JSONFormatter{})
    logger.Caller(true)  // 호출자 정보 비활성화
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("JSON 포맷팅된 메시지")
}
```

## 오류 처리

:::warning 참고 사항
성능상의 이유로 대부분의 로거 메서드는 오류를 반환하지 않습니다. 쓰기가 실패하면 로그는 조용히 폐기됩니다. 오류 인식 기능이 필요하면 사용자 정의 라이터를 사용하세요.
:::

출력 작업에 대한 오류 처리가 필요하면 사용자 정의 라이터를 구현하세요:

```go title="오류 캡처 라이터"
type ErrorCapturingWriter struct {
    writer io.Writer
    lastError error
}

func (w *ErrorCapturingWriter) Write(data []byte) (int, error) {
    n, err := w.writer.Write(data)
    if err != nil {
        w.lastError = err
    }
    return n, err
}

func (w *ErrorCapturingWriter) LastError() error {
    return w.lastError
}
```

## 스레드 안전성

:::tip 동시성 안전
모든 `Logger` 인스턴스의 메서드는 스레드 안전하며, 추가 동기화 메커니즘 없이 여러 고루틴에서 동시에 사용할 수 있습니다. 그러나 단일 `Entry` 객체는 스레드 안전하지 **않으며**, 일회용입니다.
:::

---

## 🌍 다국어 문서

이 문서는 여러 언어로 제공됩니다:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/API.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/API.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/API.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/API.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/API.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/API.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/API.md)
-   [🇰🇷 한국어](API.md) (현재)

---

**LazyGophers Log 전체 API 참조 - 탁월한 로깅으로 더 나은 애플리케이션을 구축하세요! 🚀**
