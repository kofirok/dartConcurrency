# Project 5 Concurrency
# Khalid Kofiro


This project simulates throwing darts at the unit circle to estimate the value of pi.

# Functions

* The `ThrowDart` struct keeps track of the total number of darts thrown and the number of darts
that hit the target. 

* The `NewThrowDart` function creates a new instance of `ThrowDart` with a specified
number of darts to throw. 

* The `Run` method throws the darts and counts the number that hit the target.

* The `main` function initializes a number of darts to be thrown, runs a series of independent `ThrowDart`
simulations using goroutines, and then calculates an estimate of Pi based on the proportion of darts
that hit the unit circle. Then it measures how close the estimated pi value is to the actual value of 
true pi for various numbers of darts thrown. It also measures the time taken to calculate this for each 
estimate. By increasing the thread count you can run the dart simulation concurrently because of the 
implementation of WaitGroup.


The estimate of Pi is printed to the console for each thread ran.

# Running the code

* To run the code, save it with a .go extension.

* Open a terminal or command prompt and navigate to the directory containing the file.

* Once you are in the correct directory type `go run filename.go` (dont forget to replace 'filename' with the actual name of your file).

* Adjust the amount of Darts/Threads as you please

* This will compile and run the program, displaying the estimated value of pi on the terminal. See below.


```sh
go run main.go

Estimate of Pi: 3.159100
Delta from true Pi: 0.017507
 
Estimate of Pi: 3.146260
Delta from true Pi: 0.004667
 
Estimate of Pi: 3.150300
Delta from true Pi: 0.008707
 
Time taken for 200000 Darts with 3 thread(s): 116.48725ms
```

Hope you enjoy this simulation!