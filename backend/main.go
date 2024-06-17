package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
//	defer timer("sum")()

func main() {
	var filename = "../results/test.csv"

	log.Println("Normally creating file", filename)

	// Prepare testing file
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Something went wrong with opening file", err)
		return
	}
	defer func() {
		file.Close()
	}()

	// Start clean
	file.Truncate(0)

	// Set headers
	WriteLineInFile(file, "n,durationms,total_loops,group_limit")

	// Run
	BulkTest(
		file,
		TestApiNormalWait,
		5000,
		1000,
	)
}

// NormalWait waits for a duration drawn from a normal distribution.
// func NormalWait(mean, stdDev time.Duration) {
// 	// Use math.Randn to generate a standard normal variate
// 	x := math.Randn(rand.New(rand.NewSource(time.Now().UnixNano())))

// 	// Scale and offset the variate to fit the desired distribution
// 	waitTime := time.Duration(x*stdDev) + mean

// 	// Ensure waitTime is non-negative (cannot wait negative duration)
// 	if waitTime < 0 {
// 		waitTime = 0
// 	}

// 	time.Sleep(waitTime)
// }

func TestApiNormalWait() {
	var rnd = rand.NormFloat64()*110.5 + 1000

	time.Sleep(
		time.Duration(time.Duration(rnd) * time.Millisecond),
	)

}

func BulkTest(file *os.File, ToTestFunc func(), total_loops int, group_limit int) {
	if total_loops == 0 || group_limit == 0 {
		log.Println("please add total and / or group limit")
		return
	}

	var start = GetTimer("main")

	// Set my waitgroup
	var wg sync.WaitGroup

	// Create gogroup
	group := errgroup.Group{}

	// Set concurrency limit
	group.SetLimit(group_limit)

	var total_time time.Duration

	for i := 1; i <= total_loops; i++ {

		wg.Add(1)
		group.Go(func() error {
			var endTimer = GetTimer("main" + fmt.Sprintf("%d", i))
			defer func() {
				wg.Done()
				var elaps = endTimer()
				total_time += elaps
				WriteLineInFile(
					file,
					fmt.Sprintf("%v,%v,%v,%v", i, elaps.Microseconds(), total_loops, group_limit),
				)
			}()
			ToTestFunc()
			return nil
		})
	}

	wg.Wait()

	fmt.Printf("Total function took total [%v]; in total %v runs, cum-time: %v and average time of [%v]ms\n", start(), total_loops, total_time, int(total_time.Milliseconds())/total_loops)

}

func FakeTest() {
	// Generate a random number of seconds to wait (between 1 and 10 seconds)
	delay := (rand.Intn(1500) + 1)
	fmt.Println("Waiting for ", delay, " milliseconds")

	// Wait for the random number of seconds
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
func TestApi() {

	// Define the API endpoint
	url := "https://dummyjson.com/todos/1"

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	// Set headers if necessary

	// Make the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received non-200 status code: %d\n", resp.StatusCode)
		return
	}

	// Read the response body
	io.ReadAll(resp.Body)

}
