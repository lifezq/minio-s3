package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type MinioConfig struct {
	ServerName string
	Endpoint   string
	AccessKey  string
	SecretKey  string
	UseSSL     bool
}

type Config struct {
	rest.RestConf
	CacheConf     cache.CacheConf
	Minio         MinioConfig
	Datasource    string
	StorageEngine string
}
