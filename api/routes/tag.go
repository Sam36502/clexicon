package routes

import (
	"clexicon-api/model"
	"clexicon-api/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func PutTag(ctx echo.Context) error {
	tagParam := model.Tag{}
	err := ctx.Bind(&tagParam)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_OBJECT, err)
	}

	id, err := model.G_TagBox.Put(&tagParam)
	if err != nil {
		if strings.HasPrefix(err.Error(), "Unique constraint") {
			return responses.GetAPIError(ctx, responses.APIERR_OBJECT_EXISTS, "tag", tagParam.Tag)
		}
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, responses.ModelCreationSuccess{
		ID: id,
	})
}

func GetTag(ctx echo.Context) error {
	// Try to get by id if present
	idStr := ctx.QueryParam("id")
	if idStr != "" {
		return getTagByID(ctx, idStr)
	}

	// Try to get by tag if present
	code := ctx.QueryParam("tag")
	if code != "" {
		return getTagByTag(ctx, code)
	}

	// Else, get all tags
	return getAllTags(ctx)
}

func getTagByID(ctx echo.Context, idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "'id' param must be a valid positive integer")
	}

	tag, err := model.G_TagBox.Get(id)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	if tag == nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "No Tag with that ID exists")
	}

	return ctx.JSON(http.StatusOK, tag)
}

func getTagByTag(ctx echo.Context, tag string) error {
	model.G_TagQuery.SetStringParams(model.Tag_.Tag, tag)
	tags, err := model.G_TagQuery.Find()
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	if len(tags) < 1 {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, tag, "No Tag with that code exists")
	}

	return ctx.JSON(http.StatusOK, tags[0])
}

func getAllTags(ctx echo.Context) error {
	tags, err := model.G_TagBox.GetAll()
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, tags)
}
