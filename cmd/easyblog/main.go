// main.go
// 程序初始化，运行

package main

import (
	"log"
	"net/http"

	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/repo"
	"github.com/samb233/easyblog/internal/server"
	"github.com/samb233/easyblog/internal/service"
	"github.com/samb233/easyblog/internal/usecase"
	"github.com/samb233/easyblog/pkg/config"
)

func main() {
	cfg, err := readConf()
	if err != nil {
		panic(err)
	}
	srv, cleanUp, err := initApp(cfg.Repo, cfg.Server)
	if err != nil {
		panic(err)
	}

	defer cleanUp()

	// TODO： “优雅退出”，这里Fatal的话上面defer就没用了
	log.Fatal(srv.ListenAndServe())
}

// TODO: 最好还是有个logger吧？
// TODO: 使用wire框架自动化生成依赖
func initApp(confDatabase *conf.Repo, confServer *conf.Server) (*http.Server, func(), error) {
	database, cleanUp, err := repo.NewData(confDatabase)
	if err != nil {
		return nil, nil, err
	}

	indexRepo := repo.NewIndexRepo(database)
	contentRepo := repo.NewContentRepo(database)
	indexUsecase := usecase.NewIndexUsecase(indexRepo)
	contentUsecase := usecase.NewContentUsecase(contentRepo)
	blogService := service.NewBlogService(indexUsecase, contentUsecase)

	srv := server.NewHTTPServer(confServer, blogService)

	return srv, cleanUp, nil
}

func readConf() (*conf.Bootstrap, error) {
	cfg := &conf.Bootstrap{}
	err := config.LoadFromFile("configs/config.yaml", config.Yaml, cfg)
	return cfg, err
}
