package easyjson

import (
	easyjson "github.com/mailru/easyjson"
	"github.com/mailru/easyjson/buffer"
	"github.com/mailru/easyjson/jwriter"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
)

type EasyJSONSerializer struct {
	bufferPool syncpool.BufferPool
	a          A
}

func (m *EasyJSONSerializer) Encode(v goserbench.SmallStruct) (syncpool.Buffer, error) {
	a := &m.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
	//easyjson.Marshal(a)

	zb := m.bufferPool.Get(0)
	zb.Reset()

	w := jwriter.Writer{
		Buffer: buffer.Buffer{
			Buf: zb.Bytes(),
		},
	}
	m.a.MarshalEasyJSON(&w)

	zb2 := m.bufferPool.Get(w.Size())
	zb2.Reset()
	bs, err := w.BuildBytes(zb2.Bytes()) //TODO double check
	zb.Free()
	if err != nil {
		zb2.Free()
		return zb2, err
	}
	_, err = zb2.Write(bs)
	return zb2, err
}

func (m *EasyJSONSerializer) Decode(zb syncpool.Buffer) (v goserbench.SmallStruct, err error) {
	a := &m.a
	err = easyjson.Unmarshal(zb.Bytes(), a)
	if err != nil {
		return
	}

	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return
}

func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	return &EasyJSONSerializer{
		bufferPool: syncpool.NewBufferPool(),
	}
}
