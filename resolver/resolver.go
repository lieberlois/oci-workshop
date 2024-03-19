package resolver

import (
	"io"
	"os"
	"plugin"
	"workshop/types"
)

type PluginResolver interface {
	Resolve(name string) (types.DecoderFunc, error)
}

func loadDecoderFuncFromPlugin(filepath string) (types.DecoderFunc, error) {
	plug, err := plugin.Open(filepath)
	if err != nil {
		return nil, err
	}

	symDecoderFunc, err := plug.Lookup("Decode")
	if err != nil {
		return nil, err
	}

	decoderFunc, ok := symDecoderFunc.(func(io.Reader) io.Reader)
	if !ok {
		return nil, os.ErrInvalid
	}

	return decoderFunc, nil
}
