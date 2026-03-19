---
titleSuffix: ' | LazyGophers Log'
---
# 📚 API-Dokumentation

## Überblick

LazyGophers Log bietet eine umfassende Logging-API mit Unterstützung für mehrere Log-Level, benutzerdefinierte Formatierung, asynchrones Schreiben und Build-Tag-Optimierung. Dieses Dokument deckt alle öffentlichen APIs, Konfigurationsoptionen und Nutzungsmuster ab.

## Inhaltsverzeichnis

-   [Kerntypen](#kerntypen)
-   [Logger-API](#logger-api)
-   [Globale Funktionen](#globale-funktionen)
-   [Log-Level](#log-level)
-   [Formatierer](#formatierer)
-   [Ausgabe-Writer](#ausgabe-writer)
-   [Kontext-Logging](#kontext-logging)
-   [Build-Tags](#build-tags)
-   [Leistungsoptimierung](#leistungsoptimierung)
-   [Beispiele](#beispiele)

## Kerntypen

### Logger

Die Haupt-Logger-Struktur, die alle Logging-Funktionen bereitstellt.

```go
type Logger struct {
    // Enthält private Felder für threadsichere Operationen
}
```

#### Konstruktor

```go
func New() *Logger
```

Erstellt eine neue Logger-Instanz mit Standardkonfiguration:

-   Level: `DebugLevel`
-   Ausgabe: `os.Stdout`
-   Formatierer: Standard-Textformatierer
-   Aufrufer-Verfolgung: Deaktiviert

**Beispiel:**

```go title="Logger erstellen"
logger := log.New()
logger.Info("Neuer Logger erstellt")
```

### Entry

Stellt einen einzelnen Log-Eintrag mit allen zugehörigen Metadaten dar.

```go
type Entry struct {
    Time       time.Time     // Zeitstempel bei Erstellung des Eintrags
    Level      Level         // Log-Level
    Message    string        // Log-Nachricht
    Pid        int          // Prozess-ID
    Gid        uint64       // Goroutine-ID
    TraceID    string       // Trace-ID für verteiltes Tracing
    CallerName string       // Name der Aufrufer-Funktion
    CallerFile string       // Pfad der Aufrufer-Datei
    CallerLine int          // Zeilennummer des Aufrufers
}
```

## Logger-API

### Konfigurationsmethoden

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Legt den niedrigsten Log-Level fest. Nachrichten unter diesem Level werden ignoriert.

**Parameter:**

-   `level`: Der niedrigste zu verarbeitende Log-Level

**Rückgabewert:**

-   `*Logger`: Gibt sich selbst zurück, um Method-Chaining zu unterstützen

**Beispiel:**

```go title="Log-Level festlegen"
logger.SetLevel(log.InfoLevel)
logger.Debug("Dies wird nicht angezeigt")  // Ignoriert
logger.Info("Dies wird angezeigt")        // Verarbeitet
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Legt ein oder mehrere Ausgabeziele für Log-Nachrichten fest.

**Parameter:**

-   `writers`: Ein oder mehrere `io.Writer` Ausgabeziele

**Rückgabewert:**

-   `*Logger`: Gibt sich selbst zurück, um Method-Chaining zu unterstützen

**Beispiel:**

```go title="Ausgabeziel festlegen"
// Einzelausgabe
logger.SetOutput(os.Stdout)

// Mehrere Ausgaben
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

Legt einen benutzerdefinierten Formatierer für die Log-Ausgabe fest.

**Parameter:**

-   `formatter`: Ein Formatierer, der das `Format`-Interface implementiert

**Rückgabewert:**

-   `*Logger`: Gibt sich selbst zurück, um Method-Chaining zu unterstützen

**Beispiel:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Aktiviert oder deaktiviert die Aufzeichnung von Aufruferinformationen in Log-Einträgen.

**Parameter:**

-   `enable`: Ob Aufruferinformationen aktiviert werden sollen (`true` übergeben zum Aktivieren)

**Rückgabewert:**

-   `*Logger`: Gibt sich selbst zurück, um Method-Chaining zu unterstützen

**Beispiel:**

```go
logger.EnableCaller(true)
logger.Info("Dies enthält Datei:Zeilennummer-Informationen")

logger.EnableCaller(false)
logger.Info("Dies enthält keine Datei:Zeilennummer-Informationen")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

Steuert die Aufruferinformationen im Formatierer.

**Parameter:**

-   `disable`: Ob Aufruferinformationen deaktiviert werden sollen (`true` übergeben zum Deaktivieren)

**Rückgabewert:**

-   `*Logger`: Gibt sich selbst zurück, um Method-Chaining zu unterstützen

**Beispiel:**

```go
logger.Caller(false)  // Nicht deaktivieren, Aufruferinformationen anzeigen
logger.Info("Dies enthält Datei:Zeilennummer-Informationen")

logger.Caller(true)   // Aufruferinformationen deaktivieren
logger.Info("Dies enthält keine Datei:Zeilennummer-Informationen")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Legt die Stapeltiefe für Aufruferinformationen beim Wrappen des Loggers fest.

**Parameter:**

-   `depth`: Anzahl der zu überspringenden Stapelrahmen

**Rückgabewert:**

-   `*Logger`: Gibt sich selbst zurück, um Method-Chaining zu unterstützen

**Beispiel:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // Wrapper-Funktion überspringen
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Legt Präfix- oder Suffixtext für alle Log-Nachrichten fest.

**Parameter:**

-   `prefix/suffix`: Der Text, der vor/nach der Nachricht eingefügt werden soll

**Rückgabewert:**

-   `*Logger`: Gibt sich selbst zurück, um Method-Chaining zu unterstützen

**Beispiel:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // Ausgabe: [APP] Hello [END]
```

### Logging-Methoden

Alle Logging-Methoden haben zwei Varianten: eine einfache Version und eine formatierte Version.

#### Trace-Level

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

Loggt auf Trace-Level (am ausführlichsten).

**Beispiel:**

```go
logger.Trace("Detaillierte Ausführungsverfolgung")
logger.Tracef("Verarbeite Element %d von %d", i, total)
```

#### Debug-Level

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

Loggt Entwicklungsinformationen auf Debug-Level.

**Beispiel:**

```go
logger.Debug("Variablenstatus:", variable)
logger.Debugf("Benutzer %s erfolgreich authentifiziert", username)
```

#### Info-Level

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

Loggt informative Nachrichten.

**Beispiel:**

```go
logger.Info("Anwendung gestartet")
logger.Infof("Server lauscht auf Port %d", port)
```

#### Warn-Level

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

Loggt Warnnachrichten für potenzielle Problemfälle.

**Beispiel:**

```go
logger.Warn("Veraltete Funktion aufgerufen")
logger.Warnf("Speichernutzung hoch: %d%%", memoryPercent)
```

#### Error-Level

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

Loggt Fehlermeldungen.

**Beispiel:**

```go
logger.Error("Datenbankverbindung fehlgeschlagen")
logger.Errorf("Anfrageverarbeitung fehlgeschlagen: %v", err)
```

#### Fatal-Level

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

Loggt einen fatalen Fehler und ruft `os.Exit(1)` auf.

:::danger Zerstörerische Operation
`Fatal` und `Fatalf` rufen sofort nach dem Logging `os.Exit(1)` auf, um den Prozess zu beenden. Verwenden Sie dies nur bei nicht wiederherstellbaren Fehlerzuständen. `defer`-Anweisungen werden **nicht** ausgeführt.
:::

**Beispiel:**

```go
logger.Fatal("Kritischer Systemfehler")
logger.Fatalf("Server kann nicht gestartet werden: %v", err)
```

#### Panic-Level

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

Loggt eine Fehlermeldung und ruft `panic()` auf.

:::danger Zerstörerische Operation
`Panic` und `Panicf` rufen nach dem Logging `panic()` auf. Im Gegensatz zu `Fatal` kann `panic` mit `recover()` abgefangen werden, beendet aber das Programm, wenn es nicht abgefangen wird.
:::

**Beispiel:**

```go
logger.Panic("Nicht behebbarer Fehler aufgetreten")
logger.Panicf("Ungültiger Zustand: %v", state)
```

### Hilfsmethoden

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Erstellt eine Kopie des Loggers mit derselben Konfiguration.

**Rückgabewert:**

-   `*Logger`: Neue Logger-Instanz mit kopierten Einstellungen

**Beispiel:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

Erstellt einen kontextbewussten Logger, der `context.Context` als ersten Parameter akzeptiert.

**Rückgabewert:**

-   `LoggerWithCtx`: Kontextbewusste Logger-Instanz

**Beispiel:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "Kontextbewusste Nachricht")
```

## Globale Funktionen

Paket-Level-Funktionen, die den Standard-Global-Logger verwenden.

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

**Beispiel:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("Verwende globalen Logger")
```

## Log-Level

### Level-Typ

```go
type Level int8
```

### Verfügbare Level

```go
const (
    PanicLevel Level = iota  // 0 - Panik und Beenden
    FatalLevel              // 1 - Fataler Fehler und Beenden
    ErrorLevel              // 2 - Fehlerzustand
    WarnLevel               // 3 - Warnzustand
    InfoLevel               // 4 - Informative Nachricht
    DebugLevel              // 5 - Debug-Nachricht
    TraceLevel              // 6 - Ausführlichste Ablaufverfolgung
)
```

### Level-Methoden

```go
func (l Level) String() string
```

Gibt die String-Repräsentation des Levels zurück.

**Beispiel:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## Formatierer

### Format-Interface

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

Benutzerdefinierte Formatierer müssen dieses Interface implementieren.

### Standard-Formatierer

Integrierter Textformatierer mit anpassbaren Optionen.

```go
type Formatter struct {
    // Konfigurationsoptionen
}
```

### JSON-Formatierer-Beispiel

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

// Verwendung
logger.SetFormatter(&JSONFormatter{})
```

## Ausgabe-Writer

### Dateiausgabe und Rotation

```go
func GetOutputWriterHourly(filename string) io.Writer
```

Erstellt einen Writer, der Logdateien stündlich rotiert.

**Parameter:**

-   `filename`: Der Basisdateiname der Logdatei

**Rückgabewert:**

-   `io.Writer`: Rotierender Datei-Writer

**Beispiel:**

```go title="Stündliche Log-Rotation"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// Erstellt Dateien wie: app-2024010115.log, app-2024010116.log, etc.
```

### Asynchroner Writer

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

Erstellt einen asynchronen Writer für Hochleistungs-Logging.

**Parameter:**

-   `writer`: Der zugrunde liegende Writer
-   `bufferSize`: Größe des internen Puffers

**Rückgabewert:**

-   `*AsyncWriter`: Asynchrone Writer-Instanz

**Methoden:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**Beispiel:**

```go title="Asynchroner Writer"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## Kontext-Logging

### LoggerWithCtx-Interface

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

### Kontext-Funktionen

```go
func SetTrace(traceID string)
func GetTrace() string
```

Legt die Trace-ID für die aktuelle Goroutine fest und ruft sie ab.

**Beispiel:**

```go
log.SetTrace("trace-123-456")
log.Info("Diese Nachricht enthält die Trace-ID")

traceID := log.GetTrace()
fmt.Println("Aktuelle Trace-ID:", traceID)
```

## Build-Tags

Diese Bibliothek unterstützt bedingte Kompilierung mit Build-Tags:

:::info Build-Tag-Beschreibung
Build-Tags werden mit dem Parameter `go build -tags` angegeben. Verschiedene Tags ändern das Kompilierungsverhalten und die Laufzeiteigenschaften der Log-Bibliothek. Die Wahl geeigneter Tags ermöglicht ein Gleichgewicht zwischen Entwicklungsfreundlichkeit und Produktionsleistung.
:::

### Standard-Modus

```bash
go build
```

-   Volle Funktionalität aktiviert
-   Debug-Nachrichten enthalten
-   Standardleistung

### Debug-Modus

```bash
go build -tags debug
```

-   Erweiterte Debug-Informationen
-   Zusätzliche Laufzeitprüfungen
-   Detaillierte Aufruferinformationen

### Release-Modus

```bash
go build -tags release
```

-   Für Produktionsumgebung optimiert
-   Debug-Nachrichten deaktiviert
-   Automatische Log-Rotation aktiviert

### Discard-Modus

```bash
go build -tags discard
```

-   Maximale Leistung
-   Alle Logging-Operationen sind No-Ops
-    Null Overhead

### Kombinations-Modus

```bash
go build -tags "debug,discard"    # Debug und Discard
go build -tags "release,discard"  # Release und Discard
```

## Leistungsoptimierung

:::tip Leistungs-Best-Practices
Diese Bibliothek ist tiefgehend durch Objekt-Pools, vorangestellte Level-Prüfungen und asynchrones Schreiben optimiert. In Szenarien mit hohem Durchsatz wird empfohlen, asynchrone Writer und geeignete Build-Tags kombiniert zu verwenden, um die beste Leistung zu erzielen.
:::

### Objekt-Pools

Die Bibliothek verwendet intern `sync.Pool` zur Verwaltung von:

-   Log-Eintragsobjekten
-   Byte-Puffern
-   Formatierer-Puffern

Dies reduziert den Garbage-Collection-Druck in Szenarien mit hohem Durchsatz.

### Level-Prüfung

Log-Level-Prüfungen erfolgen vor teuren Operationen:

```go
// Effizient - Nachrichtenformatierung nur, wenn Level aktiviert
logger.Debugf("Ergebnis teurer Operation: %+v", expensiveCall())

// Weniger effizient, wenn Debug in Produktion deaktiviert ist
result := expensiveCall()
logger.Debug("Ergebnis:", result)
```

### Asynchrones Schreiben

Für Anwendungen mit hohem Durchsatz:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // Großer Puffer
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### Build-Tag-Optimierung

Verwenden Sie geeignete Build-Tags je nach Umgebung:

-   Entwicklung: Standard oder Debug-Tags
-   Produktion: Release-Tags
-   Leistungskritisch: Discard-Tags

## Beispiele

### Grundlegende Verwendung

```go title="Grundlegende Verwendung"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("Anwendung wird gestartet")
    log.Warn("Dies ist eine Warnung")
    log.Error("Dies ist ein Fehler")
}
```

### Benutzerdefinierter Logger

```go title="Benutzerdefinierte Konfiguration"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // Logger konfigurieren
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  // Aufruferinformationen deaktivieren
    logger.SetPrefixMsg("[MyApp] ")

    // Ausgabe in Datei festlegen
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("Benutzerdefinierter Logger konfiguriert")
    logger.Debug("Debug-Informationen mit Aufrufer")
}
```

### Hochleistungs-Logging

```go title="Hochleistungs-Logging"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Rotierenden Datei-Writer erstellen
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // Asynchronen Writer für bessere Leistung verwenden
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // Debug-Logs in Produktion überspringen

    // Hochdurchsatz-Logging
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### Kontextbewusstes Logging

```go title="Kontextbewusstes Logging"
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

    ctxLogger.Info(ctx, "Verarbeite Benutzeranfrage")
    ctxLogger.Debug(ctx, "Validierung abgeschlossen")
}
```

### Benutzerdefinierter JSON-Formatierer

```go title="Benutzerdefinierter JSON-Formatierer"
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
    logger.Caller(true)  // Aufruferinformationen deaktivieren
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("JSON-formatierte Nachricht")
}
```

## Fehlerbehandlung

:::warning Hinweis
Aus Leistungsgründen geben die meisten Logger-Methoden keine Fehler zurück. Wenn das Schreiben fehlschlägt, werden Logs stillschweigend verworfen. Wenn Fehlerbeawareness erforderlich ist, verwenden Sie einen benutzerdefinierten Writer.
:::

Wenn Sie Fehlerbehandlung für Ausgabeoperationen benötigen, implementieren Sie einen benutzerdefinierten Writer:

```go title="Fehler-capturer Writer"
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

## Thread-Sicherheit

:::tip Gleichgangssicherheit
Alle Methoden von `Logger`-Instanzen sind threadsicher und können ohne zusätzliche Synchronisationsmechanismen in mehreren Goroutines gleichzeitig verwendet werden. Beachten Sie jedoch, dass einzelne `Entry`-Objekte **nicht** threadsicher sind und für den einmaligen Gebrauch bestimmt sind.
:::

---

## 🌍 Mehrsprachige Dokumentation

Diese Dokumentation ist in mehreren Sprachen verfügbar:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/API.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/API.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/API.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/API.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/API.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/API.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/API.md)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/API.md)
-   [🇩🇪 Deutsch](API.md) (Aktuell)

---

**LazyGophers Log vollständige API-Referenz - Bauen Sie bessere Anwendungen mit erstklassigem Logging! 🚀**
