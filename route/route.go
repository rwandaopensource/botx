package route

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rwandaopensource/botx/controller"
)

// NewRouter ...
func NewRouter() *mux.Router {
	var origin = os.Getenv("ORIGIN")
	if origin == "" {
		origin = "*"
	}

	rt := mux.NewRouter()

	// HOME ROUTE == PONG
	rt.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte("PONG"))
	}).Methods(http.MethodGet, http.MethodPost)

	// OPTIONS request
	rt.Methods(http.MethodOptions).HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.WriteHeader(http.StatusNoContent)
	})

	rt.HandleFunc("/install", controller.Install)

	rt.HandleFunc("*", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	})

	return rt
}
