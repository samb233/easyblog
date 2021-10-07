// codec.go

package config

import "gopkg.in/yaml.v3"

// yaml
func decodeYaml(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
