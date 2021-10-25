package svc

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/tal-tech/go-zero/core/logx"
	"gitlab.energy-envision.com/storage/internal/storage"
	"log"
	"os/exec"
	"time"

	"gitlab.energy-envision.com/storage/client"
	"gitlab.energy-envision.com/storage/internal/config"
	"gitlab.energy-envision.com/storage/internal/types"
	"gitlab.energy-envision.com/storage/model"

	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/syncx"
)

type ServiceContext struct {
	Config config.Config
	model.UserModel
	Cache  cache.Cache
	Client client.Client
	Loc    *time.Location
}

func NewServiceContext(c config.Config) *ServiceContext {

	loc, _ := time.LoadLocation("Asia/Shanghai")
	return &ServiceContext{
		Config: c,
		UserModel: model.NewUserModel(sqlx.NewSqlConn("mysql", c.Datasource),
			c.CacheConf),
		Cache: cache.New(c.CacheConf, syncx.NewSingleFlight(), cache.NewStat(types.CACHE_REDIS_STATE),
			sql.ErrNoRows, []cache.Option{}...),
		Client: getClientByStorageEngine(&c),
		Loc:    loc,
	}
}

func getClientByStorageEngine(c *config.Config) client.Client {

	logx.Infof("存储引擎[%s]加载中...", c.StorageEngine)

	switch c.StorageEngine {
	case types.ENGINE_LOCAL:
		return storage.NewLocalClient(c)
	case types.ENGINE_MINIO:
		return newMinioClient(c)
	}

	log.Fatalln("严重错误！！！存储引擎配置错误，程序退出...")
	return nil
}

func newMinioClient(c *config.Config) *minio.Client {

	// Initialize minio client object.
	client, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.SecretKey, ""),
		Secure: c.Minio.UseSSL,
	})
	if err != nil {
		log.Fatalf("minio connection fatal error: %s\n", err.Error())
	}

	out, err := exec.CommandContext(context.Background(), "mc",
		[]string{
			"alias", "set", c.Minio.ServerName, fmt.Sprintf("http://%s", c.Minio.Endpoint),
			c.Minio.AccessKey, c.Minio.SecretKey, "--api", "s3v4",
		}...).CombinedOutput()
	if err != nil {
		log.Fatalf("minio alias fatal error:%s %s\n", string(out), err.Error())
	}

	return client
}
