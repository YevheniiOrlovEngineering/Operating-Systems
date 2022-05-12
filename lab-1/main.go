package main

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/algorithms"
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/utils"
)

// main drives SJF algorithm
func main() {
	pList := utils.GenerateProc()
	utils.PrintTableStdOut(pList, "Arriving Order")
	doneP := algorithms.SJF(pList)
	utils.PrintTableStdOut(doneP, "SJF Order")
}
