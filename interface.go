package marshal

import "github.com/nazarifard/syncpool"

type Encoder[V any] interface {
	Encode(V) (syncpool.Buffer, error)
}

type Decoder[V any] interface {
	Decode(bs []byte) (v *V, n int, err error)
}

type Interface[V any] interface {
	Encoder[V]
	Decoder[V]
	Free(*V)
}
