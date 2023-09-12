package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductController interface {
	Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindbyCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}