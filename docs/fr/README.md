---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Une bibliothèque de journalisation Go à haute performance et flexible, construite sur zap, offrant des fonctionnalités riches et une API simple.

## 📖 Langues de la documentation

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 Chinois simplifié](https://lazygophers.github.io/log/zh-CN/)
-   [🇹🇼 Chinois traditionnel](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](README.md) (actuel)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/)
-   [🇵🇹 Português](https://lazygophers.github.io/log/pt/)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/)
-   [🇵🇱 Polski](https://lazygophers.github.io/log/pl/)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/)
-   [🇹🇷 Türkçe](https://lazygophers.github.io/log/tr/)

## ✨ Fonctionnalités

-   **🚀 Haute performance**：Construit sur zap avec pool d'objets et enregistrement conditionnel de champs
-   **📊 Niveaux de journalisation riches**：Niveaux Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configuration flexible**：
    -   Contrôle du niveau de journalisation
    -   Enregistrement des informations de l'appelant
    -   Informations de traçage (y compris l'ID de goroutine)
    -   Préfixes et suffixes personnalisés
    -   Cibles de sortie personnalisées (console, fichiers, etc.)
    -   Options de formatage du journal
-   **🔄 Rotation des fichiers**：Support de la rotation des fichiers de journal chaque heure
-   **🔌 Compatibilité Zap**：Intégration transparente avec zap WriteSyncer
-   **🎯 API simple**：API clair similaire à la bibliothèque de journalisation standard, facile à utiliser

## 🚀 Démarrage rapide

### Installation

:::tip Installation
```bash
go get github.com/lazygophers/log
```
:::

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

    // Utiliser une sortie formatée
    log.Infof("L'utilisateur %s s'est connecté avec succès", "admin")

    // Configuration personnalisée
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MonApp]")

    customLogger.Info("Ceci est un message du logger personnalisé")
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
    logger.Info("Message d'information avec informations de traçage")
}
```

### Contrôle du niveau de journalisation

```go title="Contrôle du niveau de journalisation"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Seuls les niveaux warn et supérieurs seront enregistrés
    logger.Debug("Ceci ne sera pas enregistré")  // Ignoré
    logger.Info("Ceci ne sera pas enregistré")   // Ignoré
    logger.Warn("Ceci sera enregistré")    // Enregistré
    logger.Error("Ceci sera enregistré")   // Enregistré
}
```

## 🎯 Scénarios d'utilisation

### Scénarios applicables

-   **Services web et backend API**：Traçage des requêtes, journalisation structurée, surveillance des performances
-   **Architecture microservices**：Traçage distribué (TraceID), format de journal unifié, haute volumétrie
-   **Outils en ligne de commande**：Contrôle des niveaux, sortie claire, rapports d'erreurs
-   **Tâches par lots**：Rotation des fichiers, exécution longue, optimisation des ressources

### Avantages particuliers

-   **Optimisation avec pool d'objets**：Réutilisation des objets Entry et Buffer, réduction de la pression du GC
-   **Écriture asynchrone**：Haute volumétrie (10000+ enregistrements/seconde) sans blocage
-   **Support TraceID**：Traçage des requêtes dans les systèmes distribués, intégration avec OpenTelemetry
-   **Démarrage sans configuration**：Prêt à l'emploi, configuration progressive

## 🔧 Options de configuration

:::note Options de configuration
Toutes les méthodes ci-dessous prennent en charge l'appel en chaîne et peuvent être combinées pour créer un Logger personnalisé.
:::

### Configuration du Logger

| Méthode                  | Description                | Valeur par défaut       |
| --------------------- | ------------------- | ----------- |
| `SetLevel(level)`       | Définir le niveau de journalisation minimum     | `DebugLevel` |
| `EnableCaller(enable)`  | Activer/désactiver les informations de l'appelant  | `false`      |
| `EnableTrace(enable)`   | Activer/désactiver les informations de traçage    | `false`      |
| `SetCallerDepth(depth)` | Définir la profondeur de l'appelant       | `2`          |
| `SetPrefixMsg(prefix)`  | Définir le préfixe du journal         | `""`         |
| `SetSuffixMsg(suffix)`  | Définir le suffixe du journal         | `""`         |
| `SetOutput(writers...)` | Définir les cibles de sortie         | `os.Stdout`  |

### Niveaux de journalisation

| Niveau        | Description                        |
| ----------- | --------------------------- |
| `TraceLevel` | Le plus détaillé, pour le traçage détaillé        |
| `DebugLevel` | Informations de débogage                    |
| `InfoLevel`  | Informations générales                    |
| `WarnLevel`  | Messages d'avertissement                    |
| `ErrorLevel` | Messages d'erreur                    |
| `FatalLevel` | Erreurs fatales (appelle os.Exit(1))    |
| `PanicLevel` | Erreurs de panique (appelle panic())    |

## 🏗️ Architecture

### Composants principaux

-   **Logger**：Structure principale de journalisation avec options configurables
-   **Entry**：Enregistrement individuel avec support étendu des métadonnées
-   **Level**：Définitions des niveaux de journalisation et utilitaires
-   **Format**：Interface de formatage du journal et implémentations

### Optimisation des performances

-   **Pool d'objets**：Réutilisation des objets Entry pour réduire l'allocation de mémoire
-   **Enregistrement conditionnel**：Enregistre uniquement les champs coûteux lorsque nécessaire
-   **Vérification rapide des niveaux**：Vérification du niveau de journalisation à la couche externe
-   **Design sans verrouillage**：La plupart des opérations ne nécessitent pas de verrous

## 📊 Comparaison des performances

:::info Comparaison des performances
Les données ci-dessous sont basées sur les tests de référence; les performances réelles peuvent varier selon l'environnement et la configuration.
:::

| Caractéristique          | lazygophers/log | zap    | logrus | journal standard |
| ------------- | --------------- | ------ | ------ | -------- |
| Performance      | Haute              | Haute     | Moyenne     | Faible       |
| Simplicité de l'API    | Haute              | Moyenne     | Haute     | Haute       |
| Richesse des fonctionnalités    | Moyenne          | Haute     | Haute     | Faible       |
| Flexibilité      | Moyenne          | Haute     | Haute     | Faible       |
| Courbe d'apprentissage      | Faible              | Moyenne     | Moyenne     | Faible       |

## ❓ Questions fréquemment posées

### Comment choisir le bon niveau de journalisation ?

-   **Environnement de développement**：Utiliser `DebugLevel` ou `TraceLevel` pour obtenir des informations détaillées
-   **Environnement de production**：Utiliser `InfoLevel` ou `WarnLevel` pour réduire les coûts
-   **Tests de performance**：Utiliser `PanicLevel` pour désactiver tous les journaux

### Comment optimiser les performances en production ?

:::note Attention
Dans les scénarios à haute volumétrie, il est recommandé de combiner l'écriture asynchrone avec des niveaux de journalisation raisonnables pour optimiser les performances.
:::

1. Utiliser `AsyncWriter` pour l'écriture asynchrone：

```go title="Configuration de l'écriture asynchrone"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Ajuster le niveau de journalisation pour éviter les journaux inutiles：

```go title="Optimisation du niveau"
logger.SetLevel(log.InfoLevel)  // Passer Debug et Trace
```

3. Utiliser des journaux conditionnels pour réduire les coûts：

```go title="Journaux conditionnels"
if logger.Level >= log.DebugLevel {
    logger.Debug("Informations de débogage détaillées")
}
```

### Quelle est la différence entre `Caller` et `EnableCaller` ?

-   **`EnableCaller(enable bool)`**：Contrôle si le Logger collecte les informations de l'appelant
    -   `EnableCaller(true)` active la collecte des informations de l'appelant
-   **`Caller(disable bool)`**：Contrôle si le Formatter affiche les informations de l'appelant
    -   `Caller(true)` désactive l'affichage des informations de l'appelant

Il est recommandé d'utiliser `EnableCaller` pour un contrôle global.

### Comment implémenter un formateur personnalisé ?

Implémenter l'interface `Format`：

```go title="Formateur personnalisé"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Documentation associée

-   [📚 Documentation API](API.md) - Référence API complète
-   [🤝 Guide de contribution](/fr/CONTRIBUTING) - Comment contribuer
-   [📋 Journal des modifications](/fr/CHANGELOG) - Historique des versions
-   [🔒 Politique de sécurité](/fr/SECURITY) - Guide de sécurité
-   [📜 Code de conduite](/fr/CODE_OF_CONDUCT) - Normes de la communauté

## 🚀 Obtenir de l'aide

-   **GitHub Issues**：[Rapporter un bug ou demander une fonctionnalité](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[Documentation API](https://pkg.go.dev/github.com/lazygophers/log)
-   [✓ Exemples](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](/fr/LICENSE) pour plus de détails.

## 🤝 Contribution

Nous accueillons favorablement les contributions ! Veuillez consulter notre [Guide de contribution](/fr/CONTRIBUTING) pour plus d'informations.

---

**lazygophers/log** est conçu pour être la solution de journalisation privilégiée pour les développeurs Go qui valorisent à la fois la performance et la simplicité. Que vous construisiez une petite utilité ou un système distribué à grande échelle, cette bibliothèque fournit un excellent équilibre entre fonctionnalité et facilité d'utilisation.
