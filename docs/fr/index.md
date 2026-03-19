---
pageType: home

hero:
    name: LazyGophers Log
    text: Une bibliothèque de journalisation Go performante et flexible
    tagline: Construite sur zap, offrant des fonctionnalités riches et une API simple
    actions:
        - theme: brand
          text: Démarrage rapide
          link: /fr/API
        - theme: alt
          text: Voir sur GitHub
          link: https://github.com/lazygophers/log

features:
    - title: "Haute performance"
      details: Construit sur zap avec réutilisation d'objets Entry via un pool, réduisant l'allocation mémoire
      icon: 🚀
    - title: "Niveaux de journalisation riches"
      details: Niveaux Trace, Debug, Info, Warn, Error, Fatal, Panic
      icon: 📊
    - title: "Configuration flexible"
      details: Personnalisez les niveaux de journalisation, les informations de l'appelant, les informations de trace, les préfixes, les suffixes et les cibles de sortie
      icon: ⚙️
    - title: "Rotation de fichiers"
      details: Support de la rotation horaire des fichiers journaux
      icon: 🔄
    - title: "Compatibilité Zap"
      details: Intégration transparente avec zap WriteSyncer
      icon: 🔌
    - title: "API simple"
      details: API claire similaire à la bibliothèque de journalisation standard, facile à utiliser
      icon: 🎯
---

## Démarrage rapide

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

## Documentation

-   [Référence API](API.md) - Documentation API complète
-   [Journal des modifications](/fr/CHANGELOG) - Historique des versions
-   [Guide de contribution](/fr/CONTRIBUTING) - Comment contribuer
-   [Politique de sécurité](/fr/SECURITY) - Guide de sécurité
-   [Code de conduite](/fr/CODE_OF_CONDUCT) - Normes communautaires

## Comparaison des performances

| Caractéristique     | lazygophers/log | zap    | logrus | Journal standard |
| ------------------- | --------------- | ------ | ------ | --------------- |
| Performance         | Élevée          | Élevée | Moyenne | Basse         |
| Simplicité de l'API | Élevée          | Moyenne | Élevée | Élevée         |
| Richesse des fonctionnalités | Moyenne      | Élevée | Élevée | Basse         |
| Flexibilité         | Moyenne          | Élevée | Élevée | Basse         |
| Courbe d'apprentissage | Basse          | Moyenne | Moyenne | Basse         |

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](/fr/LICENSE) pour plus de détails.
