package service

import (
	"github.com/COPPED/item-service/pkg/models"
)

func FetchItem(id string) *models.Item {
	return &models.Item{
		ItemId:      "item-id-1",
		ItemName:    "item-name-1",
		Brand:       "brand-1",
		ImgUrlList:  []string{"img-url-list-1"},
		Price:       "$19.99",
		IsAvailable: true,
		Description: "Oversized t-shirt",
		Sizing:      "Model is 6ft 6in wearing an XL",
	}
}

func FetchItemList() *[]models.Item {
	return &[]models.Item{
		{
			ItemId:      "item-id-1",
			ItemName:    "item-name-1",
			Brand:       "brand-1",
			ImgUrlList:  []string{"img-url-list-1"},
			Price:       "$19.99",
			IsAvailable: true,
			Description: "Oversized t-shirt",
			Sizing:      "Model is 6ft 6in wearing an XL",
		},
		{
			ItemId:      "item-id-2",
			ItemName:    "item-name-2",
			Brand:       "brand-2",
			ImgUrlList:  []string{"img-url-list-2"},
			Price:       "$129.99",
			IsAvailable: true,
			Description: "Oversized sweatshirt",
			Sizing:      "Model is 6ft 6in wearing an XL",
		},
	}
}
