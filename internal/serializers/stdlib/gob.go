package stdlib

import (
	"encoding/json"

	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
)

type GobTap[V any] struct {
	bufferPool pool.BufferPool
}

func (m *GobTap[V]) Encode(v V) (pool.Buffer, error) {
	zb := m.bufferPool.Get(0)
	zb.Reset()
	err := json.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m *GobTap[V]) Decode(buf pool.Buffer) (v V, err error) {
	err = json.NewDecoder((*pool.RBuffer)(buf)).Decode(&v)
	return
}

func NewGobTap[V any]() tap.TapInterface[V] {
	return &GobTap[V]{
		bufferPool: pool.NewBufferPool(),
	}
}
