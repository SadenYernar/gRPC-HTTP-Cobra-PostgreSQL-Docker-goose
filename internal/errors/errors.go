package errors

import "fmt"

var (
	ErrNoRowsFound         = fmt.Errorf("NoRowsFound")
	ErrAlreadyRegistered   = fmt.Errorf("AlreadyRegistered")
	ErrUniquenessViolation = fmt.Errorf("UniquenessViolation")
)
