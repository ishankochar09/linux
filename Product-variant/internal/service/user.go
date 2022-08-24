package service

import (
	"github.com/ishankochar09/go_pro/Product-variant/internal/store"
	"github.com/ishankochar09/go_pro/Product-variant/internal/models"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type product struct {
	ProductSVC store.ProductStore
}

func NewProductService(store store.ProductStore)ProductService{
	return &product{store}
}

func (p *product)AddProduct(ctx *gofr.Context, product *models.Product) (*models.Product, error){
return &models.Product{}, nil
}


