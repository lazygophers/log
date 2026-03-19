---
titleSuffix: " | LazyGophers Log"
---

# 📋 Changelog

Tutte le modifiche importanti di questo progetto sono documentate in questo file.

Il formato si basa su [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), e questo progetto segue [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Non Rilasciato]

### Aggiunto

-   Documentazione multilingue completa (7 lingue)
-   Template issue GitHub (segnalazioni bug, richieste funzionalità, domande)
-   Template Pull Request con controlli compatibilità build tag
-   Linee guida contributo multilingua
-   Codice di condotta con linee guida di applicazione
-   Politica di sicurezza con processo di segnalazione vulnerabilità
-   Documentazione API completa con esempi
-   Struttura progetto professionale e template

### Modificato

-   README migliorato con documentazione completa delle funzionalità
-   Copertura test migliorata per tutte le configurazioni build tag
-   Struttura progetto aggiornata per migliore manutenibilità

### Documentazione

-   Aggiunto supporto multilingue per tutti i documenti principali
-   Creata guida API completa
-   Stabilito linee guida workflow contributi
-   Implementato processo segnalazione sicurezza

## [1.0.0] - 2024-01-01

### Aggiunto

-   Funzionalità logging core con livelli multipli (Trace, Debug, Info, Warn, Error, Fatal, Panic)
-   Implementazione logger thread-safe con object pooling
-   Supporto build tag (predefinito, debug, release, modalità discard)
-   Interfaccia formatter personalizzabile con formatter testo predefinito
-   Supporto output multi-writer
-   Funzionalità scrittura asincrona per scenari ad alto throughput
-   Rotazione file log oraria automatica
-   Logging context-aware con tracking goroutine ID e trace ID
-   Informazioni chiamante con profondità stack configurabile
-   Funzioni di convenienza a livello pacchetto globale
-   Supporto integrazione logger Zap

### Prestazioni

-   Object pooling per oggetti entry e buffer con `sync.Pool`
-   Controlli livello anticipati per evitare operazioni costose
-   Writer asincrono per operazioni di scrittura log non bloccanti
-   Ottimizzazioni build tag per ambienti diversi

### Build Tag

-   **Predefinito**: Funzionalità completa con messaggi debug
-   **Debug**: Informazioni debug avanzate e dettagli chiamante
-   **Release**: Ottimizzazioni produzione, messaggi debug disabilitati
-   **Discard**: Prestazioni massime, operazioni log no-op

### Funzionalità Core

-   **Logger**: Struttura logger principale con livello, output, formatter configurabili
-   **Entry**: Struttura record log con metadati completi
-   **Livelli**: Sette livelli log da Panic (più alto) a Trace (più basso)
-   **Formatter**: Sistema formato plug-in
-   **Writer**: Supporto rotazione file e scrittura asincrona
-   **Contesto**: Supporto goroutine ID e tracing distribuito

### Punti Salienti API

-   API configurazione fluente con method chaining
-   Metodi log semplici e formattati (`.Info()` e `.Infof()`)
-   Clonazione logger per configurazioni isolate
-   Logging context-aware con `CloneToCtx()`
-   Personalizzazione messaggi prefisso e suffisso
-   Interruttore informazioni chiamante

### Test

-   Suite test completa con copertura 93,5%
-   Supporto test multi-build tag
-   Workflow test automatizzati
-   Benchmark prestazioni

## [0.9.0] - 2023-12-15

### Aggiunto

-   Struttura progetto iniziale
-   Funzionalità logging di base
-   Filtraggio basato su livelli
-   Supporto output file

### Modificato

-   Prestazioni migliorate tramite object pooling
-   Gestione errori migliorata

## [0.8.0] - 2023-12-01

### Aggiunto

-   Supporto multi-writer
-   Interfaccia formatter personalizzabile
-   Funzionalità scrittura asincrona

### Risolto

-   Memory leak in scenari ad alto throughput
-   Condizioni race nell'accesso concorrente

## [0.7.0] - 2023-11-15

### Aggiunto

-   Supporto build tag per compilazione condizionale
-   Livelli log Trace e Debug
-   Tracking informazioni chiamante

### Modificato

-   Ottimizzati pattern allocazione memoria
-   Migliorata thread-safety

## [0.6.0] - 2023-11-01

### Aggiunto

-   Funzionalità rotazione log
-   Logging context-aware
-   Tracking goroutine ID

### Deprecato

-   Vecchi metodi configurazione (verranno rimossi in v1.0.0)

## [0.5.0] - 2023-10-15

### Aggiunto

-   Formatter JSON
-   Destinazioni output multiple
-   Benchmark prestazioni

### Modificato

-   Ristrutturato engine logging core
-   Migliorata coerenza API

### Rimosso

-   Vecchi metodi log

## [0.4.0] - 2023-10-01

### Aggiunto

-   Livelli log Fatal e Panic
-   Funzioni globali pacchetto
-   Validazione configurazione

### Risolto

-   Problemi sincronizzazione output
-   Ottimizzazione uso memoria

## [0.3.0] - 2023-09-15

### Aggiunto

-   Livelli log personalizzati
-   Interfaccia formatter
-   Operazioni thread-safe

### Modificato

-   Semplificato design API
-   Documentazione estesa

## [0.2.0] - 2023-09-01

### Aggiunto

-   Supporto output file
-   Filtraggio basato su livelli
-   Opzioni formattazione base

### Risolto

-   Colli di bottiglia prestazioni
-   Memory leak

## [0.1.0] - 2023-08-15

### Aggiunto

-   Rilascio iniziale
-   Logging console base
-   Supporto livelli semplice (Info, Warn, Error)
-   Struttura logger core

## Riepilogo Storico Versioni

| Versione | Data Rilascio | Funzionalità Principali |
| -------- | ------------- | ----------------------- |
| 1.0.0    | 2024-01-01    | Sistema logging completo, build tag, scrittura asincrona, documentazione completa |
| 0.9.0    | 2023-12-15    | Miglioramenti prestazioni, object pooling |
| 0.8.0    | 2023-12-01    | Multi-writer, scrittura asincrona, formatter personalizzati |
| 0.7.0    | 2023-11-15    | Build tag, livelli Trace/Debug, informazioni chiamante |
| 0.6.0    | 2023-11-01    | Rotazione log, logging contestuale, tracking goroutine |
| 0.5.0    | 2023-10-15    | Formatter JSON, output multipli, benchmark |
| 0.4.0    | 2023-10-01    | Livelli Fatal/Panic, funzioni globali |
| 0.3.0    | 2023-09-15    | Livelli personalizzati, interfaccia formatter |
| 0.2.0    | 2023-09-01    | Output file, filtraggio livelli |
| 0.1.0    | 2023-08-15    | Rilascio iniziale, logging console base |

## Guida Migrazione

### Migrazione da v0.9.x a v1.0.0

#### Modifiche Breaking

-   Nessuna - v1.0.0 è retrocompatibile con v0.9.x

#### Nuove Funzionalità Disponibili

-   Supporto build tag esteso
-   Documentazione completa
-   Template progetto professionali
-   Processo segnalazione sicurezza

#### Aggiornamenti Consigliati

```go
// Vecchio modo (ancora supportato)
logger := log.New()
logger.SetLevel(log.InfoLevel)

// Nuovo modo consigliato, con method chaining
logger := log.New().
    SetLevel(log.InfoLevel).
    Caller(true).
    SetPrefixMsg("[MyApp] ")
```

### Migrazione da v0.8.x a v0.9.x

#### Modifiche Breaking

-   Rimossi metodi configurazione deprecati
-   Cambiata gestione buffer interna

#### Passaggi Migrazione

1. Aggiorna percorsi import se necessario
2. Sostituisci metodi deprecati:

    ```go
    // Vecchio (deprecato)
    logger.SetOutputFile("app.log")

    // Nuovo
    file, _ := os.Create("app.log")
    logger.SetOutput(file)
    ```

### Migrazione da v0.5.x e precedenti

#### Modifiche Principali

-   Ridisegno API completo per migliore coerenza
-   Miglioramenti prestazioni tramite object pooling
-   Nuovo sistema build tag

#### Migrazione Richiesta

-   Aggiorna tutte le chiamate log alla nuova API
-   Rivedi e aggiorna implementazioni formatter
-   Testa con le nuove configurazioni build tag

## Pietre Miliari Sviluppo

### 🎯 Roadmap v1.1.0 (Pianificata)

-   [ ] Logging strutturato con coppie chiave-valore
-   [ ] Campionamento log per scenari ad alto volume
-   [ ] Sistema plugin per output personalizzati
-   [ ] Metriche prestazioni avanzate
-   [ ] Integrazione cloud log

### 🎯 Roadmap v1.2.0 (Futuro)

-   [ ] Supporto file configurazione (YAML/JSON/TOML)
-   [ ] Aggregazione e filtraggio log
-   [ ] Streaming log in tempo reale
-   [ ] Funzionalità sicurezza avanzate
-   [ ] Integrazione dashboard prestazioni

## Contributi

Accogliamo con favore i contributi! Consulta la nostra [guida contributi](/it/CONTRIBUTING) per dettagli su:

-   Segnalazione bug e richiesta funzionalità
-   Processo sottomissione codice
-   Impostazione sviluppo
-   Requisiti test
-   Standard documentazione

## Sicurezza

Per vulnerabilità di sicurezza, consulta la nostra [politica sicurezza](/it/SECURITY) per:

-   Versioni supportate
-   Processo segnalazione
-   Timeline risposta
-   Best practices sicurezza

## Supporto

-   📖 [Documentazione](docs/)
-   🐛 [Issue Tracker](https://github.com/lazygophers/log/issues)
-   💬 [Discussioni](https://github.com/lazygophers/log/discussions)
-   📧 Email: support@lazygophers.com

## Licenza

Questo progetto è concesso in licenza sotto la Licenza MIT - vedi il file [LICENSE](/it/LICENSE) per dettagli.

---

## 🌍 Documentazione Multilingua

Questo changelog è disponibile in più lingue:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CHANGELOG.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/CHANGELOG.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CHANGELOG.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CHANGELOG.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CHANGELOG.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CHANGELOG.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CHANGELOG.md)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/CHANGELOG.md) (corrente)

---

**Segui ogni miglioramento e rimani aggiornato sullo sviluppo di Lazygophers Log!🚀**

---

_Questo changelog viene aggiornato automaticamente ad ogni rilascio. Per le informazioni più recenti, consulta la pagina [GitHub Releases](https://github.com/lazygophers/log/releases)._
