out/base64/decoder.so: decoders/base64/decoder.go
	@echo "Building Base64 Decoder..."
	@go build -buildmode=plugin -o out/base64/decoder.so decoders/base64/decoder.go

out/base64/sbom.json: decoders/base64/decoder.go
	@echo "Generating SBOM..."
	@trivy fs --format cyclonedx --output out/base64/sbom.json ./decoders/base64  # TODO: probably move to bash script, currently does not detect dependencies (requires go mod init & go mod tidy)

out/hex/decoder.so: decoders/hex/decoder.go
	@echo "Building Hex Decoder..."
	@go build -buildmode=plugin -o out/hex/decoder.so decoders/hex/decoder.go

out/hex/sbom.json: decoders/hex/decoder.go
	@echo "Generating SBOM..."
	@trivy fs --format cyclonedx --output out/hex/sbom.json ./decoders/hex  # TODO: probably move to bash script, currently does not detect dependencies (requires go mod init & go mod tidy)

out/json/decoder.so: decoders/json/decoder.go
	@echo "Building JSON Decoder..."
	@go build -buildmode=plugin -o out/json/decoder.so decoders/json/decoder.go

out/json/sbom.json: decoders/json/decoder.go
	@echo "Generating SBOM..."
	@trivy fs --format cyclonedx --output out/json/sbom.json ./decoders/json  # TODO: probably move to bash script, currently does not detect dependencies (requires go mod init & go mod tidy)

out/reverse/decoder.so: decoders/reverse/decoder.go
	@echo "Building Reverse Decoder..."
	@go build -buildmode=plugin -o out/reverse/decoder.so decoders/reverse/decoder.go

out/reverse/sbom.json: decoders/reverse/decoder.go
	@echo "Generating SBOM..."
	@trivy fs --format cyclonedx --output out/reverse/sbom.json ./decoders/reverse  # TODO: probably move to bash script, currently does not detect dependencies (requires go mod init & go mod tidy)

.PHONY: build
build: out/base64/decoder.so out/base64/sbom.json out/hex/decoder.so out/hex/sbom.json out/json/decoder.so out/json/sbom.json out/reverse/decoder.so out/reverse/sbom.json

.PHONY: decode
decode: build
	@go run main.go		