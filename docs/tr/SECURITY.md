---
titleSuffix: ' | LazyGophers Log'
---
# 🔒 Güvenlik Politikası

## Güvenlik Taahhüdümüz

LazyGophers Log güvenliği çok ciddiye alır. Günlük kütüphanemiz için en yüksek güvenlik standartlarını sürdürmeye, kullanıcılarımızın uygulamalarının güvenliğini korumaya kararlıyız. Güvenlik açıklarını sorumlu bir şekilde ifşa etme çabalarınızı takdir ediyoruz ve güvenlik topluluğuna katkınızı tanımak için elimizden gelenin en iyisini yapacağız.

### Güvenlik İlkeleri

-   **Güvenlik by Tasarım**: Güvenlik düşünceleri geliştirme sürecinin her yönüne dahil edilir
-   **Şeffaflık**: Güvenlik sorunları ve düzeltmeleri hakkında açık iletişim sürdürürüz
-   **Topluluk İşbirliği**: Güvenlik araştırmacıları ve kullanıcılarla işbirliği yapıyoruz
-   **Sürekli İyileştirme**: Güvenlik uygulamalarını düzenli olarak gözden geçiriyor ve geliştiriyoruz

## Desteklenen Sürümler

Aşağıdaki LazyGophers Log sürümleri için aktif olarak güvenlik güncellemeleri sağlıyoruz:

| Sürüm  | Destek Durumu | Durum   | Yaşam Sonu | Notlar           |
| ----- | -------- | ------ | ------------ | -------------- |
| 1.x.x | ✅ Evet    | Aktif   | TBD         | Tam güvenlik desteği   |
| 0.9.x | ✅ Evet    | Bakım   | 2024-06-01   | Sadece kritik güvenlik düzeltmeleri |
| 0.8.x | ⚠️ Sınırlı  | Eski   | 2024-03-01   | Sadece acil düzeltmeler     |
| 0.7.x | ❌ Hayır    | Kullanımdan Kaldırıldı | 2024-01-01   | Güvenlik desteği yok     |
| < 0.7 | ❌ Hayır    | Kullanımdan Kaldırıldı | 2023-12-01   | Güvenlik desteği yok     |

### Destek Politikası Ayrıntıları

:::info Destek Seviyesi Açıklamaları

-   **Aktif**: Tam güvenlik güncellemeleri, düzenli yamalar, proaktif izleme
-   **Bakım**: Sadece kritik ve yüksek güvenlik sorunları
-   **Eski**: Sadece kritik açıklar için acil güvenlik düzeltmeleri
-   **Kullanımdan Kaldırıldı**: Güvenlik desteği yok - kullanıcılar derhal yükseltmeli

:::

### Yükseltme Önerileri

:::warning Sürüm Yükseltme Hatırlatması

-   **Derhal Eylem**: < 0.8.x sürümünü kullanan kullanıcılar derhal 1.x.x'e yükseltmeli
-   **Geçiş Planlayın**: 0.8.x - 0.9.x sürümlerini kullanan kullanıcılar yaşam sonundan önce 1.x.x'e geçişi planlamalı
-   **Güncel Kalın**: En iyi güvenlik için her zaman en son kararlı sürümü kullanın

:::

## 🐛 Güvenlik Açıklarını Bildirme

:::danger Güvenlik Açıklarını Kamu Kanalları aracılığıyla Bildirmeyin

Lütfen güvenlik açıklarını şu yollarla bildirmeyin:

-   Genel GitHub issues
-   Genel tartışmalar
-   Sosyal medya
-   Posta listeleri
-   Topluluk forumları

:::

### Güvenlik Raporlama Kanalları

:::info Güvenlik Açığı Bildirme Kanalları

Bir güvenlik açığı bildirmek için lütfen aşağıdaki güvenli kanallardan birini kullanın:

#### Tercih Edilen İletişim Yöntemi

-   **E-posta**: security@lazygophers.com
-   **PGP Anahtarı**: İstek üzerine mevcuttur
-   **Konu**: `[SECURITY] Güvenlik Açığı Bildirimi - LazyGophers Log`

#### GitHub Güvenlik Tavsiyeleri

-   [GitHub Güvenlik Tavsiyelerimize](https://github.com/lazygophers/log/security/advisories) ziyaret edin
-   "New draft security advisory"ye tıklayın
-   Açık hakkında ayrıntılı bilgi sağlayın

#### Yedek İletişim Yöntemi

-   **E-posta**: support@lazygophers.com (gizli güvenlik sorunu olarak işaretleyin)

:::

### Rapor İçerik Gereksinimleri

Lütfen güvenlik açığı raporunuzda aşağıdaki bilgileri ekleyin:

#### Temel Bilgiler

-   **Özet**: Açığın kısa bir açıklaması
-   **Etki**: Potansiyel etki ve ciddiyet değerlendirmesi
-   **Yeniden Oluşturma Adımları**: Sorunu yeniden oluşturmak için ayrıntılı adımlar
-   **Kavram Kanıtı**: Açığı gösteren kod veya adımlar
-   **Etkilenen Sürümler**: Etkilenen belirli sürümler veya sürüm aralıkları
-   **Ortam**: İşletim sistemi, Go sürümü, kullanılan derleme etiketleri

#### İsteğe Bağlı ama Yararlı Bilgiler

-   **CVSS Puanı**: Hesaplayabilirseniz
-   **CWE Referansı**: Common Weakness Enumeration referansı
-   **Önerilen Düzeltme**: Düzeltme fikriniz varsa
-   **Zaman Çizelgesi**: Tercih edilen ifşa zaman çizelgeniz

### Rapor Şablonu Örneği

```markdown title="Güvenlik Raporu Şablonu"
Konu: [SECURITY] Günlük Formatlayıcısında Arabellek Taşması

Özet:
Günlük formatlayıcısı, aşırı uzun günlük mesajlarını işlerken bir arabellek taşması açığına sahip.

Etki:
- Potansiyel keyfi kod yürütme
- Bellek hasarı
- Hizmet reddi

Yeniden Oluşturma Adımları:
1. Bir günlükçü örneği oluşturun
2. 10.000 karakterden uzun bir mesaj günlükleyin
3. Bellek hasarını gözlemleyin

Etkilenen Sürümler:
- v1.0.0 ile v1.2.3 arası

Ortam:
- İşletim Sistemi: Ubuntu 20.04
- Go: 1.21.0
- Derleme Etiketleri: release

Kavram Kanıtı:
[Minimal kod örneği ekleyin]
```

## 📋 Güvenlik Yanıt Süreci

### Yanıt Zaman Çizelgemiz

| Zaman Çerçevesi | Eylem |
| -------- | ---- |
| 24 saat  | Rapor alındığının ilk bildirimi |
| 72 saat  | İlk değerlendirme ve sınıflandırma |
| 1 hafta     | Detaylı soruşturma başlangıcı |
| 2-4 hafta   | Düzeltme geliştirme ve test |
| 4-6 hafta   | Koordine ifşa ve yayın |

### Yanıt Süreci Adımları

#### 1. Bildirim (24 saat)

-   Güvenlik açığı raporunu alın
-   İzleme numarası atayın
-   Eksik bilgileri isteyin

#### 2. Değerlendirme (72 saat)

-   İlk ciddiyet değerlendirmesi
-   Etkilenen sürümlerin belirlenmesi
-   Etki analizi
-   CVSS puanı atama

#### 3. Soruşturma (1 hafta)

-   Detaylı teknik analiz
-   Kök neden tanımlama
-   Sömteri senaryosu analizi
-   Düzeltme stratejisi planlama

#### 4. Geliştirme (2-4 hafta)

-   Güvenlik yaması geliştirme
-   İç test
-   Desteklenen tüm sürümlerde regresyon testleri
-   Belge güncellemeleri

#### 5. İfşa (4-6 hafta)

-   Raporlayanla ifşa zaman çizelgesini koordine edin
-   Güvenlik duyurusunu hazırlayın
-   Yamalı sürümü yayınlayın
-   Kamu ifşası

### Ciddiyet Sınıflandırması

Aşağıdaki ciddiyet sınıflandırma standartlarını kullanıyoruz:

#### 🔴 Kritik (CVSS 9.0-10.0)

-   Gizlilik, bütünlük veya kullanılabilirliğe doğrudan tehdit
-   Uzaktan kod yürütme
-   Tam sistem kompromisi
-   **Yanıt**: 72 saat içinde acil yama

#### 🟠 Yüksek (CVSS 7.0-8.9)

-   Önemli güvenlik etkisi
-   Ayrıcalık yükseltme
-   Veri sızıntısı
-   **Yanıt**: 1-2 hafta içinde yama

#### 🟡 Orta (CVSS 4.0-6.9)

-   Orta güvenlik etkisi
-   Sınırlı veri sızıntısı
-   Kısmi sistemi kompromisi
-   **Yanıt**: 1 ay içinde yama

#### 🟢 Düşük (CVSS 0.1-3.9)

-   Daha az güvenlik etkisi
-   Bilgi sızıntısı
-   Sınırlı kapsam açığı
-   **Yanıt**: Sonraki düzenli sürümde düzeltme

### İletişim Tercihleri

#### Sizden Beklediklerimiz

-   **Sorumlu İfşa**: Sorunu düzeltmek için makul bir süre verin
-   **İletişim İşbirliği**: Sorularımıza ve açıklama taleplerimize yanıt verin
-   **Koordine İşbirliği**: İfşa zamanını birlikte belirlemek için bizimle çalışın
-   **Test Yardımı**: Mümkünse, düzeltmemizi doğrulamamıza yardımcı olun

#### Sizden Bekleyebilecekleriniz

-   **Zamanında Bildirim**: Raporunuzu hızlıca onaylarız
-   **Düzenli Güncellemeler**: Süre boyunca düzenli durum güncellemeleri sağlarız
-   **Kamu Tanıma**: Bulgunuzu kamuoyunda tanırız (anonim kalmak isterseniz hariç)
-   **Saygılı İletişim**: Profesyonel ve saygılı iletişim tarzı

## 🛡️ Güvenlik En İyi Uygulamaları

### Uygulama Geliştiricileri İçin

#### Dağıtım Güvenliği

-   **En Son Sürümü Kullanın**: Her zaman güvenlik yamaları içeren en son desteklenen sürümü kullanın
-   **Duyuruları İzleyin**: Güvenlik posta listemize ve GitHub güvenlik tavsiyelerine abone olun
-   **Güvenli Yapılandırma**: Güvenlik sertleştirme kılavuzumuzu izleyin
-   **Düzenli Güncellemeler**: Kritik sorunlar yayınlandıktan sonra 48 saat içinde güvenlik güncellemelerini uygulayın
-   **Sürüm Kilitleme**: Üretim ortamında sürüm aralıkları yerine belirli sürüm numaraları kullanın
-   **Güvenlik Taraması**: Uygulamanızı ve bağımlılıklarınızı düzenli olarak açıklara karşı tarayın

#### Günlük Güvenliği ve Veri Koruma

:::tip Günlük Güvenliği En İyi Uygulamaları

-   **Hassas Veriler**: Asla şifreler, API anahtarları, tokenler, kişisel kimlik bilgileri veya finansal bilgileri günlüklemeyin
-   **Veri Sınıflandırması**: Günlük içeriği için veri sınıflandırma politikası uygulayın
-   **Girdi Temizleme**: Günlüklemekten önce tüm kullanıcı girdilerini temizleyin ve doğrulayın
-   **Çıktı Kodlama**: Enjeksiyon saldırılarını önlemek için günlük çıktısını doğru şekilde kodlayın
-   **Erişim Kontrolü**: Günlük dosyaları ve dizinleri için katı erişim kontrolü uygulayın
-   **Şifreleme**: Hassas işletme verileri içeren günlük dosyalarını şifreleyin
-   **Saklama Politikası**: Uygun günlük saklama ve silme politikaları uygulayın
-   **Denetim İzi**: Günlük dosyalarına erişim ve değişiklikler için denetim izi sürün

:::

#### Derleme ve Dağıtım Güvenliği

:::tip Güvenli Derleme Kılavuzu

-   **Sağlama Doğrulama**: Her zaman paketlerin sağlama toplamlarını ve imzalarını doğrulayın
-   **Resmi Kaynaklar**: Sadece resmi GitHub sürümlerinden veya Go modül proxy'lerinden indirin
-   **Bağımlılık Yönetimi**: `go mod verify` ve bağımlılık tarama araçları kullanın
-   **Derleme Etiketleri**: Güvenlik ihtiyaçlarınıza göre uygun derleme etiketlerini kullanın:
    -   Üretim: Optimize edilmiş güvenlik derlemeleri için `release` etiketi
    -   Geliştirme: Gelişmiş hata ayıklama için `debug` etiketi (asla üretimde kullanmayın)
    -   Yüksek Güvenlik: Maksimum performans ve最小 saldırı yüzeyi için `discard` etiketi
-   **Tedarik Zinciri Güvenliği**: Tüm bağımlılık zincirinin bütünlüğünü doğrulayın

:::

#### Altyapı Güvenliği

-   **Günlük Toplama**: Uygun kimlik doğrulama ile güvenli günlük toplama sistemleri kullanın
-   **Ağ Güvenliği**: Günlük aktarımının şifreli kanallar (TLS 1.3+) kullandığından emin olun
-   **Depolama Güvenliği**: Günlükleri erişim kontrollü güvenli depolama sistemlerinde saklayın
-   **Yedekleme Güvenliği**: Günlük yedeklerini şifreleyin ve koruyun, uygun saklama süreleri belirleyin

### Katkıda Bulunanlar ve Bakımcılar İçin

#### Güvenli Geliştirme Yaşam Döngüsü

:::note Güvenli Geliştirme Standartları

-   **Tehdit Modelleme**: Günlük kütüphanesinin tehddit modelini düzenli olarak gözden geçirin ve güncelleyin
-   **Güvenlik Gereksinimleri**: Tüm özellik geliştirmesine güvenlik gereksinimlerini entegre edin
-   **Güvenli Kodlama**: Güvenli kodlama uygulamalarını ve OWASP yönergelerini izleyin
-   **Kod Güvenliği**:
    -   **Girdi Doğrulama**: Uygun sınır kontrolleri ile tüm girdileri tamamen doğrulayın
    -   **Arabellek Yönetimi**: Uygun arabellek boyutu yönetimi ve taşma koruması uygulayın
    -   **Hata İşleme**: Bilgi sızıntısını önleyen güvenli hata işleme
    -   **Bellek Güvenliği**: Arabellek taşmalarını, bellek sızıntılarını ve use-after-free hatalarını önleyin
    -   **Eşzamanlılık Güvenliği**: İş parçacığı güvenli işlemleri sağlayın ve yarış koşullarını önleyin

:::

#### Geliştirme Güvenlik Uygulamaları

-   **Güvenlik İncelemeleri**: Tüm değişiklikler güvenlik kod incelemesinden geçmeli
-   **Statik Analiz**: Çeşitli statik analiz araçları kullanın (`gosec`, `staticcheck`, `semgrep`)
-   **Dinamik Test**: Güvenliğe odaklanan dinamik test ve fuzzing ekleyin
-   **Bağımlılık Güvenliği**:
    -   Tüm bağımlılıkları en son güvenli sürümlerde tutun
    -   Düzenli olarak `govulncheck` ve `nancy` ile bağımlılık açığı taraması yapın
    -   Gereksiz bağımlılıklardan kaçınarak bağımlılık yüzeyini minimize edin
-   **Test Etme**:
    -   Kapsamlı güvenlik test durumları ekleyin
    -   Desteklenen tüm derleme etiketleri ve yapılandırmalarda test edin
    -   Sınır testleri ve girdi doğrulama testleri gerçekleştirin
    -   Hizmet reddi açıklarını belirlemek için performans testleri yapın

#### Tedarik Zinciri Güvenliği

-   **Kod İmzalama**: Tüm sürümleri doğrulanmış imzalarla imzalayın
-   **Derleme Süreci**: Tekrarlanabilir derlemeler ve güvenli derleme ortamı kullanın
-   **Yayın Yönetimi**: Uygun onaylarla güvenli yayın sürecini izleyin
-   **Açık İfşa**: Koordine edilmiş açık ifşa sürecini sürdürün

## 📚 Güvenlik Kaynakları

### İç Belgeler

-   [Katkıda Bulunma Kılavuzu](/tr/CONTRIBUTING) - Katkıda bulunanlar için güvenlik hususları
-   [Davranış Kuralları](/tr/CODE_OF_CONDUCT) - Topluluk güvenliği ve refahı
-   [API Belgeleri](API.md) - Güvenli kullanım desenleri ve örnekler
-   [Derleme Yapılandırma Kılavuzu](README.md) - Derleme etiketlerinin güvenlik etkisi

### Dış Güvenlik Standartları ve Çerçeveleri

-   [NIST Siber Güvenlik Çerçevesi](https://www.nist.gov/cyberframework) - Kapsamlı güvenlik çerçevesi
-   [OWASP Top 10](https://owasp.org/www-project-top-ten/) - En kritik web uygulaması güvenlik riskleri
-   [OWASP Günlük Çeşitlemesi](https://cheatsheetseries.owasp.org/cheatsheets/Logging_Cheat_Sheet.html) - Günlük güvenliği en iyi uygulamaları
-   [Go Güvenlik Kontrol Listesi](https://github.com/Checkmarx/Go-SCP) - Go'ya özgü güvenlik kılavuzu
-   [CIS Kontrolleri](https://www.cisecurity.org/controls/) - Kritik güvenlik kontrolleri
-   [ISO 27001](https://www.iso.org/isoiec-27001-information-security.html) - Bilgi güvenliği yönetimi

### Açık Veritabanları ve İstihbarat

-   [Common Vulnerabilities and Exposures (CVE)](https://cve.mitre.org/) - Açık veritabanı
-   [National Vulnerability Database (NVD)](https://nvd.nist.gov/) - ABD hükümeti açık veritabanı
-   [Go Açık Veritabanı](https://pkg.go.dev/vuln/) - Go'ya özgü açıklar
-   [GitHub Güvenlik Tavsiyeleri](https://github.com/advisories) - Açık kaynak güvenlik tavsiyeleri
-   [Snyk Açık Veritabanı](https://snyk.io/vuln/) - Ticari açık istihbaratı

### Güvenlik Araçları ve Tarayıcılar

#### Statik Analiz Araçları

-   **`gosec`**: Go güvenlik kontrolcü - Go kodundaki güvenlik sorunlarını tespit eder
-   **`staticcheck`**: Güvenlik kontrolleri olan gelişmiş Go kod kontrolcüsü
-   **`semgrep`**: Özel güvenlik kurallarıyla çok dilli statik analiz
-   **`CodeQL`**: GitHub'ın güvenlik açıklarını bulmak için anlamsal kod analizi
-   **`nancy`**: Go bağımlılıklarındaki bilinen açıkları kontrol eder

#### Dinamik Analiz ve Test Etme

-   **`govulncheck`**: Resmi Go açık kontrolcüsü
-   **Go Yerleşik Fuzzing**: Güvenlik sorunlarını bulmak için `go test -fuzz`
-   **`dlv` (Delve)**: Güvenlik testi için Go hata ayıklayıcı
-   **Yük Test Araçları**: Hizmet reddi açıklarını belirlemek için

#### Bağımlılık ve Tedarik Zinciri Güvenliği

-   **`go mod verify`**: Bağımlılıkların değiştirilip değiştirilmediğini doğrular
-   **Dependabot**: Otomatik bağımlılık güncellemeleri ve güvenlik uyarıları
-   **Snyk**: Ticari bağımlılık taraması ve izleme
-   **FOSSA**: Lisans uyumluluğu ve açık taraması

#### Kod Kalitesi ve Güvenliği

-   **`golangci-lint`**: Çeşitli güvenlik kontrolcüleri ile hızlı Go kod lint aracı
-   **`goreportcard`**: Go kod kalitesi değerlendirmesi
-   **`gocyclo`**: Döngüsel karmaşıklık analizi
-   **`ineffassign`**: Etkisiz atamaları tespit eder

### Güvenlik Topluluğu ve Kaynakları

#### Go Güvenlik Topluluğu

-   [Go Güvenlik Politikası](https://golang.org/security) - Resmi Go güvenlik politikası
-   [Go Geliştirme Güvenliği](https://groups.google.com/g/golang-dev) - Go geliştirme tartışmaları
-   [Golang Güvenliği](https://github.com/golang/go/wiki/Security) - Go güvenlik wiki'si

#### Genel Güvenlik Topluluğu

-   [OWASP Topluluğu](https://owasp.org/membership/) - Open Web Application Security Project
-   [SANS Enstitüsü](https://www.sans.org/) - Güvenlik eğitimi ve sertifikasyonu
-   [FIRST](https://www.first.org/) - Olay Yanıt ve Güvenlik Ekipleri Forumu
-   [CVE Projesi](https://cve.mitre.org/about/index.html) - Açık ifşa projesi

### Eğitim ve Sertifikasyon

-   **Güvenli Kodlama Eğitimi**: Platforma özgü güvenli kodlama kursları
-   **CISSP**: Certified Information Systems Security Professional
-   **GSEC**: GIAC Security Essentials Certification
-   **CEH**: Certified Ethical Hacker
-   **Go Güvenlik Kursları**: Özelleşmiş Go güvenlik eğitim programları

## 🏆 Güvenlik Şöhret Holü

Projenin güvenliğini iyileştirmeye yardımcı olan güvenlik araştırmacılarını onurlandırmak için bir güvenlik şöhret holü sürdürüyoruz:

### Katkıda Bulunanlar

_Burada sorumlu bir şekilde açık ifşa eden güvenlik araştırmacılarını listeleneceğiz (onlarının izniyle)_

### Tanıma Kriterleri

-   Geçerli güvenlik açıklarının sorumlu ifşası
-   Düzeltme sürecinde yapıcı işbirliği
-   Projenin genel güvenliğine katkı

## 📞 İletişim Bilgileri

### Güvenlik Ekibi

-   **Tercih Edilen**: security@lazygophers.com
-   **Yedek**: support@lazygophers.com
-   **PGP Anahtarı**: İstek üzerine mevcuttur

### Yanıt Ekibi

Güvenlik yanıt ekibimiz şunları içerir:

-   Çekirdek bakımcılar
-   Güvenliğe odaklanan katkıda bulunanlar
-   Dış güvenlik danışmanları (gerekirse)

## 🔄 Politika Güncellemeleri

Bu güvenlik politikası düzenli olarak gözden geçirilir ve güncellenir:

-   Süreç iyileştirmeleri için **üç aylık gözden geçirmeler**
-   Güvenlik olayları için **anında güncellemeler**
-   Kapsamlı politika güncellemeleri için **yıllık gözden geçirmeler**

Son güncelleme: 2024-01-01

---

## 🌍 Çok Dilli Belgeler

Bu belge birçok dilde mevcuttur:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/SECURITY.md)
-   [🇨🇳 简体中文](/zh-CN/SECURITY)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/SECURITY.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/SECURITY.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/SECURITY.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/SECURITY.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/SECURITY.md)

---

**Güvenlik ortak bir sorumluluktur. LazyGophers Log'un güvenliğini korumaya yardımcı olduğunuz için teşekkürler! 🔒**
