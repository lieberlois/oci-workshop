package main

import (
	"io"
	"os"
	"workshop/resolver"
)

func main() {
	var reader io.Reader
	// pluginResolver := &resolver.FSPluginResolver{}
	// plugins := []string{"./out/base64.so"}

	ociResolverConfig := resolver.OCIResolverConfig{
		Hostname:  "localhost",
		Port:      "8080",
		PluginDir: "./plugins",
	}
	pluginResolver, cleanupFunc := resolver.NewOCIPluginResolver(ociResolverConfig)
	defer cleanupFunc()

	plugins := []string{"base64:v0.0.1"}

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
