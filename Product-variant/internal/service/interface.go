package service

import (
	"github.com/ishankochar09/go_pro/Product-variant/internal/models"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	
)

type ProductService interface {
	AddProduct(ctx *gofr.Context, product *models.Product) (*models.Product, error)
}
