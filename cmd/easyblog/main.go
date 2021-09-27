// main.go
// 程序初始化，运行

package main

import (
	"log"
	"net/http"

	"github.com/samb233/easyblog/internal/conf"
	"github.com/samb233/easyblog/internal/data"
	"github.com/samb233/easyblog/internal/server"
	"github.com/samb233/easyblog/internal/service"
	"github.com/samb233/easyblog/internal/usecase"
)

func main() {
	conf := readConf()
	srv, cleanUp, err := initApp(&conf.Data.Database, &conf.Server)
	if err != nil {
		log.Fatal(err)
	}

	defer cleanUp()

	// TODO： “优雅退出”，这里Fatal的话上面defer就没用了
	log.Fatal(srv.ListenAndServe())
}

// TODO: 最好还是有个logger吧？
// TODO: 使用wire框架自动化生成依赖
func initApp(confDatabase *conf.Database, confServer *conf.Server) (*http.Server, func(), error) {
	database, cleanUp, err := data.NewData(confDatabase)
	if err != nil {
		return nil, nil, err
	}

	indexRepo := data.NewIndexRepo(database)
	contentRepo := data.NewContentRepo(database)
	indexUsecase := usecase.NewIndexUsecase(indexRepo)
	contentUsecase := usecase.NewContentUsecase(contentRepo)
	blogService := service.NewBlogService(indexUsecase, contentUsecase)

	srv := server.NewHTTPServer(confServer, blogService)

	return srv, cleanUp, nil
}

// TODO: 完善读取
func readConf() *conf.Bootstrap {
	return &conf.Bootstrap{}
}
