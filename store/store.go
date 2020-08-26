package store

import "github.com/yigitsadic/qrmenum_client/client"

type MapItem interface {
	IsExpired() bool
}

type Store interface {
	GetMapItem(key string) (*ProductMapItem, error)
	SetMapItem(key string, products []client.ProductResponse)
	GetClient() client.Client
}
