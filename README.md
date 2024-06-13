## MarshalTap
 Whenever we have to load and marshal an amount of huge stream data memory management is a challemge.
 Frequently memory allocation for data marashaling forces unnecessary load to GC and finally it causes low performance and even crash in some cases.
 MarshalTap is designed to prevent unnecessary memory allocation during marshal and or unmarshal data process.
 We should try to have Zero memory allocation per each Encoding or Decoding oprand. 
 Usually even best designed data serializers causes at leat one memory allocating per oprand. if you have to handle a big bandwith of real time data, should be care about memory allocation always.
 
 MarshalTap is a tap module that can be installed before a data serializer module for reduction of memory allocations.
 Zero memory allocation is always a ideal case and the following tables shows how Marshal-Tap prevented extra memory allocation successfully. 

 MarshalTap is a general module and can be used with of any availiable data serializer module.
 Each data serializer module that can provide an interface with 3 methods, it can connected to marshaltap easily.
 Firstly MarshalTap is designed for Fastape high performance data serializer but it designed to work with any arbitrary marshaller.
 Each data serializer is known a modem and any modem interface should provide these methods:
 ```go
  type ModemInterface[V any] interface {
  Marshal(v V, buf []byte) (err error)
  Unmarshal(buf []byte, v *V) (err error)
  Sizeof(v V) int
  }
 ```
 ## Technology
  MarshalTap uses of [syncpool](https://github.com/nazarifard/syncpool) module to reduce memory allocations.
  
 ## goserbench
  For benchmark reason first version of current project is forked and drived from [goserbench](https://github.com/alecthomas/go_serialization_benchmarks)
  However it may not be sync with goserbench in the future. Because **marshaltap** is designed to provide a tap for other data serializers not for benchmarking.
  But others including benchmarking developers can use MarshalTap to work with a various range of serializers easily.
 
 ## efficiency 
 MarshalTap may change the game. This benchmarks shows how MarshalTap reduces unneccessary memory usage and improve performance efficiently.
 ```sh
.───┬──────────────────┬─────────┬───────┬──────┬───────────.
│ # │       name       │    #    │ ns/op │ B/op │ allocs/op │
├───┼──────────────────┼─────────┼───────┼──────┼───────────┤
│ 0 │     easyjson-4   │  475309 | 2853  │ 976  │ 7         │
│ 0 │ Tap-easyjson-4   │ 1606742 │  756  │  48  │ 1         │
'───┴──────────────────┴─────────┴───────┴──────┴───────────'
```

Also some of other results show how MarshalTap can reduce memory allocations even for best and fast serializers. The following tables show results of marshallers with and without MarshalTap.
```sh
## With MarshalTap
.───┬──────────────────┬─────────┬───────┬──────┬───────────.
│ # │       name       │    #    │ ns/op │ B/op │ allocs/op │
├───┼──────────────────┼─────────┼───────┼──────┼───────────┤
│ 0 │ benc/usafe-4     │ 8546756 │ 136.7 │ 0    │ 0         │
│ 1 │ benc-4           │ 7955727 │ 146.0 │ 0    │ 0         │
│ 2 │ gencode/unsafe-4 │ 6888806 │ 160.5 │ 0    │ 0         │
│ 3 │ msgp-4           │ 6804525 │ 166.9 │ 0    │ 0         │
│ 4 │ mus/unsafe-4     │ 7116132 │ 167.1 │ 0    │ 0         │
│ 5 │ mus-4            │ 7163511 │ 167.9 │ 0    │ 0         │
│ 6 │ fastape-4        │ 6420730 │ 184.6 │ 0    │ 0         │
│ 7 │ colfer-4         │ 6094623 │ 185.7 │ 0    │ 0         │
│ 8 │ calmh/xdr-4      │ 5549302 │ 213.0 │ 0    │ 0         │
│ 9 │ flatbuffers-4    │ 2899107 │ 411.6 │ 0    │ 0         │
'───┴──────────────────┴─────────┴───────┴──────┴───────────'

## Without MarshalTap
.───┬──────────────────┬─────────┬───────┬──────┬───────────.
│ # │       name       │    #    │ ns/op │ B/op │ allocs/op │
├───┼──────────────────┼─────────┼───────┼──────┼───────────┤
│ 0 │ benc/usafe-4     │ 7776930 │ 153.7 │ 64   │ 1         │
│ 1 │ benc-4           │ 7669548 │ 158.3 │ 64   │ 1         │
│ 2 │ gencode-4        │ 6780241 │ 166.7 │ 16   │ 1         │
│ 3 │ mus-4            │ 7139898 │ 182.5 │ 48   │ 1         │
│ 4 │ colfer-4         │ 5930238 │ 198.6 │ 64   │ 1         │
│ 5 │ fastape-4        │ 6018102 │ 199.5 │ 64   │ 1         │
│ 6 │ gogo/protobuf-4  │ 5589772 │ 206.6 │ 64   │ 1         │
│ 7 │ msgp-4           │ 5561642 │ 215.1 │ 128  │ 1         │
│ 8 │ calmh/xdr-4      │ 4842514 │ 250.3 │ 64   │ 1         │
│ 9 │ mus/unsafe-4     │ 7380118 │ 268.9 │ 64   │ 1         │
'───┴──────────────────┴─────────┴───────┴──────┴───────────'
```

## Current status
  currently masrhaltap is tested with some of serializers that is used by [goserbench](https://github.com/alecthomas/go_serialization_benchmarks) project.
  but they are not provides all 3 required method easily, then I cant support them yet. 
  
## Usage
 ```go
 import (
	"github.com/nazarifard/fastape"
	"github.com/nazarifard/marshaltap/modem"
	"github.com/nazarifard/marshaltap/tap"
 )

 type S string
 
 type SModem struct { 
      S fastape.StringTape 
 }
 func (m *SModem) Sizeof(s S) int {
 	return m.S.Sizeof(s)
 }
 func (m *SModem) Marshal(s S, buf []byte) error {
 	_, err := m.S.Marshal(s, buf)
 	return err
 }
 func (m *SModem) Unmarshal(bs []byte, s S) error {
 	_, err = m.S.Unmarshal(bs, &s)
 	return
 }
 
 func main() {
 	s:="hello"
 	stap:=tap.NewTap[S, *SModem](s)	

 	buf, _ := stap.Encode(s)         //get buf
 	s2,_ := stap.Decode(buf.Bytes()) //use buf
 	buf.Free()                       //free buf	
	
 	print(s2)
 }
```
## license
 **MIT**
    
  