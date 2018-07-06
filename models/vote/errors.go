package vote

import (
	"errors"
)

var ErrVisitorNotFinishedOrWrong = errors.New("visitor not finished all question or wrong")
var ErrUndefine = errors.New("undefine")
