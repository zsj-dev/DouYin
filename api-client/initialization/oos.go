package initialization

import (
	"log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zsj-dev/DouYin/api-client/conf"
)

func RegisterOSS() {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "http://oss-cn-hangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := ""
	accessKeySecret := ""
	bucketName := "byte-douyin"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Panicf("初始化 Oss 异常: %v", err)
	}
	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Panicf("初始化 Oss 异常: %v", err)
	}
	conf.OSS = bucket
}
