package main

import (
	"fmt"
	"os"
	"time"
)

func WriteLineInFile(file *os.File, str string) {
	file.WriteString(fmt.Sprintln(str))

}

func GetFileNameWithDatePrefix(filename string) string {
	current_time := time.Now()

	return fmt.Sprintf("%d%02d%02dT%02d%02d%02d_%s.csv",
		current_time.Year(), current_time.Month(), current_time.Day(),
		current_time.Hour(), current_time.Minute(), current_time.Second(), filename)
}

func GetTimer(name string) func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		var elapsed = time.Since(start)
		fmt.Printf("%s took %v\n", name, elapsed)
		return elapsed
	}
}
