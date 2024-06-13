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

func (m *BinaryTap[V]) Encode(v V) (pool.Buffer, error) {
	zb := m.bufferPool.Get(0)
	zb.Reset()
	err := binary.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m *BinaryTap[V]) Decode(buf pool.Buffer) (v V, err error) {
	b2 := bytes.NewBuffer(buf.Bytes())
	err = binary.NewDecoder(b2).Decode(&v)
	return
}

func NewTap[V any]() tap.TapInterface[V] {
	return &BinaryTap[V]{
		bufferPool: pool.NewBufferPool(),
	}
}
