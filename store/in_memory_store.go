package store

import (
	"github.com/yigitsadic/qrmenum_client/client"
)

type InMemoryStore struct {
	Client client.Client
}

func (i *InMemoryStore) Get(key string) error {
	return nil
}

func NewInMemoryStore(client client.Client) *InMemoryStore {
	return &InMemoryStore{Client: client}
}
