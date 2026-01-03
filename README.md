# Stok Takip Sistemi

Portable Multi-Database Stock Management System - Taşınabilir, çoklu veritabanı destekli stok takip uygulaması.

## 🎯 Özellikler

- ✅ **Taşınabilir**: Flash disk üzerinden çalışır, hiçbir bağımlılık gerektirmez
- ✅ **Çoklu Veritabanı**: Birden fazla veritabanı dosyası yönetimi
- ✅ **Otomatik Stok Hesaplama**: Giriş/çıkış işlemlerine göre otomatik stok takibi
- ✅ **Gelişmiş Filtreleme**: Kategorik, durum ve aralık bazlı filtreleme
- ✅ **Modern UI**: Vue 3 ve Tailwind CSS ile modern arayüz
- ✅ **Light Mode**: Sade ve profesyonel görünüm

## 🛠️ Teknoloji Yığını

- **Backend**: Go 1.21+ (GORM, SQLite3)
- **Frontend**: Vue 3 (Composition API, Pinia, Vue Router)
- **Bridge**: Wails v2
- **Database**: SQLite3
- **Styling**: Tailwind CSS

## 📋 Gereksinimler

### Geliştirme İçin

- Go 1.21 veya üzeri - [İndir](https://go.dev/dl/)
- Node.js 18 LTS veya üzeri - [İndir](https://nodejs.org/)
- Wails CLI - Kurulum: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- Git

### Çalıştırma İçin (Build edilmiş uygulama)

- Hiçbir bağımlılık gerekmez! Sadece `.exe` dosyasını çalıştırın.

## 🚀 Kurulum ve Çalıştırma

### 1. Projeyi Klonlayın

```bash
git clone <repository-url>
cd StokTakipDesktop
```

### 2. Go Bağımlılıklarını Yükleyin

```bash
go mod download
```

### 3. Frontend Bağımlılıklarını Yükleyin

```bash
cd frontend
npm install
cd ..
```

### 4. Geliştirme Modunda Çalıştırın

```bash
wails dev
```

Uygulama otomatik olarak açılacak ve hot-reload özelliği aktif olacaktır.

### 5. Production Build

```bash
wails build
```

Build edilen dosya `build/bin/` klasöründe olacaktır.

## 📁 Proje Yapısı

```
StokTakipDesktop/
├── internal/                    # Backend Go kodu
│   ├── app/                    # Ana uygulama
│   ├── models/                 # GORM modelleri
│   ├── services/               # İş mantığı servisleri
│   ├── database/               # Veritabanı yönetimi
│   ├── config/                 # Konfigürasyon
│   └── utils/                  # Yardımcı fonksiyonlar
├── frontend/                    # Vue.js frontend
│   ├── src/
│   │   ├── views/              # Ana sayfalar
│   │   ├── stores/             # Pinia state management
│   │   ├── router/             # Vue Router
│   │   └── style.css           # Global stiller
│   ├── index.html
│   └── package.json
├── plans/                       # Proje planları ve dokümantasyon
├── main.go                      # Ana entry point
├── wails.json                   # Wails konfigürasyonu
└── go.mod                       # Go modül tanımı
```

## 💻 Geliştirme

### Wails Komutları

```bash
# Geliştirme modu (hot reload)
wails dev

# Debug modu
wails dev -debug

# Production build
wails build

# Wails doctor (sistem kontrolü)
wails doctor

# Wails binding'leri güncelle
wails generate module
```

### Frontend Geliştirme

Frontend klasöründe ayrı olarak da çalışabilirsiniz:

```bash
cd frontend
npm run dev      # Geliştirme sunucusu
npm run build    # Production build
```

## 🗄️ Veritabanı Şeması

### Categories (Kategoriler)
- `id` - Primary Key
- `name` - Kategori adı
- `color` - HEX renk kodu
- `created_at`, `updated_at`

### Products (Ürünler)
- `id` - Primary Key
- `code` - Unique ürün kodu
- `name` - Ürün adı
- `category_id` - Foreign Key
- `unit` - Birim (adet, kg, litre vb.)
- `critical_limit` - Kritik stok seviyesi
- `price` - Birim fiyat
- `created_at`, `updated_at`

### Stock Movements (Stok Hareketleri)
- `id` - Primary Key
- `product_id` - Foreign Key
- `type` - IN (Giriş) / OUT (Çıkış)
- `quantity` - Miktar
- `date` - Hareket tarihi
- `note` - Açıklama
- `created_at`

## 🎨 UI/UX

- **Renk Teması**: Light mode (varsayılan)
- **Ana Renkler**: Blue (Primary), Gray (Secondary)
- **Font**: System fonts (platform native görünüm)
- **Responsive**: Minimum 1280x768 çözünürlük öneriliği

## 📝 Yapılacaklar (Roadmap)

### Sprint 1 ✅
- [x] Proje yapısı ve temel altyapı
- [x] Database connection manager
- [x] Path manager (taşınabilirlik)
- [x] Config manager
- [x] GORM modelleri
- [x] Temel servisler
- [x] Frontend yapısı (Vue 3 + Router + Pinia)
- [x] Database selector UI

### Sprint 2 (Sonraki Adım)
- [ ] Database service tam implementasyonu
- [ ] Wails binding'leri oluşturma
- [ ] Database seçici fonksiyonel hale getirme
- [ ] Kategori CRUD işlemleri
- [ ] Kategori yönetimi UI

### Sprint 3-8
Detaylı plan için `plans/stok-takip-mimari-plan.md` dosyasına bakınız.

## 🐛 Hata Ayıklama

### Sık Karşılaşılan Sorunlar

**Problem**: Wails komutu bulunamıyor
```bash
# Çözüm: PATH'e Wails ekleyin
export PATH=$PATH:$(go env GOPATH)/bin  # Linux/Mac
# veya Windows için System Environment Variables'dan GOPATH/bin ekleyin
```

**Problem**: Frontend bağımlılıkları yüklenemiyor
```bash
# Çözüm: Node modüllerini temizleyip tekrar yükleyin
cd frontend
rm -rf node_modules package-lock.json
npm install
```

**Problem**: Build hatası alıyorum
```bash
# Çözüm: Önce temizlik yapın
wails build -clean
```

## 📄 Lisans

Bu proje özel kullanım içindir.

## 👥 Katkıda Bulunma

Proje aktif geliştirme aşamasındadır. Önerileriniz için issue açabilirsiniz.

## 📞 İletişim

Sorularınız için issue açabilirsiniz.

---

**Not**: Bu proje Wails v2 kullanılarak geliştirilmiştir. Daha fazla bilgi için [Wails Documentation](https://wails.io) sayfasını ziyaret edin.
