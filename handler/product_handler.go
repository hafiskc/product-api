package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "example.com/product-api/service"
)

type ProductHandler struct {
    service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
    return &ProductHandler{service: s}
}

func (h *ProductHandler) Search(c *gin.Context) {

    query := c.Query("q")

    if query == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "query parameter 'q' is required",
        })
        return
    }

    results := h.service.Search(query)

    c.JSON(http.StatusOK, gin.H{
        "data": results,
    })
}