#!/bin/bash

echo "==================================================="
echo "      byBooks Full-Stack App Starter"
echo "==================================================="
echo ""

echo "[1/4] Backend dependencies checking..."
cd backend
go mod download
echo "[2/4] Starting Go Backend in background..."
go run main.go &
BACKEND_PID=$!

cd ..

echo ""
echo "[3/4] Installing Frontend dependencies (npm install)..."
cd frontend
npm install
echo "[4/4] Starting React Frontend..."
npm run dev &
FRONTEND_PID=$!

echo ""
echo "==================================================="
echo "  🚀 Startup completed!"
echo ""
echo "  - Backend: http://localhost:8080"
echo "  - Swagger API: http://localhost:8080/swagger/index.html"
echo "  - Frontend: http://localhost:5173"
echo "==================================================="
echo ""
echo "Press [Ctrl+C] to stop both servers."

trap "kill $BACKEND_PID $FRONTEND_PID; exit" INT
wait
