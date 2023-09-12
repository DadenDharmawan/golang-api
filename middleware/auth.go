package middleware

import (
	"net/http"

	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") == "SECRET" {
		// ok
		middleware.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	
		webResponse := web.WebResponse{
			Code: http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
	
		helper.WriteToResponseBody(w, webResponse)
	}
}
