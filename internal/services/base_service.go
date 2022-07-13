package services

import "github.com/aligoren/fiyatine/internal/models"

type BaseService struct {
	ProductService
}

func (baseService BaseService) Search() []models.ResponseModel {
	return baseService.ProductService.searchProduct()
}
