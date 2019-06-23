package tests

import (
	"testing"
	"tp-highload-performance-test/pkg/storages/leveldb"
)

const (
	NBlocks      = 100
	NConnections = 1
	NBatch       = 256
)

func BenchmarkConnectionsWrite(b *testing.B) {
	// setup - набор данных
	blocks := setup(NBlocks, "1", NBatch)
	b.ResetTimer()

	// для каждого из NConnections открыть соеднинение и записать NBlocks блоков
	// for i = 0; i < NConnections; i++ {
	// switch to goroutines
	// }

	repo := leveldb.NewRepository()
	err := repo.OpenConnection()
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

	// write data
}

// func BenchmarkSetup(b *testing.B) {
// 	setup(NBlocks, "1", 256)
// }
