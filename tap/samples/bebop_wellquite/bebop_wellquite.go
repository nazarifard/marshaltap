package bebopwellquite

import (
	"time"

	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type BebopWellquiteSerializer struct {
	a BebopBufWellquite
}

func (s *BebopWellquiteSerializer) Sizeof(v goserbench.SmallStruct) int {
	return s.a.SizeBebop()
}

func (s *BebopWellquiteSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	_, err = a.MarshalBebop(buf)
	return
}

func (s *BebopWellquiteSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	a := &s.a
	_, err = a.UnmarshalBebop(bs)
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

func (s *BebopWellquiteSerializer) TimePrecision() time.Duration {
	return 100 * time.Nanosecond
}

func NewModem() modem.Interface[goserbench.SmallStruct] {
	return &BebopWellquiteSerializer{}
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap(modem)
}
