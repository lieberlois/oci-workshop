out/base64.so: decoders/base64/decoder.go
	@echo "Building Base64 Decoder..."
	@go build -buildmode=plugin -o out/base64.so decoders/base64/decoder.go

.PHONY: build
build: out/base64.so