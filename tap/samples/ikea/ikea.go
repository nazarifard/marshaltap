package ikea

import (
	"bytes"
	"math"
	"time"

	ikea "github.com/ikkerens/ikeapack"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
)

type IkeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    uint64 // NOTE: Ike does not support float64 - it needs to be converted to an int type.
}

type IkeaSerializer struct {
	a IkeA
	//buf *bytes.Buffer
	bufferPool syncpool.BufferPool
}

func (s IkeaSerializer) Encode(v goserbench.SmallStruct) (zb syncpool.Buffer, err error) {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)

	zb = s.bufferPool.Get(0)
	zb.Reset()
	err = ikea.Pack(zb, a)
	return
}

func (s IkeaSerializer) Decode(bs []byte) (v goserbench.SmallStruct, n int, err error) {
	a := &s.a
	err = ikea.Unpack(bytes.NewReader(bs), a)
	if err != nil {
		return
	}

	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
	return v, 0, err //TODO
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	return IkeaSerializer{
		bufferPool: syncpool.NewBufferPool(),
	}
}
