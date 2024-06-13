package bebop200sc

import (
	"time"

	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type Bebop200ScSerializer struct {
	a BebopBuf200sc
	//buf []byte
}

func (s *Bebop200ScSerializer) Sizeof(v goserbench.SmallStruct) int {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return s.a.Size()
}
func (s *Bebop200ScSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	_ = a.MarshalBebopTo(buf)
	return nil
}

func (s *Bebop200ScSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	a := &s.a
	err = a.UnmarshalBebop(bs)
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

func (s *Bebop200ScSerializer) TimePrecision() time.Duration {
	return 100 * time.Nanosecond
}

func (s *Bebop200ScSerializer) ForceUTC() bool {
	return true
}

func NewModem() modem.ModemInterface[goserbench.SmallStruct] {
	return &Bebop200ScSerializer{}
}

func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap[goserbench.SmallStruct, *Bebop200ScSerializer](modem)
}
