# [简单的分布式系统](https://www.bilibili.com/video/BV1ZU4y1577q?p=12&spm_id_from=pageDriver)
- 分布式模型:
    1. Hub & Spoke: 所有的服务都依赖一个中心服务, 有利于负载均衡, 方便做集中式的追踪和日志, 但是不好避免单点故障
    2. Peer to Peer: 点对点, 没有单点故障, 但是服务很难发现, 负载均衡比较困难
    3. Message Queues: 系统使用消息队列, 比较容易扩展, 易于消息的持久化, 但是存在单点故障的问题, 而且配置相对比较困难
- 该 Demo 使用的是混合模型
- 系统主要组件:
    - 服务注册:
        1. 服务注册
        2. 健康检查
    - 用户门户:
        1. Web 应用
        2. API 网关
    - 日志服务:
        1. 集中式日志
    - 业务服务:
        1. 业务逻辑
        2. 数据持久化
- 技术选型:
    - 语言: Go
    - 框架: 不使用框架
    - 数据传输: HTTP
    - 传输协议: JSON

```bash
# 运行项目
# 启动注册服务
cd cmd/registryservice
go run .

# 启动日志服务
cd cmd/logservice
go run .

# 启动业务服务 (学生成绩查询, 依赖 logservice)
cd cmd/gradingservice
go run .

# 启动 web 应用 (依赖 logservice & gradingservice)
cd cmd/portal
go run .
```
```bash
# 日志服务
http post localhost:4000/log a=1 b=2
# 服务注册服务
http post localhost:3000/services serviceName="Business Service" serviceURL="http://localhost:5000/business"
# 查询学生成绩
http get localhost:6000/students
```