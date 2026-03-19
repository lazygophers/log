---
titleSuffix: ' | LazyGophers Log'
---
# 🤝 Contribuindo para LazyGophers Log

Agradecemos muito sua contribuição! Queremos tornar a contribuição para LazyGophers Log o mais simples e transparente possível, seja:

-   🐛 Relatando bugs
-   💬 Discutindo o estado atual do código
-   ✨ Enviando solicitações de recursos
-   🔧 Propor soluções
-   🚀 Implementando novos recursos

## 📋 Sumário

-   [Código de Conduta](#-código-de-conduta)
-   [Processo de Desenvolvimento](#-processo-de-desenvolvimento)
-   [Primeiros Passos](#-primeiros-passos)
-   [Processo de Pull Request](#-processo-de-pull-request)
-   [Padrões de Codificação](#-padrões-de-codificação)
-   [Diretrizes de Teste](#-diretrizes-de-teste)
-   [Requisitos de Build Tags](#️-requisitos-de-build-tags)
-   [Documentação](#-documentação)
-   [Diretrizes de Issues](#-diretrizes-de-issues)
-   [Considerações de Performance](#-considerações-de-performance)
-   [Diretrizes de Segurança](#-diretrizes-de-segurança)
-   [Comunidade](#-comunidade)

## 📜 Código de Conduta

Este projeto e todos os participantes estão sujeitos ao nosso [Código de Conduta](/pt/CODE_OF_CONDUCT). Ao participar, você concorda em cumprir as regras.

## 🔄 Processo de Desenvolvimento

Usamos GitHub para hospedar código, rastrear issues e solicitações de recursos, e aceitar pull requests.

### Fluxo de Trabalho

:::note Visão geral do processo de desenvolvimento
1. **Fork** o repositório
2. **Clone** seu fork localmente
3. **Crie** um branch de recurso a partir de `master`
4. **Faça** suas alterações
5. **Teste** completamente em todas as build tags
6. **Envie** um pull request
:::

## 🚀 Primeiros Passos

### Pré-requisitos

-   **Go 1.21+** - [Instalar Go](https://golang.org/doc/install)
-   **Git** - [Instalar Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
-   **Make** (opcional mas recomendado)

### Configuração de Desenvolvimento Local

```bash title="Clonar repositório e configurar ambiente de desenvolvimento"
# 1. Fork o repositório no GitHub
# 2. Clone seu fork
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. Adicione o remoto upstream
git remote add upstream https://github.com/lazygophers/log.git

# 4. Instale dependências
go mod tidy

# 5. Verifique a instalação
make test-quick
```

### Configuração de Ambiente

:::info Configuração de Ambiente
Certifique-se de que as variáveis de ambiente Go estejam configuradas corretamente e as ferramentas de desenvolvimento recomendadas instaladas para a melhor experiência de desenvolvimento.
:::

```bash title="Configuração de ambiente"
# Configure o ambiente Go (se ainda não configurado)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Opcional: Instale ferramentas úteis
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 Processo de Pull Request

### Antes do Envio

1.  **Pesquise** PRs existentes para evitar duplicatas
2.  **Teste** suas alterações em todas as configurações de build
3.  **Documente** quaisquer breaking changes
4.  **Atualize** a documentação relevante
5.  **Adicione** testes para novos recursos

### Lista de Verificação de PR

:::warning Confirme todos os itens antes de enviar o PR
PRs que não atendem aos requisitos da lista de verificação não serão mesclados.
:::

-   [ ] **Qualidade do Código**

    -   [ ] Código segue o guia de estilo do projeto
    -   [ ] Sem novos avisos de lint
    -   [ ] Tratamento de erros correto
    -   [ ] Algoritmos e estruturas de dados eficientes

-   [ ] **Testes**

    -   [ ] Todos os testes existentes passam: `make test`
    -   [ ] Novos testes para novos recursos adicionados
    -   [ ] Cobertura de testes mantida ou melhorada
    -   [ ] Todas as build tags testadas: `make test-all`

-   [ ] **Documentação**

    -   [ ] Código tem comentários apropriados
    -   [ ] Documentação da API atualizada (se necessário)
    -   [ ] README atualizado (se necessário)
    -   [ ] Documentação multilíngue atualizada (se for voltada ao usuário)

-   [ ] **Compatibilidade de Build**
    -   [ ] Modo padrão: `go build`
    -   [ ] Modo debug: `go build -tags debug`
    -   [ ] Modo release: `go build -tags release`
    -   [ ] Modo discard: `go build -tags discard`
    -   [ ] Modos de combinação testados

### Modelo de PR

Ao enviar um pull request, use nosso [modelo de PR](https://github.com/lazygophers/log/blob/main/.github/pull_request_template.md).

## 📏 Padrões de Codificação

### Guia de Estilo Go

:::tip Padrões de código Go
Seguimos o guia de estilo Go padrão com alguns complementos. Certifique-se de que a formatação do código passa nas verificações `go fmt` e `goimports`.
:::

```go
// ✅ Good
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ Bad
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### Convenções de Nomenclatura

-   **Pacotes**: Curto, minúsculas, preferencialmente uma única palavra
-   **Funções**: PascalCase, descritivo
-   **Variáveis**: camelCase para locais, PascalCase para exportadas
-   **Constantes**: PascalCase para exportadas, camelCase para não exportadas
-   **Interfaces**: Geralmente terminam em "er" (ex: `Writer`, `Formatter`)

### Organização do Código

```
project/
├── docs/           # Documentação multilíngue
├── .github/        # Modelos e workflows do GitHub
├── logger.go       # Implementação principal do logger
├── entry.go        # Estrutura de entrada de log
├── level.go        # Níveis de log
├── formatter.go    # Formatação de logs
├── output.go       # Gerenciamento de saída
└── *_test.go      # Testes co-localizados com código fonte
```

### Tratamento de Erros

:::tip Melhores práticas de tratamento de erros
Código de biblioteca deve retornar erros, não panic, permitindo que o chamador decida como lidar com situações excepcionais.
:::

```go title="Exemplo de tratamento de erros"
// ✅ Recomendado: Retorne erros, não panic
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ Evite: Usar panic em código de biblioteca
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // Não faça isso
    }
    return &Logger{...}
}
```

## 🧪 Diretrizes de Teste

### Estrutura de Teste

```go title="Exemplo de teste tabular"
func TestLogger_Info(t *testing.T) {
    tests := []struct {
        name     string
        level    Level
        message  string
        expected bool
    }{
        {"info level allows info", InfoLevel, "test", true},
        {"warn level blocks info", WarnLevel, "test", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Implementação do teste
        })
    }
}
```

### Requisitos de Cobertura

:::warning Requisito de cobertura
PRs com cobertura abaixo de 90% para código novo não passarão nas verificações de CI.
:::

-   **Mínimo**: 90% de cobertura para código novo
-   **Meta**: 95%+ de cobertura geral
-   **Todas as build tags** devem manter cobertura
-   Use `make coverage-all` para verificar

### Comandos de Teste

```bash title="Executar testes"
# Testes rápidos em todas as build tags
make test-quick

# Suíte completa de testes com cobertura
make test-all

# Relatório de cobertura
make coverage-html

# Benchmarks
make benchmark
```

## 🏗️ Requisitos de Build Tags

:::warning Compatibilidade de Build
Todas as alterações devem ser compatíveis com nosso sistema de build tags. Código que não passa em todos os testes de build tags não será mesclado.
:::

### Build Tags Suportadas

-   **Padrão** (`go build`): Funcionalidade completa
-   **Debug** (`go build -tags debug`): Recursos de depuração aprimorados
-   **Release** (`go build -tags release`): Otimizações de produção
-   **Discard** (`go build -tags discard`): Performance máxima

### Testes de Build Tags

:::info Descrição de Build Tags
O projeto usa build tags para compilação condicional, com tags diferentes correspondendo a diferentes modos de execução. Teste em todas as tags antes de enviar.
:::

```bash title="Testes de build tags"
# Teste cada configuração de build
make test-default
make test-debug
make test-release
make test-discard

# Teste combinações
make test-debug-discard
make test-release-discard

# Teste tudo de uma vez
make test-all
```

### Diretrizes de Build Tags

```go
//go:build debug
// +build debug

package log

// Implementação específica de debug
```

## 📚 Documentação

### Documentação de Código

-   **Todas as funções exportadas** devem ter comentários claros
-   **Algoritmos complexos** precisam de explicações
-   **Exemplos** para uso não trivial
-   **Thread-safety** notas (se aplicável)

```go
// SetLevel define o nível mínimo de logging.
// Logs abaixo deste nível serão ignorados.
// Este método é thread-safe.
//
// Example:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // Não será exibido
//   logger.Info("visible")   // Será exibido
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### Atualizações do README

Ao adicionar recursos, atualize:

-   README.md principal
-   Todos os READMEs específicos de idioma em `docs/`
-   Exemplos de código
-   Lista de recursos

## 🐛 Diretrizes de Issues

### Relatórios de Bugs

Use o [modelo de relatório de bug](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/bug_report.md) e inclua:

-   **Descrição clara do problema**
-   **Passos para reproduzir**
-   **Comportamento esperado vs. real**
-   **Detalhes do ambiente** (sistema operacional, versão Go, build tags)
-   **Exemplo mínimo de código**

### Solicitações de Recursos

Use o [modelo de solicitação de recurso](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/feature_request.md) e inclua:

-   **Motivação clara do recurso**
-   **Proposta de API** design
-   **Considerações de implementação**
-   **Análise de breaking changes**

### Perguntas

Use o [modelo de pergunta](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/question.md) para:

-   Problemas de uso
-   Ajuda com configuração
-   Melhores práticas
-   Orientação de integração

## 🚀 Considerações de Performance

### Benchmarking

Sempre faça benchmark de alterações sensíveis à performance:

```bash title="Executar benchmarks"
# Execute benchmarks
go test -bench=. -benchmem

# Compare performance antes e depois
go test -bench=. -benchmem > before.txt
# Faça as alterações
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### Diretrizes de Performance

:::tip Pontos de otimização de performance
Esta é uma biblioteca de logging sensível à performance. Qualquer alteração deve considerar o impacto no hot path.
:::

-   **Minimize** alocações de memória no hot path
-   **Use pools de objetos** para objetos frequentemente criados
-   **Retorno antecipado** para níveis de log desabilitados
-   **Evite reflexão** em código crítico de performance
-   **Profile antes de otimizar**

### Gerenciamento de Memória

```go
// ✅ Recomendado: Usar pool de objetos
var entryPool = sync.Pool{
    New: func() interface{} {
        return &Entry{}
    },
}

func getEntry() *Entry {
    return entryPool.Get().(*Entry)
}

func putEntry(e *Entry) {
    e.Reset()
    entryPool.Put(e)
}
```

## 🔒 Diretrizes de Segurança

### Dados Sensíveis

:::warning Aviso de segurança
Vazar dados sensíveis em logs pode levar a incidentes graves de segurança. Certifique-se de seguir os seguintes padrões.
:::

-   **Nunca registre** senhas, tokens ou dados sensíveis
-   **Sanitize** entrada do usuário em mensagens de log
-   **Evite** registrar corpos de request/response completos
-   **Use** logging estruturado para melhor controle

```go
// ✅ Recomendado
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ Evite
logger.Infof("User login: %+v", userRequest) // Pode conter senhas
```

### Dependências

-   Mantenha dependências **atualizadas**
-   **Revise cuidadosamente** novas dependências
-   **Minimize** dependências externas
-   **Use** `go mod verify` para verificar integridade

## 👥 Comunidade

### Obter Ajuda

-   📖 [Documentação](README.md)
-   💬 [Discussões GitHub](https://github.com/lazygophers/log/discussions)
-   🐛 [Rastreador de Issues](https://github.com/lazygophers/log/issues)
-   📧 E-mail: support@lazygophers.com

### Diretrizes de Comunicação

-   **Mantenha-se respeitoso** e inclusivo
-   **Pesquise antes** de perguntar
-   **Forneça contexto** ao pedir ajuda
-   **Ajude outros** quando puder
-   **Siga** o [Código de Conduta](/pt/CODE_OF_CONDUCT)

## 🎯 Reconhecimento

Contribuidores são reconhecidos das seguintes formas:

-   **Seção README de** contribuidores
-   **Notas de Release** menções
-   **Gráfico de contribuidores** GitHub
-   **Posts de agradecimento** da comunidade

## 📝 Licença

Ao contribuir, você concorda que suas contribuições serão licenciadas sob a licença MIT.

---

## 🌍 Documentação Multilíngue

Este documento está disponível em vários idiomas:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CONTRIBUTING.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/CONTRIBUTING.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CONTRIBUTING.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CONTRIBUTING.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CONTRIBUTING.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CONTRIBUTING.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CONTRIBUTING.md)
-   [🇵🇹 Português](/pt/CONTRIBUTING)（atual）

---

**Obrigado por contribuir com LazyGophers Log!🚀**
