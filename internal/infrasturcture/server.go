package infrasturcture

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	. "github.com/mbilarusdev/fool_auth_service/internal/logger"
)

func RunServer(r *mux.Router) {
	addr := os.Getenv("AUTH_SERVICE_ADDR")

	LogInfo(fmt.Sprintf("Starting server on %v", addr))

	if err := http.ListenAndServe(addr, r); err != nil {
		PanicErrWithMsg(err, "Failed to listen and serve fool auth server!")
	}
}
