```bash
docker run -it -p 8080:5000 registry:2

oras push --plain-http localhost:8080/base64:v0.0.1 \
    --artifact-type application/vnd.oci.plugin.golang.so \
    base64.so:goplugin/so

oras pull localhost:8080/base64:v0.0.1
```
