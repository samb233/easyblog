package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testYaml = `server:
  addr: 127.0.0.1:9999
repo:
  database:
    driver: mysql
    source: root:password@tcp(127.0.0.1:23306)/easyblog?parseTime=True`
)

type testConfigStruct struct {
	Server *Server
	Repo   *Repo
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

func TestLoad(t *testing.T) {
	var (
		addr   = "127.0.0.1:9999"
		driver = "mysql"
		source = "root:password@tcp(127.0.0.1:23306)/easyblog?parseTime=True"
	)

	// test load yaml
	t.Run("test load yaml", func(t *testing.T) {
		t.Helper()
		cfg := &testConfigStruct{}
		err := Load([]byte(testYaml), Yaml, cfg)
		if err != nil {
			t.Errorf("have an error when Load() %v", err)
			return
		}
		// fmt.Println(cfg)
		assert.Equal(t, addr, cfg.Server.Addr)
		assert.Equal(t, driver, cfg.Repo.Database.Driver)
		assert.Equal(t, source, cfg.Repo.Database.Source)
	})

	// test codec not exist
	t.Run("test codec not exist", func(t *testing.T) {
		t.Helper()
		errCfg := &testConfigStruct{}
		err := Load([]byte(testYaml), "lamy", errCfg)
		if err == nil {
			t.Errorf("should have en error but not")
		}
	})

}
