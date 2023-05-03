package main

import (
	"updateRepoName/internal/handler/product"
	"updateRepoName/internal/handler/variant"

	productSVC "updateRepoName/internal/service/product"
	variantSVC "updateRepoName/internal/service/variant"
	productSTR "updateRepoName/internal/store/product"
	variantSTR "updateRepoName/internal/store/variant"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {

	app := gofr.New()
	app.Server.ValidateHeaders = false

	productStore := productSTR.NewProductRepo()
	variantStore := variantSTR.NewVariantRepo()

	productService := productSVC.NewProductService(&productStore)
	variantService := variantSVC.NewVariantService(&variantStore)

	productHandler := product.NewProductHandler(productService)
	variantHandler := variant.NewVariantHandler(variantService)

	app.POST("/product", productHandler.AddProduct)
	app.GET("/product/{pid}", productHandler.GetProduct)

	app.POST("/product/{pid}/variant", variantHandler.AddVariant)
	app.GET("/product/{pid}/variant/{vid}", variantHandler.GetVariant)

	app.Start()
}
