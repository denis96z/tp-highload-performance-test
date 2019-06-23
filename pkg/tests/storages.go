package tests

import (
	"tp-highload-performance-test/pkg/models"
)

// Setup -
func Setup(NBlocks models.ID, DocumentID models.UUID) []*models.Block {
	var blocks []*models.Block
	var i models.ID
	for i = 0; i < NBlocks; i++ {
		blocks = append(blocks, &models.Block{
			DocumentID: DocumentID,
			BlockID:    NBlocks,
			Data:       []byte{'1', '2', '3', '4'},
		})
	}
	return blocks
}
