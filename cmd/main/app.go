package main

import (
	"fmt"
	"go_test/pkg/logging"
	"go_test/pkg/utils"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"go_test/internal/config"
	"go_test/internal/user"
	"github.com/julienschmidt/httprouter"
)

const SOCKET = "sock"
const PORT = "port"
const SOCKET_FILE_NAME = "app.sock"
const UNIX = "unix"
const TCP = "tcp"

func main() {
	logger := logging.GetLogger()
	logger.Info("create router...")
	router := httprouter.New()

	conf := config.GetConfig()
	handler := user.NewHandler(logger)
	logger.Info("register user handler...")
	handler.Register(router)
	start(router, conf)
}

func start(router *httprouter.Router, conf *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application...")
	listener := getListener(conf)
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if conf.Listen.Type == PORT {
		logger.Infof("server is listening %s:%s", conf.Listen.BindIP, conf.Listen.Port)
	} else {
		logger.Info("server is listening unix socket")
	}
	
	logger.Fatal(server.Serve(listener))
}


func getListener(conf *config.Config) net.Listener {
	logger := logging.GetLogger()
	var listener net.Listener
	var listenerErr error
	if conf.Listen.Type == SOCKET {
		appDIr, err := filepath.Abs(filepath.Dir(os.Args[0]))
		utils.ErrorHandler(err)
		logger.Info("create socket...")
		socketPath := path.Join(appDIr, SOCKET_FILE_NAME)
		logger.Debugf("socker path: %s", socketPath)
		logger.Info("listen unix socket...")
		listener, listenerErr = net.Listen(UNIX, socketPath)
	} else {
		var err error
		logger.Info("listen tcp socket...")
		listener, listenerErr = net.Listen(TCP, fmt.Sprintf("%s:%s", conf.Listen.BindIP, conf.Listen.Port))
		utils.ErrorHandler(err)
	}
	utils.ErrorHandler(listenerErr)
	return listener
}
