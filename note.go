package ChartoGopher

type note struct {
	Time     int
	Colour   Button
	Duration int
	NoteChar starPowerChar
}

type Button int

//Keys
const (
	BTN_GREEN  Button = 0
	BTN_RED    Button = 1
	BTN_YELLOW Button = 2
	BTN_BLUE   Button = 3
	BTN_ORANGE Button = 4
	BTN_OPEN   Button = 7

	SP_BTN           Button = 2
	CANCEL_HAMMER_ON Button = 5
	TAP              Button = 6
)

type starPowerChar string

const (
	SP_FALSE starPowerChar = "N"
	SP_TRUE  starPowerChar = "S"
)
