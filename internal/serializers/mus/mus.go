package mus

import (
	"time"

	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/raw"
	"github.com/mus-format/mus-go/unsafe"
	"github.com/mus-format/mus-go/varint"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
)

type MUSSerializer struct{}

func (s MUSSerializer) Sizeof(v goserbench.SmallStruct) int {
	n := ord.SizeString(v.Name)
	n += raw.SizeInt64(v.BirthDay.UnixNano())
	n += ord.SizeString(v.Phone)
	n += varint.SizeInt32(int32(v.Siblings))
	n += ord.SizeBool(v.Spouse)
	n += raw.SizeFloat64(v.Money)
	return n
}

func (s MUSSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {
	n := ord.MarshalString(v.Name, buf)
	n += raw.MarshalInt64(v.BirthDay.UnixNano(), buf[n:])
	n += ord.MarshalString(v.Phone, buf[n:])
	n += varint.MarshalInt32(int32(v.Siblings), buf[n:])
	n += ord.MarshalBool(v.Spouse, buf[n:])
	raw.MarshalFloat64(v.Money, buf[n:])
	return nil
}

func (s MUSSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {

	var n int

	v.Name, n, err = ord.UnmarshalString(bs)
	if err != nil {
		return
	}
	var n1 int
	var bdayNano int64
	bdayNano, n1, err = raw.UnmarshalInt64(bs[n:])
	v.BirthDay = time.Unix(0, bdayNano)
	n += n1
	if err != nil {
		return
	}
	v.Phone, n1, err = ord.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	var sibInt32 int32
	sibInt32, n1, err = varint.UnmarshalInt32(bs[n:])
	v.Siblings = int(sibInt32)
	n += n1
	if err != nil {
		return
	}
	v.Spouse, n1, err = ord.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Money, n1, err = raw.UnmarshalFloat64(bs[n:])
	n += n1
	return
}

func NewModem() modem.ModemInterface[goserbench.SmallStruct] {
	return MUSSerializer{}
}

func NewTap() tap.TapInterface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap[goserbench.SmallStruct, MUSSerializer](modem)
}

type MUSUnsafeSerializer struct{}

func (s MUSUnsafeSerializer) Sizeof(v goserbench.SmallStruct) int {
	n := ord.SizeString(v.Name)
	n += raw.SizeInt64(v.BirthDay.UnixNano())
	n += ord.SizeString(v.Phone)
	n += varint.SizeInt32(int32(v.Siblings))
	n += ord.SizeBool(v.Spouse)
	n += raw.SizeFloat64(v.Money)
	return n
}

func (s MUSUnsafeSerializer) Marshal(v goserbench.SmallStruct, buf []byte) (err error) {
	n := unsafe.MarshalString(v.Name, buf)
	n += unsafe.MarshalInt64(v.BirthDay.UnixNano(), buf[n:])
	n += unsafe.MarshalString(v.Phone, buf[n:])
	n += unsafe.MarshalInt32(int32(v.Siblings), buf[n:])
	n += unsafe.MarshalBool(v.Spouse, buf[n:])
	unsafe.MarshalFloat64(v.Money, buf[n:])
	return nil
}

func (s MUSUnsafeSerializer) Unmarshal(bs []byte, v *goserbench.SmallStruct) (err error) {
	var n int

	v.Name, n, err = unsafe.UnmarshalString(bs)
	if err != nil {
		return
	}
	var n1 int
	var bdayNano int64
	bdayNano, n1, err = unsafe.UnmarshalInt64(bs[n:])
	v.BirthDay = time.Unix(0, bdayNano)
	n += n1
	if err != nil {
		return
	}
	v.Phone, n1, err = unsafe.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	var sibInt32 int32
	sibInt32, n1, err = unsafe.UnmarshalInt32(bs[n:])
	v.Siblings = int(sibInt32)
	n += n1
	if err != nil {
		return
	}
	v.Spouse, n1, err = unsafe.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Money, n1, err = unsafe.UnmarshalFloat64(bs[n:])
	n += n1
	return
}

func NewUnsafeModem() modem.ModemInterface[goserbench.SmallStruct] {
	return MUSUnsafeSerializer{}
}

func NewUnsafeTap() tap.TapInterface[goserbench.SmallStruct] {
	modem := NewUnsafeModem()
	return tap.NewTap[goserbench.SmallStruct, MUSUnsafeSerializer](modem)
}
