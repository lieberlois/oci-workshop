package resolver

import (
	"workshop/types"
	"workshop/util"
)

type FSPluginResolver struct{}

func (r FSPluginResolver) Resolve(name string) (types.DecoderFunc, error) {
	util.DPrintf("FSPluginResolver: Resolving %s...", name)
	return loadDecoderFuncFromPlugin(name)
}
