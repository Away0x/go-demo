# 依赖
```bash
# 开发时热加载
GO111MODULE=on go install github.com/cosmtrek/air@latest
# 测试邮件服务器
GO111MODULE=on go install github.com/mailhog/MailHog@latest
```

# 一些工具库
- gin —— 路由、路由组、中间件
- zap —— 高性能日志方案
- gorm —— ORM 数据操作
- cobra —— 命令行结构
- viper —— 配置信息
- cast —— 类型转换
- redis —— Redis 操作
- jwt —— JWT 操作
- base64Captcha —— 图片验证码
- govalidator —— 请求验证器
- limiter —— 限流器
- email —— SMTP 邮件发送
- aliyun-communicate —— 发送阿里云短信
- ansi —— 终端高亮输出
- strcase —— 字符串大小写操作
- pluralize —— 英文字符单数复数处理
- faker —— 假数据填充
- imaging —— 图片裁切

# 自定义的库
- app —— 应用对象
- auth —— 用户授权
- cache —— 缓存
- captcha —— 图片验证码
- config —— 配置信息
- console —— 终端
- database —— 数据库操作
- file —— 文件处理
- hash —— 哈希
- helpers —— 辅助方法
- jwt —— JWT 认证
- limiter —— API 限流
- logger —— 日志记录
- mail —— 邮件发送
- migrate —— 数据库迁移
- paginator —— 分页器
- redis —— Redis 数据库操作
- response —— 响应处理
- seed —— 数据填充
- sms —— 发送短信
- str —— 字符串处理
- verifycode —— 数字验证码
