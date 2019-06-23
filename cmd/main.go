package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"tp-highload-performance-test/pkg/storages/memcache"
	"tp-highload-performance-test/pkg/storages/redis"

	"tp-highload-performance-test/pkg/models"
	"tp-highload-performance-test/pkg/storages/leveldb"
)

func main() {
	numThreads := parseArgInt("num-threads", 1)
	blockSize := parseArgInt("block-size", 256)
	numUniqueDocs := parseArgInt("num-unique-docs", 1000)
	numUniqueBlocksPerDoc := parseArgInt("num-unique-blocks-per-doc", 10000)
	sizePerThread := parseArgMemGb("size-per-thread", 1)
	storage := parseArgStr("storage", "leveldb")
	address := parseArgStr("address", "./documents")

	fmt.Printf("Threads: %d\n", numThreads)
	fmt.Printf("Block size: %db\n", blockSize)
	fmt.Printf("Size per thread: %dGb\n", sizePerThread)
	fmt.Printf("Storage address: %s\n", address)

	c := leveldb.NewRepository
	switch storage {
	case "leveldb":
		c = leveldb.NewRepository
	case "redis":
		c = redis.NewRepository
	case "memcache":
		c = memcache.NewRepository
	}

	dIDs := makeDocIDs(numUniqueDocs)
	bIDs := makeBlockIDs(numUniqueBlocksPerDoc)

	blocks := make([][]models.Block, numThreads)
	for i := 0; i < numThreads; i++ {
		blocks[i] = makeBlocks(
			numUniqueDocs,
			blockSize,
			bIDs, dIDs,
		)
	}

	r := c(address)
	fmt.Println("Repository initialized...")

	if err := r.OpenConnection(); err != nil {
		panic(err)
	}
	fmt.Println("Connection opened...")

	fmt.Println("Press [ENTER] to continue...")
	_, _ = fmt.Scanln()

	defer func() {
		if err := r.CloseConnection(); err != nil {
			panic(err)
		}
		fmt.Println("Connection closed...")
	}()

	wg := &sync.WaitGroup{}

	sizePerThread *= 1024 * 1024 * 1024
	for i := 0; i < numThreads; i++ {
		wg.Add(1)

		go func(th int) {
			ops := []struct {
				Name string
				Func func(*models.Block) error
			}{
				{
					Name: "saving",
					Func: r.SaveBlock,
				},
				{
					Name: "loading",
					Func: r.LoadBlock,
				},
				{
					Name: "deleting",
					Func: r.DeleteBlock,
				},
			}

			for _, op := range ops {
				for i := 0; i < sizePerThread; i += blockSize {
					b := &blocks[th][rand.Intn(numUniqueDocs)]

					t := time.Now()
					err := op.Func(b)
					d := time.Now().Sub(t)

					status := "[ok]"
					if err != nil {
						status = fmt.Sprintf("[fail=%s]", err.Error())
					}

					fmt.Printf(
						"%s block: [status=%s, thread=%d; block=%s; time=%v]\n",
						op.Name, status, th, b.String(), d,
					)
				}
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func parseArgStr(name string, defVal string) string {
	for _, arg := range os.Args {
		prefix := fmt.Sprintf("--%s=", name)
		if strings.HasPrefix(arg, prefix) {
			return arg[strings.Index(arg, "=")+1:]
		}
	}
	return defVal
}

func parseArgInt(name string, defVal int) int {
	s := parseArgStr(name, strconv.Itoa(defVal))
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func parseArgMemGb(name string, defVal int) int {
	s := parseArgStr(name, strconv.Itoa(defVal)+"Gb")
	v, err := strconv.Atoi(s[:len(s)-2])
	if err != nil {
		panic(err)
	}
	return v
}

func makeDocIDs(n int) []models.UUID {
	ids := make([]models.UUID, n)
	for i := 0; i < n; i++ {
		ids[i] = models.RandomUUID()
	}
	return ids
}

func makeBlockIDs(n int) []models.ID {
	ids := make([]models.ID, n)
	for i := 0; i < n; i++ {
		ids[i] = models.RandomID()
	}
	return ids
}

func makeBlocks(
	n int, size int,
	bIDs []models.ID,
	dIDs []models.UUID,
) []models.Block {
	blocks := make([]models.Block, n)
	for i := 0; i < n; i++ {
		blocks[i] = models.Block{
			DocumentID: dIDs[rand.Intn(len(dIDs))],
			BlockID:    bIDs[rand.Intn(len(bIDs))],
			Data:       make([]byte, size),
		}
	}
	return blocks
}
