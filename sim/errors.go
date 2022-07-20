package sim

import "errors"

// Inserting players into lineup
var INSERT_PLAYER_INVALID_LINE error = errors.New("error inserting player into lineup. invalid values")
var INSERT_PLAYER_INVALID_POSITION error = errors.New("error inserting player into lineup. invalid values")
