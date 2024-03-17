package resolver

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"workshop/types"
	"workshop/util"

	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
)

type OCIResolverConfig struct {
	Hostname  string
	Port      string
	PluginDir string
}

type OCIPluginResolver struct {
	OCIResolverConfig
}

func NewOCIPluginResolver(config OCIResolverConfig) (*OCIPluginResolver, func()) {
	r := &OCIPluginResolver{
		OCIResolverConfig: config,
	}

	cleanupFunc := func() {
		err := os.RemoveAll(r.OCIResolverConfig.PluginDir)
		if err != nil {
			log.Fatalf("Failed to cleanup OCI plugins directory: %s", err)
		}
	}

	return r, cleanupFunc
}

func (r OCIPluginResolver) Resolve(name string) (types.DecoderFunc, error) {
	nameSplit := strings.Split(name, ":")
	repoName := nameSplit[0]
	tag := nameSplit[1]

	util.DPrintf("OCIPluginResolver: Pulling %s from registry %s:%s...", name, r.Hostname, r.Port)
	err := r.pullOciArtifact(repoName, tag)

	if err != nil {
		return nil, err
	}

	pluginPath := fmt.Sprintf("%s/%s.so", r.PluginDir, repoName)
	return loadDecoderFuncFromPlugin(pluginPath)
}

func (r OCIPluginResolver) pullOciArtifact(name string, tag string) error {
	fileStore, err := file.New(r.PluginDir)
	if err != nil {
		panic(err)
	}
	defer fileStore.Close()

	ctx := context.Background()

	artifactRef := fmt.Sprintf("%s:%s/%s", r.Hostname, r.Port, name)

	repo, err := remote.NewRepository(artifactRef)
	if err != nil {
		panic(err)
	}

	repo.PlainHTTP = true
	_, err = oras.Copy(ctx, repo, tag, fileStore, tag, oras.DefaultCopyOptions)
	if err != nil {
		return err
	}

	return nil
}
