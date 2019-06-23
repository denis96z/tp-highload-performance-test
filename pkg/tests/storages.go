package tests

import (
	"tp-highload-performance-test/pkg/models"
)

func generateBlocks(nBlocks models.ID, documentID models.UUID, blockSize int) []*models.Block {
	var blocks []*models.Block
	data := make([]byte, blockSize)
	for j := 0; j < blockSize; j++ {
		data[j] = byte(j)
	}
	var i models.ID
	for i = 0; i < nBlocks; i++ {
		blocks = append(blocks, &models.Block{
			DocumentID: documentID,
			BlockID:    i,
			Data:       data,
		})
	}
	return blocks
}
