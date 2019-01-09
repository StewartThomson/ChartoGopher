package ChartoGopher

import "strconv"

type SyncProperty interface {
	//Type of sync property
	key() string

	value() string

	location() string
}

type SyncTrack struct {
	TimeSignatures []TimeSignature
	Tempos         []Tempo
}

type TimeSignature struct {
	Numerator   int
	Denominator int
	Position    int
}

type Tempo struct {
	Bpm      int
	Position int
}

func (SyncTrack) header() string {
	return "SyncTrack"
}

func (Tempo) key() string {
	return "B"
}

func (t Tempo) value() string {
	return strconv.Itoa(t.Bpm * 1000)
}

func (t Tempo) location() string {
	return strconv.Itoa(t.Position)
}

func (TimeSignature) key() string {
	return "TS"
}

func (t TimeSignature) value() (out string) {
	out = strconv.Itoa(t.Numerator)

	if t.Denominator == 4 {
		return
	}

	//Must be greater than 0, less than or equal to 64, and a power of 2
	if t.Denominator <= 0 || t.Denominator > 64 || !(t.Denominator&(t.Denominator) == 0) {
		panic("invalid time signature denominator at position " + t.location())
	}

	out += " " + strconv.Itoa(t.Denominator)

	return
}

func (t TimeSignature) location() string {
	return strconv.Itoa(t.Position)
}
