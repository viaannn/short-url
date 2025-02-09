package app

import "time"

type MasterUrl struct {
	Key         string     `gorm:"primaryKey", json:"key"`
	TargetUrl   string     `json:"target_url"`
	CreatedDate *time.Time `json:"created_date"`
}

func (MasterUrl) TableName() string {
	return "master_url"
}

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateShortLinkRequest struct {
	TargetUrl string `json:"target_url"`
}

type CreateShortLinkResponse struct {
	ShortLink string `json:"short_link"`
	TargetUrl string `json:"targer_url"`
}
