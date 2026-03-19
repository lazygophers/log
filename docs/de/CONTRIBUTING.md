---
titleSuffix: ' | LazyGophers Log'
---
# 🤝 Zu LazyGophers Log beitragen

Wir freuen uns sehr über Ihren Beitrag! Wir möchten es so einfach und transparent wie möglich machen, zu LazyGophers Log beizutragen, sei es:

-   🐛 Fehler melden
-   💬 Über den aktuellen Status des Codes diskutieren
-   ✨ Funktionsanforderungen einreichen
-   🔧 Lösungen vorschlagen
-   🚀 Neue Funktionen implementieren

## 📋 Inhaltsverzeichnis

-   [Verhaltenskodex](#-verhaltenskodex)
-   [Entwicklungsprozess](#-entwicklungsprozess)
-   [Erste Schritte](#-erste-schritte)
-   [Pull-Request-Prozess](#-pull-request-prozess)
-   [Codierungsstandards](#-codierungsstandards)
-   [Testrichtlinien](#-testrichtlinien)
-   [Build-Tag-Anforderungen](#️-build-tag-anforderungen)
-   [Dokumentation](#-dokumentation)
-   [Issue-Richtlinien](#-issue-richtlinien)
-   [Leistungsüberlegungen](#-leistungsüberlegungen)
-   [Sicherheitsrichtlinien](#-sicherheitsrichtlinien)
-   [Community](#-community)

## 📜 Verhaltenskodex

Dieses Projekt und alle Teilnehmer unterliegen unserem [Verhaltenskodex](/de/CODE_OF_CONDUCT). Durch die Teilnahme erklären Sie sich mit den Regeln einverstanden.

## 🔄 Entwicklungsprozess

Wir verwenden GitHub, um Code zu hosten, Issues und Funktionsanforderungen zu verfolgen und Pull Requests zu akzeptieren.

### Workflow

:::note Entwicklungsprozessüberblick
1. Repository **Forken**
2. Fork **klonen**
3. Feature-Branch von `master` **erstellen**
4. **Änderungen** vornehmen
5. In allen Build-Tags **testen**
6. Pull Request **einreichen**
:::

## 🚀 Erste Schritte

### Voraussetzungen

-   **Go 1.21+** - [Go installieren](https://golang.org/doc/install)
-   **Git** - [Git installieren](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
-   **Make** (optional aber empfohlen)

### Lokale Entwicklungseinrichtung

```bash title="Repository klonen und Entwicklungsumgebung einrichten"
# 1. Repository auf GitHub forken
# 2. Fork klonen
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. Upstream-Remote hinzufügen
git remote add upstream https://github.com/lazygophers/log.git

# 4. Abhängigkeiten installieren
go mod tidy

# 5. Installation verifizieren
make test-quick
```

### Umgebungseinrichtung

:::info Umgebungskonfiguration
Stellen Sie sicher, dass die Go-Umgebungsvariablen korrekt konfiguriert sind und die empfohlenen Entwicklungstools installiert sind, um die beste Entwicklungserfahrung zu erzielen.
:::

```bash title="Umgebungseinrichtung"
# Go-Umgebung einrichten (falls noch nicht geschehen)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Optional: Nützliche Tools installieren
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 Pull-Request-Prozess

### Vor der Einreichung

1.  **Suchen** Sie nach vorhandenen PRs, um Duplikate zu vermeiden
2.  **Testen** Sie Ihre Änderungen in allen Build-Konfigurationen
3.  **Dokumentieren** Sie alle Breaking Changes
4.  **Aktualisieren** Sie die relevante Dokumentation
5.  **Fügen** Sie Tests für neue Funktionen hinzu

### PR-Checkliste

:::warning Bitte überprüfen Sie alle Punkte vor der PR-Einreichung
PRs, die die Checklistenanforderungen nicht erfüllen, werden nicht zusammengeführt.
:::

-   [ ] **Codequalität**

    -   [ ] Code folgt dem Projektstil-Leitfaden
    -   [ ] Keine neuen Lint-Warnungen
    -   [ ] Korrekte Fehlerbehandlung
    -   [ ] Effiziente Algorithmen und Datenstrukturen

-   [ ] **Tests**

    -   [ ] Alle vorhandenen Tests bestehen: `make test`
    -   [ ] Neue Tests für neue Funktionen hinzugefügt
    -   [ ] Testabdeckung gehalten oder verbessert
    -   [ ] Alle Build-Tags getestet: `make test-all`

-   [ ] **Dokumentation**

    -   [ ] Code hat angemessene Kommentare
    -   [ ] API-Dokumentation aktualisiert (falls erforderlich)
    -   [ ] README aktualisiert (falls erforderlich)
    -   [ ] Mehrsprachige Dokumentation aktualisiert (falls benutzerorientiert)

-   [ ] **Build-Kompatibilität**
    -   [ ] Standard-Modus: `go build`
    -   [ ] Debug-Modus: `go build -tags debug`
    -   [ ] Release-Modus: `go build -tags release`
    -   [ ] Discard-Modus: `go build -tags discard`
    -   [ ] Kombinationsmodi getestet

### PR-Vorlage

Verwenden Sie beim Einreichen eines Pull Requests unsere [PR-Vorlage](https://github.com/lazygophers/log/blob/main/.github/pull_request_template.md).

## 📏 Codierungsstandards

### Go-Stil-Leitfaden

:::tip Go-Codierungsstandards
Wir folgen dem Standard-GO-Stil-Leitfaden mit einigen Ergänzungen. Stellen Sie sicher, dass die Codeformatierung `go fmt` und `goimports` Prüfungen besteht.
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

### Benennungskonventionen

-   **Pakete**: Kurz, kleingeschrieben, wenn möglich ein einzelnes Wort
-   **Funktionen**: CamelCase, beschreibend
-   **Variablen**: camelCase für lokale, PascalCase für exportierte
-   **Konstanten**: PascalCase für exportierte, camelCase für nicht exportierte
-   **Schnittstellen**: Enden normalerweise auf "er" (z. B. `Writer`, `Formatter`)

### Code-Organisation

```
project/
├── docs/           # Mehrsprachige Dokumentation
├── .github/        # GitHub-Vorlagen und Workflows
├── logger.go       # Haupt-Logger-Implementierung
├── entry.go        # Log-Eintragsstruktur
├── level.go        # Log-Level
├── formatter.go    # Log-Formatierung
├── output.go       # Ausgabemanagement
└── *_test.go      # Tests zusammen mit Quellcode
```

### Fehlerbehandlung

:::tip Best Practices für Fehlerbehandlung
Bibliothekscode sollte Fehler zurückgeben, nicht panic, und dem Aufrufer die Entscheidung überlassen, wie mit Ausnahmen umgegangen wird.
:::

```go title="Fehlerbehandlungsbeispiel"
// ✅ Empfohlen: Fehler zurückgeben, nicht panic
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ Vermeiden: Panic in Bibliothekscode verwenden
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // Tun Sie dies nicht
    }
    return &Logger{...}
}
```

## 🧪 Testrichtlinien

### Teststruktur

```go title="Tabellengesteuertes Testbeispiel"
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
            // Test-Implementierung
        })
    }
}
```

### Abdeckungsanforderungen

:::warning Abdeckungs-Anforderung
PRs mit einer Abdeckung unter 90% für neuen Code werden die CI-Prüfungen nicht bestehen.
:::

-   **Minimum**: 90% Abdeckung für neuen Code
-   **Ziel**: 95%+ Gesamtüberdeckung
-   **Alle Build-Tags** müssen Abdeckung aufrechterhalten
-   Verwenden Sie `make coverage-all` zur Verifizierung

### Testbefehle

```bash title="Tests ausführen"
# Schnelle Tests in allen Build-Tags
make test-quick

# Vollständige Testsuite mit Abdeckung
make test-all

# Abdeckungsbericht
make coverage-html

# Benchmarks
make benchmark
```

## 🏗️ Build-Tag-Anforderungen

:::warning Build-Kompatibilität
Alle Änderungen müssen mit unserem Build-Tag-System kompatibel sein. Code, der nicht alle Build-Tag-Tests besteht, wird nicht zusammengeführt.
:::

### Unterstützte Build-Tags

-   **Standard** (`go build`): Vollständige Funktionalität
-   **Debug** (`go build -tags debug`): Erweiterte Debugging-Funktionen
-   **Release** (`go build -tags release`): Produktoptimierung
-   **Discard** (`go build -tags discard`): Maximale Leistung

### Build-Tag-Tests

:::info Build-Tag-Beschreibung
Das Projekt verwendet Build-Tags für bedingte Kompilierung, wobei verschiedene Tags verschiedenen Ausführungsmodi entsprechen. Testen Sie vor dem Einreichen unter allen Tags.
:::

```bash title="Build-Tag-Tests"
# Jede Build-Konfiguration testen
make test-default
make test-debug
make test-release
make test-discard

# Kombinationen testen
make test-debug-discard
make test-release-discard

# Alle auf einmal testen
make test-all
```

### Build-Tag-Richtlinien

```go
//go:build debug
// +build debug

package log

// Debug-spezifische Implementierung
```

## 📚 Dokumentation

### Code-Dokumentation

-   **Alle exportierten Funktionen** müssen klare Kommentare haben
-   **Komplexe Algorithmen** benötigen Erklärungen
-   **Beispiele** für nicht triviale Verwendung
-   **Thread-Sicherheit** Hinweise (falls zutreffend)

```go
// SetLevel legt das minimale Logging-Level fest.
// Logs unter diesem Level werden ignoriert.
// Diese Methode ist Thread-sicher.
//
// Example:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // Wird nicht ausgegeben
//   logger.Info("visible")   // Wird ausgegeben
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### README-Updates

Fügen Sie beim Hinzufügen von Funktionen Folgendes hinzu:

-   Haupt-README.md
-   Alle sprachspezifischen READMEs in `docs/`
-   Codebeispiele
-   Funktionsliste

## 🐛 Issue-Richtlinien

### Fehlerberichte

Verwenden Sie die [Fehlerbericht-Vorlage](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/bug_report.md) und fügen Sie Folgendes hinzu:

-   **Klare Problembeschreibung**
-   **Schritte zur Reproduktion**
-   **Erwartetes vs. tatsächliches Verhalten**
-   **Umgebungsdetails** (Betriebssystem, Go-Version, Build-Tags)
-   **Minimaler Code-Beispiel**

### Funktionsanforderungen

Verwenden Sie die [Funktionsanforderung-Vorlage](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/feature_request.md) und fügen Sie Folgendes hinzu:

-   **Klare Funktionsmotivation**
-   **Vorgeschlagenes API**-Design
-   **Implementierungsüberlegungen**
-   **Breaking-Change-Analyse**

### Fragen

Verwenden Sie die [Frage-Vorlage](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/question.md) für:

-   Nutzungsprobleme
-   Konfigurationshilfe
-   Best Practices
-   Integrationsanleitung

## 🚀 Leistungsüberlegungen

### Benchmarking

Führen Sie immer Benchmarks für leistungssensible Änderungen durch:

```bash title="Benchmarks ausführen"
# Benchmarks ausführen
go test -bench=. -benchmem

// Vorher-Nachher-Leistung vergleichen
go test -bench=. -benchmem > before.txt
// Änderungen vornehmen
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### Leistungsrichtlinien

:::tip Leistungsoptimierungspunkte
Dies ist eine Leistungssensible Logging-Bibliothek. Alle Änderungen sollten die Auswirkungen auf den Hot-Pfad berücksichtigen.
:::

-   **Minimieren** Sie Speicherzuweisungen im Hot-Pfad
-   **Verwenden Sie Objekt-Pools** für häufig erstellte Objekte
-   **Frühe Rückgabe** für deaktivierte Log-Level
-   **Vermeiden Sie Reflektion** in leistungskritischem Code
-   **Profiling vor Optimierung**

### Speichermanagement

```go
// ✅ Empfohlen: Objekt-Pool verwenden
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

## 🔒 Sicherheitsrichtlinien

### Sensible Daten

:::warning Sicherheitswarnung
Das Lecken sensibler Daten in Protokollen kann zu schweren Sicherheitsvorfällen führen. Bitte befolgen Sie unbedingt die folgenden Standards.
:::

-   **Protokollieren Sie niemals** Passwörter, Tokens oder sensible Daten
-   **Bereinigen Sie** Benutzereingaben aus Protokollnachrichten
-   **Vermeiden Sie** das Protokollieren ganzer Request/Response-Body
-   **Verwenden Sie** strukturiertes Logging für bessere Kontrolle

```go
// ✅ Empfohlen
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ Vermeiden
logger.Infof("User login: %+v", userRequest) // Kann Passwörter enthalten
```

### Abhängigkeiten

-   Halten Sie Abhängigkeiten **aktuell**
-   **Überprüfen Sie sorgfältig** neue Abhängigkeiten
-   **Minimieren Sie** externe Abhängigkeiten
-   **Verwenden Sie** `go mod verify` zur Integritätsprüfung

## 👥 Community

### Hilfe erhalten

-   📖 [Dokumentation](README.md)
-   💬 [GitHub-Diskussionen](https://github.com/lazygophers/log/discussions)
-   🐛 [Issue-Tracker](https://github.com/lazygophers/log/issues)
-   📧 E-Mail: support@lazygophers.com

### Kommunikationsrichtlinien

-   **Respektvoll und inklusiv bleiben**
-   **Vor dem Fragen suchen**
-   **Kontext bereitstellen, wenn um Hilfe gebeten wird**
-   **Andere helfen, wenn Sie können**
-   **Befolgen Sie** den [Verhaltenskodex](/de/CODE_OF_CONDUCT)

## 🎯 Anerkennung

Beitragende werden auf folgende Weise anerkannt:

-   **README-Beitragende**-Abschnitt
-   **Release-Notes**-Erwähnungen
-   **GitHub-Beitragende**-Diagramm
-   **Community-Dank**-Beiträge

## 📝 Lizenz

Durch einen Beitrag erklären Sie sich damit einverstanden, dass Ihr Beitrag unter der MIT-Lizenz lizenziert wird.

---

## 🌍 Mehrsprachige Dokumentation

Dieses Dokument ist in mehreren Sprachen verfügbar:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CONTRIBUTING.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/CONTRIBUTING.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CONTRIBUTING.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CONTRIBUTING.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CONTRIBUTING.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CONTRIBUTING.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CONTRIBUTING.md)
-   [🇩🇪 Deutsch](/de/CONTRIBUTING)（aktuell）

---

**Danke, dass Sie zu LazyGophers Log beitragen!🚀**
