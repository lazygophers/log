---
pageType: home

hero:
    name: LazyGophers Log
    text: Hochleistungs-, flexible Go-Protokollierungsbibliothek
    tagline: Basierend auf zap, bietet reiche Funktionen und eine einfache API
    actions:
        - theme: brand
          text: Schnellstart
          link: /API
        - theme: alt
          text: GitHub ansehen
          link: https://github.com/lazygophers/log

features:
    - title: "Hohe Leistung"
      details: Basierend auf zap mit Objektpooling und bedingter Feldaufzeichnung für optimale Leistung
      icon: 🚀
    - title: "Reichhaltige Protokollierungsstufen"
      details: Unterstützt Trace-, Debug-, Info-, Warn-, Error-, Fatal- und Panic-Stufen
      icon: 📊
    - title: "Flexible Konfiguration"
      details: Anpassbare Protokollierungsstufe, Anruferinformationen, Ablaufverfolgung, Präfixe, Suffixe und Ausgabeziele
      icon: ⚙️
    - title: "Dateirotation"
      details: Integrierte stündliche Protokolldateirotation
      icon: 🔄
    - title: "Zap-Kompatibilität"
      details: Nahtlose Integration mit zap WriteSyncer
      icon: 🔌
    - title: "Einfache API"
      details: Klare API ähnlich der Standardprotokollierungsbibliothek, einfach zu verwenden und zu integrieren
      icon: 🎯
---

## Schnellstart

### Installation

```bash
go get github.com/lazygophers/log
```

### Grundlegende Verwendung

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Standardmäßiger globaler Logger verwenden
    log.Debug("Debug-Information")
    log.Info("Allgemeine Information")
    log.Warn("Warnung")
    log.Error("Fehler")

    // Formatierungsausgabe verwenden
    log.Infof("Benutzer %s erfolgreich angemeldet", "admin")

    // Benutzerdefinierte Konfiguration
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Dies ist ein Protokoll vom benutzerdefinierten Logger")
}
```

## Dokumentation

-   [API-Referenz](API.md) - Vollständige API-Dokumentation
-   [Änderungsprotokoll](/de/CHANGELOG) - Versionshistorie
-   [Beitragsleitfaden](/de/CONTRIBUTING) - Wie beitragen
-   [Sicherheitsrichtlinie](/de/SECURITY) - Sicherheitsleitfaden
-   [Verhaltenskodex](/de/CODE_OF_CONDUCT) - Gemeinschaftsrichtlinien

## Leistungsvergleich

| Funktion       | lazygophers/log | zap | logrus | Standardprotokollierung |
| ---------- | --------------- | --- | ------ | -------- |
| Leistung       | Hoch              | Hoch  | Mittel     | Niedrig       |
| API-Einfachheit | Hoch              | Mittel  | Hoch     | Hoch       |
| Funktionsreichtum | Mittel              | Hoch  | Hoch     | Niedrig       |
| Flexibilität     | Mittel              | Hoch  | Hoch     | Niedrig       |
| Lernkurve   | Niedrig              | Mittel  | Mittel     | Niedrig       |

## Lizenz

Dieses Projekt ist unter der MIT-Lizenz lizenziert - siehe Datei [LICENSE](/de/LICENSE).
