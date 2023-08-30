package main

// Project 5 Concurrency
// Khalid Kofiro

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

/* ThrowDart is a struct for the set of darts thrown/hit */

type ThrowDart struct {
	totalThrown int // total number of darts thrown
	totalHit    int // total number of darts that hit the target
}

/* NewThrowDart creates a new instance of ThrowDart with a specified number of darts to throw */

func NewThrowDart(totalThrown int) *ThrowDart {
	return &ThrowDart{
		totalThrown: totalThrown,
		totalHit:    0,
	}
}

/* Run throws the specified number of darts and counts the number of darts that hit the target */

func (ts *ThrowDart) Run() {
	rand.Seed(time.Now().UnixNano()) // set random seed to the current time
	for i := 0; i < ts.totalThrown; i++ {
		x := rand.Float64()
		y := rand.Float64() // generates random x and y coordinates
		if x*x+y*y < 1.0 {  // checks if the dart hit the target (unit circle)
			ts.totalHit++ // if it did: increment the total number of darts that hit the target
		}
	}
	// The results will print for each thread
	piEstimate := float64(ts.totalHit) / float64(ts.totalThrown) * 4 // calculate an estimate of pi
	fmt.Printf("Estimate of Pi: %f\n", piEstimate)                   // print the estimate of pi
	piDelta := math.Abs(piEstimate - math.Pi)                        // calculate the absolute difference between piEstimate and math.Pi
	fmt.Printf("Delta from true Pi: %f\n", piDelta)                  // print the delta
	fmt.Print(" \n")
}

/*
The main function initializes a number of darts to be thrown, runs a series of independent
ThrowDart simulations, and then calculates an estimate of Pi based on the proportion of
darts that hit a circular target. It also measures how close the estimated pi value is
to the actual value of true pi for various numbers of darts thrown. By increasing the thread
count you can run the dart simulation concurrently because of the implementation of WaitGroup. */

func main() {
	start := time.Now()                        // record the start time
	NUM_THREADS := 3                           // Adjust number of concurrent threads (sets of darts) to use
	NUM_THROWS := 200000                       // Adjust number of darts to throw in each set
	samples := make([]*ThrowDart, NUM_THREADS) // create a slice of ThrowDart instances

	var wg sync.WaitGroup // initialize a WaitGroup

	for i := 0; i < NUM_THREADS; i++ {
		samples[i] = NewThrowDart(NUM_THROWS) // create a new instance of ThrowDart and add it to the slice
		wg.Add(1)                             // increment the WaitGroup counter
		go func(td *ThrowDart, wg *sync.WaitGroup) {
			defer wg.Done() // decrement the WaitGroup counter when the goroutine exits
			td.Run()        // throw the darts and count the number that hit the target
		}(samples[i], &wg)
	}

	wg.Wait() // wait for all goroutines to complete

	elapsed := time.Since(start) // calculate the elapsed time
	fmt.Printf("Time taken for %d Darts with %d thread(s): %s\n", NUM_THROWS, NUM_THREADS, elapsed)
}
