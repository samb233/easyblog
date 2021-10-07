// repo.go
// connect to database and redis

package repo

import (
	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/repo/ent"
	"github.com/samb233/easyblog/pkg/log"
)

type Repo struct {
	db *ent.Client

	// TODO: redis
}

func NewRepo(conf *conf.Repo, logger log.Logger) (*Repo, func(), error) {
	log := log.WithFields(logger, log.Fields{
		"file": "repo/repo.go",
	})
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
		log.Info("close database")
		if err := database.db.Close(); err != nil {
			log.Errorf("close database error: %v", err)
		}
	}, nil
}
