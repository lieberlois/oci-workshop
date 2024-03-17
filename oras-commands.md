```bash
docker run -it -p 8080:5000 registry:2

oras push --plain-http localhost:8080/base64:v0.0.1 \
    --artifact-type application/vnd.oci.plugin.golang.so \
    base64.so:goplugin/so

trivy fs --format cyclonedx --output sbom.json .

oras attach \
    localhost:8080/base64:v0.0.1 \
    ./sbom.json \
    --artifact-type goplugin/sbom

DIGEST=$(oras discover localhost:8080/base64:v0.0.1 -o json | jq -r '.manifests[] | select(.artifactType == "goplugin/sbom") | .digest')
oras pull "localhost:8080/base64:v0.0.1@$DIGEST"

oras pull localhost:8080/base64:v0.0.1
```
