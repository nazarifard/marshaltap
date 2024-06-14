package fastjson

import (
	"time"

	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
	"github.com/valyala/fastjson"
)

type FastJSONSerializer struct {
	arena      fastjson.Arena
	object     *fastjson.Value
	bufferPool syncpool.BufferPool
	//buf    []byte
}

func (s FastJSONSerializer) Encode(v goserbench.SmallStruct) (buf syncpool.Buffer, err error) {
	object, arena := s.object, s.arena
	object.Set("name", arena.NewString(v.Name))
	object.Set("birthday", arena.NewNumberInt(int(v.BirthDay.UnixNano())))
	object.Set("phone", arena.NewString(v.Phone))
	object.Set("siblings", arena.NewNumberInt(v.Siblings))
	var spouse *fastjson.Value
	if v.Spouse {
		spouse = arena.NewTrue()
	} else {
		spouse = arena.NewFalse()
	}
	object.Set("spouse", spouse)
	object.Set("money", arena.NewNumberFloat64(v.Money))

	zb := s.bufferPool.Get(0)
	zb.Reset()
	dest := object.MarshalTo(zb.Bytes())
	_, err = zb.Write(dest)
	return zb, err
}

func (s FastJSONSerializer) Decode(bs []byte) (v goserbench.SmallStruct, n int, err error) {
	val, err := fastjson.ParseBytes(bs)
	if err != nil {
		return
	}
	v.Name = string(val.GetStringBytes("name"))
	v.BirthDay = time.Unix(0, val.GetInt64("birthday"))
	v.Phone = string(val.GetStringBytes("phone"))
	v.Siblings = val.GetInt("siblings")
	v.Spouse = val.GetBool("spouse")
	v.Money = val.GetFloat64("money")
	return v, 0, err //TODO
}

func NewTap() tap.Interface[goserbench.SmallStruct] {
	var arena fastjson.Arena
	return FastJSONSerializer{
		object:     arena.NewObject(),
		arena:      arena,
		bufferPool: syncpool.NewBufferPool(),
	}
}
