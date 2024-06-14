package colfer

import (
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type ColferSerializer struct {
	a ColferA
}

func (s *ColferSerializer) ForceUTC() bool {
	return true
}

func (s *ColferSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {

	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	_ = a.MarshalTo(buf)
	return
}

func (s *ColferSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	a := &s.a

	// Colfer requires manually claring the fields to their default value.
	*a = ColferA{}

	err = a.UnmarshalBinary(bs)
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
func (s *ColferSerializer) Sizeof(v goserbench.SmallStruct) int {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	n, _ := s.a.MarshalLen()
	return n
}
func NewModem() modem.Interface[goserbench.SmallStruct] {
	return &ColferSerializer{}
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap[goserbench.SmallStruct, *ColferSerializer](modem)
}
