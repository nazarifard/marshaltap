package xdrcalmh

import (
	"math"
	"time"

	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type XDRCalmhSerializer struct {
	a XDRA
}

func (s *XDRCalmhSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
	return a.MarshalXDR(buf)
}

func (s *XDRCalmhSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	a := &s.a
	err = a.UnmarshalXDR(bs)
	if err != nil {
		return
	}

	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
	return
}

func (s *XDRCalmhSerializer) Sizeof(v goserbench.SmallStruct) int {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
	return s.a.XDRSize()
}

func NewModem() modem.Interface[goserbench.SmallStruct] {
	return &XDRCalmhSerializer{}
}
func NewTap() tap.Interface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap(modem)
}
