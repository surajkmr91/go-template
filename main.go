package main

import (
	"context"
	"fmt"

	"github.com/surajkmr91/go-template/api/routes"
	"github.com/surajkmr91/go-template/commons/config"
	"github.com/surajkmr91/go-template/commons/flags"
	"github.com/surajkmr91/go-template/commons/log"
)

var ctx = context.Background()
var appConfigName = "application.yml"

func init() {
	log.Info(ctx).Msg("bootstrapping the template service ....")
}

// initialising config yml file
func init() {
	err := config.InitConfig(flags.BaseConfigPath(), appConfigName)
	if err != nil {
		log.Fatal(ctx).Err(err).Msg("failed to initilize config")
	}
}

func main() {
	router := routes.DefaultRouter(ctx)
	log.Info(ctx).Msgf("listening to port: %d", flags.Port())
	err := router.Run(fmt.Sprintf(":%d", flags.Port()))
	if err != nil {
		log.Fatal(ctx).Err(err).Msg("server startup failed")
	}

}
