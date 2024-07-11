package usecase

import "go-api-rest/model"

type ProductUsecase struct {
	//repositori
}

func NewProductUsecase() ProductUsecase {
	return ProductUsecase{}
}

func (p *ProductUsecase) GetProducts() ([]model.Product, error){
	return []model.Product{}, nil
}