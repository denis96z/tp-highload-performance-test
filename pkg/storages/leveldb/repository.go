package leveldb

import (
	"tp-highload-performance-test/pkg/models"

	_ "github.com/syndtr/goleveldb/leveldb"
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

func (r *Repository) SaveBlock(block *models.Block) error {
	panic("TODO")
}

func (r *Repository) LoadBlock(block *models.Block) error {
	panic("TODO")
}

func (r *Repository) DeleteBlock(block *models.Block) error {
	panic("TODO")
}
