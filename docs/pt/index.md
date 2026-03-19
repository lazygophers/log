---
pageType: home

hero:
    name: LazyGophers Log
    text: Biblioteca de registro Go de alto desempenho e flexível
    tagline: Construída sobre zap, fornecendo recursos ricos e API simples
    actions:
        - theme: brand
          text: Início Rápido
          link: /API
        - theme: alt
          text: Ver no GitHub
          link: https://github.com/lazygophers/log

features:
    - title: "Alto Desempenho"
      details: Construída sobre zap com reutilização de pool de objetos e registro condicional de campos
      icon: 🚀
    - title: "Níveis de Rico"
      details: Suporta níveis Trace, Debug, Info, Warn, Error, Fatal, Panic
      icon: 📊
    - title: "Configuração Flexível"
      details: Personalize níveis, informações do chamador, rastreamento, prefixos, sufixos e destinos de saída
      icon: ⚙️
    - title: "Rotação de Arquivos"
      details: Suporte embutido para rotação de arquivos de log por hora
      icon: 🔄
    - title: "Compatibilidade Zap"
      details: Integração perfeita com zap WriteSyncer
      icon: 🔌
    - title: "API Simples"
      details: API clara semelhante à biblioteca de registro padrão, fácil de usar e integrar
      icon: 🎯
---

## Início Rápido

### Instalação

```bash
go get github.com/lazygophers/log
```

### Uso Básico

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Usar logger global padrão
    log.Debug("Informação de depuração")
    log.Info("Informação geral")
    log.Warn("Informação de aviso")
    log.Error("Informação de erro")

    // Usar saída formatada
    log.Infof("Usuário %s fez login com sucesso", "admin")

    // Configuração personalizada
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Este é um log do logger personalizado")
}
```

## Documentação

-   [Referência de API](API.md) - Documentação completa da API
-   [Log de Alterações](/pt/CHANGELOG) - Histórico de versões
-   [Guia de Contribuição](/pt/CONTRIBUTING) - Como contribuir
-   [Política de Segurança](/pt/SECURITY) - Guia de segurança
-   [Código de Conduta](/pt/CODE_OF_CONDUCT) - Diretrizes da comunidade

## Comparação de Desempenho

| Recurso       | lazygophers/log | zap | logrus | Log padrão |
| ---------- | --------------- | --- | ------ | -------- |
| Desempenho       | Alto              | Alto  | Médio     | Baixo       |
| Simplicidade da API    | Alto              | Médio  | Alto     | Alto       |
| Riqueza de recursos    | Médio              | Alto  | Alto     | Baixo       |
| Flexibilidade      | Médio              | Alto  | Alto     | Baixo       |
| Curva de aprendizado      | Baixo              | Médio  | Médio     | Baixo       |

## Licença

Este projeto é licenciado sob a Licença MIT - consulte o arquivo [LICENSE](/pt/LICENSE) para detalhes.
