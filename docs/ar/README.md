---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

مكتبة سجلات Go عالية الأداء وقابلة للتخصيص، بنيت على zap، وتوفر ميزات غنية وواجهة برمجة تطبيقات بسيطة.

## 📖 لغات التوثيق

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 الصينية المبسطة](README.md) (الحالي)
-   [🇹🇼 الصينية التقليدية](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇸🇦 العربية](README.md) (الحالي)

## ✨ الميزات

-   **🚀 أداء عالي**：مبني على zap مع تجميع الكائنات وتسجيل الحقول الشرطية
-   **📊 مستويات سجلات غنية**：مستويات Trace, Debug, Info, Warn, Error, Fatal, Panic
-   **⚙️ تكوين مرن**：
    -   التحكم في مستوى السجل
    -   تسجيل معلومات المتصل
    -   معلومات التتبع (بما في ذلك معرف goroutine)
    -   بادئات وسلاسل سجلات مخصصة
    -   أهداف إخراج مخصصة (الوحدة الطرفية، الملفات، إلخ)
    -   خيارات تنسيق السجلات
-   **🔄 تدوير الملفات**：دعم تدوير ملفات السجلات كل ساعة
-   **🔌 التوافق مع Zap**：تكامل سلس مع zap WriteSyncer
-   **🎯 API بسيط**：واجهة واضحة تشبه مكتبة السجلات القياسية، سهلة الاستخدام

## 🚀 البدء السريع

### التثبيت

:::tip التثبيت
```bash
go get github.com/lazygophers/log
```
:::

### الاستخدام الأساسي

```go title="البداية السريعة"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // استخدام السجل العام الافتراضي
    log.Debug("رسالة تصحيح")
    log.Info("رسالة معلومات")
    log.Warn("رسالة تحذير")
    log.Error("رسالة خطأ")

    // استخدام الإخراج المنسق
    log.Infof("تم تسجيل الدخول بنجاح للمستخدم %s", "admin")

    // التكوين المخصص
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[تطبيقي]")

    customLogger.Info("هذا سجل من سجل مخصص")
}
```

## 📚 الاستخدام المتقدم

### السجل المخصص مع الإخراج إلى ملف

```go title="تكوين الإخراج إلى ملف"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // إنشاء سجل مع إخراج إلى ملف
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("سجل تصحيح مع معلومات المتصل")
    logger.Info("سجل معلومات مع معلومات التتبع")
}
```

### التحكم في مستوى السجل

```go title="التحكم في مستوى السجل"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // سيتم تسجيل المستويات فقط warn وما فوق
    logger.Debug("هذا لن يتم تسجيله")  // تم تجاهله
    logger.Info("هذا لن يتم تسجيله")   // تم تجاهله
    logger.Warn("هذا سيتم تسجيله")    // تم تسجيله
    logger.Error("هذا سيتم تسجيله")   // تم تسجيله
}
```

## 🎯 سيناريوهات الاستخدام

### السيناريوهات القابلة للتطبيق

-   **خدمات الويب وخوادم واجهة برمجة التطبيقات**：تتبع الطلبات، السجلات الهيكلية، مراقبة الأداء
-   **بنية الخدمات المصغرة**：التتبع الموزع (TraceID)، تنسيق سجلات موحد، معدل نقل عالي
-   **أدوات سطر الأوامر**：التحكم في المستويات، إخراج نظيف، تقارير الأخطاء
-   **المهام الدفعية**：تدوير الملفات، وقت تشغيل طويل، تحسين الموارد

### المزايا الخاصة

-   **التحسين مع تجميع الكائنات**：إعادة استخدام كائنات Entry و Buffer، تقليل ضغط GC
-   **الكتابة غير المتزامنة**：معدل نقل عالي (10000+ سجل/ثانية) بدون تزامن
-   **دعم TraceID**：تتبع الطلبات في الأنظمة الموزعة، التكامل مع OpenTelemetry
-   **بدء بدون تكوين**：جاهز للاستخدام، تكوين تدريجي

## 🔧 خيارات التكوين

:::note خيارات التكوين
جميع الطرق التالية تدعم الاستدعاء السلس ويمكن دمجها لبناء سجل مخصص.
:::

### تكوين السجل

| الطريقة                  | الوصف                | القيمة الافتراضية       |
| --------------------- | ------------------- | ----------- |
| `SetLevel(level)`       | تعيين مستوى السجل الأدنى     | `DebugLevel` |
| `EnableCaller(enable)`  | تمكين/تعطيل معلومات المتصل  | `false`      |
| `EnableTrace(enable)`   | تمكين/تعطيل معلومات التتبع    | `false`      |
| `SetCallerDepth(depth)` | تعيين عمق المتصل       | `2`          |
| `SetPrefixMsg(prefix)`  | تعيين بادئة السجل         | `""`         |
| `SetSuffixMsg(suffix)`  | تعيين لاحقة السجل         | `""`         |
| `SetOutput(writers...)` | تعيين أهداف الإخراج         | `os.Stdout`  |

### مستويات السجل

| المستوى        | الوصف                        |
| ----------- | --------------------------- |
| `TraceLevel` | الأكثر تفصيلاً، للتبعية التفصيلية        |
| `DebugLevel` | معلومات التصحيح                    |
| `InfoLevel`  | معلومات عامة                    |
| `WarnLevel`  | رسائل التحذير                    |
| `ErrorLevel` | رسائل الأخطاء                    |
| `FatalLevel` | أخطاء فادحة (تستدعي os.Exit(1))    |
| `PanicLevel` | أخطاء الوهن (تستدعي panic())    |

## 🏗️ البنية المعمارية

### المكونات الأساسية

-   **Logger**：الهيكل الرئيسية للسجلات مع خيارات قابلة للتخصيص
-   **Entry**：سجل فردي مع دعم شامل للبيانات الوصفية
-   **Level**：تعريفات مستويات السجلات والوظائف المساعدة
-   **Format**：واجهة تنسيق السجلات والتنفيذ

### تحسين الأداء

-   **مجموعة الكائنات**：إعادة استخدام كائنات Entry لتقليل تخصيص الذاكرة
-   **التسجيل الشرطي**：تسجيل الحقول المكلفة فقط عند الحاجة
-   **التحقق السريع من المستويات**：التحقق من مستوى السجل في الطبقة الخارجية
-   **التصميم غير المتزامن**：معظم العمليات لا تتطلب تزامن

## 📊 مقارنة الأداء

:::info مقارنة الأداء
البيانات التالية مستندة إلى الاختبارات المقارنة؛ الأداء الفعلي قد يختلف حسب البيئة والتكوين.
:::

| الميزة          | lazygophers/log | zap    | logrus | سجل قياسي |
| ------------- | --------------- | ------ | ------ | -------- |
| الأداء      | عالي              | عالي     | متوسط     | منخفض       |
| بساطة API    | عالي              | متوسط     | عالي     | عالي       |
| غنى الميزات    | متوسط          | عالي     | عالي     | منخفض       |
| المرونة      | متوسط          | عالي     | عالي     | منخفض       |
| منحنى التعلم      | منخفض              | متوسط     | متوسط     | منخفض       |

## ❓ الأسئلة الشائعة

### كيف اختيار مستوى السجل المناسب؟

-   **بيئة التطوير**：استخدم `DebugLevel` أو `TraceLevel` للحصول على معلومات مفصلة
-   **بيئة الإنتاج**：استخدم `InfoLevel` أو `WarnLevel` لتقليل التكلفة
-   **اختبار الأداء**：استخدم `PanicLevel` لتعطيل جميع السجلات

### كيف تحسن الأداء في بيئة الإنتاج؟

:::warning ملاحظة
في السيناريوهات عالية الإنتاجية، يُنصح باستخدام الكتابة غير المتزامنة مع مستويات سجلات معقولة لتحسين الأداء.
:::

1. استخدم `AsyncWriter` للكتابة غير المتزامنة：

```go title="تكوين الكتابة غير المتزامنة"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. ضبط مستوى السجل لتجنب السجلات غير الضرورية：

```go title="تحسين المستوى"
logger.SetLevel(log.InfoLevel)  // تخطي Debug و Trace
```

3. استخدم السجلات الشرطية لتقليل التكلفة：

```go title="السجلات الشرطية"
if logger.Level >= log.DebugLevel {
    logger.Debug("معلومات تصحيح تفصيلية")
}
```

### ما الفرق بين `Caller` و `EnableCaller`؟

-   **`EnableCaller(enable bool)`**：يتحكم في ما إذا كان السجل يجمع معلومات المتصل
    -   `EnableCaller(true)` يتيج تجميع معلومات المتصل
-   **`Caller(disable bool)`**：يتحكم في ما إذا كان التنسيق يعرض معلومات المتصل
    -   `Caller(true)` يتعطي إخراج معلومات المتصل

يُنصح باستخدام `EnableCaller` للتحكم العام.

### كيف تنفذ تنسيقاً مخصصاً؟

قم بتنفيذ واجهة `Format`：

```go title="التنسيق المخصص"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 الوثائق ذات الصلة

-   [📚 وثائق API](API.md) - مرجع كامل لواجهة برمجة التطبيقات
-   [🤝 دليل المساهمة](/ar/CONTRIBUTING) - كيف تساهم
-   [📋 سجل التغييرات](/ar/CHANGELOG) - سجل التاريخ
-   [🔒 سياسة الأمان](/ar/SECURITY) - دليل الأمان
-   [📜 قانون السلوك](/ar/CODE_OF_CONDUCT) - معايير المجتمع

## 🚀 الحصول على المساعدة

-   **GitHub Issues**：[إبلاغ عن خطأ أو طلب ميزة](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[وثائق API](https://pkg.go.dev/github.com/lazygophers/log)
-   [✓ أمثلة](https://github.com/lazygophers/log/tree/main/examples)

## 📄 الترخيص

هذا المشروع مرخص تحت رخصة MIT - راجع ملف [LICENSE](/ar/LICENSE) للتفاصيل.

## 🤝 المساهمة

نرحب بالمساهمات! يرجى مراجعة [دليل المساهمة](/ar/CONTRIBUTING) للمزيد من المعلومات.

---

**lazygophers/log** مصمم ليكون الحل الأول للسجلات لمطوري Go الذين يقدرون الأداء والبساطة. سواء كنت تبني أداة صغيرة أو نظام موزع واسع النطاق، توفر هذه المكتبة توازناً ممتازاً بين الوظائف والسهولة في الاستخدام.
