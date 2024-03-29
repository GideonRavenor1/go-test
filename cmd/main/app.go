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

const Socket = "sock"
const Port = "port"
const SocketFileName = "app.sock"
const Unix = "unix"
const TCP = "tcp"

func main() {
	logger := logging.GetLogger()
	logger.Info("create router...")
	router := httprouter.New()

	conf := config.GetConfig()

	//mongoDBClient, err := mongodb.NewClient(
	//	context.Background(), conf.MongoDB.Host, conf.MongoDB.Port,
	//	conf.MongoDB.Username, conf.MongoDB.Password, conf.MongoDB.Database, conf.MongoDB.AuthDB,
	//)
	//utils.PanicHandler(err)

	//storage := db.NewStorage(mongoDBClient, conf.MongoDB.Collection, logger)

	logger.Info("register user handler...")
	handler := user.NewHandler(logger)
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
	if conf.Listen.Type == Port {
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
	if conf.Listen.Type == Socket {
		appDIr, err := filepath.Abs(filepath.Dir(os.Args[0]))
		utils.PanicHandler(err)
		logger.Info("create socket...")
		socketPath := path.Join(appDIr, SocketFileName)
		logger.Debugf("socker path: %s", socketPath)
		logger.Info("listen unix socket...")
		listener, listenerErr = net.Listen(Unix, socketPath)
	} else {
		logger.Info("listen tcp socket...")
		listener, listenerErr = net.Listen(TCP, fmt.Sprintf("%s:%s", conf.Listen.BindIP, conf.Listen.Port))
		utils.PanicHandler(listenerErr)
	}
	utils.PanicHandler(listenerErr)
	return listener
}
