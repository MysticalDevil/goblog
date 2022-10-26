# goblog
Native go language practice project
## 本项目采用 MVC 架构设计
其中 Controller 层和 View 层 基于原生 go 语言编写

## 其他第三方模块
本项目中采用的第三方模块如下
- gorilla/mux 用来替代原生路由，目的是支持命名路由和正则路由
- gorilla/sessions 使用 Session 来支持用户登录
- lib/pq 支持 postgresql
- spf13/viper 配置文件读取
- stretchr/testify 单元测试支持
- thedevsaddam/govalidator 输入字段验证支持
- crypto 密码加密
- gorm 模型层支持
