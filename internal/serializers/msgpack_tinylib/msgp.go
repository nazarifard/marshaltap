package msgpacktinylib

import (
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type MsgpSerializer struct {
	a A
}

func (m *MsgpSerializer) Marshal(v goserbench.SmallStruct, bs []byte) error {
	a := &m.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
	_, err := a.MarshalMsg(bs[:0]) //.MarshalMsg(nil)
	return err
}

func (m *MsgpSerializer) Unmarshal(d []byte, v *goserbench.SmallStruct) error {
	a := &m.a
	_, err := a.UnmarshalMsg(d)
	if err != nil {
		return err
	}

	v.Name = a.Name
	v.BirthDay = a.BirthDay
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return err
}

func (m *MsgpSerializer) Sizeof(v goserbench.SmallStruct) int {
	a := &m.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = v.Siblings
	a.Spouse = v.Spouse
	a.Money = v.Money
	return m.a.Msgsize()
}

func NewModem() modem.ModemInterface[goserbench.SmallStruct] {
	return &MsgpSerializer{}
}
func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap[goserbench.SmallStruct, *MsgpSerializer](modem)
}
