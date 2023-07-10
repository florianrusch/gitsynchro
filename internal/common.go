package internal

import (
	"fmt"
	"os"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func NestedJoin[T any](joinChar string, objects []T, function func(T) string) string {
	var returnString string

	for i, obj := range objects {
		if i == 0 {
			returnString += function(obj)

			continue
		}

		returnString += fmt.Sprintf("%s%s", joinChar, function(obj))
	}

	return returnString
}
