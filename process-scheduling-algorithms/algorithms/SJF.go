package algorithms

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"
)

// SJF produces process executing order according to SJF algorithm
func SJF(pList []process.Process) ([]process.Process, int, int) {
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

			p := wQ[0]
			SetProcessSpecs(&p, ct)
			dQ = append(dQ, p)
			wQ = wQ[1:]
			ct = p.Ft
		} else {
			ct = pList[0].At
		}
	}
	dQ = ComputeProcesses(wQ, dQ, ct)
	avgWt, avgTat := EvalAvgStats(dQ)
	return dQ, avgWt, avgTat
}
