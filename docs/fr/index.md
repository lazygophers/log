---
pageType: home

hero:
  name: LazyGophers Log
  text: Une bibliothèque de journalisation haute performance pour Go
  tagline: Basée sur zap avec des fonctionnalités riches et une API simple
  actions:
    - theme: brand
      text: Démarrer
      link: /fr/API
    - theme: alt
      text: Voir sur GitHub
      link: https://github.com/lazygophers/log

features:
  - title: 'Haute Performance'
    details: Basée sur zap avec mise en pool d'objets et enregistrement conditionnel des champs
    icon: 🚀
  - title: 'Niveaux de Journalisation Riches'
    details: Prend en charge les niveaux Trace, Debug, Info, Warn, Error, Fatal et Panic
    icon: 📊
  - title: 'Configuration Flexible'
    details: Personnalisez les niveaux, les informations de l'appelant, le traçage, les préfixes, suffixes et cibles de sortie
    icon: ⚙️
  - title: 'Rotation des Fichiers'
    details: Support intégré de la rotation horaire des fichiers de journalisation
    icon: 🔄
  - title: 'Compatibilité Zap'
    details: Intégration transparente avec zap WriteSyncer
    icon: 🔌
  - title: 'API Simple'
    details: API claire similaire à la bibliothèque de journalisation standard, facile à utiliser et à intégrer
    icon: 🎯
---

## Démarrage Rapide

### Installation

```bash
go get github.com/lazygophers/log
```

### Utilisation de Base

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
    log.Infof("Utilisateur %s connecté avec succès", "admin")

    // Configuration personnalisée
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Ceci est un journal du logger personnalisé")
}
```

## Documentation

- [Référence API](API.md) - Documentation API complète
- [Journal des Modifications](CHANGELOG.md) - Historique des versions
- [Guide de Contribution](CONTRIBUTING.md) - Comment contribuer
- [Politique de Sécurité](SECURITY.md) - Directives de sécurité
- [Code de Conduite](CODE_OF_CONDUCT.md) - Directives communautaires

## Comparaison des Performances

| Fonctionnalité      | lazygophers/log | zap    | logrus | journal standard |
| ------------------- | --------------- | ------ | ------ | ----------------- |
| Performance         | Élevée          | Élevée | Moyenne| Faible            |
| Simplicité API      | Élevée          | Moyenne| Élevée | Élevée            |
| Richesse Fonctions | Moyenne         | Élevée | Élevée | Faible            |
| Flexibilité        | Moyenne         | Élevée | Élevée | Faible            |
| Courbe d'Apprentissage| Faible       | Moyenne| Moyenne| Faible            |

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour les détails.
