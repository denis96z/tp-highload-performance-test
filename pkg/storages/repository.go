package storages

type Repository interface {
	OpenConnection() error
	CloseConnection() error

	SaveBlock(block *Block) error
	LoadBlock(block *Block) error
	DeleteBlock(block *Block) error
}
