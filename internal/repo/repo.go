// repo.go
// connect to database and redis

package repo

import (
	"fmt"

	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/repo/ent"
)

type Repo struct {
	db *ent.Client

	// TODO: redis
}

func NewData(conf *conf.Repo) (*Repo, func(), error) {
	dsn := conf.Database.Source
	driver := conf.Database.Driver

	client, err := ent.Open(driver, dsn)
	if err != nil {
		return nil, nil, err
	}

	database := &Repo{
		db: client,
	}

	return database, func() {
		fmt.Println("close db")
		if err := database.db.Close(); err != nil {
			fmt.Println("close db error: ", err)
		}
	}, nil
}
