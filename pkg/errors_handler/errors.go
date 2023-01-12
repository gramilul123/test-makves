package errors_handler

import (
	"fmt"
)

var (
	ErrReadFile        = fmt.Errorf("error read CSV file")
	ErrHttpConnection  = fmt.Errorf("cant connect with server")
	ErrInternalService = fmt.Errorf("internal service error")
)
