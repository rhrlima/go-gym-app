package controller

import (
	"gym-app/internal/model"
	"gym-app/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagUsecase usecase.TagUsecase
}

func NewTagController(usecase usecase.TagUsecase) TagController {
	return TagController{
		tagUsecase: usecase,
	}
}

func (tc *TagController) GetTags(ctx *gin.Context) {

	tags, err := tc.tagUsecase.GetTags()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, tags)
}

func (tc *TagController) CreateTag(ctx *gin.Context) {

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

func (tc *TagController) UpdateTag(ctx *gin.Context) {

	var tag model.Tag
	err := ctx.BindJSON(&tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = tc.tagUsecase.UpdateTag(tag)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (tc *TagController) GetTagByID(ctx *gin.Context) {
	id := ctx.Param("tagId")

	if id == "" {
		response := model.Response{
			Message: "Tag Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tagId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Tag Id must be an integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tag, err := tc.tagUsecase.GetTagByID(tagId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if tag == nil {
		response := model.Response{
			Message: "Tag Id not found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (tc *TagController) GetTagByName(ctx *gin.Context) {
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

func (tc *TagController) GetTagsByExerciseID(ctx *gin.Context) {
	id := ctx.Param("exerciseId")

	if id == "" {
		response := model.Response{
			Message: "Exercise Id cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	exerciseId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Exercise Id must be an integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tags, err := tc.tagUsecase.GetTagsByExerciseID(exerciseId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (tc *TagController) DeleteTagByID(ctx *gin.Context) {

	id := ctx.Param("tagId")
	if id == "" {
		response := model.Response{
			Message: "Tag ID cannot be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	tagId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Tag ID must be an integer",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = tc.tagUsecase.DeleteTagByID(tagId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
