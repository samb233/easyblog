// conf.go
// 应用配置的定义

package conf

type Bootstrap struct {
	Repo    *Repo
	Server  *Server
	Service *Service
}

//------------------------------------------------------------------------------

// 持久化设置
type Repo struct {
	Database *Database
	Redis    Redis
}

// 数据库设置
type Database struct {
	Driver string
	Source string
}

// redis设置
type Redis struct {
	Addr     string
	Password string
}

//------------------------------------------------------------------------------

// 服务器设置
type Server struct {
	Addr string
}

//------------------------------------------------------------------------------

// 服务设置（业务设置）
type Service struct {
	Pagesize int
}

func (b *Bootstrap) Check() error {
	// 检查关键配置是否为空
	// TODO:

	if b.Service == nil {
		b.Service = &Service{
			Pagesize: 50,
		}
	}
	return nil
}
