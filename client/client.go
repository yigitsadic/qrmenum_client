package client

type Client interface {
	FetchFromCMS(key string) ([]ProductResponse, error)
}
