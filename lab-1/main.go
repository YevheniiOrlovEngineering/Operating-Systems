package main

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/utils"
)

// main drives SJF algorithm
func main() {
	pArr := utils.GenerateProc()
	utils.PrintProcess(pArr)
}
