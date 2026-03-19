---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Uma biblioteca de registro Go de alto desempenho e flexível, construída sobre zap, fornecendo recursos ricos e uma API simples.

## 📖 Idiomas da documentação

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 中文 simplificado](README.md) (atual)
-   [🇹🇼 中文 tradicional](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)

## ✨ Recursos

-   **🚀 Alto desempenho**：Construído sobre zap com pool de objetos e registro condicional de campos
-   **📊 Níveis de registro ricos**：Níveis Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configuração flexível**：
    -   Controle do nível de registro
    -   Registro de informações do chamador
    -   Informações de rastreamento (incluindo ID de goroutine)
    -   Prefixos e sufixos personalizados
    -   Destinos de saída personalizados (console, arquivos, etc.)
    -   Opções de formato de registro
-   **🔄 Rotação de arquivos**：Suporte para rotação de arquivos de registro por hora
-   **🔌 Compatibilidade com Zap**：Integração perfeita com zap WriteSyncer
-   **🎯 API simples**：API clara semelhante à biblioteca de registro padrão, fácil de usar

## 🚀 Início rápido

### Instalação

:::tip Instalação
```bash
go get github.com/lazygophers/log
```
:::

### Uso básico

```go title="Início rápido"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Usar logger global padrão
    log.Debug("Mensagem de depuração")
    log.Info("Mensagem informativa")
    log.Warn("Mensagem de aviso")
    log.Error("Mensagem de erro")

    // Usar saída formatada
    log.Infof("Usuário %s entrou com sucesso", "admin")

    // Configuração personalizada
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Este é um registro do logger personalizado")
}
```

## 📚 Uso avançado

### Logger personalizado com saída de arquivo

```go title="Configuração de saída de arquivo"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Criar logger com saída de arquivo
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Registro de depuração com informações do chamador")
    logger.Info("Registro informativo com informações de rastreamento")
}
```

### Controle do nível de registro

```go title="Controle do nível de registro"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Apenas warn e acima serão registrados
    logger.Debug("Isso não será registrado")  // Ignorado
    logger.Info("Isso não será registrado")   // Ignorado
    logger.Warn("Isso será registrado")    // Registrado
    logger.Error("Isso será registrado")   // Registrado
}
```

## 🎯 Cenários de uso

### Cenários aplicáveis

-   **Serviços web e backends API**：Rastreamento de solicitações, registro estruturado, monitoramento de desempenho
-   **Arquitetura de microsserviços**：Rastreamento distribuído (TraceID), formato de registro unificado, alta taxa de transferência
-   **Ferramentas de linha de comando**：Controle de níveis, saída limpa, relatórios de erros
-   **Tarefas em lote**：Rotação de arquivos, execução longa, otimização de recursos

### Vantagens especiais

-   **Otimização com pool de objetos**：Reutilização de objetos Entry e Buffer, reduzindo a pressão do GC
-   **Escrita assíncrona**：Alta taxa de transferência (10000+ registros/segundo) sem bloqueio
-   **Suporte para TraceID**：Rastreamento de solicitações em sistemas distribuídos, integração com OpenTelemetry
-   **Início sem configuração**：Pronto para usar, configuração progressiva

## 🔧 Opções de configuração

:::note Opções de configuração
Todos os seguintes métodos suportam chamada em cadeia e podem ser combinados para construir um Logger personalizado.
:::

### Configuração do Logger

| Método                  | Descrição                | Padrão      |
| --------------------- | ------------------- | ------------ |
| `SetLevel(level)`       | Definir nível mínimo de registro     | `DebugLevel` |
| `EnableCaller(enable)`  | Ativar/desativar informações do chamador | `false`      |
| `EnableTrace(enable)`   | Ativar/desativar informações de rastreamento  | `false`      |
| `SetCallerDepth(depth)` | Definir profundidade do chamador   | `2`          |
| `SetPrefixMsg(prefix)`  | Definir prefixo de registro  | `""`         |
| `SetSuffixMsg(suffix)`  | Definir sufixo de registro  | `""`         |
| `SetOutput(writers...)` | Definir destinos de saída         | `os.Stdout`  |

### Níveis de registro

| Nível        | Descrição                        |
| ----------- | --------------------------- |
| `TraceLevel` | Mais detalhado, para rastreamento detalhado        |
| `DebugLevel` | Informações de depuração                  |
| `InfoLevel`  | Informações gerais                    |
| `WarnLevel`  | Mensagens de aviso                  |
| `ErrorLevel` | Mensagens de erro                  |
| `FatalLevel` | Erros fatais (chama os.Exit(1))    |
| `PanicLevel` | Erros de pânico (chama panic())      |

## 🏗️ Arquitetura

### Componentes principais

-   **Logger**：Estrutura principal de registro com opções configuráveis
-   **Entry**：Registro individual com suporte abrangente de campos
-   **Level**：Definições de níveis de registro e funções utilitárias
-   **Format**：Interface de formatação de registro e implementações

### Otimização de desempenho

-   **Pool de objetos**：Reutiliza objetos Entry para reduzir alocação de memória
-   **Registro condicional**：Registra apenas campos caros quando necessário
-   **Verificação rápida de níveis**：Verifica o nível de registro na camada mais externa
-   **Design sem bloqueio**：A maioria das operações não requer bloqueios

## 📊 Comparação de desempenho

:::info Comparação de desempenho
Os dados a seguir são baseados em benchmarks; o desempenho real pode variar dependendo do ambiente e da configuração.
:::

| Recurso          | lazygophers/log | zap    | logrus | registro padrão |
| ------------- | --------------- | ------ | ------ | -------- |
| Desempenho      | Alto              | Alto     | Médio     | Baixo       |
| Simplicidade da API    | Alto              | Médio     | Alto     | Alto       |
| Riqueza de recursos    | Médio          | Alto     | Alto     | Baixo       |
| Flexibilidade      | Médio          | Alto     | Alto     | Baixo       |
| Curva de aprendizado      | Baixo              | Médio     | Médio     | Baixo       |

## ❓ Perguntas frequentes

### Como escolher o nível de registro adequado?

-   **Ambiente de desenvolvimento**：Use `DebugLevel` ou `TraceLevel` para obter informações detalhadas
-   **Ambiente de produção**：Use `InfoLevel` ou `WarnLevel` para reduzir sobrecarga
-   **Testes de desempenho**：Use `PanicLevel` para desativar todos os registros

### Como otimizar o desempenho em produção?

:::warning Nota
Em cenários de alta taxa de transferência, recomenda-se combinar escrita assíncrona com níveis de registro razoáveis para otimizar o desempenho.
:::

1. Use `AsyncWriter` para escrita assíncrona：

```go title="Configuração de escrita assíncrona"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Ajuste os níveis de registro para evitar registros desnecessários：

```go title="Otimização de nível"
logger.SetLevel(log.InfoLevel)  // Pular Debug e Trace
```

3. Use registros condicionais para reduzir sobrecarga：

```go title="Registros condicionais"
if logger.Level >= log.DebugLevel {
    logger.Debug("Informações de depuração detalhadas")
}
```

### Qual é a diferença entre `Caller` e `EnableCaller`?

-   **`EnableCaller(enable bool)`**：Controla se o Logger coleta informações do chamador
    -   `EnableCaller(true)` ativa a coleta de informações do chamador
-   **`Caller(disable bool)`**：Controla se o Formatter exibe informações do chamador
    -   `Caller(true)` desativa a exibição de informações do chamador

Recomenda-se usar `EnableCaller` para controle global.

### Como implementar um formatador personalizado?

Implemente a interface `Format`：

```go title="Formatador personalizado"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Documentação relacionada

-   [📚 Documentação da API](API.md) - Referência completa da API
-   [🤝 Guia de contribuição](/pt/CONTRIBUTING) - Como contribuir
-   [📋 Log de alterações](/pt/CHANGELOG) - Histórico de versões
-   [🔒 Política de segurança](/pt/SECURITY) - Diretrizes de segurança
-   [📜 Código de conduta](/pt/CODE_OF_CONDUCT) - Diretrizes da comunidade

## 🚀 Obter ajuda

-   **GitHub Issues**：[Relatar bugs ou solicitar recursos](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[Documentação da API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Exemplos**：[Exemplos de uso](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licença

Este projeto está licenciado sob a Licença MIT - consulte o arquivo [LICENSE](/pt/LICENSE) para obter detalhes.

## 🤝 Contribuindo

Agradecemos contribuições! Consulte nosso [Guia de contribuição](/pt/CONTRIBUTING) para obter mais informações.

---

**lazygophers/log** foi projetado para ser a solução de registro preferida para desenvolvedores Go que valorizam tanto desempenho quanto simplicidade. Seja construindo um pequeno utilitário ou um sistema distribuído em grande escala, esta biblioteca fornece o equilíbrio certo entre funcionalidade e facilidade de uso.
