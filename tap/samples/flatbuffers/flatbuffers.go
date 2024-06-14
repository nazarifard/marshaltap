package flatbuffers

import (
	"time"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
)

type FlatBufferSerializer struct {
	builder    *flatbuffers.Builder
	bufferPool syncpool.BufferPool
}

func (s FlatBufferSerializer) Encode(v goserbench.SmallStruct) (buffer syncpool.Buffer, err error) {
	buffer = s.bufferPool.Get(0)
	buffer.Reset()

	builder := s.builder
	builder.Bytes = buffer.Bytes()
	builder.Reset()

	name := builder.CreateString(v.Name)
	phone := builder.CreateString(v.Phone)

	FlatBufferAStart(builder)
	FlatBufferAAddName(builder, name)
	FlatBufferAAddPhone(builder, phone)
	FlatBufferAAddBirthDay(builder, v.BirthDay.UnixNano())
	FlatBufferAAddSiblings(builder, int32(v.Siblings))
	FlatBufferAAddSpouse(builder, v.Spouse)
	FlatBufferAAddMoney(builder, v.Money)
	builder.Finish(FlatBufferAEnd(builder))
	buffer.Write(builder.Bytes[builder.Head():])
	return
}

func (s FlatBufferSerializer) Decode(bs []byte) (v goserbench.SmallStruct, n int, err error) {
	o := FlatBufferA{}
	o.Init(bs, flatbuffers.GetUOffsetT(bs))
	v.Name = string(o.Name())
	v.BirthDay = time.Unix(0, o.BirthDay())
	v.Phone = string(o.Phone())
	v.Siblings = int(o.Siblings())
	v.Spouse = o.Spouse()
	v.Money = o.Money()
	return v, 0, err //TODO
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	return FlatBufferSerializer{
		builder:    flatbuffers.NewBuilder(0),
		bufferPool: syncpool.NewBufferPool(),
	}
}
