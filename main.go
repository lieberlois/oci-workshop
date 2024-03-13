package main

import (
	"fmt"
	"io"
	"os"
	"plugin"
)

func main() {
	// TODO: Load Plugin File from OCI, either via Oras Lib or CLI
	decoderFunc, err := loadDecoderFuncFromPlugin("out/base64.so")

	// PluginLoader Interface -> either OCI or ... ? Maybe local FS?

	if err != nil {
		panic(err)
	}

	// str := "SGVsbG8gd29ybGQK"
	// reader := strings.NewReader(str)
	reader, err := os.Open("super-secret.txt")
	if err != nil {
		panic(err)
	}

	result := decoderFunc(reader)

	_, err = io.Copy(os.Stdout, result)
	if err != nil {
		panic(err)
	}
}

func loadDecoderFuncFromPlugin(path string) (func(io.Reader) io.Reader, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}

	symDecoderFunc, err := plug.Lookup("Decode")
	if err != nil {
		panic(err)
	}

	decoderFunc, ok := symDecoderFunc.(func(io.Reader) io.Reader)
	if !ok {
		fmt.Println("Unexpected type from module symbol")
		os.Exit(1)
	}
	return decoderFunc, err
}
