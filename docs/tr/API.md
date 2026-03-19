---
titleSuffix: ' | LazyGophers Log'
---
# 📚 API Dokümantasyonu

## Genel Bakış

LazyGophers Log, çoklu log seviyesi, özel biçimlendirme, zaman uyumsuz yazma ve derleme etiketi optimizasyonu desteği ile kapsamlı bir günlük kaydı API'si sağlar. Bu belge tüm genel API'leri, yapılandırma seçeneklerini ve kullanım kalıplarını kapsar.

## İçindekiler

-   [Temel Tipler](#temel-tipler)
-   [Logger API](#logger-api)
-   [Global Fonksiyonlar](#global-fonksiyonlar)
-   [Log Seviyeleri](#log-seviyeleri)
-   [Biçimlendiriciler](#biçimlendiriciler)
-   [Çıktı Yazıcıları](#çıktı-yazıcıları)
-   [Bağlam Günlük Kaydı](#bağlam-günlük-kaydı)
-   [Derleme Etiketleri](#derleme-etiketleri)
-   [Performans Optimizasyonu](#performans-optimizasyonu)
-   [Örnekler](#örnekler)

## Temel Tipler

### Logger

Tüm günlük kaydı işlevlerini sağlayan ana logger yapısı.

```go
type Logger struct {
    // İş parçacığı güvenli işlemler için özel alanlar içerir
}
```

#### Yapıcı

```go
func New() *Logger
```

Varsayılan yapılandırma ile yeni bir logger örneği oluşturur:

-   Seviye: `DebugLevel`
-   Çıktı: `os.Stdout`
-   Biçimlendirici: Varsayılan metin biçimlendirici
-   Çağıran izleme: Devre dışı

**Örnek:**

```go title="Logger oluşturma"
logger := log.New()
logger.Info("Yeni logger oluşturuldu")
```

### Entry

İlişkili tüm meta verilerle tek bir günlük girdisini temsil eder.

```go
type Entry struct {
    Time       time.Time     // Giriş oluşturulurkenki zaman damgası
    Level      Level         // Log seviyesi
    Message    string        // Log mesajı
    Pid        int          // Süreç kimliği
    Gid        uint64       // Goroutine kimliği
    TraceID    string       // Dağıtık izleme için iz kimliği
    CallerName string       // Çağıran işlev adı
    CallerFile string       // Çağıran dosya yolu
    CallerLine int          // Çağıran satır numarası
}
```

## Logger API

### Yapılandırma Metotları

#### SetLevel

```go
func (l *Logger) SetLevel(level Level) *Logger
```

Minimum log seviyesini ayarlar. Bu seviyenin altındaki mesajlar yoksayılır.

**Parametreler:**

-   `level`: İşlenecek minimum log seviyesi

**Döndürür:**

-   `*Logger`: Metod zincirlemeyi desteklemek için kendisini döndürür

**Örnek:**

```go title="Log seviyesi ayarlama"
logger.SetLevel(log.InfoLevel)
logger.Debug("Bu gösterilmez")  # Yoksayıldı
logger.Info("Bu gösterilir")    # İşlendi
```

#### SetOutput

```go
func (l *Logger) SetOutput(writers ...io.Writer) *Logger
```

Log mesajları için bir veya daha fazla çıktı hedefi ayarlar.

**Parametreler:**

-   `writers`: Bir veya daha fazla `io.Writer` çıktı hedefi

**Döndürür:**

-   `*Logger`: Metod zincirlemeyi desteklemek için kendisini döndürür

**Örnek:**

```go title="Çıktı hedefi ayarlama"
// Tek çıktı
logger.SetOutput(os.Stdout)

// Birden fazla çıktı
file, _ := os.Create("app.log")
logger.SetOutput(os.Stdout, file)
```

#### SetFormatter

```go
func (l *Logger) SetFormatter(formatter Format) *Logger
```

Log çıktısı için özel bir biçimlendirici ayarlar.

**Parametreler:**

-   `formatter`: `Format` arayüzünü uygulayan bir biçimlendirici

**Döndürür:**

-   `*Logger`: Metod zincirlemeyi desteklemek için kendisini döndürür

**Örnek:**

```go
logger.SetFormatter(&JSONFormatter{})
```

#### EnableCaller

```go
func (l *Logger) EnableCaller(enable bool) *Logger
```

Log girdilerindeki çağıran bilgisi kaydını etkinleştirir veya devre dışı bırakır.

**Parametreler:**

-   `enable`: Çağıran bilgilerinin etkinleştirilmesi (etkinleştirmek için `true` geçirin)

**Döndürür:**

-   `*Logger`: Metod zincirlemeyi desteklemek için kendisini döndürür

**Örnek:**

```go
logger.EnableCaller(true)
logger.Info("Bu dosya:satır numarası bilgisi içerecek")

logger.EnableCaller(false)
logger.Info("Bu dosya:satır numarası bilgisi içermeyecek")
```

#### Caller

```go
func (l *Logger) Caller(disable bool) *Logger
```

Biçimlendiricideki çağıran bilgilerini kontrol eder.

**Parametreler:**

-   `disable`: Çağıran bilgilerinin devre dışı bırakılması (devre dışı bırakmak için `true` geçirin)

**Döndürür:**

-   `*Logger`: Metod zincirlemeyi desteklemek için kendisini döndürür

**Örnek:**

```go
logger.Caller(false)  # Devre dışı bırakma, çağıran bilgilerini göster
logger.Info("Bu dosya:satır numarası bilgisi içerecek")

logger.Caller(true)   # Çağıran bilgilerini devre dışı bırak
logger.Info("Bu dosya:satır numarası bilgisi içermeyecek")
```

#### SetCallerDepth

```go
func (l *Logger) SetCallerDepth(depth int) *Logger
```

Logger sarmalandığında çağıran bilgileri için yığın derinliğini ayarlar.

**Parametreler:**

-   `depth`: Atlanacak yığın çerçevesi sayısı

**Döndürür:**

-   `*Logger`: Metod zincirlemeyi desteklemek için kendisini döndürür

**Örnek:**

```go
func logWrapper(msg string) {
    logger.SetCallerDepth(1).Info(msg)  # Wrapper işlevini atla
}
```

#### SetPrefixMsg / SetSuffixMsg

```go
func (l *Logger) SetPrefixMsg(prefix string) *Logger
func (l *Logger) SetSuffixMsg(suffix string) *Logger
```

Tüm log mesajları için ön ek veya son ek metni ayarlar.

**Parametreler:**

-   `prefix/suffix`: Mesaja önceden eklenecek/sonuna eklenecek metin

**Döndürür:**

-   `*Logger`: Metod zincirlemeyi desteklemek için kendisini döndürür

**Örnek:**

```go
logger.SetPrefixMsg("[APP] ").SetSuffixMsg(" [END]")
logger.Info("Hello")  # Çıktı: [APP] Hello [END]
```

### Günlük Kaydı Metotları

Tüm günlük kaydı metotlarının iki varyantı vardır: basit versiyon ve biçimlendirilmiş versiyon.

#### Trace Seviyesi

```go
func (l *Logger) Trace(v ...any)
func (l *Logger) Tracef(format string, v ...any)
```

Trace seviyesinde log kaydı yapar (en detaylı).

**Örnek:**

```go
logger.Trace("Detaylı yürütme izleme")
logger.Tracef("%d öğe işleniyor, toplam %d", i, total)
```

#### Debug Seviyesi

```go
func (l *Logger) Debug(v ...any)
func (l *Logger) Debugf(format string, v ...any)
```

Debug seviyesinde geliştirme bilgilerini kaydeder.

**Örnek:**

```go
logger.Debug("Değişken durumu:", variable)
logger.Debugf("Kullanıcı %s başarıyla kimlik doğruladı", username)
```

#### Info Seviyesi

```go
func (l *Logger) Info(v ...any)
func (l *Logger) Infof(format string, v ...any)
```

Bilgi mesajlarını kaydeder.

**Örnek:**

```go
logger.Info("Uygulama başladı")
logger.Infof("Sunucu %d portunu dinliyor", port)
```

#### Warn Seviyesi

```go
func (l *Logger) Warn(v ...any)
func (l *Logger) Warnf(format string, v ...any)
```

Olası sorunlu durumlar için uyarı mesajlarını kaydeder.

**Örnek:**

```go
logger.Warn("Kullanımdan kaldırılmış işlev çağrıldı")
logger.Warnf("Yüksek bellek kullanımı: %d%%", memoryPercent)
```

#### Error Seviyesi

```go
func (l *Logger) Error(v ...any)
func (l *Logger) Errorf(format string, v ...any)
```

Hata mesajlarını kaydeder.

**Örnek:**

```go
logger.Error("Veritabanı bağlantısı başarısız")
logger.Errorf("İstek işleme başarısız: %v", err)
```

#### Fatal Seviyesi

```go
func (l *Logger) Fatal(v ...any)
func (l *Logger) Fatalf(format string, v ...any)
```

Ölümcül bir hata kaydeder ve `os.Exit(1)` çağırır.

:::danger Yıkıcı İşlem
`Fatal` ve `Fatalf`, log kaydından hemen sonra `os.Exit(1)` çağırarak süreci sonlandırır. Yalnızca geri alınamaz hata koşullarında kullanın. `defer` ifadeleri **yürütülmez**.
:::

**Örnek:**

```go
logger.Fatal("Kritik sistem hatası")
logger.Fatalf("Sunucu başlatılamadı: %v", err)
```

#### Panic Seviyesi

```go
func (l *Logger) Panic(v ...any)
func (l *Logger) Panicf(format string, v ...any)
```

Bir hata mesajı kaydeder ve `panic()` çağırır.

:::danger Yıkıcı İşlem
`Panic` ve `Panicf`, log kaydından sonra `panic()` çağırır. `Fatal`'dan farklı olarak, `panic` `recover()` ile yakalanabilir, ancak yakalanmazsa programı sonlandırır.
:::

**Örnek:**

```go
logger.Panic("Geri alınamaz bir hata oluştu")
logger.Panicf("Geçersiz durum: %v", state)
```

### Yardımcı Metotlar

#### Clone

```go
func (l *Logger) Clone() *Logger
```

Aynı yapılandırmaya sahip bir logger kopyası oluşturur.

**Döndürür:**

-   `*Logger`: Kopyalanmış ayarlarla yeni logger örneği

**Örnek:**

```go
dbLogger := logger.Clone()
dbLogger.SetPrefixMsg("[DB] ")
```

#### CloneToCtx

```go
func (l *Logger) CloneToCtx() LoggerWithCtx
```

`context.Context`'i ilk parametre olarak kabul eden bağlam farkında bir logger oluşturur.

**Döndürür:**

-   `LoggerWithCtx`: Bağlam farkında logger örneği

**Örnek:**

```go
ctxLogger := logger.CloneToCtx()
ctxLogger.Info(ctx, "Bağlam farkında mesaj")
```

## Global Fonksiyonlar

Varsayılan global logger'ı kullanan paket düzeyi fonksiyonlar.

```go
func SetLevel(level Level)
func SetOutput(writers ...io.Writer)
func SetFormatter(formatter Format)
func Caller(disable bool)

func Trace(v ...any)
func Tracef(format string, v ...any)
func Debug(v ...any)
func Debugf(format string, v ...any)
func Info(v ...any)
func Infof(format string, v ...any)
func Warn(v ...any)
func Warnf(format string, v ...any)
func Error(v ...any)
func Errorf(format string, v ...any)
func Fatal(v ...any)
func Fatalf(format string, v ...any)
func Panic(v ...any)
func Panicf(format string, v ...any)
```

**Örnek:**

```go
import "github.com/lazygophers/log"

log.SetLevel(log.InfoLevel)
log.Info("Global logger kullanımı")
```

## Log Seviyeleri

### Level Türü

```go
type Level int8
```

### Mevcut Seviyeler

```go
const (
    PanicLevel Level = iota  // 0 - Panik ve çık
    FatalLevel              // 1 - Ölümcül hata ve çık
    ErrorLevel              // 2 - Hata koşulu
    WarnLevel               // 3 - Uyarı koşulu
    InfoLevel               // 4 - Bilgi mesajı
    DebugLevel              // 5 - Hata ayıklama mesajı
    TraceLevel              // 6 - En detaylı izleme
)
```

### Level Metotları

```go
func (l Level) String() string
```

Seviyenin dize gösterimini döndürür.

**Örnek:**

```go
fmt.Println(log.InfoLevel.String())  // "INFO"
```

## Biçimlendiriciler

### Format Arayüzü

```go
type Format interface {
    Format(entry *Entry) []byte
}
```

Özel biçimlendiriciler bu arayüzü uygulamalıdır.

### Varsayılan Biçimlendirici

Özelleştirilebilir seçeneklerle yerleşik metin biçimlendirici.

```go
type Formatter struct {
    // Yapılandırma seçenekleri
}
```

### JSON Biçimlendirici Örneği

```go
type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "caller":    fmt.Sprintf("%s:%d", entry.CallerFile, entry.CallerLine),
    }
    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }

    jsonData, _ := json.Marshal(data)
    return append(jsonData, '\n')
}

// Kullanım
logger.SetFormatter(&JSONFormatter{})
```

## Çıktı Yazıcıları

### Dosya Çıktısı ve Döndürme

```go
func GetOutputWriterHourly(filename string) io.Writer
```

Log dosyalarını saatlik olarak döndüren bir yazıcı oluşturur.

**Parametreler:**

-   `filename`: Log dosyasının temel dosya adı

**Döndürür:**

-   `io.Writer`: Döndüren dosya yazıcısı

**Örnek:**

```go title="Saatlik log döndürme"
writer := log.GetOutputWriterHourly("./logs/app.log")
logger.SetOutput(writer)
// Şu dosyaları oluşturur: app-2024010115.log, app-2024010116.log, vb.
```

### Zaman Uyumsuz Yazıcı

```go
func NewAsyncWriter(writer io.Writer, bufferSize int) *AsyncWriter
```

Yüksek performanslı günlük kaydı için zaman uyumsuz bir yazıcı oluşturur.

**Parametreler:**

-   `writer`: Alt yazıcı
-   `bufferSize`: Dahili arabellek boyutu

**Döndürür:**

-   `*AsyncWriter`: Zaman uyumsuz yazıcı örneği

**Metotlar:**

```go
func (aw *AsyncWriter) Write(data []byte) (int, error)
func (aw *AsyncWriter) Close() error
```

**Örnek:**

```go title="Zaman uyumsuz yazıcı"
file, _ := os.Create("app.log")
asyncWriter := log.NewAsyncWriter(file, 1000)
defer asyncWriter.Close()

logger.SetOutput(asyncWriter)
```

## Bağlam Günlük Kaydı

### LoggerWithCtx Arayüzü

```go
type LoggerWithCtx interface {
    Trace(ctx context.Context, v ...any)
    Tracef(ctx context.Context, format string, v ...any)
    Debug(ctx context.Context, v ...any)
    Debugf(ctx context.Context, format string, v ...any)
    Info(ctx context.Context, v ...any)
    Infof(ctx context.Context, format string, v ...any)
    Warn(ctx context.Context, v ...any)
    Warnf(ctx context.Context, format string, v ...any)
    Error(ctx context.Context, v ...any)
    Errorf(ctx context.Context, format string, v ...any)
    Fatal(ctx context.Context, v ...any)
    Fatalf(ctx context.Context, format string, v ...any)
    Panic(ctx context.Context, v ...any)
    Panicf(ctx context.Context, format string, v ...any)
}
```

### Bağlam Fonksiyonları

```go
func SetTrace(traceID string)
func GetTrace() string
```

Geçerli goroutine için iz kimliğini ayarlar ve alır.

**Örnek:**

```go
log.SetTrace("trace-123-456")
log.Info("Bu mesaj iz kimliğini içerecek")

traceID := log.GetTrace()
fmt.Println("Geçerli iz kimliği:", traceID)
```

## Derleme Etiketleri

Bu kitaplık derleme etiketleri ile koşullu derlemeyi destekler:

:::info Derleme Etiketi Açıklaması
Derleme etiketleri `go build -tags` parametresi ile belirtilir. Farklı etiketler log kitaplığının derleme davranışını ve çalışma zamanı özelliklerini değiştirir. Uygun etiketleri seçmek, geliştirme kolaylığı ve üretim performansı arasında denge sağlar.
:::

### Varsayılan Mod

```bash
go build
```

-   Tam işlevsellik etkin
-   Hata ayıklama mesajları dahil
-   Standart performans

### Hata Ayıklama Modu

```bash
go build -tags debug
```

-   Gelişmiş hata ayıklama bilgileri
-   Ek çalışma zamanı kontrolleri
-   Detaylı çağıran bilgileri

### Yayın Modu

```bash
go build -tags release
```

-   Üretim ortamı için optimize edilmiş
-   Hata ayıklama mesajları devre dışı
-   Otomatik log döndürme etkin

### Yoksayma Modu

```bash
go build -tags discard
```

-   Maksimum performans
-   Tüm günlük kaydı işlemleri no-op'lar
-   Sıfır ek yükleme

### Birleşik Mod

```bash
go build -tags "debug,discard"    # Hata ayıklama ve yoksayma
go build -tags "release,discard"  # Yayın ve yoksayma
```

## Performans Optimizasyonu

:::tip Performans En İyi Uygulamaları
Bu kitaplık nesne havuzları, ön seviye kontrolleri ve zaman uyumsuz yazma gibi mekanizmalar aracılığıyla derinlemesine optimize edilmiştir. Yüksek aktarım senaryolarında, en iyi performansı elde etmek için zaman uyumsuz yazıcılar ve uygun derleme etiketlerini birleştirmeniz önerilir.
:::

### Nesne Havuzları

Kitaplık yönetmek için dahili olarak `sync.Pool` kullanır:

-   Log girdisi nesneleri
-   Bayt arabellekleri
-   Biçimlendirici arabellekleri

Bu, yüksek aktarım senaryolarında çöp toplama basıncını azaltır.

### Seviye Kontrolü

Log seviyesi kontrolleri pahalı işlemlerden önce gerçekleşir:

```go
# Verimli - Seviye etkin olduğunda mesaj biçimlendirme
logger.Debugf("Pahalı işlem sonucu: %+v", expensiveCall())

# Prodüksiyonda hata ayıklama devre dışıyken daha az verimli
result := expensiveCall()
logger.Debug("Sonuç:", result)
```

### Zaman Uyumsuz Yazma

Yüksek aktarım uygulamaları için:

```go
asyncWriter := log.NewAsyncWriter(file, 10000)  # Büyük arabellek
logger.SetOutput(asyncWriter)
defer asyncWriter.Close()
```

### Derleme Etiketi Optimizasyonu

Ortam için uygun derleme etiketlerini kullanın:

-   Geliştirme: Varsayılan veya hata ayıklama etiketleri
-   Prodüksiyon: Yayın etiketleri
-   Performans kritik: Yoksayma etiketleri

## Örnekler

### Temel Kullanım

```go title="Temel kullanım"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    log.SetLevel(log.InfoLevel)
    log.Info("Uygulama başlatılıyor")
    log.Warn("Bu bir uyarıdır")
    log.Error("Bu bir hatadır")
}
```

### Özel Logger

```go title="Özel yapılandırma"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()

    # Logger'ı yapılandır
    logger.SetLevel(log.DebugLevel)
    logger.Caller(true)  # Çağıran bilgilerini devre dışı bırak
    logger.SetPrefixMsg("[MyApp] ")

    # Dosya çıktısını ayarla
    file, err := os.Create("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    logger.SetOutput(file)

    logger.Info("Özel logger yapılandırıldı")
    logger.Debug("Çağıran bilgili hata ayıklama bilgileri")
}
```

### Yüksek Performanslı Günlük Kaydı

```go title="Yüksek performanslı günlük kaydı"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    # Saatlik döndüren dosya yazıcısı oluştur
    writer := log.GetOutputWriterHourly("./logs/app.log")

    # Daha iyi performans için zaman uyumsuz yazıcı kullan
    asyncWriter := log.NewAsyncWriter(writer, 5000)
    defer asyncWriter.Close()

    logger := log.New()
    logger.SetOutput(asyncWriter)
    logger.SetLevel(log.InfoLevel)  # Prodüksiyonda debug loglarını atla

    # Yüksek aktarım günlük kaydı
    for i := 0; i < 10000; i++ {
        logger.Infof("Processing request %d", i)
    }
}
```

### Bağlam Farkında Günlük Kaydı

```go title="Bağlam farkında günlük kaydı"
package main

import (
    "context"
    "github.com/lazygophers/log"
)

func main() {
    logger := log.New()
    ctxLogger := logger.CloneToCtx()

    ctx := context.Background()
    log.SetTrace("trace-123-456")

    ctxLogger.Info(ctx, "Kullanıcı isteği işleniyor")
    ctxLogger.Debug(ctx, "Doğrulama tamamlandı")
}
```

### Özel JSON Biçimlendirici

```go title="Özel JSON biçimlendirici"
package main

import (
    "encoding/json"
    "os"
    "time"
    "github.com/lazygophers/log"
)

type JSONFormatter struct{}

func (f *JSONFormatter) Format(entry *log.Entry) []byte {
    data := map[string]interface{}{
        "timestamp": entry.Time.Format(time.RFC3339Nano),
        "level":     entry.Level.String(),
        "message":   entry.Message,
        "pid":       entry.Pid,
        "gid":       entry.Gid,
    }

    if entry.TraceID != "" {
        data["trace_id"] = entry.TraceID
    }

    if entry.CallerName != "" {
        data["caller"] = map[string]interface{}{
            "function": entry.CallerName,
            "file":     entry.CallerFile,
            "line":     entry.CallerLine,
        }
    }

    jsonData, _ := json.MarshalIndent(data, "", "  ")
    return append(jsonData, '\n')
}

func main() {
    logger := log.New()
    logger.SetFormatter(&JSONFormatter{})
    logger.Caller(true)  # Çağıran bilgilerini devre dışı bırak
    logger.SetOutput(os.Stdout)

    log.SetTrace("request-456")
    logger.Info("JSON biçimlendirilmiş mesaj")
}
```

## Hata İşleme

:::warning Uyarı
Performans nedenleriyle, çoğu logger metodu hata döndürmez. Yazma başarısız olursa, loglar sessizce atılır. Hata farkındalığına ihtiyacınız varsa, özel bir yazıcı kullanın.
:::

Çıktı işlemleri için hata işlemeeye ihtiyacınız varsa, özel bir yazıcı uygulayın:

```go title="Hata yakalayan yazıcı"
type ErrorCapturingWriter struct {
    writer io.Writer
    lastError error
}

func (w *ErrorCapturingWriter) Write(data []byte) (int, error) {
    n, err := w.writer.Write(data)
    if err != nil {
        w.lastError = err
    }
    return n, err
}

func (w *ErrorCapturingWriter) LastError() error {
    return w.lastError
}
```

## İş Parçacığı Güvenliği

:::tip Eşzamanlılık Güvenliği
Tüm `Logger` örneklerinin metotları iş parçacığı güvenlidir ve ek synchronizasyon mekanizmaları olmadan birden çok goroutine'de eşzamanlı olarak kullanılabilir. Ancak, tek tek `Entry` nesnelerinin iş parçacığı güvenli **olmadığını** ve tek kullanımlık olduğunu unutmayın.
:::

---

## 🌍 Çok Dilli Dokümantasyon

Bu belge birden çok dilde mevcuttur:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/API.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/API.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/API.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/API.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/API.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/API.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/API.md)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/API.md)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/API.md)
-   [🇧🇷 Português](https://lazygophers.github.io/log/pt/API.md)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/API.md)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/API.md)
-   [🇵🇱 Polski](https://lazygophers.github.io/log/pl/API.md)
-   [🇹🇷 Türkçe](API.md) (Mevcut)

---

**LazyGophers Log tam API referansı - Olağanüstü günlük kaydı ile daha iyi uygulamalar oluşturun! 🚀**
