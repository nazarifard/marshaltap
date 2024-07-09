package marshal

import (
	"github.com/nazarifard/fastape"
	"github.com/nazarifard/syncpool"
)

func NewFastap[V any](tape fastape.Tape[V], pool *syncpool.Pool[V]) Interface[V] {
	return fastape.NewMarshalTap(tape, pool)
}
