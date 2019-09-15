package helpers

import (
	"fmt"
)

// ErrCheck checks if err is preset. If so, the program panics.
func ErrCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}
