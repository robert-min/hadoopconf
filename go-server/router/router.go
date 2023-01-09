package router

import (
	"github.com/gorilla/mux"
	"github.com/hadoopconf/go-server/middleware"
)

func MakeHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/task", middleware.GetAllConfig).Methods("GET", "OPTIONS")

	return router
}
