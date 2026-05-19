@echo off
title byBooks - Full Stack Starter
color 0A

echo ===================================================
echo            byBooks Full-Stack App Starter
echo ===================================================
echo.

echo [1/4] Backend dependencies checking...
cd backend
call go mod download
echo [2/4] Starting Go Backend in a new window...
start "byBooks - Go Backend" cmd /k "go run main.go"

cd ..

echo.
echo [3/4] Installing Frontend dependencies (npm install)...
cd frontend
call npm install
echo [4/4] Starting React Frontend in a new window...
start "byBooks - React Frontend" cmd /k "npm run dev"

echo.
echo ===================================================
echo   🚀 Startup completed!
echo.
echo   - Backend: http://localhost:8080
echo   - Swagger API: http://localhost:8080/swagger/index.html
echo   - Frontend: http://localhost:5173
echo ===================================================
pause
