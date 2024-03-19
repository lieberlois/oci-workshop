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
	"oras.land/oras-go/v2/registry"
	"oras.land/oras-go/v2/registry/remote"
)

type OCIResolverConfig struct {
	Hostname     string
	Port         string
	PluginDir    string
	ValidateSbom bool
}

type OCIPluginResolver struct {
	config OCIResolverConfig
}

type Option func(*OCIResolverConfig)

func WithHostname(hostname string) Option {
	return func(opts *OCIResolverConfig) {
		opts.Hostname = hostname
	}
}

func WithPort(port string) Option {
	return func(opts *OCIResolverConfig) {
		opts.Port = port
	}
}

func WithPluginDir(pluginDir string) Option {
	return func(opts *OCIResolverConfig) {
		opts.PluginDir = pluginDir
	}
}

func WithValidateSbom(validateSbom bool) Option {
	return func(opts *OCIResolverConfig) {
		opts.ValidateSbom = validateSbom
	}
}

func NewOCIPluginResolver(options ...Option) (*OCIPluginResolver, func()) {
	r := &OCIPluginResolver{}

	for _, opt := range options {
		opt(&r.config)
	}

	cleanupFunc := func() {
		err := os.RemoveAll(r.config.PluginDir)
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

	util.DPrintf(
		"OCIPluginResolver: Pulling %s from registry %s:%s...",
		name,
		r.config.Hostname,
		r.config.Port,
	)
	err := r.pullOciArtifact(repoName, tag)

	if err != nil {
		return nil, err
	}

	artifactFolderPath := fmt.Sprintf("%s/%s", r.config.PluginDir, repoName)
	pluginPath := fmt.Sprintf("%s/decoder.so", artifactFolderPath)
	sbomPath := fmt.Sprintf("%s/sbom.json", artifactFolderPath)

	if r.config.ValidateSbom {
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

func (r OCIPluginResolver) pullOciArtifact(name string, tag string) error {
	fileStore, err := file.New(r.config.PluginDir + "/" + name) // note: folder per plugin for sbom etc.
	if err != nil {
		return err
	}
	defer fileStore.Close()

	ctx := context.Background() // TODO: move to central Context

	// Build Aritfact Ref: hostname:port/repo
	artifactRef := fmt.Sprintf("%s:%s/%s", r.config.Hostname, r.config.Port, name)

	repo, err := remote.NewRepository(artifactRef)
	if err != nil {
		panic(err)
	}

	repo.PlainHTTP = true

	// Fetch OCI artifact
	manifestDescriptor, err := oras.Copy(ctx, repo, tag, fileStore, tag, oras.DefaultCopyOptions)
	if err != nil {
		return err
	}

	// Get artifacts with subject pointing to main artifact
	refs, err := registry.Referrers(ctx, repo, manifestDescriptor, "goplugin/sbom")
	if err != nil {
		panic(err)
	}

	// Downloads all attached SBOM artifacts
	for _, ref := range refs {
		util.DPrintf("Resolving SBOM ref %s...", ref.Digest)

		_, err := oras.Copy(ctx, repo, ref.Digest.String(), fileStore, "", oras.DefaultCopyOptions)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
