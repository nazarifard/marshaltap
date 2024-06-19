package marshal

import (
	"testing"
	"time"

	"github.com/nazarifard/marshaltap/tap"
	"github.com/nazarifard/marshaltap/tap/samples/fastape"
	//"github.com/nazarifard/syncpool"
	"github.com/stretchr/testify/assert"
)

func TestTape(t *testing.T) {
	s := S{
		Name:     "Bahador Nazarifard",
		BirthDay: time.Now(),
		Phone:    "09876543210",
		Siblings: 123,
		Spouse:   false,
		Money:    0.1234567890,
	}
	sModem := fastape.NewModem()
	sTap := tap.NewTap(sModem)
	buff, err := sTap.Encode(s)
	//assert.Error(t, err, "tape.Encoder failed")
	s2, _, err := sTap.Decode(buff.Bytes())
	//assert.Error(t, err, "tape.Decoder failed")
	buff.Free()
	assert.Equal(t, s.Money, s2.Money)
	assert.Equal(t, s2.BirthDay.Equal(s.BirthDay), true)
	_ = err
}

func BenchmarkFastape_MarshalTap(b *testing.B) {
	s := S{
		Name:     "Bahador Nazarifard",
		BirthDay: time.Now(),
		Phone:    "09876543210",
		Siblings: 123,
		Spouse:   false,
		Money:    0.1234567890,
	}
	sModem := fastape.NewModem()
	sTap := tap.NewTap(sModem)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buff, err := sTap.Encode(s)
		if err != nil {
			b.Errorf("tape.Encoder failed")
		}
		// _, err = sTap.Decode(buff)
		// if err != nil {
		// 	b.Errorf("tape.Encoder failed")
		// }
		if false {
			_ = buff
		}
		buff.Free()
	}
}

func BenchmarkFastape(b *testing.B) {
	s := S{
		Name:     "Bahador Nazarifard",
		BirthDay: time.Now(),
		Phone:    "09876543210",
		Siblings: 123,
		Spouse:   false,
		Money:    0.1234567890,
	}
	sModem := fastape.NewModem()
	//sTap := tap.NewTap[S, MS](sModem)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, sModem.Sizeof(s))
		err := sModem.Marshal(s, buf)
		if err != nil {
			b.Errorf("tape.Encoder failed")
		}
		if false {
			_ = buf
		}
		//buff.Free()
	}
}

