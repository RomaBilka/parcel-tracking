package response_errors

import "errors"

var InvalidNumber = errors.New("invalid number")
var NotFound = errors.New("not found")
var SomethingIsWrong = errors.New("something is wrong")
