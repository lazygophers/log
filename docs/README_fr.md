# 🚀 LazyGophers Log

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.5%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Une bibliothèque de journalisation hautes performances et riche en fonctionnalités pour les applications Go avec support multi-balises de construction, écriture asynchrone et options de personnalisation étendues.

## 📖 Langues de documentation

- [🇺🇸 English](../README.md)
- [🇨🇳 简体中文](README.zh-CN.md)
- [🇹🇼 繁體中文](README.zh-TW.md)
- [🇫🇷 Français](README.fr.md) (Actuel)
- [🇷🇺 Русский](README.ru.md)
- [🇪🇸 Español](README.es.md)
- [🇸🇦 العربية](README.ar.md)

## ✨ Fonctionnalités

- **🚀 Hautes performances**: Support du pooling d'objets et écriture asynchrone
- **🏗️ Support des balises de construction**: Différents comportements pour les modes debug, release et discard
- **🔄 Rotation des journaux**: Rotation automatique des fichiers de journaux par heure
- **🎨 Formatage riche**: Formats de journaux personnalisables avec support des couleurs
- **🔍 Traçage contextuel**: Suivi des ID de Goroutine et ID de trace
- **🔌 Intégration de frameworks**: Intégration native du logger Zap
- **⚙️ Hautement configurable**: Niveaux flexibles, sorties et formatage
- **🧪 Bien testé**: 93.5% de couverture de test à travers toutes les configurations de construction

## 🚀 Démarrage rapide

### Installation

```bash
go get github.com/lazygophers/log
```

### Utilisation de base

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Journalisation simple
    log.Info("Bonjour, Monde!")
    log.Debug("Ceci est un message de débogage")
    log.Warn("Ceci est un avertissement")
    log.Error("Ceci est une erreur")

    // Journalisation formatée
    log.Infof("L'utilisateur %s s'est connecté avec l'ID %d", "jean", 123)
    
    // Avec un logger personnalisé
    logger := log.New()
    logger.SetLevel(log.InfoLevel)
    logger.Info("Message du logger personnalisé")
}
```

### Utilisation avancée

```go
package main

import (
    "context"
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Créer un logger avec sortie vers fichier
    logger := log.New()
    
    // Définir la sortie vers un fichier avec rotation horaire
    writer := log.GetOutputWriterHourly("./logs/app.log")
    logger.SetOutput(writer)
    
    // Configurer le formatage
    logger.SetLevel(log.DebugLevel)
    logger.SetPrefixMsg("[APP] ")
    logger.Caller(true) // Activer les informations d'appelant
    
    // Journalisation contextuelle
    ctxLogger := logger.CloneToCtx()
    ctxLogger.Info(context.Background(), "Journalisation sensible au contexte")
    
    // Journalisation asynchrone pour les scénarios à haut débit
    asyncWriter := log.NewAsyncWriter(writer, 1000)
    logger.SetOutput(asyncWriter)
    defer asyncWriter.Close()
    
    logger.Info("Journalisation asynchrone hautes performances")
}
```

## 🏗️ Balises de construction

La bibliothèque prend en charge différents modes de construction via les balises de construction Go :

### Mode par défaut (Aucune balise)
```bash
go build
```
- Fonctionnalité complète de journalisation
- Messages de débogage activés
- Performances standard

### Mode débogage
```bash
go build -tags debug
```
- Informations de débogage améliorées
- Informations détaillées sur l'appelant
- Support du profilage de performance

### Mode release
```bash
go build -tags release
```
- Optimisé pour la production
- Messages de débogage désactivés
- Rotation automatique des fichiers de journaux

### Mode discard
```bash
go build -tags discard
```
- Performances maximales
- Tous les journaux sont supprimés
- Zéro surcharge de journalisation

### Modes combinés
```bash
go build -tags "debug,discard"    # Debug avec discard
go build -tags "release,discard"  # Release avec discard
```

## 📊 Niveaux de journaux

La bibliothèque prend en charge 7 niveaux de journaux (de la priorité la plus haute à la plus basse) :

| Niveau | Valeur | Description |
|--------|--------|-------------|
| `PanicLevel` | 0 | Journalise puis appelle panic |
| `FatalLevel` | 1 | Journalise puis appelle os.Exit(1) |
| `ErrorLevel` | 2 | Conditions d'erreur |
| `WarnLevel` | 3 | Conditions d'avertissement |
| `InfoLevel` | 4 | Messages informatifs |
| `DebugLevel` | 5 | Messages de niveau débogage |
| `TraceLevel` | 6 | Journalisation la plus détaillée |

## 🔌 Intégration de frameworks

### Intégration Zap

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/lazygophers/log"
)

// Créer un logger zap qui écrit dans notre système de journaux
logger := log.New()
hook := log.NewZapHook(logger)

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    hook,
    zapcore.InfoLevel,
)
zapLogger := zap.New(core)

zapLogger.Info("Message de Zap", zap.String("key", "value"))
```

## 🧪 Tests

La bibliothèque est livrée avec un support de test complet :

```bash
# Exécuter tous les tests
make test

# Exécuter les tests avec couverture pour toutes les balises de construction
make coverage-all

# Test rapide à travers toutes les balises de construction
make test-quick

# Générer les rapports de couverture HTML
make coverage-html
```

### Résultats de couverture par balise de construction

| Balise de construction | Couverture |
|------------------------|------------|
| Par défaut | 92.9% |
| Debug | 93.1% |
| Release | 93.5% |
| Discard | 93.1% |
| Debug+Discard | 93.1% |
| Release+Discard | 93.3% |

## ⚙️ Options de configuration

### Configuration du logger

```go
logger := log.New()

// Définir le niveau minimum de journalisation
logger.SetLevel(log.InfoLevel)

// Configurer la sortie
logger.SetOutput(os.Stdout) // Un seul writer
logger.SetOutput(writer1, writer2, writer3) // Plusieurs writers

// Personnaliser les messages
logger.SetPrefixMsg("[MonApp] ")
logger.SetSuffixMsg(" [FIN]")
logger.AppendPrefixMsg("Extra: ")

// Configurer le formatage
logger.ParsingAndEscaping(false) // Désactiver les séquences d'échappement
logger.Caller(true) // Activer les informations d'appelant
logger.SetCallerDepth(4) // Ajuster la profondeur de pile d'appelant
```

## 📁 Rotation des journaux

Rotation automatique des journaux avec intervalles configurables :

```go
// Rotation horaire
writer := log.GetOutputWriterHourly("./logs/app.log")

// La bibliothèque créera des fichiers comme :
// - app-2024010115.log (2024-01-01 15:00)
// - app-2024010116.log (2024-01-01 16:00)
// - app-2024010117.log (2024-01-01 17:00)
```

## 🔍 Contexte et traçage

Support intégré pour la journalisation sensible au contexte et le traçage distribué :

```go
// Définir l'ID de trace pour la goroutine actuelle
log.SetTrace("trace-123-456")

// Obtenir l'ID de trace
traceID := log.GetTrace()

// Journalisation sensible au contexte
ctx := context.Background()
ctxLogger := log.CloneToCtx()
ctxLogger.Info(ctx, "Requête traitée", "user_id", 123)

// Suivi automatique des ID de goroutine
log.Info("Ce journal inclut automatiquement l'ID de goroutine")
```

## 📈 Performances

La bibliothèque est conçue pour les applications hautes performances :

- **Pooling d'objets**: Réutilise les objets d'entrée de journal pour réduire la pression GC
- **Écriture asynchrone**: Écritures de journaux non bloquantes pour les scénarios à haut débit
- **Filtrage de niveau**: Le filtrage précoce évite les opérations coûteuses
- **Optimisation des balises de construction**: Optimisation au moment de la compilation pour différents environnements

### Benchmarks

```bash
# Exécuter les benchmarks de performance
make benchmark

# Benchmark des différents modes de construction
make benchmark-debug
make benchmark-release  
make benchmark-discard
```

## 🤝 Contribuer

Nous accueillons les contributions ! Veuillez consulter notre [Guide de contribution](CONTRIBUTING.md) pour plus de détails.

### Configuration de développement

1. **Fork et Clone**
   ```bash
   git clone https://github.com/your-username/log.git
   cd log
   ```

2. **Installer les dépendances**
   ```bash
   go mod tidy
   ```

3. **Exécuter les tests**
   ```bash
   make test-all
   ```

4. **Soumettre une Pull Request**
   - Suivez notre [Modèle PR](../.github/pull_request_template.md)
   - Assurez-vous que les tests passent
   - Mettez à jour la documentation si nécessaire

## 📋 Exigences

- **Go**: 1.21 ou supérieur
- **Dépendances**: 
  - `go.uber.org/zap` (pour l'intégration Zap)
  - `github.com/petermattis/goid` (pour l'ID de goroutine)
  - `github.com/lestrrat-go/file-rotatelogs` (pour la rotation des journaux)
  - `github.com/google/uuid` (pour les ID de trace)

## 📄 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](../LICENSE) pour plus de détails.

## 🙏 Remerciements

- [Zap](https://github.com/uber-go/zap) pour l'inspiration et le support d'intégration
- [Logrus](https://github.com/sirupsen/logrus) pour les modèles de conception de niveaux
- La communauté Go pour les commentaires continus et les améliorations

## 📞 Support

- 📖 [Documentation](../docs/)
- 🐛 [Suivi des problèmes](https://github.com/lazygophers/log/issues)
- 💬 [Discussions](https://github.com/lazygophers/log/discussions)
- 📧 Email: support@lazygophers.com

---

**Fait avec ❤️ par l'équipe LazyGophers**