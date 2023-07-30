package internal

import (
	"fmt"
)

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
