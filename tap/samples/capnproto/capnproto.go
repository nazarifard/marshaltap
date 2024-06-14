package capnproto

import (
	"time"

	capn "github.com/glycerine/go-capnproto"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
)

type CapNProtoSerializer struct {
	bufferPool syncpool.BufferPool
}

func (x *CapNProtoSerializer) Encode(v goserbench.SmallStruct) (syncpool.Buffer, error) {
	seg := capn.NewBuffer(nil)
	c := NewRootCapnpA(seg)
	c.SetName(v.Name)
	c.SetBirthDay(v.BirthDay.UnixNano())
	c.SetPhone(v.Phone)
	c.SetSiblings(int32(v.Siblings))
	c.SetSpouse(v.Spouse)
	c.SetMoney(v.Money)

	zb := x.bufferPool.Get(0)
	zb.Reset()
	_, err := seg.WriteTo(zb)
	return zb, err
}

func (x *CapNProtoSerializer) Decode(bs []byte) (v goserbench.SmallStruct, n int, err error) {
	s, _, err := capn.ReadFromMemoryZeroCopy(bs)
	if err != nil {
		return
	}
	o := ReadRootCapnpA(s)
	v.Name = o.Name()
	v.BirthDay = time.Unix(0, o.BirthDay())
	v.Phone = o.Phone()
	v.Siblings = int(o.Siblings())
	v.Spouse = o.Spouse()
	v.Money = o.Money()
	return v, 0, err //TODO
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	return &CapNProtoSerializer{
		bufferPool: syncpool.NewBufferPool(),
	}
}
