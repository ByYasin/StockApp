# Stok Takip Masaüstü Uygulaması (Windows, MacOS ve Linux için derlenebilir.)

Modern ve kullanıcı dostu bir masaüstü stok takip uygulaması. Küçük ve orta ölçekli işletmeler için tasarlanmış, internet bağlantısı gerektirmeyen yerli bir çözüm.

## Nedir?

Bu uygulama, işletmelerin ürün stoklarını yönetmelerine yardımcı olan bir masaüstü yazılımıdır. Wails framework'ü kullanılarak geliştirilmiştir ve Windows, macOS ve Linux işletim sistemlerinde çalışır.

## Temel Özellikler

**Ürün Yönetimi**
- Ürün ekleme, düzenleme ve silme
- Kategori bazlı organizasyon
- Kritik stok seviyesi takibi
- Detaylı ürün filtreleme

**Stok Hareketi**
- Stok giriş ve çıkış kayıtları
- Otomatik stok hesaplama
- Hareket geçmişi takibi
- Not ekleme imkanı

**Kategori Sistemi**
- Renk kodlamalı kategori yönetimi
- Hiyerarşik organizasyon
- Esnek kategori yapısı

**Dashboard**
- Anlık stok durumu görüntüleme
- Kritik stok uyarıları
- Günlük hareket özeti
- Hızlı erişim kısayolları

**Teknik Özellikler**
- Çoklu veritabanı desteği
- Portable kullanım (flash disk'ten çalıştırma)
- Dark/Light tema desteği
- Responsive arayüz
- Hızlı ve hafif

## Teknoloji Stack'i

**Backend**
- Go (Golang)
- GORM ORM
- SQLite veritabanı

**Frontend**
- Vue.js 3
- Tailwind CSS
- Pinia state management

**Desktop Framework**
- Wails v2

## Kurulum

Gereksinimler:
- Go 1.19 veya üzeri
- Node.js 16 veya üzeri
- Wails CLI

```bash
# Wails CLI kurulumu
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Proje bağımlılıklarını yükleme
wails doctor

# Geliştirme modunda çalıştırma
wails dev

# Production build
wails build
```

## Kullanım

1. Uygulamayı başlatın
2. İlk açılışta yeni bir veritabanı oluşturun veya mevcut olanı seçin
3. Dashboard'dan kategoriler oluşturun
4. Ürünlerinizi sisteme ekleyin
5. Stok giriş/çıkış işlemlerini kaydedin

## Proje Yapısı

```
StokTakipDesktop/
├── internal/          # Go backend kodu
│   ├── models/       # Veritabanı modelleri
│   ├── services/     # İş mantığı katmanı
│   ├── database/     # Veritabanı yönetimi
│   ├── config/       # Konfigürasyon
│   └── utils/        # Yardımcı fonksiyonlar
├── frontend/          # Vue.js frontend
│   ├── src/
│   │   ├── views/    # Sayfa bileşenleri
│   │   ├── stores/   # State management
│   │   └── router/   # Routing
│   └── wailsjs/      # Wails binding'leri
└── build/            # Build çıktıları
```

## Lisans

MIT