package algorithms

import "github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"

// SetProcessSpecs computes single process specs
func SetProcessSpecs(p *process.Process, t int) {
	p.St = t
	p.Ft = t + p.Bt
	p.Tat = p.Ft - p.At
	p.Wt = p.Tat - p.Bt
}

// ComputeProcesses computes specs for a pool of processes
func ComputeProcesses(wQ []process.Process, dQ []process.Process, t int) []process.Process {
	for _, p := range wQ {
		SetProcessSpecs(&p, t)
		dQ = append(dQ, p)
		t += p.Bt
	}
	return dQ
}

// EvalAvgStats computes average value of Waiting time and Turn Around time
func EvalAvgStats(pList []process.Process) (int, int) {
	avgWT, avgTAT := 0, 0
	for _, p := range pList {
		avgWT += p.Wt
		avgTAT += p.Tat
	}
	return avgWT / len(pList), avgTAT / len(pList)
}
