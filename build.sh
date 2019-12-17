go build -ldflags="-s -w" -o build/jca main/crypto.go main/shared.go
go build -ldflags="-s -w" -o build/jca-keygen main/keygen.go main/shared.go
go build -ldflags="-s -w" -o build/jca-integrity main/integrity.go main/shared.go
zip -q build/jca-linux.zip build/jca build/jca-keygen build/jca-integrity

# for file in $(find build/ -executable -type f); do sudo cp $file /opt/jca/; done
