Portable Multi-Database Stock Management System - PRD

🎯 Proje Vizyonu

Flash disk üzerinden çalışan, hiçbir bağımlılığı olmayan (Zero-dependency), her bilgisayarda açıldığı anda kendi dizinindeki veritabanlarını yönetebilen, gelişmiş filtreleme özelliklerine sahip profesyonel bir stok takip uygulaması.

🛠 Teknoloji Yığını (Tech Stack)

Backend: Golang (Hız, tek binary çıktı ve düşük kaynak tüketimi için).

Frontend: Vue.js (Wails aracılığıyla native görünümlü UI).

Bridge: Wails (latest) (Go ve Frontend'i birleştiren modern framework).

Database: SQLite3 (Dosya tabanlı, taşınabilir ve sunucu gerektirmeyen yapı).

ORM: GORM (Hızlı veritabanı işlemleri ve otomatik migrasyon için).

📁 Dosya ve Taşınabilirlik Mimarisi

Uygulama, çalıştığı konumu dinamik olarak algılamalıdır.

Plaintext

/Root/
├── StockApp.exe           # Ana uygulama
├── config.json            # Tercihler ve son açılan DB bilgisi
└── /Data/                 # Veritabanı dosyaları (.db)
    ├── Depo_A.db
    └── Yedek_Parca.db

🚀 Temel Fonksiyonel Gereksinimler

1. Veritabanı Seçici (Launcher)

Uygulama açılışında /Data klasörünü tarar.

Mevcut .db dosyalarını listeler.

Yeni veritabanı oluşturma imkanı tanır (Tabloları otomatik migrate eder).

Seçilen veritabanına dinamik olarak bağlanır.

2. Stok Yönetimi & Otomatik Hesaplama

Ürün Kartı: Ürün Kodu (Manuel), Ürün Adı, Kategori, Birim, Kritik Stok Seviyesi, Birim Fiyat.

Hareket Kaydı: Giriş (+) veya Çıkış (-) işlemleri.

Otomatik Bakiye: Mevcut stok = (Tüm Girişler) - (Tüm Çıkışlar).

Negatif Stok Kontrolü: Stok miktarını aşan çıkışlarda uyarı mekanizması.

3. Gelişmiş Filtreleme Sistemi (Detaylı)

Global Arama: İsim ve kod içerisinde "Fuzzy Search".

Kategorik Filtreleme: Çoklu kategori seçimi.

Durum Filtreleri: "Kritik Stoktakiler", "Stokta Olmayanlar", "Son 7 Gün Hareket Görenler".

Aralık Filtreleri: Tarih aralığı, Miktar aralığı (Örn: 10-50 arası), Fiyat aralığı.

Sıralama: En çok azalanlar, en son eklenenler, alfabetik.

🏗 Veritabanı Şeması (Database Schema)

products table

id (Primary Key)

code (Unique Index)

name (Index)

category_id (Foreign Key)

unit (string)

critical_limit (int)

price (float)

stock_movements table
id (Primary Key)

product_id (Foreign Key)

type (enum: IN, OUT)

quantity (int)

date (datetime, Index)

note (text)

🛠 Geliştirme Talimatları (Agent İçin)

Dinamik Yol Yönetimi: Tüm dosya okuma/yazma işlemleri os.Executable() baz alınarak "relative path" ile yapılmalı.

Veritabanı Bağlantısı: Singleton pattern kullanılmalı ancak veritabanı değiştirildiğinde bağlantı güvenli bir şekilde kapatılıp yeni dosyaya yönlenmeli.

UI/UX: Karanlık mod desteği istemiyorum bu projede sadece light olsun varsayılan olarak, veri tablolarında "lazy loading" veya "pagination" kullanılmalı.

Hata Yönetimi: Flash disk sökülmesine karşı veritabanı işlemlerinde Panic Recovery mekanizması kurulmalı.