package main

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/algorithms"
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/test"
	"log"
	"os"
)

// setLogEnv sets logger config
func setLogEnv() (*log.Logger, *log.Logger) {
	if err := os.RemoveAll("./logs"); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("./logs", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	logFile, err1 := os.Create("./logs/out.log")
	errFile, err2 := os.Create("./logs/err.log")

	if err1 != nil {
		log.Fatal(err1)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

	errorLogger := log.New(errFile, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(logFile, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLogger, errorLogger
}

// main drives SJF & SRTF algorithms
func main() {
	stdOut, stdErr := setLogEnv()
	if ok, err := test.ValidateAlgorithms(stdErr); ok && err == nil {

		var SjfPool []process.Process
		var srtfPool []process.Process

		pList := process.GenerateProc()
		SjfPool = append(SjfPool, pList...)
		srtfPool = append(srtfPool, pList...)

		pSJF, avgWt1, avgTat1 := algorithms.SJF(SjfPool)
		pSRTF, avgWt2, avgTat2 := algorithms.SRTF(srtfPool)

		process.PrintProcessTableStdOut(pList, "Process Pool", stdOut)
		process.PrintProcessTableStdOut(pSJF, "SJF Order", stdOut, avgWt1, avgTat1)
		process.PrintProcessTableStdOut(pSRTF, "SRTF Order", stdOut, avgWt2, avgTat2)
	} else {
		stdOut.Println(err)
	}
}
