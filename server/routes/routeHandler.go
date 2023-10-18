package routes

import (
	"net/http"

	db "github.com/hktrib/RetailGo/db/sqlc"
	"github.com/hktrib/RetailGo/util"
)

type RouteHandler struct {
	Store  *db.Store
	Config *util.Config
}

func (rh *RouteHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World -> RetailGo!!"))
}
