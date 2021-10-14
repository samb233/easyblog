# easyblog

普通的，并不是很好用的博客项目

如果你正在寻找一个好用的博客网站，我推荐使用[Hugo](https://github.com/gohugoio/hugo)。

另外[ArkInfra](https://github.com/samb233/ArkInfra)也在开坑了！

## 计划

- [x] 搭建仿clean-arch和ddd的项目结构
- [x] 一些工具库的编写
  - [x] 日志库
  - [x] 配置文件读取
  - [ ] error处理
  - [x] 工具库单元测试
- [x] api设计
  - [x] api层的设计
  - [x] 基本的CRUD功能
  - [x] api接口的设计
- [ ] 添加redis作为缓存
  - [ ] 一致性
- [ ] 添加新功能
  - [ ] tag与分类
- [ ] 业务上的一些思考
  - [ ] 学习前端，了解如何做一个用的爽的文章增删改查功能

## 业务定义

- 获取文章列表

  GET `/api/v1/articles?page=1`

- 获取文章

  GET `/api/v1/articles/:id`

- 新增文章

- 修改文章

- 删除文章

DELETE `/api/v1/articles/:id`
