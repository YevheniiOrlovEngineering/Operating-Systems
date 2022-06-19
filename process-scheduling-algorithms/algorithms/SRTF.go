package algorithms

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"
)

// SRTF produces process executing order according to SRTF algorithm
func SRTF(pList []process.Process) ([]process.Process, int, int) {
	var wQ []process.Process
	var dQ []process.Process
	ct := 0

	process.SortArrivalBurst(pList)

	for len(pList) != 0 {
		if pList[0].At <= ct || len(wQ) != 0 {
			pNotArrIdx := process.GetIdxByAt(pList, ct) + 1
			for j := 0; j < pNotArrIdx; j++ {
				wQ = append(wQ, pList[j])
			}
			process.SortBurstArrival(wQ)
			pList = process.RemoveProcesses(pList, 0, pNotArrIdx-1)
			pToDo := wQ[0]

			if pLastIdx := process.GetIdxByAt(pList, ct+pToDo.Bt-1); pLastIdx != -1 {
				for i := 0; i != pLastIdx+1; i++ {
					p := pList[i]
					if p.Bt < pToDo.Bt {
						pToDo = p
						pLastIdx = process.GetIdxByAt(pList, p.At+p.Bt-1)
					}
				}
			}
			if pToDo == wQ[0] {
				wQ = wQ[1:]
			} else {
				pToDoIdx := process.GetIdxById(pList, pToDo.Id)
				pList = process.RemoveProcesses(pList, pToDoIdx, pToDoIdx)
			}
			if len(dQ) != 0 && pToDo.At <= dQ[len(dQ)-1].Ft {
				SetProcessSpecs(&pToDo, dQ[len(dQ)-1].Ft)
			} else {
				SetProcessSpecs(&pToDo, pToDo.At)
			}
			dQ = append(dQ, pToDo)
			ct = pToDo.Ft
		} else {
			ct = pList[0].At
		}
	}
	dQ = ComputeProcesses(wQ, dQ, ct)
	avgWt, avgTat := EvalAvgStats(dQ)
	return dQ, avgWt, avgTat
}
