package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go_test/internal/user"
)

func IndexHendler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
}

func main() {
	log.Println("create router...")
	router := httprouter.New()

	handler :=user.NewHandler()
	log.Println("register user handler...")
	handler.Register(router)
	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application...")
	port := ":1234"
	listener, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Printf("server is listening 0.0.0.0%s", port)
	log.Fatalln(server.Serve(listener))
}
