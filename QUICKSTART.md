# 🚀 Hızlı Başlangıç

## Tek Komutla Kurulum (Windows)

Proje kök dizininde çift tıklayın:
```
setup.bat
```

Bu script otomatik olarak:
- Go bağımlılıklarını yükler
- NPM paketlerini yükler
- Gerekli klasörleri oluşturur

## Manuel Kurulum

### 1. Go Bağımlılıkları
```bash
go mod download
```

### 2. Frontend Bağımlılıkları
```bash
cd frontend
npm install
cd ..
```

### 3. Dist Klasörü (İlk çalıştırma için)
```bash
mkdir frontend\dist
```

## Çalıştırma

### Geliştirme Modu
```bash
wails dev
```

### Production Build
```bash
wails build
```

Build edilen dosya: `build/bin/StockApp.exe`

## ⚠️ Import Hataları

VSCode'da kırmızı çizgiler görüyorsanız **NORMAL**. Bunlar:
- `go mod download` çalıştırılmadığı için görünür
- `setup.bat` çalıştırdıktan sonra kaybolur

## 🐛 Sorun Giderme

### "Wails bulunamadı" hatası
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### "Pattern all:frontend/dist: no matching files found"
```bash
mkdir frontend\dist
echo. > frontend\dist\.gitkeep
```

### Go import hataları
```bash
go mod tidy
go mod download
```

## 📌 Önemli Notlar

1. **İlk çalıştırma** `wails dev` biraz uzun sürebilir (binding'ler oluşturulur)
2. **Hot reload** aktiftir - hem Go hem Vue değişiklikleri otomatik yansır
3. **Frontend dist klasörü** boş olabilir, sorun değil

## ✅ Başarılı Kurulum Kontrolü

Aşağıdaki komutlar çalışmalı:
```bash
go version      # 1.21+
node --version  # 18+
wails doctor    # Tüm checkler yeşil olmalı
```

## 🎯 Sonraki Adım

Kurulum tamamlandıktan sonra:
```bash
wails dev
```

Uygulama açılacak ve database seçici ekranı görünecek!

---

**Yardım gerekirse**: `docs/SETUP.md` dosyasına bakın.
