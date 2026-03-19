---
pageType: home

hero:
    name: LazyGophers Log
    text: Libreria di logging Go ad alte prestazioni e flessibile
    tagline: Costruita su zap, fornisce funzionalità ricche e API semplice
    actions:
        - theme: brand
          text: Avvio Rapido
          link: /API
        - theme: alt
          text: Vedi su GitHub
          link: https://github.com/lazygophers/log

features:
    - title: "Alte Prestazioni"
      details: Costruita su zap con riutilizzo del pool di oggetti e registrazione condizionale dei campi
      icon: 🚀
    - title: "Livelli di Log Ricchi"
      details: Supporta livelli Trace, Debug, Info, Warn, Error, Fatal, Panic
      icon: 📊
    - title: "Configurazione Flessibile"
      details: Personalizza livelli, informazioni del chiamante, tracciamento, prefissi, suffissi e destinazioni di output
      icon: ⚙️
    - title: "Rotazione File"
      details: Supporto integrato per la rotazione oraria dei file di log
      icon: 🔄
    - title: "Compatibilità Zap"
      details: Integrazione perfetta con zap WriteSyncer
      icon: 🔌
    - title: "API Semplice"
      details: API chiara simile alla libreria di log standard, facile da usare e integrare
      icon: 🎯
---

## Avvio Rapido

### Installazione

```bash
go get github.com/lazygophers/log
```

### Uso Base

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Usa il logger globale predefinito
    log.Debug("Informazioni di debug")
    log.Info("Informazioni generali")
    log.Warn("Informazioni di avviso")
    log.Error("Informazioni di errore")

    // Usa output formattato
    log.Infof("Utente %s ha effettuato l'accesso", "admin")

    // Configurazione personalizzata
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Questo è un log dal logger personalizzato")
}
```

## Documentazione

-   [Riferimento API](API.md) - Documentazione API completa
-   [Registro delle Modifiche](/it/CHANGELOG) - Storico versioni
-   [Guida ai Contributi](/it/CONTRIBUTING) - Come contribuire
-   [Politica di Sicurezza](/it/SECURITY) - Guida alla sicurezza
-   [Codice di Condotta](/it/CODE_OF_CONDUCT) - Linee guida della community

## Confronto Prestazioni

| Caratteristica       | lazygophers/log | zap | logrus | Log standard |
| ---------- | --------------- | --- | ------ | -------- |
| Prestazioni       | Alto              | Alto  | Medio     | Basso       |
| Semplicità API    | Alto              | Medio  | Alto     | Alto       |
| Ricchezza Funzionalità    | Medio              | Alto  | Alto     | Basso       |
| Flessibilità      | Medio              | Alto  | Alto     | Basso       |
| Curva di Apprendimento      | Basso              | Medio  | Medio     | Basso       |

## Licenza

Questo progetto è concesso in licenza sotto la Licenza MIT - consulta il file [LICENSE](/it/LICENSE) per dettagli.
