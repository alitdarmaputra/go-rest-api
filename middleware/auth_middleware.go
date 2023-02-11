package middleware

import (
	"net/http"

	"github.com/alitdarmaputra/belajar-golang-rest-api/helper"
	"github.com/alitdarmaputra/belajar-golang-rest-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorize",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
