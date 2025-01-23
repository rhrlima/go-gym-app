package controller

import (
	"gym-app/model"
	"gym-app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tagController struct {
	tagUsecase usecase.TagUsecase
}

func NewTagController(usecase usecase.TagUsecase) tagController {
	return tagController{
		tagUsecase: usecase,
	}
}

func (tc *tagController) CreateTag(ctx *gin.Context) {

	var tag model.Tag
	err := ctx.BindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedTag, err := tc.tagUsecase.CreateTag(tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedTag)
}

func (tc *tagController) GetTags(ctx *gin.Context) {

	tags, err := tc.tagUsecase.GetTags()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, tags)
}

func (tc *tagController) GetTagByName(ctx *gin.Context) {
	tagName := ctx.Param("tagName")

	if tagName == "" {
		response := model.Response{
			Message: "Tag name cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tag, err := tc.tagUsecase.GetTagByName(tagName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if tag == nil {
		response := model.Response{
			Message: "Tag Name not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, tag)
}
