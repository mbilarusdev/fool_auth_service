package app

import (
	"github.com/gorilla/mux"
	"github.com/mbilarusdev/fool_auth_service/internal/utils"
	"github.com/mbilarusdev/fool_base/v2/infra/cache"
	"github.com/mbilarusdev/fool_base/v2/infra/db"
	"github.com/mbilarusdev/fool_base/v2/infra/network"
	"github.com/mbilarusdev/fool_base/v2/log"
)

func RunApp() {
	serviceName := "Auth Service"

	log.Init()
	defer log.Sync()

	utils.ParseConfig()

	db.ConnPGX(serviceName)

	cache.PingRDB()

	r := mux.NewRouter()

	gatewayAddr := network.GetGatewayAddr()
	gatewayMiddleware := network.GetCheckGatewayMiddleware(gatewayAddr)
	r.Use(gatewayMiddleware)

	network.Run(r, serviceName)
}
