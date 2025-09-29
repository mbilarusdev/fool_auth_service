package infrasturcture

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func RunServer(r *mux.Router) {
	addr := os.Getenv("FOOL_AUTH_SERVICE_ADDR")

	if err := http.ListenAndServe(addr, r); err != nil {
		errMsg := "Failed to listen and serve fool auth server!"
		LogErr(fmt.Errorf("%s: %v", errMsg, err))
		panic(errMsg)
	}
}
