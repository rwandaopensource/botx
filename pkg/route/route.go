package route

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Router the
func Router() *mux.Router {
	var origin = os.Getenv("ORIGIN")
	if origin == "" {
		origin = "*"
	}

	r := mux.NewRouter()

	// HOME ROUTE == PONG
	r.HandleFunc("/", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte("PONG"))
	}).Methods(http.MethodGet, http.MethodPost)

	// OPTIONS request
	r.Methods(http.MethodOptions).HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.WriteHeader(http.StatusNoContent)
	})

	r.HandleFunc("*", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	}).Methods(http.MethodGet, http.MethodPost)

	return r
}
