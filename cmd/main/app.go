package main

import (
	"go_test/pkg/logging"
	"go_test/pkg/utils"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go_test/internal/user"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router...")
	router := httprouter.New()

	handler := user.NewHandler(logger)
	logger.Info("register user handler...")
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application...")
	port := ":1234"
	listener, err := net.Listen("tcp", port)
	utils.ErrorHandler(err)

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Infof("server is listening 0.0.0.0%s", port)
	logger.Fatal(server.Serve(listener))
}
