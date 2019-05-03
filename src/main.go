package main

import (
	"fmt"
	"os/exec"
	"time"
	"math"
	"os"
)

var logFile = "/.Timerwood.dat"

func scaler(unixSec int, secType string) int{
	timeType := map[string]int {
		"DAY":       86400,
		"HOUR":      3600,
		"MINUTE":    60,
	}
	if unixSec > timeType[secType] {
		per := math.Floor(float64(unixSec / timeType[secType]))
		return int(per)
	} else {
		return 0
	}
}

func outputTimeUsage(time map[string]int, timeUsageNano float64) string{
	if time["minute"] == 0 {
		fmt.Printf("Time Usage : %f Sec\n", timeUsageNano / 1e9)
		return fmt.Sprintf("%f Sec", timeUsageNano / 1e9)
	}

	var tips string = "Time Usage : "
	var timestring string = ""
	if time["day"] > 0 {
		tips += "%d Day "
		timestring += "%d Day "
	}

	if time["hour"] > 0 {
		tips += "%d Hour "
		timestring += "%d Hour "
	}

	if time["minute"] > 0 {
		tips += "%d Min "
		timestring += "%d Min "
	}

	if time["second"] >= 0 {
		tips += "%d Sec "
		timestring += "%d Sec "
	}

	tips += "\n"

	if time["day"] > 0 && time["hour"] > 0 && time["minute"] > 0 {
		fmt.Printf(tips, time["day"], time["hour"], time["minute"], time["second"])
		return fmt.Sprintf(timestring, time["day"], time["hour"], time["minute"], time["second"])
	}

	if time["hour"] > 0 && time["minute"] > 0 {
		fmt.Printf(tips, time["hour"], time["minute"], time["second"])
		return fmt.Sprintf(timestring, time["hour"], time["minute"], time["second"])
	}

	if time["minute"] > 0 {
		fmt.Printf(tips, time["minute"], time["second"])
		return fmt.Sprintf(timestring, time["minute"], time["second"])
	}
	return ""
}

func outputFormat(execError error) bool{
	fmt.Printf("#####################################\n")
	fmt.Printf("#          T I M E R W O O D        #\n")
	fmt.Printf("#####################################\n")
	fmt.Printf("\n")
	if execError != nil {
		fmt.Printf("[Fail] Progam Has Error : %s\n", execError)
		return false
	} else {
		fmt.Printf("[Success] ")
		return true
	}
}

func log(execContent []byte) {
	fileName := fmt.Sprintf("output_%d.log", time.Now().UnixNano())
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write(execContent)
		if err != nil {
			fmt.Printf("[Warning] Write Progam Ouput Error : %s\n", err)
		}
	}
}

func getTimerWoodLog(timestr string, execFile string) string{
	var log string

	timeLength := len(timestr)
	space := " "
	for i := 29 - timeLength; i > 0 ; i-- {
		space += " "
	}
	date := time.Now().Format("2006-01-02 15:04:05")
	log += "| " + date + " | " + timestr + space + "| " + execFile + "\n"
	return log
}

func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("[Warning] Write Timerwood.dat Error : %s\n", err)
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

func writeTimerWood(timestr string, execFile string) {
	fileName := "/.Timerwood.dat"

	_, err := os.Stat(fileName)

	if err == nil {
		// FILE EXIST : APPEND
		_ = appendToFile(fileName, getTimerWoodLog(timestr, execFile))
	} else {
		fmt.Printf("[Warning] Write Timerwood.dat Error : %s\n", err)
	}
}

func main()  {
	_, err := os.Stat(logFile)
	if err != nil {
		// 文件存在
		fmt.Printf("[Exit] Timerwood is not installed in your System\n")
		return
	}

	if len(os.Args) != 2 {
		fmt.Printf("[Exit] Please add your exec Command\n")
		return
	}

	args := os.Args[1]

	startTime := time.Now().UnixNano()
	execContent, execError := exec.Command("/bin/sh", "-c", args).Output()
	endTime := time.Now().UnixNano()

	timeUsage := float64((endTime - startTime) / 1e9)
	timeUsageNano := float64(endTime - startTime)

	timeDay := scaler(int(timeUsage), "DAY")
	timeHour := scaler(int(timeUsage) - timeDay * 86400, "HOUR")
	timerMinute := scaler(int(timeUsage) - timeDay * 86400 - timeHour * 3600, "MINUTE")
	timerSecond := int(timeUsage) - timeDay * 86400 - timeHour * 3600 - timerMinute * 60

	timeTypeParam := map[string]int {
		"day": timeDay,
		"hour": timeHour,
		"minute": timerMinute,
		"second": timerSecond,
	}
	ifSucc := outputFormat(execError)

	if ifSucc {
		/* 打印执行时间 */
		timerWoodLogString := outputTimeUsage(timeTypeParam, timeUsageNano)

		/* 打印输出日志 */
		log(execContent)

		/* 输出数据文件 */
		writeTimerWood(timerWoodLogString, string(args))
	}
	return
}
