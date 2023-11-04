package main

import (
	"blaze/internal/todoservice"
	"blaze/pkg/util"
	"net/http"
)

func main() {
	router := todoservice.InitService()

	util.Log.Info("Server started at http://localhost:3000")
	http.ListenAndServe("0.0.0.0:3000", router)
}
