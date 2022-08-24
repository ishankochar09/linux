package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ishankochar09/go_pro/Product-variant/internal/models"
)

type productRepo struct {
}

func NewProductRepo() ProductStore {
	return &productRepo{}
}

func (p *productRepo) AddProduct(ctx *gofr.Context, product *models.Product) (*models.Product, error) {
	return &models.Product{}, nil
}
