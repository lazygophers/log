---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Una libreria di registrazione Go ad alte prestazioni e flessibile, costruita su zap, che fornisce funzionalità ricche e un'API semplice.

## 📖 Lingue della documentazione

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 Cinese semplificato](https://lazygophers.github.io/log/zh-CN/)
-   [🇹🇼 Cinese tradizionale](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/)
-   [🇵🇹 Português](https://lazygophers.github.io/log/pt/)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/)
-   [🇵🇱 Polski](https://lazygophers.github.io/log/pl/)
-   [🇮🇹 Italiano](README.md) (attuale)
-   [🇹🇷 Türkçe](https://lazygophers.github.io/log/tr/)

## ✨ Caratteristiche

-   **🚀 Alte prestazioni**：Costruito su zap con object pooling e registrazione condizionale dei campi
-   **📊 Ricchi livelli di log**：Livelli Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configurazione flessibile**：
    -   Controllo del livello di log
    -   Registrazione delle informazioni del chiamante
    -   Informazioni di tracciamento (incluso ID goroutine)
    -   Prefissi e suffissi personalizzati
    -   Destinazioni di output personalizzate (console, file, ecc.)
    -   Opzioni di formattazione del log
-   **🔄 Rotazione dei file**：Supporto per la rotazione oraria dei file di log
-   **🔌 Compatibilità Zap**：Integrazione trasparente con zap WriteSyncer
-   **🎯 API semplice**：API chiara simile alla libreria di log standard, facile da usare

## 🚀 Avvio rapido

### Installazione

:::tip Installazione
```bash
go get github.com/lazygophers/log
```
:::

### Utilizzo di base

```go title="Avvio rapido"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Usare il logger globale predefinito
    log.Debug("Messaggio di debug")
    log.Info("Messaggio informativo")
    log.Warn("Messaggio di avviso")
    log.Error("Messaggio di errore")

    // Usare l'output formattato
    log.Infof("Utente %s ha effettuato l'accesso", "admin")

    // Configurazione personalizzata
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Questo è un log dal logger personalizzato")
}
```

## 📚 Utilizzo avanzato

### Logger personalizzato con output su file

```go title="Configurazione output su file"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Creare logger con output su file
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Log di debug con informazioni del chiamante")
    logger.Info("Log informativo con informazioni di tracciamento")
}
```

### Controllo del livello di log

```go title="Controllo livello di log"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Solo warn e superiori verranno registrati
    logger.Debug("Questo non verrà registrato")  // Ignorato
    logger.Info("Questo non verrà registrato")   // Ignorato
    logger.Warn("Questo verrà registrato")    // Registrato
    logger.Error("Questo verrà registrato")   // Registrato
}
```

## 🎯 Scenari di utilizzo

### Scenari applicabili

-   **Servizi web e backend API**：Tracciamento delle richieste, logging strutturato, monitoraggio delle prestazioni
-   **Architettura a microservizi**：Tracciamento distribuito (TraceID), formato di log unificato, alto throughput
-   **Strumenti a riga di comando**：Controllo dei livelli, output pulito, segnalazione errori
-   **Attività in batch**：Rotazione dei file, esecuzione prolungata, ottimizzazione delle risorse

### Vantaggi speciali

-   **Ottimizzazione con object pool**：Riutilizzo degli oggetti Entry e Buffer, riducendo la pressione del GC
-   **Scrittura asincrona**：Alto throughput (10000+ log/secondo) senza blocco
-   **Supporto TraceID**：Tracciamento delle richieste in sistemi distribuiti, integrazione con OpenTelemetry
-   **Avvio senza configurazione**：Pronto all'uso, configurazione progressiva

## 🔧 Opzioni di configurazione

:::note Opzioni di configurazione
Tutti i seguenti metodi supportano le chiamate a catena e possono essere combinati per costruire un Logger personalizzato.
:::

### Configurazione Logger

| Metodo                  | Descrizione                | Predefinito      |
| --------------------- | ------------------- | ------------ |
| `SetLevel(level)`       | Imposta il livello minimo di log     | `DebugLevel` |
| `EnableCaller(enable)`  | Abilita/disabilita informazioni del chiamante | `false`      |
| `EnableTrace(enable)`   | Abilita/disabilita informazioni di tracciamento  | `false`      |
| `SetCallerDepth(depth)` | Imposta la profondità del chiamante   | `2`          |
| `SetPrefixMsg(prefix)`  | Imposta il prefisso del log  | `""`         |
| `SetSuffixMsg(suffix)`  | Imposta il suffisso del log  | `""`         |
| `SetOutput(writers...)` | Imposta le destinazioni di output         | `os.Stdout`  |

### Livelli di log

| Livello        | Descrizione                        |
| ----------- | --------------------------- |
| `TraceLevel` | Più dettagliato, per tracciamento dettagliato        |
| `DebugLevel` | Informazioni di debug                  |
| `InfoLevel`  | Informazioni generali                    |
| `WarnLevel`  | Messaggi di avviso                  |
| `ErrorLevel` | Messaggi di errore                  |
| `FatalLevel` | Errori fatali (chiama os.Exit(1))    |
| `PanicLevel` | Errori di panico (chiama panic())      |

## 🏗️ Architettura

### Componenti principali

-   **Logger**：Struttura di registrazione principale con opzioni configurabili
-   **Entry**：Registrazione individuale con supporto completo dei campi
-   **Level**：Definizioni dei livelli di log e funzioni di utilità
-   **Format**：Interfaccia di formattazione del log e implementazioni

### Ottimizzazione delle prestazioni

-   **Object pooling**：Riutilizza gli oggetti Entry per ridurre l'allocazione di memoria
-   **Registrazione condizionale**：Registra solo campi costosi quando necessario
-   **Controllo rapido dei livelli**：Controlla il livello di log al livello più esterno
-   **Design senza lock**：La maggior parte delle operazioni non richiede lock

## 📊 Confronto delle prestazioni

:::info Confronto delle prestazioni
I seguenti dati si basano su benchmark; le prestazioni effettive possono variare a seconda dell'ambiente e della configurazione.
:::

| Caratteristica          | lazygophers/log | zap    | logrus | log standard |
| ------------- | --------------- | ------ | ------ | -------------- |
| Prestazioni      | Alto              | Alto     | Medio     | Basso       |
| Semplicità API    | Alto              | Medio     | Alto     | Alto       |
| Ricchezza funzionalità    | Medio          | Alto     | Alto     | Basso       |
| Flessibilità      | Medio          | Alto     | Alto     | Basso       |
| Curva di apprendimento      | Basso              | Medio     | Medio     | Basso       |

## ❓ Domande frequenti

### Come scegliere il livello di log appropriato?

-   **Ambiente di sviluppo**：Usa `DebugLevel` o `TraceLevel` per informazioni dettagliate
-   **Ambiente di produzione**：Usa `InfoLevel` o `WarnLevel` per ridurre l'overhead
-   **Test delle prestazioni**：Usa `PanicLevel` per disabilitare tutti i log

### Come ottimizzare le prestazioni in produzione?

:::warning Nota
In scenari ad alto throughput, si raccomanda di combinare la scrittura asincrona con livelli di log ragionevoli per ottimizzare le prestazioni.
:::

1. Usa `AsyncWriter` per la scrittura asincrona：

```go title="Configurazione scrittura asincrona"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Regola i livelli di log per evitare registrazioni non necessarie：

```go title="Ottimizzazione livello"
logger.SetLevel(log.InfoLevel)  // Salta Debug e Trace
```

3. Usa la registrazione condizionale per ridurre l'overhead：

```go title="Registrazione condizionale"
if logger.Level >= log.DebugLevel {
    logger.Debug("Informazioni di debug dettagliate")
}
```

### Qual è la differenza tra `Caller` e `EnableCaller`?

-   **`EnableCaller(enable bool)`**：Controlla se il Logger raccoglie le informazioni del chiamante
    -   `EnableCaller(true)` abilita la raccolta delle informazioni del chiamante
-   **`Caller(disable bool)`**：Controlla se il Formatter emette le informazioni del chiamante
    -   `Caller(true)` disabilita l'emissione delle informazioni del chiamante

Si raccomanda di usare `EnableCaller` per il controllo globale.

### Come implementare un formattatore personalizzato?

Implementa l'interfaccia `Format`：

```go title="Formattatore personalizzato"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Documentazione correlata

-   [📚 Documentazione API](API.md) - Riferimento API completo
-   [🤝 Guida al contributo](/it/CONTRIBUTING) - Come contribuire
-   [📋 Registro delle modifiche](/it/CHANGELOG) - Storico delle versioni
-   [🔒 Politica di sicurezza](/it/SECURITY) - Linee guida sulla sicurezza
-   [📜 Codice di condotta](/it/CODE_OF_CONDUCT) - Linee guida della comunità

## 🚀 Ottenere aiuto

-   **GitHub Issues**：[Segnalare bug o richiedere funzionalità](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[Documentazione API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Esempi**：[Esempi di utilizzo](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licenza

Questo progetto è concesso in licenza sotto la Licenza MIT - vedere il file [LICENSE](/it/LICENSE) per i dettagli.

## 🤝 Contribuire

Accogliamo i contributi! Consulta la nostra [Guida al contributo](/it/CONTRIBUTING) per maggiori informazioni.

---

**lazygophers/log** è progettato per essere la soluzione di logging preferita per gli sviluppatori Go che apprezzano sia le prestazioni che la semplicità. Che tu stia costruendo una piccola utility o un sistema distribuito su larga scala, questa libreria fornisce il giusto equilibrio tra funzionalità e facilità d'uso.
