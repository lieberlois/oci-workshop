#!/bin/bash

if [[ -z "$OCI_REGISTRY" ]]; then
    echo Environment variable OCI_REGISTRY not set...
    exit 1
fi

http_option=""
if [[ "$OCI_REGISTRY" != *azurecr.io ]]; then
    http_option="--plain-http"
fi

for dir in out/*; do
    plugin_name="$(echo $dir | sed 's/out\///g')"
    (cd $dir && oras attach \
        ${http_option} \
        "${OCI_REGISTRY}/${plugin_name}:v0.0.1" \
        ./sbom.json \
        --artifact-type goplugin/sbom)
done