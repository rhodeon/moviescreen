// Package prettylog provides functions for logging info and errors in distinguished colours.
package prettylog

import (
	"fmt"
	"github.com/rhodeon/prettylog/colors"
	"log"
	"os"
)

// infoLogger is a logger which prints message to standard output.
// Also displays the date and time.
var infoLogger = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

// colorizeInfo prints the info log and resets text colour afterwards.
func colorizeInfo(logInfo func()) {
	fmt.Fprint(os.Stdout, colors.Yellow)
	logInfo()
	fmt.Fprint(os.Stdout, colors.Reset)
}

// Info is equivalent to fmt.Print.
func Info(info ...interface{}) {
	colorizeInfo(func() {
		infoLogger.Print(info...)
	})
}

// InfoF is equivalent to fmt.Printf.
func InfoF(format string, info ...interface{}) {
	colorizeInfo(func() {
		infoLogger.Printf(format, info...)
	})
}

// InfoLn is equivalent to fmt.Println.
func InfoLn(info ...interface{}) {
	colorizeInfo(func() {
		infoLogger.Println(info...)
	})
}
