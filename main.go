package main

import (
	"cronExpParser/cron"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("Not enough argument")
	}

	fmt.Println("Input : ", argsWithoutProg[0])
	exp, err := cron.Parse(argsWithoutProg[0])
	if err != nil {
		fmt.Println("ERROR : ", err)
		return
	}

	fmt.Println("Output : ")
	expToPrint := exp.ToString()
	for _, t := range expToPrint {
		fmt.Println(t)
	}
}
