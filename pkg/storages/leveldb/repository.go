package leveldb

import (
	"encoding/binary"
	"fmt"

	"tp-highload-performance-test/pkg/models"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type Repository struct {
	openReadOptions  opt.Options
	openWriteOptions opt.Options

	readOptions  opt.ReadOptions
	writeOptions opt.WriteOptions
}

func NewRepository() *Repository {
	return &Repository{
		openReadOptions: opt.Options{
			Strict:   opt.NoStrict,
			ReadOnly: true,
		},
		openWriteOptions: opt.Options{
			Strict:   opt.NoStrict,
			ReadOnly: false,
		},
		readOptions: opt.ReadOptions{
			Strict:        opt.NoStrict,
			DontFillCache: false,
		},
		writeOptions: opt.WriteOptions{
			Sync:         false,
			NoWriteMerge: true,
		},
	}
}

func (r *Repository) OpenConnection() error {
	return nil
}

func (r *Repository) CloseConnection() error {
	return nil
}

func (r *Repository) SaveBlock(block *models.Block) error {
	db, err := r.openDB(block.DocumentID, &r.openWriteOptions)
	if err != nil {
		return err
	}
	defer r.closeDB(db)

	return db.Put(
		binBlockID(block.BlockID),
		block.Data,
		&r.writeOptions,
	)
}

func (r *Repository) LoadBlock(block *models.Block) error {
	db, err := r.openDB(block.DocumentID, &r.openReadOptions)
	if err != nil {
		return err
	}
	defer r.closeDB(db)

	block.Data, err = db.Get(
		binBlockID(block.BlockID),
		&r.readOptions,
	)
	return err
}

func (r *Repository) DeleteBlock(block *models.Block) error {
	db, err := r.openDB(block.DocumentID, &r.openWriteOptions)
	if err != nil {
		return err
	}
	defer r.closeDB(db)

	return db.Delete(
		binBlockID(block.BlockID),
		&r.writeOptions,
	)
}

const (
	DocumentsPath = "./documents"
)

func (r *Repository) openDB(
	documentID models.UUID, options *opt.Options,
) (*leveldb.DB, error) {
	return leveldb.OpenFile(
		fmt.Sprintf(DocumentsPath+"/%s", documentID), options,
	)
}

func (r *Repository) closeDB(db *leveldb.DB) {
	_ = db.Close()
}

func binBlockID(key models.ID) []byte {
	d := make([]byte, 8)
	binary.LittleEndian.PutUint64(d, uint64(key))
	return d
}
