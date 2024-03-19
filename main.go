package main

import (
	"io"
	"os"
	"workshop/resolver"
)

func main() {
	sbomValidation := true

	pluginResolver := resolver.NewFSPluginResolver(sbomValidation)
	plugins := []string{"./out/json", "./out/reverse", "./out/base64"}

	// pluginResolver, cleanupFunc := resolver.NewOCIPluginResolver(
	// 	resolver.WithHostname("localhost"),
	// 	resolver.WithPort("8080"),
	// 	resolver.WithPluginDir("./plugins"),
	// 	resolver.WithValidateSbom(sbomValidation),
	// )
	// defer cleanupFunc()
	// plugins := []string{"json:v0.0.1", "reverse:v0.0.1", "base64:v0.0.1"}

	// Reader variable for plugin chain
	var reader io.Reader

	// Initialize with file
	reader, err := os.Open("super-secret.json")
	if err != nil {
		panic(err)
	}

	for _, plugin := range plugins {
		decoderFunc, err := pluginResolver.Resolve(plugin)
		if err != nil {
			panic(err)
		}

		reader = decoderFunc(reader)
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		panic(err)
	}
}
