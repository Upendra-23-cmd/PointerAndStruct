package main

import (
	"fmt"
	"sync"
)

// Memory Pool Stores Reuseable byte slices
type MemoryPool struct {
	BlockSize int
	Mu        sync.Mutex
	Free      [][]byte
}

// NewMemoryPool creates a pool with preallocated blocks
func NewMemoryPool(blockSize, initial int) *MemoryPool {
	p := &MemoryPool{BlockSize: blockSize}
	for i := 0; i < initial; i++ {
		p.Free = append(p.Free, make([]byte, blockSize))
	}
	return p
}

// Get retrieves a buffer from the pool or allocate a new one
func (p *MemoryPool) Get() []byte {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	if len(p.Free) == 0 {
		return make([]byte, p.BlockSize)
	}

	// pop from slice
	buf := p.Free[len(p.Free)-1]
	p.Free = p.Free[:len(p.Free)-1]
	return buf
}

// Put returns a buffer back to the pool
func (p *MemoryPool) Put(buf []byte) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Free = append(p.Free, buf)
}

func main() {
	pool := NewMemoryPool(8, 2)

	// Get Buffers
	b1 := pool.Get()
	b2 := pool.Get()

	copy(b1, []byte("Golang"))
	copy(b2, []byte("Memory"))

	fmt.Println("Buffer1 :", string(b1))
	fmt.Println("Buffer2 :", string(b2))

	// Return them
	pool.Put(b1)
	pool.Put(b2)

	fmt.Println("Pool has Buffer available again :", len(pool.Free))
}
