package main

import (
	"fmt"
)

var (
	version   = "v1.0.0"
	buildTime string
)

// displayVersion prints out the version number and build time of the program before exiting.
func displayVersion() {
	fmt.Printf("Version:\t%s\n", version)
	fmt.Printf("Build time:\t%s\n", buildTime)
}
