package bson

import (
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

type BsonTap[V any] struct {
	bufferPool pool.BufferPool
}

func (m *BsonTap[V]) Encode(v V) (buf pool.Buffer, err error) {
	buf = m.bufferPool.Get(0)
	buf.Reset()
	vw, err := bsonrw.NewBSONValueWriter(buf)
	if err != nil {
		buf.Free()
		return
	}
	encoder, err := bson.NewEncoder(vw)
	if err != nil {
		return
	}
	err = encoder.Encode(v)
	return
}

func (m *BsonTap[V]) Decode(bs []byte) (v V, err error) {
	decoder, err := bson.NewDecoder(bsonrw.NewBSONDocumentReader(bs))
	if err != nil {
		return
	}
	err = decoder.Decode(&v)
	return v, err
}

func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	return &BsonTap[goserbench.SmallStruct]{
		bufferPool: pool.NewBufferPool(),
	}
}
