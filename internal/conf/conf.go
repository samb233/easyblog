// conf.go
// 应用配置的定义

package conf

type Bootstrap struct {
	Repo   *Repo
	Server *Server
}

type Repo struct {
	Database *Database
	// Redis    Redis
}

type Server struct {
	Addr string
}

type Database struct {
	Driver string
	Source string
}
