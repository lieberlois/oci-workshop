package main

import (
	"io"
	"os"
	"workshop/resolver"
)

func main() {
	// pluginResolver := &resolver.FSPluginResolver{}
	// plugins := []string{"./out/base64.so"}

	pluginResolver, cleanupFunc := resolver.NewOCIPluginResolver(
		resolver.WithHostname("localhost"),
		resolver.WithPort("8080"),
		resolver.WithPluginDir("./plugins"),
	)
	defer cleanupFunc()

	plugins := []string{"base64:v0.0.1"}

	// Reader variable for plugin chain
	var reader io.Reader

	// Initialize with file
	reader, err := os.Open("super-secret.txt")
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
