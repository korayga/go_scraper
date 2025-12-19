package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/chromedp/chromedp"
)

func main() {
	// Komut satırından URL kontrolü
	if len(os.Args) < 2 {
		fmt.Println("Kullanım: go run main.go <URL>")
		return
	}
	url := os.Args[1]
	scraping(url)
}

func scraping(url string) {

	// Chrome tarayıcı ayarları
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
		chromedp.Flag("headless", true),                                
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // Bot tespitini engelle
		chromedp.Flag("exclude-switches", "enable-automation"),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-features", "IsolateOrigins,site-per-process"),
		chromedp.WindowSize(1920, 1080),
	)

	// Chrome context oluştur
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 90 saniye timeout
	ctx, cancel = context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	// Değişkenler
	var htmlContent string        // Sayfanın HTML içeriği
	var buf []byte                // Screenshot buffer
	var links []map[string]string // Sayfadaki tüm linkler

	fmt.Println("İstek gönderiliyor:", url)

	// Chromedp komutları
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),                                         // Sayfayı aç
		chromedp.Sleep(10*time.Second),                                 // Sayfa yüklensin
		chromedp.OuterHTML("html", &htmlContent),                       // HTML içeriğini al
		chromedp.FullScreenshot(&buf, 90),                              // Screenshot al
		chromedp.AttributesAll(`a[href]`, &links, chromedp.ByQueryAll), // Linkleri topla
	)

	if err != nil {
		fmt.Println("Chromedp hatası:", err)
		return
	}

	// HTML'i dosyaya kaydet
	htmlFilename := fmt.Sprintf("html_%d.html", time.Now().Unix())
	err = os.WriteFile(htmlFilename, []byte(htmlContent), 0644)
	if err != nil {
		fmt.Println("HTML kaydetme hatası:", err)
		return
	}
	fmt.Println("HTML kaydedildi:", htmlFilename)

	// Screenshot'ı dosyaya kaydet
	screenshotFilename := fmt.Sprintf("screenshot_%d.png", time.Now().Unix())
	err = os.WriteFile(screenshotFilename, buf, 0644)
	if err != nil {
		fmt.Println("Screenshot kaydetme hatası:", err)
		return
	}
	fmt.Println("Screenshot kaydedildi:", screenshotFilename)

	// URL'leri temizle ve filtrele
	fmt.Println("Bulunan Tam URL'ler")

	var cleanURLs []string
	seen := make(map[string]bool) // Aynı linki bir kere olması için

	for _, attrs := range links {
		href := attrs["href"]

		// Boşlukları temizle
		href = strings.TrimSpace(href)

		if !strings.HasPrefix(href, "http://") && !strings.HasPrefix(href, "https://") {
			continue
		}
		// Duplicate kontrolü - aynı link varsa geç
		if seen[href] {
			continue
		}

		// Temiz listeye ekle
		cleanURLs = append(cleanURLs, href)
		seen[href] = true
	}

	// URL'leri dosyaya kaydet
	urlsFilename := fmt.Sprintf("urls_%d.txt", time.Now().Unix())
	urlsFile, err := os.Create(urlsFilename)
	if err != nil {
		fmt.Println("URL dosyası oluşturma hatası:", err)
		return
	}
	defer urlsFile.Close()

	// Her bir URL'yi ekrana yazdır ve dosyaya yaz
	for i, href := range cleanURLs {
		// Ekrana yazdır
		fmt.Println(i+1, ".", href)
		// Dosyaya yaz
		fmt.Fprintln(urlsFile, href)
	}

	fmt.Println("\nToplam", len(cleanURLs), "URL bulundu")
	fmt.Println("URL listesi kaydedildi:", urlsFilename)
}
