---
titleSuffix: " | LazyGophers Log"
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Una biblioteca de registro Go de alto rendimiento y flexible, construida sobre zap, que ofrece funciones ricas y una API simple.

## 📖 Idiomas de documentación

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](README.md) (actual)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)

## 🚀 Documentación en línea

Visite nuestra [documentación GitHub Pages](https://lazygophers.github.io/log/) para una mejor experiencia de lectura.

## ✨ Características

-   **🚀 Alto rendimiento**: Construido sobre zap con reutilización de objetos Entry a través de un pool, reduciendo la asignación de memoria
-   **📊 Niveles de registro ricos**: Niveles Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configuración flexible**:
    -   Control de nivel de registro
    -   Registro de información del llamador
    -   Información de traza (incluyendo ID de goroutine)
    -   Prefijos y sufijos de registro personalizados
    -   Destinos de salida personalizados (consola, archivos, etc.)
    -   Opciones de formato de registro
-   **🔄 Rotación de archivos**: Soporte para rotación horaria de archivos de registro
-   **🔌 Compatibilidad con Zap**: Integración perfecta con zap WriteSyncer
-   **🎯 API simple**: API limpia similar a la biblioteca de registro estándar, fácil de usar

## 🚀 Inicio rápido

### Instalación

:::tip Instalación
```bash
go get github.com/lazygophers/log
```
:::

### Uso básico

```go title="Inicio rápido"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Usar el logger global por defecto
    log.Debug("Información de depuración")
    log.Info("Información general")
    log.Warn("Advertencia")
    log.Error("Error")

    // Usar salida formateada
    log.Infof("El usuario %s inició sesión exitosamente", "admin")

    // Configuración personalizada
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MiApp]")

    customLogger.Info("Este es un registro del logger personalizado")
}
```

## 📚 Uso avanzado

### Logger personalizado con salida de archivo

```go title="Configuración de salida de archivo"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Crear un logger con salida de archivo
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Registro de depuración con información del llamador")
    logger.Info("Registro general con información de traza")
}
```

### Control de niveles de registro

```go title="Control de niveles de registro"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Solo los registros warn y superiores serán registrados
    logger.Debug("Esto no será registrado")  // Ignorado
    logger.Info("Esto no será registrado")   // Ignorado
    logger.Warn("Esto será registrado")    // Registrado
    logger.Error("Esto será registrado")   // Registrado
}
```

## 🎯 Casos de uso

### Casos de uso adecuados

-   **Servicios web y API backend**: Seguimiento de solicitudes, registros estructurados, monitoreo de rendimiento
-   **Arquitectura de microservicios**: Seguimiento distribuido (TraceID), formato de registro uniforme, alta吞吐量
-   **Herramientas de línea de comandos**: Control de niveles, salida concisa, informe de errores
-   **Tareas de procesamiento por lotes**: Rotación de archivos, ejecución prolongada, optimización de recursos

### Ventajas especiales

-   **Optimización de pool de objetos**: Reutilización de objetos Entry y Buffer, reducción de presión GC
-   **Escritura asíncrona**: Escenarios de alta吞吐量 (10000+ registros/segundo) sin bloqueo
-   **Soporte TraceID**: Seguimiento de solicitudes de sistemas distribuidos, integración con OpenTelemetry
-   **Inicio sin configuración**: Listo para usar, configuración progresiva

## 🔧 Opciones de configuración

:::note Opciones de configuración
Todos los métodos admiten encadenamiento, pueden combinarse para construir un Logger personalizado.
:::

### Configuración del Logger

| Método                  | Descripción                         | Valor predeterminado |
| ----------------------- | ----------------------------------- | ------------------- |
| `SetLevel(level)`       | Establecer el nivel mínimo de registro | `DebugLevel`        |
| `EnableCaller(enable)`  | Habilitar/deshabilitar registro de información del llamador | `false`      |
| `EnableTrace(enable)`   | Habilitar/deshabilitar información de traza | `false`      |
| `SetCallerDepth(depth)` | Establecer la profundidad del llamador | `2`          |
| `SetPrefixMsg(prefix)`  | Establecer prefijo de registro      | `""`         |
| `SetSuffixMsg(suffix)`  | Establecer sufijo de registro      | `""`         |
| `SetOutput(writers...)` | Establecer destinos de salida      | `os.Stdout`  |

### Niveles de registro

| Nivel       | Descripción                                  |
| ----------- | -------------------------------------------- |
| `TraceLevel` | El más detallado, para seguimiento detallado |
| `DebugLevel` | Información de depuración                    |
| `InfoLevel`  | Mensajes informativos                        |
| `WarnLevel`  | Condiciones de advertencia                   |
| `ErrorLevel` | Condiciones de error                         |
| `FatalLevel` | Errores fatales (llama a os.Exit(1))        |
| `PanicLevel` | Errores de pánico (llama a panic())         |

## 🏗️ Arquitectura

### Componentes principales

-   **Logger**: Estructura de registro principal con nivel, salida, formateador y profundidad de llamador configurables
-   **Entry**: Registro individual con soporte completo de metadatos
-   **Level**: Definiciones de nivel de registro y funciones de utilidad
-   **Format**: Interfaz e implementación de formato de registro

### Optimizaciones de rendimiento

-   **Pool de objetos**: Reutilización de objetos Entry para reducir asignación de memoria
-   **Registro condicional**: Registro de campos costosos solo cuando es necesario
-   **Verificación rápida de niveles**: Verificación de niveles en el nivel más externo
-   **Diseño sin bloqueo**: La mayoría de las operaciones no necesitan bloqueos

## 📊 Comparación de rendimiento

:::info Comparación de rendimiento
Los siguientes datos se basan en benchmarks, el rendimiento real puede variar según el entorno y la configuración.
:::

| Característica       | lazygophers/log | zap    | logrus | Registro estándar |
| -------------------- | --------------- | ------ | ------ | ----------------- |
| Rendimiento          | Alto           | Alto   | Medio  | Bajo             |
| Simplicidad de API   | Alto           | Medio  | Alto   | Alto             |
| Riqueza de funciones | Medio          | Alto   | Alto   | Bajo             |
| Flexibilidad         | Medio          | Alto   | Alto   | Bajo             |
| Curva de aprendizaje | Baja          | Medio  | Medio  | Bajo             |

## ❓ Preguntas frecuentes

### ¿Cómo elegir el nivel de registro adecuado?

-   **Entorno de desarrollo**: Usar `DebugLevel` o `TraceLevel` para obtener información detallada
-   **Entorno de producción**: Usar `InfoLevel` o `WarnLevel` para reducir la sobrecarga
-   **Pruebas de rendimiento**: Usar `PanicLevel` para deshabilitar todos los registros

### ¿Cómo optimizar el rendimiento en producción?

:::warning Advertencia
En escenarios de alta吞吐量, se recomienda combinar escritura asíncrona y niveles de registro apropiados para optimizar el rendimiento.
:::

1. Usar `AsyncWriter` para escritura asíncrona:

```go title="Configuración de escritura asíncrona"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Ajustar niveles de registro para evitar registros innecesarios:

```go title="Optimización de niveles"
logger.SetLevel(log.InfoLevel)  // Omitir Debug y Trace
```

3. Usar registros condicionales para reducir sobrecarga:

```go title="Registros condicionales"
if logger.Level >= log.DebugLevel {
    logger.Debug("Información de depuración detallada")
}
```

### ¿Cuál es la diferencia entre `Caller` y `EnableCaller`?

-   **`EnableCaller(enable bool)`**: Controla si el Logger recopila información del llamante
    -   `EnableCaller(true)` habilita la recopilación de información del llamante
-   **`Caller(disable bool)`**: Controla si el Formato muestra información del llamante
    -   `Caller(true)` deshabilita la visualización de información del llamante

Se recomienda usar `EnableCaller` para control global.

### ¿Cómo implementar un formateador personalizado?

Implementar la interfaz `Format`:

```go title="Formateador personalizado"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 Documentación relacionada

-   [📚 Documentación API](API.md) - Referencia API completa
-   [🤝 Guía de contribución](CONTRIBUTING.md) - Cómo contribuir
-   [📋 Registro de cambios](CHANGELOG.md) - Historial de versiones
-   [🔒 Política de seguridad](SECURITY.md) - Guía de seguridad
-   [📜 Código de conducta](CODE_OF_CONDUCT.md) - Estándares comunitarios

## 🚀 Obtener ayuda

-   **GitHub Issues**: [Informar de un bug o solicitar una función](https://github.com/lazygophers/log/issues)
-   **GoDoc**: [Documentación API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Ejemplos**: [Ejemplos de uso](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licencia

Este proyecto está bajo licencia MIT - ver archivo [LICENSE] para detalles.

## 🤝 Contribuir

¡Las contribuciones son bienvenidas! Por favor, consulte nuestra [guía de contribución](CONTRIBUTING.md) para más detalles.

---

**lazygophers/log** tiene como ser la solución de registro preferida para desarrolladores Go que valoran el rendimiento y la simplicidad. Ya sea que esté construyendo herramientas pequeñas o sistemas distribuidos a gran escala, esta biblioteca ofrece un equilibrio ideal entre funcionalidad y facilidad de uso.