// conf.go
// 读取配置文件

package conf

type Config struct {
	Database database
	Server   server
}

type database struct {
	source string
	driver string
}

type server struct {
	addr string
}

func NewConf() *Config {
	return &Config{}
}
