package routes

import (
	"clexicon-api/model"
	"clexicon-api/responses"
	"net/http"
	"strconv"

	"github.com/blevesearch/bleve/v2"
	"github.com/labstack/echo/v4"
)

func PutWord(ctx echo.Context) error {
	lang, err := getLang(ctx)
	if err != nil {
		return err
	}

	wordParam := model.Word{}
	wordParam.Language = lang
	err = ctx.Bind(&wordParam)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_OBJECT, err)
	}

	newID, err := model.G_WordBox.Put(&wordParam)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	err = model.G_WordIndex.Index(strconv.FormatUint(wordParam.ID, 10), wordParam)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, responses.ModelCreationSuccess{
		ID: newID,
	})
}

func GetWord(ctx echo.Context) error {

	// Try with id, if present
	idStr := ctx.QueryParam("id")
	if idStr != "" {
		return getWordByID(ctx, idStr)
	}

	// Else search based on params
	model.G_WordSearchQuery.SetStringParams(model.Word_.Orthography, ctx.QueryParam("orth"))
	model.G_WordSearchQuery.SetStringParams(model.Word_.Romanisation, ctx.QueryParam("rom"))
	model.G_WordSearchQuery.SetStringParams(model.Word_.Pronunciation, ctx.QueryParam("ipa"))
	model.G_WordSearchQuery.SetStringParams(model.Word_.Meanings, ctx.QueryParam("m"))

	words, err := model.G_WordSearchQuery.Find()
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	return ctx.JSON(http.StatusOK, words)
}

func SearchWord(ctx echo.Context) error {
	/*
		lang, err := getLang(ctx)
		if err != nil {
			return err
		}
	*/

	queryStr := ctx.QueryParam("q")
	if queryStr != "" {
		query := bleve.NewFuzzyQuery(queryStr)
		query.SetFuzziness(2)
		search := bleve.NewSearchRequest(query)
		results, err := model.G_WordIndex.Search(search)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, results)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func getWordByID(ctx echo.Context, idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "'id' param must be a valid positive integer")
	}

	lang, err := model.G_WordBox.Get(id)
	if err != nil {
		return responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	if lang == nil {
		return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "No Word with that ID exists")
	}

	return ctx.JSON(http.StatusOK, lang)
}

func getLang(ctx echo.Context) (*model.Lang, error) {
	idStr := ctx.Param("lid")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "'lang_id' param must be a valid positive integer")
	}

	lang, err := model.G_LangBox.Get(id)
	if err != nil {
		return nil, responses.GetAPIError(ctx, responses.APIERR_MODEL_FAIL, err)
	}

	if lang == nil {
		return nil, responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "No Language with that ID exists")
	}

	return lang, nil
}

// TODO: Add Endpoint for "family-tree" of a word's etymology
