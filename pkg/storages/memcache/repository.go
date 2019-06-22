package memcache

import (
	"tp-highload-performance-test/pkg/models"
	"tp-highload-performance-test/pkg/utils"

	"github.com/bradfitz/gomemcache/memcache"
)

const (
	bufferLen = models.IDLen + models.UUIDLen
)

const (
	expirationTime = 0
)

type Repository struct {
	client *memcache.Client
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) OpenConnection() error {
	r.client = memcache.New("127.0.0.1:11211")
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
	keyBuf := make([]byte, bufferLen)
	utils.PrintKeyToBuffer(keyBuf, block.DocumentID, block.BlockID)
	return utils.BytesToString(keyBuf)
}
