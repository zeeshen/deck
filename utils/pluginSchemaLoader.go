package utils

import (
	"net/http"

	"github.com/hbagdi/go-kong/kong"
	"github.com/pkg/errors"
)

// PluginSchemaStore retrives a schema of a Plugin from Kong.
type PluginSchemaStore struct {
	config  KongClientConfig
	client  *kong.Client
	schemas map[string]map[string]interface{}
}

// NewPluginSchemaStore creates a PluginSchemaStore.
func NewPluginSchemaStore(config KongClientConfig) (*PluginSchemaStore, error) {
	client, err := GetKongClient(config)
	if err != nil {
		return nil, err
	}
	return &PluginSchemaStore{
		config:  config,
		client:  client,
		schemas: make(map[string]map[string]interface{}),
	}, nil
}

// Schema retrives schema of a plugin.
// A cache is used to save the responses and subsequent queries are served from
// the cache.
func (p *PluginSchemaStore) Schema(pluginName string) (map[string]interface{}, error) {
	// TODO add a force refetch option
	if pluginName == "" {
		return nil, errors.New("pluginName can not be empty")
	}

	// lookup in cache
	if schema, ok := p.schemas[pluginName]; ok {
		return schema, nil
	}

	// not present in cache, lookup
	req, err := http.NewRequest("GET",
		p.config.Address+"/plugins/schema/"+pluginName, nil)
	if err != nil {
		return nil, err
	}
	schema := make(map[string]interface{})
	_, err = p.client.Do(nil, req, &schema)
	if err != nil {
		return nil, err
	}
	return schema, nil
}
