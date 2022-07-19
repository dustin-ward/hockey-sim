package main

import "errors"

// Status of current contract
type SigningStatus uint8

const (
	UNSIGNED_UFA SigningStatus = iota
	UNSIGNED_RFA
	SIGNED_UFA
	SIGNED_RFA
)

// Contract datatype
type Contract struct {
	Term         uint8
	Two_Way      bool
	Values_NHL   []float32
	Values_AHL   []float32
	Current_year uint16
}

// Constructor for Contract
func NewContract(term uint8, two_Way bool, nhl_values, ahl_values []float32) *Contract {
	c := new(Contract)
	c.Term = term
	c.Current_year = 0
	copy(c.Values_NHL, nhl_values)
	if two_Way {
		c.Two_Way = true
		copy(c.Values_AHL, ahl_values)
	} else {
		ahl_values = nil
	}
	return c
}

// Calculate AAV of contract
func (c *Contract) AAV(league League) (float32, error) {
	var sum float32 = 0.0
	if league == NHL {
		for _, v := range c.Values_NHL {
			sum += v
		}
	} else if c.Two_Way {
		for _, v := range c.Values_AHL {
			sum += v
		}
	} else {
		return -1, errors.New("error calculating AAV of contract")
	}
	return sum / float32(c.Term), nil
}
