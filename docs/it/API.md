---
titleSuffix: ' | LazyGophers Log'
---
# 📚 Documentazione API

## Panoramica

LazyGophers Log fornisce un'API di logging completa con supporto per più livelli, formattazione personalizzata, scrittura asincrona e ottimizzazione dei tag di build. Questo documento copre tutte le API pubbliche, le opzioni di configurazione e i pattern di utilizzo.

## Indice

-   [Tipi Principali](#tipi-principali)
-   [API del Logger](#api-del-logger)
-   [Funzioni Globali](#funzioni-globali)
-   [Livelli di Log](#livelli-di-log)
-   [Formatter](#formatter)
-   [Writer di Output](#writer-di-output)
-   [Logging con Contesto](#logging-con-contesto)
-   [Tag di Build](#tag-di-build)
-   [Ottimizzazione delle Prestazioni](#ottimizzazione-delle-prestazioni)
-   [Esempi](#esempi)

## Tipi Principali

### Logger

La struttura principale del logger che fornisce tutte le funzionalità di logging.

```go
type Logger struct {
    // Contiene campi privati per operazioni thread-safe
}
```

#### Costruttore

```go
func New() *Logger
```

Crea una nuova istanza del logger con configurazione predefinita:

-   Livello: `DebugLevel`
-   Output: `os.Stdout`
-   Formatter: Formattatore di testo predefinito
-   Tracciamento del chiamante: Disattivato

**Esempio:**

```go title="Creazione di un logger"
logger := log.New()
logger.Info("Nuovo logger creato")
```

### Entry

Rappresenta una singola voce di log con tutti i metadati associati.

```go
type Entry struct {
    Time       time.Time     // Timestamp di creazione della voce
    Level      Level         // Livello di log
    Message    string        // Messaggio di log
    Pid        int          // ID del processo
    Gid        uint64       // ID della goroutine
    TraceID    string       // ID di tracciamento per tracing distribuito
    CallerName string       // Nome della funzione chiamante
    CallerFile string       // Percorso del file chiamante
    CallerLine int          // Numero di riga del chiamante
}
```

## API del Logger

### Metodi di Configurazione

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Imposta il livello minimo di log. I messaggi sotto questo livello verranno ignorati.

**Parametri:**

-   `level`: Il livello minimo di log da elaborare

**Restituisce:**

-   `*Logger`: Restituisce se stesso per supportare il method chaining

**Esempio:**

```go title="Impostazione del livello di log"
logger.SetLevel(log.InfoLevel)
logger.Debug("Questo non verrà visualizzato")  // Ignorato
logger.Info("Questo verrà visualizzato")        // Elaborato
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Imposta una o più destinazioni di output per i messaggi di log.

**Parametri:**

-   `writers`: Una o più destinazioni di output `io.Writer`

**Restituisce:**

-   `*Logger`: Restituisce se stesso per supportare il method chaining

**Esempio:**

```go title="Impostazione della destinazione di output"
// Output singolo
logger.SetOutput(os.Stdout)

// Output multipli
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

Imposta un formattatore personalizzato per l'output di log.

**Parametri:**

-   `formatter`: Un formattatore che implementa l'interfaccia `Format`

**Restituisce:**

-   `*Logger`: Restituisce se stesso per supportare il method chaining

**Esempio:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Attiva o disattiva la registrazione delle informazioni del chiamante nelle voci di log.

**Parametri:**

-   `enable`: Se attivare le informazioni del chiamante (passa `true` per attivare)

**Restituisce:**

-   `*Logger`: Restituisce se stesso per supportare il method chaining

**Esempio:**

```go
logger.EnableCaller(true)
logger.Info("Questo includerà informazioni file:riga")

logger.EnableCaller(false)
logger.Info("Questo non includerà informazioni file:riga")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

Controlla le informazioni del chiamante nel formattatore.

**Parametri:**

-   `disable`: Se disabilitare le informazioni del chiamante (passa `true` per disabilitare)

**Restituisce:**

-   `*Logger`: Restituisce se stesso per supportare il method chaining

**Esempio:**

```go
logger.Caller(false)  // Non disabilitare, mostra informazioni del chiamante
logger.Info("Questo includerà informazioni file:riga")

logger.Caller(true)   // Disabilita informazioni del chiamante
logger.Info("Questo non includerà informazioni file:riga")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Imposta la profondità dello stack per le informazioni del chiamante quando si wrappa il logger.

**Parametri:**

-   `depth`: Il numero di frame dello stack da saltare

**Restituisce:**

-   `*Logger`: Restituisce se stesso per supportare il method chaining

**Esempio:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // Salta la funzione wrapper
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Imposta il testo di prefisso o suffisso per tutti i messaggi di log.

**Parametri:**

-   `prefix/suffix`: Il testo da anteporre/accodare al messaggio

**Restituisce:**

-   `*Logger`: Restituisce se stesso per supportare il method chaining

**Esempio:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // Output: [APP] Hello [END]
```

### Metodi di Logging

Tutti i metodi di logging hanno due varianti: una versione semplice e una versione formattata.

#### Livello Trace

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

Registra i log al livello trace (più dettagliato).

**Esempio:**

```go
logger.Trace("Tracciamento dettagliato dell'esecuzione")
logger.Tracef("Elaborazione elemento %d di %d", i, total)
```

#### Livello Debug

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

Registra informazioni di sviluppo al livello debug.

**Esempio:**

```go
logger.Debug("Stato variabile:", variable)
logger.Debugf("Utente %s autenticato con successo", username)
```

#### Livello Info

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

Registra messaggi informativi.

**Esempio:**

```go
logger.Info("Applicazione avviata")
logger.Infof("Server in ascolto sulla porta %d", port)
```

#### Livello Warn

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

Registra messaggi di avviso per situazioni potenzialmente problematiche.

**Esempio:**

```go
logger.Warn("Funzione deprecata chiamata")
logger.Warnf("Uso memoria elevato: %d%%", memoryPercent)
```

#### Livello Error

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

Registra messaggi di errore.

**Esempio:**

```go
logger.Error("Connessione al database fallita")
logger.Errorf("Elaborazione richiesta fallita: %v", err)
```

#### Livello Fatal

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

Registra un errore fatale e chiama `os.Exit(1)`.

:::danger Operazione Distruttiva
`Fatal` e `Fatalf` chiameranno `os.Exit(1)` immediatamente dopo il logging, terminando il processo. Usare solo in condizioni di errore non recuperabili. Le istruzioni `defer` **non** verranno eseguite.
:::

**Esempio:**

```go
logger.Fatal("Errore critico del sistema")
logger.Fatalf("Impossibile avviare il server: %v", err)
```

#### Livello Panic

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

Registra un messaggio di errore e chiama `panic()`.

:::danger Operazione Distruttiva
`Panic` e `Panicf` chiameranno `panic()` dopo il logging. A differenza di `Fatal`, `panic` può essere recuperato con `recover()`, ma terminerà il programma se non viene catturato.
:::

**Esempio:**

```go
logger.Panic("Si è verificato un errore non recuperabile")
logger.Panicf("Stato non valido: %v", state)
```

### Metodi Utili

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Crea una copia del logger con la stessa configurazione.

**Restituisce:**

-   `*Logger`: Nuova istanza del logger con impostazioni copiate

**Esempio:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

Crea un logger context-aware che accetta `context.Context` come primo parametro.

**Restituisce:**

-   `LoggerWithCtx`: Istanza del logger context-aware

**Esempio:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "Messaggio context-aware")
```

## Funzioni Globali

Funzioni a livello di pacchetto che usano il logger globale predefinito.

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

**Esempio:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("Uso del logger globale")
```

## Livelli di Log

### Tipo Level

```go
type Level int8
```

### Livelli Disponibili

```go
const (
    PanicLevel Level = iota  // 0 - Panic ed esci
    FatalLevel              // 1 - Errore fatale ed esci
    ErrorLevel              // 2 - Condizione di errore
    WarnLevel               // 3 - Condizione di avviso
    InfoLevel               // 4 - Messaggio informativo
    DebugLevel              // 5 - Messaggio di debug
    TraceLevel              // 6 - Tracciamento più dettagliato
)
```

### Metodi Level

```go
func (l Level) String() string
```

Restituisce la rappresentazione stringa del livello.

**Esempio:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## Formatter

### Interfaccia Format

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

I formatter personalizzati devono implementare questa interfaccia.

### Formatter Predefinito

Formattatore di testo integrato con opzioni personalizzabili.

```go
type Formatter struct {
    // Opzioni di configurazione
}
```

### Esempio di Formatter JSON

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

// Uso
logger.SetFormatter(&JSONFormatter{})
```

## Writer di Output

### Output File e Rotazione

```go
func GetOutputWriterHourly(filename string) io.Writer
```

Crea un writer che ruota i file di log orariamente.

**Parametri:**

-   `filename`: Il nome base del file di log

**Restituisce:**

-   `io.Writer`: Writer di file con rotazione

**Esempio:**

```go title="Rotazione log oraria"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// Crea file come: app-2024010115.log, app-2024010116.log, etc.
```

### Writer Asincrono

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

Crea un writer asincrono per logging ad alte prestazioni.

**Parametri:**

-   `writer`: Il writer sottostante
-   `bufferSize`: Dimensione del buffer interno

**Restituisce:**

-   `*AsyncWriter`: Istanza del writer asincrono

**Metodi:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**Esempio:**

```go title="Writer asincrono"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## Logging con Contesto

### Interfaccia LoggerWithCtx

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

### Funzioni di Contesto

```go
func SetTrace(traceID string)
func GetTrace() string
```

Imposta e ottiene l'ID di tracciamento per la goroutine corrente.

**Esempio:**

```go
log.SetTrace("trace-123-456")
log.Info("Questo messaggio includerà l'ID di tracciamento")

traceID := log.GetTrace()
fmt.Println("ID di tracciamento corrente:", traceID)
```

## Tag di Build

Questa libreria supporta la compilazione condizionale con tag di build:

:::info Descrizione dei Tag di Build
I tag di build sono specificati tramite il parametro `go build -tags`. Tag diversi cambiano il comportamento di compilazione e le caratteristiche runtime della libreria di log. Scegliere i tag appropriati permette di bilanciare la comodità di sviluppo e le prestazioni di produzione.
:::

### Modalità Predefinita

```bash
go build
```

-   Funzionalità completa attivata
-   Messaggi di debug inclusi
-   Prestazioni standard

### Modalità Debug

```bash
go build -tags debug
```

-   Informazioni di debug avanzate
-   Controlli runtime aggiuntivi
-   Informazioni dettagliate del chiamante

### Modalità Release

```bash
go build -tags release
```

-   Ottimizzato per l'ambiente di produzione
-   Messaggi di debug disattivati
-   Rotazione automatica dei log attivata

### Modalità Discard

```bash
go build -tags discard
```

-   Prestazioni massime
-   Tutte le operazioni di logging sono no-op
-   Zero overhead

### Modalità Combinata

```bash
go build -tags "debug,discard"    # Debug e Discard
go build -tags "release,discard"  # Release e Discard
```

## Ottimizzazione delle Prestazioni

:::tip Best Practices per le Prestazioni
Questa libreria è profondamente ottimizzata attraverso meccanismi come object pool, controlli di livello preventivi e scrittura asincrona. In scenari ad alto throughput, si consiglia di combinare writer asincroni e tag di build appropriati per ottenere le migliori prestazioni.
:::

### Object Pool

La libreria usa internamente `sync.Pool` per gestire:

-   Oggetti di voce di log
-   Buffer di byte
-   Buffer del formattatore

Ciò riduce la pressione del garbage collection in scenari ad alto throughput.

### Controllo del Livello

I controlli del livello di log avvengono prima di operazioni costose:

```go
// Efficiente - formattazione messaggio solo quando livello attivato
logger.Debugf("Risultato operazione costosa: %+v", expensiveCall())

// Meno efficiente quando debug è disattivato in produzione
result := expensiveCall()
logger.Debug("Risultato:", result)
```

### Scrittura Asincrona

Per applicazioni ad alto throughput:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // Buffer grande
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### Ottimizzazione dei Tag di Build

Usa tag di build appropriati per l'ambiente:

-   Sviluppo: Tag predefiniti o debug
-   Produzione: Tag release
-   Critico per le prestazioni: Tag discard

## Esempi

### Uso Base

```go title="Uso base"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("Avvio applicazione")
    log.Warn("Questo è un avviso")
    log.Error("Questo è un errore")
}
```

### Logger Personalizzato

```go title="Configurazione personalizzata"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // Configura logger
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  // Disabilita informazioni del chiamante
    logger.SetPrefixMsg("[MyApp] ")

    // Imposta output su file
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("Logger personalizzato configurato")
    logger.Debug("Informazioni di debug con chiamante")
}
```

### Logging ad Alte Prestazioni

```go title="Logging ad alte prestazioni"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Crea writer con rotazione oraria
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // Usa writer asincrono per migliori prestazioni
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // Salta log di debug in produzione

    // Logging ad alto throughput
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### Logging Context-Aware

```go title="Logging context-aware"
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

    ctxLogger.Info(ctx, "Elaborazione richiesta utente")
    ctxLogger.Debug(ctx, "Validazione completata")
}
```

### Formatter JSON Personalizzato

```go title="Formatter JSON personalizzato"
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
    logger.Caller(true)  // Disabilita informazioni del chiamante
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("Messaggio formattato in JSON")
}
```

## Gestione degli Errori

:::warning Attenzione
Per motivi di prestazioni, la maggior parte dei metodi del logger non restituisce errori. Se la scrittura fallisce, i log verranno silently scartati. Se hai bisogno di error awareness, usa un writer personalizzato.
:::

Se hai bisogno di gestione degli errori per le operazioni di output, implementa un writer personalizzato:

```go title="Writer di cattura errori"
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

## Thread Safety

:::tip Sicurezza della Concorrenza
Tutti i metodi delle istanze `Logger` sono thread-safe e possono essere usati simultaneamente in più goroutine senza meccanismi di sincronizzazione aggiuntivi. Tuttavia, nota che singoli oggetti `Entry` **non** sono thread-safe e sono per uso singolo.
:::

---

## 🌍 Documentazione Multilingua

Questa documentazione è disponibile in più lingue:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/API.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/API.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/API.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/API.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/API.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/API.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/API.md)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/API.md)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/API.md)
-   [🇧🇷 Português](https://lazygophers.github.io/log/pt/API.md)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/API.md)
-   [🇮🇹 Italiano](API.md) (Attuale)

---

**Riferimento completo API LazyGophers Log - Costruisci migliori applicazioni con un eccezionale logging! 🚀**
