package app

import (
	"github.com/gorilla/mux"
	"github.com/mbilarusdev/fool_auth_service/internal/infrasturcture"
	. "github.com/mbilarusdev/fool_auth_service/internal/logger"
)

func RunApp() {
	InitLogger()
	defer SyncLogger()
	infrasturcture.ConnectPostgres()
	infrasturcture.PingRedis()
	r := mux.NewRouter()

	infrasturcture.RunServer(r)
}
