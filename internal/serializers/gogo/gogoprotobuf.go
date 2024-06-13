package gogo

import (
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
)

type GogoProtoSerializer struct {
	a          GogoProtoBufA
	bufferPool syncpool.BufferPool

	// marshaller and unmarshaller are set on creation to either binary
	// or json marshallers.
	marshaller   func(proto.Message) (syncpool.Buffer, error)
	unmarshaller func(syncpool.Buffer, proto.Message) error
}

func (s *GogoProtoSerializer) Encode(v goserbench.SmallStruct) (buf syncpool.Buffer, err error) {
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return s.marshaller(a)
}

func (s *GogoProtoSerializer) Decode(buf syncpool.Buffer) (v goserbench.SmallStruct, err error) {
	// NOTE: gogoproto serialization in jsonpb mode does not automatically
	// clear fields with their empty value.
	a := &s.a
	*a = GogoProtoBufA{}

	err = s.unmarshaller(buf, a)
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

// func NewProtoModem() modem.ModemInterface[goserbench.SmallStruct] {
// 	return &GogoProtoSerializer{
// 		marshaller:   proto.Marshal,
// 		unmarshaller: proto.Unmarshal,
// 	}
// }

func NewJSonTap() tap.TapInterface[goserbench.SmallStruct] {
	marshaller := &jsonpb.Marshaler{}
	bufferPool := syncpool.NewBufferPool()

	x := GogoProtoSerializer{
		marshaller: func(a proto.Message) (syncpool.Buffer, error) {
			buf := bufferPool.Get(0)
			buf.Reset()
			err := marshaller.Marshal(buf, a)
			return buf, err
		},
		unmarshaller: func(b syncpool.Buffer, a proto.Message) (err error) {
			err = jsonpb.Unmarshal((*syncpool.RBuffer)(b), a)
			return err
		},
	}
	return &x
}
