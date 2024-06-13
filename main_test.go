package marshaltap

import (
	"fmt"
	"os"
	"testing"

	"github.com/nazarifard/marshaltap/goserbench"
	"github.com/nazarifard/marshaltap/internal/serializers/avro"
	"github.com/nazarifard/marshaltap/internal/serializers/baseline"
	bebop200sc "github.com/nazarifard/marshaltap/internal/serializers/bebop_200sc"
	bebopwellquite "github.com/nazarifard/marshaltap/internal/serializers/bebop_wellquite"
	"github.com/nazarifard/marshaltap/internal/serializers/benc"
	binaryalecthomas "github.com/nazarifard/marshaltap/internal/serializers/binary_alecthomas"
	"github.com/nazarifard/marshaltap/internal/serializers/bson"
	"github.com/nazarifard/marshaltap/internal/serializers/capnproto"
	"github.com/nazarifard/marshaltap/internal/serializers/colfer"
	"github.com/nazarifard/marshaltap/internal/serializers/easyjson"
	"github.com/nazarifard/marshaltap/internal/serializers/fastape"
	"github.com/nazarifard/marshaltap/internal/serializers/fastjson"
	"github.com/nazarifard/marshaltap/internal/serializers/flatbuffers"
	"github.com/nazarifard/marshaltap/internal/serializers/gencode"
	"github.com/nazarifard/marshaltap/internal/serializers/gogo"
	"github.com/nazarifard/marshaltap/internal/serializers/jsoniter"
	"github.com/nazarifard/marshaltap/internal/serializers/mongobson"

	//"github.com/nazarifard/marshaltap/internal/serializers/gotiny"
	"github.com/nazarifard/marshaltap/internal/serializers/hprose"
	"github.com/nazarifard/marshaltap/internal/serializers/hprose2"
	"github.com/nazarifard/marshaltap/internal/serializers/ikea"
	msgpacktinylib "github.com/nazarifard/marshaltap/internal/serializers/msgpack_tinylib"
	"github.com/nazarifard/marshaltap/internal/serializers/mus"
	"github.com/nazarifard/marshaltap/internal/serializers/ssz"
	"github.com/nazarifard/marshaltap/internal/serializers/stdlib"
	xdrcalmh "github.com/nazarifard/marshaltap/internal/serializers/xdr_calmh"
	"github.com/nazarifard/marshaltap/tap"
)

var (
	validate = os.Getenv("VALIDATE") != ""
)

func TestMessage(t *testing.T) {
	fmt.Print(`A test suite for benchmarking various Go serialization methods.

See README.md for details on running the benchmarks.

`)
}

type BenchmarkCase struct {
	Name string
	URL  string
	//New  func() goserbench.Serializer
	New func() tap.TapInterface[goserbench.SmallStruct]
}

var benchmarkCases = []BenchmarkCase{
	{
		// 	Name: "gotiny",
		// 	URL:  "github.com/niubaoshu/gotiny",
		// 	New:  gotiny.,
		// }, {
		Name: "msgp",
		URL:  "github.com/tinylib/msgp",
		New:  msgpacktinylib.NewTap, //.NewMsgpSerializer,
		// }, {
		// 	Name: "msgpack",
		// 	URL:  "github.com/vmihailenco/msgpack",
		// 	New:  msgpackvmihailenco.VmihailencoMsgpackSerializer(),
	}, {
		Name: "json",
		URL:  "pkg.go/dev/encoding/json",
		New:  stdlib.NewJsonTap[goserbench.SmallStruct],
	}, {
		Name: "jsoniter",
		URL:  "github.com/json-iterator/go",
		New:  jsoniter.NewTap,
	}, {
		Name: "easyjson",
		URL:  "github.com/mailru/easyjson",
		New:  easyjson.NewTap,
	}, {
		Name: "bson",
		URL:  "gopkg.in/mgo.v2/bson",
		New:  bson.NewTap,
	}, {
		Name: "mongobson",
		URL:  "go.mongodb.org/mongo-driver/mongo",
		New:  mongobson.NewTap,
	}, {
		Name: "gob",
		URL:  "pkg.go.dev/encoding/gob",
		New:  stdlib.NewGobTap[goserbench.SmallStruct],
		// }, {
		// 	Name: "davecgh/xdr",
		// 	URL:  "github.com/davecgh/go-xdr/xdr",
		// 	New:  xdrdavecgh.NewXDRDavecghSerializer,
		// }, {
		// 	Name: "ugorji/msgpack",
		// 	URL:  "github.com/ugorji/go/codec",
		// 	New:  ugorji.NewMsgPackTap,
		// }, {
		// 	Name: "ugorji/binc",
		// 	URL:  "github.com/ugorji/go/codec",
		// 	New:  ugorji.NewTap,
		// }, {
		// 	Name: "sereal",
		// 	URL:  "github.com/Sereal/Sereal/Go/sereal",
		// 	New:  sereal.NewSerealSerializer,
	}, {
		Name: "alecthomas/binary",
		URL:  "github.com/alecthomas/binary",
		New:  binaryalecthomas.NewTap[goserbench.SmallStruct],
	}, {
		Name: "flatbuffers",
		URL:  "github.com/google/flatbuffers/go",
		New:  flatbuffers.NewTap,
	}, {
		Name: "capnproto",
		URL:  "github.com/glycerine/go-capnproto",
		New:  capnproto.NewTap,
	}, {
		Name: "hprose",
		URL:  "github.com/hprose/hprose-go/io",
		New:  hprose.NewTap,
	}, {
		Name: "hprose2",
		URL:  "github.com/hprose/hprose-golang/io",
		New:  hprose2.NewTap,
		// }, {
		// 	Name: "dedis/protobuf",
		// 	URL:  "go.dedis.ch/protobuf",
		// 	New:  protobufdedis.NewProtobufSerializer,
		// }, {
		// 	Name: "pulsar",
		// 	URL:  "github.com/cosmos/cosmos-proto",
		// 	New:  pulsar.NewPulsarSerializer,
		// }, {
		// 	Name: "gogo/protobuf",
		// 	URL:  "github.com/gogo/protobuf/proto",
		// 	New:  gogo.NewGogoProtoSerializer,
	}, {
		Name: "gogo/jsonpb",
		URL:  "github.com/gogo/protobuf/proto",
		New:  gogo.NewJSonTap,
	}, {
		Name: "colfer",
		URL:  "github.com/pascaldekloe/colfer",
		New:  colfer.NewTap,
	}, {
		Name: "gencode",
		URL:  "github.com/andyleap/gencode",
		New:  gencode.NewTap,
	}, {
		Name: "gencode/unsafe",
		URL:  "github.com/andyleap/gencode",
		New:  gencode.NewTapUnsafe,
	}, {
		Name: "calmh/xdr",
		URL:  "github.com/calmh/xdr",
		New:  xdrcalmh.NewTap,
	}, {
		Name: "goavro",
		URL:  "gopkg.in/linkedin/goavro.v1",
		New:  avro.NewAvroATap,
		// }, {
		// 	Name: "avro2/text",
		// 	URL:  "github.com/linkedin/goavro",
		// 	New:  avro.NewAvroATap(),
		// }, {
		// 	Name: "avro2/binary",
		// 	URL:  "github.com/linkedin/goavro",
		// 	New:  avro.NewAvro2Bin,
	}, {
		Name: "ikea",
		URL:  "github.com/ikkerens/ikeapack",
		New:  ikea.NewTap,
		// }, {
		// 	Name: "shamaton/msgpack/map",
		// 	URL:  "github.com/shamaton/msgpack",
		// 	New:  shamaton.NewShamatonMapMsgpackSerializer,
		// }, {
		// 	Name: "shamaton/msgpack/array",
		// 	URL:  "github.com/shamaton/msgpack",
		// 	New:  shamaton.NewShamatonArrayMsgPackSerializer,
		// }, {
		// 	Name: "shamaton/msgpackgen/map",
		// 	URL:  "github.com/shamaton/msgpack",
		// 	New:  shamaton.NewShamatonMapMsgPackgenSerializer,
		// }, {
		// 	Name: "shamaton/msgpackgen/array",
		// 	URL:  "github.com/shamaton/msgpack",
		// 	New:  shamaton.NewShamatonArrayMsgpackgenSerializer,
	}, {
		Name: "ssz",
		URL:  "github.com/prysmaticlabs/go-ssz",
		New:  ssz.NewTap,
	}, {
		Name: "200sc/bebop",
		URL:  "github.com/200sc/bebop",
		New:  bebop200sc.NewTap,
	}, {
		Name: "wellquite/bebop",
		URL:  "wellquite.org/bebop",
		New:  bebopwellquite.NewTap,
	}, {
		Name: "fastjson",
		URL:  "github.com/valyala/fastjson",
		New:  fastjson.NewTap,
	}, {
		Name: "benc",
		URL:  "github.com/deneonet/benc",
		New:  benc.NewTap,
	}, {
		Name: "benc/usafe",
		URL:  "github.com/deneonet/benc",
		New:  benc.NewUnsafeTap,
	}, {
		Name: "mus",
		URL:  "github.com/mus-format/mus-go",
		New:  mus.NewTap,
	}, {
		Name: "mus/unsafe",
		URL:  "github.com/mus-format/mus-go",
		New:  mus.NewTap,
	}, {
		Name: "baseline",
		URL:  "",
		New:  baseline.NewTap,
	}, {
		Name: "fastape",
		URL:  "github.com/nazarifard/fastape",
		New:  fastape.NewTap,
	},
}

func BenchmarkSerializers(b *testing.B) {
	for i := range benchmarkCases[:2] {
		bc := benchmarkCases[i]
		b.Run("marshal/"+bc.Name, func(b *testing.B) {
			goserbench.BenchMarshalSmallStruct(b, bc.New())
		})
		b.Run("unmarshal/"+bc.Name, func(b *testing.B) {
			goserbench.BenchUnmarshalSmallStruct(b, bc.New(), validate)
		})
	}
}
