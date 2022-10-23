package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// Output system performance
func checkPerformance() {
	// CPU Performance
	ps, _ := cpu.Percent(100*time.Millisecond, false)
	cpuPercent := int(ps[0])
	// Memory Performance
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	heapMem := int(toKb(ms.HeapAlloc))
	sysMem := int(toKb(ms.Sys))
	// Counting Goroutines
	goroutineCount := runtime.NumGoroutine()
	// Log output of system performance
	fmt.Printf("CPU : %d percent, HeapMemory : %d KB, SystemMemory : %d KB, GoroutineCount : %d \n", cpuPercent, heapMem, sysMem, goroutineCount)
}

func toKb(bytes uint64) uint64 {
	return bytes / 1024
}

// Prime Number detection function
func checkPrime(primeCount int) bool {
	for i := 2; i <= primeCount/2; i++ {
		if primeCount%i == 0 || primeCount == 0 || primeCount == 1 {
			return false
		}
	}
	return true
}

// Function to find one prime number greater than or equal to a certain number
func primeRoutine(targetLine int, primeChan chan int) {
	fmt.Printf("%d Line Target Goroutine Wake UP \n", targetLine)
	primeCount := 0
	for {
		primeCount++
		if checkPrime(primeCount) && targetLine < primeCount {
			checkPerformance()
			primeChan <- primeCount
			return
		}
	}
}

func main() {

	primeChan1 := make(chan int, 1024)
	primeChan2 := make(chan int, 1024)
	primeChan3 := make(chan int, 1024)
	primeChan4 := make(chan int, 1024)
	primeChan5 := make(chan int, 1024)

	go primeRoutine(100000, primeChan1)
	go primeRoutine(150000, primeChan2)
	go primeRoutine(200000, primeChan3)
	go primeRoutine(250000, primeChan4)
	go primeRoutine(300000, primeChan5)

STOP:
	for {
		select {
		case primeResult1 := <-primeChan1:
			fmt.Printf("PrimeAnswer1 : %d \n", primeResult1)
		case primeResult2 := <-primeChan2:
			fmt.Printf("PrimeAnswer2 : %d \n", primeResult2)
		case primeResult3 := <-primeChan3:
			fmt.Printf("PrimeAnswer3 : %d \n", primeResult3)
		case primeResult4 := <-primeChan4:
			fmt.Printf("PrimeAnswer4 : %d \n", primeResult4)
		case primeResult5 := <-primeChan5:
			fmt.Printf("PrimeAnswer5 : %d \n", primeResult5)
			break STOP
		default:
		}
	}

	checkPerformance()
}

// ==========================
//       Output Sample
// ==========================
// go build -o SystemPerformanceCheck SystemPerformanceCheck.go
// ./SystemPerformanceCheck
// 100000 Line Target Goroutine Wake UP
// 200000 Line Target Goroutine Wake UP
// 250000 Line Target Goroutine Wake UP
// 150000 Line Target Goroutine Wake UP
// 300000 Line Target Goroutine Wake UP
// CPU : 63 percent, HeapMemory : 180 KB, SystemMemory : 9043 KB, GoroutineCount : 6
// PrimeAnswer1 : 100003
// CPU : 51 percent, HeapMemory : 182 KB, SystemMemory : 9299 KB, GoroutineCount : 5
// PrimeAnswer2 : 150001
// CPU : 40 percent, HeapMemory : 182 KB, SystemMemory : 9299 KB, GoroutineCount : 4
// PrimeAnswer3 : 200003
// CPU : 26 percent, HeapMemory : 183 KB, SystemMemory : 9299 KB, GoroutineCount : 3
// PrimeAnswer4 : 250007
// CPU : 12 percent, HeapMemory : 184 KB, SystemMemory : 9299 KB, GoroutineCount : 2
// PrimeAnswer5 : 300007
// CPU : 5 percent, HeapMemory : 184 KB, SystemMemory : 9299 KB, GoroutineCount : 1
