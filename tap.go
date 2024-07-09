package marshal

import (
	"github.com/nazarifard/syncpool"
)

type tap[V any] struct {
	modem      Modem[V]
	pool       *syncpool.Pool[V]
	bufferPool syncpool.BufferPool
}

func (t tap[V]) Encode(v V) (buf syncpool.Buffer, err error) {
	size := t.modem.Sizeof(v)
	buf = t.bufferPool.Get(size)
	err = t.modem.Marshal(v, buf.Bytes())
	return
}

func (t tap[V]) Decode(bs []byte) (v *V, n int, err error) {
	v = t.pool.Get()
	err = t.modem.Unmarshal(bs, v)
	return v, t.modem.Sizeof(*v), err
}

func (t tap[V]) Free(v *V) {
	t.pool.Put(v)
}

func NewTap[V any](m Modem[V], pool *syncpool.Pool[V]) tap[V] {
	return tap[V]{
		pool:       pool,
		modem:      m,
		bufferPool: syncpool.NewBufferPool(),
	}
}
