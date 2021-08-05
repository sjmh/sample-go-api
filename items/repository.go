package items

import "context"

type Repository interface {
	AddItems(ctx context.Context, items []Item) error
	AddItem(ctx context.Context, item Item) error
	GetItemByID(ctx context.Context, id string) (*Item, error)
	GetItems(ctx context.Context) []Item
}
