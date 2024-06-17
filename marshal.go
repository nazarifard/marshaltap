package marshal

import (
	pool "github.com/nazarifard/syncpool"
)

type Encoder[V any] interface {
	Encode(V) (pool.Buffer, error)
}

type Decoder[V any] interface {
	Decode(bs []byte) (v V, n int, err error)
}

type Interface[V any] interface {
	Encoder[V]
	Decoder[V]
}
