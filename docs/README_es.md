# 🚀 LazyGophers Log

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.5%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Una biblioteca de logging de alto rendimiento y rica en características para aplicaciones Go con soporte multi-etiquetas de construcción, escritura asíncrona y amplias opciones de personalización.

## 📖 Idiomas de documentación

- [🇺🇸 English](../README.md)
- [🇨🇳 简体中文](README.zh-CN.md)
- [🇹🇼 繁體中文](README.zh-TW.md)
- [🇫🇷 Français](README.fr.md)
- [🇷🇺 Русский](README.ru.md)
- [🇪🇸 Español](README.es.md) (Actual)
- [🇸🇦 العربية](README.ar.md)

## ✨ Características

- **🚀 Alto rendimiento**: Soporte de pooling de objetos y escritura asíncrona
- **🏗️ Soporte de etiquetas de construcción**: Diferentes comportamientos para modos debug, release y discard
- **🔄 Rotación de logs**: Rotación automática de archivos de log por hora
- **🎨 Formateo enriquecido**: Formatos de log personalizables con soporte de colores
- **🔍 Trazado contextual**: Seguimiento de ID de Goroutine e ID de traza
- **🔌 Integración de frameworks**: Integración nativa con logger Zap
- **⚙️ Altamente configurable**: Niveles flexibles, salidas y formateo
- **🧪 Bien probado**: 93.5% de cobertura de pruebas en todas las configuraciones de construcción

## 🚀 Inicio rápido

### Instalación

```bash
go get github.com/lazygophers/log
```

### Uso básico

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Logging simple
    log.Info("¡Hola, Mundo!")
    log.Debug("Este es un mensaje de depuración")
    log.Warn("Esta es una advertencia")
    log.Error("Este es un error")

    // Logging formateado
    log.Infof("El usuario %s se conectó con ID %d", "juan", 123)
    
    // Con logger personalizado
    logger := log.New()
    logger.SetLevel(log.InfoLevel)
    logger.Info("Mensaje de logger personalizado")
}
```

### Uso avanzado

```go
package main

import (
    "context"
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Crear logger con salida a archivo
    logger := log.New()
    
    // Configurar salida a archivo con rotación horaria
    writer := log.GetOutputWriterHourly("./logs/app.log")
    logger.SetOutput(writer)
    
    // Configurar formateo
    logger.SetLevel(log.DebugLevel)
    logger.SetPrefixMsg("[APP] ")
    logger.Caller(true) // Habilitar información del caller
    
    // Logging contextual
    ctxLogger := logger.CloneToCtx()
    ctxLogger.Info(context.Background(), "Logging consciente del contexto")
    
    // Logging asíncrono para escenarios de alto rendimiento
    asyncWriter := log.NewAsyncWriter(writer, 1000)
    logger.SetOutput(asyncWriter)
    defer asyncWriter.Close()
    
    logger.Info("Logging asíncrono de alto rendimiento")
}
```

## 🏗️ Etiquetas de construcción

La biblioteca soporta diferentes modos de construcción a través de etiquetas de construcción de Go:

### Modo por defecto (Sin etiquetas)
```bash
go build
```
- Funcionalidad completa de logging
- Mensajes de depuración habilitados
- Rendimiento estándar

### Modo debug
```bash
go build -tags debug
```
- Información de depuración mejorada
- Información detallada del caller
- Soporte de perfilado de rendimiento

### Modo release
```bash
go build -tags release
```
- Optimizado para producción
- Mensajes de depuración deshabilitados
- Rotación automática de archivos de log

### Modo discard
```bash
go build -tags discard
```
- Rendimiento máximo
- Todos los logs son descartados
- Cero sobrecarga de logging

### Modos combinados
```bash
go build -tags "debug,discard"    # Debug con discard
go build -tags "release,discard"  # Release con discard
```

## 📊 Niveles de log

La biblioteca soporta 7 niveles de log (de mayor a menor prioridad):

| Nivel | Valor | Descripción |
|-------|-------|-------------|
| `PanicLevel` | 0 | Registra y luego llama a panic |
| `FatalLevel` | 1 | Registra y luego llama a os.Exit(1) |
| `ErrorLevel` | 2 | Condiciones de error |
| `WarnLevel` | 3 | Condiciones de advertencia |
| `InfoLevel` | 4 | Mensajes informativos |
| `DebugLevel` | 5 | Mensajes de nivel de depuración |
| `TraceLevel` | 6 | Logging más detallado |

## 🔌 Integración de frameworks

### Integración con Zap

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/lazygophers/log"
)

// Crear un logger zap que escriba a nuestro sistema de logs
logger := log.New()
hook := log.NewZapHook(logger)

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    hook,
    zapcore.InfoLevel,
)
zapLogger := zap.New(core)

zapLogger.Info("Mensaje de Zap", zap.String("key", "value"))
```

## 🧪 Pruebas

La biblioteca viene con soporte completo de pruebas:

```bash
# Ejecutar todas las pruebas
make test

# Ejecutar pruebas con cobertura para todas las etiquetas de construcción
make coverage-all

# Prueba rápida a través de todas las etiquetas de construcción
make test-quick

# Generar reportes de cobertura HTML
make coverage-html
```

### Resultados de cobertura por etiqueta de construcción

| Etiqueta de construcción | Cobertura |
|--------------------------|-----------|
| Por defecto | 92.9% |
| Debug | 93.1% |
| Release | 93.5% |
| Discard | 93.1% |
| Debug+Discard | 93.1% |
| Release+Discard | 93.3% |

## ⚙️ Opciones de configuración

### Configuración del logger

```go
logger := log.New()

// Establecer nivel mínimo de log
logger.SetLevel(log.InfoLevel)

// Configurar salida
logger.SetOutput(os.Stdout) // Un solo writer
logger.SetOutput(writer1, writer2, writer3) // Múltiples writers

// Personalizar mensajes
logger.SetPrefixMsg("[MiApp] ")
logger.SetSuffixMsg(" [FIN]")
logger.AppendPrefixMsg("Extra: ")

// Configurar formateo
logger.ParsingAndEscaping(false) // Deshabilitar secuencias de escape
logger.Caller(true) // Habilitar información del caller
logger.SetCallerDepth(4) // Ajustar profundidad del stack del caller
```

## 📁 Rotación de logs

Rotación automática de logs con intervalos configurables:

```go
// Rotación horaria
writer := log.GetOutputWriterHourly("./logs/app.log")

// La biblioteca creará archivos como:
// - app-2024010115.log (2024-01-01 15:00)
// - app-2024010116.log (2024-01-01 16:00)
// - app-2024010117.log (2024-01-01 17:00)
```

## 🔍 Contexto y trazado

Soporte incorporado para logging consciente del contexto y trazado distribuido:

```go
// Establecer ID de traza para la goroutine actual
log.SetTrace("trace-123-456")

// Obtener ID de traza
traceID := log.GetTrace()

// Logging consciente del contexto
ctx := context.Background()
ctxLogger := log.CloneToCtx()
ctxLogger.Info(ctx, "Solicitud procesada", "user_id", 123)

// Seguimiento automático de ID de goroutine
log.Info("Este log incluye automáticamente el ID de goroutine")
```

## 📈 Rendimiento

La biblioteca está diseñada para aplicaciones de alto rendimiento:

- **Pooling de objetos**: Reutiliza objetos de entrada de log para reducir presión de GC
- **Escritura asíncrona**: Escrituras de log no bloqueantes para escenarios de alto rendimiento
- **Filtrado de nivel**: El filtrado temprano previene operaciones costosas
- **Optimización de etiquetas de construcción**: Optimización en tiempo de compilación para diferentes entornos

### Benchmarks

```bash
# Ejecutar benchmarks de rendimiento
make benchmark

# Benchmark de diferentes modos de construcción
make benchmark-debug
make benchmark-release  
make benchmark-discard
```

## 🤝 Contribuir

¡Damos la bienvenida a las contribuciones! Por favor consulta nuestra [Guía de Contribución](CONTRIBUTING.md) para detalles.

### Configuración de desarrollo

1. **Fork y Clone**
   ```bash
   git clone https://github.com/your-username/log.git
   cd log
   ```

2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Ejecutar pruebas**
   ```bash
   make test-all
   ```

4. **Enviar Pull Request**
   - Sigue nuestro [Template de PR](../.github/pull_request_template.md)
   - Asegúrate de que las pruebas pasen
   - Actualiza la documentación si es necesario

## 📋 Requisitos

- **Go**: 1.21 o superior
- **Dependencias**: 
  - `go.uber.org/zap` (para integración con Zap)
  - `github.com/petermattis/goid` (para ID de goroutine)
  - `github.com/lestrrat-go/file-rotatelogs` (para rotación de logs)
  - `github.com/google/uuid` (para ID de traza)

## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT - consulta el archivo [LICENSE](../LICENSE) para detalles.

## 🙏 Agradecimientos

- [Zap](https://github.com/uber-go/zap) por la inspiración y soporte de integración
- [Logrus](https://github.com/sirupsen/logrus) por los patrones de diseño de niveles
- La comunidad Go por la retroalimentación continua y mejoras

## 📞 Soporte

- 📖 [Documentación](../docs/)
- 🐛 [Rastreador de problemas](https://github.com/lazygophers/log/issues)
- 💬 [Discusiones](https://github.com/lazygophers/log/discussions)
- 📧 Email: support@lazygophers.com

---

**Hecho con ❤️ por el equipo LazyGophers**