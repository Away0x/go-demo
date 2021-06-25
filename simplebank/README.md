# Postgres
> [Image](https://hub.docker.com/_/postgres)

```bash
# pull image
docker pull postgres:12-alpine

# create container(容器名 postgres12)
docker run --name postgres12 -p 5432:5432 \
    -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret \
    -d postgres:12-alpine

# 查看容器日志
docker logs postgres12

# 启动/停止
docker start/stop postgres12
```
```bash
# 进入 postgres (该容器默认在本地设置了信任身份验证，所以 localhost 连接时不需要密码)
docker exec -it postgres12 psql -U root
# shell
docker exec -it postgres12 /bin/sh


# 容器提供了一些命令方便 shell 与 postgres 交互,
#  - 创建数据库: createdb
#  - 删除数据库: dropdb
docker exec -it postgres12 createdb --username=root --owner=root demodb
docker exec -it postgres12 psql -U root demodb # 进入数据库
docker exec -it postgres12 dropdb demodb
```

# Migration
> [migrate](https://github.com/golang-migrate/migrate)

```bash
brew install golang-migrate
```
```bash
# 创建迁移
migrate create -ext sql -dir db/migration -seq init_schema

# 运行迁移
migrate -path db/migration \
    -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" \
    -verbose \ # 打印日志
    up
```

# SQLC
> [sqlc](https://github.com/kyleconroy/sqlc)
