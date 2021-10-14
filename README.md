# easyblog

普通的，并不是很好用的博客项目，主要用于实践类似于clean-arch和ddd的项目结构。使用`gin`作为`http`框架实现`RESTful API`，使用`ent`作为`orm`框架。数据库使用`mysql`。

如果你正在寻找一个好用的博客网站，我推荐使用[Hugo](https://github.com/gohugoio/hugo)。

## 项目结构

```
├── api
├── cmd
├── configs
├── internal
│   ├── conf
│   ├── domain
│   ├── repo
│   ├── server
│   ├── service
│   └── usecase
├── pkg
└── test

```

- `api`：`api`的定义，`gin`相关的所有逻辑都在这一层实现，只做表单验证，调用`internal/service`中的接口实现相应逻辑
- `cmd`：存放`main.go`
- `configs`：存放配置文件
- `internal`：存放应用代码
  - `domain`：领域模型建模，理应使用充血模型，但是博客网站实在是没什么逻辑，所以很空旷，只提供了`repo`接口和`usecase`接口
  - `repo`：持久化层。实现数据库与缓存的调用逻辑，实现`domain`层的`repo`接口。
  - `usecase`：用例层。只用来对上层提供`repo`层的使用用例。不应有复杂逻辑。
  - `service`：服务层。用于提供给网络服务调用的接口。做数据组装，调用`usecase`来访问`repo`。
  - `server`：调用`api`层的注册方法，连接`service`与`api`层。并提供`http server`用于启动。
  - `conf`：应用内配置。可以在这里实现`Check`方法检测配置正确与否。
- `pkg`：存放公用代码，在这里放了`log`与`config`包，分别用于日志输出与配置文件读取。`log`库使用`logrus`实现，为了松耦合，提供接口供应用调用，但这部分我做的不够好。
- `test`：存放测试相关文件，`sql`文件与`docker-compose.yml`等



## 使用

```bash
git clone https://github.com/samb233/easyblog.git
```

编译，`main.go`文件位于`cmd`目录下

```bash
go build -o easyblog cmd/easyblog/main.go
```

运行

```bash
./easyblog
```



## api

- 获取文章列表

  GET `http://localhost:9999/api/v1/articles?page=1`

- 获取文章

  GET `http://localhost:9999/api/v1/articles/1`

- 新增文章

  POST `http://localhost:9999/api/v1/articles`，可用 `form-data` 或 `x-www-form-urlencoded` ，其中`title`为`required`

- 修改文章

  POST `http://localhost:9999/api/v1/articles/1`

- 删除文章

  DELETE `http://localhost:9999/api/v1/articles/1`
