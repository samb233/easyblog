// data.go
// 提供数据库连接

package data

import (
	"fmt"

	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/data/ent"
)

type Data struct {
	db *ent.Client

	// TODO: redis
}

func NewData(conf *conf.Database) (*Data, func(), error) {
	// TODO: read from config
	dsn := conf.Source
	driver := conf.Driver

	client, err := ent.Open(driver, dsn)
	if err != nil {
		return nil, nil, err
	}

	database := &Data{
		db: client,
	}

	return database, func() {
		fmt.Println("close db")
		if err := database.db.Close(); err != nil {
			fmt.Println("close db error: ", err)
		}
	}, nil
}
