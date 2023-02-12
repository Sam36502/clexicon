package routes

import (
	"clexicon-api/model"
	"clexicon-api/responses"
	"net/http"

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
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, responses.ModelCreationSuccess{
		ID: id,
	})
}

func GetLangs(ctx echo.Context) error {

	langs, err := model.G_LangBox.GetAll()
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, langs)
}
