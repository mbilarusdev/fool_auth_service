package infrasturcture

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	. "github.com/mbilarusdev/fool_auth_service/internal/logger"
)

func RunServer(r *mux.Router) {
	addr := os.Getenv("FOOL_AUTH_SERVICE_ADDR")

	if err := http.ListenAndServe(addr, r); err != nil {
		PanicErrWithMsg(err, "Failed to listen and serve fool auth server!")
	}
}
