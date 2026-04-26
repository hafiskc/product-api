package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "example.com/product-api/repository"
)

type HistoryHandler struct {
    repo *repository.HistoryRepository
}

func NewHistoryHandler(r *repository.HistoryRepository) *HistoryHandler {
    return &HistoryHandler{repo: r}
}

func (h *HistoryHandler) GetHistory(c *gin.Context) {

    history := h.repo.GetLatest(10)

    c.JSON(http.StatusOK, gin.H{
        "data": history,
    })
}
func (h *HistoryHandler) GetHistoryByID(c *gin.Context) {

    id := c.Param("id")

    history, found := h.repo.GetByID(id)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "history not found",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": history,
    })
}