package services

type BaseService struct {
	ProductService
}

func (baseService BaseService) Search() {
	baseService.ProductService.searchProduct()
}
