package ChartoGopher

type Instrument string

//Instrument names
const (
	GUITAR       Instrument = "single"
	GUITARCOOP   Instrument = "DoubleGuitar"
	BASS         Instrument = "DoubleBass"
	RHYTHM       Instrument = "DoubleRhythm"
	DRUMS        Instrument = "Drums"
	KEYBOARD     Instrument = "Keyboard"
	GHLIVEGUITAR Instrument = "GHLGuitar"
	GHLIVEBASS   Instrument = "GHLBass"
)

type Difficulty string

const (
	EASY   Difficulty = "Easy"
	MEDIUM Difficulty = "Medium"
	HARD   Difficulty = "Hard"
	EXPERT Difficulty = "Expert"
)

type Track struct {
	Difficulty Difficulty
	Instrument Instrument
	Notes      []note
}

func (t Track) header() string {
	return string(t.Difficulty) + string(t.Instrument)
}

func CreateTrack(difficulty Difficulty, instrument Instrument) Track {
	return Track{
		Difficulty: difficulty,
		Instrument: instrument,
		Notes:      make([]note, 0),
	}
}

func (t *Track) AddNote(note note) {
	t.Notes = append(t.Notes, note)
}
