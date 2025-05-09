// Package blockbuffer is a simple heap like structure for keeping a
// few recent blocks in memory for averaging and analysis.
package blockbuffer

import (
	"container/heap"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/0xPolygon/panoptichain/log"
)

// BufferedBlock abstracts the block for queueing and sorting. The only thing
// needed is a block number in order to keep the blocks ordered.
type BufferedBlock interface {
	Number() *big.Int
}

// BlockBuffer represents a min-heap data structure to keep a set of recent
// blocks in memory.
type BlockBuffer struct {
	size    uint                     // the number of blocks to keep in the queue
	blocks  map[uint64]BufferedBlock // the actual blocks stored in a map
	numbers BigIntHeap               // A heap for keeping track of the blocks
	rw      sync.RWMutex
}

// BigIntHeap implements the heap interface for big.Ints.
type BigIntHeap []*big.Int

// Len returns the length of the heap.
func (h BigIntHeap) Len() int { return len(h) }

// Less compares to big.Ints.
func (h BigIntHeap) Less(i, j int) bool { return h[i].Cmp(h[j]) < 0 }

// Swap exchanges two elements in the heap.
func (h BigIntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *BigIntHeap) Push(x any) {
	*h = append(*h, x.(*big.Int))
}

// Pop removes the last element in the heap and returns it.
func (h *BigIntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// NewBlockBuffer returns a heap that's constrained to the given size.
func NewBlockBuffer(size uint) *BlockBuffer {
	return &BlockBuffer{
		blocks:  make(map[uint64]BufferedBlock, 0),
		numbers: make(BigIntHeap, 0),
		size:    size,
	}
}

// GetBlock returns a block or an error if it doesn't exist.
func (b *BlockBuffer) GetBlock(number uint64) (BufferedBlock, error) {
	b.rw.RLock()
	defer b.rw.RUnlock()

	block, ok := b.blocks[number]
	if !ok {
		return nil, fmt.Errorf("failed to fetch buffered block %d", number)
	}

	return block, nil
}

// PutBlock pushes a new block into the heap. If adding that block expand the
// heap beyond the max size, delete the oldest block (based on number).
func (b *BlockBuffer) PutBlock(block BufferedBlock) error {
	b.rw.Lock()
	defer b.rw.Unlock()

	n := block.Number()
	if n == nil {
		return errors.New("failed to get buffered block number")
	}

	// Add block to map
	b.blocks[n.Uint64()] = block

	// Add block number to min-heap
	heap.Push(&b.numbers, n)

	// If size exceeds n, remove the smallest block number
	if len(b.blocks) > int(b.size) {
		minBlockNumber := heap.Pop(&b.numbers).(*big.Int)

		log.Debug().
			Uint64("block_number", minBlockNumber.Uint64()).
			Msg("Block buffer is full, dropping earliest block")

		delete(b.blocks, minBlockNumber.Uint64())
	}

	return nil
}
