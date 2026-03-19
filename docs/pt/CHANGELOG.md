---
titleSuffix: " | LazyGophers Log"
---
# 📋 Registro de Alterações

Todas as alterações importantes deste projeto serão registradas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) e este projeto adere ao [Versionamento Semântico](https://semver.org/spec/v2.0.0.html).

## [Não Lançado]

### Novo

-   Documentação multilíngue completa (7 idiomas)
-   Modelos de problemas do GitHub (relatório de bugs, solicitação de recursos, perguntas)
-   Modelo de Pull Request com verificação de compatibilidade de tags de compilação
-   Guia de contribuição multilíngue
-   Política de segurança com processo de relatório de vulnerabilidades
-   Documentação completa de API com exemplos
-   Estrutura e modelos profissionais de projeto

### Alterado

-   README melhorado com documentação abrangente de recursos
-   Cobertura de testes aprimorada para todas as configurações de tags de compilação
-   Estrutura do projeto atualizada para melhor manutenibilidade

### Documentação

-   Adicionado suporte multilíngue a todos os documentos principais
-   Criado referência abrangente de API
-   Estabelecido guia de fluxo de trabalho de contribuição
-   Implementado processo de relatório de segurança

## [1.0.0] - 2024-01-01

### Novo

-   Funcionalidade principal de registro com múltiplos níveis (Trace, Debug, Info, Warn, Error, Fatal, Panic)
-   Implementação de logger thread-safe com pool de objetos
-   Suporte a tags de compilação (padrão, debug, release, descartar)
-   Interface de formatador personalizado com formatador de texto padrão
-   Suporte a múltiplas saídas
-   Funcionalidade de gravação assíncrona para cenários de alta vazão
-   Rotação automática de arquivo de log por hora
-   Registro com reconhecimento de contexto com rastreamento de ID de goroutine e rastreamento
-   Informações de chamador com profundidade de pilha configurável
-   Funções de conveniência de nível de pacote global
-   Suporte de integração de logger Zap

### Desempenho

-   Uso de `sync.Pool` para pool de objetos de entrada e buffers
-   Verificação antecipada de nível para evitar operações caras
-   Gravador assíncrono para gravação de log sem bloqueio
-   Otimizações de tags de compilação para diferentes ambientes

### Tags de Compilação

-   **Padrão**: Funcionalidade completa com mensagens de depuração
-   **Depuração**: Informações de depuração aprimoradas e detalhes do chamador
-   **Lançamento**: Otimizado para produção, mensagens de depuração desabilitadas
-   **Descartar**: Desempenho máximo, operações de log sem efeito

### Funcionalidades Principais

-   **Logger**: Estrutura principal de logger com nível, saída e formatador configuráveis
-   **Entrada**: Estrutura de registro de log com metadados abrangentes
-   **Níveis**: Sete níveis de log de Panic (mais alto) para Trace (mais baixo)
-   **Formatadores**: Sistema de formatação conectável
-   **Gravadores**: Suporte a rotação de arquivo e gravação assíncrona
-   **Contexto**: Suporte a ID de Goroutine e rastreamento distribuído

### Destaques da API

-   API de configuração fluida com encadeamento de métodos
-   Métodos de registro simples e formatados (`.Info()` e `.Infof()`)
-   Isolamento de configuração com clonagem de logger
-   Registro com reconhecimento de contexto com `CloneToCtx()`
-   Personalização de mensagens de prefixo e sufixo
-   Alternância de informações do chamador

### Testes

-   Suíte de testes abrangente com cobertura de 93,5%
-   Suporte a testes de múltiplas tags de compilação
-   Fluxos de trabalho de teste automatizados
-   Benchmarks de desempenho

## [0.9.0] - 2023-12-15

### Novo

-   Estrutura inicial do projeto
-   Funcionalidade básica de registro
-   Filtragem baseada em nível
-   Suporte de saída de arquivo

### Alterado

-   Melhorias de desempenho por meio de pool de objetos
-   Tratamento de erros aprimorado

## [0.8.0] - 2023-12-01

### Novo

-   Suporte a múltiplos gravadores
-   Interface de formatador personalizado
-   Funcionalidade de gravação assíncrona

### Corrigido

-   Vazamento de memória em cenários de alta vazão
-   Condições de corrida no acesso concorrente

## [0.7.0] - 2023-11-15

### Novo

-   Suporte a tags de compilação para compilação condicional
-   Registro de níveis Trace e Debug
-   Rastreamento de informações do chamador

### Alterado

-   Otimização de padrões de alocação de memória
-   Melhoria na segurança de threads

## [0.6.0] - 2023-11-01

### Novo

-   Funcionalidade de rotação de logs
-   Registro com reconhecimento de contexto
-   Rastreamento de ID de Goroutine

### Obsoleto

-   Métodos de configuração antigos (serão removidos no v1.0.0)

## [0.5.0] - 2023-10-15

### Novo

-   Formatador JSON
-   Múltiplos destinos de saída
-   Benchmarks de desempenho

### Alterado

-   Refatoração do mecanismo de registro principal
-   Melhoria na consistência da API

### Removido

-   Métodos de registro antigos

## [0.4.0] - 2023-10-01

### Novo

-   Registro de níveis Fatal e Panic
-   Funções de pacote global
-   Validação de configuração

### Corrigido

-   Problemas de sincronização de saída
-   Otimização de uso de memória

## [0.3.0] - 2023-09-15

### Novo

-   Níveis de log personalizados
-   Interface de formatador
-   Operações thread-safe

### Alterado

-   Simplificação do design da API
-   Documentação aprimorada

## [0.2.0] - 2023-09-01

### Novo

-   Suporte de saída de arquivo
-   Filtragem baseada em nível
-   Opções básicas de formatação

### Corrigido

-   Gargalos de desempenho
-   Vazamentos de memória

## [0.1.0] - 2023-08-15

### Novo

-   Lançamento inicial
-   Registro básico de console
-   Suporte simples de níveis (Info, Warn, Error)
-   Estrutura principal do logger

## Resumo do Histórico de Versões

| Versão  | Data de Lançamento   | Principais Recursos                                       |
| ----- | ---------- | ---------------------------------------------- |
| 1.0.0 | 2024-01-01 | Sistema completo de registro, tags de compilação, gravação assíncrona, documentação abrangente |
| 0.9.0 | 2023-12-15 | Melhorias de desempenho, pool de objetos                               |
| 0.8.0 | 2023-12-01 | Múltiplos gravadores, gravação assíncrona, formatadores personalizados             |
| 0.7.0 | 2023-11-15 | Tags de compilação, níveis Trace/Debug, informações do chamador           |
| 0.6.0 | 2023-11-01 | Rotação de logs, registro de contexto, rastreamento de goroutine           |
| 0.5.0 | 2023-10-15 | Formatador JSON, múltiplas saídas, benchmarks              |
| 0.4.0 | 2023-10-01 | Níveis Fatal/Panic, funções globais                     |
| 0.3.0 | 2023-09-15 | Níveis personalizados, interface de formatador                       |
| 0.2.0 | 2023-09-01 | Saída de arquivo, filtragem de nível                             |
| 0.1.0 | 2023-08-15 | Lançamento inicial, registro básico de console                       |

## Guia de Migração

### Migrando do v0.9.x para v1.0.0

#### Mudanças Interruptivas

-   Nenhuma - v1.0.0 é compatível com v0.9.x

#### Novos Recursos Disponíveis

-   Suporte aprimorado a tags de compilação
-   Documentação abrangente
-   Modelos profissionais de projeto
-   Processo de relatório de segurança

#### Atualizações Recomendadas

```go
// Maneira antiga (ainda suportada)
logger := log.New()
logger.SetLevel(log.InfoLevel)

// Nova maneira recomendada, usando encadeamento de métodos
logger := log.New().
    SetLevel(log.InfoLevel).
    Caller(true).
    SetPrefixMsg("[MyApp] ")
```

### Migrando do v0.8.x para v0.9.x

#### Mudanças Interruptivas

-   Removidos métodos de configuração obsoletos
-   Alterado gerenciamento de buffer interno

#### Passos de Migração

1. Atualizar caminhos de importação, se necessário
2. Substituir métodos obsoletos:

    ```go
    // Antigo (obsoleto)
    logger.SetOutputFile("app.log")

    // Novo
    file, _ := os.Create("app.log")
    logger.SetOutput(file)
    ```

### Migrando do v0.5.x ou anterior

#### Mudanças Principais

-   API completamente reprojetada para melhor consistência
-   Desempenho aprimorado por meio de pool de objetos
-   Novo sistema de tags de compilação

#### Migração Necessária

-   Atualizar todas as chamadas de log para a nova API
-   Revisar e atualizar implementações de formatador
-   Testar com novas configurações de tags de compilação

## Marcos de Desenvolvimento

### 🎯 Roadmap v1.1.0 (Planejado)

-   [ ] Registro estruturado com pares chave-valor
-   [ ] Amostragem de log para cenários de alto volume
-   [ ] Sistema de plugins para saída personalizada
-   [ ] Métricas de desempenho aprimoradas
-   [ ] Integração com logs em nuvem

### 🎯 Roadmap v1.2.0 (Futuro)

-   [ ] Suporte a arquivos de configuração (YAML/JSON/TOML)
-   [ ] Agregação e filtragem de logs
-   [ ] Streaming de logs em tempo real
-   [ ] Funcionalidades de segurança aprimoradas
-   [ ] Integração com dashboard de desempenho

## Contribuindo

Contribuições são bem-vindas! Por favor, consulte nosso [Guia de Contribuição](/pt/CONTRIBUTING) para detalhes sobre:

-   Relatar bugs e solicitar recursos
-   Processo de submissão de código
-   Configuração de desenvolvimento
-   Requisitos de teste
-   Padrões de documentação

## Segurança

Para vulnerabilidades de segurança, consulte nossa [Política de Segurança](/pt/SECURITY) para:

-   Versões suportadas
-   Processo de relatório
-   Linha do tempo de resposta
-   Melhores práticas de segurança

## Suporte

-   📖 [Documentação](docs/)
-   🐛 [Rastreador de Problemas](https://github.com/lazygophers/log/issues)
-   💬 [Discussões](https://github.com/lazygophers/log/discussions)
-   📧 Email: support@lazygophers.com

## Licença

Este projeto está licenciado sob a Licença MIT - consulte o arquivo [LICENSE](/pt/LICENSE) para detalhes.

---

## 🌍 Documentação Multilíngue

Este registro de alterações está disponível em vários idiomas:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CHANGELOG.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/CHANGELOG.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CHANGELOG.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CHANGELOG.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CHANGELOG.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CHANGELOG.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CHANGELOG.md)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/CHANGELOG.md)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/CHANGELOG.md)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/CHANGELOG.md)
-   [🇵🇹 Português](/pt/CHANGELOG) (atual)

---

**Acompanhe cada melhoria e fique por dentro do desenvolvimento do LazygoPHers Log!🚀**

---

_Este registro de alterações é atualizado automaticamente a cada lançamento. Para as informações mais recentes, consulte a página [GitHub Releases](https://github.com/lazygophers/log/releases)._
