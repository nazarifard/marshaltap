package ssz

import (
	"math"
	"reflect"
	"time"

	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/pkg/errors"
	ssz "github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/go-ssz/types"
)

type SSZA struct {
	Name     string
	BirthDay uint64 // ssz does not support int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    uint64 // ssz does not support float64
}

type SSZSerializer struct {
	a SSZA
}

// this function is drived from "github.com/prysmaticlabs/go-ssz"
// TODO double check with main developer
func ssz_Marshal(val interface{}, buf []byte) error {
	if val == nil {
		return errors.New("untyped-value nil cannot be marshaled")
	}
	rval := reflect.ValueOf(val)

	// We pre-allocate a buffer-size depending on the value's calculated total byte size.
	//buf := make([]byte, types.DetermineSize(rval))
	factory, err := types.SSZFactory(rval, rval.Type())
	if err != nil {
		return err
	}
	if rval.Type().Kind() == reflect.Ptr {
		if _, err := factory.Marshal(rval.Elem(), rval.Elem().Type(), buf, 0 /* start offset */); err != nil {
			return errors.Wrapf(err, "failed to marshal for type: %v", rval.Elem().Type())
		}
		return nil
	}
	if _, err := factory.Marshal(rval, rval.Type(), buf, 0 /* start offset */); err != nil {
		return errors.Wrapf(err, "failed to marshal for type: %v", rval.Type())
	}
	return nil
}

func (s *SSZSerializer) Sizeof(v goserbench.SmallStruct) int {
	val := any(v)
	rval := reflect.ValueOf(val)
	return int(types.DetermineSize(rval))
}

func (s *SSZSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = uint64(v.BirthDay.UnixNano())
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = math.Float64bits(v.Money)
	return ssz_Marshal(a, buf)
}

func (s *SSZSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	a := &s.a
	err = ssz.Unmarshal(bs, a)
	if err != nil {
		return
	}

	v.Name = a.Name
	v.BirthDay = time.Unix(0, int64(a.BirthDay))
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = math.Float64frombits(a.Money)
	return
}

func NewModem() modem.Interface[goserbench.SmallStruct] {
	return &SSZSerializer{}
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	modem := &SSZSerializer{}
	return tap.NewTap(modem)
}
