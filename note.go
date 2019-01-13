package ChartoGopher

type note struct {
	Time     int
	Colour   Button
	Duration int
	//Hammer on / pull off
	Hopo bool
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
)
