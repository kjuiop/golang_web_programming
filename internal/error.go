package internal

import "errors"

var errEmptyUsername = errors.New("empty username")
var errEmptyMemberShip = errors.New("empty membership")
var errNotApplyMemberShip = errors.New("bad request membership type")
var errAlreadyExistUsername = errors.New("exist already username")
