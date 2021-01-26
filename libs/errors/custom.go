package errors

import (
	"fmt"
)

type WithContext struct {
	Context  string
	Original error
}

func (e WithContext) Error() string {
	return fmt.Sprintf("[Context]: %s. [Original]: %s", e.Context, e.Original)
}

func (e WithContext) Unwrap() error {
	return e.Original
}
