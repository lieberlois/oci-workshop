out/base64/decoder.so: decoders/base64/decoder.go
	@echo "Building Base64 Decoder..."
	@go build -buildmode=plugin -o out/base64/decoder.so decoders/base64/decoder.go

out/base64/sbom.json: decoders/base64/decoder.go
	@echo "Generating SBOM..."
	@trivy fs --format cyclonedx --output out/base64/sbom.json ./decoders/base64

.PHONY: build
build: out/base64/decoder.so out/base64/sbom.json  # Maybe change to out/base64/decoder.so and add task for SBOM

.PHONY: decode
decode: build
	@go run main.go		