package config

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func GetOSSClient() {
	client, err := oss.New("Endpoint", "AccessKeyId", "AccessKeySecret")
}
