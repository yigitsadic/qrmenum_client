package store

import (
	"github.com/yigitsadic/qrmenum_client/client"
	"github.com/yigitsadic/qrmenum_client/shared"
	"time"
)

type ProductMapItem struct {
	Products   []client.ProductResponse
	LastAccess time.Time
}

// TODO: Implement is expired control.
func (p ProductMapItem) IsExpired() bool {
	return false
}

type InMemoryStore struct {
	Client     *client.Client
	ProductMap map[string]ProductMapItem
}

// Sets a map item.
func (i *InMemoryStore) SetMapItem(key string, products []client.ProductResponse) {
	i.ProductMap[key] = ProductMapItem{
		Products:   products,
		LastAccess: time.Now(),
	}
}

// Returns pointer of http client.
func (i *InMemoryStore) GetClient() client.Client {
	return *i.Client
}

// Tries to return pointer to ProductMapItem if it is present in ProductMap map.
func (i *InMemoryStore) GetMapItem(key string) (*ProductMapItem, error) {
	p, ok := i.ProductMap[key]
	if !ok {
		return nil, shared.KeyNotFoundOnInMemoryStore
	}

	return &p, nil
}

// Initializes a new in memory store.
func NewInMemoryStore(client client.Client) *InMemoryStore {
	return &InMemoryStore{Client: &client}
}
