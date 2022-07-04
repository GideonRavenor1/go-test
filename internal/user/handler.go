package user

import (
	"net/http"
	"log"

	"github.com/julienschmidt/httprouter"
	"go_test/internal/handlers"
)

const (
	usersURL = "/users"
	userURL =  "/user/:uuid"
)

type handler struct {

}

func NewHandler() handlers.Handler{
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)


}


func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))
	log.Println("get list users")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("this is get user"))
	log.Println("get user by uuid")	
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(201)
	w.Write([]byte("this is create user"))
	log.Println("create user")	
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("this is update user"))
	log.Println("update user")
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(200)
	w.Write([]byte("this is pathupdate user"))
	log.Println("partially update user")
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))
	log.Println("delete user")
}