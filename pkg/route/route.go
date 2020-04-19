package route

import (
	"github.com/gorilla/mux"
)

var R *mux.Router

func Router() *mux.Router {
	R := mux.NewRouter()
	return R
}
