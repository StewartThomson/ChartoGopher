package ChartoGopher

type Instrument string

//Instrument names
const (
	INSTR_GUITAR       Instrument = "single"
	INSTR_GUITARCOOP   Instrument = "DoubleGuitar"
	INSTR_BASS         Instrument = "DoubleBass"
	INSTR_RHYTHM       Instrument = "DoubleRhythm"
	INSTR_DRUMS        Instrument = "Drums"
	INSTR_KEYBOARD     Instrument = "Keyboard"
	INSTR_GHLIVEGUITAR Instrument = "GHLGuitar"
	INSTR_GHLIVEBASS   Instrument = "GHLBass"
)

type Difficulty string

const (
	DIFF_EASY   Difficulty = "Easy"
	DIFF_MEDIUM Difficulty = "Medium"
	DIFF_HARD   Difficulty = "Hard"
	DIFF_EXPERT Difficulty = "Expert"
)

type Track struct {
	Difficulty Difficulty
	Instrument Instrument
	Notes      []note
}

func (t Track) Header() string {
	return string(t.Difficulty) + string(t.Instrument)
}

func CreateTrack(difficulty Difficulty, instrument Instrument) *Track {
	return &Track{
		Difficulty: difficulty,
		Instrument: instrument,
		Notes:      make([]note, 0),
	}
}

func (t *Track) AddNote(time int, colour int, duration int, hopo bool) {
	t.Notes = append(t.Notes, note{
		Time:     time,
		Colour:   colour,
		Duration: duration,
		Hopo:     hopo,
	})
}
