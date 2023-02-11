package controller

import (
	"net/http"
	"strconv"

	"github.com/alitdarmaputra/belajar-golang-rest-api/helper"
	"github.com/alitdarmaputra/belajar-golang-rest-api/model/web"
	"github.com/alitdarmaputra/belajar-golang-rest-api/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	categoryCreateRequest := web.CategoryCreateRequest{}

	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}

	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = int64(id)

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), int64(id))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), int64(id))
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
