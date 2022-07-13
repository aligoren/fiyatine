package services

import "github.com/aligoren/fiyatine/internal/models"

type ProductService interface {
	// searchProduct amazon, hepsiburada ya da diğer providerlar hiç fark etmiyor. hepsi bu methodu implement etmek zorundalar.
	searchProduct() []models.ResponseModel
}
