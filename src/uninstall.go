package main

import (
	"os"
	"fmt"
)

var logFile = "/.Timerwood.dat"

func main() {
	fmt.Printf("#####################################\n")
	fmt.Printf("#          T I M E R W O O D        #\n")
	fmt.Printf("#####################################\n")
	fmt.Printf("\n")

	err := os.Remove(logFile)

	if err != nil {
		fmt.Printf("[Fail] TIMERWOOD is not Uninstalled\n")
	} else {
		fmt.Printf("[Success] Uninstall Successfully , WELCOME to use TIMERWOOD Again\n")
	}
}