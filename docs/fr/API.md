---
titleSuffix: " | LazyGophers Log"
---

# 📚 Documentation API

## Vue d'ensemble

LazyGophers Log fournit une API de journalisation complète qui supporte plusieurs niveaux de journalisation, un formatage personnalisé, l'écriture asynchrone et l'optimisation des balises de construction. Ce document couvre toutes les API publiques, les options de configuration et les modèles d'utilisation.

## Table des matières

-   [Types de base](#types-de-base)
-   [API Logger](#api-logger)
-   [Fonctions globales](#fonctions-globales)
-   [Niveaux de journalisation](#niveaux-de-journalisation)
-   [Formateurs](#formateurs)
-   [Rédacteurs de sortie](#redacteurs-de-sortie)
-   [Journalisation contextuelle](#journalisation-contextuelle)
-   [Balises de construction](#balises-de-construction)
-   [Optimisation des performances](#optimisation-des-performances)
-   [Exemples](#exemples)

## Types de base

### Logger

La structure de journalisation principale qui fournit toutes les fonctionnalités de journalisation.

```go
type Logger struct {
    // Contient des champs privés pour les opérations thread-safe
}
```

#### Constructeur

```go
func New() *Logger
```

Crée une nouvelle instance de journalisation avec la configuration par défaut :

-   Niveau : `DebugLevel`
-   Sortie : `os.Stdout`
-   Formateur : Formateur de texte par défaut
-   Suivi de l'appelant : Désactivé

**Exemple :**

```go
logger := log.New()
logger.Info("Nouveau journalisation créé")
```

### Entry

Représente une seule entrée de journal avec toutes les métadonnées associées.

```go
type Entry struct {
    Time       time.Time     // Horodatage lors de la création de l'entrée
    Level      Level         // Niveau de journalisation
    Message    string        // Message de journalisation
    Pid        int          // ID de processus
    Gid        uint64       // ID de goroutine
    TraceID    string       // ID de trace pour le traçage distribué
    CallerName string       // Nom de la fonction de l'appelant
    CallerFile string       // Chemin du fichier de l'appelant
    CallerLine int          // Numéro de ligne de l'appelant
}
```

## API Logger

### Méthodes de configuration

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Définit le niveau minimum de journalisation. Les messages en dessous de ce niveau seront ignorés.

**Paramètres :**

-   `level` : Le niveau minimum de journalisation à traiter

**Retour :**

-   `*Logger` : Renvoie lui-même pour supporter le chaînage de méthodes

**Exemple :**

```go
logger.SetLevel(log.InfoLevel)
logger.Debug("Ceci ne sera pas affiché")  // Ignoré
logger.Info("Ceci sera affiché")    // Traité
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Active ou désactive l'enregistrement des informations de l'appelant.

**Paramètres :**

-   `enable` : `true` pour activer, `false` pour désactiver

**Retour :**

-   `*Logger` : Renvoie lui-même pour supporter le chaînage de méthodes

**Exemple :**

```go
logger.EnableCaller(true)
logger.Info("Ceci inclura les informations de l'appelant")
```

#### EnableTrace

```go
func (l *Logger) EnableTrace(enable bool) *Logger
```

Active ou désactive l'enregistrement des informations de trace (y compris l'ID de goroutine).

**Paramètres :**

-   `enable` : `true` pour activer, `false` pour désactiver

**Retour :**

-   `*Logger` : Renvoie lui-même pour supporter le chaînage de méthodes

**Exemple :**

```go
logger.EnableTrace(true)
logger.Info("Ceci inclura les informations de trace")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Définit la profondeur de la pile d'appels pour l'enregistrement de l'appelant.

**Paramètres :**

-   `depth` : La profondeur de la pile d'appels (défaut : 2)

**Retour :**

-   `*Logger` : Renvoie lui-même pour supporter le chaînage de méthodes

**Exemple :**

```go
logger.SetCallerDepth(3)
logger.Info("Ceci utilisera une profondeur de 3")
```

#### SetPrefixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
```

Définit le préfixe pour tous les messages de journalisation.

**Paramètres :**

-   `prefix` : Le préfixe à ajouter avant chaque message

**Retour :**

-   `*Logger` : Renvoie lui-même pour supporter le chaînage de méthodes

**Exemple :**

```go
logger.SetPrefixMsg("[MyApp] ")
logger.Info("Ceci aura un préfixe")
```

#### SetSuffixMsg

```go
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Définit le suffixe pour tous les messages de journalisation.

**Paramètres :**

-   `suffix` : Le suffixe à ajouter après chaque message

**Retour :**

-   `*Logger` : Renvoie lui-même pour supporter le chaînage de méthodes

**Exemple :**

```go
logger.SetSuffixMsg(" [END]")
logger.Info("Ceci aura un suffixe")
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Définit les cibles de sortie pour la journalisation.

**Paramètres :**

-   `writers` : Une ou plusieurs cibles de sortie (par exemple, `os.Stdout`, fichiers)

**Retour :**

-   `*Logger` : Renvoie lui-même pour supporter le chaînage de méthodes

**Exemple :**

```go
// Sortie vers stdout et un fichier
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

## Fonctions globales

Les fonctions globales fournissent un accès rapide à une instance de journalisation par défaut.

### Niveaux de journalisation

```go
const (
    PanicLevel Level = iota
    FatalLevel
    ErrorLevel
    WarnLevel
    InfoLevel
    DebugLevel
    TraceLevel
)
```

| Niveau       | Description                      | Utilisation                                               |
| ------------ | -------------------------------- | --------------------------------------------------------- |
| `PanicLevel` | Le plus élevé, appelle `panic()` | Erreurs critiques qui arrêtent l'application              |
| `FatalLevel` | Élevé, appelle `os.Exit(1)`      | Erreurs fatales qui nécessitent une terminaison immédiate |
| `ErrorLevel` | Élevé                            | Erreurs générales                                         |
| `WarnLevel`  | Moyen                            | Messages d'avertissement                                  |
| `InfoLevel`  | Normal                           | Messages d'information généraux                           |
| `DebugLevel` | Bas                              | Informations de débogage                                  |
| `TraceLevel` | Le plus bas                      | Informations de trace détaillées                          |

### Méthodes de journalisation

```go
func Trace(args ...interface{})
func Debug(args ...interface{})
func Info(args ...interface{})
func Warn(args ...interface{})
func Error(args ...interface{})
func Fatal(args ...interface{})
func Panic(args ...interface{})
```

**Exemple :**

```go
log.Trace("Message de trace")
log.Debug("Message de débogage")
log.Info("Message d'information")
log.Warn("Message d'avertissement")
log.Error("Message d'erreur")
log.Fatal("Message fatal")  // Appelle os.Exit(1)
log.Panic("Message de panique")  // Appelle panic()
```

### Méthodes formatées

```go
func Tracef(format string, args ...interface{})
func Debugf(format string, args ...interface{})
func Infof(format string, args ...interface{})
func Warnf(format string, args ...interface{})
func Errorf(format string, args ...interface{})
func Fatalf(format string, args ...interface{})
func Panicf(format string, args ...interface{})
```

**Exemple :**

```go
log.Infof("L'utilisateur %s s'est connecté", "admin")
log.Errorf("Échec de la connexion : %v", err)
```

## Balises de construction

LazyGophers Log prend en charge les balises de construction pour optimiser les performances dans différents environnements.

### Balises disponibles

| Balise    | Description                                                      | Utilisation           |
| --------- | ---------------------------------------------------------------- | --------------------- |
| (défaut)  | Fonctionnalité complète avec messages de débogage                | Développement général |
| `debug`   | Informations de débogage améliorées et détails de l'appelant     | Débogage approfondi   |
| `release` | Optimisé pour la production avec messages de débogage désactivés | Production            |
| `discard` | Performance maximale avec opérations de journalisation no-op     | Tests de performance  |

**Utilisation :**

```bash
# Développement (par défaut)
go build

# Débogage approfondi
go build -tags=debug

# Production
go build -tags=release

# Tests de performance
go build -tags=discard
```

## Optimisation des performances

### Pool d'objets

LazyGophers Log utilise `sync.Pool` pour réutiliser les objets `Entry` et les tampons, réduisant l'allocation mémoire et la pression sur le ramasse-miettes.

### Enregistrement conditionnel

Les champs coûteux (comme les informations de l'appelant et de trace) ne sont enregistrés que si le niveau de journalisation le permet, évitant les calculs inutiles.

### Vérification rapide du niveau

Le niveau de journalisation est vérifié à la couche la plus externe, permettant un retour rapide sans allocation mémoire pour les messages qui seront ignorés.

### Conception sans verrou

La plupart des opérations de journalisation ne nécessitent pas de verrou, offrant une meilleure performance en environnement concurrent.

## Exemples

### Journalisation simple

```go
package main

import "github.com/lazygophers/log"

func main() {
    log.Info("Application démarrée")
    log.Warn("Ceci est un avertissement")
    log.Error("Ceci est une erreur")
}
```

### Journalisation avec sortie fichier

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Créer un journalisation avec sortie fichier
    logger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Message de débogage avec informations de l'appelant")
    logger.Info("Message d'information avec informations de trace")
}
```

### Journalisation conditionnelle

```go
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Seuls les messages warn et supérieurs seront journalisés
    logger.Debug("Ceci ne sera pas journalisé")  // Ignoré
    logger.Info("Ceci ne sera pas journalisé")   // Ignoré
    logger.Warn("Ceci sera journalisé")    // Journalisé
    logger.Error("Ceci sera journalisé")   // Journalisé
}
```

### Journalisation avec préfixe personnalisé

```go
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().
        SetLevel(log.InfoLevel).
        SetPrefixMsg("[MyApp] ").
        SetSuffixMsg(" [DONE]")

    logger.Info("Ceci aura un préfixe et un suffixe")
}
```

## Intégration Zap

LazyGophers Log peut être utilisé comme `zap.WriteSyncer` pour une intégration transparente avec les applications existantes.

```go
package main

import (
    "go.uber.org/zap"
    "github.com/lazygophers/log"
)

func main() {
    // Créer un logger zap
    zapLogger, _ := zap.NewProduction()

    // Utiliser lazygophers/log comme WriteSyncer
    logger := log.New().SetOutput(zapLogger)

    logger.Info("Ceci sera écrit via zap")
}
```

## Support

-   📖 [Documentation](/)
-   🐛 [Suivi des problèmes](https://github.com/lazygophers/log/issues)
-   💬 [Discussions](https://github.com/lazygophers/log/discussions)
-   📧 Email: support@lazygophers.com

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

---

**Pour plus d'informations, consultez la [documentation complète](https://lazygophers.github.io/log/).**
