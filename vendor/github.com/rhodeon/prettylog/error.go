package prettylog

import (
	"fmt"
	"github.com/rhodeon/prettylog/colors"
	"log"
	"os"
)

// errorLogger is a Logger which prints message to standard error.
// Also displays the date, time and relative path of the affected file.
var errorLogger = log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

// colorizeError prints the error log and resets text colour afterwards.
func colorizeError(logError func()) {
	fmt.Fprint(os.Stderr, colors.Red)
	logError()
	fmt.Fprint(os.Stderr, colors.Reset)
}

// Error is equivalent to fmt.Print.
func Error(error ...interface{}) {
	colorizeError(func() {
		errorLogger.Print(error...)
	})
}

// ErrorF is equivalent to fmt.PrintF.
func ErrorF(format string, error ...interface{}) {
	colorizeError(func() {
		errorLogger.Printf(format, error...)
	})
}

// ErrorLn is equivalent to fmt.PrintLn.
func ErrorLn(error ...interface{}) {
	colorizeError(func() {
		errorLogger.Println(error...)
	})
}

// FatalError logs the error then exits the program.
func FatalError(error ...interface{}) {
	colorizeError(func() {
		errorLogger.Fatal(error...)
	})
}
