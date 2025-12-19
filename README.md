#  Go-Chromedp Web Scraper

![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=for-the-badge&logo=go)
![Chromedp](https://img.shields.io/badge/Chromedp-v0.14.2-4CAF50?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey?style=for-the-badge)

> **Siber Tehdit İstihbaratı (CTI)** süreçleri için geliştirilmiş basit, scraping aracıdır.


## 🌟 Özellikler

### 📸 Full-Page Screenshot
- Sayfanın **tamamını** %90 kalitede PNG formatında kaydetme
- Sadece görünen kısmı değil, tüm scroll edilebilir alanı yakalar
- 1920x1080 çözünürlük desteği

### ⚡ Dinamik İçerik Desteği
- JavaScript (React, Vue, Angular, Next.js) ile yüklenen içerikleri tam olarak yakalar
- 10 saniye sayfa yükleme bekleme süresi
- Lazy-loaded elementleri destekler

### 🧹 Akıllı URL Filtreleme
- Sayfadaki tüm `<a href>` etiketlerini tarar
- **Sadece tam URL'leri** (http/https ile başlayanları) raporlar
- Duplicate (tekrarlayan) URL'leri otomatik kaldırır


### 💾 Otomatik Arşivleme
- Çıktıları UNIX timestamp ile adlandırır
- Aynı siteyi farklı zamanlarda taradığınızda verilerin üzerine yazılmaz
- HTML, Screenshot ve URL listesini ayrı dosyalarda saklar

### 🔒 Hata Yönetimi
- Timeout koruması (90 saniye)
- Dosya yazma hatalarını yakalar
- Chrome hatalarını loglayıp güvenli şekilde sonlanır

---

## 🛠️ Kurulum

### Ön Gereksinimler

- **Go 1.25.5** veya üzeri ([İndir](https://golang.org/dl/))
- **Chrome/Chromium** tarayıcı (Chromedp otomatik indirir)

### Kurulum Adımları

```bash
# 1. Projeyi klonlayın
git clone https://github.com/kullaniciadiniz/go-scraper.git
cd go-scraper

# 2. Bağımlılıkları yükleyin
go mod download

# 3. Çalıştırın (opsiyonel test)
go run main.go https://example.com
```

**Alternatif: Binary Oluşturma**

```bash
# Executable oluşturun
go build -o scraper.exe main.go

# Windows'ta
scraper.exe https://example.com

# Linux/macOS'ta
chmod +x scraper
./scraper https://example.com
```

---

## 🚀 Kullanım

### Temel Kullanım

```bash
go run main.go <HEDEF_URL>
```



## 📦 Çıktı Dosyaları

Program her çalıştırıldığında **3 adet dosya** oluşturur:

### 1. HTML Dosyası
```
📄 html_1734623456.html
```
- Sayfanın **tam HTML kaynak kodu**
- JavaScript render sonrası içerik
- Dosya boyutu: ~50KB - 5MB (siteye göre değişir)

### 2. Screenshot Dosyası
```
🖼️ screenshot_1734623456.png
```
- **Tam sayfa ekran görüntüsü** (PNG formatı)
- Kalite: %90
- Çözünürlük: 1920x1080 viewport
- Dosya boyutu: ~200KB - 3MB

### 3. URL Listesi
```
🔗 urls_1734623456.txt
```
- Sayfadaki **tüm geçerli tam URL'ler**
- Her satırda bir URL
- Temizlenmiş ve filtrelenmiş

---

## 🔧 Teknik Detaylar

### Kullanılan Teknolojiler

| Kütüphane/Paket | Versiyon | Kullanım Amacı |
|----------------|----------|----------------|
| `chromedp/chromedp` | v0.14.2 | Headless Chrome otomasyonu |
| `chromedp/cdproto` | latest | Chrome DevTools Protocol |
| `context` | std | Timeout ve iptal yönetimi |
| `fmt` | std | Formatlı çıktı |
| `os` | std | Dosya sistemi işlemleri |
| `strings` | std | String manipülasyonu |
| `time` | std | Zaman işlemleri |

---

### Chrome Bayrakları (Flags)

Program şu Chrome ayarlarını kullanır:

```go
chromedp.Flag("headless", true)                                 // Pencere açmadan çalışır
chromedp.Flag("disable-blink-features", "AutomationControlled") // Bot tespitini gizler
chromedp.Flag("exclude-switches", "enable-automation")          // Otomasyon işaretini kaldırır
chromedp.Flag("no-sandbox", true)                               // Linux sandbox bypass
chromedp.Flag("disable-gpu", false)                             // GPU hızlandırması aktif
chromedp.Flag("disable-web-security", true)                     // CORS sorunlarını önler
chromedp.WindowSize(1920, 1080)                                 // Full HD viewport
```


---

## 🛠️ Sorun Giderme

### ❌ "Chromedp hatası: context deadline exceeded"

**Sebep:** Site 90 saniyede yüklenemedi  
**Çözüm:**
```go
// main.go'da timeout süresini artırın
ctx, cancel = context.WithTimeout(ctx, 180*time.Second) // 180 saniye yap
```

### ❌ "ERR_NAME_NOT_RESOLVED"

**Sebep:** URL geçersiz veya internet bağlantısı yok  
**Çözüm:**
- URL'nin doğruluğunu kontrol edin
- İnternet bağlantınızı test edin
- DNS ayarlarınızı kontrol edin (8.8.8.8)


### ❌ Cloudflare "Checking your browser" sonsuz loop

**Sebep:** Headless mod Cloudflare tarafından algılandı  
**Çözüm:**
```go
// main.go'da headless'ı kapat
chromedp.Flag("headless", false), // true → false
```

### ❌ "Permission denied" hatası

**Sebep:** Klasöre yazma izni yok  
**Çözüm:**
```bash
# Linux/macOS
sudo chmod 777 .

# Windows (Administrator PowerShell)
icacls . /grant Users:F
```

---


### 🔴 Yasal Uyarı

> **UYARI:** Bu araç **sadece eğitim, araştırma ve yasal penetrasyon testleri** için geliştirilmiştir. 

### 🛡️ Güvenli Kullanım İpuçları

1. **VPN Kullanın:** IP adresinizi korumak için
2. **Rate Limiting:** Aynı siteyi sürekli taramayın (5-10 dakika arayla)
3. **robots.txt Saygısı:** Sitelerin tarama kurallarını kontrol edin
4. **Test Ortamı:** İlk testleri kendi sitenizde yapın

---

## 📞 İletişim

- **GitHub Issues:** [Sorun bildir](https://github.com/korayga/go-scraper/issues)
- **Linkedin:** [korayga](https://www.linkedin.com/in/koray-garip/)

---

