---
titleSuffix: " | LazyGophers Log"
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

مكتبة تسجيل Go عالية الأداء والمرنة، مبنية على zap، توفر ميزات غنية وواجهة برمجة تطبيقات بسيطة.

## 📖 لغات الوثائق

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇸🇦 العربية](README.md) (الحالي)

## 🚀 التوثيق الإلكتروني

زور [توثيق GitHub Pages](https://lazygophers.github.io/log/) لتحصل على تجربة قراءة أفضل.

## ✨ الميزات

-   **🚀 أداء عالي**: مبنية على zap مع إعادة استخدام كائنات Entry عبر مجموعة، مما يقلل من تخصيص الذاكرة
-   **📊 مستويات تسجيل غنية**: مستويات Trace، Debug، Info، Warn، Error، Fatal، Panic
-   **⚙️ تكوين مرن**:
    -   التحكم بمستوى التسجيل
    -   تسجيل معلومات المتصل
    -   معلومات التتبع (بما في ذلك معرف goroutine)
    -   بادئات وخلاصات تسجيل مخصصة
    -   أهداف إخراج مخصصة (الوحدة التحكم، الملفات، إلخ)
    -   خيارات تنسيق التسجيل
-   **🔄 دورة الملفات**: دعم لدورة الملفات التسجيلية كل ساعة
-   **🔌 توافق مع Zap**: تكامل سلس مع zap WriteSyncer
-   **🎯 واجهة برمجة تطبيقات بسيطة**: واجهة برمجة تطبيقات نظيفة تشبه مكتبة التسجيل القياسية، سهلة الاستخدام

## 🚀 بداية سريعة

### التثبيت

:::tip التثبيت
```bash
go get github.com/lazygophers/log
```
:::

### الاستخدام الأساسي

```go title="بداية سريعة"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // استخدام المسجل العام الافتراضي
    log.Debug("معلومات تصحيح")
    log.Info("معلومات عادية")
    log.Warn("تحذير")
    log.Error("خطأ")

    // استخدام المخرجات المنسقة
    log.Infof("المستخدم %s سجل الدخول بنجاح", "admin")

    // تكوين مخصص
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[تطبيقي]")

    customLogger.Info("هذا تسجيل من المسجل المخصص")
}
```

## 📚 استخدام متقدم

### مسجل مخصص مع مخرجات ملف

```go title="تكوين مخرجات الملف"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // إنشاء مسجل مع مخرجات ملف
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("تسجيل تصحيح مع معلومات المتصل")
    logger.Info("تسجيل عادي مع معلومات التتبع")
}
```

### التحكم بمستويات التسجيل

```go title="التحكم بمستويات التسجيل"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // سيتم تسجيل مستويات warn وما فوق فقط
    logger.Debug("لن يتم تسجيل هذا")  // سيتم تجاهله
    logger.Info("لن يتم تسجيل هذا")   // سيتم تجاهله
    logger.Warn("سيتم تسجيل هذا")    // سيتم تسجيله
    logger.Error("سيتم تسجيل هذا")   // سيتم تسجيله
}
```

## � حالات الاستخدام

### حالات الاستخدم المناسبة

-   **خدمات الويب وواجهات برمجة التطبيقات الخلفية**: تتبع الطلبات، تسجيل هيكلي، مراقبة الأداء
-   **هندسة الخدمات الدقيقة**: تتبع موزع (TraceID)، تنسيق تسجيل موحد، إنتاجية عالية
-   **أدوات سطر الأوامر**: تحكم المستويات، مخرجات موجزة، تقارير الأخطاء
-   **مهام المعالجة الدفعية**: تدوير الملفات، تشغيل لفترة طويلة، تحسين الموارد

### ميزات خاصة

-   **تحسين مجموعة الكائنات**: إعادة استخدام كائنات Entry و Buffer، تقليل ضغط garbage collection
-   **الكتابة غير المتزامنة**: سيناريوهات إنتاجية عالية (10000+ سجل/ثانية) بدون حجب
-   **دعم TraceID**: تتبع طلبات الأنظمة الموزعة، تكامل مع OpenTelemetry
-   **بدون تكوين**: جاهز للاستخدام، تكوين تدريجي

## 🔧 خيارات التكوين

:::note خيارات التكوين
جميع الطرق تدعم السلسلة، يمكن دمجها لبناء مسجل مخصص.
:::

### تكوين المسجل

| الطريقة                  | الوصف                         | القيمة الافتراضية |
| ------------------------- | ----------------------------- | ----------------- |
| `SetLevel(level)`         | تعيين الحد الأدنى لمستوى التسجيل | `DebugLevel`      |
| `EnableCaller(enable)`    | تفعيل/تعطيل تسجيل معلومات المتصل | `false`           |
| `EnableTrace(enable)`     | تفعيل/تعطيل معلومات التتبع   | `false`           |
| `SetCallerDepth(depth)`   | تعيين عمق المتصل             | `2`               |
| `SetPrefixMsg(prefix)`    | تعيين بادئة التسجيل          | `""`              |
| `SetSuffixMsg(suffix)`    | تعيين خاتمة التسجيل          | `""`              |
| `SetOutput(writers...)`   | تعيين أهداف المخرجات         | `os.Stdout`       |

### مستويات التسجيل

| المستوى      | الوصف                        |
| ------------ | ---------------------------- |
| `TraceLevel` | الأكثر تفصيلاً، للتتبع التفصيلي |
| `DebugLevel` | معلومات التصحيح              |
| `InfoLevel`  | رسائل معلومات               |
| `WarnLevel`  | شروط تحذير                   |
| `ErrorLevel` | شروط خطأ                     |
| `FatalLevel` | أخطاء قاتلة (تستدعي os.Exit(1)) |
| `PanicLevel` | أخطاء ذعر (تستدعي panic())   |

## 🏗️ هندسة

### المكونات الرئيسية

-   **Logger**: هيكل تسجيل رئيسي ذو مستوى ومخرجات ومنسق وعمق متصل قابل للتكوين
-   **Entry**: تسجيل فردي بدعم كامل للبيانات الوصفية
-   **Level**: تعريفات مستوى التسجيل والدالات المساعدة
-   **Format**: واجهة وتنفيذ تنسيق التسجيل

### تحسينات الأداء

-   **مجموعة الكائنات**: إعادة استخدام كائنات Entry لتقليل تخصيص الذاكرة
-   **التسجيل الشرطي**: تسجيل الحقول المكلفة فقط عند الضرورة
-   **فحص سريع للمستويات**: فحص المستويات في المستوى الخارجي
-   **تصميم بدون قفل**: معظم العمليات لا تحتاج إلى قفل

## 📊 مقارنة الأداء

:::info مقارنة الأداء
البيانات التالية تعتمد على المقاييس، الأداء الفعلي قد يختلف حسب البيئة والتكوين.
:::

| الميزة          | lazygophers/log | zap    | logrus | التسجيل القياسي |
| --------------- | --------------- | ------ | ------ | -------------- |
| الأداء         | عالي           | عالي  | متوسط | منخفض         |
| بساطة API      | عالي           | متوسط | عالي  | عالي          |
| غنى الميزات   | متوسط          | عالي  | عالي  | منخفض         |
| المرونة        | متوسط          | عالي  | عالي  | منخفض         |
| منحنى التعلم   | منخفض          | متوسط | متوسط | منخفض         |

## ❓ أسئلة شائعة

### كيف أختار مستوى تسجيل مناسب؟

-   **بيئة التطوير**: استخدم `DebugLevel` أو `TraceLevel` للحصول على معلومات مفصلة
-   **بيئة الإنتاج**: استخدم `InfoLevel` أو `WarnLevel` لتقليل الحمل
-   **اختبارات الأداء**: استخدم `PanicLevel` لتعطيل جميع التسجيلات

### كيف أحسن الأداء في الإنتاج؟

:::تحذير
في سيناريوهات الإنتاجية العالية، يُنصح بالدمج بين الكتابة غير المتزامنة ومستويات التسجيل المناسبة لتحسين الأداء.
:::

1. استخدام `AsyncWriter` للكتابة غير المتزامنة:

```go title="تكوين الكتابة غير المتزامنة"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. ضبط مستويات التسجيل لتجنب التسجيلات غير الضرورية:

```go title="تحسين المستويات"
logger.SetLevel(log.InfoLevel)  // تخطي Debug و Trace
```

3. استخدام تسجيلات شرطية لتقليل الحمل:

```go title="تسجيلات شرطية"
if logger.Level >= log.DebugLevel {
    logger.Debug("معلومات تصحيح مفصلة")
}
```

### ما الفرق بين `Caller` و `EnableCaller`؟

-   **`EnableCaller(enable bool)`**: يتحكم ما إذا كان المسجل يجمع معلومات المتصل
    -   `EnableCaller(true)` يفعيل جمع معلومات المتصل
-   **`Caller(disable bool)`**: يتحكم ما إذا كان المنسق يعرض معلومات المتصل
    -   `Caller(true)` يعطل عرض معلومات المتصل

يُنصح باستخدام `EnableCaller` للتحكم العام.

### كيف أنفذ منسقًا مخصصًا؟

نفذ واجهة `Format`:

```go title="منسق مخصص"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔌 وثائق ذات صلة

-   [📚 وثائق API](API.md) - مرجع API كامل
-   [🤝 د المساهمة](CONTRIBUTING.md) - كيفية المساهمة
-   [📋 سجل التغييرات](CHANGELOG.md) - تاريخ الإصدارات
-   [🔒 سياسة الأمان](SECURITY.md) - دليل الأمان
-   [📜 مدونة السلوك](CODE_OF_CONDUCT.md) - معايير المجتمع

## 🚀 الحصول على مساعدة

-   **GitHub Issues**: [الإبلاغ عن خطأ أو طلب ميزة](https://github.com/lazygophers/log/issues)
-   **GoDoc**: [وثائق API](https://pkg.go.dev/github.com/lazygophers/log)
-   **أمثلة**: [أمثلة الاستخدام](https://github.com/lazygophers/log/tree/main/examples)

## 📄 الترخيص

هذا المشروع مرخص بموجب ترخيص MIT - انظر ملف [LICENSE] للتفاصيل.

## 🤝 المساهمة

نحن نرحب بالمساهمات! يرجى مراجعة [دليل المساهمة](CONTRIBUTING.md) للمزيد من التفاصيل.

---

**lazygophers/log** يهدف إلى أن يكون حل تسجيل المفضل لمطوري Go الذين يقدرون الأداء والبساطة. سواء كنت تبني أدوات صغيرة أو أنظمة موزعة كبيرة، توفر هذه المكتبة توازنًا مثاليًا بين الوظيفة وسهولة الاستخدام.