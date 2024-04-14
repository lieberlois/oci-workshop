#!/bin/bash

for dir in out/*; do
    plugin_name="$(echo $dir | sed 's/out\///g')"
    (cd $dir && oras push \
        --plain-http "162.55.221.56:5000/${plugin_name}:v0.0.1" \
        --artifact-type application/vnd.oci.plugin.golang.so \
        decoder.so:goplugin/so)
done