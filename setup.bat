@echo off
echo ================================================
echo Stok Takip Sistemi - Kurulum Script
echo ================================================
echo.

echo [1/5] Go bagimliliklar yukleniyor...
go mod download
if %errorlevel% neq 0 (
    echo HATA: Go bagimliliklari yuklenemedi!
    pause
    exit /b 1
)
echo ✓ Go bagimliliklari yuklendi
echo.

echo [2/5] Frontend klasorune geciliyor...
cd frontend
if %errorlevel% neq 0 (
    echo HATA: Frontend klasoru bulunamadi!
    cd ..
    pause
    exit /b 1
)
echo.

echo [3/5] NPM bagimliliklar yukleniyor...
call npm install
if %errorlevel% neq 0 (
    echo HATA: NPM bagimliliklari yuklenemedi!
    cd ..
    pause
    exit /b 1
)
echo ✓ NPM bagimliliklari yuklendi
echo.

echo [4/5] Ana dizine donuluyor...
cd ..
echo.

echo [5/5] Frontend dist klasoru olusturuluyor...
if not exist "frontend\dist" (
    mkdir frontend\dist
    echo. > frontend\dist\.gitkeep
)
echo ✓ Dist klasoru hazir
echo.

echo ================================================
echo KURULUM TAMAMLANDI!
echo ================================================
echo.
echo Simdi asagidaki komutu calistirabilirsiniz:
echo   wails dev
echo.
echo veya production build icin:
echo   wails build
echo.
pause
