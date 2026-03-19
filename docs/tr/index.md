---
pageType: home

hero:
    name: LazyGophers Log
    text: Yüksek performanslı ve esnek Go günlük kütüphanesi
    tagline: Zap üzerine kurulmuş, zengin özellikler ve basit API sunar
    actions:
        - theme: brand
          text: Hızlı Başlangıç
          link: /API
        - theme: alt
          text: GitHub'da Görüntüle
          link: https://github.com/lazygophers/log

features:
    - title: "Yüksek Performans"
      details: Zap üzerine kurulmuş, nesne havuzu yeniden kullanımı ve koşullu alan kaydırmayı içerir
      icon: 🚀
    - title: "Zengin Günlük Seviyeleri"
      details: Trace, Debug, Info, Warn, Error, Fatal, Panic seviyelerini destekler
      icon: 📊
    - title: "Esnek Yapılandırma"
      details: Günlük seviyeleri, arayan bilgileri, izleme, önek, sonek ve çıktı hedeflerini özelleştirin
      icon: ⚙️
    - title: "Dosya Döndürme"
      details: Saatlik günlük dosyası döndürme için yerleşik destek
      icon: 🔄
    - title: "Zap Uyumluluğu"
      details: Zap WriteSyncer ile sorunsuz entegrasyon
      icon: 🔌
    - title: "Basit API"
      details: Standart günlük kütüphanesine benzer net API, kullanımı ve entegrasyonu kolay
      icon: 🎯
---

## Hızlı Başlangıç

### Kurulum

```bash
go get github.com/lazygophers/log
```

### Temel Kullanım

```go
package main

import (
    "github.com/lazygophers/log"
)

func main() {
    // Varsayılan global logger kullan
    log.Debug("Hata ayıklama bilgisi")
    log.Info("Genel bilgiler")
    log.Warn("Uyarı bilgisi")
    log.Error("Hata bilgisi")

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

## Dokümantasyon

-   [API Referansı](API.md) - Tam API belgelemesi
-   [Değişiklik Günlüğü](/tr/CHANGELOG) - Sürüm geçmişi
-   [Katkı Kılavuzu](/tr/CONTRIBUTING) - Nasıl katkı yapılır
-   [Güvenlik Politikası](/tr/SECURITY) - Güvenlik kılavuzu
-   [Davranış Kuralları](/tr/CODE_OF_CONDUCT) - Topluluk yönergeleri

## Performans Karşılaştırması

| Özellik       | lazygophers/log | zap | logrus | Standart günlük |
| ---------- | --------------- | --- | ------ | -------- |
| Performans       | Yüksek              | Yüksek  | Orta     | Düşük       |
| API Basitliği    | Yüksek              | Orta  | Yüksek     | Yüksek       |
| Özellik Zenginliği    | Orta              | Yüksek  | Yüksek     | Düşük       |
| Esneklik      | Orta              | Yüksek  | Yüksek     | Düşük       |
| Öğrenme Eğrisi      | Düşük              | Orta  | Orta     | Düşük       |

## Lisans

Bu proje MIT Lisansı altında lisanslanmıştır - detaylar için [LICENSE](/tr/LICENSE) dosyasına bakın.
