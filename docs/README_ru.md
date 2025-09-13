# 🚀 LazyGophers Log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![DeepWiki](https://img.shields.io/badge/DeepWiki-documented-blue?logo=bookstack&logoColor=white)](https://deepwiki.ai/docs/lazygophers/log)
[![Go.Dev Downloads](https://pkg.go.dev/badge/github.com/lazygophers/log.svg)](https://pkg.go.dev/github.com/lazygophers/log)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/log)
[![Goproxy.io](https://goproxy.io/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.io/stats/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Высокопроизводительная, многофункциональная библиотека логирования для Go приложений с поддержкой мультитегов сборки, асинхронной записи и обширными возможностями настройки.

## 📖 Языки документации

- [🇺🇸 English](../README.md)
- [🇨🇳 简体中文](README.zh-CN.md)
- [🇹🇼 繁體中文](README.zh-TW.md)
- [🇫🇷 Français](README.fr.md)
- [🇷🇺 Русский](README.ru.md) (Текущий)
- [🇪🇸 Español](README.es.md)
- [🇸🇦 العربية](README.ar.md)

## ✨ Особенности

- **🚀 Высокая производительность**: Поддержка пула объектов и асинхронной записи
- **🏗️ Поддержка тегов сборки**: Различное поведение для режимов debug, release и discard
- **🔄 Ротация логов**: Автоматическая почасовая ротация файлов логов
- **🎨 Богатое форматирование**: Настраиваемые форматы логов с поддержкой цветов
- **🔍 Контекстное трассирование**: Отслеживание ID горутин и ID трассировки
- **🔌 Интеграция с фреймворками**: Нативная интеграция с Zap логгером
- **⚙️ Высокая настраиваемость**: Гибкие уровни, выводы и форматирование
- **🧪 Хорошо протестировано**: 93.0% покрытие тестами во всех конфигурациях сборки

## 🚀 Быстрый старт

### Установка

```bash
go get github.com/lazygophers/log
```

### Базовое использование

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Простое логирование
    log.Info("Привет, Мир!")
    log.Debug("Это отладочное сообщение")
    log.Warn("Это предупреждение")
    log.Error("Это ошибка")

    // Форматированное логирование
    log.Infof("Пользователь %s вошел в систему с ID %d", "ivan", 123)
    
    // С кастомным логгером
    logger := log.New()
    logger.SetLevel(log.InfoLevel)
    logger.Info("Сообщение кастомного логгера")
}
```

### Продвинутое использование

```go
package main

import (
    "context"
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Создать логгер с выводом в файл
    logger := log.New()
    
    // Установить вывод в файл с почасовой ротацией
    writer := log.GetOutputWriterHourly("./logs/app.log")
    logger.SetOutput(writer)
    
    // Настроить форматирование
    logger.SetLevel(log.DebugLevel)
    logger.SetPrefixMsg("[APP] ")
    logger.Caller(true) // Включить информацию о вызывающем коде
    
    // Контекстное логирование
    ctxLogger := logger.CloneToCtx()
    ctxLogger.Info(context.Background(), "Контекстно-зависимое логирование")
    
    // Асинхронное логирование для высокопроизводительных сценариев
    asyncWriter := log.NewAsyncWriter(writer, 1000)
    logger.SetOutput(asyncWriter)
    defer asyncWriter.Close()
    
    logger.Info("Высокопроизводительное асинхронное логирование")
}
```

## 🏗️ Теги сборки

Библиотека поддерживает различные режимы сборки через теги сборки Go:

### Режим по умолчанию (Без тегов)
```bash
go build
```
- Полная функциональность логирования
- Отладочные сообщения включены
- Стандартная производительность

### Режим отладки
```bash
go build -tags debug
```
- Расширенная отладочная информация
- Детальная информация о вызывающем коде
- Поддержка профилирования производительности

### Режим релиза
```bash
go build -tags release
```
- Оптимизировано для продакшена
- Отладочные сообщения отключены
- Автоматическая ротация файлов логов

### Режим отбрасывания
```bash
go build -tags discard
```
- Максимальная производительность
- Все логи отбрасываются
- Нулевые накладные расходы на логирование

### Комбинированные режимы
```bash
go build -tags "debug,discard"    # Отладка с отбрасыванием
go build -tags "release,discard"  # Релиз с отбрасыванием
```

## 📊 Уровни логов

Библиотека поддерживает 7 уровней логов (от высшего к низшему приоритету):

| Уровень | Значение | Описание |
|---------|----------|----------|
| `PanicLevel` | 0 | Логирует и затем вызывает panic |
| `FatalLevel` | 1 | Логирует и затем вызывает os.Exit(1) |
| `ErrorLevel` | 2 | Условия ошибок |
| `WarnLevel` | 3 | Условия предупреждений |
| `InfoLevel` | 4 | Информационные сообщения |
| `DebugLevel` | 5 | Сообщения уровня отладки |
| `TraceLevel` | 6 | Самое подробное логирование |

## 🔌 Интеграция с фреймворками

### Интеграция с Zap

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/lazygophers/log"
)

// Создать zap логгер, который пишет в нашу систему логирования
logger := log.New()
hook := log.NewZapHook(logger)

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    hook,
    zapcore.InfoLevel,
)
zapLogger := zap.New(core)

zapLogger.Info("Сообщение от Zap", zap.String("key", "value"))
```

## 🧪 Тестирование

Библиотека поставляется с комплексной поддержкой тестирования:

```bash
# Запустить все тесты
make test

# Запустить тесты с покрытием для всех тегов сборки
make coverage-all

# Быстрое тестирование всех тегов сборки
make test-quick

# Генерировать HTML отчеты о покрытии
make coverage-html
```

### Результаты покрытия по тегам сборки

| Тег сборки | Покрытие |
|------------|----------|
| По умолчанию | 92.9% |
| Отладка | 93.1% |
| Релиз | 93.5% |
| Отбрасывание | 93.1% |
| Отладка+Отбрасывание | 93.1% |
| Релиз+Отбрасывание | 93.3% |

## ⚙️ Опции конфигурации

### Конфигурация логгера

```go
logger := log.New()

// Установить минимальный уровень логирования
logger.SetLevel(log.InfoLevel)

// Настроить вывод
logger.SetOutput(os.Stdout) // Один писатель
logger.SetOutput(writer1, writer2, writer3) // Несколько писателей

// Настроить сообщения
logger.SetPrefixMsg("[МоеПриложение] ")
logger.SetSuffixMsg(" [КОНЕЦ]")
logger.AppendPrefixMsg("Дополнительно: ")

// Настроить форматирование
logger.ParsingAndEscaping(false) // Отключить escape-последовательности
logger.Caller(true) // Включить информацию о вызывающем коде
logger.SetCallerDepth(4) // Настроить глубину стека вызовов
```

## 📁 Ротация логов

Автоматическая ротация логов с настраиваемыми интервалами:

```go
// Почасовая ротация
writer := log.GetOutputWriterHourly("./logs/app.log")

// Библиотека создаст файлы типа:
// - app-2024010115.log (2024-01-01 15:00)
// - app-2024010116.log (2024-01-01 16:00)
// - app-2024010117.log (2024-01-01 17:00)
```

## 🔍 Контекст и трассировка

Встроенная поддержка контекстно-зависимого логирования и распределенной трассировки:

```go
// Установить ID трассировки для текущей горутины
log.SetTrace("trace-123-456")

// Получить ID трассировки
traceID := log.GetTrace()

// Контекстно-зависимое логирование
ctx := context.Background()
ctxLogger := log.CloneToCtx()
ctxLogger.Info(ctx, "Запрос обработан", "user_id", 123)

// Автоматическое отслеживание ID горутины
log.Info("Этот лог автоматически включает ID горутины")
```

## 📈 Производительность

Библиотека спроектирована для высокопроизводительных приложений:

- **Пул объектов**: Переиспользует объекты записей логов для снижения давления на GC
- **Асинхронная запись**: Неблокирующая запись логов для высокопроизводительных сценариев
- **Фильтрация по уровням**: Раннее фильтрование предотвращает дорогие операции
- **Оптимизация тегов сборки**: Оптимизация во время компиляции для различных сред

### Бенчмарки

```bash
# Запустить бенчмарки производительности
make benchmark

# Бенчмарки различных режимов сборки
make benchmark-debug
make benchmark-release  
make benchmark-discard
```

## 🤝 Вклад в проект

Мы приветствуем вклад в проект! Пожалуйста, ознакомьтесь с нашим [Руководством по вкладу](CONTRIBUTING.md) для деталей.

### Настройка среды разработки

1. **Fork и Clone**
   ```bash
   git clone https://github.com/your-username/log.git
   cd log
   ```

2. **Установить зависимости**
   ```bash
   go mod tidy
   ```

3. **Запустить тесты**
   ```bash
   make test-all
   ```

4. **Отправить Pull Request**
   - Следуйте нашему [Шаблону PR](../.github/pull_request_template.md)
   - Убедитесь, что тесты проходят
   - Обновите документацию при необходимости

## 📋 Требования

- **Go**: 1.19 или выше
- **Зависимости**: 
  - `go.uber.org/zap` (для интеграции с Zap)
  - `github.com/petermattis/goid` (для ID горутин)
  - `github.com/lestrrat-go/file-rotatelogs` (для ротации логов)
  - `github.com/google/uuid` (для ID трассировки)

## 📄 Лицензия

Этот проект лицензирован под MIT License - см. файл [LICENSE](../LICENSE) для деталей.

## 🙏 Благодарности

- [Zap](https://github.com/uber-go/zap) за вдохновение и поддержку интеграции
- [Logrus](https://github.com/sirupsen/logrus) за шаблоны дизайна уровней
- Сообщество Go за непрерывную обратную связь и улучшения

## 📞 Поддержка

- 📖 [Документация](../docs/)
- 🐛 [Трекер проблем](https://github.com/lazygophers/log/issues)
- 💬 [Обсуждения](https://github.com/lazygophers/log/discussions)
- 📧 Email: support@lazygophers.com

---

**Сделано с ❤️ командой LazyGophers**