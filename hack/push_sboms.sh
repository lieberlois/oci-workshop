#!/bin/bash

if [[ -z "$OCI_REGISTRY" ]]; then
    echo Environment variable OCI_REGISTRY not set...
    exit 1
fi

for dir in out/*; do
    plugin_name="$(echo $dir | sed 's/out\///g')"
    (cd $dir && oras attach \
        --plain-http \
        "${OCI_REGISTRY}/${plugin_name}:v0.0.1" \
        ./sbom.json \
        --artifact-type goplugin/sbom)
done