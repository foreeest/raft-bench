package util

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	// "strconv"
	"time"
	// "github.com/foreeest/raftbench/dragonboat"
)

type TestParams struct {
	NumKeys   int
	Mil       int
	Runs      int
	Wait      time.Duration
	FirstWait time.Duration
	Step      time.Duration
	MaxTries  int
	Enabled   bool
	LogFile   string
}

func Bench(testParams TestParams, read func(string) bool, write func(string, string) bool) {
	defer WaitForCtrlC()
	if !testParams.Enabled {
		return
	}

	f, err := os.Create(testParams.LogFile)
	if err != nil {
		log.Fatal("unable to create csv log")
	}
	defer f.Close()

	time.Sleep(testParams.FirstWait) // 10s now
	log.Printf("Starting benchmark...\n")
	for i := 0; i < testParams.Runs; i++ {
		log.Printf("BENCHMARK %v OF %v\n", i+1, testParams.Runs)
		time.Sleep(testParams.Wait) // 3s now

		start := time.Now()
		failure := 0
		for k := 0; k < testParams.NumKeys*testParams.Mil; k++ { // for 1* 1000 times
			v := rand.Int()
			tries := 0
			// tries := 1
			// ok := false
			// ok = write(fmt.Sprintf("%d", k), fmt.Sprintf("%d", v))
			for ok := false; !ok; ok = write(fmt.Sprintf("%d", k), fmt.Sprintf("%d", v)) {
				time.Sleep(testParams.Step) // 0.1s by default, but now set to 0.001s
				tries++
				if tries > testParams.MaxTries { // 10
					break
				}
			}
			failure += tries
		}
		duration := time.Since(start).Seconds()
		// dragonboat.RequestDuration.WithLabelValues("w_operation" + strconv.Itoa(i+1)).Observe(duration)
		_, _ = f.WriteString(fmt.Sprintf("write,%v,%v,%v,%v\n", i+1, failure, testParams.NumKeys*testParams.Mil, duration))
		//dragonboat.requestDuration.WithLabelValues(r.URL.Path).Observe(duration)

		time.Sleep(testParams.Wait) // 3s
		start = time.Now()
		failure = 0
		for k := 0; k < testParams.NumKeys*testParams.Mil; k++ {
			tries := 0
			// tries := 1
			// ok := false
			// ok = read(fmt.Sprintf("%d", k))
			for ok := false; !ok; ok = read(fmt.Sprintf("%d", k)) {
				time.Sleep(testParams.Step)
				tries++
				if tries > testParams.MaxTries {
					break
				}
			}
			failure += tries
		}
		duration = time.Since(start).Seconds()
		// dragonboat.RequestDuration.WithLabelValues("r_operation" + strconv.Itoa(i+1)).Observe(duration)
		_, _ = f.WriteString(fmt.Sprintf("read,%v,%v,%v,%v\n", i+1, failure, testParams.NumKeys*testParams.Mil, duration))
	}
	log.Printf("BENCHMARK COMPLETE\n")
}
