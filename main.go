package main

import (
	"net/http"

	"github.com/alitdarmaputra/belajar-golang-rest-api/helper"
	"github.com/alitdarmaputra/belajar-golang-rest-api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializeServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
