package tap

import (
	. "github.com/nazarifard/marshaltap/modem"
	pool "github.com/nazarifard/syncpool"
)

type Encoder[V any] interface {
	Encode(V) (pool.Buffer, error)
}
type Decoder[V any] interface {
	Decode(bs []byte) (V, error)
}
type TapInterface[V any] interface {
	Encoder[V]
	Decoder[V]
}

var tapBufferPool = pool.NewBufferPool()

type Tap[V any, M ModemInterface[V]] struct {
	Modem ModemInterface[V]
}

func (t *Tap[V, M]) Encode(v V) (buf pool.Buffer, err error) {
	size := t.Modem.Sizeof(v)
	buf = tapBufferPool.Get(size)
	err = t.Modem.Marshal(v, buf.Bytes())
	return
}
func (t *Tap[V, M]) Decode(bs []byte) (v V, err error) {
	err = t.Modem.Unmarshal(bs, &v)
	return
}

func NewTap[V any, M ModemInterface[V]](modem ModemInterface[V]) *Tap[V, M] {
	return &Tap[V, M]{
		Modem: modem,
	}
}
