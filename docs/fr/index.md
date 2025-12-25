---
titleSuffix: " | LazyGophers Log"
---

# LazyGophers Log

Une bibliothèque de journalisation haute performance pour Go

API simple, performance excellente et configuration flexible

![LazyGophers Log Logo](/log/public/logo.svg)

[![CI Status](https://github.com/lazygophers/log/actions/workflows/ci.yml/badge.svg)](https://github.com/lazygophers/log/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/log.svg)](https://pkg.go.dev/github.com/lazygophers/log)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[Commencer](#quick-start) | [Référence API](/API)

## ✨ Fonctionnalités principales

### Haute performance

Construit sur zap, utilisant la mise en pool d'objets et l'enregistrement conditionnel de champs pour assurer une performance excellente

### Niveaux de journalisation riches

Supporte sept niveaux de journalisation : Trace, Debug, Info, Warn, Error, Fatal, Panic

### Configuration flexible

Supporte le contrôle du niveau de journalisation, l'enregistrement des informations de l'appelant, les informations de trace, les préfixes et suffixes personnalisés, etc.

### Rotation des fichiers

Fonction de rotation des fichiers de journal intégrée, supportant la rotation automatique horaire des fichiers de journal

### Compatibilité Zap

Intégration transparente avec zap WriteSyncer, supportant les cibles de sortie personnalisées

### API simple

API conçue similaire à la bibliothèque de journalisation standard, facile à utiliser et à migrer

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

    customLogger.Info("Ceci est un message du logger personnalisé")
}
```

## 📚 Navigation de la documentation

| Document                            | Description                              |
| ----------------------------------- | ---------------------------------------- |
| [Référence API](/API)               | Documentation API détaillée               |
| [Journal des modifications](/CHANGELOG)             | Voir tous les enregistrements de mise à jour de version          |
| [Guide de contribution](/CONTRIBUTING) | Comment contribuer du code au projet    |
| [Code de conduite](/CODE_OF_CONDUCT) | Code de conduite de la communauté                |
| [Politique de sécurité](/SECURITY)        | Processus de signalement des vulnérabilités de sécurité |

## 🌍 Documentation multilingue

-   [🇺🇸 English](/)
-   [🇨🇳 简体中文](/zh-CN/)
-   [🇹🇼 繁體中文](/zh-TW/)
-   [🇫🇷 Français](/fr/)

## 📄 Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](/LICENSE) pour plus de détails.

## 🤝 Contribution

Nous accueillons les contributions ! Veuillez consulter le [Guide de contribution](/CONTRIBUTING) pour plus de détails.

---

**LazyGophers Log** vise à être la solution de journalisation préférée des développeurs Go, se concentrant à la fois sur la performance et la facilité d'utilisation. Que vous construisiez de petits utilitaires ou de grands systèmes distribués, cette bibliothèque offre l'équilibre parfait entre fonctionnalités et facilité d'utilisation.

[⭐ Star sur GitHub](https://github.com/lazygophers/log)