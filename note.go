package ChartoGopher

type note struct {
	Time     int
	Colour   int
	Duration int
	//Hammer on / pull off
	Hopo bool
}

//Keys
const (
	GREEN  = 0
	RED    = 1
	YELLOW = 2
	BLUE   = 3
	ORANGE = 4
	OPEN   = 7
)
