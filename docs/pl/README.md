---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Wysokowydajna i elastyczna biblioteka logowania Go, zbudowana na zap, zapewniająca bogate funkcje i prosty API.

## 📖 Języki dokumentacji

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 Chiński uproszczony](https://lazygophers.github.io/log/zh-CN/)
-   [🇹🇼 Chiński tradycyjny](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/)
-   [🇵🇹 Português](https://lazygophers.github.io/log/pt/)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/)
-   [🇵🇱 Polski](README.md) (aktualny)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/)
-   [🇹🇷 Türkçe](https://lazygophers.github.io/log/tr/)

## ✨ Funkcje

-   **🚀 Wysoka wydajność**：Zbudowany na zap z pulą obiektów i warunkowym rejestrowaniem pól
-   **📊 Bogate poziomy logowania**：Poziomy Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Elastyczna konfiguracja**：
    -   Kontrola poziomu logowania
    -   Rejestrowanie informacji o wywołującym
    -   Informacje o śledzeniu (w tym ID goroutine)
    -   Niestandardowe prefiksy i sufiksy
    -   Niestandardowe cele wyjściowe (konsola, pliki itd.)
    -   Opcje formatowania logów
-   **🔄 Rotacja plików**：Obsługa godzinnej rotacji plików dziennika
-   **🔌 Kompatybilność z Zap**：Bezproblemowa integracja z zap WriteSyncer
-   **🎯 Proste API**：Jasne API podobne do standardowej biblioteki logowania, łatwe w użyciu

## 🚀 Szybki start

### Instalacja

:::tip Instalacja
```bash
go get github.com/lazygophers/log
```
:::

### Podstawowe użycie

```go title="Szybki start"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Użyj domyślnego globalnego loggera
    log.Debug("Komunikat debugowania")
    log.Info("Komunikat informacyjny")
    log.Warn("Komunikat ostrzegawczy")
    log.Error("Komunikat o błędzie")

    // Użyj sformatowanego wyjścia
    log.Infof("Użytkownik %s pomyślnie zalogowany", "admin")

    // Niestandardowa konfiguracja
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("To jest log z niestandardowego loggera")
}
```

## 📚 Zaawansowane użycie

### Niestandardowy logger z wyjściem do pliku

```go title="Konfiguracja wyjścia do pliku"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Utwórz logger z wyjściem do pliku
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Log debugowania z informacjami o wywołującym")
    logger.Info("Log informacyjny z informacjami o śledzeniu")
}
```

### Kontrola poziomu logowania

```go title="Kontrola poziomu logowania"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Tylko warn i powyżej będą rejestrowane
    logger.Debug("To nie zostanie zarejestrowane")  // Zignorowane
    logger.Info("To nie zostanie zarejestrowane")   // Zignorowane
    logger.Warn("To zostanie zarejestrowane")    // Zarejestrowane
    logger.Error("To zostanie zarejestrowane")   // Zarejestrowane
}
```

## 🎯 Scenariusze użycia

### Odpowiednie scenariusze

-   **Serwisy WWW i backendy API**：Śledzenie żądań, ustrukturyzowane logowanie, monitorowanie wydajności
-   **Architektura mikrousług**：Rozproszone śledzenie (TraceID), ujednolicony format logów, wysoka przepustowość
-   **Narzędzia wiersza poleceń**：Kontrola poziomów, czyste wyjście, raportowanie błędów
-   **Zadania wsadowe**：Rotacja plików, długie wykonywanie, optymalizacja zasobów

### Specjalne zalety

-   **Optymalizacja z pulą obiektów**：Ponowne wykorzystanie obiektów Entry i Buffer, zmniejszenie presji GC
-   **Asynchroniczny zapis**：Wysoka przepustowość (10000+ logów/sekundę) bez blokowania
-   **Obsługa TraceID**：Śledzenie żądań w systemach rozproszonych, integracja z OpenTelemetry
-   **Start bez konfiguracji**：Gotowy do użycia, progresywna konfiguracja

## 🔧 Opcje konfiguracji

:::note Opcje konfiguracji
Wszystkie poniższe metody obsługują łączenie łańcuchowe i mogą być łączone w celu budowy niestandardowego Loggera.
:::

### Konfiguracja Loggera

| Metoda                  | Opis                | Domyślna      |
| --------------------- | ------------------- | ------------ |
| `SetLevel(level)`       | Ustaw minimalny poziom logowania     | `DebugLevel` |
| `EnableCaller(enable)`  | Włącz/wyłącz informacje o wywołującym | `false`      |
| `EnableTrace(enable)`   | Włącz/wyłącz informacje o śledzeniu  | `false`      |
| `SetCallerDepth(depth)` | Ustaw głębokość wywołującego   | `2`          |
| `SetPrefixMsg(prefix)`  | Ustaw prefiks logu  | `""`         |
| `SetSuffixMsg(suffix)`  | Ustaw sufiks logu  | `""`         |
| `SetOutput(writers...)` | Ustaw cele wyjściowe         | `os.Stdout`  |

### Poziomy logowania

| Poziom        | Opis                        |
| ----------- | --------------------------- |
| `TraceLevel` | Najbardziej szczegółowy, do szczegółowego śledzenia        |
| `DebugLevel` | Informacje o debugowaniu                  |
| `InfoLevel`  | Informacje ogólne                    |
| `WarnLevel`  | Komunikaty ostrzegawcze                  |
| `ErrorLevel` | Komunikaty o błędach                  |
| `FatalLevel` | Błędy krytyczne (wywołuje os.Exit(1))    |
| `PanicLevel` | Błędy paniki (wywołuje panic())      |

## 🏗️ Architektura

### Główne komponenty

-   **Logger**：Główna struktura logowania z konfigurowalnymi opcjami
-   **Entry**：Indywidualny rekord logu z obszerną obsługą pól
-   **Level**：Definicje poziomów logowania i funkcje użytkowe
-   **Format**：Interfejs formatowania logu i implementacje

### Optymalizacja wydajności

-   **Pula obiektów**：Ponownie wykorzystuje obiekty Entry w celu zmniejszenia alokacji pamięci
-   **Warunkowe rejestrowanie**：Rejestruje tylko kosztowne pola gdy konieczne
-   **Szybka kontrola poziomów**：Sprawdza poziom logowania na zewnętrznej warstwie
-   **Projekt bez blokad**：Większość operacji nie wymaga blokad

## 📊 Porównanie wydajności

:::info Porównanie wydajności
Poniższe dane opierają się na benchmarkach; rzeczywista wydajność może się różnić w zależności od środowiska i konfiguracji.
:::

| Funkcja          | lazygophers/log | zap    | logrus | Standardowy log |
| ------------- | --------------- | ------ | ------ | ---------------- |
| Wydajność      | Wysoka              | Wysoka     | Średnia     | Niska       |
| Prostota API    | Wysoka              | Średnia     | Wysoka     | Wysoka       |
| Bogactwo funkcji    | Średnia          | Wysoka     | Wysoka     | Niska       |
| Elastyczność      | Średnia          | Wysoka     | Wysoka     | Niska       |
| Krzywa uczenia      | Niska              | Średnia     | Średnia     | Niska       |

## ❓ Często zadawane pytania

### Jak wybrać odpowiedni poziom logowania?

-   **Środowisko programistyczne**：Użyj `DebugLevel` lub `TraceLevel` dla szczegółowych informacji
-   **Środowisko produkcyjne**：Użyj `InfoLevel` lub `WarnLevel` w celu zmniejszenia narzutu
-   **Testy wydajności**：Użyj `PanicLevel` w celu wyłączenia wszystkich logów

### Jak zoptymalizować wydajność w środowisku produkcyjnym?

:::warning Uwaga
W scenariuszach o wysokiej przepustowości zaleca się połączenie asynchronicznego zapisu z rozsądnymi poziomami logowania w celu optymalizacji wydajności.
:::

1. Użyj `AsyncWriter` do asynchronicznego zapisu：

```go title="Konfiguracja asynchronicznego zapisu"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Dostosuj poziomy logowania, aby uniknąć niepotrzebnych logów：

```go title="Optymalizacja poziomu"
logger.SetLevel(log.InfoLevel)  // Pomiń Debug i Trace
```

3. Użyj warunkowego logowania w celu zmniejszenia narzutu：

```go title="Warunkowe logowanie"
if logger.Level >= log.DebugLevel {
    logger.Debug("Szczegółowe informacje debugowania")
}
```

### Jaka jest różnica między `Caller` a `EnableCaller`?

-   **`EnableCaller(enable bool)`**：Kontroluje, czy Logger zbiera informacje o wywołującym
    -   `EnableCaller(true)` włącza zbieranie informacji o wywołującym
-   **`Caller(disable bool)`**：Kontroluje, czy Formatter wypisuje informacje o wywołującym
    -   `Caller(true)` wyłącza wypisywanie informacji o wywołującym

Zaleca się używanie `EnableCaller` do globalnej kontroli.

### Jak zaimplementować niestandardowy formater?

Zaimplementuj interfejs `Format`：

```go title="Niestandardowy formater"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Powiązana dokumentacja

-   [📚 Dokumentacja API](API.md) - Pełne odwołanie API
-   [🤝 Przewodnik po wkładzie](/pl/CONTRIBUTING) - Jak wnieść wkład
-   [📋 Dziennik zmian](/pl/CHANGELOG) - Historia wersji
-   [🔒 Polityka bezpieczeństwa](/pl/SECURITY) - Wytyczne bezpieczeństwa
-   [📜 Zasady postępowania](/pl/CODE_OF_CONDUCT) - Wytyczne społeczności

## 🚀 Uzyskanie pomocy

-   **GitHub Issues**：[Zgłaszanie błędów lub żądanie funkcji](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[Dokumentacja API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Przykłady**：[Przykłady użycia](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licencja

Ten projekt jest licencjonowany na licencji MIT - zobacz plik [LICENSE](/pl/LICENSE) dla szczegółów.

## 🤝 Wkład

Doceniamy wkład! Zapoznaj się z naszym [Przewodnikiem po wkładzie](/pl/CONTRIBUTING) aby uzyskać więcej informacji.

---

**lazygophers/log** został zaprojektowany, aby być preferowanym rozwiązaniem logowania dla programistów Go, którzy cenią zarówno wydajność, jak i prostotę. Niezależnie od tego, czy budujesz małe narzędzie czy wielkoskalowy system rozproszony, ta biblioteka zapewnia odpowiednią równowagę między funkcjonalnością a łatwością użycia.
