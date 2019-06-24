package memcache

import (
	"fmt"

	"tp-highload-performance-test/pkg/models"
	"tp-highload-performance-test/pkg/storages"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	expirationTime = 0
)

type Repository struct {
	address string

	client *memcache.Client
}

func NewRepository(address string) storages.Repository {
	return &Repository{
		address: address,
	}
}

func (r *Repository) OpenConnection() error {
	r.client = memcache.New(r.address)
	return nil
}

func (r *Repository) CloseConnection() error {
	return nil
}

func (r *Repository) SaveBlock(block *models.Block) error {
	key := makeKey(block)
	return r.client.Set(&memcache.Item{
		Key:        key,
		Value:      block.Data,
		Expiration: expirationTime,
	})
}

func (r *Repository) LoadBlock(block *models.Block) error {
	key := makeKey(block)

	item, err := r.client.Get(key)
	if err != nil {
		return err
	}

	block.Data = item.Value
	return nil
}

func (r *Repository) DeleteBlock(block *models.Block) error {
	key := makeKey(block)
	return r.client.Delete(key)
}

func makeKey(block *models.Block) string {
	return fmt.Sprintf(
		"%s:%d",
		block.DocumentID.String(),
		block.BlockID.Int64(),
	)
}
