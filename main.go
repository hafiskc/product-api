package main

import (
	"github.com/gin-gonic/gin"

	"example.com/product-api/handler"
	"example.com/product-api/provider"
	"example.com/product-api/repository"
	"example.com/product-api/service"
)

func main() {

    // initialize providers
    providers := []provider.Provider{
        &provider.ProviderA{},
        &provider.ProviderB{},
        &provider.ProviderC{},
    }

    // create service
	repo := repository.NewHistoryRepository()
    productService := service.NewProductService(providers,repo)

    // create handler
    productHandler := handler.NewProductHandler(productService)
	historyHandler := handler.NewHistoryHandler(repo)
    healthHandler := handler.NewHealthHandler()

    // setup gin
    router := gin.Default()

    // route
    router.GET("/search", productHandler.Search)
	router.GET("/history", historyHandler.GetHistory)
	router.GET("/history/:id", historyHandler.GetHistoryByID)
    router.GET("/health", healthHandler.Check)

    // run server
    router.Run(":8089")
}