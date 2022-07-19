package main

const PERIOD_LENGTH = 20 * 60

// Enumeration for all possible game states
type GameState uint8

const (
	PRE_GAME GameState = iota
	FIRST_PERIOD
	SECOND_PERIOD
	THIRD_PERIOD
	POST_GAME
)

// Representation of the official in-game clock
type Clock struct {
	Current_State     GameState
	Seconds_Remaining uint32
}

// Constructor for Clock type
func NewClock() *Clock {
	c := new(Clock)
	c.Current_State = PRE_GAME
	c.Seconds_Remaining = PERIOD_LENGTH
	return c
}

func (c *Clock) reset() bool {
	c.Current_State++
	if c.Current_State == POST_GAME {
		return false
	}
	c.Seconds_Remaining = PERIOD_LENGTH
	return true
}

func (c *Clock) Tick() bool {
	if c.Seconds_Remaining == 0 {
		return c.reset()
	}

	if c.Current_State == PRE_GAME {
		return c.reset()
	}

	c.Seconds_Remaining--
	return true
}
