package models

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

type ID int64
type UUID uuid.UUID

func RandomID() ID {
	return ID(rand.Int63())
}

func (v ID) Int64() int64 {
	return int64(v)
}

func RandomUUID() UUID {
	v, _ := uuid.NewRandom()
	return UUID(v)
}

func (v UUID) String() string {
	return uuid.UUID(v).String()
}

const (
	IDLen   = 4
	UUIDLen = 16
)

type BlockData []byte

type Block struct {
	DocumentID UUID
	BlockID    ID
	Data       BlockData
}

func (v *Block) String() string {
	return fmt.Sprintf(
		"[document_id=%s; block_id=%d; data_size=%d]",
		v.DocumentID.String(), v.BlockID.Int64(), len(v.Data),
	)
}

func (v BlockData) MarshalBinary() ([]byte, error) {
	return []byte(v), nil
}

func (v *BlockData) UnmarshalBinary(data []byte) error {
	*v = data
	return nil
}
