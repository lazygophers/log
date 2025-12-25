---
titleSuffix: " | LazyGophers Log"
---

# 📋 Journal des modifications

Toutes les modifications notables de ce projet seront documentées dans ce fichier.

Le format est basé sur [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), et ce projet adhère au [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Non publié]

### Ajouté

-   Documentation multilingue complète (7 langues)
-   Modèles d'issues GitHub (Rapport de bug, Demande de fonctionnalité, Questions)
-   Modèle de Pull Request avec vérifications de compatibilité des balises de construction
-   Directives de contribution en plusieurs langues
-   Code de conduite avec lignes directrices d'application
-   Politique de sécurité avec processus de signalement des vulnérabilités
-   Documentation API complète avec exemples
-   Structure de projet professionnelle et modèles

### Modifié

-   README amélioré avec documentation complète des fonctionnalités
-   Couverture de tests améliorée pour toutes les configurations de balises de construction
-   Structure de projet mise à jour pour une meilleure maintenabilité

### Documentation

-   Ajout du support multilingue pour toute la documentation principale
-   Création d'une référence API complète
-   Établissement des directives de workflow de contribution
-   Mise en œuvre des procédures de signalement de sécurité

## [1.0.0] - 2024-01-01

### Ajouté

-   Fonctionnalité de journalisation de base avec plusieurs niveaux (Trace, Debug, Info, Warn, Error, Fatal, Panic)
-   Implémentation de journalisation thread-safe avec mise en pool d'objets
-   Support des balises de construction (par défaut, débogage, publication, abandon)
-   Interface de formateur personnalisé avec formateur de texte par défaut
-   Support de sortie multi-rédacteur
-   Capacités d'écriture asynchrone pour les scénarios à haut débit
-   Rotation automatique horaire des fichiers de journalisation
-   Journalisation contextuelle avec suivi de l'ID de goroutine et de l'ID de trace
-   Informations de l'appelant avec profondeur de pile configurable
-   Fonctions de commodité au niveau du package global
-   Support d'intégration du journaliseur Zap

### Performance

-   Mise en pool d'objets avec `sync.Pool` pour les objets d'entrée et les tampons
-   Vérification rapide du niveau pour éviter les opérations coûteuses
-   Rédacteur asynchrone pour les écritures de journalisation non bloquantes
-   Optimisations des balises de construction pour différents environnements

### Balises de construction

-   **Par défaut** : Fonctionnalité complète avec messages de débogage
-   **Débogage** : Informations de débogage améliorées et détails de l'appelant
-   **Publication** : Optimisé pour la production avec messages de débogage désactivés
-   **Abandon** : Performance maximale avec opérations de journalisation no-op

### Fonctionnalités principales

-   **Logger** : Structure de journalisation principale avec niveau, sortie et formateur configurables
-   **Entry** : Structure d'enregistrement de journal avec métadonnées complètes
-   **Levels** : Sept niveaux de journalisation de Panic (le plus élevé) à Trace (le plus bas)
-   **Formatters** : Système de formatage enfichable
-   **Writers** : Support de rotation de fichiers et d'écriture asynchrone
-   **Context** : Support de l'ID de goroutine et du traçage distribué

### Points forts de l'API

-   API de configuration fluide avec chaînage de méthodes
-   Méthodes de journalisation simples et formatées (`.Info()` et `.Infof()`)
-   Clonage de journaliseur pour configurations isolées
-   Journalisation contextuelle avec `CloneToCtx()`
-   Personnalisation des messages de préfixe et de suffixe
-   Basculement des informations de l'appelant

### Tests

-   Suite de tests complète avec une couverture de 93.5%
-   Support de tests multi-balises de construction
-   Workflows de tests automatisés
-   Tests de performance

## [0.9.0] - 2023-12-15

### Ajouté

-   Structure de projet initiale
-   Fonctionnalité de journalisation de base
-   Filtrage basé sur le niveau
-   Support de sortie fichier

### Modifié

-   Performance améliorée avec mise en pool d'objets
-   Gestion des erreurs améliorée

## [0.8.0] - 2023-12-01

### Ajouté

-   Support multi-rédacteur
-   Interface de formateur personnalisé
-   Capacités d'écriture asynchrone

### Corrigé

-   Fuites de mémoire dans les scénarios à haut débit
-   Conditions de course dans l'accès concurrent

## [0.7.0] - 2023-11-15

### Ajouté

-   Support des balises de construction pour la compilation conditionnelle
-   Journalisation de niveau Trace et débogage
-   Suivi des informations de l'appelant

### Modifié

-   Modèles d'allocation mémoire optimisés
-   Sécurité des threads améliorée

## [0.6.0] - 2023-11-01

### Ajouté

-   Fonctionnalité de rotation des fichiers de journalisation
-   Journalisation contextuelle
-   Suivi de l'ID de goroutine

### Obsolète

-   Anciennes méthodes de configuration (seront supprimées dans v1.0.0)

## [0.5.0] - 2023-10-15

### Ajouté

-   Formateur JSON
-   Plusieurs destinations de sortie
-   Tests de performance

### Modifié

-   Refactorisation du moteur de journalisation principal
-   Cohérence de l'API améliorée

### Supprimé

-   Anciennes méthodes de journalisation

## [0.4.0] - 2023-10-01

### Ajouté

-   Journalisation de niveau Fatal et Panic
-   Fonctions globales du package
-   Validation de la configuration

### Corrigé

-   Problèmes de synchronisation de la sortie
-   Optimisation de l'utilisation mémoire

## [0.3.0] - 2023-09-15

### Ajouté

-   Niveaux de journalisation personnalisés
-   Interface de formateur
-   Opérations thread-safe

### Modifié

-   Conception de l'API simplifiée
-   Documentation améliorée

## [0.2.0] - 2023-09-01

### Ajouté

-   Support de sortie fichier
-   Filtrage basé sur le niveau
-   Options de formatage de base

### Corrigé

-   Goulots d'étranglement de performance
-   Fuites de mémoire

## [0.1.0] - 2023-08-15

### Ajouté

-   Publication initiale
-   Journalisation de console de base
-   Support de niveau simple (Info, Warn, Error)
-   Structure de journalisation principale

## Résumé de l'historique des versions

| Version | Date de publication | Fonctionnalités principales                                                                             |
| ------- | ------------------- | ------------------------------------------------------------------------------------------------------- |
| 1.0.0   | 2024-01-01          | Système de journalisation complet, balises de construction, écriture asynchrone, documentation complète |
| 0.9.0   | 2023-12-15          | Améliorations de performance, mise en pool d'objets                                                     |
| 0.8.0   | 2023-12-01          | Multi-rédacteur, écriture asynchrone, formateurs personnalisés                                          |
| 0.7.0   | 2023-11-15          | Balises de construction, niveaux Trace/débogage, informations de l'appelant                             |
| 0.6.0   | 2023-11-01          | Rotation des fichiers, journalisation contextuelle, suivi de goroutine                                  |
| 0.5.0   | 2023-10-15          | Formateur JSON, sorties multiples, tests de performance                                                 |
| 0.4.0   | 2023-10-01          | Niveaux Fatal/Panic, fonctions globales                                                                 |
| 0.3.0   | 2023-09-15          | Niveaux personnalisés, interface de formateur                                                           |
| 0.2.0   | 2023-09-01          | Sortie fichier, filtrage par niveau                                                                     |
| 0.1.0   | 2023-08-15          | Publication initiale, journalisation de console de base                                                 |

## Guides de migration

### Migration de v0.9.x vers v1.0.0

#### Modifications avec rupture

-   Aucune - v1.0.0 est rétrocompatible avec v0.9.x

#### Nouvelles fonctionnalités disponibles

-   Support des balises de construction amélioré
-   Documentation complète
-   Modèles de projet professionnels
-   Procédures de signalement de sécurité

#### Mises à jour recommandées

```go
// Ancienne méthode (toujours supportée)
logger := log.New()
logger.SetLevel(log.InfoLevel)

// Nouvelle méthode recommandée avec chaînage de méthodes
logger := log.New().
    SetLevel(log.InfoLevel).
    Caller(true).
    SetPrefixMsg("[MyApp] ")
```

### Migration de v0.8.x vers v0.9.x

#### Modifications avec rupture

-   Suppression des méthodes de configuration obsolètes
-   Modification de la gestion des tampons internes

#### Étapes de migration

1. Mettre à jour les chemins d'importation si nécessaire
2. Remplacer les méthodes obsolètes :

    ```go
    // Ancienne (obsolète)
    logger.SetOutputFile("app.log")

    // Nouvelle
    file, _ := os.Create("app.log")
    logger.SetOutput(file)
    ```

### Migration de v0.5.x et versions antérieures

#### Modifications majeures

-   Conception complète de l'API pour une meilleure cohérence
-   Performance améliorée avec mise en pool d'objets
-   Nouveau système de balises de construction

#### Migration requise

-   Mettre à jour tous les appels de journalisation vers la nouvelle API
-   Réviser et mettre à jour les implémentations de formateur
-   Tester avec les nouvelles configurations de balises de construction

## Jalons de développement

### 🎯 Feuille de route v1.1.0 (Planifié)

-   [ ] Journalisation structurée avec paires clé-valeur
-   [ ] Échantillonnage de journalisation pour les scénarios à grand volume
-   [ ] Système de plugins pour les sorties personnalisées
-   [ ] Métriques de performance améliorées
-   [ ] Intégrations de journalisation cloud

### 🎯 Feuille de route v1.2.0 (Futur)

-   [ ] Support des fichiers de configuration (YAML/JSON/TOML)
-   [ ] Agrégation et filtrage de journalisation
-   [ ] Streaming de journalisation en temps réel
-   [ ] Fonctionnalités de sécurité améliorées
-   [ ] Intégration du tableau de bord de performance

## Contribution

Nous accueillons les contributions ! Veuillez consulter notre [Guide de contribution](CONTRIBUTING.md) pour plus de détails sur :

-   Signalement de bugs et demandes de fonctionnalités
-   Processus de soumission de code
-   Configuration du développement
-   Exigences de tests
-   Normes de documentation

## Sécurité

Pour les vulnérabilités de sécurité, veuillez consulter notre [Politique de sécurité](SECURITY.md) pour :

-   Versions supportées
-   Procédures de signalement
-   Chronologie de réponse
-   Bonnes pratiques de sécurité

## Support

-   📖 [Documentation](/)
-   🐛 [Suivi des problèmes](https://github.com/lazygophers/log/issues)
-   💬 [Discussions](https://github.com/lazygophers/log/discussions)
-   📧 Email: support@lazygophers.com

## Licence

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

---

## 🌍 Documentation multilingue

Ce journal des modifications est disponible en plusieurs langues :

-   [🇺🇸 English](CHANGELOG.md)
-   [🇨🇳 简体中文](../zh-CN/CHANGELOG.md)
-   [🇹🇼 繁體中文](../zh-TW/CHANGELOG.md)
-   [🇫🇷 Français](CHANGELOG.md) (Courant)
-   [🇷🇺 Русский](../README_ru.md)
-   [🇪🇸 Español](../README_es.md)
-   [🇸🇦 العربية](../README_ar.md)

---

**Suivez chaque amélioration et restez informé de l'évolution de LazygoPHers Log ! 🚀**

---

_Ce journal des modifications est mis à jour automatiquement avec chaque publication. Pour les informations les plus récentes, consultez la page [GitHub Releases](https://github.com/lazygophers/log/releases)._
