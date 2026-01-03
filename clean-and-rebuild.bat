@echo off
echo ====================================
echo Stok Takip - Temizleme ve Yeniden Derleme
echo ====================================
echo.

echo 1. Wails cache temizleniyor...
wails clean

echo.
echo 2. Frontend node_modules temizleniyor...
cd frontend
if exist node_modules rmdir /s /q node_modules
if exist dist rmdir /s /q dist
if exist .vite rmdir /s /q .vite

echo.
echo 3. Frontend bagimliliklari yeniden yukleniyor...
call npm install

echo.
echo 4. Ana dizine donuluyor...
cd ..

echo.
echo 5. Go modulleri guncelleniyor...
go mod tidy

echo.
echo 6. Wailsjs binding'leri yeniden olusturuluyor...
echo NOT: wails dev komutunu manuel olarak calistirin!
echo.

echo ====================================
echo Temizleme tamamlandi!
echo Simdi 'wails dev' komutunu calistirin
echo ====================================
pause
