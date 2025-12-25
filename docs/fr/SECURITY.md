---
titleSuffix: " | LazyGophers Log"
---

# 🔒 Politique de sécurité

## Versions supportées

| Version  | Supportée |
| -------- | --------- |
| >= 1.0.0 | ✅ Oui    |
| < 1.0.0  | ❌ Non    |

## Signalement d'une vulnérabilité

Nous prenons la sécurité de LazyGophers Log au sérieux et apprécions les rapports de vulnérabilité de la communauté.

### Comment signaler

Pour signaler une vulnérabilité de sécurité, veuillez envoyer un email à [security@lazygophers.com](mailto:security@lazygophers.com).

Veuillez inclure les informations suivantes dans votre rapport :

-   **Description** : Une description détaillée de la vulnérabilité
-   **Version affectée** : La ou les versions de LazyGophers Log concernées
-   **Environnement** : Les détails de l'environnement (version de Go, système d'exploitation, etc.)
-   **Étapes de reproduction** : Les étapes pour reproduire la vulnérabilité
-   **Impact potentiel** : L'impact potentiel de la vulnérabilité
-   **Preuve de concept** : Si possible, une preuve de concept ou un exemple d'exploitation
-   **Solution proposée** : Si possible, une solution ou un correctif suggéré

### Ce qu'attendre après le signalement

-   **Accusé de réception** : Vous recevrez un accusé de réception automatique
-   **Évaluation** : Nous évaluerons la vulnérabilité et déterminerons la priorité
-   **Mises à jour** : Nous vous informerons de l'état et des progrès
-   **Publication** : Après résolution, nous publierons un correctif et mettrons à jour les notes de version

## Chronologie de réponse

| Type de vulnérabilité      | Temps de réponse   |
| -------------------------- | ------------------ |
| Critique (sécurité active) | Dans les 48 heures |
| Haute                      | Dans les 7 jours   |
| Moyenne                    | Dans les 30 jours  |
| Basse                      | Dans les 90 jours  |

## Bonnes pratiques de sécurité

### Pour les utilisateurs

-   Utilisez toujours la dernière version stable de LazyGophers Log
-   Configurez les niveaux de journalisation de manière appropriée
-   Ne journalisez jamais de mots de passe ou de données sensibles
-   Validez toutes les entrées externes avant de les journaliser
-   Utilisez les niveaux de journalisation appropriés pour les différents environnements

### Pour les développeurs

-   Suivez les [principes de développement sécurisé](https://cheatsheetseries.owasp.org/cheatsheets/Go/Logging_and_Monitoring_Cheat_Sheet)
-   Effectuez des revues de code régulières
-   Utilisez l'analyse statique et les outils de fuzzing
-   Testez les modifications de performance pour les régressions de sécurité
-   Maintenez les dépendances à jour

## Divulgation coordonnée

Les divulgations coordonnées de vulnérabilités de sécurité doivent suivre ce processus :

1. **Préparation** : Préparez le correctif et les notes de version
2. **Coordination** : Coordonnez avec l'équipe de sécurité si nécessaire
3. **Publication** : Publiez le correctif et mettez à jour les notes de version
4. **Communication** : Informez les utilisateurs de manière claire et opportune

### Attribution

Les rapports de vulnérabilités de sécurité seront attribués à :

-   Le découvreur de la vulnérabilité
-   L'équipe de maintenance qui a résolu la vulnérabilité
-   Toute autre partie qui a contribué à la résolution

## Politique de divulgation

### Cas d'exception

Dans certains cas rares, nous pouvons décider de divulguer une vulnérabilité sans correctif immédiat :

-   Si la vulnérabilité est déjà publiquement connue
-   Si un correctif tiers est déjà disponible
-   Si la vulnérabilité est dans une version non supportée
-   Si la vulnérabilité a une gravité faible et un impact minimal

### Critères de divulgation

Toute divulgation sans correctif doit être approuvée par au moins deux mainteneurs du projet et doit inclure :

-   Justification détaillée de la décision
-   Plan de communication pour les utilisateurs
-   Évaluation des risques pour les utilisateurs

---

## 📞 Contact

Pour toute question concernant cette politique de sécurité, veuillez contacter :

-   **Email** : [security@lazygophers.com](mailto:security@lazygophers.com)
-   **GitHub Issues** : [Signaler un problème](https://github.com/lazygophers/log/issues/new/choose)

---

**La sécurité est une responsabilité partagée. Ensemble, nous pouvons créer un écosystème plus sécurisé.** 🔒
