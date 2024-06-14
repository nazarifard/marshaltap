package json

import (
	"encoding/json"

	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
	pool "github.com/nazarifard/syncpool"
)

type JsonTap[V any] struct {
	bufferPool pool.BufferPool
}

// Json Encode does not improve performance. its done just for compatibilty
func (m JsonTap[V]) Encode(v V) (zb pool.Buffer, err error) {
	bs, err := json.Marshal(v) //
	zb = m.bufferPool.Get(len(bs))
	zb.Reset()
	zb.Write(bs)
	return zb, err
}

func (m JsonTap[V]) Decode(bs []byte) (v V, n int, err error) {
	err = json.Unmarshal(bs, &v)
	return v, 0, err //TODO
}

func NewJsonTap[V any]() tap.Interface[V] {
	return JsonTap[V]{
		bufferPool: syncpool.NewBufferPool(),
	}
}
