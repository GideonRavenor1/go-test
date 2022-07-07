package user

import (
	"go_test/pkg/logging"
	"go_test/pkg/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go_test/internal/handlers"
)

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
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
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is list of users"))
	utils.ErrorHandler(err)
	h.logger.Info("get list users")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is get user"))
	utils.ErrorHandler(err)
	h.logger.Info("get user by uuid")
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("this is create user"))
	utils.ErrorHandler(err)
	h.logger.Info("create user")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is update user"))
	utils.ErrorHandler(err)
	h.logger.Info("update user")
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is pathupdate user"))
	utils.ErrorHandler(err)
	h.logger.Info("partially update user")
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte("this is delete user"))
	utils.ErrorHandler(err)
	h.logger.Info("delete user")
}
