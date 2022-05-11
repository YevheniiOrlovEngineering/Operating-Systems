package utils

import (
	"github.com/YevheniiOrlovEngineering/Operating-Systems/lab-1/process"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// GenerateProc generates process list with all specs
func GenerateProc() []process.Process {
	rand.Seed(time.Now().UnixNano())
	pList := make([]process.Process, rand.Intn(process.PNumMax-1)+1)

	for i := 0; i < len(pList); i++ {
		id := "P" + strconv.Itoa(i)
		at := rand.Intn(process.PArrTimeMax)
		bt := rand.Intn(process.PBurTimeMax-1) + 1

		pList[i] = process.Process{
			Id: id,
			At: at,
			Bt: bt,
		}
	}
	sortArrivalBurst(pList)
	return pList
}

// RemoveProcesses pops out processes from i_th to j_th
func RemoveProcesses(s []process.Process, i int, j int) []process.Process {
	return append(s[:i], s[j+1:]...)
}

// RemoveInts pops out integers from i_th to j_th
func RemoveInts(s []int, i int, j int) []int {
	return append(s[:i], s[j+1:]...)
}

// GetIdxById gets process index by id
func GetIdxById(pList []process.Process, id string) int {
	for i := range pList {
		if pList[i].Id == id {
			return i
		}
	}
	return -1
}

// GetIdxByAt gets process index by arrival time
func GetIdxByAt(pList []process.Process, at int) int {
	return sort.Search(len(pList), func(i int) bool {
		return pList[i].At > at
	})
}

// GetIdxByBt gets process index by burst time
func GetIdxByBt(pList []process.Process, bt int) int {
	return sort.Search(len(pList), func(i int) bool {
		return pList[i].Bt == bt
	})
}

// sortArrival sorts slice by arrival time
func sortArrival(pList []process.Process) {
	sort.SliceStable(pList, func(i, j int) bool {
		return pList[i].At < pList[j].At
	})
}

// sortBurst sorts slice by burst time
func sortBurst(pList []process.Process) {
	sort.SliceStable(pList, func(i, j int) bool {
		return pList[i].Bt < pList[j].Bt
	})
}

// sortArrivalBurst sorts initially by arrival time and then by burst time
func sortArrivalBurst(pList []process.Process) {
	var atList []int
	sortArrival(pList)
	for i := range pList {
		atList = append(atList, pList[i].At)
	}
	for k, v := range GetDuplicates(atList) {
		i := sort.SearchInts(atList, k)
		pToSort := pList[i : i+v]
		sortBurst(pToSort)
	}
}

// SortBurstArrival sorts initially by burst time and then by arrival time
func SortBurstArrival(pList []process.Process) {
	var btList []int
	sortBurst(pList)
	for i := range pList {
		btList = append(btList, pList[i].Bt)
	}
	for k, v := range GetDuplicates(btList) {
		i := sort.SearchInts(btList, k)
		pToSort := pList[i : i+v]
		sortArrival(pToSort)
	}
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
