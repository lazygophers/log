---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Una biblioteca de registro Go de alto rendimiento y flexible, construida sobre zap, que ofrece funciones ricas y una API sencilla.

## 📖 Idiomas de documentación

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 Chino simplificado](https://lazygophers.github.io/log/zh-CN/)
-   [🇹🇼 Chino tradicional](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇪🇸 Español](README.md) (actual)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/)
-   [🇵🇹 Português](https://lazygophers.github.io/log/pt/)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/)
-   [🇵🇱 Polski](https://lazygophers.github.io/log/pl/)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/)
-   [🇹🇷 Türkçe](https://lazygophers.github.io/log/tr/)

## ✨ Características

-   **🚀 Alto rendimiento**: Construido sobre zap con agrupación de objetos y registro de campos condicionales
-   **📊 Niveles de registro ricos**: Niveles Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configuración flexible**:
    -   Control del nivel de registro
    -   Registro de información del llamador
    -   Información de trazado (incluido ID de goroutine)
    -   Prefijos y sufijos de registro personalizados
    -   Objetivos de salida personalizados (consola, archivos, etc.)
    -   Opciones de formato de registro
-   **🔄 Rotación de archivos**: Soporte para rotación de archivos de registro cada hora
-   **🔌 Compatibilidad con Zap**: Integración perfecta con zap WriteSyncer
-   **🎯 API sencilla**: Interfaz clara similar a la biblioteca de registro estándar, fácil de usar

## 🚀 Empezando

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
    // Usar el logger global predeterminado
    log.Debug("Mensaje de depuración")
    log.Info("Mensaje informativo")
    log.Warn("Mensaje de advertencia")
    log.Error("Mensaje de error")

    // Usar salida formateada
    log.Infof("Usuario %s ha iniciado sesión correctamente", "admin")

    // Configuración personalizada
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MiApp]")

    customLogger.Info("Este es un registro desde un logger personalizado")
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
    // Crear logger con salida de archivo
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Registro de depuración con información del llamador")
    logger.Info("Registro informativo con información de trazado")
}
```

### Control del nivel de registro

```go title="Control del nivel de registro"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Solo se registrarán nivel warn y superiores
    logger.Debug("Esto no se registrará")  // Ignorado
    logger.Info("Esto no se registrará")   // Ignorado
    logger.Warn("Esto se registrará")    // Registrado
    logger.Error("Esto se registrará")   // Registrado
}
```

## 🎯 Escenarios de uso

### Escenarios aplicables

-   **Servicios web y backend API**: Seguimiento de solicitudes, registro estructurado, monitoreo de rendimiento
-   **Arquitectura de microservicios**: Trazado distribuido (TraceID), formato de registro unificado, alto rendimiento
-   **Herramientas de línea de comandos**: Control de niveles, salida limpia, reportes de errores
-   **Tareas por lotes**: Rotación de archivos, ejecución prolongada, optimización de recursos

### Ventajas especiales

-   **Optimización con agrupación de objetos**: Reutilización de objetos Entry y Buffer, reduciendo la presión del GC
-   **Escritura asíncrona**: Alta tasa de transferencia (10000+ registros/segundo) sin bloqueos
-   **Soporte para TraceID**: Seguimiento de solicitudes en sistemas distribuidos, integración con OpenTelemetry
-   **Inicio sin configuración**: Listo para usar, configuración progresiva

## 🔧 Opciones de configuración

:::note Opciones de configuración
Todos los siguientes métodos admiten llamada en cadena y se pueden combinar para construir un Logger personalizado.
:::

### Configuración del Logger

| Método                  | Descripción                | Valor por defecto       |
| --------------------- | ------------------- | ----------- |
| `SetLevel(level)`       | Establecer el nivel de registro mínimo     | `DebugLevel` |
| `EnableCaller(enable)`  | Habilitar/deshabilitar información del llamador  | `false`      |
| `EnableTrace(enable)`   | Habilitar/deshabilitar información de trazado    | `false`      |
| `SetCallerDepth(depth)` | Establecer profundidad del llamador       | `2`          |
| `SetPrefixMsg(prefix)`  | Establecer prefijo de registro         | `""`         |
| `SetSuffixMsg(suffix)`  | Establecer sufijo de registro         | `""`         |
| `SetOutput(writers...)` | Establecer objetivos de salida         | `os.Stdout`  |

### Niveles de registro

| Nivel        | Descripción                        |
| ----------- | --------------------------- |
| `TraceLevel` | Más detallado, para seguimiento detallado        |
| `DebugLevel` | Información de depuración                    |
| `InfoLevel`  | Información general                    |
| `WarnLevel`  | Mensajes de advertencia                    |
| `ErrorLevel` | Mensajes de error                    |
| `FatalLevel` | Errores fatales (llama os.Exit(1))    |
| `PanicLevel` | Errores de pánico (llama panic())    |

## 🏗️ Arquitectura

### Componentes principales

-   **Logger**: Estructura principal de registro con opciones configurables
-   **Entry**: Registro individual con soporte amplio para metadatos
-   **Level**: Definiciones de niveles de registro y funciones de utilidad
-   **Format**: Interfaz de formato de registro e implementaciones

### Optimización de rendimiento

-   **Agrupación de objetos**: Reutilización de objetos Entry para reducir la asignación de memoria
-   **Registro condicional**: Solo registra campos costosos cuando son necesarios
-   **Verificación rápida de niveles**: Verifica el nivel de registro en la capa externa
-   **Diseño sin bloqueos**: La mayoría de las operaciones no requieren bloqueos

## 📊 Comparación de rendimiento

:::info Comparación de rendimiento
Los siguientes datos se basan en pruebas de referencia; el rendimiento real puede variar según el entorno y la configuración.
:::

| Característica          | lazygophers/log | zap    | logrus | registro estándar |
| ------------- | --------------- | ------ | ------ | -------- |
| Rendimiento      | Alto              | Alto     | Medio     | Bajo       |
| Simplicidad de API    | Alto              | Medio     | Alto     | Alto       |
| Riqueza de funciones    | Medio          | Alto     | Alto     | Bajo       |
| Flexibilidad      | Medio          | Alto     | Alto     | Bajo       |
| Curva de aprendizaje      | Bajo              | Medio     | Medio     | Bajo       |

## ❓ Preguntas frecuentes

### ¿Cómo elegir el nivel de registro adecuado?

-   **Entorno de desarrollo**：Usar `DebugLevel` o `TraceLevel` para obtener información detallada
-   **Entorno de producción**：Usar `InfoLevel` o `WarnLevel` para reducir el costo
-   **Pruebas de rendimiento**：Usar `PanicLevel` para deshabilitar todos los registros

### ¿Cómo optimizar el rendimiento en producción?

:::warning Nota
En escenarios de alta tasa de transferencia, se recomienda combinar escritura asíncrona con niveles de registro razonables para optimizar el rendimiento.
:::

1. Usar `AsyncWriter` para escritura asíncrona：

```go title="Configuración de escritura asíncrona"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Ajustar el nivel de registro para evitar registros innecesarios：

```go title="Optimización de nivel"
logger.SetLevel(log.InfoLevel)  // Saltar Debug y Trace
```

3. Usar registros condicionales para reducir el costo：

```go title="Registros condicionales"
if logger.Level >= log.DebugLevel {
    logger.Debug("Información de depuración detallada")
}
```

### ¿Cuál es la diferencia entre `Caller` y `EnableCaller`?

-   **`EnableCaller(enable bool)`**：Controla si el Logger recopila información del llamador
    -   `EnableCaller(true)` habilita la recopilación de información del llamador
-   **`Caller(disable bool)`**：Controla si el Formatter muestra información del llamador
    -   `Caller(true)` deshabilita la salida de información del llamador

Se recomienda usar `EnableCaller` para control global.

### ¿Cómo implementar un formateador personalizado?

Implemente la interfaz `Format`：

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
-   [🤝 Guía de contribución](/es/CONTRIBUTING) - Cómo contribuir
-   [📋 Registro de cambios](/es/CHANGELOG) - Historial de versiones
-   [🔒 Política de seguridad](/es/SECURITY) - Guía de seguridad
-   [📜 Código de conducta](/es/CODE_OF_CONDUCT) - Normas de la comunidad

## 🚀 Obtener ayuda

-   **GitHub Issues**：[Reportar bugs o solicitar características](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[Documentación API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Ejemplos**：[Ejemplos de uso](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT - consulte el archivo [LICENSE](/es/LICENSE) para obtener más detalles.

## 🤝 Contribución

¡Bienvenidos los contributions! Consulte nuestra [Guía de contribución](/es/CONTRIBUTING) para obtener más información.

---

**lazygophers/log** está diseñado para ser la solución de registro preferida para desarrolladores Go que valoran tanto el rendimiento como la simplicidad. Ya sea que esté construyendo una pequeña utilidad o un sistema distribuido a gran escala, esta biblioteca proporciona un excelente equilibrio entre funcionalidad y facilidad de uso.
