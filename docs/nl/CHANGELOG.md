---
titleSuffix: " | LazyGophers Log"
---

# 📋 Wijzigingslog

Alle belangrijke wijzigingen in dit project worden in dit bestand bijgehouden.

Het formaat is gebaseerd op [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), en dit project volgt [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Niet Gepubliceerd]

### Toegevoegd

-   Uitgebreide meertalige documentatie (7 talen)
-   GitHub issue-sjablonen (bugrapporten, functieverzoeken, vragen)
-   Pull Request-sjabloon met build tag compatibiliteitscontroles
-   Meertalige bijdragergids
-   Gedragscode met uitvoeringsrichtlijnen
-   Beveiligingsbeleid met kwetsbaarheidsrapportageproces
-   Volledige API-documentatie met voorbeelden
-   Professionele projectstructuur en sjablonen

### Gewijzigd

-   README verbeterd met uitgebreide functiedocumentatie
-   Testdekking verbeterd voor alle build tag-configuraties
-   Projectstructuur bijgewerkt voor beter onderhoud

### Documentatie

-   Meertalige ondersteuning toegevoegd voor alle belangrijke documenten
-   Uitgebreide API-referentie gemaakt
-   Richtlijnen voor bijdrage-werkstromen vastgesteld
-   Beveiligingsrapportageproces geïmplementeerd

## [1.0.0] - 2024-01-01

### Toegevoegd

-   Kern log-functionaliteit met meerdere logniveaus (Trace, Debug, Info, Warn, Error, Fatal, Panic)
-   Thread-safe logger implementatie met object pooling
-   Build tag ondersteuning (standaard, debug, release, discard modi)
-   Aanpasbare formatter interface met standaard tekstformatter
-   Multi-writer uitvoerondersteuning
-   Async schrijven functies voor hoge doorvoerscenario's
-   Automatisch uurlijkse logbestandsrotatie
-   Context-bewust loggen met goroutine ID en trace ID tracking
-   Caller informatie met configureerbare stack diepte
-   Globale pakket-niveau gemaksfuncties
-   Zap logger integratieondersteuning

### Prestaties

-   Object pooling voor entry-objecten en buffers met `sync.Pool`
-   Vroegtijdige niveaucontroles om dure bewerkingen te vermijden
-   Async writer voor niet-blokkerende log-schrijfbewerkingen
-   Build tag optimalisaties voor verschillende omgevingen

### Build Tags

-   **Standaard**: Volledige functionaliteit met debugberichten
-   **Debug**: Uitgebreide debug-informatie en caller-details
-   **Release**: Productieoptimalisaties, debugberichten uitgeschakeld
-   **Discard**: Maximale prestaties, no-op log-bewerkingen

### Kernfuncties

-   **Logger**: Hoofd logger structuur met configureerbaar niveau, uitvoer, formatter
-   **Entry**: Log record structuur met uitgebreide metagegevens
-   **Levels**: Zeven logniveaus van Panic (hoogste) tot Trace (laagste)
-   **Formatters**: Insteekformaat systeem
-   **Writers**: Bestandsrotatie en async schrijfondersteuning
-   **Context**: Goroutine ID en gedistribueerde tracing ondersteuning

### API Hoogtepunten

-   Vloeiende configuratie API met method chaining
-   Eenvoudige en geformatteerde logmethoden (`.Info()` en `.Infof()`)
-   Logger klonen voor geïsoleerde configuraties
-   Context-bewust loggen met `CloneToCtx()`
-   Voorvoegsel- en achtervoegselbericht aanpassing
-   Caller informatie schakelaar

### Tests

-   Uitgebreide testsuite met 93,5% dekking
-   Multi-build tag testondersteuning
-   Geautomatiseerde test workflows
-   Prestatiebenchmarks

## [0.9.0] - 2023-12-15

### Toegevoegd

-   Initiële projectstructuur
-   Basis log-functionaliteit
-   Niveau-gebaseerde filtering
-   Bestandsuitvoerondersteuning

### Gewijzigd

-   Prestaties verbeterd via object pooling
-   Foutafhandeling verbeterd

## [0.8.0] - 2023-12-01

### Toegevoegd

-   Multi-writer ondersteuning
-   Aanpasbare formatter interface
-   Async schrijf-functionaliteit

### Opgelost

-   Geheugenlek in hoge doorvoerscenario's
-   Race condition bij gelijktijdige toegang

## [0.7.0] - 2023-11-15

### Toegevoegd

-   Build tag ondersteuning voor conditionele compilatie
-   Trace en Debug niveau loggen
-   Caller informatie tracking

### Gewijzigd

-   Geheugentoewijzingspatronen geoptimaliseerd
-   Thread-veiligheid verbeterd

## [0.6.0] - 2023-11-01

### Toegevoegd

-   Log rotatie functionaliteit
-   Context-bewust loggen
-   Goroutine ID tracking

### Afgekeurd

-   Oude configuratiemethoden (wordt verwijderd in v1.0.0)

## [0.5.0] - 2023-10-15

### Toegevoegd

-   JSON formatter
-   Meerdere uitvoerdoelen
-   Prestatiebenchmarks

### Gewijzigd

-   Kern log engine herstructureerd
-   API consistentie verbeterd

### Verwijderd

-   Oude logmethoden

## [0.4.0] - 2023-10-01

### Toegevoegd

-   Fatal en Panic niveau loggen
-   Globale pakketfuncties
-   Configuratievalidatie

### Opgelost

-   Uitvoersynchronisatieproblemen
-   Geheugengebruik optimalisatie

## [0.3.0] - 2023-09-15

### Toegevoegd

-   Aanpasbare logniveaus
-   Formatter interface
-   Thread-veilige bewerkingen

### Gewijzigd

-   API ontwerp versimpeld
-   Documentatie uitgebreid

## [0.2.0] - 2023-09-01

### Toegevoegd

-   Bestandsuitvoerondersteuning
-   Niveau-gebaseerde filtering
-   Basis formatteeropties

### Opgelost

-   Prestatieknelpunten
-   Geheugenlekken

## [0.1.0] - 2023-08-15

### Toegevoegd

-   Eerste publicatie
-   Basis console loggen
-   Eenvoudige niveau ondersteuning (Info, Warn, Error)
-   Kern logger structuur

## Samenvatting Versiegeschiedenis

| Versie | Publicatiedatum | Belangrijkste functies |
| ------ | --------------- | ---------------------- |
| 1.0.0  | 2024-01-01      | Volledig logsysteem, build tags, async schrijven, uitgebreide documentatie |
| 0.9.0  | 2023-12-15      | Prestatieverbeteringen, object pooling |
| 0.8.0  | 2023-12-01      | Multi-writer, async schrijven, aanpasbare formatters |
| 0.7.0  | 2023-11-15      | Build tags, Trace/Debug niveaus, caller informatie |
| 0.6.0  | 2023-11-01      | Log rotatie, context loggen, goroutine tracking |
| 0.5.0  | 2023-10-15      | JSON formatter, meerdere uitvoeren, benchmarks |
| 0.4.0  | 2023-10-01      | Fatal/Panic niveaus, globale functies |
| 0.3.0  | 2023-09-15      | Aanpasbare niveaus, formatter interface |
| 0.2.0  | 2023-09-01      | Bestandsuitvoer, niveau filtering |
| 0.1.0  | 2023-08-15      | Eerste publicatie, basis console loggen |

## Migratiegids

### Migreren van v0.9.x naar v1.0.0

#### Breaking Changes

-   Geen - v1.0.0 is achterwaarts compatibel met v0.9.x

#### Nieuw Beschikbare Functies

-   Uitgebreide build tag ondersteuning
-   Volledige documentatie
-   Professionele projectsjablonen
-   Beveiligingsrapportageproces

#### Aanbevolen Updates

```go
// Oude manier (nog steeds ondersteund)
logger := log.New()
logger.SetLevel(log.InfoLevel)

// Aanbevolen nieuwe manier, met method chaining
logger := log.New().
    SetLevel(log.InfoLevel).
    Caller(true).
    SetPrefixMsg("[MyApp] ")
```

### Migreren van v0.8.x naar v0.9.x

#### Breaking Changes

-   Verouderde configuratiemethoden verwijderd
-   Interne bufferbeheer gewijzigd

#### Migratiestappen

1. Werk importpaden indien nodig bij
2. Vervang verouderde methoden:

    ```go
    // Oud (verouderd)
    logger.SetOutputFile("app.log")

    // Nieuw
    file, _ := os.Create("app.log")
    logger.SetOutput(file)
    ```

### Migreren van v0.5.x en eerder

#### Belangrijkste Wijzigingen

-   Volledige API herontwerp voor betere consistentie
-   Prestatieverbeteringen via object pooling
-   Nieuw build tag systeem

### Vereiste Migratie

-   Werk alle log-aanroepen bij naar de nieuwe API
-   Beoordeel en update formatter implementaties
-   Test met de nieuwe build tag configuraties

## Ontwikkelpijlpalen

### 🎯 v1.1.0 Routekaart (Gepland)

-   [] Gestructureerd loggen met sleutel-waarde paren
-   [ ] Log sampling voor hoge volume scenario's
-   [ ] Plugin systeem voor aangepaste uitvoer
-   [ ] Uitgebreide prestatie-metrics
-   [ ] Cloud log integratie

### 🎯 v1.2.0 Routekaart (Toekomst)

-   [ ] Configuratiebestand ondersteuning (YAML/JSON/TOML)
-   [ ] Logaggregatie en filtering
-   [ ] Real-time log streaming
-   [ ] Uitgebreide beveiligingsfuncties
-   [ ] Prestatie dashboard integratie

## Bijdragen

We verwelkomen bijdragen! Bekijk onze [bijdragergids](/nl/CONTRIBUTING) voor details over:

-   Bugs rapporteren en functies aanvragen
-   Code inleveringsproces
-   Ontwikkelingsinstelling
-   Testvereisten
-   Documentatiestandaarden

## Beveiliging

Voor beveiligingslekken, zie onze [beveiligingsbeleid](/nl/SECURITY) voor:

-   Ondersteunde versies
-   Rapportageproces
-   Responstijdlijn
-   Beveiligings best practices

## Ondersteuning

-   📖 [Documentatie](docs/)
-   🐛 [Issue Tracker](https://github.com/lazygophers/log/issues)
-   💬 [Discussies](https://github.com/lazygophers/log/discussions)
-   📧 Email: support@lazygophers.com

## Licentie

Dit project is gelicentieerd onder de MIT Licentie - zie het [LICENSE](/nl/LICENSE) bestand voor details.

---

## 🌍 Meertalige Documentatie

Dit wijzigingslog is beschikbaar in meerdere talen:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CHANGELOG.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/CHANGELOG.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CHANGELOG.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CHANGELOG.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CHANGELOG.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CHANGELOG.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CHANGELOG.md)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/CHANGELOG.md) (huidige)

---

**Volg elke verbetering en blijf op de hoogte van de ontwikkeling van Lazygophers Log!🚀**

---

_Dit wijzigingslog wordt automatisch bijgewerkt bij elke publicatie. Voor de meest recente informatie, zie de [GitHub Releases](https://github.com/lazygophers/log/releases) pagina._
