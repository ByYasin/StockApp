# Test Talimatları

## Uygulamayı Çalıştırma

### 1. Development Mode (Geliştirme)
Terminal'de şu komutu çalıştırın:
```bash
wails dev
```

**Beklenen Davranış:**
- Wails CLI binding'leri oluşturacak
- Frontend bağımlılıkları yüklenecek
- Frontend derlenecek
- Uygulama penceresi açılacak

**Not:** İlk çalıştırmada 1-2 dakika sürebilir.

### 2. Production Build
Terminal'de şu komutu çalıştırın:
```bash
wails build
```

Çıktı: `build/bin/StokTakip.exe`

## Test Senaryoları

### ✅ Senaryo 1: İlk Kullanım
1. Uygulamayı çalıştırın
2. "Yeni Veritabanı Oluştur" butonuna tıklayın
3. Veritabanı adı girin (örn: "TestDB")
4. Ana ekran açılmalı

**Beklenen Sonuç:** 
- Dashboard açılmalı
- Tüm istatistikler 0 olmalı
- "Henüz hareket kaydı yok" mesajı görülmeli

### ✅ Senaryo 2: Kategori Ekleme
1. Üst menüden "Kategoriler" sekmesine tıklayın
2. "+ Yeni Kategori" butonuna tıklayın
3. Form doldur:
   - Ad: "Elektronik"
   - Açıklama: "Elektronik ürünler"
   - Renk: Mavi seç (#3B82F6)
4. "Kaydet" butonuna tıklayın

**Beklenen Sonuç:**
- Modal kapanmalı
- Yeni kategori kartı görünmeli
- "0 ürün" yazısı olmalı

**Tekrarlayın:**
- "Gıda" kategorisi ekleyin (Yeşil)
- "Kırtasiye" kategorisi ekleyin (Turuncu)

### ✅ Senaryo 3: Ürün Ekleme
1. Üst menüden "Ürünler" sekmesine tıklayın
2. "+ Yeni Ürün" butonuna tıklayın
3. Form doldur:
   - Ürün Kodu: "ELK001"
   - Ürün Adı: "Mouse"
   - Kategori: "Elektronik" seç
   - Birim: "Adet"
   - Kritik Stok Seviyesi: 5
   - Birim Fiyat: 150.00
4. "Kaydet" butonuna tıklayın

**Beklenen Sonuç:**
- Modal kapanmalı
- Tabloda yeni ürün görünmeli
- Stok: 0 (kırmızı renkte)
- Kategori: Mavi "Elektronik" badge

**Tekrarlayın:**
- "Klavye" ekleyin (ELK002, Elektronik)
- "Ekmek" ekleyin (GID001, Gıda)

### ✅ Senaryo 4: Stok Giriş İşlemi
1. Üst menüden "Hareketler" sekmesine tıklayın
2. "+ Yeni Hareket" butonuna tıklayın
3. Form doldur:
   - Ürün: "Mouse (Mevcut: 0)" seç
   - Hareket Tipi: "Giriş" seç (yeşil)
   - Miktar: 20
   - Not: "İlk alım"
4. "Kaydet" butonuna tıklayın

**Beklenen Sonuç:**
- Modal kapanmalı
- Tabloda yeni hareket görünmeli
- İstatistikler güncellenmiş olmalı (Toplam Giriş: 20)

### ✅ Senaryo 5: Stok Çıkış İşlemi
1. "Hareketler" sekmesinde "+ Yeni Hareket"
2. Form doldur:
   - Ürün: "Mouse (Mevcut: 20)" seç
   - Hareket Tipi: "Çıkış" seç (kırmızı)
   - Miktar: 15
   - Not: "Satış"
3. "Kaydet"

**Beklenen Sonuç:**
- Çıkış hareketi eklenmeli
- İstatistikler: Toplam Giriş: 20, Toplam Çıkış: 15

### ✅ Senaryo 6: Dashboard Kontrolü
1. "Ana Sayfa" sekmesine dön
2. İstatistikleri kontrol et:
   - Toplam Ürün: 3
   - Kritik Stok: 0 (çünkü Mouse 5 olmalı, şu an 5)
   - Kategoriler: 3
   - Bugünkü Hareketler: 2

**Beklenen Sonuç:**
- Tüm istatistikler doğru
- "Son Hareketler" tablosunda 2 kayıt

### ✅ Senaryo 7: Kritik Stok Uyarısı
1. "Hareketler" sekmesine git
2. Mouse için çıkış yap: 1 adet
   - Mouse kritik seviyede olmalı (5)
3. "Ana Sayfa"ya dön

**Beklenen Sonuç:**
- Kritik Stok: 1 (Mouse)
- Dashboard'da Mouse turuncu renkte görünmeli

### ✅ Senaryo 8: Filtreleme Testi
1. "Ürünler" sekmesine git
2. Arama kutusuna "Mouse" yaz
   - Sadece Mouse görünmeli
3. Arama kutusunu temizle
4. Kategori filtresinden "Elektronik" seç
   - Sadece Mouse ve Klavye görünmeli
5. Stok Filtresi: "Kritik Stok" seç
   - Sadece Mouse görünmeli

### ✅ Senaryo 9: Düzenleme İşlemi
1. "Kategoriler" sekmesinde "Elektronik"in düzenle butonuna tıkla
2. Adı "Elektronik Ürünler" olarak değiştir
3. Kaydet

**Beklenen Sonuç:**
- Kategori adı güncellenmiş olmalı
- Ürünler sekmesinde de yeni ad görünmeli

### ✅ Senaryo 10: Silme İşlemi
1. "Hareketler" sekmesinde bir hareketi sil
2. Onay mesajı gelecek
3. "Evet" de

**Beklenen Sonuç:**
- Hareket silinmeli
- İstatistikler güncellenmiş olmalı
- Dashboard'da sayılar değişmiş olmalı

### ✅ Senaryo 11: Çoklu Veritabanı
1. Dashboard'da veritabanı adının yanındaki alana dikkat et
2. Uygulamayı kapat
3. Tekrar aç
4. Son kullanılan veritabanı otomatik yüklenmeli

**Alternatif:** 
- Yeni bir veritabanı oluştur
- Veritabanları arasında geçiş yap

## Hata Kontrolleri

### ❌ Test 1: Stok Yetersizliği
1. Çıkış yapmaya çalış (örn: Mouse için 100 adet)
2. "Yetersiz stok!" uyarısı almalısınız

### ❌ Test 2: Boş Alan Kontrolü
1. Yeni ürün eklerken boş alanları bırakın
2. "Bu alan gereklidir" uyarısı almalısınız

### ❌ Test 3: Kategori Silme Engeli
1. Ürünleri olan bir kategoriyi silmeye çalışın
2. "Bu kategori X ürün tarafından kullanılıyor" uyarısı almalısınız

## Performans Testleri

### 📊 Test 1: Çok Sayıda Kayıt
1. 100+ ürün ekleyin (script veya manuel)
2. Filtreleme hızını test edin
3. Arama performansını kontrol edin

**Beklenen:** Anında sonuç

### 📊 Test 2: Çok Sayıda Hareket
1. 1000+ hareket oluşturun
2. Dashboard yüklenme hızı
3. Hareketler sayfası yüklenme hızı

**Beklenen:** 1-2 saniye içinde

## Portable Test

### 💾 Test: Flash Disk Taşınabilirlik
1. Uygulamayı build edin: `wails build`
2. `build/bin/StokTakip.exe` dosyasını flash diske kopyalayın
3. Flash diskten çalıştırın
4. Yeni veritabanı oluşturun
5. Veriler ekleyin
6. Uygulamayı kapatın
7. Flash diski başka bir bilgisayara takın
8. Tekrar açın

**Beklenen Sonuç:**
- Veriler korunmalı
- Son veritabanı otomatik açılmalı

## Sorun Giderme

### Uygulama Açılmıyor
```bash
# Temiz build deneyin
wails clean
wails dev
```

### Binding Hataları
```bash
# Wails'i güncelleyin
wails doctor
go mod tidy
```

### Frontend Hataları
```bash
cd frontend
npm install
npm run build
cd ..
wails dev
```

## Başarı Kriterleri

✅ Tüm sayfalar "Yakında gelecek" yerine gerçek içerik gösteriyor
✅ CRUD işlemleri çalışıyor
✅ Filtreleme ve arama çalışıyor
✅ İstatistikler doğru hesaplanıyor
✅ Stok otomatik güncelleniyor
✅ Kritik stok uyarıları çalışıyor
✅ Multi-database desteği çalışıyor
✅ Portable olarak çalışıyor

## Rapor

Test sonuçlarını kaydedin:
- ✅ Başarılı testler
- ❌ Başarısız testler
- 🐛 Bulunan hatalar
- 💡 İyileştirme önerileri
