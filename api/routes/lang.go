package routes

import (
	"clexicon-api/model"
	"clexicon-api/responses"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func PutLang(ctx echo.Context) error {
	langParam := model.Lang{}
	err := ctx.Bind(&langParam)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_OBJECT, err)
	}

	id, err := model.G_LangBox.Put(&langParam)
	if err != nil {
		if strings.HasPrefix(err.Error(), "Unique constraint") {
			return responses.GetAPIError(ctx, responses.APIERR_OBJECT_EXISTS, "code", langParam.Code)
		}
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, responses.ModelCreationSuccess{
		ID: id,
	})
}

func GetLang(ctx echo.Context) error {

	// Try to get by id if present
	idStr := ctx.QueryParam("id")
	if idStr != "" {
		return getLangByID(ctx, idStr)
	}

	// Try to get by code if present
	code := ctx.QueryParam("code")
	if code != "" {
		return getLangByCode(ctx, code)
	}

	// Else, get all langs
	return getAllLangs(ctx)
}

func getLangByID(ctx echo.Context, idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "'id' param must be a valid positive integer")
	}

	lang, err := model.G_LangBox.Get(id)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	if lang == nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "No Language with that ID exists")
	}

	return ctx.JSON(http.StatusOK, lang)
}

func getLangByCode(ctx echo.Context, code string) error {
	model.G_LangCodeQuery.SetStringParams(model.Lang_.Code, code)
	langs, err := model.G_LangCodeQuery.Find()
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	if len(langs) < 1 {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, code, "No Language with that code exists")
	}

	return ctx.JSON(http.StatusOK, langs[0])
}

func getAllLangs(ctx echo.Context) error {
	langs, err := model.G_LangBox.GetAll()
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, langs)
}

// TODO: Add endpoint to build a "family-tree" from `ancestor-codes` field
