---
titleSuffix: ' | LazyGophers Log'
---
# lazygophers/log

[![Go Version](https://img.shields.io/badge/go-1.19+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.0%25-brightgreen.svg)](https://github.com/lazygophers/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/lazygophers/log)](https://goreportcard.com/report/github.com/lazygophers/log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Yüksek performanslı ve esnek bir Go günlük kütüphanesi, zap üzerine inşa edilmiş, zengin özellikler ve basit bir API sunar.

## 📖 Belgelerin dilleri

-   [🇺🇸 English](https://lazygophers.github.io/log/en/)
-   [🇨🇳 Basitleştirilmiş Çince](https://lazygophers.github.io/log/zh-CN/)
-   [🇹🇼 Geleneksel Çince](https://lazygophers.github.io/log/zh-TW/)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/)
-   [🇯🇵 日本語](https://lazygophers.github.io/log/ja/)
-   [🇩🇪 Deutsch](https://lazygophers.github.io/log/de/)
-   [🇰🇷 한국어](https://lazygophers.github.io/log/ko/)
-   [🇵🇹 Português](https://lazygophers.github.io/log/pt/)
-   [🇳🇱 Nederlands](https://lazygophers.github.io/log/nl/)
-   [🇵🇱 Polski](https://lazygophers.github.io/log/pl/)
-   [🇮🇹 Italiano](https://lazygophers.github.io/log/it/)
-   [🇹🇷 Türkçe](README.md) (şu an)

## ✨ Özellikler

-   **🚀 Yüksek performans**：Nesne havuzlama ve koşullu alan kaydı ile zap üzerine inşa edilmiş
-   **📊 Zengin günlük seviyeleri**：Trace, Debug, Info, Warn, Error, Fatal, Panic seviyeleri
-   **⚙️ Esnek yapılandırma**：
    -   Günlük seviyesi kontrolü
    -   Çağıran bilgi kaydı
    -   İzleme bilgileri (goroutine ID dahil)
    -   Özel önek ve sonek
    -   Özel çıktı hedefleri (konsol, dosyalar vb.)
    -   Günlük biçimlendirme seçenekleri
-   **🔄 Dosya döndürme**：Saatlik günlük dosyası döndürme desteği
-   **🔌 Zap uyumluluğu**：Zap WriteSyncer ile sorunsuz entegrasyon
-   **🎯 Basit API**：Standart günlük kütüphanesine benzer net API, kullanımı kolay

## 🚀 Hızlı başlangıç

### Kurulum

:::tip Kurulum
```bash
go get github.com/lazygophers/log
```
:::

### Temel kullanım

```go title="Hızlı başlangıç"
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Varsayılan global logger kullan
    log.Debug("Hata ayıklama mesajı")
    log.Info("Bilgi mesajı")
    log.Warn("Uyarı mesajı")
    log.Error("Hata mesajı")

    // Biçimlendirilmiş çıktı kullan
    log.Infof("Kullanıcı %s başarıyla giriş yaptı", "admin")

    // Özel yapılandırma
    customLogger := log.New().
        SetLevel(log.InfoLevel).
        EnableCaller(false).
        SetPrefixMsg("[MyApp]")

    customLogger.Info("Bu özel logger'dan bir günlüktür")
}
```

## 📚 Gelişmiş kullanım

### Dosya çıktılı özel logger

```go title="Dosya çıktısı yapılandırması"
package main

import (
    "os"
    "github.com/lazygophers/log"
)

func main() {
    // Dosya çıktılı logger oluştur
    logger := log.New().
        SetLevel(log.DebugLevel).
        EnableCaller(true).
        EnableTrace(true).
        SetOutput(os.Stdout, log.GetOutputWriterHourly("/var/log/myapp.log"))

    logger.Debug("Çağıran bilgili hata ayıklama günlüğü")
    logger.Info("İzleme bilgili bilgi günlüğü")
}
```

### Günlük seviyesi kontrolü

```go title="Günlük seviyesi kontrolü"
package main

import "github.com/lazygophers/log"

func main() {
    logger := log.New().SetLevel(log.WarnLevel)

    // Sadece warn ve üzeri kaydedilecektir
    logger.Debug("Bu kaydedilmeyecek")  // Yoksayıldı
    logger.Info("Bu kaydedilmeyecek")   // Yoksayıldı
    logger.Warn("Bu kaydedilecek")    // Kaydedildi
    logger.Error("Bu kaydedilecek")   // Kaydedildi
}
```

## 🎯 Kullanım senaryoları

### Uygulanabilir senaryolar

-   **Web servisleri ve API arka uçları**：İstek izleme, yapılandırılmış günlük, performans izleme
-   **Mikroservis mimarisi**：Dağıtılmış izleme (TraceID), birleştirilmiş günlük biçimi, yüksek verim
-   **Komut satırı araçları**：Seviye kontrolü, temiz çıktı, hata raporlama
-   **Toplu iş görevleri**：Dosya döndürme, uzun çalışma, kaynak optimizasyonu

### Özel avantajlar

-   **Nesne havuzu optimizasyonu**：Entry ve Buffer nesnelerinin yeniden kullanımı, GC baskısını azaltır
-   **Eşzamanlı olmayan yazma**：Yüksek verim senaryoları (10000+ günlük/saniye) engelleme olmadan
-   **TraceID desteği**：Dağıtılmış sistemlerde istek izleme, OpenTelemetry ile entegrasyon
-   **Sıfır yapılandırma başlangıcı**：Kullanıma hazır, aşamalı yapılandırma

## 🔧 Yapılandırma seçenekleri

:::note Yapılandırma seçenekleri
Aşağıdaki tüm yöntemler zincirleme çağrıyı destekler ve özel Logger oluşturmak için birleştirilebilir.
:::

### Logger yapılandırması

| Yöntem                  | Açıklama                | Varsayılan      |
| --------------------- | ------------------- | ------------ |
| `SetLevel(level)`       | Minimum günlük seviyesi ayarla     | `DebugLevel` |
| `EnableCaller(enable)`  | Çağıran bilgilerini etkinleştir/devre dışı bırak | `false`      |
| `EnableTrace(enable)`   | İzleme bilgilerini etkinleştir/devre dışı bırak  | `false`      |
| `SetCallerDepth(depth)` | Çağıran derinliğini ayarla   | `2`          |
| `SetPrefixMsg(prefix)`  | Günlük öneki ayarla  | `""`         |
| `SetSuffixMsg(suffix)`  | Günlük soneki ayarla  | `""`         |
| `SetOutput(writers...)` | Çıktı hedeflerini ayarla         | `os.Stdout`  |

### Günlük seviyeleri

| Seviye        | Açıklama                        |
| ----------- | --------------------------- |
| `TraceLevel` | En ayrıntılı, detaylı izleme için        |
| `DebugLevel` | Hata ayıklama bilgileri                  |
| `InfoLevel`  | Genel bilgiler                    |
| `WarnLevel`  | Uyarı mesajları                  |
| `ErrorLevel` | Hata mesajları                  |
| `FatalLevel` | Ölümcül hatalar (os.Exit(1) çağırır)    |
| `PanicLevel` | Panik hataları (panic() çağırır)      |

## 🏗️ Mimari

### Temel bileşenler

-   **Logger**：Yapılandırılabilir seçeneklere sahip ana günlük yapısı
-   **Entry**：Kapsamlı alan desteğine sahip bireysel günlük kaydı
-   **Level**：Günlük seviyesi tanımları ve yardımcı işlevler
-   **Format**：Günlük biçimlendirme arayüzü ve uygulamaları

### Performans optimizasyonu

-   **Nesne havuzu**：Bellek ayırmayı azaltmak için Entry nesnelerini yeniden kullanır
-   **Koşullu kayıt**：Gerekli olduğunda sadece maliyetli alanları kaydeder
-   **Hızlı seviye kontrolü**：Günlük seviyesini en dış katmanda kontrol eder
-   **Kilit yok tasarım**：Çoğu işlem kilit gerektirmez

## 📊 Performans karşılaştırması

:::info Performans karşılaştırması
Aşağıdaki veriler benchmark'lara dayanmaktadır; gerçek performans ortam ve yapılandırmaya göre değişebilir.
:::

| Özellik          | lazygophers/log | zap    | logrus | Standart günlük |
| ------------- | --------------- | ------ | ------ | -------------- |
| Performans      | Yüksek              | Yüksek     | Orta     | Düşük       |
| API basitliği    | Yüksek              | Orta     | Yüksek     | Yüksek       |
| Özellik zenginliği    | Orta          | Yüksek     | Yüksek     | Düşük       |
| Esneklik      | Orta          | Yüksek     | Yüksek     | Düşük       |
| Öğrenme eğrisi      | Düşük              | Orta     | Orta     | Düşük       |

## ❓ Sık sorulan sorular

### Uygun günlük seviyesini nasıl seçerim?

-   **Geliştirme ortamı**：Detaylı bilgi için `DebugLevel` veya `TraceLevel` kullanın
-   **Üretim ortamı**：Yükü azaltmak için `InfoLevel` veya `WarnLevel` kullanın
-   **Performans testleri**：Tüm günlükleri devre dışı bırakmak için `PanicLevel` kullanın

### Üretim ortamında performansı nasıl optimize ederim?

:::warning Not
Yüksek verim senaryolarında, performansı optimize etmek için eşzamanlı olmayan yazmayı ve makul günlük seviyelerini birleştirmeniz önerilir.
:::

1. Eşzamanlı olmayan yazma için `AsyncWriter` kullanın：

```go title="Eşzamanlı olmayan yazma yapılandırması"
writer := log.GetOutputWriterHourly("./logs/app.log")
asyncWriter := log.NewAsyncWriter(writer, 5000)
logger.SetOutput(asyncWriter)
```

2. Gereksiz günlükleri önlemek için günlük seviyelerini ayarlayın：

```go title="Seviye optimizasyonu"
logger.SetLevel(log.InfoLevel)  // Debug ve Trace'i atla
```

3. Yükü azaltmak için koşullu günlük kullanın：

```go title="Koşullu günlük"
if logger.Level >= log.DebugLevel {
    logger.Debug("Detaylı hata ayıklama bilgileri")
}
```

### `Caller` ile `EnableCaller` arasındaki fark nedir?

-   **`EnableCaller(enable bool)`**：Logger'ın çağıran bilgilerini toplayıp toplamadığını kontrol eder
    -   `EnableCaller(true)` çağıran bilgisi toplamayı etkinleştirir
-   **`Caller(disable bool)`**：Formatter'ın çağıran bilgilerini çıkarıp çıkarmadığını kontrol eder
    -   `Caller(true)` çağıran bilgisi çıkarmayı devre dışı bırakır

Genel kontrol için `EnableCaller` kullanmanız önerilir.

### Özel bir formatleyici nasıl uygularım?

`Format` arayüzünü uygulayın：

```go title="Özel formatleyici"
type MyFormatter struct{}

func (f *MyFormatter) Format(entry *log.Entry) []byte {
    return []byte(fmt.Sprintf("[%s] %s\n",
        entry.Level.String(), entry.Message))
}

logger.SetFormatter(&MyFormatter{})
```

## 🔗 İlgili belgeler

-   [📚 API belgeleri](API.md) - Tam API başvurusu
-   [🤝 Katkı rehberi](/tr/CONTRIBUTING) - Nasıl katkı yapılır
-   [📋 Değişiklik günlüğü](/tr/CHANGELOG) - Sürüm geçmişi
-   [🔒 Güvenlik politikası](/tr/SECURITY) - Güvenlik yönergeleri
-   [📜 Davranış kuralları](/tr/CODE_OF_CONDUCT) - Topluluk yönergeleri

## 🚀 Yardım alma

-   **GitHub Issues**：[Hata bildirin veya özellik isteyin](https://github.com/lazygophers/log/issues)
-   **GoDoc**：[API belgeleri](https://pkg.go.dev/github.com/lazygophers/log)
-   **Örnekler**：[Kullanım örnekleri](https://github.com/lazygophers/log/tree/main/examples)

## 📄 Lisans

Bu proje MIT Lisansı altında lisanslanmıştır - ayrıntılar için [LICENSE](/tr/LICENSE) dosyasına bakın.

## 🤝 Katkıda bulunma

Katkıları memnuniyetle karşılıyoruz! [Katkı rehberimize](/tr/CONTRIBUTING) bakın.

---

**lazygophers/log**, hem performans hem de basitliği önemseyen Go geliştiricileri için tercih edilen günlük çözümü olacak şekilde tasarlanmıştır. Küçük bir araç mı yoksa büyük ölçekli dağıtılmış bir sistem mi inşa ediyorsanız, bu kütüphane işlevsellik ve kullanım kolaylığı arasında doğru dengeyi sağlar.
