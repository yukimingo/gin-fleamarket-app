package controllers

import (
	"gin-fleamarket/models"
	"gin-fleamarket/services"
)

type IItemController interface {
	FindAll() (*[]models.Item, error)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (c ItemController) FindAll() (*[]models.Item, error) {
	return c.service.FindAll()
}

jfdks;ljfdklsa
