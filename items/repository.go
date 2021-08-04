package items

import "context"

type Repository interface {
	AddItem(ctx context.Context, item Item) bool
	GetItemByID(ctx context.Context, id string) (Item, error)
	GetItems(ctx context.Context) []Item
}