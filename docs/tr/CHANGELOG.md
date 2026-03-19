---
titleSuffix: " | LazyGophers Log"
---

# 📋 Değişiklik Günlüğü

Bu projedeki tüm önemli değişiklikler bu dosyada kaydedilir.

Format [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) tabanlıdır, proje [Semantic Versioning](https://semver.org/spec/v2.0.0.html) takip eder.

## [Yayınlanmamış]

### Eklendi

-   Kapsamlı çok dilli belgeler (7 dil)
-   GitHub issue şablonları (hata raporları, özellik istekleri, sorular)
-   Derleme etiketi uyumluluk kontrolü ile Pull Request şablonu
-   Çok dilli katkı kılavuzu
-   Yürütme yönergeleriyle davranış kuralları
-   Güvenlik açığı bildirme süreciyle güvenlik politikası
-   Örneklerle tam API belgeleri
-   Profesyonel proje yapısı ve şablonlar

### Değiştirildi

-   Kapsamlı özellik belgeleriyle README geliştirildi
-   Tüm derleme etiketi yapılandırmaları için test kapsamı iyileştirildi
-   Bakım kolaylığı için proje yapısı güncellendi

### Belgeler

-   Tüm ana belgeler için çok dil desteği eklendi
-   Kapsamlı API referansı oluşturuldu
-   Katkı iş akışı kılavuzları kuruldu
-   Güvenlik raporlama süreci uygulandı

## [1.0.0] - 2024-01-01

### Eklendi

-   Çoklu günlük seviyeli temel günlük özellikleri (Trace, Debug, Info, Warn, Error, Fatal, Panic)
-   Nesne havuzu ile iş parçacığı güvenlı günlükçü uygulaması
-   Derleme etiketi desteği (varsayılan, hata ayıklama, yayın, atma modları)
-   Varsayılan metin formatlayıcı ile özel formatlayıcı arayüzü
-   Çoklu yazıcı çıkış desteği
-   Yüksek aktarım senaryoları için zaman uyumsuz yazma özelliği
-   Otomatik saatlik günlük dosyası döndürme
-   Goroutine ID ve izleme ID izleme ile bağlam duyarlı günlük
-   Yapılandırılabilir yığın derinliği ile çağıran bilgisi
-   Genel paket seviyesi kolaylık fonksiyonları
-   Zap günlükçü entegrasyon desteği

### Performans

-   `sync.Pool` kullanarak giriş nesneleri ve arabellekler için nesne havuzlama
-   Pahalı işlemlerden kaçınmak için erken seviye kontrolü
-   Engellenemez günlük yazması için zaman uyumsuz yazıcı
-   Farklı ortamlar için derleme etiketi optimizasyonları

### Derleme Etiketleri

-   **Varsayılan**: Hata ayıklama mesajlarıyla tam işlevsellik
-   **Hata ayıklama**: Gelişmiş hata ayıklama bilgileri ve çağıran ayrıntıları
-   **Yayın**: Üretim optimizasyonu, hata ayıklama mesajları devre dışı
-   **Atma**: Maksimum performans, no-op günlük işlemleri

### Temel Özellikler

-   **Logger**: Yapılandırılabilir seviye, çıktı, formatlayıcı ile ana günlükçü yapısı
-   **Entry**: Kapsamlı meta verilerle günlük kaydı yapısı
-   **Seviyeler**: Panic (en yüksek) Trace (en düşük) yedi günlük seviyesi
-   **Formatlayıcılar**: Takılabilir formatlama sistemi
-   **Yazıcılar**: Dosya döndürme ve zaman uyumsuz yazma desteği
-   **Bağlam**: Goroutine ID ve dağıtılmış izleme desteği

### API Öne Çıkanlar

-   Metod zinciri ile akıcı yapılandırma API'si
-   Basit ve formatlı günlük metodları (`.Info()` ve `.Infof()`)
-   Yalıtılmış yapılandırma için günlükçü klonlama
-   `CloneToCtx()` ile bağlam duyarlı günlük
-   Önek ve sonek mesajı özelleştirme
-   Çağıran bilgisi anahtarı

### Test Etme

-   %93.5 kapsama oranıyla kapsamlı test paketi
-   Çoklu derleme etiketi test desteği
-   Otomatik test iş akışları
-   Performans kıyaslamaları

## [0.9.0] - 2023-12-15

### Eklendi

-   İlk proje yapısı
-   Temel günlük işlevleri
-   Seviye tabanlı filtreleme
-   Dosya çıkış desteği

### Değiştirildi

-   Nesne havuzu ile performans iyileştirildi
-   Hata işleme geliştirildi

## [0.8.0] - 2023-12-01

### Eklendi

-   Çoklu yazıcı desteği
-   Özel formatlayıcı arayüzü
-   Zaman uyumsuz yazma özelliği

### Düzeltildi

-   Yüksek aktarım senaryolarındaki bellek sızıntıları
-   Eşzamanlı erişimdeki yarış koşulları

## [0.7.0] - 2023-11-15

### Eklendi

-   Koşullu derleme için derleme etiketi desteği
-   Trace ve hata ayıklama seviyesi günlükleri
-   Çağıran bilgisi izleme

### Değiştirildi

-   Bellek ayırma desenleri optimize edildi
-   İş parçacığı güvenliği iyileştirildi

## [0.6.0] - 2023-11-01

### Eklendi

-   Günlük döndürme özelliği
-   Bağlam duyarlı günlük
-   Goroutine ID izleme

### Kullanımdan Kaldırıldı

-   Eski yapılandırma metodları (v1.0.0'da kaldırılacak)

## [0.5.0] - 2023-10-15

### Eklendi

-   JSON formatlayıcı
-   Çoklu çıkış hedefleri
-   Performans kıyaslamaları

### Değiştirildi

-   Çekirdek günlük motoru yeniden düzenlendi
-   API tutarlılığı iyileştirildi

### Kaldırıldı

-   Eski günlük metodları

## [0.4.0] - 2023-10-01

### Eklendi

-   Fatal ve Panic seviyesi günlükleri
-   Genel paket fonksiyonları
-   Yapılandırma doğrulama

### Düzeltildi

-   Çıkış senkronizasyon sorunları
-   Bellek kullanımı optimizasyonu

## [0.3.0] - 2023-09-15

### Eklendi

-   Özel günlük seviyeleri
-   Formatlayıcı arayüzü
-   İş parçacığı güvenli işlemler

### Değiştirildi

-   API tasarımı basitleştirildi
-   Belgeler genişletildi

## [0.2.0] - 2023-09-01

### Eklendi

-   Dosya çıkış desteği
-   Seviye tabanlı filtreleme
-   Temel biçimlendirme seçenekleri

### Düzeltildi

-   Performans darboğazları
-   Bellek sızıntıları

## [0.1.0] - 2023-08-15

### Eklendi

-   İlk sürüm
-   Temel konsol günlüğü
-   Basit seviye desteği (Info, Warn, Error)
-   Çekirdek günlükçü yapısı

## Sürüm Geçmişi Özeti

| Sürüm  | Yayın Tarihi   | Ana Özellikler                                       |
| ----- | ---------- | ---------------------------------------------- |
| 1.0.0 | 2024-01-01 | Tam günlük sistemi, derleme etiketleri, zaman uyumsuz yazma, kapsamlı belgeler |
| 0.9.0 | 2023-12-15 | Performans iyileştirmeleri, nesne havuzları |
| 0.8.0 | 2023-12-01 | Çoklu yazıcı, zaman uyumsuz yazma, özel formatlayıcılar |
| 0.7.0 | 2023-11-15 | Derleme etiketleri, Trace/hata ayıklama seviyeleri, çağıran bilgisi |
| 0.6.0 | 2023-11-01 | Günlük döndürme, bağlam günlüğü, goroutine izleme |
| 0.5.0 | 2023-10-15 | JSON formatlayıcı, çoklu çıkış, kıyaslamalar |
| 0.4.0 | 2023-10-01 | Fatal/Panic seviyeleri, genel fonksiyonlar |
| 0.3.0 | 2023-09-15 | Özel seviyeler, formatlayıcı arayüzü |
| 0.2.0 | 2023-09-01 | Dosya çıkışı, seviye filtreleme |
| 0.1.0 | 2023-08-15 | İlk sürüm, temel konsol günlüğü |

## Geçiş Kılavuzu

### v0.9.x'ten v1.0.0'a Geçiş

#### Bozan Değişiklikler

-   Yok - v1.0.0, v0.9.x ile geriye uyumludur

### Yeni Kullanılabilir Özellikler

-   Gelişmiş derleme etiketi desteği
-   Kapsamlı belgeler
-   Profesyonel proje şablonları
-   Güvenlik açığı raporlama süreci

#### Önerilen Güncellemeler

```go
// Eski yol (halen destekleniyor)
logger := log.New()
logger.SetLevel(log.InfoLevel)

// Önerilen yeni yol, metod zinciri kullanarak
logger := log.New().
    SetLevel(log.InfoLevel).
    Caller(true).
    SetPrefixMsg("[MyApp] ")
```

### v0.8.x'ten v0.9.x'e Geçiş

#### Bozan Değişiklikler

-   Kullanımdan kaldırılmış yapılandırma metodları kaldırıldı
-   Dahili arabellek yönetimi değiştirildi

#### Geçiş Adımları

1. Gerekirse içe aktarma yollarını güncelleyin
2. Kullanımdan kaldırılmış metodları değiştirin:

    ```go
    // Eski (kullanımdan kaldırıldı)
    logger.SetOutputFile("app.log")

    // Yeni
    file, _ := os.Create("app.log")
    logger.SetOutput(file)
    ```

### v0.5.x ve Öncesinden Geçiş

#### Ana Değişiklikler

-   Daha iyi tutarlılık için API'nın tamamen yeniden tasarlanması
-   Nesne havuzu ile performans artırması
-   Yeni derleme etiketi sistemi

#### Geçiş Gerekli

-   Tüm günlük çağrılarını yeni API'ye güncelleyin
-   Formatlayıcı uygulamalarını gözden geçirin ve güncelleyin
-   Yeni derleme etiketi yapılandırmalarıyla test edin

## Geliştirme Kilometre Taşları

### 🎯 v1.1.0 Yol Haritası (planlandı)

-   [ ] Anahtar-değer çiftleri ile yapılandırılmış günlük
-   [ ] Yüksek hacim senaryoları için günlük örnekleme
-   [ ] Özel çıkışlar için eklenti sistemi
-   [ ] Gelişmiş performans metrikleri
-   [ ] Bulut günlük entegrasyonu

### 🎯 v1.2.0 Yol Haritası (gelecek)

-   [ ] Yapılandırma dosyası desteği (YAML/JSON/TOML)
-   [ ] Günlük birleştirme ve filtreleme
-   [ ] Gerçek zamanlı günlük akışı
-   [ ] Gelişmiş güvenlik özellikleri
-   [ ] Performans gösterge paneli entegrasyonu

## Katkıda Bulunma

Katkıları memnuniyetle karşılıyoruz! Ayrıntılar için lütfen [Katkıda Bulunma Kılavuzumuzu](/tr/CONTRIBUTING) inceleyin:

-   Hata raporlama ve özellik talepleri
-   Kod gönderme iş akışı
-   Geliştirme kurulumu
-   Test gereksinimleri
-   Belge standartları

## Güvenlik

Güvenlik açıkları için lütfen [Güvenlik Politikamızı](/tr/SECURITY) inceleyin:

-   Desteklenen sürümler
-   Raporlama süreci
-   Yanıt zaman çizelgesi
-   Güvenlik en iyi uygulamaları

## Destek

-   📖 [Belgeler](docs/)
-   🐛 [Sorun İzleyici](https://github.com/lazygophers/log/issues)
-   💬 [Tartışmalar](https://github.com/lazygophers/log/discussions)
-   📧 E-posta: support@lazygophers.com

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır - ayrıntılar için [LICENSE](/tr/LICENSE) dosyasına bakın.

---

## 🌍 Çok Dilli Belgeler

Bu değişiklik günlüğü birçok dilde mevcuttur:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CHANGELOG.md)
-   [🇨🇳 简体中文](/zh-CN/CHANGELOG)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CHANGELOG.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CHANGELOG.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CHANGELOG.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CHANGELOG.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CHANGELOG.md)

---

**Her iyileştirmeyi takip edin ve LazygoPHers Log'un gelişiminden haberdar olun! 🚀**

---

_Bu değişiklik günlüğü her yayında otomatik olarak güncellenir. En son bilgiler için [GitHub Releases](https://github.com/lazygophers/log/releases) sayfasını kontrol edin._
