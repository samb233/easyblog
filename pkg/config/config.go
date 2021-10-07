// config.go

package config

import (
	"fmt"
	"io/ioutil"
)

// codec type of config file. yaml etc.
type Type string

// supported codec types
const (
	Yaml Type = "yaml"
)

// decode func defination
// TODO: 这种直接将配置解码到config结构体的做法存在问题
// 例如yaml，当字段名与结构体名不匹配时，会将传递空指针到结构体，并不会报错，在使用时会panic
// 关键信息缺失时，应该在这里报错而不是在使用时panic
// 两种思路
// 设置default，使用时判断nil后赋值
// 或是default + option func 的模式
type DecodeFunc func(data []byte, v interface{}) error

var decodeMap map[Type]DecodeFunc

// load all decodeFunc into decodeMap
func init() {
	decodeMap = make(map[Type]DecodeFunc)
	decodeMap[Yaml] = decodeYaml
}

// check whether a decodeFunc exist
func ExistCodec(codec Type) bool {
	_, ok := decodeMap[codec]
	return ok
}

// load from bytes
func Load(data []byte, codec Type, cfg interface{}) error {
	if !ExistCodec(codec) {
		return fmt.Errorf("codec not supported")
	}
	return decodeMap[codec](data, cfg)
}

// load from file
func LoadFromFile(filePath string, codec Type, cfg interface{}) error {
	if !ExistCodec(codec) {
		return fmt.Errorf("codec not supported")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return Load(data, codec, cfg)
}
