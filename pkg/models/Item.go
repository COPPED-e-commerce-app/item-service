package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemId      string   `json:"item_id"`
	ItemName    string   `json:"item_name"`
	Brand       string   `json:"brand"`
	ImgUrlList  []string `json:"img_url_list"`
	Price       string   `json:"price"`
	IsAvailable bool     `json:"is_available"`
	Description string   `json:"description"`
	Sizing      string   `json:"sizing"`
}
