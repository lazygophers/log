# 🤝 المساهمة في LazyGophers Log

نحن نحب مساهماتك! نريد أن نجعل المساهمة في LazyGophers Log سهلة وشفافة قدر الإمكان، سواء كانت:

- 🐛 الإبلاغ عن خطأ
- 💬 مناقشة الحالة الحالية للكود
- ✨ تقديم طلب ميزة
- 🔧 اقتراح إصلاح
- 🚀 تنفيذ ميزة جديدة

## 📋 جدول المحتويات

- [قواعد السلوك](#-قواعد-السلوك)
- [عملية التطوير](#-عملية-التطوير)
- [البداية السريعة](#-البداية-السريعة)
- [عملية Pull Request](#-عملية-pull-request)
- [معايير البرمجة](#-معايير-البرمجة)
- [إرشادات الاختبار](#-إرشادات-الاختبار)
- [متطلبات علامات البناء](#️-متطلبات-علامات-البناء)
- [التوثيق](#-التوثيق)
- [إرشادات المشاكل](#-إرشادات-المشاكل)
- [اعتبارات الأداء](#-اعتبارات-الأداء)
- [إرشادات الأمان](#-إرشادات-الأمان)
- [المجتمع](#-المجتمع)

## 📜 قواعد السلوك

هذا المشروع وجميع المشاركين فيه محكومون بـ [قواعد السلوك](CODE_OF_CONDUCT_ar.md) الخاصة بنا. بالمشاركة، من المتوقع أن تلتزم بهذه القواعد.

## 🔄 عملية التطوير

نستخدم GitHub لاستضافة الكود وتتبع المشاكل وطلبات الميزات، بالإضافة إلى قبول pull requests.

### سير العمل

1. **Fork** المستودع
2. **استنسخ** الـ fork محلياً
3. **أنشئ** فرع ميزة من `master`
4. **قم** بالتغييرات
5. **اختبر** بدقة على جميع علامات البناء
6. **أرسل** pull request

## 🚀 البداية السريعة

### المتطلبات الأساسية

- **Go 1.21+** - [تثبيت Go](https://golang.org/doc/install)
- **Git** - [تثبيت Git](https://git-scm.com/book/ar/v2)
- **Make** (اختياري لكن موصى به)

### إعداد بيئة التطوير المحلية

```bash
# 1. قم بعمل fork للمستودع على GitHub
# 2. استنسخ fork الخاص بك
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. أضف remote upstream
git remote add upstream https://github.com/lazygophers/log.git

# 4. ثبت التبعيات
go mod tidy

# 5. تحقق من التثبيت
make test-quick
```

### إعداد البيئة

```bash
# قم بإعداد بيئة Go الخاصة بك (إذا لم تكن معدة بالفعل)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# اختياري: ثبت أدوات مفيدة
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 عملية Pull Request

### قبل التقديم

1. **ابحث** عن PR موجودة لتجنب التكرار
2. **اختبر** التغييرات على جميع تكوينات البناء
3. **وثق** أي تغييرات مهمة
4. **حدث** التوثيق ذي الصلة
5. **أضف** اختبارات للوظائف الجديدة

### قائمة التحقق من PR

- [ ] **جودة الكود**
  - [ ] يتبع الكود إرشادات أسلوب المشروع
  - [ ] لا توجد تحذيرات linting جديدة
  - [ ] معالجة أخطاء مناسبة
  - [ ] خوارزميات وهياكل بيانات فعالة

- [ ] **الاختبار**
  - [ ] جميع الاختبارات الموجودة تمر: `make test`
  - [ ] اختبارات جديدة مضافة للوظائف الجديدة
  - [ ] تغطية الاختبار محافظة أو محسنة
  - [ ] جميع علامات البناء مختبرة: `make test-all`

- [ ] **التوثيق**
  - [ ] الكود معلق بشكل مناسب
  - [ ] توثيق API محدث (عند الحاجة)
  - [ ] README محدث (عند الحاجة)
  - [ ] التوثيق متعدد اللغات محدث (إذا كان موجهاً للمستخدم)

- [ ] **توافق البناء**
  - [ ] الوضع الافتراضي: `go build`
  - [ ] وضع التصحيح: `go build -tags debug`
  - [ ] وضع الإصدار: `go build -tags release`
  - [ ] وضع التجاهل: `go build -tags discard`
  - [ ] الأوضاع المدمجة مختبرة

### قالب PR

يرجى استخدام [قالب PR](.github/pull_request_template.md) عند تقديم pull requests.

## 📏 معايير البرمجة

### دليل أسلوب Go

نتبع دليل أسلوب Go القياسي مع بعض الإضافات:

```go
// ✅ جيد
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ سيء
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### اتفاقيات التسمية

- **الحزم**: قصيرة، أحرف صغيرة، كلمة واحدة عند الإمكان
- **الدوال**: CamelCase، وصفية
- **المتغيرات**: camelCase للمحلية، CamelCase للمصدرة
- **الثوابت**: CamelCase للمصدرة، camelCase لغير المصدرة
- **الواجهات**: تنتهي عادة بـ "er" (مثل `Writer`، `Formatter`)

### تنظيم الكود

```
project/
├── docs/           # توثيق بلغات متعددة
├── .github/        # قوالب GitHub وسير العمل
├── logger.go       # تنفيذ logger الرئيسي
├── entry.go        # هيكل إدخال السجل
├── level.go        # مستويات السجل
├── formatter.go    # تنسيق السجلات
├── output.go       # إدارة الإخراج
└── *_test.go      # اختبارات موضوعة مع المصدر
```

### معالجة الأخطاء

```go
// ✅ مفضل: إرجاع الأخطاء، لا تستخدم panic
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ تجنب: panic في كود المكتبة
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // لا تفعل هذا
    }
    return &Logger{...}
}
```

## 🧪 إرشادات الاختبار

### هيكل الاختبار

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
            // تنفيذ الاختبار
        })
    }
}
```

### متطلبات التغطية

- **الحد الأدنى**: تغطية 90% للكود الجديد
- **الهدف**: تغطية 95%+ إجمالية
- **جميع علامات البناء** يجب أن تحافظ على التغطية
- استخدم `make coverage-all` للتحقق

### أوامر الاختبار

```bash
# اختبار سريع على جميع علامات البناء
make test-quick

# مجموعة اختبار كاملة مع التغطية
make test-all

# تقارير التغطية
make coverage-html

# المعايير القياسية
make benchmark
```

## 🏗️ متطلبات علامات البناء

جميع التغييرات يجب أن تكون متوافقة مع نظام علامات البناء الخاص بنا:

### علامات البناء المدعومة

- **الافتراضي** (`go build`): وظائف كاملة
- **التصحيح** (`go build -tags debug`): تصحيح محسن
- **الإصدار** (`go build -tags release`): محسن للإنتاج
- **التجاهل** (`go build -tags discard`): أداء أقصى

### اختبار علامات البناء

```bash
# اختبار كل تكوين بناء
make test-default
make test-debug  
make test-release
make test-discard

# اختبار التركيبات
make test-debug-discard
make test-release-discard

# الكل في واحد
make test-all
```

### إرشادات علامات البناء

```go
//go:build debug
// +build debug

package log

// تنفيذات خاصة بالتصحيح
```

## 📚 التوثيق

### توثيق الكود

- **جميع الدوال المصدرة** يجب أن تحتوي على تعليقات واضحة
- **الخوارزميات المعقدة** تحتاج شرح
- **أمثلة** للاستخدام غير البديهي
- **ملاحظات أمان الخيوط** حيث ينطبق

```go
// SetLevel يحدد الحد الأدنى لمستوى السجل.
// السجلات أسفل هذا المستوى ستتجاهل.
// هذه الطريقة آمنة للخيوط.
//
// مثال:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // لن تظهر
//   logger.Info("visible")   // ستظهر
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### تحديثات README

عند إضافة ميزات، حدث:
- README.md الرئيسي
- جميع README الخاصة باللغة في `docs/`
- أمثلة الكود
- قوائم الميزات

## 🐛 إرشادات المشاكل

### تقارير الأخطاء

استخدم [قالب تقرير الخطأ](.github/ISSUE_TEMPLATE/bug_report.md) واتضمن:

- **وصف واضح** للمشكلة
- **خطوات الإعادة**
- **السلوك المتوقع مقابل الفعلي**
- **تفاصيل البيئة** (نظام التشغيل، إصدار Go، علامات البناء)
- **عينة كود أدنى**

### طلبات الميزات

استخدم [قالب طلب الميزة](.github/ISSUE_TEMPLATE/feature_request.md) واتضمن:

- **دافع واضح** للميزة
- **تصميم API** المقترح
- **اعتبارات التنفيذ**
- **تحليل التغييرات المهمة**

### الأسئلة

استخدم [قالب السؤال](.github/ISSUE_TEMPLATE/question.md) لـ:

- أسئلة الاستخدام
- مساعدة التكوين
- أفضل الممارسات
- إرشادات التكامل

## 🚀 اعتبارات الأداء

### القياس المعياري

قس دائماً التغييرات الحساسة للأداء:

```bash
# تشغيل المعايير
go test -bench=. -benchmem

# مقارنة قبل/بعد
go test -bench=. -benchmem > before.txt
# إجراء تغييرات
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### إرشادات الأداء

- **تقليل التخصيصات** في المسارات الحرجة
- **استخدام مجمعات الكائنات** للكائنات المنشأة بكثرة
- **الإرجاع المبكر** لمستويات السجل المعطلة
- **تجنب الانعكاس** في الكود الحرج للأداء
- **قياس قبل التحسين**

### إدارة الذاكرة

```go
// ✅ جيد: استخدام مجمعات الكائنات
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

## 🔒 إرشادات الأمان

### البيانات الحساسة

- **لا تسجل أبداً** كلمات المرور أو الرموز أو البيانات الحساسة
- **طهر** مدخلات المستخدم في رسائل السجل
- **تجنب** تسجيل نصوص الطلب/الاستجابة الكاملة
- **استخدم** التسجيل المهيكل للتحكم الأفضل

```go
// ✅ جيد
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ سيء
logger.Infof("User login: %+v", userRequest) // قد تحتوي على كلمة مرور
```

### التبعيات

- أبق التبعيات **محدثة**
- **راجع** التبعيات الجديدة بعناية
- **قلل** التبعيات الخارجية
- **استخدم** `go mod verify` للتحقق من السلامة

## 👥 المجتمع

### الحصول على المساعدة

- 📖 [التوثيق](../README_ar.md)
- 💬 [مناقشات GitHub](https://github.com/lazygophers/log/discussions)
- 🐛 [متتبع المشاكل](https://github.com/lazygophers/log/issues)
- 📧 البريد الإلكتروني: support@lazygophers.com

### إرشادات التواصل

- كن **محترماً** وشاملاً
- **ابحث** قبل طرح الأسئلة
- **قدم السياق** عند طلب المساعدة
- **ساعد الآخرين** عندما تستطيع
- **اتبع** [قواعد السلوك](CODE_OF_CONDUCT_ar.md)

## 🎯 التقدير

يتم الاعتراف بالمساهمين بعدة طرق:

- قسم **مساهمي README**
- إشارات في **ملاحظات الإصدار**
- رسوم **مساهمي GitHub** البيانية
- منشورات **تقدير المجتمع**

## 📝 الرخصة

بالمساهمة، توافق على أن مساهماتك ستكون مرخصة تحت رخصة MIT.

---

## 🌍 التوثيق متعدد اللغات

هذا المستند متوفر بعدة لغات:

- [🇺🇸 English](CONTRIBUTING.md)
- [🇨🇳 简体中文](CONTRIBUTING_zh-CN.md)
- [🇹🇼 繁體中文](CONTRIBUTING_zh-TW.md)
- [🇫🇷 Français](CONTRIBUTING_fr.md)
- [🇷🇺 Русский](CONTRIBUTING_ru.md)
- [🇪🇸 Español](CONTRIBUTING_es.md)
- [🇸🇦 العربية](CONTRIBUTING_ar.md) (الحالي)

---

**شكراً لمساهمتك في LazyGophers Log! 🚀**