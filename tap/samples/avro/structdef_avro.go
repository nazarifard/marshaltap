package avro

import (
	"bytes"
	"time"

	goavro "github.com/linkedin/goavro"
	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/tap"
	pool "github.com/nazarifard/syncpool"
)

type AvroA[V goserbench.SmallStruct] struct {
	record     *goavro.Record
	codec      goavro.Codec
	bufferPool pool.BufferPool
}

var avroSchemaJSON = `
		{
		  "type": "record",
		  "name": "AvroA",
		  "doc:": "Schema for encoding/decoding sample message",
		  "namespace": "com.example",
		  "fields": [
		    {
		      "name": "name",
		      "type": "string"
		    },
		    {
		      "name": "birthday",
		      "type": "long"
		    },
		    {
		      "name": "phone",
		      "type": "string"
		    },
		    {
		      "name": "siblings",
		      "type": "int"
		    },
		    {
		      "name": "spouse",
		      "type": "boolean"
		    },
		    {
		      "name": "money",
		      "type": "double"
		    }
		  ]
		}
	`

func NewAvroATap() tap.Interface[goserbench.SmallStruct] {
	rec, err := goavro.NewRecord(goavro.RecordSchema(avroSchemaJSON))
	if err != nil {
		panic(err)
	}
	codec, err := goavro.NewCodec(avroSchemaJSON)
	if err != nil {
		panic(err)
	}
	return &AvroA[goserbench.SmallStruct]{record: rec, codec: codec, bufferPool: pool.NewBufferPool()}
}

func (a *AvroA[V]) Encode(v goserbench.SmallStruct) (pool.Buffer, error) {
	a.record.Set("name", v.Name)
	a.record.Set("birthday", int64(v.BirthDay.UnixNano()))
	a.record.Set("phone", v.Phone)
	a.record.Set("siblings", int32(v.Siblings))
	a.record.Set("spouse", v.Spouse)
	a.record.Set("money", v.Money)

	zb := a.bufferPool.Get(0)
	zb.Reset()
	err := a.codec.Encode(zb, a.record)
	return zb, err
}

func (a *AvroA[V]) Decode(bs []byte) (v goserbench.SmallStruct, n int, err error) {
	i, err := a.codec.Decode(bytes.NewReader(bs)) //(*pool.RBuffer)(buf))
	if err != nil {
		return
	}
	rec := i.(*goavro.Record)
	temp, _ := rec.Get("name")
	v.Name = temp.(string)
	temp, _ = rec.Get("birthday")
	v.BirthDay = time.Unix(0, temp.(int64))
	temp, _ = rec.Get("phone")
	v.Phone = temp.(string)
	temp, _ = rec.Get("siblings")
	v.Siblings = int(temp.(int32))
	temp, _ = rec.Get("spouse")
	v.Spouse = temp.(bool)
	temp, _ = rec.Get("money")
	v.Money = temp.(float64)
	return v, 0, err //TODO
}

func (a *AvroA[V]) String() string {
	return "GoAvro"
}
