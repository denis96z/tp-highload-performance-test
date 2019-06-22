package models

import (
	"github.com/google/uuid"
)

type ID int64
type UUID uuid.UUID

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
