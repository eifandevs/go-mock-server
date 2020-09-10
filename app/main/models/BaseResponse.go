package models

type BaseResponse struct {
    Code string `json:"code"`
    ErrorMessage string `json:"errorMessage"`
}