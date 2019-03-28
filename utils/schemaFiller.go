package utils

import (
	"encoding/json"

	"github.com/hbagdi/go-kong/kong"
	"github.com/tidwall/gjson"
)

func fill(schema map[string]interface{},
	config kong.Configuration) (kong.Configuration, error) {
	if config == nil {
		return nil, nil
	}
	res := config.DeepCopy()
	jsonb, err := json.Marshal(&schema)
	if err != nil {
		return nil, err
	}
	// Get all the fields in the schema
	value := gjson.Get(string(jsonb), "fields")

	value.ForEach(func(key, value gjson.Result) bool {
		// get the key name
		ms := value.Map()
		fname := ""
		for k := range ms {
			fname = k
			break
		}
		// check if key is already set in the config
		if _, ok := config[fname]; ok {
			// yes, don't set it
			return true
		}
		// no, set it
		value = value.Get(fname + ".default")
		if value.Exists() {
			res[fname] = value.Value()
		} else {
			// if no default exists, set an explicit nil
			res[fname] = nil
		}
		return true
	})

	return res, nil
}
