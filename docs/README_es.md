# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Una biblioteca de registro Go de alto rendimiento y flexible, construida sobre zap, que ofrece funciones ricas y una API simple.

## 📖 Idiomas de documentación

-   [🇺🇸 English](README.md)
-   [🇨🇳 简体中文](README_zh-CN.md)
-   [🇹🇼 繁體中文](README_zh-TW.md)
-   [🇫🇷 Français](README_fr.md)
-   [🇷🇺 Русский](README_ru.md)
-   [🇪🇸 Español](README_es.md)
-   [🇸🇦 العربية](README_ar.md)

## ✨ Características

-   **🚀 Alto rendimiento** : Construido sobre zap con reutilización de objetos Entry a través de un pool, reduciendo la asignación de memoria
-   **📊 Niveles de registro ricos** : Niveles Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ Configuración flexible** :
    -   Control de nivel de registro
    -   Registro de información del llamador
    -   Información de traza (incluyendo ID de goroutine)
    -   Prefijos y sufijos de registro personalizados
    -   Destinos de salida personalizados (consola, archivos, etc.)
    -   Opciones de formato de registro
-   **🔄 Rotación de archivos** : Soporte para rotación horaria de archivos de registro
-   **🔌 Compatibilidad con Zap** : Integración perfecta con zap WriteSyncer
-   **🎯 API simple** : API limpia similar a la biblioteca de registro estándar, fácil de usar

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
    // Usar el registrador global por defecto
    log.Debug("Mensaje de depuración")
    log.Info("Mensaje de información")
    log.Warn("Mensaje de advertencia")
    log.Error("Mensaje de error")

    // Usar salida formateada
    log.Infof("El usuario %s ha iniciado sesión correctamente", "admin")

    // Configuración personalizada
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Este es un registro del registrador personalizado")
}
```

## 📚 Uso avanzado

### Registrador personalizado con salida a archivo

```go
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Crear registrador con salida a archivo
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Mensaje de depuración con información del llamador")
    logger.Info("Mensaje de información con información de traza")
}
```

### Control de nivel de registro

```go
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Solo se registrarán los mensajes warn y superiores
    logger.Debug("Esto no se registrará")  // Ignorado
    logger.Info("Esto no se registrará")   // Ignorado
    logger.Warn("Esto se registrará")    // Registrado
    logger.Error("Esto se registrará")   // Registrado
}
```

## 🔧 Opciones de configuración

### Configuración del Logger

| Método                 | Descripción                  | Valor por defecto |
| ---------------------- | ---------------------------- | ---------------- |
| `SetLevel(level)`      | Establecer nivel mínimo de registro | `DebugLevel` |
| `EnableCaller(enable)` | Habilitar/deshabilitar información del llamador | `false` |
| `EnableTrace(enable)`  | Habilitar/deshabilitar información de traza | `false` |
| `SetCallerDepth(depth)` | Establecer profundidad del llamador | `2` |
| `SetPrefixMsg(prefix)` | Establecer prefijo de registro | `""` |
| `SetSuffixMsg(suffix)` | Establecer sufijo de registro | `""` |
| `SetOutput(writers...)` | Establecer destinos de salida | `os.Stdout` |

### Niveles de registro

| Nivel        | Descripción                        |
| ------------- | ---------------------------------- |
| `TraceLevel`  | El más detallado, para traza detallada |
| `DebugLevel`  | Información de depuración          |
| `InfoLevel`   | Información general                |
| `WarnLevel`   | Mensajes de advertencia            |
| `ErrorLevel`  | Mensajes de error                  |
| `FatalLevel`  | Errores fatales (llama a os.Exit(1)) |
| `PanicLevel`  | Errores de pánico (llama a panic()) |

## 🏗️ Arquitectura

### Componentes principales

-   **Logger** : Estructura principal de registro con niveles, salidas, formateadores y profundidad del llamador configurables
-   **Entry** : Registro individual de registro con soporte completo de metadatos
-   **Level** : Definiciones de niveles de registro y funciones auxiliares
-   **Format** : Interfaz e implementaciones de formato de registro

### Optimizaciones de rendimiento

-   **Pool de objetos** : Reutilización de objetos Entry para reducir la asignación de memoria
-   **Registro condicional** : Registro de campos costosos solo cuando sea necesario
-   **Verificación rápida de nivel** : Verificación del nivel de registro en la capa más externa
-   **Diseño sin bloqueo** : La mayoría de las operaciones no requieren bloqueo

## 📊 Comparación de rendimiento

| Característica         | lazygophers/log | zap    | logrus | registro estándar |
| ---------------------- | --------------- | ------ | ------ | ---------------- |
| Rendimiento            | Alto            | Alto   | Medio  | Bajo             |
| Simplicidad de API     | Alto            | Medio  | Alto   | Alto             |
| Riqueza de funciones   | Medio           | Alto   | Alto   | Bajo             |
| Flexibilidad           | Medio           | Alto   | Alto   | Bajo             |
| Curva de aprendizaje   | Baja            | Medio  | Medio  | Baja             |

## 🔗 Documentación relacionada

-   [📚 Documentación API](API.md) - Documentación de referencia API completa
-   [🤝 Guía de contribución](CONTRIBUTING.md) - Cómo contribuir
-   [📋 Registro de cambios](../CHANGELOG.md) - Historial de versiones
-   [🔒 Política de seguridad](SECURITY.md) - Directrices de seguridad
-   [📜 Código de conducta](CODE_OF_CONDUCT.md) - Directrices de la comunidad

## 🚀 Obtener ayuda

-   **GitHub Issues** : [Informar errores o solicitar funciones](https://github.com/lazygophers/log/issues)
-   **GoDoc** : [Documentación API](https://pkg.go.dev/github.com/lazygophers/log)
-   **Ejemplos** : [Ejemplos de uso](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Licencia

Este proyecto está licenciado bajo la Licencia MIT - consulte el archivo [LICENSE](../LICENSE) para más detalles.

## 🤝 Contribuir

¡Bienvenidas las contribuciones! Por favor, consulte nuestra [Guía de contribución](CONTRIBUTING.md) para más detalles.

---

**lazygophers/log** está diseñado para ser la solución de registro predeterminada para desarrolladores Go que valoran tanto el rendimiento como la simplicidad. Ya sea que esté construyendo una pequeña utilidad o un sistema distribuido de gran escala, esta biblioteca ofrece el equilibrio adecuado entre funcionalidad y facilidad de uso.