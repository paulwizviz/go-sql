package dbop

import (
	"errors"
)

var (
	ErrPrepareStatement = errors.New(("unable to prepare statement"))
	ErrTableCreation    = errors.New("unable to create table")
)
