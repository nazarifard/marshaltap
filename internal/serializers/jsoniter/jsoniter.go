package jsoniter

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	syncpool "github.com/nazarifard/syncpool"
)

var (
	jsoniterFast = jsoniter.ConfigFastest
)

type JsonIterSerializer[V any] struct {
	bufferPool syncpool.BufferPool
}

func (m *JsonIterSerializer[V]) Encode(v V) (syncpool.Buffer, error) {
	zb := m.bufferPool.Get(0)
	zb.Reset()
	err := jsoniter.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m *JsonIterSerializer[V]) Decode(buf syncpool.Buffer) (v V, err error) {
	err = jsoniter.NewDecoder((*syncpool.RBuffer)(buf)).Decode(&v)
	return
}

func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	return &JsonIterSerializer[goserbench.SmallStruct]{
		bufferPool: syncpool.NewBufferPool(),
	}
}
