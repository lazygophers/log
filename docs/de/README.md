---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Eine Hochleistungs- und flexible Go-Logging-Bibliothek. Basierend auf zap mit umfangreichen Funktionen und einem einfachen API.

## 📖 Sprachen

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 Vereinfachtes Chinesisch](README.md) (Aktuell)
-   [🇹🇼 Traditionelles Chinesisch](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)

## ✨ Funktionen

-   **🚀 Hochleistung**: Object Pooling und bedingtes Feld-Recording basierend auf zap
-   **📊 Verschiedene Logstufen**: Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Flexible Konfiguration**:
    -   Steuerung der Logstufen
    -   Aufzeichnung von Caller-Informationen
    -   Trace-Informationen (inklusive Goroutine-ID)
    -   Custom-Präfix und Suffix
    -   Custom-Output-Ziele (Konsole, Datei, etc.)
    -   Log-Format-Optionen
-   **🔄 Datei-Rotation**: Hourly-Logfile-Rotation
-   **🔌 Zap-Kompatibilität**: Seamless Integration mit zap WriteSyncer
-   **🎯 Einfaches API**: Klare API, einfach zu benutzen

## 🚀 Quick Start

### Installation

:::tip Installation
```bash
go get github.com/lazygophers/log
```
:::

### Grundlegende Verwendung

```go title="Quick Start"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Verwenden Sie den globalen Standard-Logger
    log.Debug("Debug-Nachricht")
    log.Info("Info-Nachricht")
    log.Warn("Warn-Nachricht")
    log.Error("Fehler-Nachricht")

    // Mit Format-Ausgabe
    log.Infof("Benutzer %s hat sich erfolgreich angemeldet", "admin")

    // Custom-Konfiguration
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp] ")

    customLogger.Info("Das ist ein Log vom benutzerdefinierten Logger")
}
```

## 📚 Ausführliche Verwendung

### Logger mit Datei-Output

```go title="Datei-Output-Konfiguration"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Erstellen Sie einen Logger mit Datei-Output
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Debug-Log mit Caller-Informationen")
    logger.Info("Info-Log mit Trace-Informationen")
}
```

### Logstufen-Steuerung

```go title="Logstufen-Steuerung"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Nur warn und höhere Stufen werden aufgezeichnet
    logger.Debug("Dies wird nicht aufgezeichnet")  // Ignoriert
    logger.Info("Dies wird nicht aufgezeichnet")   // Ignoriert
    logger.Warn("Dies wird aufgezeichnet")         // Aufgezeichnet
    logger.Error("Dies wird aufgezeichnet")        // Aufgezeichnet
}
```

## 🎯 Anwendungsfälle

### Einsatzszenarien

-   **Web Services und API-Backends**: Request-Tracking, strukturiertes Logging, Performance-Monitoring
-   **Microservices-Architektur**: Verteiltes Tracking (TraceID), einheitliches Log-Format, hoher Durchsatz
-   **Command-Line-Tools**: Stufen-Steuerung, klare Ausgabe, Fehlerberichte
-   **Batch-Aufgabentasks**: Datei-Rotation, lange Laufzeit, Ressourcenoptimierung

### Besondere Vorteile

-   **Object Pooling-Optimierung**: Wiederverwendung von Entry- und Buffer-Objekten, Reduzierung von GC-Druck
-   **Asynchrone Schreibvorgänge**: Keine Blockierung bei Hochdurchsatz-Szenarien (10000+ Logs/Sekunde)
-   **TraceID-Unterstützung**: Request-Tracking in verteilten Systemen, Integration mit OpenTelemetry
-   **Zero-Config-Start**: Sofort einsatzbereit, schrittweise Konfiguration

## 🔧 Konfigurationsoptionen

:::note Konfigurationsoptionen
Alle diese Methoden unterstützen Chain-Calling und können für die Erstellung benutzerdefinierter Logger kombiniert werden.
:::

### Logger-Konfiguration

| Methode | Beschreibung | Standard |
| --- | --- | --- |
| `SetLevel(level)` | Minimale Logstufe setzen | `DebugLevel` |
| `EnableCaller(enable)` | Caller-Informationen aktivieren/deaktivieren | `false` |
| `EnableTrace(enable)` | Trace-Informationen aktivieren/deaktivieren | `false` |
| `SetCallerDepth(depth)` | Caller-Depth setzen | `2` |
| `SetPrefixMsg(prefix)` | Log-Präfix setzen | `""` |
| `SetSuffixMsg(suffix)` | Log-Suffix setzen | `""` |
| `SetOutput(writers...)` | Output-Ziele setzen | `os.Stdout` |

### Logstufen

| Stufe | Beschreibung |
| --- | --- |
| `TraceLevel` | Am detailliertesten, für detailliertes Tracking |
| `DebugLevel` | Debug-Informationen |
| `InfoLevel` | Allgemeine Informationen |
| `WarnLevel` | Warnmeldungen |
| `ErrorLevel` | Fehlermeldungen |
| `FatalLevel` | Kritische Fehler (ruft os.Exit(1) auf) |
| `PanicLevel` | Panic-Fehler (ruft panic() auf) |

## 🏗️ Architektur

### Kernkomponenten

-   **Logger**: Haupt-Logger-Struktur mit konfigurierbaren Optionen
-   **Entry**: Einzelne Log-Records mit umfassenden Feldern
-   **Level**: Definition von Logstufen und Utility-Funktionen
-   **Format**: Log-Format-Schnittstelle und -Implementierung

### Leistungsoptimierung

-   **Object Pooling**: Wiederverwendung von Entry-Objekten zum Reduzieren der Speicherzuweisung
-   **Bedingtes Recording**: Nur teure Felder aufzeichnen, wenn erforderlich
-   **Schnelles Level-Checking**: Level am äußeren Rand prüfen
-   **Lock-free Design**: Die meisten Operationen benötigen keine Sperren

## ❓ Häufig gestellte Fragen

### Wie wähle ich die richtige Logstufe?

-   **Entwicklungsumgebung**: Verwenden Sie `DebugLevel` oder `TraceLevel` für detaillierte Informationen
-   **Produktionsumgebung**: Verwenden Sie `InfoLevel` oder `WarnLevel` zum Reduzieren des Overheads
-   **Performance-Tests**: Verwenden Sie `PanicLevel` zum Deaktivieren aller Logs

### Wie optimiere ich die Leistung in der Produktion?

:::warning Hinweis
Bei Hochdurchsatz-Szenarien empfiehlt es sich, asynchrone Schreibvorgänge und geeignete Logstufen zu kombinieren.
:::

1. Verwenden Sie `AsyncWriter` für asynchrone Schreibvorgänge:

```go title="Asynchrone Schreibvorgänge"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Passen Sie die Logstufen an, um unnötige Logs zu vermeiden:

```go title="Stufen-Optimierung"
logger.SetLevel(log.InfoLevel)  // Debug und Trace überspringen
```

3. Verwenden Sie conditionales Logging zum Reduzieren des Overheads:

```go title="Conditionales Logging"
if logger.Level >= log.DebugLevel {
    logger.Debug("Detaillierte Debug-Informationen")
}
```

### Was ist der Unterschied zwischen `Caller` und `EnableCaller`?

-   **`EnableCaller(enable bool)`**: Steuert, ob der Logger Caller-Informationen sammelt
    -   `EnableCaller(true)` aktiviert die Caller-Informationssammlung
-   **`Caller(disable bool)`**: Steuert, ob der Formatter Caller-Informationen ausgibt
    -   `Caller(true)` deaktiviert die Caller-Informationsausgabe

Für die globale Steuerung empfiehlt es sich, `EnableCaller` zu verwenden.

### Wie implementiere ich einen Custom-Formatter?

Implementieren Sie die `Format`-Schnittstelle:

```go title="Custom-Formatter"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Verwandte Dokumentation

-   [📚 API-Dokumentation](API.md) - Vollständige API-Referenz
-   [🤝 Kontributionsleitfaden](/de/CONTRIBUTING) - Wie man beiträgt
-   [📋 Änderungsprotokoll](/de/CHANGELOG) - Versionshistorie
-   [🔒 Sicherheitsrichtlinie](/de/SECURITY) - Sicherheits-Leitlinien
-   [📜 Verhaltenskodex](/de/CODE_OF_CONDUCT) - Community-Leitlinien

## 🚀 Hilfe bekommen

-   **GitHub Issues**: [Bugs melden oder Funktionen anfordern](https://github.com/lazygophers/log/issues)
-   **GoDoc**: [API-Dokumentation](https://pkg.go.dev/github.com/lazygophers/log)
-   [✓ Beispiele](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Lizenz

Dieses Projekt ist unter der MIT-Lizenz lizenziert - Weitere Informationen finden Sie in der [LICENSE](/de/LICENSE) Datei.

## 🤝 Beitrag

Beiträge sind willkommen! Weitere Informationen finden Sie im [Kontributionsleitfaden](/de/CONTRIBUTING).

---

**lazygophers/log** ist eine Haupt-Logging-Lösung für Go-Entwickler, die Leistung und Einfachheit priorisieren. Egal ob Sie an kleinen Utilitys arbeiten oder an großen verteilten Systemen - diese Bibliothek bietet ein gutes Gleichgewicht zwischen Funktionen und Benutzbarkeit.
