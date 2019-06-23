package tests

import (
	"testing"

	"tp-highload-performance-test/pkg/models"
	"tp-highload-performance-test/pkg/storages/leveldb"

	"github.com/google/uuid"
)

const (
	NBlocks      = 100
	NConnections = 1
	NBatch       = 256
)

func BenchmarkConnectionsWrite(b *testing.B) {
	id, err := uuid.Parse("00000000-0000-0000-0000-000000000001")
	if err != nil {
		panic(err)
	}
	blocks := generateBlocks(NBlocks, models.UUID(id), NBatch)
	b.ResetTimer()

	repo := leveldb.NewRepository()
	err = repo.OpenConnection()

	if err != nil {
		panic("Failed to connect")
	}
	defer repo.CloseConnection()

	for i := 0; i < NBlocks; i++ {
		err = repo.SaveBlock(blocks[i])
		if err != nil {
			panic("Failed to save block")
		}
	}
}

func BenchmarkGenerateBlocks(b *testing.B) {
	id := uuid.MustParse("00000000-0000-0000-0000-000000000001")
	generateBlocks(NBlocks, models.UUID(id), 256)
}
