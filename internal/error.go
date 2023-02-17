package internal

import "errors"

var errEmptyUsername = errors.New("empty username")
var errEmptyMemberShip = errors.New("empty membership")
var errEmptyId = errors.New("empty ID")
var errNotApplyMemberShip = errors.New("bad request membership type")
var errAlreadyExistUsername = errors.New("exist already username")
var errNotFoundId = errors.New("not found id")
var errNotFoundException = errors.New("not found info")
