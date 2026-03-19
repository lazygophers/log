---
pageType: home

hero:
    name: LazyGophers Log
    text: Krachtige en flexibele Go logboekbibliotheek
    tagline: Gebouwd op zap, biedt rijke functies en eenvoudige API
    actions:
        - theme: brand
          text: Snelstart
          link: /API
        - theme: alt
          text: Bekijk op GitHub
          link: https://github.com/lazygophers/log

features:
    - title: "Hoog Prestaties"
      details: Gebouwd op zap met hergebruik van objectpool en conditionele veldopname
      icon: 🚀
    - title: "Rijke Logniveaus"
      details: Ondersteunt Trace, Debug, Info, Warn, Error, Fatal, Panic niveaus
      icon: 📊
    - title: "Flexibele Configuratie"
      details: Personaliseer niveaus, aanroepersinformatie, tracering, voorvoegsels, achtervoegsels en uitvoerdoelen
      icon: ⚙️
    - title: "Bestandsrotatie"
      details: Ingebouwde ondersteuning voor uurlijkse logbestandsrotatie
      icon: 🔄
    - title: "Zap Compatibiliteit"
      details: Naadloze integratie met zap WriteSyncer
      icon: 🔌
    - title: "Eenvoudige API"
      details: Duidelijke API vergelijkbaar met standaard logboekbibliotheek, eenvoudig te gebruiken en integreren
      icon: 🎯
---

## Snelstart

### Installatie

```bash
go get github.com/lazygophers/log
```

### Basisgebruik

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Standaard globale logger gebruiken
    log.Debug("Foutopsporingsinformatie")
    log.Info("Algemene informatie")
    log.Warn("Waarschuwingsinformatie")
    log.Error("Foutinformatie")

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

## Documentatie

-   [API Referentie](API.md) - Volledige API documentatie
-   [Wijzigingslog](/nl/CHANGELOG) - Versiegeschiedenis
-   [Bijdragengids](/nl/CONTRIBUTING) - Hoe bij te dragen
-   [Beveiligingsbeleid](/nl/SECURITY) - Beveiligingsgids
-   [Gedragscode](/nl/CODE_OF_CONDUCT) - Community richtlijnen

## Prestatievergelijking

| Kenmerk       | lazygophers/log | zap | logrus | Standaard log |
| ---------- | --------------- | --- | ------ | -------- |
| Prestaties       | Hoog              | Hoog  | Gemiddeld     | Laag       |
| API-eenvoud    | Hoog              | Gemiddeld  | Hoog     | Hoog       |
| Functionaliteit    | Gemiddeld              | Hoog  | Hoog     | Laag       |
| Flexibiliteit      | Gemiddeld              | Hoog  | Hoog     | Laag       |
| Leercurve      | Laag              | Gemiddeld  | Gemiddeld     | Laag       |

## Licentie

Dit project is gelicentieerd onder de MIT Licentie - zie het [LICENSE](/nl/LICENSE) bestand voor details.
