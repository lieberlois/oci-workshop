```bash
docker run -it -p 8080:5000 registry:2

# Base64
oras push --plain-http localhost:8080/base64:v0.0.1 --artifact-type application/vnd.oci.plugin.golang.so out/base64/decoder.so:goplugin/so
oras attach --plain-http localhost:8080/base64:v0.0.1 ./sbom.json --artifact-type goplugin/sbom

# Json
oras push --plain-http localhost:8080/json:v0.0.1 --artifact-type application/vnd.oci.plugin.golang.so out/json/decoder.so:goplugin/so
oras attach --plain-http localhost:8080/json:v0.0.1 ./sbom.json --artifact-type goplugin/sbom

# Reverse
oras push --plain-http localhost:8080/reverse:v0.0.1 --artifact-type application/vnd.oci.plugin.golang.so out/reverse/decoder.so:goplugin/so
oras attach --plain-http localhost:8080/reverse:v0.0.1 ./sbom.json --artifact-type goplugin/sbom

# Hex
oras push --plain-http localhost:8080/hex:v0.0.1 --artifact-type application/vnd.oci.plugin.golang.so out/hex/decoder.so:goplugin/so
oras attach --plain-http localhost:8080/hex:v0.0.1 ./sbom.json --artifact-type goplugin/sbom

oras discover --plain-http localhost:8080/base64:v0.0.1 -o tree

DIGEST=$(oras discover --plain-http localhost:8080/base64:v0.0.1 -o json | jq -r '.manifests[] | select(.artifactType == "goplugin/sbom") | .digest')
oras pull "localhost:8080/base64:v0.0.1@$DIGEST"

oras pull localhost:8080/base64:v0.0.1



# For Gitea:
oras login 128.140.45.106:3000 --username luis.schweigard@gmail.com --password password123 --plain-http
export OCI_REGISTRY=128.140.45.106:3000/lieberlois/ocitest
# ---> Also adjust the artifact names in the code e.g. from json:v0.0.1 to lieberlois/ocitest/json:v0.0.1

# For ACR:
export OCI_REGISTRY=<acrname.azurecr.io>
oras login $OCI_REGISTRY --username <acrname> --password <password>
docker login $OCI_REGISTRY

```
