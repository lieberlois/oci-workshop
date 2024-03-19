package resolver

import (
	"fmt"
	"io"
	"os"
	"plugin"
	"workshop/types"
	"workshop/util"

	cdx "github.com/CycloneDX/cyclonedx-go"
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

func validateSbom(reader io.Reader) error {
	bom := cdx.BOM{}
	decoder := cdx.NewBOMDecoder(reader, cdx.BOMFileFormatJSON)
	if err := decoder.Decode(&bom); err != nil {
		return err
	}

	if len(*bom.Components) > 0 {
		return fmt.Errorf("expected no external dependencies, found %d", len(*bom.Components))
	}

	util.DPrintf("no external dependencies detected")

	return nil
}
