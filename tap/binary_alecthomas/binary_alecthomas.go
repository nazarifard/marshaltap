package binaryalecthomas

import (
	"bytes"

	"github.com/alecthomas/binary"
	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
)

type BinaryTap[V any] struct {
	bufferPool pool.BufferPool
}

func (m BinaryTap[V]) Encode(v V) (pool.Buffer, error) {
	zb := m.bufferPool.Get(0)
	zb.Reset()
	err := binary.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m BinaryTap[V]) Decode(bs []byte) (v V, n int, err error) {
	b2 := bytes.NewBuffer(bs)
	err = binary.NewDecoder(b2).Decode(&v)
	return v, len(bs) - b2.Len(), err
}

func NewTap[V any]() tap.Interface[V] {
	return BinaryTap[V]{
		bufferPool: pool.NewBufferPool(),
	}
}
