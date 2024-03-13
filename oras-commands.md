```bash
docker run -it -p 8080:5000 registry:2

oras push --plain-http localhost:8080/base64decoder:v0.0.1 \
    --artifact-type application/vnd.acme.rocket.config \
    base64.so:goplugin/so

oras pull localhost:8080/base64decoder:v0.0.1
```
