package repository

import (
	"sync"

	"example.com/product-api/model"
)

type HistoryRepository struct {
	data []model.SearchHistory
	mu   sync.Mutex
}

func NewHistoryRepository() *HistoryRepository {
	return &HistoryRepository{
		data: []model.SearchHistory{},
	}
}

func (r *HistoryRepository) Save(history model.SearchHistory) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data = append(r.data, history)
}

func (r *HistoryRepository) GetLatest(limit int) []model.SearchHistory {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.data) <= limit {
		return r.data
	}

	return r.data[len(r.data)-limit:]
}
func (r *HistoryRepository) GetByID(id string) (*model.SearchHistory, bool) {

	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.data {
		if r.data[i].ID == id {
			return &r.data[i], true
		}
	}

	return nil, false
}
