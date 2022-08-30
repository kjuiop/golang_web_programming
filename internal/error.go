package internal

import "errors"

var (
	errEmptyUsername        = errors.New("empty username")
	errEmptyMemberShip      = errors.New("empty membership")
	errEmptyId              = errors.New("empty ID")
	errNotApplyMemberShip   = errors.New("bad request membership type")
	errAlreadyExistUsername = errors.New("exist already username")
	errNotFoundId           = errors.New("not found id")
	errNotFoundException    = errors.New("not found info")
)
