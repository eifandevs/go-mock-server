package controllers

import (
	"net/http"

	"github.com/eifandevs/amby/models"
	"github.com/labstack/echo"
)

func LoginHandler() echo.HandlerFunc {
  return func(c echo.Context) error {
    post := new(models.PostUserRequest)
		if err := c.Bind(post); err != nil {
      return c.JSON(http.StatusOK, models.BaseResponse{Code: "", ErrorMessage: ""})
    }
    
    // ユーザー情報の登録
    user, err := models.CreateUser((*post).Data)
    if err != nil {
      return c.JSON(http.StatusOK, models.BaseResponse{Code: "", ErrorMessage: ""})
    }

    return c.JSON(http.StatusOK, models.PostUserResponse{BaseResponse: models.BaseResponse{Code: "", ErrorMessage: ""}, Data: models.UserToken{AccessToken: user.AccessToken}})
  }
}
