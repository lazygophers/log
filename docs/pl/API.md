---
titleSuffix: ' | LazyGophers Log'
---
# 📚 Dokumentacja API

## Przegląd

LazyGophers Log zapewnia kompleksowe API rejestrowania z obsługą wielu poziomów dziennika, niestandardowego formatowania, asynchronicznego zapisu i optymalizacji tagów kompilacji. Ten dokument obejmuje wszystkie publiczne interfejsy API, opcje konfiguracji i wzorce użytkowania.

## Spis treści

-   [Podstawowe Typy](#podstawowe-typy)
-   [API Loggera](#api-loggera)
-   [Funkcje Globalne](#funkcje-globalne)
-   [Poziomy Dziennika](#poziomy-dziennika)
-   [Formatery](#formatery)
-   [Writery Wyjściowe](#writery-wyjściowe)
-   [Logowanie z Kontekstem](#logowanie-z-kontekstem)
-   [Tagi Kompilacji](#tagi-kompilacji)
-   [Optymalizacja Wydajności](#optymalizacja-wydajności)
-   [Przykłady](#przykłady)

## Podstawowe Typy

### Logger

Główna struktura loggera zapewniająca wszystkie funkcje rejestrowania.

```go
type Logger struct {
    // Zawiera pola prywatne dla operacji thread-safe
}
```

#### Konstruktor

```go
func New() *Logger
```

Tworzy nową instancję loggera z domyślną konfiguracją:

-   Poziom: `DebugLevel`
-   Wyjście: `os.Stdout`
-   Formater: Domyślny formater tekstowy
-   Śledzenie wywołującego: Wyłączone

**Przykład:**

```go title="Tworzenie loggera"
logger := log.New()
logger.Info("Utworzono nowy logger")
```

### Entry

Reprezentuje pojedynczy wpis dziennika ze wszystkimi powiązanymi metadanymi.

```go
type Entry struct {
    Time       time.Time     // Znacznik czasu utworzenia wpisu
    Level      Level         // Poziom dziennika
    Message    string        // Komunikat dziennika
    Pid        int          // ID procesu
    Gid        uint64       // ID goroutine
    TraceID    string       // ID śledzenia dla rozproszonego śledzenia
    CallerName string       // Nazwa funkcji wywołującej
    CallerFile string       // Ścieżka pliku wywołującego
    CallerLine int          // Numer wiersza wywołującego
}
```

## API Loggera

### Metody Konfiguracji

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Ustawia minimalny poziom dziennika. Komunikaty poniżej tego poziomu będą ignorowane.

**Parametry:**

-   `level`: Minimalny poziom dziennika do przetworzenia

**Zwraca:**

-   `*Logger`: Zwraca sam siebie w celu obsługi łańcuchowania metod

**Przykład:**

```go title="Ustawianie poziomu dziennika"
logger.SetLevel(log.InfoLevel)
logger.Debug("To nie zostanie wyświetlone")  // Zignorowane
logger.Info("To zostanie wyświetlone")        // Przetworzone
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Ustawia jeden lub więcej celów wyjściowych dla komunikatów dziennika.

**Parametry:**

-   `writers`: Jeden lub więcej celów wyjściowych `io.Writer`

**Zwraca:**

-   `*Logger`: Zwraca sam siebie w celu obsługi łańcuchowania metod

**Przykład:**

```go title="Ustawianie celu wyjściowego"
// Pojedyncze wyjście
logger.SetOutput(os.Stdout)

// Wiele wyjść
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

Ustawia niestandardowy formater dla wyjścia dziennika.

**Parametry:**

-   `formatter`: Formater implementujący interfejs `Format`

**Zwraca:**

-   `*Logger`: Zwraca sam siebie w celu obsługi łańcuchowania metod

**Przykład:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Włącza lub wyłącza rejestrowanie informacji o wywołującym w wpisach dziennika.

**Parametry:**

-   `enable`: Czy włączyć informacje o wywołującym (przekaż `true`, aby włączyć)

**Zwraca:**

-   `*Logger`: Zwraca sam siebie w celu obsługi łańcuchowania metod

**Przykład:**

```go
logger.EnableCaller(true)
logger.Info("To będzie zawierać informacje plik:wiersz")

logger.EnableCaller(false)
logger.Info("To nie będzie zawierać informacji plik:wiersz")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

Kontroluje informacje o wywołującym w formaterze.

**Parametry:**

-   `disable`: Czy wyłączyć informacje o wywołującym (przekaż `true`, aby wyłączyć)

**Zwraca:**

-   `*Logger`: Zwraca sam siebie w celu obsługi łańcuchowania metod

**Przykład:**

```go
logger.Caller(false)  // Nie wyłączaj, pokaż informacje o wywołującym
logger.Info("To będzie zawierać informacje plik:wiersz")

logger.Caller(true)   // Wyłącz informacje o wywołującym
logger.Info("To nie będzie zawierać informacji plik:wiersz")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Ustawia głębokość stosu dla informacji o wywołującym przy opakowywaniu loggera.

**Parametry:**

-   `depth`: Liczba ramek stosu do pominięcia

**Zwraca:**

-   `*Logger`: Zwraca sam siebie w celu obsługi łańcuchowania metod

**Przykład:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // Pomiń funkcję wrapper
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Ustawia tekst przedrostka lub przyrostka dla wszystkich komunikatów dziennika.

**Parametry:**

-   `prefix/suffix`: Tekst do dodania przed/po komunikacie

**Zwraca:**

-   `*Logger`: Zwraca sam siebie w celu obsługi łańcuchowania metod

**Przykład:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // Wyjście: [APP] Hello [END]
```

### Metody Rejestrowania

Wszystkie metody rejestrowania mają dwie warianty: prostą wersję i wersję sformatowaną.

#### Poziom Trace

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

Rejestruje dziennik na poziomie trace (najbardziej szczegółowy).

**Przykład:**

```go
logger.Trace("Szczegółowe śledzenie wykonania")
logger.Tracef("Przetwarzanie elementu %d z %d", i, total)
```

#### Poziom Debug

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

Rejestruje informacje programistyczne na poziomie debug.

**Przykład:**

```go
logger.Debug("Stan zmiennej:", variable)
logger.Debugf("Użytkownik %s pomyślnie uwierzytelniony", username)
```

#### Poziom Info

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

Rejestruje komunikaty informacyjne.

**Przykład:**

```go
logger.Info("Aplikacja uruchomiona")
logger.Infof("Serwer nasłuchuje na porcie %d", port)
```

#### Poziom Warn

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

Rejestruje komunikaty ostrzeżeń dla potencjalnie problematycznych sytuacji.

**Przykład:**

```go
logger.Warn("Wywołano przestarzałą funkcję")
logger.Warnf("Wysokie użycie pamięci: %d%%", memoryPercent)
```

#### Poziom Error

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

Rejestruje komunikaty błędów.

**Przykład:**

```go
logger.Error("Nieudane połączenie z bazą danych")
logger.Errorf("Nieudane przetwarzanie żądania: %v", err)
```

#### Poziom Fatal

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

Rejestruje błąd krytyczny i wywołuje `os.Exit(1)`.

:::danger Operacja Destrukcyjna
`Fatal` i `Fatalf` wywołają `os.Exit(1)` natychmiast po rejestrowaniu, kończąc proces. Używaj tylko w warunkach błędów nieodwracalnych. Instrukcje `defer` **nie** zostaną wykonane.
:::

**Przykład:**

```go
logger.Fatal("Krytyczny błąd systemu")
logger.Fatalf("Nie można uruchomić serwera: %v", err)
```

#### Poziom Panic

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

Rejestruje komunikat błędu i wywołuje `panic()`.

:::danger Operacja Destrukcyjna
`Panic` i `Panicf` wywołają `panic()` po rejestrowaniu. W przeciwieństwie do `Fatal`, `panic` może zostać odzyskany za pomocą `recover()`, ale zakończy program, jeśli nie zostanie przechwycony.
:::

**Przykład:**

```go
logger.Panic("Wystąpił nieodwracalny błąd")
logger.Panicf("Nieprawidłowy stan: %v", state)
```

### Metody Narzędziowe

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Tworzy kopię loggera z tą samą konfiguracją.

**Zwraca:**

-   `*Logger`: Nowa instancja loggera ze skopiowanymi ustawieniami

**Przykład:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

Tworzy logger świadomy kontekstu, który akceptuje `context.Context` jako pierwszy parametr.

**Zwraca:**

-   `LoggerWithCtx`: Instancja loggera świadomego kontekstu

**Przykład:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "Komunikat świadomy kontekstu")
```

## Funkcje Globalne

Funkcje na poziomie pakietu używające domyślnego globalnego loggera.

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

**Przykład:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("Używanie globalnego loggera")
```

## Poziomy Dziennika

### Typ Level

```go
type Level int8
```

### Dostępne Poziomy

```go
const (
    PanicLevel Level = iota  // 0 - Panic i wyjdź
    FatalLevel              // 1 - Błąd krytyczny i wyjdź
    ErrorLevel              // 2 - Warunek błędu
    WarnLevel               // 3 - Warunek ostrzeżenia
    InfoLevel               // 4 - Komunikat informacyjny
    DebugLevel              // 5 - Komunikat debug
    TraceLevel              // 6 - Najbardziej szczegółowe śledzenie
)
```

### Metody Level

```go
func (l Level) String() string
```

Zwraca reprezentację tekstową poziomu.

**Przykład:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## Formatery

### Interfejs Format

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

Niestandardowe formatery muszą implementować ten interfejs.

### Domyślny Formater

Wbudowany formater tekstowy z konfigurowalnymi opcjami.

```go
type Formatter struct {
    // Opcje konfiguracji
}
```

### Przykład Formatera JSON

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

// Użycie
logger.SetFormatter(&JSONFormatter{})
```

## Writery Wyjściowe

### Wyjście Plikowe i Rotacja

```go
func GetOutputWriterHourly(filename string) io.Writer
```

Tworzy writer, który obraca pliki dziennika co godzinę.

**Parametry:**

-   `filename`: Podstawowa nazwa pliku dziennika

**Zwraca:**

-   `io.Writer`: Writer pliku z rotacją

**Przykład:**

```go title="Rotacja dziennika co godzinę"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// Tworzy pliki takie jak: app-2024010115.log, app-2024010116.log, itd.
```

### Writer Asynchroniczny

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

Tworzy asynchroniczny writer do rejestrowania o wysokiej wydajności.

**Parametry:**

-   `writer`: Podstawowy writer
-   `bufferSize`: Rozmiar bufora wewnętrznego

**Zwraca:**

-   `*AsyncWriter`: Instancja asynchronicznego writera

**Metody:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**Przykład:**

```go title="Writer asynchroniczny"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## Logowanie z Kontekstem

### Interfejs LoggerWithCtx

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

### Funkcje Kontekstowe

```go
func SetTrace(traceID string)
func GetTrace() string
```

Ustawia i pobiera ID śledzenia dla bieżącej goroutine.

**Przykład:**

```go
log.SetTrace("trace-123-456")
log.Info("Ten komunikat będzie zawierał ID śledzenia")

traceID := log.GetTrace()
fmt.Println("Bieżące ID śledzenia:", traceID)
```

## Tagi Kompilacji

Ta biblioteka obsługuje warunkową kompilację z tagami kompilacji:

:::info Opis Tagów Kompilacji
Tagi kompilacji są określane przez parametr `go build -tags`. Różne tagi zmieniają zachowanie kompilacji i cechy runtime biblioteki dziennika. Wybór odpowiednich tagów pozwala zrównoważyć wygodę programowania i wydajność produkcji.
:::

### Tryb Domyślny

```bash
go build
```

-   Pełna funkcjonalność włączona
-   Komunikaty debugowania włączone
-   Standardowa wydajność

### Tryb Debug

```bash
go build -tags debug
```

-   Rozszerzone informacje debugowania
-   Dodatkowe kontrole runtime
-   Szczegółowe informacje o wywołującym

### Tryb Release

```bash
go build -tags release
```

-   Zoptymalizowane dla środowiska produkcyjnego
-   Komunikaty debugowania wyłączone
-   Automatyczna rotacja dzienników włączona

### Tryb Discard

```bash
go build -tags discard
```

-   Maksymalna wydajność
-   Wszystkie operacje rejestrowania to no-ops
-   Zero narzutu

### Tryb Połączony

```bash
go build -tags "debug,discard"    # Debug i Discard
go build -tags "release,discard"  # Release i Discard
```

## Optymalizacja Wydajności

:::tip Najlepsze Praktyki Wydajności
Ta biblioteka jest głęboko zoptymalizowana przez mechanizmy takie jak puli obiektów, wstępne kontrole poziomu i asynchroniczny zapis. W scenariuszach o dużej przepustowości zaleca się łączenie asynchronicznych writerów i odpowiednich tagów kompilacji dla uzyskania najlepszej wydajności.
:::

### Puli Obiektów

Biblioteka używa wewnętrznie `sync.Pool` do zarządzania:

-   Obiektami wpisów dziennika
-   Buforami bajtów
-   Buforami formatera

To zmniejsza presję garbage collection w scenariuszach o dużej przepustowości.

### Kontrola Poziomu

Kontrole poziomu dziennika występują przed kosztownymi operacjami:

```go
// Efektywne - formatowanie komunikatu tylko gdy poziom włączony
logger.Debugf("Wynik kosztownej operacji: %+v", expensiveCall())

// Mniej efektywne gdy debug jest wyłączony w produkcji
result := expensiveCall()
logger.Debug("Wynik:", result)
```

### Asynchroniczny Zapis

Dla aplikacji o dużej przepustowości:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // Duży bufor
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### Optymalizacja Tagów Kompilacji

Użyj odpowiednich tagów kompilacji dla środowiska:

-   Programowanie: Tagi domyślne lub debug
-   Produkcja: Tagi release
-   Krytyczne dla wydajności: Tagi discard

## Przykłady

### Podstawowe Użycie

```go title="Podstawowe użycie"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("Uruchamianie aplikacji")
    log.Warn("To jest ostrzeżenie")
    log.Error("To jest błąd")
}
```

### Niestandardowy Logger

```go title="Konfiguracja niestandardowa"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // Konfiguruj logger
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  // Wyłącz informacje o wywołującym
    logger.SetPrefixMsg("[MyApp] ")

    // Ustaw wyjście na plik
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("Skonfigurowano niestandardowy logger")
    logger.Debug("Informacje debugowania z wywołującym")
}
```

### Rejestrowanie o Wysokiej Wydajności

```go title="Rejestrowanie o wysokiej wydajności"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Utwórz writer z rotacją co godzinę
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // Użyj writera asynchronicznego dla lepszej wydajności
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // Pomiń logi debugowania w produkcji

    // Rejestrowanie o dużej przepustowości
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### Logowanie Świadome Kontekstu

```go title="Logowanie świadome kontekstu"
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

    ctxLogger.Info(ctx, "Przetwarzanie żądania użytkownika")
    ctxLogger.Debug(ctx, "Walidacja zakończona")
}
```

### Niestandardowy Formater JSON

```go title="Niestandardowy formater JSON"
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
    logger.Caller(true)  // Wyłącz informacje o wywołującym
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("Komunikat sformatowany w JSON")
}
```

## Obsługa Błędów

:::warning Uwaga
Ze względów wydajności większość metod loggera nie zwraca błędów. Jeśli zapis nie powiedzie się, dzienniki będą cicho odrzucone. Jeśli potrzebujesz świadomości błędów, użyj niestandardowego writera.
:::

Jeśli potrzebujesz obsługi błędów dla operacji wyjściowych, zaimplementuj niestandardowy writer:

```go title="Writer przechwytujący błędy"
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

## Bezpieczeństwo Wątków

:::tip Bezpieczeństwo Współbieżności
Wszystkie metody instancji `Logger` są thread-safe i mogą być używane jednocześnie w wielu goroutines bez dodatkowych mechanizmów synchronizacji. Należy jednak pamiętać, że pojedyncze obiekty `Entry` **nie są** thread-safe i są do jednokrotnego użytku.
:::

---

## 🌍 Dokumentacja Wielojęzyczna

Ta dokumentacja jest dostępna w wielu językach:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/API.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/API.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/API.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/API.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/API.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/API.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/API.md)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/API.md)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/API.md)
-   [🇧🇷 Português](https://lazygophers.github.io/log/pt/API.md)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/API.md)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/API.md)
-   [🇵🇱 Polski](API.md) (Aktualny)

---

**Pełne odwołanie API LazyGophers Log - Buduj lepsze aplikacje z wyjątkowym rejestrowaniem! 🚀**
