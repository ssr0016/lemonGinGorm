package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"golang/data/request"
	"golang/data/response"
	"golang/helper"
	"golang/service"
)

// Manage of the REST interface to the bussiness logic

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

// Create Controller
func (controller *TagsController) Create(ctx *gin.Context) {
	CreateTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&CreateTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(CreateTagsRequest)

	webResponse := response.Response{

		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Update Controller
func (controller *TagsController) Update(ctx *gin.Context) {
	UpdateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&UpdateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	UpdateTagsRequest.Id = id

	controller.tagsService.Update(UpdateTagsRequest)

	webResponse := response.Response{

		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Controller
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagsService.Delete(id)

	webResponse := response.Response{

		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindById Controller
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagsResponse := controller.tagsService.FindById(id)

	webResponse := response.Response{

		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagsResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll Controller
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagRespose := controller.tagsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagRespose,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
