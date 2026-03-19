---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Een krachtige en flexibele Go-logboekbibliotheek, gebouwd op zap, die rijke functies en een eenvoudige API biedt.

## 📖 Documentatietalen

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 Vereenvoudigd Chinees](README.md) (huidige)
-   [🇹🇼 Traditioneel Chinees](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)

## ✨ Functies

-   **🚀 Hoge prestaties**：Gebouwd op zap met object pooling en conditionele veldopname
-   **📊 Rijk logniveaus**：Trace, Debug, Info, Warn, Error, Fatal, Panic niveaus
-   **⚙️ Flexibele configuratie**：
    -   Logniveau controle
    -   Aanroepersinformatie vastlegging
    -   Trace-informatie (inclusief goroutine ID)
    -   Aangepaste voorvoegsels en achtervoegsels
    -   Aangepaste uitvoerdoelen (console, bestanden, enz.)
    -   Lognotatieopties
-   **🔄 Bestandsrotatie**：Ondersteuning voor uurlijkse logbestandsrotatie
-   **🔌 Zap-compatibiliteit**：Naadloze integratie met zap WriteSyncer
-   **🎯 Eenvoudige API**：Heldere API vergelijkbaar met de standaard logboekbibliotheek, eenvoudig te gebruiken

## 🚀 Snelstart

### Installatie

:::tip Installatie
```bash
go get github.com/lazygophers/log
```
:::

### Basisgebruik

```go title="Snelstart"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Standaard globale logger gebruiken
    log.Debug("Foutopsporingsbericht")
    log.Info("Informatiebericht")
    log.Warn("Waarschuwingsbericht")
    log.Error("Foutbericht")

    // Opgemaakte uitvoer gebruiken
    log.Infof("Gebruiker %s succesvol ingelogd", "admin")

    // Aangepaste configuratie
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Dit is een logboek van de aangepaste logger")
}
```

## 📚 Geavanceerd gebruik

### Aangepaste logger met bestandsuitvoer

```go title="Bestandsuitvoerconfiguratie"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Logger maken met bestandsuitvoer
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Foutopsporingslog met aanroepersinformatie")
    logger.Info("Informatielog met trace-informatie")
}
```

### Logniveau controle

```go title="Logniveau controle"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Alleen warn en hoger worden gelogd
    logger.Debug("Dit wordt niet gelogd")  // Genegeerd
    logger.Info("Dit wordt niet gelogd")   // Genegeerd
    logger.Warn("Dit wordt gelogd")    // Gelogd
    logger.Error("Dit wordt gelogd")   // Gelogd
}
```

## 🎯 Gebruiksscenario's

### Toepasselijke scenario's

-   **Webservices en API-backends**：Verzoeken traceren, gestructureerd loggen, prestatiebewaking
-   **Microservices-architectuur**：Gedistribueerde tracering (TraceID), uniforme lognotatie, hoge doorvoer
-   **Commandoregelhulpmiddelen**：Niveau controle, schone uitvoer, foutrapportage
-   **Batch-taken**：Bestandsrotatie, lange uitvoering, resourceoptimalisatie

### Speciale voordelen

-   **Object pool optimalisatie**：Hergebruik van Entry- en Buffer-objecten, vermindert GC-druk
-   **Asynchroon schrijven**：Hoge doorvoerscenario's (10000+ logs/seconde) zonder blokkering
-   **TraceID-ondersteuning**：Verzoeken traceren in gedistribueerde systemen, integratie met OpenTelemetry
-   **Nul configuratie start**：Klaar voor gebruik, progressieve configuratie

## 🔧 Configuratieopties

:::note Configuratieopties
Alle volgende methoden ondersteunen kettingaanroep en kunnen worden gecombineerd om een aangepaste Logger te bouwen.
:::

### Logger-configuratie

| Methode                  | Beschrijving                | Standaard      |
| --------------------- | ------------------- | ------------ |
| `SetLevel(level)`       | Minimaal logniveau instellen     | `DebugLevel` |
| `EnableCaller(enable)`  | Aanroepersinformatie in/uitschakelen | `false`      |
| `EnableTrace(enable)`   | Trace-informatie in/uitschakelen  | `false`      |
| `SetCallerDepth(depth)` | Aanroeperdiepte instellen   | `2`          |
| `SetPrefixMsg(prefix)`  | Logvoorvoegsel instellen  | `""`         |
| `SetSuffixMsg(suffix)`  | Logachtervoegsel instellen  | `""`         |
| `SetOutput(writers...)` | Uitvoerdoelen instellen         | `os.Stdout`  |

### Logniveaus

| Niveau        | Beschrijving                        |
| ----------- | --------------------------- |
| `TraceLevel` | Meest gedetailleerd, voor gedetailleerde tracering        |
| `DebugLevel` | Foutopsporingsinformatie                  |
| `InfoLevel`  | Algemene informatie                    |
| `WarnLevel`  | Waarschuwingsberichten                  |
| `ErrorLevel` | Foutberichten                  |
| `FatalLevel` | Fatale fouten (roept os.Exit(1) aan)    |
| `PanicLevel` | Paniekfouten (roept panic() aan)      |

## 🏗️ Architectuur

### Kerncomponenten

-   **Logger**：Hoofdlogboekstructuur met configureerbare opties
-   **Entry**：Individueel logboekregistratie met uitgebreide veldondersteuning
-   **Level**：Logniveau-definities en hulpfuncties
-   **Format**：Lognotatie-interface en implementaties

### Prestatieoptimalisatie

-   **Object pooling**：Hergebruikt Entry-objecten om geheugentoewijzing te verminderen
-   **Conditionele opname**：Neemt alleen kostbare velden op indien nodig
-   **Snelle niveaucontrole**：Controleert logniveau op buitenste laag
-   **Vergrendelingsvrij ontwerp**：De meeste bewerkingen vereisen geen vergrendelingen

## 📊 Prestatievergelijking

:::info Prestatievergelijking
De volgende gegevens zijn gebaseerd op benchmarks; werkelijke prestaties kunnen variëren afhankelijk van omgeving en configuratie.
:::

| Kenmerk          | lazygophers/log | zap    | logrus | Standaard log |
| ------------- | --------------- | ------ | ------ | -------------- |
| Prestaties      | Hoog              | Hoog     | Gemiddeld     | Laag       |
| API-eenvoud    | Hoog              | Gemiddeld     | Hoog     | Hoog       |
| Functionaliteit    | Gemiddeld          | Hoog     | Hoog     | Laag       |
| Flexibiliteit      | Gemiddeld          | Hoog     | Hoog     | Laag       |
| Leercurve      | Laag              | Gemiddeld     | Gemiddeld     | Laag       |

## ❓ Veelgestelde vragen

### Hoe kies ik het juiste logniveau?

-   **Ontwikkelingsomgeving**：Gebruik `DebugLevel` of `TraceLevel` voor gedetailleerde informatie
-   **Productieomgeving**：Gebruik `InfoLevel` of `WarnLevel` om overhead te verminderen
-   **Prestatietests**：Gebruik `PanicLevel` om alle logboekregistratie uit te schakelen

### Hoe optimaliseer ik prestaties in productie?

:::warning Let op
In scenario's met hoge doorvoer wordt aanbevolen om asynchroon schrijven te combineren met redelijke logniveaus om prestaties te optimaliseren.
:::

1. Gebruik `AsyncWriter` voor asynchroon schrijven：

```go title="Asynchroon schrijven configuratie"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Pas logniveaus aan om onnodige logboekregistratie te vermijden：

```go title="Niveau optimalisatie"
logger.SetLevel(log.InfoLevel)  // Debug en Trace overslaan
```

3. Gebruik conditionele logboekregistratie om overhead te verminderen：

```go title="Conditionele logboekregistratie"
if logger.Level >= log.DebugLevel {
    logger.Debug("Gedetailleerde foutopsporingsinformatie")
}
```

### Wat is het verschil tussen `Caller` en `EnableCaller`?

-   **`EnableCaller(enable bool)`**：Bepaalt of de Logger aanroepersinformatie verzamelt
    -   `EnableCaller(true)` schakelt aanroepersinformatieverzameling in
-   **`Caller(disable bool)`**：Bepaalt of de Formatter aanroepersinformatie uitvoert
    -   `Caller(true)` schakelt aanroepersinformatie-uitvoer uit

Het wordt aanbevolen om `EnableCaller` te gebruiken voor globale controle.

### Hoe implementeer ik een aangepaste formatter?

Implementeer de `Format`-interface：

```go title="Aangepaste formatter"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Gerelateerde documentatie

-   [📚 API-documentatie](API.md) - Volledige API-referentie
-   [🤝 Bijdragengids](/nl/CONTRIBUTING) - Hoe bij te dragen
-   [📋 Wijzigingslog](/nl/CHANGELOG) - Versiegeschiedenis
-   [🔒 Beveiligingsbeleid](/nl/SECURITY) - Beveiligingsrichtlijnen
-   [📜 Gedragscode](/nl/CODE_OF_CONDUCT) - Gemeenschapsrichtlijnen

## 🚀 Hulp krijgen

-   **GitHub Issues**：[Bugs melden of functies aanvragen](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[API-documentatie](https://pkg.go.dev/github.com/lazygophers/log)
-   **Voorbeelden**：[Gebruiksvoorbeelden](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licentie

Dit project is gelicentieerd onder de MIT Licentie - zie het [LICENSE](/nl/LICENSE) bestand voor details.

## 🤝 Bijdragen

Waardering voor bijdragen! Raadpleeg onze [Bijdragengids](/nl/CONTRIBUTING) voor meer informatie.

---

**lazygophers/log** is ontworpen om de voorkeurslogboekoplossing te zijn voor Go-ontwikkelaars die zowel prestaties als eenvoud waarderen. Of u nu een klein hulpprogramma bouwt of een grootschalig gedistribueerd systeem, deze bibliotheek biedt de juiste balans tussen functionaliteit en gebruiksgemak.
