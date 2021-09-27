// config.go
// 配置文件结构的定义

// TODO: 直接定义在这里真的难受
// 寻找更适合的方法/地方

package conf

type Bootstrap struct {
	Data   Data
	Server Server
}

type Data struct {
	Database Database
	// Redis    Redis
}

type Server struct {
	Addr string
}

type Database struct {
	Driver string
	Source string
}
