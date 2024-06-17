package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
)

var bufferPool = pool.NewBufferPool()

type GobTap[V any] struct {
	//	bufferPool pool.BufferPool
}

func (m GobTap[V]) Encode(v V) (pool.Buffer, error) {
	zb := bufferPool.Get(0)
	zb.Reset()
	err := gob.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m GobTap[V]) Decode(bs []byte) (v V, n int, err error) {
	buf := bytes.NewBuffer(bs)
	err = gob.NewDecoder(buf).Decode(&v)
	return v, len(bs) - buf.Len(), err
}

func NewGobTap[V any]() tap.Interface[V] {
	return GobTap[V]{
		//	bufferPool: pool.NewBufferPool(),
	}
}
