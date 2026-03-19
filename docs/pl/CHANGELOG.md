---
titleSuffix: " | LazyGophers Log"
---

# 📋 Dziennik zmian

Wszystkie ważne zmiany w projekcie są rejestrowane w tym pliku.

Format oparty na [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), projekt Following [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Nieopublikowane]

### Dodano

-   Pełna wielojęzyczna dokumentacja (7 języków)
-   Szablony GitHub issue (zgłoszenia błędów, prośby o funkcje, pytania)
-   Szablon Pull Request ze sprawdzaniem kompatybilności tagów kompilacji
-   Wielojęzyczny przewodnik współtworzenia
-   Kodeks postępowania z wytycznymi egzekwowania
-   Polityka bezpieczeństwa z procesem zgłaszania luk
-   Pełna dokumentacja API z przykładami
-   Profesjonalna struktura projektu i szablony

### Zmieniono

-   Rozszerzono README o pełną dokumentację funkcji
-   Poprawiono pokrycie testów dla wszystkich konfiguracji tagów kompilacji
-   Zaktualizowano strukturę projektu w celu poprawy maintainability

### Dokumentacja

-   Dodano obsługę wielojęzyczną dla wszystkich głównych dokumentów
-   Utworzono kompletne odniesienie API
-   Ustanowiono wytyczne przepływu pracy współtworzenia
-   Wdrożono proces zgłaszania problemów bezpieczeństwa

## [1.0.0] - 2024-01-01

### Dodano

-   Podstawowe funkcje rejestrowania z wieloma poziomami (Trace, Debug, Info, Warn, Error, Fatal, Panic)
-   Implementacja loggera bezpieczna dla wątków z pulą obiektów
-   Obsługa tagów kompilacji (domyślny, debug, release, discard)
-   Niestandardowy interfejs formatowania z domyślnym formaterem tekstowym
-   Obsługa wyjścia wielowriterowego
-   Funkcja asynchronicznego pisania dla scenariuszy o wysokiej przepustowości
-   Automatyczna rotacja plików dziennika co godzinę
-   Rejestrowanie świadome kontekstu z śledzeniem ID goroutine i ID śledzenia
-   Informacje o wywołującym z konfigurowalną głębokością stosu
-   Globalne funkcje wygody na poziomie pakietu
-   Obsługa integracji z loggerem Zap

### Wydajność

-   Buforowanie obiektów obiektów wpisów i buforów przy użyciu `sync.Pool`
-   Wczesne sprawdzanie poziomów w celu uniknięcia kosztownych operacji
-   Writer asynchroniczny do nieblokującego zapisu dziennika
-   Optymalizacja tagów kompilacji dla różnych środowisk

### Tagi kompilacji

-   **Domyślne**: Pełna funkcjonalność z komunikatami debugowania
-   **Debug**: Rozszerzone informacje debugowania i szczegóły wywołującego
-   **Release**: Optymalizacja produkcyjna, wyłączone komunikaty debugowania
-   **Discard**: Maksymalna wydajność, operacje dziennika no-op

### Podstawowe funkcje

-   **Logger**: Główna struktura loggera z konfigurowalnym poziomem, wyjściem, formatowaniem
-   **Entry**: Struktura wpisu dziennika z pełnymi metadanymi
-   **Poziomy**: Siedem poziomów dziennika od Panic (najwyższy) do Trace (najniższy)
-   **Formatery**: System formatowania z możliwością podłączania
-   **Writery**: Obsługa rotacji plików i pisania asynchronicznego
-   **Kontekst**: Obsługa ID goroutine i śledzenia rozproszonego

### API highlights

-   Płynne API konfiguracji z łańcuchem metod
-   Proste i sformatowane metody rejestrowania (`.Info()` i `.Infof()`)
-   Klonowanie loggera do odizolowanej konfiguracji
-   Rejestrowanie świadome kontekstu z `CloneToCtx()`
-   Niestandardowe komunikaty prefiksu i sufiksu
-   Przełącznik informacji o wywołującym

### Testowanie

-   Kompletny zestaw testów z pokryciem 93,5%
-   Obsługa testów wielokrotnych tagów kompilacji
-   Zautomatyzowane przepływy pracy testowania
-   Testy porównawcze wydajności

## [0.9.0] - 2023-12-15

### Dodano

-   Początkowa struktura projektu
-   Podstawowe funkcje rejestrowania
-   Filtrowanie na podstawie poziomów
-   Obsługa wyjścia plikowego

### Zmieniono

-   Poprawiono wydajność poprzez buforowanie obiektów
-   Rozszerzono obsługę błędów

## [0.8.0] - 2023-12-01

### Dodano

-   Obsługa wielu writerów
-   Niestandardowy interfejs formatowania
-   Funkcja pisania asynchronicznego

### Naprawiono

-   Wycieki pamięci w scenariuszach o wysokiej przepustowości
-   Warunki wyścigowe w dostępie współbieżnym

## [0.7.0] - 2023-11-15

### Dodano

-   Obsługa tagów kompilacji do kompilacji warunkowej
-   Poziomy dziennika Trace i Debug
-   Śledzenie informacji o wywołującym

### Zmieniono

-   Zoptymalizowano wzorce alokacji pamięci
-   Poprawiono bezpieczeństwo wątków

## [0.6.0] - 2023-11-01

### Dodano

-   Funkcja rotacji dziennika
-   Rejestrowanie świadome kontekstu
-   Śledzenie ID goroutine

### Przestarzałe

-   Stare metody konfiguracji (zostaną usunięte w v1.0.0)

## [0.5.0] - 2023-10-15

### Dodano

-   Formater JSON
-   Wiele celów wyjściowych
-   Testy porównawcze wydajności

### Zmieniono

-   Zrefaktoryzowano podstawowy silnik rejestrowania
-   Poprawiono spójność API

### Usunięto

-   Stare metody rejestrowania

## [0.4.0] - 2023-10-01

### Dodano

-   Poziomy dziennika Fatal i Panic
-   Globalne funkcje pakietu
-   Walidacja konfiguracji

### Naprawiono

-   Problemy z synchronizacją wyjścia
-   Optymalizacja użycia pamięci

## [0.3.0] - 2023-09-15

### Dodano

-   Niestandardowe poziomy dziennika
-   Interfejs formatowania
-   Operacje bezpieczne dla wątków

### Zmieniono

-   Uproszczono projekt API
-   Rozszerzono dokumentację

## [0.2.0] - 2023-09-01

### Dodano

-   Obsługa wyjścia plikowego
-   Filtrowanie na podstawie poziomów
-   Podstawowe opcje formatowania

### Naprawiono

-   Wąskie gardeł wydajności
-   Wycieki pamięci

## [0.1.0] - 2023-08-15

### Dodano

-   Pierwsze wydanie
-   Podstawowe rejestrowanie konsoli
-   Prosta obsługa poziomów (Info, Warn, Error)
-   Podstawowa struktura loggera

## Podsumowanie historii wersji

| Wersja | Data wydania | Główne funkcje |
| ----- | ---------- | ---------------------------------------------- |
| 1.0.0 | 2024-01-01 | Pełny system dziennika, tagi kompilacji, pisanie asynchroniczne, pełna dokumentacja |
| 0.9.0 | 2023-12-15 | Ulepszenia wydajności, buforowanie obiektów |
| 0.8.0 | 2023-12-01 | Wiele writerów, pisanie asynchroniczne, niestandardowe formatery |
| 0.7.0 | 2023-11-15 | Tagi kompilacji, poziomy Trace/Debug, informacje o wywołującym |
| 0.6.0 | 2023-11-01 | Rotacja dziennika, dziennik kontekstowy, śledzenie goroutine |
| 0.5.0 | 2023-10-15 | Formater JSON, wiele wyjść, testy porównawcze |
| 0.4.0 | 2023-10-01 | Poziomy Fatal/Panic, funkcje globalne |
| 0.3.0 | 2023-09-15 | Niestandardowe poziomy, interfejs formatowania |
| 0.2.0 | 2023-09-01 | Wyjście plikowe, filtrowanie poziomów |
| 0.1.0 | 2023-08-15 | Pierwsze wydanie, podstawowe rejestrowanie konsoli |

## Przewodnik migracji

### Migracja z v0.9.x do v1.0.0

#### Złamane zmiany

-   Brak - v1.0.0 jest kompatybilna wstecz z v0.9.x

#### Nowe dostępne funkcje

-   Rozszerzona obsługa tagów kompilacji
-   Pełna dokumentacja
-   Profesjonalne szablony projektu
-   Proces zgłaszania problemów bezpieczeństwa

#### Zalecane aktualizacje

```go
// Stary sposób (nadal obsługiwany)
logger := log.New()
logger.SetLevel(log.InfoLevel)

// Zalecany nowy sposób, używając łańcucha metod
logger := log.New().
    SetLevel(log.InfoLevel).
    Caller(true).
    SetPrefixMsg("[MyApp] ")
```

### Migracja z v0.8.x do v0.9.x

#### Złamane zmiany

-   Usunięto przestarzałe metody konfiguracji
-   Zmieniono wewnętrzne zarządzanie buforami

#### Kroki migracji

1. Zaktualizuj ścieżki importu, jeśli to konieczne
2. Zastąp przestarzałe metody:

    ```go
    // Stare (przestarzałe)
    logger.SetOutputFile("app.log")

    // Nowe
    file, _ := os.Create("app.log")
    logger.SetOutput(file)
    ```

### Migracja z v0.5.x i wcześniejszych

#### Główne zmiany

-   Całkowite przeprojektowanie API dla lepszej spójności
-   Ulepszenie wydajności przez buforowanie obiektów
-   Nowy system tagów kompilacji

#### Wymagana migracja

-   Zaktualizuj wszystkie wywołania dziennika do nowego API
-   Przejrzyj i zaktualizuj implementacje formaterów
-   Przetestuj używając nowych konfiguracji tagów kompilacji

## Kamienie milowe rozwoju

### 🎯 v1.1.0 Plan rozwoju (planowane)

-   [ ] Rejestrowanie ustrukturyzowane z parami klucz-wartość
-   [ ] Próbkowanie dzienników dla scenariuszy o dużej objętości
-   [ ] System wtyczek do niestandardowych wyjść
-   [ ] Rozszerzone metryki wydajności
-   [ ] Integracja z dziennikami chmurowymi

### 🎯 v1.2.0 Plan rozwoju (przyszłość)

-   [ ] Obsługa plików konfiguracyjnych (YAML/JSON/TOML)
-   [ ] Agregacja i filtrowanie dzienników
-   [ ] Przesyłanie strumieniowe dzienników w czasie rzeczywistym
-   [ ] Rozszerzone funkcje bezpieczeństwa
-   [ ] Integracja z pulpitem wydajności

## Wkład

Witamy wkład! Zapoznaj się z naszym [Przewodnikiem współtworzenia](/pl/CONTRIBUTING) dla szczegółów dotyczących:

-   Zgłaszania błędów i żądania funkcji
-   Przepływu pracy przesyłania kodu
-   Konfiguracji rozwoju
-   Wymagań testowania
-   Standardów dokumentacji

## Bezpieczeństwo

W przypadku luk w zabezpieczeniach zapoznaj się z naszą [Polityką bezpieczeństwa](/pl/SECURITY), aby uzyskać informacje o:

-   Wspieranych wersjach
-   Procesie zgłaszania
-   Osiach czasu odpowiedzi
-   Najlepszych praktykach bezpieczeństwa

## Wsparcie

-   📖 [Dokumentacja](docs/)
-   🐛 [Śledzenie problemów](https://github.com/lazygophers/log/issues)
-   💬 [Dyskusje](https://github.com/lazygophers/log/discussions)
-   📧 Email: support@lazygophers.com

## Licencja

Ten projekt jest licencjonowany na licencji MIT - zobacz plik [LICENSE](/pl/LICENSE) dla szczegółów.

---

## 🌍 Dokumentacja wielojęzyczna

Ten dziennik zmian jest dostępny w wielu językach:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CHANGELOG.md)
-   [🇨🇳 简体中文](/zh-CN/CHANGELOG)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CHANGELOG.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CHANGELOG.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CHANGELOG.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CHANGELOG.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CHANGELOG.md)

---

**Śledź każdą poprawę i bądź na bieżąco z rozwojem LazygoPHers Log! 🚀**

---

_Ten dziennik zmian jest aktualizowany automatycznie przy każdym wydaniu. Aby uzyskać najnowsze informacje, sprawdź stronę [GitHub Releases](https://github.com/lazygophers/log/releases)._
