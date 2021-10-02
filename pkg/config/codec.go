// yaml.go
// 提供yaml文件的DecodeFunc方法

package config

import "gopkg.in/yaml.v3"

func decodeYaml(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}
