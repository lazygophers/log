---
titleSuffix: ' | LazyGophers Log'
---
# 📚 API-documentatie

## Overzicht

LazyGophers Log biedt een uitgebreide logging-API met ondersteuning voor meerdere logniveaus, aangepaste opmaak, asynchroon schrijven en build-tag optimalisatie. Dit document behandelt alle openbare API's, configuratieopties en gebruikspatronen.

## Inhoudsopgave

-   [Kern Types](#kern-types)
-   [Logger API](#logger-api)
-   [Globale Functies](#globale-functies)
-   [Logniveaus](#logniveaus)
-   [Formatters](#formatters)
-   [Uitvoer Writers](#uitvoer-writers)
-   [Context Logging](#context-logging)
-   [Build Tags](#build-tags)
-   [Prestatieoptimalisatie](#prestatieoptimalisatie)
-   [Voorbeelden](#voorbeelden)

## Kern Types

### Logger

De belangrijkste logger structuur die alle logging functionaliteit biedt.

```go
type Logger struct {
    // Bevat private velden voor thread-safe operaties
}
```

#### Constructor

```go
func New() *Logger
```

Maakt een nieuwe logger instantie aan met standaardconfiguratie:

-   Niveau: `DebugLevel`
-   Uitvoer: `os.Stdout`
-   Formatter: Standaard tekst formatter
-   Aanroeper tracing: Uitgeschakeld

**Voorbeeld:**

```go title="Logger aanmaken"
logger := log.New()
logger.Info("Nieuwe logger aangemaakt")
```

### Entry

Vertegenwoordigt een enkele loginvoer met alle bijbehorende metagegevens.

```go
type Entry struct {
    Time       time.Time     // Tijdstempel bij aanmaken van invoer
    Level      Level         // Logniveau
    Message    string        // Logbericht
    Pid        int          // Proces ID
    Gid        uint64       // Goroutine ID
    TraceID    string       // Trace ID voor gedistribueerde tracing
    CallerName string       // Naam van de aanroepende functie
    CallerFile string       // Pad van het aanroepende bestand
    CallerLine int          // Regelnummer van de aanroeper
}
```

## Logger API

### Configuratie Methoden

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Stelt het minimum logniveau in. Berichten onder dit niveau worden genegeerd.

**Parameters:**

-   `level`: Het minimum logniveau dat moet worden verwerkt

**Retourneert:**

-   `*Logger`: Retourneert zichzelf om method chaining te ondersteunen

**Voorbeeld:**

```go title="Logniveau instellen"
logger.SetLevel(log.InfoLevel)
logger.Debug("Dit wordt niet weergegeven")  // Genegeerd
logger.Info("Dit wordt weergegeven")        // Verwerkt
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Stelt een of meer uitvoerbestemmingen in voor logberichten.

**Parameters:**

-   `writers`: Een of meer `io.Writer` uitvoerbestemmingen

**Retourneert:**

-   `*Logger`: Retourneert zichzelf om method chaining te ondersteunen

**Voorbeeld:**

```go title="Uitvoerbestemming instellen"
// Enkele uitvoer
logger.SetOutput(os.Stdout)

// Meerdere uitvoer
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

Stelt een aangepaste formatter in voor de loguitvoer.

**Parameters:**

-   `formatter`: Een formatter die de `Format` interface implementeert

**Retourneert:**

-   `*Logger`: Retourneert zichzelf om method chaining te ondersteunen

**Voorbeeld:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Activeert of deactiveert het vastleggen van aanroeperinformatie in loginvoeren.

**Parameters:**

-   `enable`: Of aanroeperinformatie moet worden ingeschakeld (geef `true` op om in te schakelen)

**Retourneert:**

-   `*Logger`: Retourneert zichzelf om method chaining te ondersteunen

**Voorbeeld:**

```go
logger.EnableCaller(true)
logger.Info("Dit bevat bestand:regelnummer informatie")

logger.EnableCaller(false)
logger.Info("Dit bevat geen bestand:regelnummer informatie")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

Controleert de aanroeperinformatie in de formatter.

**Parameters:**

-   `disable`: Of aanroeperinformatie moet worden uitgeschakeld (geef `true` op om uit te schakelen)

**Retourneert:**

-   `*Logger`: Retourneert zichzelf om method chaining te ondersteunen

**Voorbeeld:**

```go
logger.Caller(false)  // Niet uitschakelen, toon aanroeperinformatie
logger.Info("Dit bevat bestand:regelnummer informatie")

logger.Caller(true)   // Aanroeperinformatie uitschakelen
logger.Info("Dit bevat geen bestand:regelnummer informatie")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Stelt de stackdiepte in voor aanroeperinformatie bij het wrappen van de logger.

**Parameters:**

-   `depth`: Het aantal stack frames dat moet worden overgeslagen

**Retourneert:**

-   `*Logger`: Retourneert zichzelf om method chaining te ondersteunen

**Voorbeeld:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // Wrapper functie overslaan
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Stelt voor- of achtervoegsel tekst in voor alle logberichten.

**Parameters:**

-   `prefix/suffix`: De tekst die voor/na het bericht moet worden geplaatst

**Retourneert:**

-   `*Logger`: Retourneert zichzelf om method chaining te ondersteunen

**Voorbeeld:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // Uitvoer: [APP] Hello [END]
```

### Logging Methoden

Alle logging methoden hebben twee varianten: een eenvoudige versie en een geformatteerde versie.

#### Trace Niveau

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

Logt op trace niveau (meest gedetailleerd).

**Voorbeeld:**

```go
logger.Trace("Gedetailleerde uitvoeringstracering")
logger.Tracef("Verwerken item %d van %d", i, total)
```

#### Debug Niveau

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

Logt ontwikkelingsinformatie op debug niveau.

**Voorbeeld:**

```go
logger.Debug("Variabele status:", variable)
logger.Debugf("Gebruiker %s succesvol geverifieerd", username)
```

#### Info Niveau

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

Logt informatieve berichten.

**Voorbeeld:**

```go
logger.Info("Applicatie gestart")
logger.Infof("Server luistert op poort %d", port)
```

#### Warn Niveau

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

Logt waarschuwingsberichten voor potentiële probleemsituaties.

**Voorbeeld:**

```go
logger.Warn("Verouderde functie aangeroepen")
logger.Warnf("Geheugengebruik hoog: %d%%", memoryPercent)
```

#### Error Niveau

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

Logt foutberichten.

**Voorbeeld:**

```go
logger.Error("Databaseverbinding mislukt")
logger.Errorf("Verwerken van aanvraag mislukt: %v", err)
```

#### Fatal Niveau

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

Logt een fatale fout en roept `os.Exit(1)` aan.

:::danger Destructieve Operatie
`Fatal` en `Fatalf` roepen onmiddellijk na het loggen `os.Exit(1)` aan om het proces te beëindigen. Gebruik dit alleen bij onherstelbare foutcondities. `defer` instructies worden **niet** uitgevoerd.
:::

**Voorbeeld:**

```go
logger.Fatal("Kritieke systeemfout")
logger.Fatalf("Kan server niet starten: %v", err)
```

#### Panic Niveau

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

Logt een foutbericht en roept `panic()` aan.

:::danger Destructieve Operatie
`Panic` en `Panicf` roepen na het loggen `panic()` aan. In tegenstelling tot `Fatal` kan `panic` worden opgevangen met `recover()`, maar beëindigt het het programma als het niet wordt opgevangen.
:::

**Voorbeeld:**

```go
logger.Panic("Er is een onherstelbare fout opgetreden")
logger.Panicf("Ongeldige status: %v", state)
```

### Hulpmiddelen

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Maakt een kopie van de logger met dezelfde configuratie.

**Retourneert:**

-   `*Logger`: Nieuwe logger instantie met gekopieerde instellingen

**Voorbeeld:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

Maakt een context-bewuste logger die `context.Context` als eerste parameter accepteert.

**Retourneert:**

-   `LoggerWithCtx`: Context-bewuste logger instantie

**Voorbeeld:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "Context-bewust bericht")
```

## Globale Functies

Pakket-niveau functies die de standaard globale logger gebruiken.

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

**Voorbeeld:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("Gebruik van globale logger")
```

## Logniveaus

### Level Type

```go
type Level int8
```

### Beschikbare Niveaus

```go
const (
    PanicLevel Level = iota  // 0 - Paniek en afsluiten
    FatalLevel              // 1 - Fatale fout en afsluiten
    ErrorLevel              // 2 - Foutconditie
    WarnLevel               // 3 - Waarschuwingsconditie
    InfoLevel               // 4 - Informatief bericht
    DebugLevel              // 5 - Debug bericht
    TraceLevel              // 6 - Meest gedetailleerde tracering
)
```

### Level Methoden

```go
func (l Level) String() string
```

Retourneert de teksrepresentatie van het niveau.

**Voorbeeld:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## Formatters

### Format Interface

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

Aangepaste formatters moeten deze interface implementeren.

### Standaard Formatter

Ingebouwde tekst formatter met aanpasbare opties.

```go
type Formatter struct {
    // Configuratie opties
}
```

### JSON Formatter Voorbeeld

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

// Gebruik
logger.SetFormatter(&JSONFormatter{})
```

## Uitvoer Writers

### Bestandsuitvoer en Rotatie

```go
func GetOutputWriterHourly(filename string) io.Writer
```

Maakt een writer aan die logbestanden elk uur roteert.

**Parameters:**

-   `filename`: De basis bestandsnaam van het logbestand

**Retourneert:**

-   `io.Writer`: Roterende bestand writer

**Voorbeeld:**

```go title="Uurlijkse logrotatie"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// Maakt bestanden zoals: app-2024010115.log, app-2024010116.log, etc.
```

### Asynchrone Writer

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

Maakt een asynchrone writer aan voor hoogwaardige logging.

**Parameters:**

-   `writer`: De onderliggende writer
-   `bufferSize`: Grootte van de interne buffer

**Retourneert:**

-   `*AsyncWriter`: Asynchrone writer instantie

**Methoden:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**Voorbeeld:**

```go title="Asynchrone writer"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## Context Logging

### LoggerWithCtx Interface

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

### Context Functies

```go
func SetTrace(traceID string)
func GetTrace() string
```

Stelt de trace ID in en haalt deze op voor de huidige goroutine.

**Voorbeeld:**

```go
log.SetTrace("trace-123-456")
log.Info("Dit bericht bevat de trace ID")

traceID := log.GetTrace()
fmt.Println("Huidige trace ID:", traceID)
```

## Build Tags

Deze bibliotheek ondersteunt conditionele compilatie met build tags:

:::info Build Tag Beschrijving
Build tags worden gespecificeerd via de parameter `go build -tags`. Verschillende tags veranderen het compilatiegedrag en runtime kenmerken van de log bibliotheek. Het kiezen van geschikte tags maakt een balans mogelijk tussen ontwikkelingsgemak en productieprestaties.
:::

### Standaard Modus

```bash
go build
```

-   Volledige functionaliteit ingeschakeld
-   Debug berichten inbegrepen
-   Standaard prestaties

### Debug Modus

```bash
go build -tags debug
```

-   Uitgebreide debug informatie
-   Extra runtime controles
-   Gedetailleerde aanroeperinformatie

### Release Modus

```bash
go build -tags release
```

-   Geoptimaliseerd voor productieomgeving
-   Debug berichten uitgeschakeld
-   Automatische logrotatie ingeschakeld

### Discard Modus

```bash
go build -tags discard
```

-   Maximale prestaties
-   Alle logging operaties zijn no-ops
-   Zero overhead

### Gecombineerde Modus

```bash
go build -tags "debug,discard"    # Debug en Discard
go build -tags "release,discard"  # Release en Discard
```

## Prestatieoptimalisatie

:::tip Prestatie Best Practices
Deze bibliotheek is diep geoptimaliseerd door mechanismen zoals object pools, voorafgaande niveaucontroles en asynchroon schrijven. In scenario's met hoge doorvoer wordt aanbevolen om asynchrone writers en geschikte build tags te combineren voor de beste prestaties.
:::

### Object Pools

De bibliotheek gebruikt intern `sync.Pool` om te beheren:

-   Log invoer objecten
-   Byte buffers
-   Formatter buffers

Dit reduceert de garbage collection druk in scenario's met hoge doorvoer.

### Niveau Controle

Logniveau controles vinden plaats vóór dure operaties:

```go
// Efficiënt - berichtopmaak alleen wanneer niveau ingeschakeld
logger.Debugf("Resultaat van dure operatie: %+v", expensiveCall())

// Minder efficiënt wanneer debug in productie is uitgeschakeld
result := expensiveCall()
logger.Debug("Resultaat:", result)
```

### Asynchroon Schrijven

Voor toepassingen met hoge doorvoer:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // Grote buffer
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### Build Tag Optimalisatie

Gebruik geschikte build tags voor de omgeving:

-   Ontwikkeling: Standaard of debug tags
-   Productie: Release tags
-   Prestatie kritisch: Discard tags

## Voorbeelden

### Basis Gebruik

```go title="Basis gebruik"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("Applicatie starten")
    log.Warn("Dit is een waarschuwing")
    log.Error("Dit is een fout")
}
```

### Aangepaste Logger

```go title="Aangepaste configuratie"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // Logger configureren
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  // Aanroeperinformatie uitschakelen
    logger.SetPrefixMsg("[MyApp] ")

    // Uitvoer naar bestand instellen
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("Aangepaste logger geconfigureerd")
    logger.Debug("Debug informatie met aanroeper")
}
```

### Hoogwaardige Logging

```go title="Hoogwaardige logging"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Roterende bestand writer aanmaken
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // Asynchrone writer gebruiken voor betere prestaties
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // Debug logs overslaan in productie

    // Hoge doorvoer logging
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### Context-bewust Logging

```go title="Context-bewust logging"
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

    ctxLogger.Info(ctx, "Verwerken van gebruikersaanvraag")
    ctxLogger.Debug(ctx, "Validatie voltooid")
}
```

### Aangepaste JSON Formatter

```go title="Aangepaste JSON formatter"
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
    logger.Caller(true)  // Aanroeperinformatie uitschakelen
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("JSON geformatteerd bericht")
}
```

## Foutafhandeling

:::warning Let op
Om prestatieredenen geven de meeste logger methoden geen fouten terug. Als schrijven mislukt, worden logs stilzwijgend weggegooid. Als u foutbewustzijn nodig hebt, gebruik dan een aangepaste writer.
:::

Als u foutafhandeling nodig hebt voor uitvoeroperaties, implementeer dan een aangepaste writer:

```go title="Foutopvangende writer"
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

## Thread Veiligheid

:::tip Concurrentie Veilig
Alle methoden van `Logger` instanties zijn thread-safe en kunnen gelijktijdig worden gebruikt in meerdere goroutines zonder extra synchronisatiemechanismen. Merk echter op dat individuele `Entry` objecten **niet** thread-safe zijn en voor eenmalig gebruik zijn.
:::

---

## 🌍 Meertalige Documentatie

 Deze documentatie is beschikbaar in meerdere talen:

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
-   [🇳🇱 Nederlands](API.md) (Huidige)

---

**LazyGophers Log volledige API referentie - Bouw betere applicaties met uitzonderlijke logging! 🚀**
