# DouYin


- api_client：gin API 接口(兼 直接上传视频)    8080
- basic_server：视频 Feed 流、视频投稿、个人信息    50051
- interaction_server：点赞列表、用户评论    50052
- relation_server：关注列表、粉丝列表    50053


接口文档地址：https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

视频演示地址：https://github.com/zsj-dev/DouYin

## 使用技术栈

gin , gorm , grpc , jwt/go , protoc-gen-go ,oos 等
## 开始使用

安装安装包：app-release.apk
### 设置服务端地址

为方便测试登录和注册，及修改网络请求的服
务器地址，提供了退出登录和高级设置两个能
力。

点击退出登录会自动重启

在高级设置中可以配置自己的服务端项目的，前缀地址，如http://192.168.1.7:8080，在app中访问上述某个接口时就会拼接该前缀地址，例如访问 http://192.168.1.7:8080/douyin/feed/ 拉取视频列表

另外在未登录情况下，双击右下角的“我”可以打开高级设置

### 服务器开启

打开四个控制台分别输入 `go run api-client/main.go` 和 `go run basic-server/main.go`  和   `go run interaction-server/main.go` 和  `go run relation-server/main.go`



