package providers

import (
	_ "embed"
	"github.com/loft-sh/devpod/pkg/provider"
	"github.com/loft-sh/devpod/pkg/provider/providerimplementation"
	"github.com/pkg/errors"
	"strings"
)

//go:embed docker/provider.yaml
var DockerProvider string

//go:embed gcloud/provider.yaml
var GCloudProvider string

// GetBuiltInProviders retrieves the built in providers
func GetBuiltInProviders() (map[string]provider.Provider, error) {
	providers := []string{DockerProvider, GCloudProvider}
	retProviders := map[string]provider.Provider{}

	// parse providers
	for _, providerConfig := range providers {
		parsedConfig, err := provider.ParseProvider(strings.NewReader(providerConfig))
		if err != nil {
			return nil, errors.Wrap(err, "parse provider")
		}

		retProviders[parsedConfig.Name] = providerimplementation.NewProvider(parsedConfig)
	}

	return retProviders, nil
}