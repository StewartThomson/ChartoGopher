package ChartoGopher

import (
	"math"
	"strconv"
)

type SyncProperty interface {
	//Type of sync property
	Key() string

	Value() string

	Location() string
}

type syncTrack struct {
	TimeSignatures []timeSignature
	Tempos         []tempo
}

type timeSignature struct {
	Numerator   int
	Denominator int
	Position    int
}

type tempo struct {
	Bpm      int
	Position int
}

func (tempo) Key() string {
	return "B"
}

func (t tempo) Value() string {
	return strconv.Itoa(t.Bpm * 1000)
}

func (t tempo) Location() string {
	return strconv.Itoa(t.Position)
}

func (timeSignature) Key() string {
	return "TS"
}

func validateDenominator(denominator int, position int) {
	//Must be greater than 0, less than or equal to 64, and a power of 2
	if denominator <= 0 || denominator > 64 || denominator&(denominator-1) != 0 {
		panic("invalid time signature denominator at position " + strconv.Itoa(position))
	}
}

func (t timeSignature) Value() (out string) {
	out = strconv.Itoa(t.Numerator)

	if t.Denominator == 4 {
		return
	}

	validateDenominator(t.Denominator, t.Position)

	out += " " + strconv.FormatFloat(math.Log2(float64(t.Denominator)), 'f', 0, 64)

	return
}

func (t timeSignature) Location() string {
	return strconv.Itoa(t.Position)
}
