FROM golang:1.22-alpine as go-base
FROM aquasec/trivy:0.50.1 as trivy-base
FROM ghcr.io/oras-project/oras:v1.1.0 as oras-base

FROM alpine:3.19.1

RUN apk --no-cache add ca-certificates bash gcc libc-dev make

COPY --from=go-base /usr/local/go/ /usr/local/go/
COPY --from=trivy-base /usr/local/bin/trivy /usr/local/bin/trivy
COPY --from=oras-base /bin/oras /usr/local/bin/oras

ENV PATH="/usr/local/go/bin:${PATH}"

