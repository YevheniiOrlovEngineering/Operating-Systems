package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Process struct {
	id string
	// Arrival time
	at int
	// Burst time
	bt int
	// Completion time
	ct int
	// Start time
	st int
	// Finish time
	ft int
	// Waiting time
	wt int
	// Turn Around time
	tat int
}

// return slice of range {min-max} inclusively
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func GenerateProc() []Process {
	const (
		PNumMax     int = 10
		PArrTimeMax int = 20
		PBurTimeMax int = 10
	)

	rand.Seed(time.Now().UnixNano())

	pArr := make([]Process, rand.Intn(PNumMax-1)+1)
	aTimes := makeRange(0, PArrTimeMax)

	for i := 0; i < len(pArr); i++ {
		pid := "P" + strconv.Itoa(i)
		aTime := aTimes[rand.Intn(len(aTimes))]
		bTime := rand.Intn(PBurTimeMax)

		pArr[i] = Process{
			id: pid,
			at: aTime,
			bt: bTime,
		}
		aTimes = remove(aTimes, sort.IntSlice(aTimes).Search(aTime))
	}

	sortProcesses(pArr)
	return pArr
}

func PrintProcess(pArr []Process) {
	for _, v := range pArr {
		fmt.Println(v.id, ":", "Arrival Time", v.at, "\tBurst Time,", v.bt)
	}
}

func Task(d int) {
	print("Working ...")
	print("Is cost", d)
	//time.Sleep(time.Duration(d) * time.Second)
}

func sortProcesses(pArr []Process) {
	sort.Slice(pArr, func(i, j int) bool {
		return pArr[i].at < pArr[j].at
	})
}
