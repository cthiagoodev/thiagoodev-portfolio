package common

import "errors"

var ErrNotFound = errors.New("not found")
var ErrInvalidData = errors.New("invalid data")
var ErrConnection = errors.New("connection error")
