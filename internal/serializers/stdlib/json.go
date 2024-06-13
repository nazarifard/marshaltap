package stdlib

import (
	"bytes"
	"encoding/json"

	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
	pool "github.com/nazarifard/syncpool"
)

type JsonTap[V any] struct {
	bufferPool pool.BufferPool
	w          *bytes.Buffer
	r          *bytes.Buffer
	enc        *json.Encoder
	dec        *json.Decoder
}

func (m *JsonTap[V]) Encode(v V) (zb pool.Buffer, err error) {
	m.w.Reset()
	err = m.enc.Encode(v)
	if err != nil {
		return nil, err
	}
	zb = m.bufferPool.Get(m.w.Len())
	zb.Reset()
	_, err = zb.Write(m.w.Bytes())
	return zb, err
}

func (m *JsonTap[V]) Decode(buf pool.Buffer) (v V, err error) {
	err = json.NewDecoder((*pool.RBuffer)(buf)).Decode(&v)
	return
}

func NewJsonTap[V any]() tap.TapInterface[V] {
	wbuf := make([]byte, 1024)
	rbuf := make([]byte, 1024)
	w := bytes.NewBuffer(wbuf)
	r := bytes.NewBuffer(rbuf)
	return &JsonTap[V]{
		bufferPool: syncpool.NewBufferPool(),
		r:          r,
		w:          w,
		enc:        json.NewEncoder(w),
		dec:        json.NewDecoder(r),
	}
}
