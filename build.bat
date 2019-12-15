@echo off
cls
title Building ...

echo [*] Building windows ...
go build -ldflags="-s -w" -o build\jca.exe main\crypto.go main\shared.go
go build -ldflags="-s -w" -o build\jca-keygen.exe main\keygen.go main\shared.go

echo [*] Build linux ...
bash -c 'go build -ldflags="-s -w" -o build/jca main/crypto.go main/shared.go'
bash -c 'go build -ldflags="-s -w" -o build/jca-keygen main/keygen.go main/shared.go'

echo [*] Zipping ...
bash -c "zip -q build/jca-linux.zip build/jca build/jca-keygen"
bash -c "zip -q build/jca-win.zip build/jca.exe build/jca-keygen.exe"

echo [*] Build successful!

pause