package utils

/*
#include <stdint.h>
#include <string.h>

#define UUID_SIZE 16

void
print_to_buffer(void *buffer,
                const void *document_id,
                int64_t block_id)
{
	memmove(buffer, document_id, UUID_SIZE);
	memmove(buffer + UUID_SIZE, &block_id, sizeof(block_id));
}
*/
import "C"

import (
	"unsafe"

	"tp-highload-performance-test/pkg/models"
)

func PrintKeyToBuffer(buffer []byte, documentID models.UUID, blockID models.ID) {
	C.print_to_buffer(
		unsafe.Pointer(&buffer[0]),
		unsafe.Pointer(&documentID),
		C.long(blockID),
	)
}
