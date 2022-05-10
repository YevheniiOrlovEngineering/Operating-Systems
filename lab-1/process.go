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

// GenerateProc generates process list with all specs
func GenerateProc() []Process {
	const (
		PNumMax     int = 30
		PArrTimeMax int = 7
		PBurTimeMax int = 10
	)
	rand.Seed(time.Now().UnixNano())
	pArr := make([]Process, rand.Intn(PNumMax-1)+1)

	for i := 0; i < len(pArr); i++ {
		pid := "P" + strconv.Itoa(i)
		aTime := rand.Intn(PArrTimeMax)
		bTime := rand.Intn(PBurTimeMax-1) + 1

		pArr[i] = Process{
			id: pid,
			at: aTime,
			bt: bTime,
		}
	}
	sortArrivalBurst(pArr)
	return pArr
}

// PrintProcess prints process list in StdOut
func PrintProcess(pArr []Process) {
	for _, p := range pArr {
		fmt.Println(p.id, ":", "Arrival Time", p.at, "\tBurst Time,", p.bt)
	}
}

// Remove pops out i_th element from slice s
func Remove(s []Process, i int) []Process {
	return append(s[:i], s[i+1:]...)
}

func GetIdxById(pArr []Process, id string) int {
	for i := range pArr {
		if pArr[i].id == id {
			return i
		}
	}
	return -1
}

// sortArrival sorts slice by arrival time
func sortArrival(pArr []Process) {
	sort.SliceStable(pArr, func(i, j int) bool {
		return pArr[i].at < pArr[j].at
	})
}

// sortBurst sorts slice by burst time
func sortBurst(pArr []Process) {
	sort.SliceStable(pArr, func(i, j int) bool {
		return pArr[i].bt < pArr[j].bt
	})
}

// sortArrivalBurst sorts initially by arrival time and then by burst time
func sortArrivalBurst(pArr []Process) {
	var atList []int
	sortArrival(pArr)
	for i := range pArr {
		atList = append(atList, pArr[i].at)
	}
	for k, v := range GetDuplicates(atList) {
		i := sort.SearchInts(atList, k)
		pToSort := pArr[i : i+v]
		sortBurst(pToSort)
	}
}

// SortBurstArrival sorts initially by burst time and then by arrival time
func SortBurstArrival(pArr []Process) {
	var btList []int
	sortBurst(pArr)
	for i := range pArr {
		btList = append(btList, pArr[i].bt)
	}
	for k, v := range GetDuplicates(btList) {
		i := sort.SearchInts(btList, k)
		pToSort := pArr[i : i+v]
		sortArrival(pToSort)
	}
}

// Task emulates some task for process
func Task(d int) {
	//time.Sleep(time.Duration(d) * time.Second)
}

// GetDuplicates produces map with numbers which occurs more than once
func GetDuplicates(n []int) map[int]int {
	countMap := make(map[int]int)
	for _, t := range n {
		if _, ok := countMap[t]; !ok {
			countMap[t] = 1
		} else {
			countMap[t] += 1
		}
	}
	for k, v := range countMap {
		if v < 2 {
			delete(countMap, k)
		}
	}
	return countMap
}
