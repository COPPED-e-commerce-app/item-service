package rest

import (
	"github.com/COPPED/item-service/internal/item-service/service"
	"github.com/gofiber/fiber/v2"
)

func FetchItem(c *fiber.Ctx) error {
	id := c.Params("id")
	item := service.FetchItem(id)
	return c.JSON(item)
}

func FetchItemList(c *fiber.Ctx) error {
	itemList := service.FetchItemList()
	return c.JSON(itemList)
}
