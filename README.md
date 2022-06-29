# byte-douyin


- api_client：gin API 接口    8080
- basic_server：视频 Feed 流、视频投稿、个人信息    50051
- interaction_server：点赞列表、用户评论    50052
- relation_server：关注列表、粉丝列表    50053

安装 [protoc](https://github.com/protocolbuffers/protobuf)

```
$ protoc --version
libprotoc 3.19.4
```

## 下载 protoc-gen-go 和 protoc-gen-go-grpc

```shell
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### 生成
```
protoc --go_out=./pb --go-grpc_out=./pb pb/*.proto
```
## 支持库
add missing and remove unused modules

```
go mod tidy
```

## 开始使用
打开四个控制台分别输入 `go run api-client/main.go` 和 `go run basic-server/main.go`  和   `go run interaction-server/main.go` 和  `go run relation-server/main.go`



