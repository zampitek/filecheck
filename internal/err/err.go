package err

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidConfig = errors.New("invalid configuration")
	ErrMissingFile   = errors.New("missing file or directory")
	ErrParseError    = errors.New("failed to parse the file")
	ErrCLIError      = errors.New("invalid command")
)

type BaseError struct {
	Op  string
	Err error
}

func (e *BaseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Op, e.Err)
	}

	return e.Op
}

func (e *BaseError) Unwrap() error {
	return e.Err
}

func Wrap(op string, err error) error {
	if err == nil {
		return nil
	}

	return &BaseError{Op: op, Err: err}
}

func ExitWithError(msg string) {
	fmt.Fprintln(os.Stderr, "Error: "+msg)
	os.Exit(1)
}
