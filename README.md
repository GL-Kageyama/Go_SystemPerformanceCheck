# System Performance Check

## Overview
This sample is used to output logs to measure system performance.  
Five goroutines are started to find prime numbers, and the goroutines are terminated when the calculation is finished.  
The system log at the end of each calculation is output.  

## Library
The library used is "gopsutil".   
https://github.com/shirou/gopsutil    

This library can retrieve a variety of information, including CPU and memory.  

Please install as needed.   
$ go get github.com/shirou/gopsutil/cpu  

## Code
```Go
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
```

## Output Sample
$ go build -o SystemPerformanceCheck SystemPerformanceCheck.go  
$ ./SystemPerformanceCheck  
100000 Line Target Goroutine Wake UP  
200000 Line Target Goroutine Wake UP  
250000 Line Target Goroutine Wake UP  
150000 Line Target Goroutine Wake UP  
300000 Line Target Goroutine Wake UP  
CPU : 63 percent, HeapMemory : 180 KB, SystemMemory : 9043 KB, GoroutineCount : 6  
PrimeAnswer1 : 100003  
CPU : 51 percent, HeapMemory : 182 KB, SystemMemory : 9299 KB, GoroutineCount : 5  
PrimeAnswer2 : 150001  
CPU : 40 percent, HeapMemory : 182 KB, SystemMemory : 9299 KB, GoroutineCount : 4  
PrimeAnswer3 : 200003  
CPU : 26 percent, HeapMemory : 183 KB, SystemMemory : 9299 KB, GoroutineCount : 3  
PrimeAnswer4 : 250007  
CPU : 12 percent, HeapMemory : 184 KB, SystemMemory : 9299 KB, GoroutineCount : 2  
PrimeAnswer5 : 300007  
CPU : 5 percent, HeapMemory : 184 KB, SystemMemory : 9299 KB, GoroutineCount : 1  
