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
	SiblingsTape fastape.UnitTape[int]
	SpouseTape   fastape.UnitTape[bool]
	MoneyTape    fastape.UnitTape[float64]
}

func (cp smallStructTape) Sizeof(p goserbench.SmallStruct) int {
	return cp.NameTape.Sizeof(p.Name) +
		cp.BirthDayTape.Sizeof(p.BirthDay) +
		cp.PhoneTape.Sizeof(p.Phone) +
		cp.SiblingsTape.Sizeof(p.Siblings) +
		cp.SpouseTape.Sizeof(p.Spouse) +
		cp.MoneyTape.Sizeof(p.Money)
}

func (cp smallStructTape) Marshal(p goserbench.SmallStruct, buf []byte) (err error) {
	k, n := 0, 0
	k, _ = cp.NameTape.Marshal(p.Name, buf[n:])
	n += k
	k, _ = cp.BirthDayTape.Marshal(p.BirthDay, buf[n:])
	n += k
	k, _ = cp.PhoneTape.Marshal(p.Phone, buf[n:])
	n += k
	k, _ = cp.SiblingsTape.Marshal(p.Siblings, buf[n:])
	n += k
	k, _ = cp.SpouseTape.Marshal(p.Spouse, buf[n:])
	n += k
	k, _ = cp.MoneyTape.Marshal(p.Money, buf[n:])
	n += k
	return
}

func (cp smallStructTape) Unmarshal(bs []byte, p *goserbench.SmallStruct) (err error) {
	k, n := 0, 0
	k, err = cp.NameTape.Unmarshal(bs[n:], &p.Name)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.BirthDayTape.Unmarshal(bs[n:], &p.BirthDay)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.PhoneTape.Unmarshal(bs[n:], &p.Phone)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.SiblingsTape.Unmarshal(bs[n:], &p.Siblings)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.SpouseTape.Unmarshal(bs[n:], &p.Spouse)
	n += k
	if err != nil {
		return err
	}

	k, err = cp.MoneyTape.Unmarshal(bs[n:], &p.Money)
	n += k
	if err != nil {
		return err
	}

	return
}

func NewModem() modem.ModemInterface[goserbench.SmallStruct] {
	return smallStructTape{}
}

func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	m := NewModem()
	return tap.NewTap[goserbench.SmallStruct, smallStructTape](m)
}
