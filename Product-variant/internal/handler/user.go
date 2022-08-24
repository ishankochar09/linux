package handler

import (
	// "developer.zopsmart.com/go/gofr/pkg/service"
	// "github.com/ishankochar09/go_pro/Product-variant/internal/models"
	"github.com/ishankochar09/go_pro/Product-variant/internal/models"
	"github.com/ishankochar09/go_pro/Product-variant/internal/service"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	// "github.com/vikash/gofr/pkg/gofr"
)

type Product struct {
	service service.ProductService
}

func NewProductHandler(han service.ProductService) *Product {
	return &Product{han}
}

func (p *Product) AddProduct(c *gofr.Context) (interface{}, error) {
	product := &models.Product{}
	err := c.Bind(product)
	if err != nil {
		return nil, err
	}

	return p.service.AddProduct(c, product)
}
