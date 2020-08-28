package controllers

import (
	"net/http"

	"github.com/eifandevs/amby/models"
	"github.com/labstack/echo"
)

func GetFavoriteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := c.Request().Header.Get("Access-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.GetFavoriteResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		user, err := models.GetUser(accessToken)
		if err != nil {
			return c.JSON(http.StatusOK, models.GetFavoriteResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		favorites := models.GetFavorite(user.ID)

		return c.JSON(http.StatusOK, favorites)
	}
}

func PostFavoriteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		post := new(models.PostFavoriteRequest)
		if err := c.Bind(post); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("Access-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Code: "", ErrorMessage: ""})
		}

		user, err := models.GetUser(accessToken)
		if err != nil {
			return c.JSON(http.StatusOK, models.GetFavoriteResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		response := models.PostFavorite(user.ID, *post)
		return c.JSON(http.StatusOK, response)
	}
}

func DeleteFavoriteHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		delete := new(models.DeleteFavoriteRequest)
		if err := c.Bind(delete); err != nil {
			return err
		}

		accessToken := c.Request().Header.Get("Access-Token")
		if accessToken == "" {
			return c.JSON(http.StatusOK, models.BaseResponse{Code: "", ErrorMessage: ""})
		}

		user, err := models.GetUser(accessToken)
		if err != nil {
			return c.JSON(http.StatusOK, models.GetFavoriteResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: nil})
		}

		response := models.DeleteFavorite(user.ID, *delete)
		return c.JSON(http.StatusOK, response)
	}
}
