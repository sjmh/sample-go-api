package main

import (
	"context"
	"fmt"
	"github.com/sjmh/testapi/items"
	"sync"
)

type MemoryItemRepository struct {
	memoryItems []items.Item
	mut sync.RWMutex
}

func (m *MemoryItemRepository) AddItem(ctx context.Context, item items.Item) bool {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.memoryItems = append(m.memoryItems, item)
	return true
}

func (m MemoryItemRepository) GetItemByID(ctx context.Context, id string) (items.Item, error) {
	for _, item := range m.memoryItems {
		if item.ID == id {
			return item, nil
		}
	}
	error := fmt.Errorf("item ID %v not found", id)
	return items.Item{}, error
}

func (m MemoryItemRepository) GetItems(ctx context.Context) []items.Item {
	return m.memoryItems
}

func NewMemoryItemRepository() *MemoryItemRepository {
	m := &MemoryItemRepository{}
	m.AddItem(context.Background(), items.Item{ID: "1", Name: "Rusty Dagger"})
	m.AddItem(context.Background(), items.Item{ID: "2", Name: "Old Spear"})
	m.AddItem(context.Background(), items.Item{ID: "3", Name: "Holy Grail"})
	return m
}
