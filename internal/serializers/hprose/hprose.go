package hprose

import (
	"bytes"

	"github.com/hprose/hprose-go"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
)

type HproseSerializer struct {
	bufferPool pool.BufferPool
}

func (s *HproseSerializer) Encode(a goserbench.SmallStruct) (pool.Buffer, error) {
	zb := s.bufferPool.Get(0)
	zb.Reset()

	writer := hprose.NewWriter(bytes.NewBuffer(zb.Bytes()), true)
	writer.WriteString(a.Name)
	writer.WriteTime(a.BirthDay)
	writer.WriteString(a.Phone)
	writer.WriteInt64(int64(a.Siblings))
	writer.WriteBool(a.Spouse)
	writer.WriteFloat64(a.Money)
	_, err := zb.Write(writer.Stream.(*bytes.Buffer).Bytes())
	return zb, err
}

func (s *HproseSerializer) Decode(zb pool.Buffer) (v goserbench.SmallStruct, err error) {
	o := &v
	reader := hprose.NewReader(bytes.NewBuffer(zb.Bytes()), true)
	o.Name, err = reader.ReadString()
	if err != nil {
		return
	}
	o.BirthDay, err = reader.ReadDateTime()
	if err != nil {
		return
	}
	o.Phone, err = reader.ReadString()
	if err != nil {
		return
	}
	o.Siblings, err = reader.ReadInt()
	if err != nil {
		return
	}
	o.Spouse, err = reader.ReadBool()
	if err != nil {
		return
	}
	o.Money, err = reader.ReadFloat64()
	return
}

func NewModem() tap.TapInterface[goserbench.SmallStruct] {
	// buf := new(bytes.Buffer)
	// reader := hprose.NewReader(buf, true)
	// bufw := new(bytes.Buffer)
	// writer := hprose.NewWriter(bufw, true)
	return &HproseSerializer{
		bufferPool: pool.NewBufferPool(),
	}
}

func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	return &HproseSerializer{
		bufferPool: pool.NewBufferPool(),
	}
}
