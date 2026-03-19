---
titleSuffix: ' | LazyGophers Log'
---
# 📚 Documentação da API

## Visão Geral

O LazyGophers Log fornece uma API abrangente de registro de logs com suporte para múltiplos níveis, formatação personalizada, escrita assíncrona e otimização de tags de compilação. Este documento cobre todas as APIs públicas, opções de configuração e padrões de uso.

## Índice

-   [Tipos Principais](#tipos-principais)
-   [API do Logger](#api-do-logger)
-   [Funções Globais](#funções-globais)
-   [Níveis de Log](#níveis-de-log)
-   [Formatadores](#formatadores)
-   [Writers de Saída](#writers-de-saída)
-   [Logs com Contexto](#logs-com-contexto)
-   [Tags de Compilação](#tags-de-compilação)
-   [Otimização de Desempenho](#otimização-de-desempenho)
-   [Exemplos](#exemplos)

## Tipos Principais

### Logger

A estrutura principal do logger que fornece todas as funcionalidades de registro.

```go
type Logger struct {
    // Contém campos privados para operações thread-safe
}
```

#### Construtor

```go
func New() *Logger
```

Cria uma nova instância do logger com configuração padrão:

-   Nível: `DebugLevel`
-   Saída: `os.Stdout`
-   Formatador: Formatador de texto padrão
-   Rastreamento de chamador: Desativado

**Exemplo:**

```go title="Criando um logger"
logger := log.New()
logger.Info("Novo logger criado")
```

### Entry

Representa uma única entrada de log com todos os metadados associados.

```go
type Entry struct {
    Time       time.Time     // Timestamp de criação da entrada
    Level      Level         // Nível de log
    Message    string        // Mensagem de log
    Pid        int          // ID do processo
    Gid        uint64       // ID da goroutine
    TraceID    string       // ID de rastreamento para tracing distribuído
    CallerName string       // Nome da função chamadora
    CallerFile string       // Caminho do arquivo chamador
    CallerLine int          // Número da linha do chamador
}
```

## API do Logger

### Métodos de Configuração

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Define o nível mínimo de log. Mensagens abaixo deste nível serão ignoradas.

**Parâmetros:**

-   `level`: O nível mínimo de log a ser processado

**Retorno:**

-   `*Logger`: Retorna a si mesmo para suportar encadeamento de métodos

**Exemplo:**

```go title="Definindo o nível de log"
logger.SetLevel(log.InfoLevel)
logger.Debug("Isso não será exibido")  // Ignorado
logger.Info("Isso será exibido")        // Processado
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Define um ou mais destinos de saída para as mensagens de log.

**Parâmetros:**

-   `writers`: Um ou mais destinos de saída `io.Writer`

**Retorno:**

-   `*Logger`: Retorna a si mesmo para suportar encadeamento de métodos

**Exemplo:**

```go title="Definindo destino de saída"
// Saída única
logger.SetOutput(os.Stdout)

// Múltiplas saídas
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

Define um formatador personalizado para a saída de log.

**Parâmetros:**

-   `formatter`: Um formatador que implementa a interface `Format`

**Retorno:**

-   `*Logger`: Retorna a si mesmo para suportar encadeamento de métodos

**Exemplo:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Ativa ou desativa a gravação de informações do chamador nas entradas de log.

**Parâmetros:**

-   `enable`: Se deve ativar as informações do chamador (passe `true` para ativar)

**Retorno:**

-   `*Logger`: Retorna a si mesmo para suportar encadeamento de métodos

**Exemplo:**

```go
logger.EnableCaller(true)
logger.Info("Isso incluirá informações de arquivo:linha")

logger.EnableCaller(false)
logger.Info("Isso não incluirá informações de arquivo:linha")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

Controla as informações do chamador no formatador.

**Parâmetros:**

-   `disable`: Se deve desativar as informações do chamador (passe `true` para desativar)

**Retorno:**

-   `*Logger`: Retorna a si mesmo para suportar encadeamento de métodos

**Exemplo:**

```go
logger.Caller(false)  // Não desativa, mostra informações do chamador
logger.Info("Isso incluirá informações de arquivo:linha")

logger.Caller(true)   // Desativa informações do chamador
logger.Info("Isso não incluirá informações de arquivo:linha")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Define a profundidade da pilha para informações do chamador ao encapsular o logger.

**Parâmetros:**

-   `depth`: O número de frames da pilha para pular

**Retorno:**

-   `*Logger`: Retorna a si mesmo para suportar encadeamento de métodos

**Exemplo:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  // Pula a função wrapper
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Define texto de prefixo ou sufixo para todas as mensagens de log.

**Parâmetros:**

-   `prefix/suffix`: O texto para prefixar/sufixar à mensagem

**Retorno:**

-   `*Logger`: Retorna a si mesmo para suportar encadeamento de métodos

**Exemplo:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  // Saída: [APP] Hello [END]
```

### Métodos de Logging

Todos os métodos de logging têm duas variantes: uma versão simples e uma versão formatada.

#### Nível Trace

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

Registra logs no nível trace (mais detalhado).

**Exemplo:**

```go
logger.Trace("Rastreamento detalhado de execução")
logger.Tracef("Processando item %d de %d", i, total)
```

#### Nível Debug

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

Registra informações de desenvolvimento no nível debug.

**Exemplo:**

```go
logger.Debug("Estado da variável:", variable)
logger.Debugf("Usuário %s autenticado com sucesso", username)
```

#### Nível Info

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

Registra mensagens informativas.

**Exemplo:**

```go
logger.Info("Aplicação iniciada")
logger.Infof("Servidor ouvindo na porta %d", port)
```

#### Nível Warn

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

Registra mensagens de aviso para situações potencialmente problemáticas.

**Exemplo:**

```go
logger.Warn("Função depreciada chamada")
logger.Warnf("Uso de memória alto: %d%%", memoryPercent)
```

#### Nível Error

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

Registra mensagens de erro.

**Exemplo:**

```go
logger.Error("Falha na conexão com banco de dados")
logger.Errorf("Falha ao processar requisição: %v", err)
```

#### Nível Fatal

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

Registra um erro fatal e chama `os.Exit(1)`.

:::danger Operação Destrutiva
`Fatal` e `Fatalf` chamarão `os.Exit(1)` imediatamente após o logging, terminando o processo. Use apenas em situações de erro irrecuperável. Instruções `defer` **não** serão executadas.
:::

**Exemplo:**

```go
logger.Fatal("Erro crítico do sistema")
logger.Fatalf("Não foi possível iniciar o servidor: %v", err)
```

#### Nível Panic

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

Registra uma mensagem de erro e chama `panic()`.

:::danger Operação Destrutiva
`Panic` e `Panicf` chamarão `panic()` após o logging. Ao contrário de `Fatal`, `panic` pode ser recuperado com `recover()`, mas terminará o programa se não for capturado.
:::

**Exemplo:**

```go
logger.Panic("Ocorreu um erro irrecuperável")
logger.Panicf("Estado inválido: %v", state)
```

### Métodos Utilitários

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Cria uma cópia do logger com a mesma configuração.

**Retorno:**

-   `*Logger`: Nova instância do logger com configurações copiadas

**Exemplo:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

Cria um logger com reconhecimento de contexto que aceita `context.Context` como primeiro parâmetro.

**Retorno:**

-   `LoggerWithCtx`: Instância de logger com reconhecimento de contexto

**Exemplo:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "Mensagem com reconhecimento de contexto")
```

## Funções Globais

Funções a nível de pacote que usam o logger global padrão.

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

**Exemplo:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("Usando logger global")
```

## Níveis de Log

### Tipo Level

```go
type Level int8
```

### Níveis Disponíveis

```go
const (
    PanicLevel Level = iota  // 0 - Panic e sai
    FatalLevel              // 1 - Erro fatal e sai
    ErrorLevel              // 2 - Condição de erro
    WarnLevel               // 3 - Condição de aviso
    InfoLevel               // 4 - Mensagem informativa
    DebugLevel              // 5 - Mensagem de debug
    TraceLevel              // 6 - Rastreamento mais detalhado
)
```

### Métodos do Level

```go
func (l Level) String() string
```

Retorna a representação em string do nível.

**Exemplo:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## Formatadores

### Interface Format

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

Formatadores personalizados devem implementar esta interface.

### Formatador Padrão

Formatador de texto embutido com opções personalizáveis.

```go
type Formatter struct {
    // Opções de configuração
}
```

### Exemplo de Formatador JSON

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

## Writers de Saída

### Saída de Arquivo e Rotação

```go
func GetOutputWriterHourly(filename string) io.Writer
```

Cria um writer que rotaciona arquivos de log horariamente.

**Parâmetros:**

-   `filename`: O nome base do arquivo de log

**Retorno:**

-   `io.Writer`: Writer de arquivo com rotação

**Exemplo:**

```go title="Rotação de log horária"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// Cria arquivos como: app-2024010115.log, app-2024010116.log, etc.
```

### Writer Assíncrono

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

Cria um writer assíncrono para logging de alto desempenho.

**Parâmetros:**

-   `writer`: O writer subjacente
-   `bufferSize`: Tamanho do buffer interno

**Retorno:**

-   `*AsyncWriter`: Instância de writer assíncrono

**Métodos:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**Exemplo:**

```go title="Writer assíncrono"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## Logs com Contexto

### Interface LoggerWithCtx

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

### Funções de Contexto

```go
func SetTrace(traceID string)
func GetTrace() string
```

Define e obtém o ID de rastreamento para a goroutine atual.

**Exemplo:**

```go
log.SetTrace("trace-123-456")
log.Info("Esta mensagem incluirá o ID de rastreamento")

traceID := log.GetTrace()
fmt.Println("ID de rastreamento atual:", traceID)
```

## Tags de Compilação

Esta biblioteca suporta compilação condicional com tags de compilação:

:::info Descrição das Tags de Compilação
Tags de compilação são especificadas através do parâmetro `go build -tags`. Diferentes tags alteram o comportamento de compilação e as características de runtime da biblioteca de log. Escolher as tags apropriadas permite equilibrar a conveniência de desenvolvimento e o desempenho de produção.
:::

### Modo Padrão

```bash
go build
```

-   Funcionalidade completa ativada
-   Mensagens de debug incluídas
-   Desempenho padrão

### Modo Debug

```bash
go build -tags debug
```

-   Informações de debug aprimoradas
-   Verificações adicionais de runtime
-   Informações detalhadas do chamador

### Modo Release

```bash
go build -tags release
```

-   Otimizado para ambiente de produção
-   Mensagens de debug desativadas
-   Rotação automática de logs ativada

### Modo Discard

```bash
go build -tags discard
```

-   Desempenho máximo
-   Todas as operações de log são no-ops
-   Zero overhead

### Modo Combinado

```bash
go build -tags "debug,discard"    # Debug e Discard
go build -tags "release,discard"  # Release e Discard
```

## Otimização de Desempenho

:::tip Melhores Práticas de Desempenho
Esta biblioteca é profundamente otimizada através de mecanismos como pools de objetos, verificações de nível antecipadas e escrita assíncrona. Em cenários de alta taxa de transferência, recomenda-se combinar writers assíncronos e tags de compilação apropriadas para obter o melhor desempenho.
:::

### Pools de Objetos

A biblioteca usa internamente `sync.Pool` para gerenciar:

-   Objetos de entrada de log
-   Buffers de bytes
-   Buffers de formatador

Isso reduz a pressão de coleta de lixo em cenários de alta taxa de transferência.

### Verificação de Nível

As verificações de nível de log ocorrem antes de operações caras:

```go
// Eficiente - formatação de mensagem apenas quando nível ativado
logger.Debugf("Resultado de operação cara: %+v", expensiveCall())

// Menos eficiente quando debug está desativado em produção
result := expensiveCall()
logger.Debug("Resultado:", result)
```

### Escrita Assíncrona

Para aplicações de alta taxa de transferência:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  // Buffer grande
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### Otimização de Tags de Compilação

Use tags de compilação apropriadas para o ambiente:

-   Desenvolvimento: Tags padrão ou debug
-   Produção: Tags de release
-   Crítico para desempenho: Tags de descarte

## Exemplos

### Uso Básico

```go title="Uso básico"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("Iniciando aplicação")
    log.Warn("Este é um aviso")
    log.Error("Este é um erro")
}
```

### Logger Personalizado

```go title="Configuração personalizada"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    // Configurar logger
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  // Desativar informações do chamador
    logger.SetPrefixMsg("[MyApp] ")

    // Definir saída para arquivo
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("Logger personalizado configurado")
    logger.Debug("Informações de debug com chamador")
}
```

### Logging de Alto Desempenho

```go title="Logging de alto desempenho"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Criar writer com rotação horária
    writer := log.GetOutputWriterHourly("./logs/app.log")

    // Usar writer assíncrono para melhor desempenho
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  // Pular logs de debug em produção

    // Logging de alta taxa de transferência
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### Logging com Reconhecimento de Contexto

```go title="Logging com reconhecimento de contexto"
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

    ctxLogger.Info(ctx, "Processando requisição do usuário")
    ctxLogger.Debug(ctx, "Validação concluída")
}
```

### Formatador JSON Personalizado

```go title="Formatador JSON personalizado"
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
    logger.Caller(true)  // Desativar informações do chamador
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("Mensagem formatada em JSON")
}
```

## Tratamento de Erros

:::warning Aviso
Por motivos de desempenho, a maioria dos métodos do logger não retorna erros. Se a gravação falhar, os logs serão silenciosamente descartados. Se você precisar de consciência de erros, use um writer personalizado.
:::

Se você precisar de tratamento de erros para operações de saída, implemente um writer personalizado:

```go title="Writer de captura de erros"
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

:::tip Segurança de Concorrência
Todos os métodos de instâncias `Logger` são thread-safe e podem ser usados simultaneamente em múltiplas goroutines sem mecanismos adicionais de sincronização. No entanto, observe que objetos `Entry` individuais **não** são thread-safe e são para uso único.
:::

---

## 🌍 Documentação Multilíngue

Esta documentação está disponível em vários idiomas:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/API.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/API.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/API.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/API.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/API.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/API.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/API.md)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/API.md)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/API.md)
-   [🇧🇷 Português](API.md) (Atual)

---

**Referência completa da API LazyGophers Log - Construa melhores aplicativos com logging excepcional! 🚀**
