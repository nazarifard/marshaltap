package msgpackvmihailenco

import (
	"encoding/json"

	"github.com/nazarifard/syncpool"
)

type VmihailencoMsgpackSerializer[V any] struct {
	bufferPool syncpool.BufferPool
}

func (m *VmihailencoMsgpackSerializer[V]) Encode(v V) (syncpool.Buffer, error) {
	zb := m.bufferPool.Get(0)
	zb.Reset()
	err := json.NewEncoder(zb).Encode(v)
	if err != nil {
		zb.Free()
	}
	return zb, err
}

func (m *VmihailencoMsgpackSerializer[V]) Decode(buf syncpool.Buffer) (v V, err error) {
	err = json.NewDecoder((*syncpool.RBuffer)(buf)).Decode(&v)
	return
}
