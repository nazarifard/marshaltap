package hprose2

import (
	hprose2 "github.com/hprose/hprose-golang/io"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
)

type Hprose2Serializer struct {
	bufferPool pool.BufferPool
	//writer *hprose2.Writer
	//reader *hprose2.Reader
}

func (s Hprose2Serializer) Encode(v goserbench.SmallStruct) (pool.Buffer, error) {
	zb := s.bufferPool.Get(0)
	zb.Reset()

	bw := hprose2.NewByteWriter(zb.Bytes())
	writer := hprose2.Writer{
		ByteWriter: *bw,
		Simple:     true,
	}
	writer.Clear()
	writer.WriteString(v.Name)
	writer.WriteTime(&v.BirthDay)
	writer.WriteString(v.Phone)
	writer.WriteInt(int64(v.Siblings))
	writer.WriteBool(v.Spouse)
	writer.WriteFloat(v.Money, 64)
	_, err := zb.Write(writer.Bytes())
	return zb, err
}

func (s Hprose2Serializer) Decode(bs []byte) (v goserbench.SmallStruct, n int, err error) {
	o := &v
	//r:=hprose2.NewByteReader(buffer.Bytes())
	reader := hprose2.Reader{
		Simple:    true,
		RawReader: *hprose2.NewRawReader(bs),
	}
	reader.Init(bs)
	o.Name = reader.ReadString()
	o.BirthDay = reader.ReadTime()
	o.Phone = reader.ReadString()
	o.Siblings = int(reader.ReadInt())
	o.Spouse = reader.ReadBool()
	o.Money = reader.ReadFloat64()
	return v, 0, err //TODO
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	return Hprose2Serializer{
		bufferPool: pool.NewBufferPool(),
	}
}
