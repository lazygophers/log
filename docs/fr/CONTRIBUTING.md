---
titleSuffix: " | LazyGophers Log"
---

# 🤝 Contribution à LazyGophers Log

Nous accueillons vos contributions ! Nous voulons rendre la contribution à LazyGophers Log aussi simple et transparente que possible, que ce soit :

-   🐛 Signalement de bugs
-   💬 Discussion de l'état actuel du code
-   ✨ Demande de fonctionnalités
-   🔧 Proposition de correctifs
-   🚀 Implémentation de nouvelles fonctionnalités

## 📋 Table des matières

-   [Code de conduite](#-code-de-conduite)
-   [Processus de développement](#-processus-de-développement)
-   [Démarrage rapide](#-démarrage-rapide)
-   [Processus de Pull Request](#-processus-de-pull-request)
-   [Normes de codage](#-normes-de-codage)
-   [Directives de tests](#-directives-de-tests)
-   [Balises de construction](#-balises-de-construction)
-   [Documentation](#-documentation)
-   [Directives d'issues](#-directives-dissues)
-   [Considérations de performance](#-considérations-de-performance)
-   [Directives de sécurité](#-directives-de-sécurité)
-   [Communauté](#-communauté)

## 📜 Code de conduite

Ce projet et tous les participants sont régis par notre [Code de conduite](CODE_OF_CONDUCT.md). En participant, vous acceptez de vous conformer à ce code.

## 🔄 Processus de développement

Nous utilisons GitHub pour héberger le code, suivre les issues et les demandes de fonctionnalités, et accepter les pull requests.

### Workflow

1. **Fork** le dépôt
2. **Clone** votre fork localement
3. **Créer** une branche de fonctionnalité à partir de `master`
4. **Apporter** vos modifications
5. **Tester** soigneusement sous toutes les balises de construction
6. **Soumettre** une pull request

## 🚀 Démarrage rapide

### Prérequis

-   **Go 1.21+** - [Installer Go](https://golang.org/doc/install)

### Installation

```bash
# Cloner le dépôt
git clone https://github.com/lazygophers/log.git
cd log

# Installer les dépendances
go mod download
```

### Exécution des tests

```bash
# Exécuter tous les tests
go test ./...

# Exécuter les tests avec une balise de construction spécifique
go test -tags=debug ./...

# Exécuter les tests de performance
go test -bench=. -benchmem ./...
```

## � Processus de Pull Request

### Avant de soumettre

1. Vérifiez que vos tests passent sous toutes les balises de construction
2. Exécutez `go fmt` sur vos modifications
3. Assurez-vous que votre code est propre et bien documenté
4. Mettez à jour la documentation si nécessaire
5. Ajoutez des tests pour les nouvelles fonctionnalités

### Vérifications des balises de construction

LazyGophers Log prend en charge les balises de construction pour optimiser les performances dans différents environnements. Assurez-vous de tester avec toutes les balises :

```bash
# Tester avec la balise par défaut
go test ./...

# Tester avec la balise de débogage
go test -tags=debug ./...

# Tester avec la balise de publication
go test -tags=release ./...

# Tester avec la balise d'abandon
go test -tags=discard ./...
```

### Format du titre de la PR

Utilisez un titre clair et descriptif pour votre Pull Request :

-   `feat: Ajouter la fonctionnalité X`
-   `fix: Corriger le bug Y`
-   `docs: Mettre à jour la documentation`
-   `perf: Améliorer les performances`
-   `refactor: Refactoriser le code`

## 📝 Normes de codage

### Formatage du code

Exécutez `go fmt` avant de soumettre :

```bash
go fmt ./...
```

### Linting

Nous utilisons golangci-lint pour assurer la qualité du code :

```bash
# Installer golangci-lint si nécessaire
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Exécuter le linting
golangci-lint run
```

## 🧪 Directives de tests

### Couverture des tests

Visez une couverture de tests élevée. Les nouvelles fonctionnalités doivent inclure des tests.

```bash
# Exécuter les tests avec couverture
go test -coverprofile=coverage.out ./...

# Vérifier la couverture
go tool cover -func=coverage.out
```

### Tests de performance

Les modifications de performance doivent inclure des benchmarks :

```go
func BenchmarkLogger(b *testing.B) {
    logger := log.New()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        logger.Info("Message de test")
    }
}
```

## 🔧 Balises de construction

LazyGophers Log utilise des balises de construction pour optimiser les performances dans différents environnements :

| Balise    | Description                                                      | Utilisation           |
| --------- | ---------------------------------------------------------------- | --------------------- |
| (défaut)  | Fonctionnalité complète avec messages de débogage                | Développement général |
| `debug`   | Informations de débogage améliorées et détails de l'appelant     | Débogage approfondi   |
| `release` | Optimisé pour la production avec messages de débogage désactivés | Production            |
| `discard` | Performance maximale avec opérations de journalisation no-op     | Tests de performance  |

### Test avec des balises de construction

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

## 📚 Documentation

### Mises à jour requises

-   Mettre à jour les commentaires de code pour les nouvelles fonctions
-   Mettre à jour la documentation API si nécessaire
-   Mettre à jour les exemples si nécessaire

### Normes de documentation

-   Utiliser des descriptions claires et concises
-   Inclure des exemples d'utilisation
-   Documenter les paramètres et les valeurs de retour

## 🐛 Directives d'issues

### Signalement de bugs

Lorsque vous signalez un bug, incluez :

-   Version de Go utilisée
-   Version de lazygophers/log
-   Balises de construction utilisées
-   Description détaillée du problème
-   Exemple de code minimal pour reproduire
-   Sortie attendue vs sortie réelle

### Demandes de fonctionnalités

Pour les demandes de fonctionnalités, incluez :

-   Description claire de la fonctionnalité souhaitée
-   Cas d'utilisation proposés
-   Avantages de cette fonctionnalité
-   Solutions alternatives considérées

## ⚡ Considérations de performance

### Optimisations à éviter

-   Évitez les allocations inutiles dans les chemins chauds
-   Utilisez les chaînes de caractères au lieu de la concaténation excessive
-   Réutilisez les objets Entry via le pool interne
-   Évitez les conversions de type inutiles

### Bonnes pratiques

-   Utilisez les niveaux de journalisation appropriés
-   Évitez la journalisation dans les boucles serrées
-   Utilisez le journalisation conditionnelle pour les champs coûteux
-   Testez les modifications de performance avec des benchmarks

## 🔒 Directives de sécurité

### Signalement des vulnérabilités

Pour les vulnérabilités de sécurité, veuillez consulter notre [Politique de sécurité](SECURITY.md) pour :

-   Versions supportées
-   Procédures de signalement
-   Chronologie de réponse
-   Bonnes pratiques de sécurité

### Bonnes pratiques

-   Validez toutes les entrées externes
-   Évitez l'injection de données dans les messages de journalisation
-   Utilisez les niveaux de journalisation appropriés
-   Ne journalisez jamais de mots de passe ou de données sensibles

## 👥 Communauté

### Canaux de communication

-   📖 [Documentation](/)
-   � [Suivi des problèmes](https://github.com/lazygophers/log/issues)
-   💬 [Discussions](https://github.com/lazygophers/log/discussions)

### Reconnaissance

Nous reconnaissons et apprécions toutes les contributions. Les contributeurs seront crédités dans les notes de version.

## 📄 Licence

En contribuant, vous acceptez que vos contributions seront sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

---

Merci de contribuer à LazyGophers Log ! �
