package app

import (
	"github.com/gorilla/mux"
	"github.com/mbilarusdev/fool_auth_service/internal/infrasturcture"
)

func RunApp() {
	infrasturcture.InitLogger()
	defer infrasturcture.SyncLogger()
	infrasturcture.PingRedis()
	r := mux.NewRouter()
	infrasturcture.RunServer(r)
}
