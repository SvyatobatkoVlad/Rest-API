package user

import (
	"github.com/SvyatobatkoVlad/Rest-API/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/users"
	userURL = "/users:uuid"
	)

type handler struct{

}

func NewHandler() handlers.Handler{
	return &handler{}
}

func(h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("this is list of users"))
}

func (h *handler)GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("This need TO DO"))
}
func (h *handler)CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("This need TO DO"))
}
func (h *handler)UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("This need TO DO"))
}
func (h *handler)PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("This need TO DO"))
}
func (h *handler)DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("This need TO DO"))
}