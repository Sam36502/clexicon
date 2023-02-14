package routes

import (
	"clexicon-api/model"
	"clexicon-api/responses"
	"fmt"
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
	idStr := ctx.QueryParam("id")
	if idStr != "" {
		return getWordByID(ctx, idStr)
	}

	return responses.GetAPIError(ctx, responses.APIERR_INVALID_PARAM, idStr, "Word-ID must be a valid positive integer")
}

func SearchWord(ctx echo.Context) error {
	lang, err := getLang(ctx)
	if err != nil {
		return err
	}

	queryStr := ctx.QueryParam("q")
	if queryStr != "" {

		// Search index for words matching
		query := bleve.NewFuzzyQuery(queryStr)
		query.SetFuzziness(2)

		idFloat := float64(lang.ID)
		inc := true
		langQ := bleve.NewNumericRangeInclusiveQuery(&idFloat, &idFloat, &inc, &inc)
		langQ.SetField("language.id")

		combQ := bleve.NewConjunctionQuery(langQ, query)

		search := bleve.NewSearchRequest(combQ)
		results, err := model.G_WordIndex.Search(search)
		if err != nil {
			return responses.GetAPIError(ctx, responses.APIERR_INDEX_FAIL, err)
		}

		// Fetch words that match the query
		words := make([]*model.Word, results.Total)
		for i, hit := range results.Hits {
			id, err := strconv.ParseUint(hit.ID, 10, 64)
			if err != nil {
				return responses.GetAPIError(
					ctx,
					responses.APIERR_INDEX_FAIL,
					fmt.Errorf("Failed to fetch word for index ID '%s'; invalid ID", hit.ID),
				)
			}

			word, err := model.G_WordBox.Get(id)
			if err != nil {
				return responses.GetAPIError(
					ctx,
					responses.APIERR_INDEX_FAIL,
					fmt.Errorf("Failed to fetch word for index ID '%s'; %v", hit.ID, err),
				)
			}

			if word == nil {
				return responses.GetAPIError(
					ctx,
					responses.APIERR_INDEX_FAIL,
					fmt.Errorf("Failed to fetch word for index ID '%s'; No word with that ID exists", hit.ID),
				)
			}

			words[i] = word
		}

		return ctx.JSON(http.StatusOK, words)
	}

	return ctx.JSON(http.StatusNoContent, []model.Word{})
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
