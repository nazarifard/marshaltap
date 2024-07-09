package marshal

import (
	"bytes"
	"encoding/gob"

	"github.com/nazarifard/syncpool"
)

type gobTap[V any] tap[V]

func (m gobTap[V]) Encode(v V) (syncpool.Buffer, error) {
	zb := m.bufferPool.Get(0)
	zb.Reset()
	err := gob.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m gobTap[V]) Decode(bs []byte) (v *V, n int, err error) {
	buf := bytes.NewBuffer(bs)
	err = gob.NewDecoder(buf).Decode(v)
	return v, len(bs) - buf.Len(), err
}

func (m gobTap[V]) Free(v *V) {
	m.pool.Put(v)
}

func NewGobTap[V any](pool *syncpool.Pool[V]) gobTap[V] {
	return gobTap[V]{
		bufferPool: syncpool.NewBufferPool(),
		pool:       pool,
		modem:      nil,
	}
}
