package main

import (
	"os"
	"fmt"
	"os/exec"
)

var logFile = "/.Timerwood.dat"

func main() {
	fmt.Printf("#####################################\n")
	fmt.Printf("#          T I M E R W O O D        #\n")
	fmt.Printf("#####################################\n")
	fmt.Printf("\n")

	tag := createLog()
	exec.Command("/bin/sh", "-c", "cp ")
	if tag == 1 {
		fmt.Printf("[Success] Install Successfully , WELCOME to use TIMERWOOD\n")
	} else if tag == 3 {
		fmt.Printf("[Fail] Install Failed , Please Check the error above\n")
	}
}

func createLog() int{
	var log string
	log += "#####################################\n"
	log += "#          T I M E R W O O D        #\n"
	log += "#####################################\n"
	log += "\n"
	log += "| Date                | Time Usage                    | Exec File \n"

	_, err := os.Stat(logFile)
	if err == nil {
		// 文件存在
		fmt.Printf("[Warning] Timerwood is installed in your System\n")
		return 2
	} else {
		f, err := os.Create(logFile)
		defer f.Close()
		if err != nil {
			fmt.Printf("[Warning] Write Timerwood.dat Error : %s\n", err)
			return 3
		} else {
			_, err = f.Write([]byte(log))
			if err != nil {
				fmt.Printf("[Warning] Write Timerwood.dat Error : %s\n", err)
				return 3
			}
		}
	}
	return 1
}