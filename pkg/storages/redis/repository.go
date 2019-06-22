package redis

import (
	"tp-highload-performance-test/pkg"

	_ "github.com/go-redis/redis"
)

type Repository struct{}

func NewRepository() *Repository {
	panic("TODO")
}

func (r *Repository) OpenConnection() error {
	panic("TODO")
}

func (r *Repository) CloseConnection() error {
	panic("TODO")
}

func (r *Repository) SaveBlock(block *pkg.Block) error {
	panic("TODO")
}

func (r *Repository) LoadBlock(block *pkg.Block) error {
	panic("TODO")
}

func (r *Repository) DeleteBlock(block *pkg.Block) error {
	panic("TODO")
}
