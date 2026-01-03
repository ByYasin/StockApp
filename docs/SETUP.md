# Stok Takip Sistemi - Kurulum ve Başlangıç Rehberi

## 📦 İlk Kurulum

### 1. Sistem Gereksinimlerini Kontrol Edin

```bash
# Go versiyonunu kontrol edin (1.21+ olmalı)
go version

# Node.js versiyonunu kontrol edin (18+ olmalı)
node --version

# Wails kurulu mu kontrol edin
wails doctor
```

### 2. Wails CLI Kurulumu (Eğer kurulu değilse)

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 3. Proje Bağımlılıklarını Yükleyin

```bash
# Go bağımlılıkları
go mod download

# Frontend bağımlılıkları
cd frontend
npm install
cd ..
```

## 🚀 Geliştirme Ortamını Başlatma

### İlk Çalıştırma

```bash
# Tüm binding'leri oluştur
wails generate module

# Geliştirme modunda çalıştır
wails dev
```

### Sorun Giderme

Eğer `wails generate module` hatası alırsanız:

1. `frontend/wailsjs` klasörünü silin (eğer varsa)
2. Önce build deneyin: `wails build`
3. Sonra tekrar dev modunda çalıştırın: `wails dev`

## 📝 Geliştirme Süreci

### Sonraki Adımlar (Sprint 2)

1. **Backend'i Tamamlayın**
   - Product Service implementasyonu
   - Movement Service implementasyonu
   - Stok hesaplama fonksiyonları

2. **Wails Binding'leri Oluşturun**
   ```bash
   # Her backend değişikliğinden sonra
   wails generate module
   ```

3. **Frontend'i Güncelleyin**
   - Wails binding'lerini import edin
   - API çağrılarını aktif hale getirin
   - Store'ları tamamlayın

### Örnek: Database Service'i Kullanma

**Backend** (`internal/services/database_service.go`):
```go
func (s *DatabaseService) ListDatabases() ([]DatabaseInfo, error) {
    // Implementation
}
```

**Frontend** (`frontend/src/stores/database.js`):
```javascript
import { ListDatabases } from '@wails/go/services/DatabaseService'

async function loadDatabases() {
    const result = await ListDatabases()
    databases.value = result
}
```

## 🧪 Test Etme

### Manuel Test

1. Database oluşturma
2. Database seçme
3. Dashboard'a geçiş
4. Navigasyon kontrolü

### Build Test

```bash
# Production build
wails build

# Build edilen dosyayı test edin
./build/bin/StockApp.exe
```

## 📊 Proje Durumu

### Tamamlanan (Sprint 1) ✅

- [x] Proje yapısı kuruldu
- [x] Backend temel altyapısı (models, database, config, utils)
- [x] Frontend temel yapısı (Vue 3, Router, Pinia)
- [x] Database selector UI
- [x] Dashboard UI (placeholder)
- [x] Dokümantasyon

### Yapılacaklar (Sprint 2) 🚧

- [ ] Wails binding'leri oluştur
- [ ] Database service'i frontend ile bağla
- [ ] Kategori CRUD backend
- [ ] Kategori CRUD frontend
- [ ] Kategori yönetimi UI

### Gelecek Sprintler (3-8) 📅

Detaylı plan için [`plans/stok-takip-mimari-plan.md`](plans/stok-takip-mimari-plan.md) dosyasına bakınız.

## 💡 İpuçları

### Hot Reload
`wails dev` modunda hem Go hem de Vue değişiklikleri otomatik yansır.

### Debug
```bash
wails dev -debug
```
Bu komut Chrome DevTools'u etkinleştirir.

### Build Optimizasyonu
```bash
# Küçük binary için
wails build -ldflags="-s -w"
```

## 🔧 Sık Kullanılan Komutlar

```bash
# Geliştirme
wails dev                    # Hot reload ile çalıştır
wails dev -debug             # Debug mode
wails generate module        # Binding'leri yeniden oluştur

# Build
wails build                  # Production build
wails build -clean           # Temiz build
wails build -platform windows/amd64  # Spesifik platform

# Utility
wails doctor                 # Sistem kontrolü
go mod tidy                  # Go bağımlılıklarını temizle
```

## 📁 Çalışma Dizini

Uygulama çalıştığı konumu otomatik algılar:

```
/YourFlashDisk/
├── StockApp.exe       # Ana uygulama
├── config.json        # Otomatik oluşturulur
└── Data/              # Otomatik oluşturulur
    ├── Depo_A.db
    └── Yedek_Parca.db
```

## 🎯 Sonraki Sprint için Hazırlık

1. Go backend kodlarını yazın
2. `wails generate module` ile binding'leri oluşturun
3. Frontend'de Wails import'larını ekleyin:
   ```javascript
   import { ListDatabases, CreateDatabase } from '@wails/go/services/DatabaseService'
   ```
4. Store fonksiyonlarını aktif hale getirin
5. Test edin!

## 📚 Kaynaklar

- [Wails Documentation](https://wails.io)
- [Vue 3 Documentation](https://vuejs.org)
- [Pinia Documentation](https://pinia.vuejs.org)
- [GORM Documentation](https://gorm.io)
- [Tailwind CSS](https://tailwindcss.com)

---

**Önemli**: Her backend değişikliğinden sonra `wails generate module` komutunu çalıştırmayı unutmayın!
