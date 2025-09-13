# 🚀 LazyGophers Log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![DeepWiki](https://img.shields.io/badge/DeepWiki-documented-blue?logo=bookstack&logoColor=white)](https://deepwiki.ai/docs/lazygophers/log)
[![Go.Dev Downloads](https://pkg.go.dev/badge/github.com/lazygophers/log.svg)](https://pkg.go.dev/github.com/lazygophers/log)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.cn/stats/github.com/lazygophers/log)
[![Goproxy.io](https://goproxy.io/stats/github.com/lazygophers/log/badges/download-count.svg)](https://goproxy.io/stats/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

مكتبة تسجيل عالية الأداء وغنية بالميزات لتطبيقات Go مع دعم العلامات المتعددة للبناء والكتابة غير المتزامنة وخيارات التخصيص الواسعة.

## 📖 لغات التوثيق

- [🇺🇸 English](../README.md)
- [🇨🇳 简体中文](README.zh-CN.md)
- [🇹🇼 繁體中文](README.zh-TW.md)
- [🇫🇷 Français](README.fr.md)
- [🇷🇺 Русский](README.ru.md)
- [🇪🇸 Español](README.es.md)
- [🇸🇦 العربية](README.ar.md) (الحالي)

## ✨ الميزات

- **🚀 أداء عالي**: دعم تجميع الكائنات والكتابة غير المتزامنة
- **🏗️ دعم علامات البناء**: سلوكيات مختلفة لأوضاع التصحيح والإصدار والتجاهل
- **🔄 دوران السجلات**: دوران تلقائي لملفات السجلات كل ساعة
- **🎨 تنسيق غني**: تنسيقات سجلات قابلة للتخصيص مع دعم الألوان
- **🔍 التتبع السياقي**: تتبع معرف Goroutine ومعرف التتبع
- **🔌 تكامل الإطارات**: تكامل أصلي مع مسجل Zap
- **⚙️ قابل للتكوين بدرجة عالية**: مستويات وإخراجات وتنسيقات مرنة
- **🧪 مُختبر جيداً**: تغطية اختبار 93.0% عبر جميع تكوينات البناء

## 🚀 البداية السريعة

### التثبيت

```bash
go get github.com/lazygophers/log
```

### الاستخدام الأساسي

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // تسجيل بسيط
    log.Info("مرحباً، العالم!")
    log.Debug("هذه رسالة تصحيح")
    log.Warn("هذا تحذير")
    log.Error("هذا خطأ")

    // تسجيل منسق
    log.Infof("المستخدم %s سجل الدخول بالمعرف %d", "أحمد", 123)
    
    // مع مسجل مخصص
    logger := log.New()
    logger.SetLevel(log.InfoLevel)
    logger.Info("رسالة المسجل المخصص")
}
```

### الاستخدام المتقدم

```go
package main

import (
    "context"
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // إنشاء مسجل مع إخراج للملف
    logger := log.New()
    
    // تعيين الإخراج لملف مع دوران كل ساعة
    writer := log.GetOutputWriterHourly("./logs/app.log")
    logger.SetOutput(writer)
    
    // تكوين التنسيق
    logger.SetLevel(log.DebugLevel)
    logger.SetPrefixMsg("[APP] ")
    logger.Caller(true) // تمكين معلومات المستدعي
    
    // التسجيل السياقي
    ctxLogger := logger.CloneToCtx()
    ctxLogger.Info(context.Background(), "تسجيل واعٍ بالسياق")
    
    // التسجيل غير المتزامن للسيناريوهات عالية الإنتاجية
    asyncWriter := log.NewAsyncWriter(writer, 1000)
    logger.SetOutput(asyncWriter)
    defer asyncWriter.Close()
    
    logger.Info("تسجيل غير متزامن عالي الأداء")
}
```

## 🏗️ علامات البناء

تدعم المكتبة أوضاع بناء مختلفة من خلال علامات بناء Go:

### الوضع الافتراضي (بدون علامات)
```bash
go build
```
- وظائف التسجيل الكاملة
- رسائل التصحيح مفعلة
- أداء قياسي

### وضع التصحيح
```bash
go build -tags debug
```
- معلومات تصحيح محسنة
- معلومات مفصلة عن المستدعي
- دعم تحليل الأداء

### وضع الإصدار
```bash
go build -tags release
```
- محسن للإنتاج
- رسائل التصحيح معطلة
- دوران تلقائي لملفات السجلات

### وضع التجاهل
```bash
go build -tags discard
```
- أقصى أداء
- جميع السجلات يتم تجاهلها
- صفر أعباء تسجيل

### الأوضاع المدمجة
```bash
go build -tags "debug,discard"    # التصحيح مع التجاهل
go build -tags "release,discard"  # الإصدار مع التجاهل
```

## 📊 مستويات السجل

تدعم المكتبة 7 مستويات للسجل (من الأعلى إلى الأدنى أولوية):

| المستوى | القيمة | الوصف |
|---------|---------|--------|
| `PanicLevel` | 0 | يسجل ثم يستدعي panic |
| `FatalLevel` | 1 | يسجل ثم يستدعي os.Exit(1) |
| `ErrorLevel` | 2 | حالات الأخطاء |
| `WarnLevel` | 3 | حالات التحذير |
| `InfoLevel` | 4 | الرسائل الإعلامية |
| `DebugLevel` | 5 | رسائل مستوى التصحيح |
| `TraceLevel` | 6 | أكثر التسجيلات تفصيلاً |

## 🔌 تكامل الإطارات

### تكامل Zap

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "github.com/lazygophers/log"
)

// إنشاء مسجل zap يكتب إلى نظام السجلات الخاص بنا
logger := log.New()
hook := log.NewZapHook(logger)

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    hook,
    zapcore.InfoLevel,
)
zapLogger := zap.New(core)

zapLogger.Info("رسالة من Zap", zap.String("key", "value"))
```

## 🧪 الاختبار

تأتي المكتبة مع دعم اختبار شامل:

```bash
# تشغيل جميع الاختبارات
make test

# تشغيل الاختبارات مع التغطية لجميع علامات البناء
make coverage-all

# اختبار سريع عبر جميع علامات البناء
make test-quick

# توليد تقارير تغطية HTML
make coverage-html
```

### نتائج التغطية حسب علامة البناء

| علامة البناء | التغطية |
|--------------|----------|
| الافتراضي | 92.9% |
| التصحيح | 93.1% |
| الإصدار | 93.5% |
| التجاهل | 93.1% |
| التصحيح+التجاهل | 93.1% |
| الإصدار+التجاهل | 93.3% |

## ⚙️ خيارات التكوين

### تكوين المسجل

```go
logger := log.New()

// تعيين الحد الأدنى لمستوى السجل
logger.SetLevel(log.InfoLevel)

// تكوين الإخراج
logger.SetOutput(os.Stdout) // كاتب واحد
logger.SetOutput(writer1, writer2, writer3) // عدة كتاب

// تخصيص الرسائل
logger.SetPrefixMsg("[تطبيقي] ")
logger.SetSuffixMsg(" [النهاية]")
logger.AppendPrefixMsg("إضافي: ")

// تكوين التنسيق
logger.ParsingAndEscaping(false) // تعطيل تسلسلات الإفلات
logger.Caller(true) // تمكين معلومات المستدعي
logger.SetCallerDepth(4) // ضبط عمق مكدس المستدعي
```

## 📁 دوران السجلات

دوران تلقائي للسجلات مع فترات قابلة للتكوين:

```go
// دوران كل ساعة
writer := log.GetOutputWriterHourly("./logs/app.log")

// ستنشئ المكتبة ملفات مثل:
// - app-2024010115.log (2024-01-01 15:00)
// - app-2024010116.log (2024-01-01 16:00)
// - app-2024010117.log (2024-01-01 17:00)
```

## 🔍 السياق والتتبع

دعم مدمج للتسجيل الواعي بالسياق والتتبع الموزع:

```go
// تعيين معرف التتبع للـ goroutine الحالي
log.SetTrace("trace-123-456")

// الحصول على معرف التتبع
traceID := log.GetTrace()

// التسجيل الواعي بالسياق
ctx := context.Background()
ctxLogger := log.CloneToCtx()
ctxLogger.Info(ctx, "تم معالجة الطلب", "user_id", 123)

// تتبع معرف goroutine التلقائي
log.Info("هذا السجل يتضمن تلقائياً معرف goroutine")
```

## 📈 الأداء

تم تصميم المكتبة للتطبيقات عالية الأداء:

- **تجميع الكائنات**: إعادة استخدام كائنات إدخال السجل لتقليل ضغط GC
- **الكتابة غير المتزامنة**: كتابات سجل غير محجوبة للسيناريوهات عالية الإنتاجية
- **ترشيح المستوى**: الترشيح المبكر يمنع العمليات المكلفة
- **تحسين علامات البناء**: تحسين وقت التجميع لبيئات مختلفة

### المعايير المرجعية

```bash
# تشغيل معايير الأداء المرجعية
make benchmark

# معايير مرجعية لأوضاع البناء المختلفة
make benchmark-debug
make benchmark-release  
make benchmark-discard
```

## 🤝 المساهمة

نرحب بالمساهمات! يرجى مراجعة [دليل المساهمة](CONTRIBUTING.md) للتفاصيل.

### إعداد بيئة التطوير

1. **Fork و Clone**
   ```bash
   git clone https://github.com/your-username/log.git
   cd log
   ```

2. **تثبيت التبعيات**
   ```bash
   go mod tidy
   ```

3. **تشغيل الاختبارات**
   ```bash
   make test-all
   ```

4. **إرسال Pull Request**
   - اتبع [قالب PR](../.github/pull_request_template.md) الخاص بنا
   - تأكد من نجاح الاختبارات
   - حدث التوثيق إذا لزم الأمر

## 📋 المتطلبات

- **Go**: 1.19 أو أعلى
- **التبعيات**: 
  - `go.uber.org/zap` (لتكامل Zap)
  - `github.com/petermattis/goid` (لمعرف goroutine)
  - `github.com/lestrrat-go/file-rotatelogs` (لدوران السجلات)
  - `github.com/google/uuid` (لمعرفات التتبع)

## 📄 الترخيص

هذا المشروع مرخص تحت رخصة MIT - راجع ملف [LICENSE](../LICENSE) للتفاصيل.

## 🙏 شكر وتقدير

- [Zap](https://github.com/uber-go/zap) للإلهام ودعم التكامل
- [Logrus](https://github.com/sirupsen/logrus) لأنماط تصميم المستويات
- مجتمع Go للتغذية الراجعة المستمرة والتحسينات

## 📞 الدعم

- 📖 [التوثيق](../docs/)
- 🐛 [متتبع المشاكل](https://github.com/lazygophers/log/issues)
- 💬 [المناقشات](https://github.com/lazygophers/log/discussions)
- 📧 البريد الإلكتروني: support@lazygophers.com

---

**صُنع بـ ❤️ من قبل فريق LazyGophers**