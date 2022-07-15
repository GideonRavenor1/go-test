package user

import (
	"go_test/internal/apperror"
	"go_test/pkg/logging"
	"go_test/pkg/utils"
	"net/http"

	"go_test/internal/handlers"

	"github.com/julienschmidt/httprouter"
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
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is list of users"))
	utils.PanicHandler(err)
	h.logger.Info("get list users")

	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is get user"))
	utils.PanicHandler(err)
	h.logger.Info("get user by uuid")

	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("this is create user"))
	utils.PanicHandler(err)
	h.logger.Info("create user")

	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is update user"))
	utils.PanicHandler(err)
	h.logger.Info("update user")

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is pathupdate user"))
	utils.PanicHandler(err)
	h.logger.Info("partially update user")

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte("this is delete user"))
	utils.PanicHandler(err)
	h.logger.Info("delete user")

	return nil
}
