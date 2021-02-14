// show how work can be distributed
// onto multiple cores, via a simple
// for loop.
// printing on one single terminal
// disturbs a little, but it seems
// sufficient to show the principle.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	count()
	fmt.Println("work done.")
}


func count() {
	// calculate what needs to be done 
	// on every CPU core
	var nCPU = runtime.NumCPU()
	batchsize := 1000 / nCPU

	var wg sync.WaitGroup
	wg.Add(nCPU)

	// loop over the CPU cores, and start one
	// goroutine for every core to process its batch
	for k :=  0; k < nCPU; k++ {
		start := (k * batchsize) + 1
		end := (k + 1) * batchsize
		fmt.Println("start:", start, "end:", end)
		go func(start, end int) {
			// on return, notify the waitgroup we are done
			defer wg.Done()
			// do work on the batch passed
			for j := start; j <= end; j++ {
				fmt.Println(j)
			}
		}(start, end)
	}
	// wait until every batch is done
	wg.Wait()
}
