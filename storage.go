package main

import (
	"flag"
	"fmt"
	"net/http"

	"gitlab.energy-envision.com/storage/internal/config"
	"gitlab.energy-envision.com/storage/internal/handler"
	"gitlab.energy-envision.com/storage/internal/middleware"
	"gitlab.energy-envision.com/storage/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	_ "net/http/pprof"
)

var configFile = flag.String("f", "etc/storage-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", c.PprofPort), nil)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.NewAuthorization(c).AuthorizationHandle)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
