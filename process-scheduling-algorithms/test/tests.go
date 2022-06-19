package test

import (
	"errors"
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/algorithms"
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"
	"log"
)

const testRunNum = 1000

// ValidateAlgorithms drives tests
func ValidateAlgorithms(stdErr *log.Logger) (bool, error) {
	validated := true

	ok1, err1 := testArrBurst()
	ok2, err2 := testFinStartSJF()
	ok3, err3 := testFinStartSRTF()
	ok4, err4 := testArrStartSJF()
	ok5, err5 := testArrStartSRTF()

	if !ok1 && err1 != nil {
		stdErr.Println(err1)
		validated = false
	}
	if !ok2 && err2 != nil {
		stdErr.Println(err2)
		validated = false
	}
	if !ok3 && err3 != nil {
		stdErr.Println(err3)
		validated = false
	}
	if !ok4 && err4 != nil {
		stdErr.Println(err4)
		validated = false
	}
	if !ok5 && err5 != nil {
		stdErr.Println(err5)
		validated = false
	}
	if !validated {
		return false, errors.New("validation tests failed. Refer to logs for details")
	} else {
		return true, nil
	}
}

func testArrBurst() (bool, error) {
	//bugC := 0
	for j := 0; j < testRunNum; j++ {
		var SjfPool []process.Process
		var srtfPool []process.Process

		pList := process.GenerateProc()

		SjfPool = append(SjfPool, pList...)
		srtfPool = append(srtfPool, pList...)

		pSJF, _, _ := algorithms.SJF(SjfPool)
		pSRTF, _, _ := algorithms.SRTF(srtfPool)

		if len(pSRTF) == len(pSJF) {
			for i := range pSJF {
				idx := process.GetIdxById(pSRTF, pSJF[i].Id)
				if pSJF[i].At != pSRTF[idx].At || pSJF[i].Bt != pSRTF[idx].Bt {
					/*
						fmt.Println("[ERROR]", pSRTF[idx].Id)
						bugC++
					*/
					return false, errors.New("same processes have different initial specs")
				}
			}
		} else {
			return false, errors.New("algorithms process pools have different size")
		}
	}
	//fmt.Println(bugC)
	return true, nil
}

func testFinStartSJF() (bool, error) {
	//bugC := 0
	for j := 0; j < testRunNum; j++ {
		var SjfPool []process.Process

		pList := process.GenerateProc()
		SjfPool = append(SjfPool, pList...)
		pSJF, _, _ := algorithms.SJF(SjfPool)

		for i := 0; i < len(pSJF)-1; i++ {
			if pSJF[i].Ft != pSJF[i+1].St && pSJF[i+1].At <= pSJF[i].Ft {
				/*
					fmt.Println(pSJF[i].Id, pSJF[i+1].Id, "\tERROR")
					process.PrintProcessTableStdOut(pSJF, "SJF order")
					var debugList []process.Process
					debugList = append(debugList, pList...)
					algorithms.SJF(debugList)
					bugC++
				*/
				return false, errors.New("|SJF| " +
					"Arrived processes have gap in execution flow")
			}
		}
	}
	//fmt.Println(bugC)
	return true, nil
}

func testFinStartSRTF() (bool, error) {
	//bugC := 0
	for j := 0; j < testRunNum; j++ {
		var srtfPool []process.Process

		pList := process.GenerateProc()
		srtfPool = append(srtfPool, pList...)
		pSRTF, _, _ := algorithms.SRTF(srtfPool)

		for i := 0; i < len(pSRTF)-1; i++ {
			if pSRTF[i].Ft != pSRTF[i+1].St && pSRTF[i+1].At <= pSRTF[i].Ft {
				/*
					fmt.Println(pSRTF[i].Id, pSRTF[i+1].Id, "\tERROR")
					process.PrintProcessTableStdOut(pSRTF, "SRTF order")
					var debugList []process.Process
					debugList = append(debugList, pList...)
					algorithms.SRTF(debugList)
					bugC++
				*/
				return false, errors.New("|SRTF| " +
					"Arrived processes have gap in execution flow")
			}
		}
	}
	//fmt.Println(bugC)
	return true, nil
}

func testArrStartSJF() (bool, error) {
	//bugC := 0
	for j := 0; j < testRunNum; j++ {
		var sjfPool []process.Process

		pList := process.GenerateProc()
		sjfPool = append(sjfPool, pList...)
		pSJF, _, _ := algorithms.SJF(sjfPool)

		for i := range pSJF {
			if pSJF[i].At > pSJF[i].St {
				/*
					fmt.Println(pSJF[i].Id, "\tERROR")
					process.PrintProcessTableStdOut(pSJF, "SJF order")
					var debugList []process.Process
					debugList = append(debugList, pList...)
					algorithms.SJF(debugList)
					bugC++
				*/
				return false, errors.New("|SJF| " +
					"Process started earlier than arrived")
			}
		}
	}
	//fmt.Println(bugC)
	return true, nil
}

func testArrStartSRTF() (bool, error) {
	//bugC := 0
	for j := 0; j < testRunNum; j++ {
		var srtfPool []process.Process

		pList := process.GenerateProc()
		srtfPool = append(srtfPool, pList...)
		pSRTF, _, _ := algorithms.SRTF(srtfPool)

		for i := range pSRTF {
			if pSRTF[i].At > pSRTF[i].St {
				/*
					fmt.Println(pSRTF[i].Id, "\tERROR")
					process.PrintProcessTableStdOut(pSRTF, "SRTF order")
					var debugList []process.Process
					debugList = append(debugList, pList...)
					algorithms.SRTF(debugList)
					bugC++
				*/
				return false, errors.New("|SRTF| " +
					"Process started earlier than arrived")
			}
		}
	}
	//fmt.Println(bugC)
	return true, nil
}
