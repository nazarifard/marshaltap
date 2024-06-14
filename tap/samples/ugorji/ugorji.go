package ugorji

import (
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/ugorji/go/codec"
)

type UgorjiCodecSerializer[V any] struct {
	codec.Handle
}

func (u *UgorjiCodecSerializer[V]) Marshal(o V, bs []byte) error {
	return codec.NewEncoderBytes(&bs, u.Handle).Encode(o)
}

func (u *UgorjiCodecSerializer[V]) Unmarshal(d []byte, o *V) error {
	return codec.NewDecoderBytes(d, u.Handle).Decode(o)
}

func (u *UgorjiCodecSerializer[V]) Sizeof(v V) int {
	panic("UgorjiCodecSerializer.Sizeof method not implemented yet")
}

func NewMsgPackModem() modem.Interface[goserbench.SmallStruct] {
	return &UgorjiCodecSerializer[goserbench.SmallStruct]{&codec.MsgpackHandle{}}
}

func NewModem() modem.Interface[goserbench.SmallStruct] {
	h := &codec.BincHandle{}
	h.AsSymbols = 0
	return &UgorjiCodecSerializer[goserbench.SmallStruct]{h}
}

func NewMsgPackTap() tap.Interface[goserbench.SmallStruct] {
	modem := NewMsgPackModem()
	return tap.NewTap[goserbench.SmallStruct, *UgorjiCodecSerializer[goserbench.SmallStruct]](modem)
	//{&codec.MsgpackHandle{}}
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	modem := NewModem()
	return tap.NewTap[goserbench.SmallStruct, *UgorjiCodecSerializer[goserbench.SmallStruct]](modem)
	//{&codec.MsgpackHandle{}}
}
