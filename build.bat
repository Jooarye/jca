@echo off
cls
title Building ...

echo [*] Building windows ...
go build -ldflags="-s -w" -o build\jca.exe main\crypto.go main\shared.go
go build -ldflags="-s -w" -o build\jca-keygen.exe main\keygen.go main\shared.go
go build -ldflags="-s -w" -o build\jca-integrity.exe main\integrity.go main\shared.go

echo [*] Build successful!

pause
