package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
)

type GobTap[V any] struct {
	bufferPool pool.BufferPool
}

func (m *GobTap[V]) Encode(v V) (pool.Buffer, error) {
	zb := m.bufferPool.Get(0)
	zb.Reset()
	err := gob.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m *GobTap[V]) Decode(bs []byte) (v V, err error) {
	err = gob.NewDecoder(bytes.NewBuffer(bs)).Decode(&v)
	return
}

func NewGobTap[V any]() tap.TapInterface[V] {
	return &GobTap[V]{
		bufferPool: pool.NewBufferPool(),
	}
}
