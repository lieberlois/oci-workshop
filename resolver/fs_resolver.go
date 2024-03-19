package resolver

import (
	"fmt"
	"os"
	"workshop/types"
	"workshop/util"
)

type FSPluginResolver struct {
	ValidateSbom bool
}

func NewFSPluginResolver(validateSbom bool) *FSPluginResolver {
	return &FSPluginResolver{
		ValidateSbom: validateSbom,
	}
}

func (r FSPluginResolver) Resolve(name string) (types.DecoderFunc, error) {
	pluginPath := fmt.Sprintf("%s/decoder.so", name)
	sbomPath := fmt.Sprintf("%s/sbom.json", name)

	util.DPrintf("FSPluginResolver: Resolving %s...", name)

	if r.ValidateSbom {
		file, err := os.Open(sbomPath)
		if err != nil {
			return nil, err
		}

		err = validateSbom(file)
		if err != nil {
			return nil, err
		}
	}

	return loadDecoderFuncFromPlugin(pluginPath)
}
