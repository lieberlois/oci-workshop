#!/bin/bash

for dir in decoders/*; do
    plugin_name="$(echo $dir | sed 's/decoders\///g')"
    echo "Building plugin ${plugin_name}..."
    go build -buildmode=plugin -o "out/${plugin_name}/decoder.so" "${dir}/decoder.go"
done