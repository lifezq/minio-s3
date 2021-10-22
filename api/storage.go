package main

import (
	"flag"
	"fmt"
	"github.com/lifezq/minio-s3/api/internal/config"
	"github.com/lifezq/minio-s3/api/internal/handler"
	"github.com/lifezq/minio-s3/api/internal/middleware"
	"github.com/lifezq/minio-s3/api/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	_ "net/http/pprof"
)

var configFile = flag.String("f", "etc/storage-api.yaml", "the config file")

func main() {
	flag.Parse()

	go http.ListenAndServe("0.0.0.0:6060", nil)

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.NewAuthorization(c).AuthorizationHandle)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
