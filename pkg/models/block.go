package models

type ID int64
type UUID string

type BlockData []byte

type Block struct {
	DocumentID UUID
	BlockID    ID
	Data       BlockData
}
