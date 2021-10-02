// config.go
// 1. 打开配置文件
// 2. 调用codec将配置输入Config结构体

// TODO: 能不能把decodeMap优化一下？
// TODO: 大的开源项目用proto是怎麼做的？
// TODO: 配置文件的validate怎么做？放在哪里做？

package config

import "io/ioutil"

// type: 配置文件的格式
type Type string

// DecodeFunc: 解析相应格式的配置文件到v中
// v: 空的map[string]interface{}
type DecodeFunc func(data []byte, v interface{}) error

// decodeMap提供type与DecodeFunc之间的映射
var decodeMap map[Type]DecodeFunc

func init() {
	decodeMap = make(map[Type]DecodeFunc)
	decodeMap["yaml"] = decodeYaml
}

func Load(path string, codec Type, cfg interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return decodeMap[codec](data, cfg)
}
