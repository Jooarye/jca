cd /mnt/c/Users/Jooarye/go/JCA/

go build -ldflags="-s -w" -o build/jca main/crypto.go main/shared.go
go build -ldflags="-s -w" -o build/jca-keygen main/keygen.go main/shared.go