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
	GREEN  Button = 0
	RED    Button = 1
	YELLOW Button = 2
	BLUE   Button = 3
	ORANGE Button = 4
	OPEN   Button = 7

	CANCEL_HAMMER_ON Button = 5
	HAMMER_ON_BTN    Button = 2
)

type starPowerChar string

const (
	SP_FALSE starPowerChar = "N"
	SP_TRUE  starPowerChar = "S"
)
