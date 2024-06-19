package fastape

import (
	"github.com/nazarifard/fastape"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type smallStructTape struct {
	NameTape     fastape.StringTape
	BirthDayTape fastape.TimeTape
	PhoneTape    fastape.StringTape
	SiblingsTape fastape.UnitTape[byte]
	SpouseTape   fastape.UnitTape[bool]
	MoneyTape    fastape.UnitTape[float64]
}

func (cp *smallStructTape) Sizeof(p goserbench.SmallStruct) int {
	return cp.NameTape.Sizeof(p.Name) +
		cp.BirthDayTape.Sizeof(p.BirthDay) +
		cp.PhoneTape.Sizeof(p.Phone) +
		cp.SiblingsTape.Sizeof(byte(p.Siblings)) +
		cp.SpouseTape.Sizeof(p.Spouse) +
		cp.MoneyTape.Sizeof(p.Money)
}

func (cp *smallStructTape) Marshal(p goserbench.SmallStruct, buf []byte) (err error) {
	k, n := 0, 0
	k, _ = cp.NameTape.Roll(p.Name, buf[n:])
	n += k
	k, _ = cp.BirthDayTape.Roll(p.BirthDay, buf[n:])
	n += k
	k, _ = cp.PhoneTape.Roll(p.Phone, buf[n:])
	n += k
	k, _ = cp.SiblingsTape.Roll(byte(p.Siblings), buf[n:])
	n += k
	k, _ = cp.SpouseTape.Roll(p.Spouse, buf[n:])
	n += k
	k, _ = cp.MoneyTape.Roll(p.Money, buf[n:])
	n += k
	return
}

func (cp *smallStructTape) Unmarshal(bs []byte, p *goserbench.SmallStruct) (err error) {
	k, n := 0, 0
	k, err = cp.NameTape.Unroll(bs[n:], &p.Name)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.BirthDayTape.Unroll(bs[n:], &p.BirthDay)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.PhoneTape.Unroll(bs[n:], &p.Phone)
	n += k
	if err != nil {
		return err
	}

	var sib byte
	k, err = cp.SiblingsTape.Unroll(bs[n:], &sib)
	p.Siblings = int(sib)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.SpouseTape.Unroll(bs[n:], &p.Spouse)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.MoneyTape.Unroll(bs[n:], &p.Money)
	n += k
	if err != nil {
		return err
	}

	return
}

func NewModem() modem.Interface[goserbench.SmallStruct] {
	return &smallStructTape{}
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	m := NewModem()
	return tap.NewTap(m)
}
