package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/sjmh/testapi/items"
)

type MemoryItemRepository struct {
	memoryItems []items.Item
	mut         sync.RWMutex
}

func (m *MemoryItemRepository) AddItems(ctx context.Context, items []items.Item) error {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.memoryItems = append(m.memoryItems, items...)
	return nil
}

func (m *MemoryItemRepository) AddItem(ctx context.Context, item items.Item) error {
	return m.AddItems(ctx, []items.Item{item})
}

func (m *MemoryItemRepository) GetItemByID(ctx context.Context, id string) (*items.Item, error) {
	m.mut.Lock()
	defer m.mut.Unlock()
	for _, item := range m.memoryItems {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("item ID %v not found", id)
}

func (m *MemoryItemRepository) GetItems(ctx context.Context) []items.Item {
	return m.memoryItems
}

func NewMemoryItemRepository() *MemoryItemRepository {
	return &MemoryItemRepository{
		memoryItems: []items.Item{
			{ID: "1", Name: "Rusty Dagger"},
			{ID: "2", Name: "Old Spear"},
			{ID: "3", Name: "Holy Grail"},
		},
	}
}
