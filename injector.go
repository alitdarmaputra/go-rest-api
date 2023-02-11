//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/alitdarmaputra/belajar-golang-rest-api/app"
	"github.com/alitdarmaputra/belajar-golang-rest-api/controller"
	"github.com/alitdarmaputra/belajar-golang-rest-api/middleware"
	"github.com/alitdarmaputra/belajar-golang-rest-api/repository"
	"github.com/alitdarmaputra/belajar-golang-rest-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
