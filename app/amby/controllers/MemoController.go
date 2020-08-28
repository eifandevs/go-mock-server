package controllers

import (
	"net/http"

	"github.com/eifandevs/amby/models"
	"github.com/labstack/echo"
)

func GetMemoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Access-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.GetMemoResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		user, err := models.GetUser(accessToken)
		if err != nil {
			return c.JSON(http.StatusOK, models.GetMemoResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		memos := models.GetMemo(user.ID)

		return c.JSON(http.StatusOK, memos)
	}
}

func PostMemoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		post := new(models.PostMemoRequest)
		if err := c.Bind(post); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("Access-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Code: "", ErrorMessage: ""})
		}

		user, err := models.GetUser(accessToken)
		if err != nil {
			return c.JSON(http.StatusOK, models.GetMemoResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		response := models.PostMemo(user.ID, *post)
		return c.JSON(http.StatusOK, response)
	}
}

func DeleteMemoHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		delete := new(models.DeleteMemoRequest)
		if err := c.Bind(delete); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("Access-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Code: "", ErrorMessage: ""})
		}

		user, err := models.GetUser(accessToken)
		if err != nil {
			return c.JSON(http.StatusOK, models.GetMemoResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		response := models.DeleteMemo(user.ID, *delete)
		return c.JSON(http.StatusOK, response)
	}
}
