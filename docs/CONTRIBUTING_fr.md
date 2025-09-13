# 🤝 Contribuer à LazyGophers Log

Nous adorons vos contributions ! Nous voulons rendre la contribution à LazyGophers Log aussi facile et transparente que possible, qu'il s'agisse de :

- 🐛 Signaler un bug
- 💬 Discuter de l'état actuel du code
- ✨ Soumettre une demande de fonctionnalité
- 🔧 Proposer une correction
- 🚀 Implémenter une nouvelle fonctionnalité

## 📋 Table des matières

- [Code de conduite](#-code-de-conduite)
- [Processus de développement](#-processus-de-développement)
- [Démarrage rapide](#-démarrage-rapide)
- [Processus de Pull Request](#-processus-de-pull-request)
- [Standards de codage](#-standards-de-codage)
- [Directives de test](#-directives-de-test)
- [Exigences des tags de build](#️-exigences-des-tags-de-build)
- [Documentation](#-documentation)
- [Directives pour les issues](#-directives-pour-les-issues)
- [Considérations de performance](#-considérations-de-performance)
- [Directives de sécurité](#-directives-de-sécurité)
- [Communauté](#-communauté)

## 📜 Code de conduite

Ce projet et tous ceux qui y participent sont régis par notre [Code de conduite](CODE_OF_CONDUCT_fr.md). En participant, vous vous engagez à respecter ce code.

## 🔄 Processus de développement

Nous utilisons GitHub pour héberger le code, suivre les issues et les demandes de fonctionnalités, ainsi qu'accepter les pull requests.

### Flux de travail

1. **Fork** le dépôt
2. **Clone** votre fork localement
3. **Créez** une branche de fonctionnalité depuis `master`
4. **Effectuez** vos modifications
5. **Testez** minutieusement sur tous les tags de build
6. **Soumettez** une pull request

## 🚀 Démarrage rapide

### Prérequis

- **Go 1.21+** - [Installer Go](https://golang.org/doc/install)
- **Git** - [Installer Git](https://git-scm.com/book/fr/v2/Démarrage-rapide-Installation-de-Git)
- **Make** (optionnel mais recommandé)

### Configuration de l'environnement de développement local

```bash
# 1. Forkez le dépôt sur GitHub
# 2. Clonez votre fork
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. Ajoutez le remote upstream
git remote add upstream https://github.com/lazygophers/log.git

# 4. Installez les dépendances
go mod tidy

# 5. Vérifiez l'installation
make test-quick
```

### Configuration de l'environnement

```bash
# Configurez votre environnement Go (si ce n'est pas déjà fait)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Optionnel : Installez des outils utiles
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 Processus de Pull Request

### Avant de soumettre

1. **Recherchez** les PR existantes pour éviter les doublons
2. **Testez** vos modifications sur toutes les configurations de build
3. **Documentez** toute modification majeure
4. **Mettez à jour** la documentation pertinente
5. **Ajoutez** des tests pour les nouvelles fonctionnalités

### Liste de contrôle PR

- [ ] **Qualité du code**
  - [ ] Le code suit les directives de style du projet
  - [ ] Aucun nouvel avertissement de linting
  - [ ] Gestion d'erreur appropriée
  - [ ] Algorithmes et structures de données efficaces

- [ ] **Tests**
  - [ ] Tous les tests existants passent : `make test`
  - [ ] Nouveaux tests ajoutés pour les nouvelles fonctionnalités
  - [ ] Couverture de test maintenue ou améliorée
  - [ ] Tous les tags de build testés : `make test-all`

- [ ] **Documentation**
  - [ ] Le code est correctement commenté
  - [ ] Documentation API mise à jour (si nécessaire)
  - [ ] README mis à jour (si nécessaire)
  - [ ] Documentation multilingue mise à jour (si orientée utilisateur)

- [ ] **Compatibilité de build**
  - [ ] Mode par défaut : `go build`
  - [ ] Mode debug : `go build -tags debug`
  - [ ] Mode release : `go build -tags release`
  - [ ] Mode discard : `go build -tags discard`
  - [ ] Modes combinés testés

### Template de PR

Veuillez utiliser notre [template de PR](.github/pull_request_template.md) lors de la soumission des pull requests.

## 📏 Standards de codage

### Guide de style Go

Nous suivons le guide de style Go standard avec quelques ajouts :

```go
// ✅ Bon
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ Mauvais
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### Conventions de nommage

- **Packages** : Courts, minuscules, un seul mot quand possible
- **Fonctions** : CamelCase, descriptives
- **Variables** : camelCase pour locales, CamelCase pour exportées
- **Constantes** : CamelCase pour exportées, camelCase pour non-exportées
- **Interfaces** : Se terminent généralement par "er" (ex. `Writer`, `Formatter`)

### Organisation du code

```
project/
├── docs/           # Documentation en plusieurs langues
├── .github/        # Templates GitHub et workflows
├── logger.go       # Implémentation principale du logger
├── entry.go        # Structure d'entrée de log
├── level.go        # Niveaux de log
├── formatter.go    # Formatage des logs
├── output.go       # Gestion de sortie
└── *_test.go      # Tests co-localisés avec le source
```

### Gestion d'erreurs

```go
// ✅ Préféré : Retourner les erreurs, ne pas paniquer
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ Éviter : Paniquer dans le code de bibliothèque
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // Ne pas faire cela
    }
    return &Logger{...}
}
```

## 🧪 Directives de test

### Structure de test

```go
func TestLogger_Info(t *testing.T) {
    tests := []struct {
        name     string
        level    Level
        message  string
        expected bool
    }{
        {"info level allows info", InfoLevel, "test", true},
        {"warn level blocks info", WarnLevel, "test", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Implémentation du test
        })
    }
}
```

### Exigences de couverture

- **Minimum** : 90% de couverture pour le nouveau code
- **Objectif** : 95%+ de couverture globale
- **Tous les tags de build** doivent maintenir la couverture
- Utilisez `make coverage-all` pour vérifier

### Commandes de test

```bash
# Test rapide sur tous les tags de build
make test-quick

# Suite de tests complète avec couverture
make test-all

# Rapports de couverture
make coverage-html

# Benchmarks
make benchmark
```

## 🏗️ Exigences des tags de build

Toutes les modifications doivent être compatibles avec notre système de tags de build :

### Tags de build supportés

- **Par défaut** (`go build`) : Fonctionnalité complète
- **Debug** (`go build -tags debug`) : Débogage amélioré
- **Release** (`go build -tags release`) : Optimisé pour la production
- **Discard** (`go build -tags discard`) : Performance maximale

### Test des tags de build

```bash
# Tester chaque configuration de build
make test-default
make test-debug  
make test-release
make test-discard

# Tester les combinaisons
make test-debug-discard
make test-release-discard

# Tout en un
make test-all
```

### Directives pour les tags de build

```go
//go:build debug
// +build debug

package log

// Implémentations spécifiques au debug
```

## 📚 Documentation

### Documentation du code

- **Toutes les fonctions exportées** doivent avoir des commentaires clairs
- **Les algorithmes complexes** ont besoin d'explications
- **Exemples** pour l'utilisation non-triviale
- **Notes de sécurité des threads** le cas échéant

```go
// SetLevel définit le niveau de log minimum.
// Les logs en dessous de ce niveau seront ignorés.
// Cette méthode est thread-safe.
//
// Exemple :
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // Ne sortira pas
//   logger.Info("visible")   // Sortira
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### Mises à jour du README

Lors de l'ajout de fonctionnalités, mettez à jour :
- Le README.md principal
- Tous les README spécifiques à la langue dans `docs/`
- Les exemples de code
- Les listes de fonctionnalités

## 🐛 Directives pour les issues

### Rapports de bugs

Utilisez le [template de rapport de bug](.github/ISSUE_TEMPLATE/bug_report.md) et incluez :

- **Description claire** du problème
- **Étapes pour reproduire**
- **Comportement attendu vs réel**
- **Détails de l'environnement** (OS, version Go, tags de build)
- **Échantillon de code minimal**

### Demandes de fonctionnalités

Utilisez le [template de demande de fonctionnalité](.github/ISSUE_TEMPLATE/feature_request.md) et incluez :

- **Motivation claire** pour la fonctionnalité
- **Design d'API** proposé
- **Considérations d'implémentation**
- **Analyse des changements majeurs**

### Questions

Utilisez le [template de question](.github/ISSUE_TEMPLATE/question.md) pour :

- Questions d'utilisation
- Aide de configuration
- Meilleures pratiques
- Conseils d'intégration

## 🚀 Considérations de performance

### Benchmarking

Toujours benchmarker les modifications sensibles aux performances :

```bash
# Exécuter les benchmarks
go test -bench=. -benchmem

# Comparer avant/après
go test -bench=. -benchmem > before.txt
# Faire des modifications
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### Directives de performance

- **Minimiser les allocations** dans les chemins critiques
- **Utiliser les pools d'objets** pour les objets fréquemment créés
- **Retour anticipé** pour les niveaux de log désactivés
- **Éviter la réflexion** dans le code critique
- **Profiler avant d'optimiser**

### Gestion mémoire

```go
// ✅ Bon : Utiliser les pools d'objets
var entryPool = sync.Pool{
    New: func() interface{} {
        return &Entry{}
    },
}

func getEntry() *Entry {
    return entryPool.Get().(*Entry)
}

func putEntry(e *Entry) {
    e.Reset()
    entryPool.Put(e)
}
```

## 🔒 Directives de sécurité

### Données sensibles

- **Ne jamais logger** les mots de passe, tokens ou données sensibles
- **Assainir** l'entrée utilisateur dans les messages de log
- **Éviter** de logger les corps de requête/réponse complets
- **Utiliser** le logging structuré pour un meilleur contrôle

```go
// ✅ Bon
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ Mauvais
logger.Infof("User login: %+v", userRequest) // Peut contenir un mot de passe
```

### Dépendances

- Garder les dépendances **à jour**
- **Réviser** soigneusement les nouvelles dépendances
- **Minimiser** les dépendances externes
- **Utiliser** `go mod verify` pour vérifier l'intégrité

## 👥 Communauté

### Obtenir de l'aide

- 📖 [Documentation](../README_fr.md)
- 💬 [Discussions GitHub](https://github.com/lazygophers/log/discussions)
- 🐛 [Tracker d'issues](https://github.com/lazygophers/log/issues)
- 📧 Email : support@lazygophers.com

### Directives de communication

- Être **respectueux** et inclusif
- **Rechercher** avant de poser des questions
- **Fournir du contexte** quand on demande de l'aide
- **Aider les autres** quand on peut
- **Suivre** le [Code de conduite](CODE_OF_CONDUCT_fr.md)

## 🎯 Reconnaissance

Les contributeurs sont reconnus de plusieurs façons :

- Section **contributeurs du README**
- Mentions dans les **notes de release**
- Graphiques des **contributeurs GitHub**
- Posts d'**appréciation de la communauté**

## 📝 Licence

En contribuant, vous acceptez que vos contributions soient sous licence MIT.

---

## 🌍 Documentation multilingue

Ce document est disponible en plusieurs langues :

- [🇺🇸 English](CONTRIBUTING.md)
- [🇨🇳 简体中文](CONTRIBUTING_zh-CN.md)
- [🇹🇼 繁體中文](CONTRIBUTING_zh-TW.md)
- [🇫🇷 Français](CONTRIBUTING_fr.md) (Actuel)
- [🇷🇺 Русский](CONTRIBUTING_ru.md)
- [🇪🇸 Español](CONTRIBUTING_es.md)
- [🇸🇦 العربية](CONTRIBUTING_ar.md)

---

**Merci de contribuer à LazyGophers Log ! 🚀**