# 🤝 Contribuir a LazyGophers Log

¡Amamos tu contribución! Queremos hacer que contribuir a LazyGophers Log sea tan fácil y transparente como sea posible, ya sea:

- 🐛 Reportar un error
- 💬 Discutir el estado actual del código
- ✨ Enviar una solicitud de característica
- 🔧 Proponer una corrección
- 🚀 Implementar una nueva característica

## 📋 Tabla de Contenidos

- [Código de Conducta](#-código-de-conducta)
- [Proceso de Desarrollo](#-proceso-de-desarrollo)
- [Inicio Rápido](#-inicio-rápido)
- [Proceso de Pull Request](#-proceso-de-pull-request)
- [Estándares de Codificación](#-estándares-de-codificación)
- [Guías de Pruebas](#-guías-de-pruebas)
- [Requisitos de Etiquetas de Construcción](#️-requisitos-de-etiquetas-de-construcción)
- [Documentación](#-documentación)
- [Guías para Issues](#-guías-para-issues)
- [Consideraciones de Rendimiento](#-consideraciones-de-rendimiento)
- [Guías de Seguridad](#-guías-de-seguridad)
- [Comunidad](#-comunidad)

## 📜 Código de Conducta

Este proyecto y todos los que participan en él se rigen por nuestro [Código de Conducta](CODE_OF_CONDUCT_es.md). Al participar, se espera que mantengas este código.

## 🔄 Proceso de Desarrollo

Usamos GitHub para alojar el código, rastrear issues y solicitudes de características, así como aceptar pull requests.

### Flujo de Trabajo

1. **Fork** el repositorio
2. **Clonar** tu fork localmente
3. **Crear** una rama de característica desde `master`
4. **Hacer** tus cambios
5. **Probar** exhaustivamente en todas las etiquetas de construcción
6. **Enviar** un pull request

## 🚀 Inicio Rápido

### Prerrequisitos

- **Go 1.21+** - [Instalar Go](https://golang.org/doc/install)
- **Git** - [Instalar Git](https://git-scm.com/book/es/v2/Inicio---Sobre-el-Control-de-Versiones-Instalación-de-Git)
- **Make** (opcional pero recomendado)

### Configuración del Entorno de Desarrollo Local

```bash
# 1. Haz fork del repositorio en GitHub
# 2. Clona tu fork
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. Agrega el remoto upstream
git remote add upstream https://github.com/lazygophers/log.git

# 4. Instala las dependencias
go mod tidy

# 5. Verifica la instalación
make test-quick
```

### Configuración del Entorno

```bash
# Configura tu entorno Go (si no está hecho ya)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Opcional: Instala herramientas útiles
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 Proceso de Pull Request

### Antes de Enviar

1. **Buscar** PRs existentes para evitar duplicados
2. **Probar** tus cambios en todas las configuraciones de construcción
3. **Documentar** cualquier cambio importante
4. **Actualizar** la documentación relevante
5. **Agregar** pruebas para nueva funcionalidad

### Lista de Verificación PR

- [ ] **Calidad del Código**
  - [ ] El código sigue las guías de estilo del proyecto
  - [ ] No hay nuevas advertencias de linting
  - [ ] Manejo de errores apropiado
  - [ ] Algoritmos y estructuras de datos eficientes

- [ ] **Pruebas**
  - [ ] Todas las pruebas existentes pasan: `make test`
  - [ ] Nuevas pruebas agregadas para nueva funcionalidad
  - [ ] Cobertura de pruebas mantenida o mejorada
  - [ ] Todas las etiquetas de construcción probadas: `make test-all`

- [ ] **Documentación**
  - [ ] El código está apropiadamente comentado
  - [ ] Documentación API actualizada (si es necesario)
  - [ ] README actualizado (si es necesario)
  - [ ] Documentación multiidioma actualizada (si es orientada al usuario)

- [ ] **Compatibilidad de Construcción**
  - [ ] Modo por defecto: `go build`
  - [ ] Modo debug: `go build -tags debug`
  - [ ] Modo release: `go build -tags release`
  - [ ] Modo discard: `go build -tags discard`
  - [ ] Modos combinados probados

### Plantilla de PR

Por favor usa nuestra [plantilla de PR](.github/pull_request_template.md) al enviar pull requests.

## 📏 Estándares de Codificación

### Guía de Estilo Go

Seguimos la guía de estilo Go estándar con algunas adiciones:

```go
// ✅ Bueno
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ Malo
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### Convenciones de Nomenclatura

- **Paquetes**: Cortos, minúsculas, una palabra cuando sea posible
- **Funciones**: CamelCase, descriptivas
- **Variables**: camelCase para locales, CamelCase para exportadas
- **Constantes**: CamelCase para exportadas, camelCase para no exportadas
- **Interfaces**: Usualmente terminan en "er" (ej. `Writer`, `Formatter`)

### Organización del Código

```
project/
├── docs/           # Documentación en múltiples idiomas
├── .github/        # Plantillas GitHub y workflows
├── logger.go       # Implementación principal del logger
├── entry.go        # Estructura de entrada de log
├── level.go        # Niveles de log
├── formatter.go    # Formateo de logs
├── output.go       # Gestión de salida
└── *_test.go      # Pruebas co-ubicadas con el código fuente
```

### Manejo de Errores

```go
// ✅ Preferido: Retornar errores, no hacer panic
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ Evitar: Hacer panic en código de biblioteca
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // No hagas esto
    }
    return &Logger{...}
}
```

## 🧪 Guías de Pruebas

### Estructura de Pruebas

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
            // Implementación de la prueba
        })
    }
}
```

### Requisitos de Cobertura

- **Mínimo**: 90% de cobertura para código nuevo
- **Objetivo**: 95%+ de cobertura global
- **Todas las etiquetas de construcción** deben mantener cobertura
- Usa `make coverage-all` para verificar

### Comandos de Pruebas

```bash
# Prueba rápida en todas las etiquetas de construcción
make test-quick

# Suite completa de pruebas con cobertura
make test-all

# Reportes de cobertura
make coverage-html

# Benchmarks
make benchmark
```

## 🏗️ Requisitos de Etiquetas de Construcción

Todos los cambios deben ser compatibles con nuestro sistema de etiquetas de construcción:

### Etiquetas de Construcción Soportadas

- **Por defecto** (`go build`): Funcionalidad completa
- **Debug** (`go build -tags debug`): Depuración mejorada
- **Release** (`go build -tags release`): Optimizado para producción
- **Discard** (`go build -tags discard`): Máximo rendimiento

### Pruebas de Etiquetas de Construcción

```bash
# Probar cada configuración de construcción
make test-default
make test-debug  
make test-release
make test-discard

# Probar combinaciones
make test-debug-discard
make test-release-discard

# Todo en uno
make test-all
```

### Guías de Etiquetas de Construcción

```go
//go:build debug
// +build debug

package log

// Implementaciones específicas de debug
```

## 📚 Documentación

### Documentación del Código

- **Todas las funciones exportadas** deben tener comentarios claros
- **Algoritmos complejos** necesitan explicación
- **Ejemplos** para uso no trivial
- **Notas de seguridad de hilos** donde sea aplicable

```go
// SetLevel establece el nivel mínimo de logging.
// Los logs debajo de este nivel serán ignorados.
// Este método es thread-safe.
//
// Ejemplo:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // No se mostrará
//   logger.Info("visible")   // Se mostrará
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### Actualizaciones del README

Al agregar características, actualiza:
- El README.md principal
- Todos los READMEs específicos de idioma en `docs/`
- Ejemplos de código
- Listas de características

## 🐛 Guías para Issues

### Reportes de Errores

Usa la [plantilla de reporte de error](.github/ISSUE_TEMPLATE/bug_report.md) e incluye:

- **Descripción clara** del problema
- **Pasos para reproducir**
- **Comportamiento esperado vs actual**
- **Detalles del entorno** (OS, versión de Go, etiquetas de construcción)
- **Muestra de código mínima**

### Solicitudes de Características

Usa la [plantilla de solicitud de característica](.github/ISSUE_TEMPLATE/feature_request.md) e incluye:

- **Motivación clara** para la característica
- **Diseño de API** propuesto
- **Consideraciones de implementación**
- **Análisis de cambios importantes**

### Preguntas

Usa la [plantilla de pregunta](.github/ISSUE_TEMPLATE/question.md) para:

- Preguntas de uso
- Ayuda de configuración
- Mejores prácticas
- Guía de integración

## 🚀 Consideraciones de Rendimiento

### Benchmarking

Siempre haz benchmark de cambios sensibles al rendimiento:

```bash
# Ejecutar benchmarks
go test -bench=. -benchmem

# Comparar antes/después
go test -bench=. -benchmem > before.txt
# Hacer cambios
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### Guías de Rendimiento

- **Minimizar asignaciones** en rutas críticas
- **Usar pools de objetos** para objetos creados frecuentemente
- **Retorno temprano** para niveles de log deshabilitados
- **Evitar reflexión** en código crítico de rendimiento
- **Perfilar antes de optimizar**

### Gestión de Memoria

```go
// ✅ Bueno: Usar pools de objetos
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

## 🔒 Guías de Seguridad

### Datos Sensibles

- **Nunca loguear** contraseñas, tokens o datos sensibles
- **Limpiar** entrada de usuario en mensajes de log
- **Evitar** loguear cuerpos completos de solicitud/respuesta
- **Usar** logging estructurado para mejor control

```go
// ✅ Bueno
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ Malo
logger.Infof("User login: %+v", userRequest) // Puede contener contraseña
```

### Dependencias

- Mantener dependencias **actualizadas**
- **Revisar** cuidadosamente nuevas dependencias
- **Minimizar** dependencias externas
- **Usar** `go mod verify` para verificar integridad

## 👥 Comunidad

### Obtener Ayuda

- 📖 [Documentación](../README_es.md)
- 💬 [Discusiones GitHub](https://github.com/lazygophers/log/discussions)
- 🐛 [Rastreador de Issues](https://github.com/lazygophers/log/issues)
- 📧 Email: support@lazygophers.com

### Guías de Comunicación

- Ser **respetuoso** e inclusivo
- **Buscar** antes de hacer preguntas
- **Proporcionar contexto** al pedir ayuda
- **Ayudar a otros** cuando puedas
- **Seguir** el [Código de Conducta](CODE_OF_CONDUCT_es.md)

## 🎯 Reconocimiento

Los contribuyentes son reconocidos de varias maneras:

- Sección de **contribuyentes del README**
- Menciones en **notas de lanzamiento**
- Gráficos de **contribuyentes de GitHub**
- Posts de **apreciación de la comunidad**

## 📝 Licencia

Al contribuir, aceptas que tus contribuciones serán licenciadas bajo la Licencia MIT.

---

## 🌍 Documentación Multiidioma

Este documento está disponible en múltiples idiomas:

- [🇺🇸 English](CONTRIBUTING.md)
- [🇨🇳 简体中文](CONTRIBUTING_zh-CN.md)
- [🇹🇼 繁體中文](CONTRIBUTING_zh-TW.md)
- [🇫🇷 Français](CONTRIBUTING_fr.md)
- [🇷🇺 Русский](CONTRIBUTING_ru.md)
- [🇪🇸 Español](CONTRIBUTING_es.md) (Actual)
- [🇸🇦 العربية](CONTRIBUTING_ar.md)

---

**¡Gracias por contribuir a LazyGophers Log! 🚀**