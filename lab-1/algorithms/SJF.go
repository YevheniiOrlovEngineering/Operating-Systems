package algorithms

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/utils"
)

// setProcessStats computes process specs after its completion
func setProcessStats(p *process.Process, t int) {
	p.St = t
	p.Ft = t + p.Bt
	p.Wt = p.At + t
	p.Tat = p.At + p.Ft
}

// SJF produces process executing order according to SJF algorithm
func SJF(pList []process.Process) []process.Process {
	var wQ []process.Process
	var dQ []process.Process
	pNum, ct := len(pList), 0

	for len(dQ) != pNum {
		if len(pList) > 0 && pList[0].At <= ct {
			notA := utils.GetIdxByAt(pList, ct)
			for j := 0; j < notA; j++ {
				wQ = append(wQ, pList[j])
			}
			utils.SortBurstArrival(wQ)
			pList = utils.RemoveProcesses(pList, 0, notA-1)
		} else if len(pList) != 0 {
			ct++
			continue
		}
		p := wQ[0]
		setProcessStats(&p, ct)
		dQ = append(dQ, p)
		wQ = wQ[1:]
		ct += p.Bt
	}
	return dQ
}
