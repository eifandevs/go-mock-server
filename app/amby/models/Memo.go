package models

import (
	"github.com/eifandevs/amby/repo"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
)

type MemoInfo struct {
	FID      int    `json:"fid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Memo struct {
    gorm.Model
    FID int `gorm:"type:int unsigned;not null;unique;primary_key;auto_increment:false"`
    UserID uint
    Title string
	Content string `gorm:"type:varchar(1000)"`
}

type GetMemoResponse struct {
    BaseResponse
    Data  []MemoInfo `json:"data"`
}

type PostMemoRequest struct {
    Data  []MemoInfo `json:"data"`
}

type DeleteMemoRequest struct {
    Data  []MemoInfo `json:"data"`
}

func GetMemo(userID uint) GetMemoResponse {
    db := repo.Connect("development")
    defer db.Close()

    memos := []Memo{}
    if err := db.Where("user_id = ?", userID).Find(&memos).Error; err != nil {
        return GetMemoResponse{BaseResponse: BaseResponse{Code: "", ErrorMessage: ""}, Data: nil}
    }

    items := funk.Map(memos, func(memo Memo) MemoInfo {
        return MemoInfo{FID: memo.FID, Title: memo.Title, Content: memo.Content}
    })
    
    if castedItems, ok := items.([]MemoInfo); ok {
        return GetMemoResponse{BaseResponse: BaseResponse{Code: "200", ErrorMessage: ""}, Data: castedItems}
    } else {
        panic("cannot cast memo item.")
    }
}

func PostMemo(userID uint, request PostMemoRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Data {
        if err := db.Create(&Memo{FID: item.FID, UserID: userID, Title: item.Title, Content: item.Content}).Error; err != nil {
            return BaseResponse{Code: "", ErrorMessage: ""}
        }
    }

	return BaseResponse{Code: "200", ErrorMessage: ""}
}

func DeleteMemo(userID uint, request DeleteMemoRequest) BaseResponse {
    db := repo.Connect("development")
    defer db.Close()

    for _, item := range request.Data {
        deletingRecord := Memo{}
        deletingRecord.FID = item.FID
        deletingRecord.UserID = userID
        db.First(&deletingRecord)
        if err := db.Unscoped().Delete(&deletingRecord).Error; err != nil {
            return BaseResponse{Code: "", ErrorMessage: ""}
        }
    }

	return BaseResponse{Code: "200", ErrorMessage: ""}
}
