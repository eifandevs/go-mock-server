package models

import (
	"github.com/eifandevs/amby/repo"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
)

type FavoriteInfo struct {
    FID int `json:"fid"`
    Title string `json:"title"`
    Url string `json:"url"`
}

type Favorite struct {
    gorm.Model
    FID int `gorm:"type:int unsigned;not null;unique;primary_key;auto_increment:false"`
    UserID uint
    Title string
    Url string
}

type GetFavoriteResponse struct {
    BaseResponse
    Data  []FavoriteInfo `json:"data"`
}

type PostFavoriteRequest struct {
    Data  []FavoriteInfo `json:"data"`
}

type DeleteFavoriteRequest struct {
    Data  []FavoriteInfo `json:"data"`
}

func GetFavorite(userID uint) GetFavoriteResponse {
    db := repo.Connect("development")
    defer db.Close()

    favorites := []Favorite{}
    if err := db.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
        return GetFavoriteResponse{BaseResponse: BaseResponse{Code: "", ErrorMessage: ""}, Data: nil}
    }

    items := funk.Map(favorites, func(favorite Favorite) FavoriteInfo {
        return FavoriteInfo{FID: favorite.FID, Title: favorite.Title, Url: favorite.Url}
    })
    
    if castedItems, ok := items.([]FavoriteInfo); ok {
        return GetFavoriteResponse{BaseResponse: BaseResponse{Code: "200", ErrorMessage: ""}, Data: castedItems}
    } else {
        panic("cannot cast favorite item.")
    }
}

func PostFavorite(userID uint, request PostFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Data {
        if err := db.Create(&Favorite{FID: item.FID, UserID: userID, Title: item.Title, Url: item.Url}).Error; err != nil {
            return BaseResponse{Code: "", ErrorMessage: ""}
        }
    }

	return BaseResponse{Code: "200", ErrorMessage: ""}
}

func DeleteFavorite(userID uint, request DeleteFavoriteRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Data {
        deletingRecord := Favorite{}
        deletingRecord.FID = item.FID
        deletingRecord.UserID = userID
        db.First(&deletingRecord)
        if err := db.Unscoped().Delete(&deletingRecord).Error; err != nil {
            return BaseResponse{Code: "", ErrorMessage: ""}
        }
    }

	return BaseResponse{Code: "200", ErrorMessage: ""}
}