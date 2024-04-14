#!/bin/bash

rm -f "decoders/**/go.mod" > /dev/null
rm -f "decoders/**/go.sum" > /dev/null

for dir in decoders/*; do
    plugin_name="$(echo $dir | sed 's/decoders\///g')"

    if ! [[ -d "out/${plugin_name}" ]]; then
        echo "Output for Plugin ${plugin_name} not found..."
        exit 1
    fi

    (cd $dir && go mod init decoder && go mod tidy)
    
    trivy fs \
        --format cyclonedx \
        --output "out/${plugin_name}/sbom.json" ${dir}

    rm "${dir}/go.mod"
done