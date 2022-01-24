# 依赖
```bash
# 开发时热加载
GO111MODULE=on go install github.com/cosmtrek/air@latest
# 测试邮件服务器
GO111MODULE=on go install github.com/mailhog/MailHog@latest
```

# 一些工具库
- go-pluralize: 处理英文单复数
- strcase: 处理大小写
- lumberjack: 一套滚动日志的实现方案，帮助我们管理日志文件
- ansi: 设置终端输出颜色
- faker: 生成假数据