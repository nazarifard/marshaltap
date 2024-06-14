package pulsar

// import (
// 	"time"

// 	"github.com/nazarifard/marshaltap/goserbench"
// 	"github.com/nazarifard/marshaltap/modem"
// 	pproto "google.golang.org/protobuf/proto"
// )

// type PulsarSerializer struct {
// 	a PulsarBufA
// }

// func (s *PulsarSerializer) ForceUTC() bool {
// 	return true
// }

// func (s *PulsarSerializer) Sizeof(v goserbench.SmallStruct) int {
// 	a := &s.a
// 	a.Name = v.Name
// 	a.BirthDay = v.BirthDay.UnixNano()
// 	a.Phone = v.Phone
// 	a.Siblings = int32(v.Siblings)
// 	a.Spouse = v.Spouse
// 	a.Money = v.Money
// 	return pproto.Size(a)
// }

// func (s *PulsarSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {
// 	a := &s.a
// 	a.Name = v.Name
// 	a.BirthDay = v.BirthDay.UnixNano()
// 	a.Phone = v.Phone
// 	a.Siblings = int32(v.Siblings)
// 	a.Spouse = v.Spouse
// 	a.Money = v.Money
// 	pproto.MarshalOptions() .Marshal(a) //.Marshal(a)
// }

// func (s *PulsarSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
// 	a := &s.a

// 	// Pulsar requires manually claring the fields to their default value.
// 	// *a = PulsarA{}

// 	err = pproto.Unmarshal(bs, a)
// 	if err != nil {
// 		return
// 	}

// 	v.Name = a.Name
// 	v.BirthDay = time.Unix(0, a.BirthDay)
// 	v.Phone = a.Phone
// 	v.Siblings = int(a.Siblings)
// 	v.Spouse = a.Spouse
// 	v.Money = a.Money
// 	return
// }

// func NewModem() modem.ModemInterface[goserbench.SmallStruct] {
// 	return PulsarSerializer{}
// }
