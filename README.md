# DouYin


- api_client：gin API 接口(兼 直接上传视频)    8080
- basic_server：视频 Feed 流、视频投稿、个人信息    50051
- interaction_server：点赞列表、用户评论    50052
- relation_server：关注列表、粉丝列表    50053


接口文档地址：https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145

## 使用技术栈

gin , gorm , grpc , jwt/go , protoc-gen-go , protoc-gen-go-grpc
## 开始使用

安装安装包：app-release.apk

打开四个控制台分别输入 `go run api-client/main.go` 和 `go run basic-server/main.go`  和   `go run interaction-server/main.go` 和  `go run relation-server/main.go`



