---
titleSuffix: " | LazyGophers Log"
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Une bibliothèque de journalisation Go performante et flexible, construite sur zap, offrant des fonctionnalités riches et une API simple.

## 📖 Langues de documentation

-   [🇺🇸 English](../en/README.md)
-   [🇨🇳 简体中文](../zh-CN/README.md)
-   [🇹🇼 繁體中文](../zh-TW/README.md)
-   [🇫🇷 Français](README.md) (actuel)
-   [🇷🇺 Русский](../ru/README.md)
-   [🇪🇸 Español](../es/README.md)
-   [🇸🇦 العربية](../ar/README.md)

## 🚀 Documentation en ligne

Visitez notre [documentation GitHub Pages](https://lazygophers.github.io/log/) pour une meilleure expérience de lecture.

## ✨ Fonctionnalités

-   **🚀 Haute performance** : Construit sur zap avec réutilisation d'objets Entry via un pool, réduisant l'allocation mémoire
-   **📊 Niveaux de journalisation riches** : Niveaux Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configuration flexible** :
    -   Contrôle du niveau de journalisation
    -   Enregistrement des informations de l'appelant
    -   Informations de trace (y compris l'ID de goroutine)
    -   Préfixes et suffixes de journalisation personnalisés
    -   Cibles de sortie personnalisées (console, fichiers, etc.)
    -   Options de formatage de journalisation
-   **🔄 Rotation de fichiers** : Support de la rotation horaire des fichiers journaux
-   **🔌 Compatibilité Zap** : Intégration transparente avec zap WriteSyncer
-   **🎯 API simple** : API claire similaire à la bibliothèque de journalisation standard, facile à utiliser

## 🚀 Démarrage rapide

### Installation

```bash
go get github.com/lazygophers/log
```

### Utilisation de base

```go title="Démarrage rapide"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Utiliser le logger global par défaut
    log.Debug("Message de débogage")
    log.Info("Message d'information")
    log.Warn("Message d'avertissement")
    log.Error("Message d'erreur")

    // Utiliser la sortie formatée
    log.Infof("L'utilisateur %s s'est connecté avec succès", "admin")

    // Configuration personnalisée
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Ceci est un journal du logger personnalisé")
}
```

## 📚 Utilisation avancée

### Logger personnalisé avec sortie fichier

```go title="Configuration de sortie fichier"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Créer un logger avec sortie fichier
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Message de débogage avec informations de l'appelant")
    logger.Info("Message d'information avec informations de trace")
}
```

### Contrôle des niveaux de journalisation

```go title="Contrôle des niveaux de journalisation"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Seuls les messages warn et supérieur seront enregistrés
    logger.Debug("Ceci ne sera pas enregistré")  // Ignoré
    logger.Info("Ceci ne sera pas enregistré")   // Ignoré
    logger.Warn("Ceci sera enregistré")    // Enregistré
    logger.Error("Ceci sera enregistré")   // Enregistré
}
```

## 🎯 Cas d'usage

### Cas d'usage appropriés

-   **Services Web et API backend** : Suivi des requêtes, journaux structurés, surveillance des performances
-   **Architecture microservices** : Suivi distribué (TraceID), format de journalisation uniforme, haut débit
-   **Outils en ligne de commande** : Contrôle des niveaux, sortie concise, rapport d'erreurs
-   **Tâches de traitement par lots** : Rotation des fichiers, exécution de longue durée, optimisation des ressources

### Avantages particuliers

-   **Optimisation par pool d'objets** : Réutilisation des objets Entry et Buffer, réduction de la pression GC
-   **Écriture asynchrone** : Scénarios haut débit (10000+ messages/seconde) sans blocage
-   **Support TraceID** : Suivi des requêtes systèmes distribués, intégration avec OpenTelemetry
-   **Démarrage zéro configuration** : Prêt à l'emploi, configuration progressive

## 🔧 Options de configuration

:::note Options de configuration
Toutes les méthodes prennent en charge l'enchaînement, peuvent être combinées pour construire un Logger personnalisé.
:::

### Configuration du Logger

| Méthode                  | Description                             | Valeur par défaut |
| ------------------------ | --------------------------------------- | ---------------- |
| `SetLevel(level)`         | Définir le niveau de journalisation minimum | `DebugLevel`     |
| `EnableCaller(enable)`    | Activer/désactiver l'enregistrement des informations de l'appelant | `false`      |
| `EnableTrace(enable)`    | Activer/désactiver les informations de trace | `false`      |
| `SetCallerDepth(depth)`   | Définir la profondeur de l'appelant       | `2`          |
| `SetPrefixMsg(prefix)`    | Définir le préfixe du journal           | `""`         |
| `SetSuffixMsg(suffix)`    | Définir le suffixe du journal           | `""`         |
| `SetOutput(writers...)`   | Définir les cibles de sortie            | `os.Stdout`  |

### Niveaux de journalisation

| Niveau        | Description                         |
| ------------- | ----------------------------------- |
| `TraceLevel`  | Le plus détaillé, pour un suivi détaillé |
| `DebugLevel`  | Informations de débogage             |
| `InfoLevel`   | Messages d'information               |
| `WarnLevel`   | Conditions d'avertissement          |
| `ErrorLevel`  | Conditions d'erreur                  |
| `FatalLevel`  | Erreurs fatales (appelle os.Exit(1)) |
| `PanicLevel`  | Erreurs de panique (appelle panic()) |

## 🏗️ Architecture

### Composants principaux

-   **Logger** : Structure de journalisation principale avec niveau, sortie, formateur et profondeur d'appel configurables
-   **Entry** : Enregistrement de journal unique avec métadonnées complètes
-   **Level** : Définition et fonctions utilitaires des niveaux de journalisation
-   **Format** : Interface et implémentation du formatage des journaux

### Optimisations de performance

-   **Pool d'objets** : Réutilisation des objets Entry pour réduire l'allocation mémoire
-   **Enregistrement conditionnel** : Enregistrement des champs coûteux uniquement lorsque nécessaire
-   **Vérification rapide des niveaux** : Vérification des niveaux au niveau le plus externe
-   **Conception sans verrou** : La plupart des opérations n'ont pas besoin de verrous

## 📊 Comparaison des performances

:::info Comparaison des performances
Les données suivantes sont basées sur des benchmarks, les performances réelles peuvent varier selon l'environnement et la configuration.
:::

| Caractéristique     | lazygophers/log | zap    | logrus | Journal standard |
| ------------------- | --------------- | ------ | ------ | --------------- |
| Performance         | Élevée          | Élevée | Moyenne | Basse         |
| Simplicité de l'API | Élevée          | Moyenne | Élevée | Élevée         |
| Richesse des fonctionnalités | Moyenne      | Élevée | Élevée | Basse         |
| Flexibilité         | Moyenne          | Élevée | Élevée | Basse         |
| Courbe d'apprentissage | Basse          | Moyenne | Moyenne | Basse         |

## ❓ Questions fréquentes

### Comment choisir le niveau de journalisation approprié ?

-   **Environnement de développement** : Utiliser `DebugLevel` ou `TraceLevel` pour obtenir des informations détaillées
-   **Environnement de production** : Utiliser `InfoLevel` ou `WarnLevel` pour réduire la surcharge
-   **Tests de performance** : Utiliser `PanicLevel` pour désactiver tous les journaux

### Comment optimiser les performances en production ?

:::warning Attention
Dans les scénarios haut débit, il est recommandé de combiner l'écriture asynchrone et des niveaux de journalisation appropriés pour optimiser les performances.
:::

1. Utiliser `AsyncWriter` pour l'écriture asynchrone :

```go title="Configuration d'écriture asynchrone"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Ajuster les niveaux de journalisation pour éviter les journaux inutiles :

```go title="Optimisation des niveaux"
logger.SetLevel(log.InfoLevel)  // Sauter Debug et Trace
```

3. Utiliser des journaux conditionnels pour réduire la surcharge :

```go title="Journaux conditionnels"
if logger.Level >= log.DebugLevel {
    logger.Debug("Message de débogage détaillé")
}
```

### Quelle est la différence entre `Caller` et `EnableCaller` ?

-   **`EnableCaller(enable bool)`** : Contrôle si le Logger collecte les informations de l'appelant
    -   `EnableCaller(true)` active la collecte des informations de l'appelant
-   **`Caller(disable bool)`** : Contrôle si le formateur affiche les informations de l'appelant
    -   `Caller(true)` désactive l'affichage des informations de l'appelant

Il est recommandé d'utiliser `EnableCaller` pour un contrôle global.

### Comment implémenter un formateur personnalisé ?

Implémenter l'interface `Format` :

```go title="Formateur personnalisé"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Documentation connexe

-   [📚 Documentation API](API.md) - Référence API complète
-   [🤝 Guide de contribution](CONTRIBUTING.md) - Comment contribuer
-   [📋 Journal des modifications](CHANGELOG.md) - Historique des versions
-   [🔒 Politique de sécurité](SECURITY.md) - Guide de sécurité
-   [📜 Code de conduite](CODE_OF_CONDUCT.md) - Normes communautaires

## 🚀 Obtenir de l'aide

-   **GitHub Issues** : [Signaler un bug ou demander une fonctionnalité](https://github.com/lazygophers/log/issues)
-   **GoDoc** : [Documentation API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Exemples** : [Exemples d'utilisation](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

## 🤝 Contribuer

Les contributions sont les bienvenues ! Veuillez consulter notre [guide de contribution](CONTRIBUTING.md) pour plus de détails.

---

**lazygophers/log** vise à être la solution de journalisation de choix pour les développeurs Go qui valorisent les performances et la simplicité. Que vous construisiez de petits outils ou de grands systèmes distribués, cette bibliothèque offre un équilibre idéal entre fonctionnalité et facilité d'utilisation.