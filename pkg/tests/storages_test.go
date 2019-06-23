package tests

import (
	"testing"
)

const (
	NBlocks      = 100
	NConnections = 1
)

func BenchmarkConnectionsRead(b *testing.B) {
	// setup - набор данных
	// для каждого из NConnections открыть соеднинение и записать NBlocks блоков

	// for i = 0; i < NConnections; i++ {

	// }
}

func BenchmarkSetup(b *testing.B) {
	Setup(NBlocks, "1")
}
