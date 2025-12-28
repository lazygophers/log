---
titleSuffix: " | LazyGophers Log"
pageType: home

hero:
  name: LazyGophers Log
  text: Une bibliothèque de journalisation haute performance pour Go
  tagline: API simple, performance excellente et configuration flexible
  image:
    src: /log/public/logo.svg
    alt: LazyGophers Log Logo
  actions:
    - theme: brand
      text: Commencer
      link: /fr/README
    - theme: alt
      text: Référence API
      link: /fr/API
    - theme: alt
      text: Voir sur GitHub
      link: https://github.com/lazygophers/log

features:
  - title: Haute performance
    details: Construit sur zap, utilisant la mise en pool d'objets et l'enregistrement conditionnel de champs pour assurer une performance excellente
    icon: ⚡
  - title: Niveaux de journalisation riches
    details: Supporte sept niveaux de journalisation : Trace, Debug, Info, Warn, Error, Fatal, Panic
    icon: 📊
  - title: Configuration flexible
    details: Supporte le contrôle du niveau de journalisation, l'enregistrement des informations de l'appelant, les informations de trace, les préfixes et suffixes personnalisés, etc.
    icon: ⚙️
  - title: Rotation des fichiers
    details: Fonction de rotation des fichiers de journal intégrée, supportant la rotation automatique horaire des fichiers de journal
    icon: 🔄
  - title: Compatibilité Zap
    details: Intégration transparente avec zap WriteSyncer, supportant les cibles de sortie personnalisées
    icon: 🔗
  - title: API simple
    details: API conçue similaire à la bibliothèque de journalisation standard, facile à utiliser et à migrer
    icon: 🚀
  - title: Thread-safe
    details: Conception sans verrou pour la plupart des opérations, assurant la sécurité des threads sans surcharge de performance
    icon: 🔒
  - title: Prêt pour la production
    details: Testé en conditions réelles dans des environnements de production avec une couverture de tests complète
    icon: ✅
