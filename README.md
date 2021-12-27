# blog-microservices

[![test](https://github.com/stonecutter/blog-microservices/actions/workflows/test.yaml/badge.svg)](https://github.com/stonecutter/blog-microservices/actions/workflows/test.yaml)

blog microservices deployed in an Istio-enabled kubernetes cluster.

### 架构

![architecture](./assets/architecture.png)

### 目录结构

主要遵循 [Standard Go Project Layout](https://github.com/golang-standards/project-layout) 推荐的目录分层。

### 使用的依赖:

* [gRPC](https://github.com/grpc/grpc-go) 微服务通信协议
* [GORM](https://github.com/jackc/pgx) 数据库 ORM
* [DTM](https://github.com/dtm-labs/dtm) 分布式事务管理器
* [Jaeger](https://www.jaegertracing.io/) 分布式追踪
* [Prometheus](https://prometheus.io/) 监控系统
* [Grafana](https://grafana.com/) 数据可视化
* [Kiali](https://kiali.io/) 可观察性工具
* [Kubernetes](https://kubernetes.io/) 容器编排
* [Istio](https://istio.io/) 服务网格

### Makefile 简介

| 命令                    | 说明                                           |
|-----------------------|----------------------------------------------|
| `make init`           | 安装各类 protoc-gen-* 、 wire 以及 migrate          |
| `make protoc`         | 基于 *.proto 文件，生成各类 *_pb.go                   |
| `make wire`           | 基于 wire.go 文件，生成 wire_gen.go                 |
| `make test`           | 测试                                           |
| `make migrate-up`     | 迁移数据库                                        |
| `make migrate-down`   | 回滚数据库                                        |
| `make blog-server`    | 启动 blog 服务                                   |
| `make user-server`    | 启动 user 服务                                   |
| `make post-server`    | 启动 post 服务                                   |
| `make comment-server` | 启动 comment 服务                                |
| `make auth-server`    | 启动 auth 服务                                   |
| `make dtm-server`     | DTM 为外部依赖，启动需浏览官方文档                          |
| `make docker-build`   | 构建 Docker 镜像                                 |
| `make kube-deploy`    | 在集群中部署 blog、user、post、comment、auth 以及 dtm 服务 |
| `make kube-delete`    | 在集群中删除上述服务                                   |
| `make kube-redeploy`  | 在集群中重新部署服务(数据库服务不删除)                         |



### 访问服务

推荐使用 [BloomRPC](https://github.com/bloomrpc/bloomrpc) 或者 [Insomnia](https://github.com/Kong/insomnia)