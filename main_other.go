// +build !windows

package main

import (
	"log"
)

func RunWindows(cfgPath string) {
	log.Fatal("RunWindows should never be called on a platform other than Windows!")
}
