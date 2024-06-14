package gencode

import (
	"time"

	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type GencodeSerializer struct {
	//buf []byte
	a GencodeA
}

func (s *GencodeSerializer) Marshal(v goserbench.SmallStruct, bs []byte) error {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	_, err := a.Marshal(bs)
	return err
}

func (s *GencodeSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	a := &s.a
	_, err = a.Unmarshal(bs)
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

func (s *GencodeSerializer) Sizeof(v goserbench.SmallStruct) int {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return int(s.a.Size())
}

func NewModem() modem.Interface[goserbench.SmallStruct] {
	return &GencodeSerializer{}
}
func NewTap() tap.Interface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap[goserbench.SmallStruct, *GencodeSerializer](modem)
}

type GencodeUnsafeSerializer struct {
	//buf []byte
	a GencodeUnsafeA
}

func (s *GencodeUnsafeSerializer) Marshal(v goserbench.SmallStruct, bs []byte) error {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	_, err := a.Marshal(bs)
	return err
}

func (s *GencodeUnsafeSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	a := &s.a
	_, err = a.Unmarshal(bs)
	if err != nil {
		return
	}

	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return
}
func (s *GencodeUnsafeSerializer) Sizeof(v goserbench.SmallStruct) int {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return int(s.a.Size())
}

func NewUnsafeModem() modem.Interface[goserbench.SmallStruct] {
	return &GencodeUnsafeSerializer{}
}
func NewTapUnsafe() tap.Interface[goserbench.SmallStruct] {
	modem := NewUnsafeModem()
	return tap.NewTap[goserbench.SmallStruct, *GencodeUnsafeSerializer](modem)
}
