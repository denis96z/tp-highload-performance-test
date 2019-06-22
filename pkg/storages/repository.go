package storages

import (
	"tp-highload-performance-test/pkg/models"
)

type Repository interface {
	OpenConnection() error
	CloseConnection() error

	SaveBlock(block *models.Block) error
	LoadBlock(block *models.Block) error
	DeleteBlock(block *models.Block) error
}
