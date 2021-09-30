package router

import (
	"github.com/7-Rohan/HW-Server_Deployment/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", middleware.Home)
	router.HandleFunc("/{userid}", middleware.GetByUserid)
	router.HandleFunc("/follower/{username}", middleware.GetByUsername)
	return router
}
