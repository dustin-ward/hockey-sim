package sim

// League Enum
type League uint8

const (
	NHL League = iota
	AHL
)

// Full name of player
type Name struct {
	First string
	Last  string
}

// Player shoots/catches left/right
type Handedness uint8

const (
	LEFT Handedness = iota
	RIGHT
)

// Main player datatype
type Player struct {
	Name
	Age int
	Position

	Stats

	url string
}
