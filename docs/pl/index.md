---
pageType: home

hero:
    name: LazyGophers Log
    text: Wysokowydajna i elastyczna biblioteka logowania Go
    tagline: Zbudowana na zap, zapewnia bogate funkcje i prosty interfejs API
    actions:
        - theme: brand
          text: Szybki Start
          link: /API
        - theme: alt
          text: Zobacz na GitHub
          link: https://github.com/lazygophers/log

features:
    - title: "Wysoka Wydajność"
      details: Zbudowana na zap z ponownym użyciem puli obiektów i warunkowym rejestrowaniem pól
      icon: 🚀
    - title: "Bogate Poziomy Logowania"
      details: Obsługuje poziomy Trace, Debug, Info, Warn, Error, Fatal, Panic
      icon: 📊
    - title: "Elastyczna Konfiguracja"
      details: Dostosuj poziomy, informacje o wywołującym, śledzenie, prefiksy, sufiksy i cele wyjściowe
      icon: ⚙️
    - title: "Rotacja Plików"
      details: Wbudowana obsługa rotacji plików logowania co godzinę
      icon: 🔄
    - title: "Kompatybilność z Zap"
      details: Bezproblemowa integracja z zap WriteSyncer
      icon: 🔌
    - title: "Proste API"
      details: Jasny interfejs API podobny do standardowej biblioteki logowania, łatwy w użyciu i integracji
      icon: 🎯
---

## Szybki Start

### Instalacja

```bash
go get github.com/lazygophers/log
```

### Podstawowe Użycie

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Użyj domyślnego globalnego loggera
    log.Debug("Informacje o debugowaniu")
    log.Info("Informacje ogólne")
    log.Warn("Informacje ostrzegawcze")
    log.Error("Informacje o błędzie")

    // Użyj sformatowanego wyjścia
    log.Infof("Użytkownik %s pomyślnie zalogowany", "admin")

    // Niestandardowa konfiguracja
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("To jest log z niestandardowego loggera")
}
```

## Dokumentacja

-   [Dokumentacja API](API.md) - Pełna dokumentacja interfejsu API
-   [Dziennik Zmian](/pl/CHANGELOG) - Historia wersji
-   [Przewodnik po Wkładzie](/pl/CONTRIBUTING) - Jak wnosić wkład
-   [Polityka Bezpieczeństwa](/pl/SECURITY) - Przewodnik bezpieczeństwa
-   [Kodeks Postępowania](/pl/CODE_OF_CONDUCT) - Wytyczne społeczności

## Porównanie Wydajności

| Funkcja       | lazygophers/log | zap | logrus | Standardowy log |
| ---------- | --------------- | --- | ------ | -------- |
| Wydajność       | Wysoka              | Wysoka  | Średnia     | Niska       |
| Prostota API    | Wysoka              | Średnia  | Wysoka     | Wysoka       |
| Bogactwo Funkcji    | Średnia              | Wysoka  | Wysoka     | Niska       |
| Elastyczność      | Średnia              | Wysoka  | Wysoka     | Niska       |
| Krzywa Uczenia      | Niska              | Średnia  | Średnia     | Niska       |

## Licencja

Ten projekt jest licencjonowany na Licencji MIT - zobacz plik [LICENSE](/pl/LICENSE) dla szczegółów.
