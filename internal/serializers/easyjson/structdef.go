package easyjson

import (
	"time"
)

// This is a copy of the goserbench.SmallStruct because easyjson code generation
// adds methods to it.

//easyjson:json
type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

type NoTimeA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int
	Spouse   bool
	Money    float64
}

type NoTimeNoStringNoFloatA struct {
	Name     []byte
	BirthDay uint64
	Phone    []byte
	Siblings uint32
	Spouse   bool
	Money    uint64
}
