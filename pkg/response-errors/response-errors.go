package response_errors

import "errors"

var InvalidNumber = errors.New("invalid number")
var NotFound = errors.New("not found")
var CarrierNotFound = errors.New("carrier is not detected")
