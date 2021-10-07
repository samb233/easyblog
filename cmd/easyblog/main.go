// main.go
// 程序初始化，运行

package main

import (
	"net/http"
	"os"

	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/repo"
	"github.com/samb233/easyblog/internal/server"
	"github.com/samb233/easyblog/internal/service"
	"github.com/samb233/easyblog/internal/usecase"
	"github.com/samb233/easyblog/pkg/config"
	"github.com/samb233/easyblog/pkg/log"
	"github.com/sirupsen/logrus"
)

func main() {
	// init logger
	logger := log.NewLogger(os.Stdout, log.LogFormat(&logrus.JSONFormatter{}))
	log := log.WithFields(logger, log.Fields{
		"app": "easyblog",
	})

	// read conf
	cfg, err := readConf()
	if err != nil {
		log.Errorf("read conf error : %v", err)
		return
	}

	// init app
	srv, cleanUp, err := initApp(cfg.Repo, cfg.Server, log)
	if err != nil {
		log.Errorf("initApp error: %v", err)
	}

	defer cleanUp()

	// serve
	err = srv.ListenAndServe()
	if err != nil {
		log.Errorf("server error: %v", err)
	}
}

// TODO: use wire to generate automatically
func initApp(confDatabase *conf.Repo, confServer *conf.Server, logger log.Logger) (*http.Server, func(), error) {
	database, cleanUp, err := repo.NewRepo(confDatabase, logger)
	if err != nil {
		return nil, nil, err
	}

	indexRepo := repo.NewIndexRepo(database, logger)
	contentRepo := repo.NewContentRepo(database, logger)
	indexUsecase := usecase.NewIndexUsecase(indexRepo)
	contentUsecase := usecase.NewContentUsecase(contentRepo)
	blogService := service.NewBlogService(indexUsecase, contentUsecase)

	srv := server.NewHTTPServer(confServer, blogService)

	return srv, cleanUp, nil
}

func readConf() (*conf.Bootstrap, error) {
	cfg := &conf.Bootstrap{}
	err := config.LoadFromFile("../../configs/config.yaml", config.Yaml, cfg)
	return cfg, err
}
