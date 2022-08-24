package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ishankochar09/go_pro/Product-variant/internal/handler"
	"github.com/ishankochar09/go_pro/Product-variant/internal/service"
	"github.com/ishankochar09/go_pro/Product-variant/internal/store"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	productStore := store.NewProductRepo()
	productService := service.NewProductService(productStore)
	productHandler := handler.NewProductHandler(productService)

	app.POST("/product", productHandler.AddProduct)

	app.Start()

}
