package marshal

import (
	"encoding/json"

	"github.com/nazarifard/syncpool"
)

type jsonTap[V any] tap[V]

func (m jsonTap[V]) Encode(v V) (zb syncpool.Buffer, err error) {
	bs, err := json.Marshal(v) //
	zb = m.bufferPool.Get(len(bs))
	zb.Reset()
	zb.Write(bs)
	return zb, err
}

func (m jsonTap[V]) Decode(bs []byte) (v *V, n int, err error) {
	v = m.pool.Get()
	err = json.Unmarshal(bs, v)
	if err != nil {
		return nil, 0, err
	}
	return v, len(bs), err
}

func (m jsonTap[V]) Free(v *V) {
	m.pool.Put(v)
}

func NewJsonTap[V any](pool *syncpool.Pool[V]) jsonTap[V] {
	return jsonTap[V]{
		bufferPool: syncpool.NewBufferPool(),
		pool:       pool,
		modem:      nil,
	}
}
