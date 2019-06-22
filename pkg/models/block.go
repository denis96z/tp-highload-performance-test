package models

type ID int64
type UUID string

const (
	BlockSize = 4
)

type BlockData [BlockSize]byte

type Block struct {
	DocumentID UUID
	BlockID    ID
	Data       BlockData
}
