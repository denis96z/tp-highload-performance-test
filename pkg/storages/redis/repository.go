package redis

import (
	"tp-highload-performance-test/pkg/models"
	"tp-highload-performance-test/pkg/utils"

	"github.com/go-redis/redis"
)

const (
	bufferLen = models.IDLen + models.UUIDLen
)

const (
	expirationTime = 0
)

type Repository struct {
	client *redis.Client
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) OpenConnection() error {
	r.client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	return nil
}

func (r *Repository) CloseConnection() error {
	return r.client.Close()
}

func (r *Repository) SaveBlock(block *models.Block) error {
	key := makeKey(block)

	cmd := r.client.Set(
		key,
		block.Data,
		expirationTime,
	)

	return cmd.Err()
}

func (r *Repository) LoadBlock(block *models.Block) error {
	key := makeKey(block)
	cmd := r.client.Get(key)

	var err error
	block.Data, err = cmd.Bytes()

	return err
}

func (r *Repository) DeleteBlock(block *models.Block) error {
	key := makeKey(block)
	cmd := r.client.Del(key)
	return cmd.Err()
}

func makeKey(block *models.Block) string {
	keyBuf := make([]byte, bufferLen)
	utils.PrintKeyToBuffer(keyBuf, block.DocumentID, block.BlockID)
	return utils.BytesToString(keyBuf)
}
