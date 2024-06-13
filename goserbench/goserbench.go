package goserbench

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/syncpool"
)

const (
	// MaxSmallStructNameSize is the max size of a name used in the small
	// struct benchmarks.
	MaxSmallStructNameSize = 16

	// MaxSmallStructPhoneSize is the max size of a phone used in the small
	// struct benchmarks.
	MaxSmallStructPhoneSize = 10
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generateSmallStruct() []*SmallStruct {
	a := make([]*SmallStruct, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &SmallStruct{
			Name:     randString(MaxSmallStructNameSize),
			BirthDay: time.Now(),
			Phone:    randString(MaxSmallStructPhoneSize),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

// BenchMarshalSmallStruct benchmarks marshalling the [SmallStruct] type.
func BenchMarshalSmallStruct(b *testing.B, s tap.TapInterface[SmallStruct]) {
	b.Helper()
	data := generateSmallStruct()

	b.ReportAllocs()
	b.ResetTimer()
	var serialSize int
	for i := 0; i < b.N; i++ {
		o := data[rand.Intn(len(data))]
		buf, err := s.Encode(*o)
		if err != nil {
			b.Fatalf("marshal error %s for %#v", err, o)
		}
		serialSize += buf.Len()
		buf.Free()
	}
	b.ReportMetric(float64(serialSize)/float64(b.N), "B/serial")
}

// BenchUnmarshalSmallStruct benchmarks unmarshalling the [SmallStruct] type.
// If validate is true, then the unmarshalled struct is verified to be correct
// against the source struct.
func BenchUnmarshalSmallStruct(b *testing.B, s tap.TapInterface[SmallStruct], validate bool) {
	b.Helper()

	var timePrecision time.Duration
	if stp, ok := s.(SerializerTimePrecision); ok {
		timePrecision = stp.TimePrecision()
	}
	var forcesUTC bool
	if set, ok := s.(SerializerEnforcesTimezone); ok {
		forcesUTC = set.ForcesUTC()
	}

	data := generateSmallStruct()
	ser := make([]syncpool.Buffer, len(data))
	var serialSize int
	for i, d := range data {
		// Reduce the precision of the Birthday field when the
		// serializer cannot represent time with nanosecond precision.
		if timePrecision > 0 {
			d.BirthDay = d.BirthDay.Truncate(timePrecision)
		}

		// Enforce Timezone when serializer requires it.
		if forcesUTC {
			d.BirthDay = d.BirthDay.UTC()
		}

		// Reduce precision of fractional fields when the serializer
		// cannot represent the full float64 range.
		if slfp, ok := s.(SerializerLimitsFloat64Precision); ok {
			fracDigits := slfp.ReduceFloat64Precision()
			i, f := math.Modf(d.Money)
			power := math.Pow(10, float64(fracDigits))
			newf := math.Trunc(f*power) / power
			d.Money = float64(i) + newf
		}

		buf, err := s.Encode(*d)
		if err != nil {
			b.Fatal(err)
		}
		//t := make([]byte, buf.Len())
		//serialSize += copy(t, buf.Bytes())
		serialSize += buf.Len()
		ser[i] = buf
	}
	//o := &SmallStruct{}

	b.ReportMetric(float64(serialSize)/float64(len(data)), "B/serial")
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := rand.Intn(len(ser))
		v, err := s.Decode(ser[n])
		if err != nil {
			b.Fatalf("unmarshal error %s for %#x / %q", err, ser[n], ser[n])
		}
		// Validate unmarshalled data.
		if validate {
			i := data[n]
			correct := v.Name == i.Name && v.Phone == i.Phone && v.Siblings == i.Siblings && v.Spouse == i.Spouse && v.Money == i.Money && v.BirthDay.Equal(i.BirthDay) //&& cmpTags(o.Tags, i.Tags) && cmpAliases(o.Aliases, i.Aliases)
			if !correct {
				b.Fatalf("unmarshaled object differed:\n%v\n%v", i, v)
			}
		}
	}

	for i := range len(ser) {
		ser[i].Free()
	}
}
