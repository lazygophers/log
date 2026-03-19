---
titleSuffix: " | LazyGophers Log"
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Une bibliotheque de journalisation Go performante et flexible, construite sur zap, offrant des fonctionnalites riches et une API simple.

## Langues de documentation

-   [English](https://lazygophers.github.io/log/en/)
-   [简体中文](https://lazygophers.github.io/log/)
-   [繁體中文](https://lazygophers.github.io/log/zh-TW/)
-   [Français](README.md) (actuel)
-   [Русский](https://lazygophers.github.io/log/ru/)
-   [Español](https://lazygophers.github.io/log/es/)
-   [العربية](https://lazygophers.github.io/log/ar/)

## Documentation en ligne

Visitez notre [documentation GitHub Pages](https://lazygophers.github.io/log/) pour une meilleure expérience de lecture.

## Fonctionnalités

-   **Haute performance** : Construit sur zap avec réutilisation d'objets Entry via un pool
-   **Niveaux de journalisation riches** : Niveaux Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **Configuration flexible** : Contrôle du niveau, infos de l'appelant, trace, préfixes, suffixes
-   **Rotation de fichiers** : Support de la rotation horaire des fichiers journaux
-   **Compatibilité Zap** : Intégration transparente avec zap WriteSyncer
-   **API simple** : API claire similaire à la bibliothèque standard

## Installation

```bash
go get github.com/lazygophers/log
```

## Utilisation de base

```go
package main

import "github.com/lazygophers/log"

func main() {
    log.Debug("Message de débogage")
    log.Info("Message d'information")
    log.Warn("Message d'avertissement")
    log.Error("Message d'erreur")

    log.Infof("L'utilisateur %s s'est connecté", "admin")

    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Ceci est un journal personnalisé")
}
```

## Documentation

-   [Documentation API](API.md)
-   [Journal des modifications](/fr/CHANGELOG)
-   [Guide de contribution](/fr/CONTRIBUTING)
-   [Politique de sécurité](/fr/SECURITY)
-   [Code de conduite](/fr/CODE_OF_CONDUCT)

## Licence

Ce projet est sous licence MIT.
