// repo.go
// connect to database and redis

package repo

import (
	"github.com/go-redis/redis/v8"
	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/repo/ent"
	"github.com/samb233/easyblog/pkg/log"

	_ "github.com/go-sql-driver/mysql"
)

type Repo struct {
	db  *ent.Client
	rdb *redis.Client
}

func NewRepo(conf *conf.Repo, logger log.Logger) (*Repo, func(), error) {
	log := log.WithFields(logger, log.Fields{
		"file": "repo/repo.go",
	})

	// database
	dsn := conf.Database.Source
	driver := conf.Database.Driver

	db, err := ent.Open(driver, dsn)
	if err != nil {
		return nil, nil, err
	}

	// redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
	})

	repo := &Repo{
		db:  db,
		rdb: rdb,
	}

	return repo, func() {
		log.Info("closing repositories: ")

		if err := repo.db.Close(); err != nil {
			log.Errorf("close database error: %v", err)
		} else {
			log.Info("close database success")
		}

		if err := repo.rdb.Close(); err != nil {
			log.Errorf("close redis client error: %v", err)
		} else {
			log.Info("close redis success")
		}
	}, nil
}
