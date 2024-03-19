```bash
docker run -it -p 8080:5000 registry:2

# Base64
oras push --plain-http localhost:8080/base64:v0.0.1 --artifact-type application/vnd.oci.plugin.golang.so decoder.so:goplugin/so
oras attach localhost:8080/base64:v0.0.1 ./sbom.json --artifact-type goplugin/sbom

# Json
oras push --plain-http localhost:8080/json:v0.0.1 --artifact-type application/vnd.oci.plugin.golang.so decoder.so:goplugin/so
oras attach localhost:8080/json:v0.0.1 ./sbom.json --artifact-type goplugin/sbom

# Reverse
oras push --plain-http localhost:8080/reverse:v0.0.1 --artifact-type application/vnd.oci.plugin.golang.so decoder.so:goplugin/so
oras attach localhost:8080/reverse:v0.0.1 ./sbom.json --artifact-type goplugin/sbom

oras discover localhost:8080/base64:v0.0.1 -o tree

DIGEST=$(oras discover localhost:8080/base64:v0.0.1 -o json | jq -r '.manifests[] | select(.artifactType == "goplugin/sbom") | .digest')
oras pull "localhost:8080/base64:v0.0.1@$DIGEST"

oras pull localhost:8080/base64:v0.0.1
```
