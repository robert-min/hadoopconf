package router

import (
	"github.com/gorilla/mux"
	"github.com/hadoopconf/go-server/middleware"
)

func MakeHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/hdfs", middleware.GetAllHdfsConfig).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/core", middleware.GetAllCoreConfig).Methods("GET", "OPTIONS")

	return router
}
