---
titleSuffix: ' | LazyGophers Log'
---
# 🤝 LazyGophers Log'a Katkıda Bulunma

Katkılarınızı memnuniyetle karşılıyoruz! LazyGophers Log'a katkıda bulunmayı mümkün olduğunca basit ve şeffaf hale getirmek istiyoruz, ister:

-   🐛 Hata bildirin
-   💬 Kodun mevcut durumunu tartışın
-   ✨ Özellik isteyin
-   🔧 Düzeltme önerin
-   🚀 Yeni özellikler uygulayın

## 📋 İçindekiler

-   [Davranış Kuralları](#-davranış-kuralları)
-   [Geliştirme Süreci](#-geliştirme-süreci)
-   [Başlarken](#-başlarken)
-   [Çekme İsteği Süreci](#-çekme-isteği-süreci)
-   [Kodlama Standartları](#-kodlama-standartları)
-   [Test Kılavuzu](#-test-kılavuzu)
-   [Derleme Etiketi Gereksinimleri](#️-derleme-etiketi-gereksinimleri)
-   [Belgelendirme](#-belgelendirme)
-   [Soru Kılavuzu](#-soru-kılavuzu)
-   [Performans Göz Önünde Bulundurma](#-performans-göz-önünde-bulundurma)
-   [Güvenlik Kılavuzu](#-güvenlik-kılavuzu)
-   [Topluluk](#-topluluk)

## 📜 Davranış Kuralları

Bu proje ve tüm katılımcıları [Davranış Kurallarımız](/tr/CODE_OF_CONDUCT) ile bağlıdır. Katılarak bu kurallara uymayı kabul edersiniz.

## 🔄 Geliştirme Süreci

Kodu barındırmak, sorunları ve özellik isteklerini takip etmek ve çekme isteklerini kabul etmek için GitHub kullanıyoruz.

### İş Akışı

:::note Geliştirme süreci genel bakış
1. **Fork** yapın deposu
2. **Clone** yapın forkunuzu yerel olarak
3. **Create** yapın `master` dalından özellik dalı oluşturun
4. **Make** değişikliklerinizi yapın
5. **Test** tüm derleme etiketleri altında kapsamlı test yapın
6. **Submit** çekme isteği gönderin
:::

## 🚀 Başlarken

### Ön Koşullar

-   **Go 1.21+** - [Go'yu yükleyin](https://golang.org/doc/install)
-   **Git** - [Git'i yükleyin](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
-   **Make** (isteğe bağlı ancak önerilir)

### Yerel Geliştirme Kurulumu

```bash title="Depoyu klonlayın ve geliştirme ortamını kurun"
# 1. GitHub'da fork yapın
# 2. Forkunuzu klonlayın
git clone https://github.com/YOUR_USERNAME/log.git
cd log

# 3. Yukarı akış uzak deposunu ekleyin
git remote add upstream https://github.com/lazygophers/log.git

# 4. Bağımlılıkları yükleyin
go mod tidy

# 5. Kurulumu doğrulayın
make test-quick
```

### Ortam Kurulumu

:::info Ortam Yapılandırması
Go ortam değişkenlerinin doğru yapılandırıldığından ve en iyi geliştirme deneyimi için önerilen geliştirme araçlarının yüklendiğinden emin olun.
:::

```bash title="Ortam kurulumu"
# Go ortamını ayarlayın (henüz yapılmadıysa)
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# İsteğe bağlı: Yararlı araçları yükleyin
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 📨 Çekme İsteği Süreci

### Gönderimden Önce

1.  **Ara** mevcut PR'leri tekrarları önlemek için
2.  **Test** tüm derleme yapılandırmaları altında değişikliklerinizi
3.  **Belgeleyin** herhangi bir kırıcı değişikliği
4.  **Güncelleyin** ilgili belgeleri
5.  **Ekleyin** yeni özellikler için testler

### PR Kontrol Listesi

:::warning PR göndermeden önce aşağıdaki tüm öğeleri kontrol edin
Kontrol listesi gereksinimlerini karşılamayan PR'lar birleştirilmeyecektir.
:::

-   [ ] **Kod Kalitesi**

    -   [ ] Kod proje stil kılavuzuna uyar
    -   [ ] Yeni lint uyarıları yok
    -   [ ] Doğru hata işleme
    -   [ ] Verimli algoritmalar ve veri yapıları

-   [ ] **Testler**

    -   [ ] Tüm mevcut testler geçiyor: `make test`
    -   [ ] Yeni özellikler için yeni testler eklendi
    -   [ ] Test kapsamı korundu veya artırıldı
    -   [ ] Tüm derleme etiketleri test edildi: `make test-all`

-   [ ] **Belgelendirme**

    -   [ ] Kodda uygun yorumlar var
    -   [ ] API belgeleri güncellendi (gerekirse)
    -   [ ] README güncellendi (gerekirse)
    -   [ ] Çok dilli belgeler güncellendi (kullanıcıya yönelikse)

-   [ ] **Derleme Uyumluluğu**
    -   [ ] Varsayılan mod: `go build`
    -   [ ] Hata ayıklama modu: `go build -tags debug`
    -   [ ] Yayın modu: `go build -tags release`
    -   [ ] Atma modu: `go build -tags discard`
    -   [ ] Kombine modlar test edildi

### PR Şablonu

Çekme isteği gönderirken lütfen [PR şablonumuzu](https://github.com/lazygophers/log/blob/main/.github/pull_request_template.md) kullanın.

## 📏 Kodlama Standartları

### Go Stil Kılavuzu

:::tip Go kod standartları
Standart Go stil kılavuzunu takip ediyoruz ve bazı eklemelerimiz var. Lütfen kod formatlamasının `go fmt` ve `goimports` kontrollerinden geçtiğinden emin olun.
:::

```go
// ✅ Good
func (l *Logger) Info(v ...any) {
    if l.level > InfoLevel {
        return
    }
    l.log(InfoLevel, fmt.Sprint(v...))
}

// ❌ Bad
func (l *Logger) Info(v ...any){
    if l.level>InfoLevel{
        return
    }
    l.log(InfoLevel,fmt.Sprint(v...))
}
```

### Adlandırma Kuralları

-   **Paket**: Kısa, küçük harf, mümkünse tek kelime
-   **Fonksiyon**: CamelCase, tanımlayıcı
-   **Değişken**: Yerel değişkenler için camelCase, dışa aktarılanlar için CamelCase
-   **Sabitler**: Dışa aktarılanlar için CamelCase, dışa aktarılmayanlar için camelCase
-   **Arabirim**: Genellikle "er" ile biter (örneğin `Writer`, `Formatter`)

### Kod Organizasyonu

```
project/
├── docs/           # Çok dilli belgeler
├── .github/        # GitHub şablonları ve iş akışları
├── logger.go       # Ana logger uygulaması
├── entry.go        # Log girişi yapısı
├── level.go        # Log seviyeleri
├── formatter.go    # Log formatlama
├── output.go       # Çıktı yönetimi
└── *_test.go      # Kaynak kodla birlikte testler
```

### Hata İşleme

:::tip Hata işleme en iyi uygulamalar
Kütüphane kodu hata dönmeli, panic atmamalıdır, çağıranın nasıl yanıt vereceğine karar vermesine izin vermelidir.
:::

```go title="Hata işleme örneği"
// ✅ Önerilir: Hata döndürün, panic yapmayın
func NewLogger(config Config) (*Logger, error) {
    if err := config.Validate(); err != nil {
        return nil, fmt.Errorf("invalid config: %w", err)
    }
    return &Logger{...}, nil
}

// ❌ Kaçının: Kütüphane kodunda panic kullanın
func NewLogger(config Config) *Logger {
    if err := config.Validate(); err != nil {
        panic(err) // Bunu yapmayın
    }
    return &Logger{...}
}
```

## 🧪 Test Kılavuzu

### Test Yapısı

```go title="Tablo sürücülü test örneği"
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
            // Test implementation
        })
    }
}
```

### Kapsam Gereksinimleri

:::warning Kapsam sert gereksinimleri
Yeni kod kapsamı %90'ın altında olan PR'lar CI kontrolünü geçmeyecektir.
:::

-   **Minimum gereksinim**: Yeni kod kapsamı %90
-   **Hedef**: Genel kapsam %95+
-   **Tüm derleme etiketleri** kapsamı korumalıdır
-   Doğrulama için `make coverage-all` kullanın

### Test Komutları

```bash title="Testleri çalıştırın"
# Tüm derleme etiketleri altında hızlı test
make test-quick

# Kapsamlı test paketi, kapsam dahil
make test-all

# Kapsam raporu
make coverage-html

# Kıyaslamalar
make benchmark
```

## 🏗️ Derleme Etiketi Gereksinimleri

:::warning Derleme uyumluluğu
Tüm değişiklikler derleme etiketi sistemimizle uyumlu olmalıdır, tüm derleme etiketi testlerini geçmeyen kod birleştirilmeyecektir.
:::

### Desteklenen Derleme Etiketleri

-   **Varsayılan** (`go build`): Tam işlevsellik
-   **Hata ayıklama** (`go build -tags debug`): Geliştirilmiş hata ayıklama özellikleri
-   **Yayın** (`go build -tags release`): Üretim optimizasyonları
-   **Atma** (`go build -tags discard`): Maksimum performans

### Derleme Etiketi Testleri

:::info Derleme etiketi açıklaması
Proje koşullu derleme için derleme etiketleri kullanır, farklı etiketler farklı çalışma modlarına karşılık gelir. Göndermeden önce tüm etiketler altında test ettiğinizden emin olun.
:::

```bash title="Derleme etiketi testleri"
# Her derleme yapılandırmasını test edin
make test-default
make test-debug
make test-release
make test-discard

# Kombine testler
make test-debug-discard
make test-release-discard

# Hepsi bir arada test
make test-all
```

### Derleme Etiketi Kılavuzu

```go
//go:build debug
// +build debug

package log

// Hata ayıklamaya özgü uygulama
```

## 📚 Belgeler

### Kod Belgeleri

-   **Tüm dışa aktarılan fonksiyonlar** açık yorumlara sahip olmalıdır
-   **Karmaşık algoritmalar** açıklama gerektirir
-   **Örnekler** sıradan kullanım için kullanılır
-   **İş parçacığı güvenliği** notları (uygunsa)

```go
// SetLevel sets the minimum logging level.
// Logs below this level will be ignored.
// This method is thread-safe.
//
// Example:
//   logger.SetLevel(log.InfoLevel)
//   logger.Debug("ignored")  // Won't output
//   logger.Info("visible")   // Will output
func (l *Logger) SetLevel(level Level) *Logger {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
    return l
}
```

### README Güncellemeleri

Özellik eklerken, lütfen güncelleyin:

-   Ana README.md
-   `docs/` içindeki tüm dillere özgü README'ler
-   Kod örnekleri
-   Özellik listesi

## 🐛 Sorun Kılavuzu

### Hata Raporları

[Hata raporu şablonunu](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/bug_report.md) kullanın ve şunları ekleyin:

-   **Açık sorun açıklaması**
-   **Yeniden oluşturma adımları**
-   **Beklenen ve gerçek davranış**
-   **Ortam ayrıntıları** (işletim sistemi, Go sürümü, derleme etiketleri)
-   **Minimum kod örneği**

### Özellik İstekleri

[Özellik isteme şablonunu](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/feature_request.md) kullanın ve şunları ekleyin:

-   **Açık özellik motivasyonu**
-   **Önerilen API** tasarımı
-   **Uygulama düşünceleri**
-   **Kırıcı değişiklik analizi**

### Sorular

[Soru şablonunu](https://github.com/lazygophers/log/blob/main/.github/ISSUE_TEMPLATE/question.md) şunlar için kullanın:

-   Kullanım sorunları
-   Yapılandırma yardımı
-   En iyi uygulamalar
-   Entegrasyon rehberliği

## 🚀 Performans Göz Önünde Bulundurma

### Kıyaslamalar

Performans hassas değişiklikleri her zaman kıyaslayın:

```bash title="Kıyaslamaları çalıştırın"
# Kıyaslamaları çalıştırın
go test -bench=. -benchmem

# Öncesi sonrası performansı karşılaştırın
go test -bench=. -benchmem > before.txt
# Değişiklikleri yapın
go test -bench=. -benchmem > after.txt
benchcmp before.txt after.txt
```

### Performans Kılavuzu

:::tip Performans optimizasyonu noktaları
Bu performans hassas bir günlük kütüphanesidir, her değişiklik sıcak yol üzerindeki etkileri göz önünde bulundurmalıdır.
:::

-   **En aza indirin** sıcak yol üzerindeki bellek ayırmaları
-   **Nesne havuzlarını kullanın** sık oluşturulan nesneler için
-   **Erken dönüş** devre dışı bırakılmış log seviyeleri için
-   **Yansımadan kaçının** performans kritik kodda
-   **Optimizasyondan önce profil oluşturun**

### Bellek Yönetimi

```go
// ✅ Önerilir: Nesne havuzu kullanın
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

## 🔒 Güvenlik Kılavuzu

### Hassas Veriler

:::warning Güvenlik dikkat edilmesi gerekenler
Günlüklerde hassas veri sızıntısı ciddi güvenlik olaylarına yol açabilir, lütfen aşağıdaki standartlara uyun.
:::

-   **Asla günlüğe mayın** şifreler, jetonlar veya hassas veriler
-   **Temizleyin** günlük mesajlarındaki kullanıcı girdilerini
-   **Kaçının** tüm istek/yanıt gövdesini günlüğe kaydetmekten
-   **Kullanın** daha iyi kontrol için yapılandırılmış günlük kaydı

```go
// ✅ Önerilir
logger.Info("User login attempt", "user_id", userID, "ip", clientIP)

// ❌ Kaçının
logger.Infof("User login: %+v", userRequest) // Şifre içerebilir
```

### Bağımlılıklar

-   Bağımlılıkları **güncel tutun**
-   **Yeni bağımlılıkları dikkatlice inceleyin**
-   **En aza indirin** dış bağımlılıkları
-   **Kullanın** `go mod verify` bütünlüğü kontrol etmek için

## 👥 Topluluk

### Yardım Alın

-   📖 [Belgeler](README.md)
-   💬 [GitHub Tartışmalar](https://github.com/lazygophers/log/discussions)
-   🐛 [Soru İzleyici](https://github.com/lazygophers/log/issues)
-   📧 E-posta: support@lazygophers.com

### İletişim Kılavuzu

-   **Saygılı ve kapsayıcı kalın**
-   **Sormadan önce arayın**
-   **Yardım istediğinizde bağlam sağlayın**
-   **Gücünüz yettiğinde başkalarına yardımcı olun**
-   **Uyun** [Davranış Kuralları](/tr/CODE_OF_CONDUCT)

## 🎯 Takdir

Katkıda bulunanlar birkaç yolla takdir edilir:

-   **README katkıda bulunanlar** bölümü
-   **Yayın notları** içerisinde atıflar
-   **GitHub katkıda bulunanlar** grafiği
-   **Topluluk teşekkür** gönderileri

## 📝 Lisans

Katkıda bulunarak, katkılarınızın MIT lisansı altında lisanslandığını kabul edersiniz.

---

## 🌍 Çok Dilli Belgeler

Bu belge birden fazla dilde sunulmaktadır:

-   [🇺🇸 English](https://lazygophers.github.io/log/en/CONTRIBUTING.md)
-   [🇨🇳 简体中文](https://lazygophers.github.io/log/zh-CN/CONTRIBUTING.md)
-   [🇹🇼 繁體中文](https://lazygophers.github.io/log/zh-TW/CONTRIBUTING.md)
-   [🇫🇷 Français](https://lazygophers.github.io/log/fr/CONTRIBUTING.md)
-   [🇷🇺 Русский](https://lazygophers.github.io/log/ru/CONTRIBUTING.md)
-   [🇪🇸 Español](https://lazygophers.github.io/log/es/CONTRIBUTING.md)
-   [🇸🇦 العربية](https://lazygophers.github.io/log/ar/CONTRIBUTING.md)
-   [🇹🇷 Türkçe](/tr/CONTRIBUTING) (güncel)

---

**LazyGophers Log'a katkıda bulunduğunuz için teşekkürler!🚀**
