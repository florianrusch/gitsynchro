package log

import "fmt"

// Infof should be used to describe the example commands that are about to run.
func Debugf(format string, args ...interface{}) {
	fmt.Printf("DEBU: %s\n", fmt.Sprintf(format, args...))
}

// Infof should be used to describe the example commands that are about to run.
func Infof(format string, args ...interface{}) {
	fmt.Printf("INFO: %s\n", fmt.Sprintf(format, args...))
}

// Warningf should be used to display a warning.
func Warningf(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1mWARN: %s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Errorf should be used to display errors.
func Errorf(err error) {
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}
