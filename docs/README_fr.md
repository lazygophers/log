# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Une bibliothèque de journalisation Go performante et flexible, construite sur zap, offrant des fonctionnalités riches et une API simple.

## 📖 Langues de documentation

-   [🇺🇸 English](README.md)
-   [🇨🇳 简体中文](README_zh-CN.md)
-   [🇹🇼 繁體中文](README_zh-TW.md)
-   [🇫🇷 Français](README_fr.md)
-   [🇷🇺 Русский](README_ru.md)
-   [🇪🇸 Español](README_es.md)
-   [🇸🇦 العربية](README_ar.md)

## ✨ Fonctionnalités

-   **🚀 Haute performance** : Construite sur zap avec réutilisation d'objets Entry via un pool, réduisant l'allocation mémoire
-   **📊 Niveaux de journalisation riches** : Niveaux Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configuration flexible** :
    -   Contrôle du niveau de journalisation
    -   Enregistrement des informations d'appelant
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

```go
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

```go
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

    logger.Debug("Message de débogage avec informations d'appelant")
    logger.Info("Message d'information avec informations de trace")
}
```

### Contrôle du niveau de journalisation

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

## 🔧 Options de configuration

### Configuration du Logger

| Méthode                 | Description                 | Valeur par défaut |
| ----------------------- | --------------------------- | ---------------- |
| `SetLevel(level)`       | Définir le niveau minimum de journalisation | `DebugLevel` |
| `EnableCaller(enable)`  | Activer/désactiver les informations d'appelant | `false` |
| `EnableTrace(enable)`   | Activer/désactiver les informations de trace | `false` |
| `SetCallerDepth(depth)` | Définir la profondeur de l'appelant | `2` |
| `SetPrefixMsg(prefix)`  | Définir le préfixe du journal | `""` |
| `SetSuffixMsg(suffix)`  | Définir le suffixe du journal | `""` |
| `SetOutput(writers...)` | Définir les cibles de sortie | `os.Stdout` |

### Niveaux de journalisation

| Niveau        | Description                        |
| ------------- | ---------------------------------- |
| `TraceLevel`  | Le plus verbeux, pour le suivi détaillé |
| `DebugLevel`  | Informations de débogage           |
| `InfoLevel`   | Informations générales             |
| `WarnLevel`   | Messages d'avertissement           |
| `ErrorLevel`  | Messages d'erreur                  |
| `FatalLevel`  | Erreurs fatales (appelle os.Exit(1)) |
| `PanicLevel`  | Erreurs de panique (appelle panic()) |

## 🏗️ Architecture

### Composants principaux

-   **Logger** : Structure de journalisation principale avec niveaux, sorties, formateurs et profondeur d'appelant configurables
-   **Entry** : Enregistrement de journal unique avec support complet de métadonnées
-   **Level** : Définitions de niveaux de journalisation et fonctions utilitaires
-   **Format** : Interface et implémentations de formatage de journalisation

### Optimisations de performance

-   **Pool d'objets** : Réutilisation des objets Entry pour réduire l'allocation mémoire
-   **Enregistrement conditionnel** : Enregistrement des champs coûteux uniquement lorsque nécessaire
-   **Vérification rapide du niveau** : Vérification du niveau de journalisation à la couche la plus externe
-   **Conception sans verrou** : La plupart des opérations ne nécessitent pas de verrou

## 📊 Comparaison des performances

| Caractéristique      | lazygophers/log | zap    | logrus | journalisation standard |
| -------------------- | --------------- | ------ | ------ | ----------------------- |
| Performance          | Haute           | Haute  | Moyenne | Basse                   |
| Simplicité de l'API  | Haute           | Moyenne | Haute  | Haute                   |
| Richesse de fonctionnalités | Moyenne | Haute  | Haute  | Basse                   |
| Flexibilité          | Moyenne         | Haute  | Haute  | Basse                   |
| Courbe d'apprentissage | Basse          | Moyenne | Moyenne | Basse                   |

## 🔗 Documentation associée

-   [📚 Documentation API](API.md) - Référence API complète
-   [🤝 Guide de contribution](CONTRIBUTING.md) - Comment contribuer
-   [📋 Journal des modifications](../CHANGELOG.md) - Historique des versions
-   [🔒 Politique de sécurité](SECURITY.md) - Directives de sécurité
-   [📜 Code de conduite](CODE_OF_CONDUCT.md) - Directives de communauté

## 🚀 Obtenir de l'aide

-   **GitHub Issues** : [Signaler des bugs ou demander des fonctionnalités](https://github.com/lazygophers/log/issues)
-   **GoDoc** : [Documentation API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Exemples** : [Exemples d'utilisation](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](../LICENSE) pour plus de détails.

## 🤝 Contribution

Nous accueillons les contributions ! Veuillez consulter notre [Guide de contribution](CONTRIBUTING.md) pour plus de détails.

---

**lazygophers/log** est conçu pour être la solution de journalisation de choix pour les développeurs Go qui valorisent à la fois la performance et la simplicité. Que vous construisiez un petit utilitaire ou un système distribué à grande échelle, cette bibliothèque offre le bon équilibre entre fonctionnalités et facilité d'utilisation.