package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter Counter
var tracks Tracks
var winner chan int
var sleepTime int
var distanceGoal int
var numRacers int

func newCounter(size int) Counter {
	return Counter{racers: make([]int, size+1)} // not using the zero'th index
}

func newTracks(size int) Tracks {
	return Tracks{racers: make([]string, size+1)}
}

func incrementCounter(n int) {
	counter.racers[n]++
}

func equalCounter() bool {
	rv := true
	for i, _ := range counter.racers {
		if counter.racers[1] != counter.racers[i] {
			rv = false
			break
		}
	}
	return rv
}

func smallerCounter(racerNumber int) bool {
	var rv bool
	for v := range counter.racers {
		if counter.racers[racerNumber] < v {
			rv = true
		} else {
			rv = false
		}
	}

	return rv
}

func checkCounter(racerNumber int) bool {
	var rv bool
	if equalCounter() || smallerCounter(racerNumber) {
		rv = true
	} else {
		rv = false
	}

	return rv
}

func print() {
	clearScreen()
	for i, v := range tracks.racers {
		if 0 < i {
			fmt.Println(v + "[" + strconv.Itoa(i) + "]")
		}
	}
}

func racer(racerNum int) {
	steps := rand.Intn(sleepTime) + 1
	distance := 0
	if checkCounter(racerNum) {
		for j := distance; distance < distanceGoal; j++ {
			for i := 0; i < steps; i++ {
				tracks.racers[racerNum] += "x"
			}
			distance += steps
			duration := rand.Int31n(int32(sleepTime))
			time.Sleep(time.Duration(duration) * time.Millisecond)
			print()
		}

		incrementCounter(racerNum)
	}

	if distance >= distanceGoal {
		winner <- racerNum
	}

	wg.Done()
}

func init() {
	numRacers = 5
	sleepTime = 20
	distanceGoal = 200
	counter = newCounter(numRacers)
	tracks = newTracks(numRacers)
	winner = make(chan int, numRacers)
}

func main() {
	for i := 1; i <= numRacers; i++ {
		wg.Add(1)
		go racer(i)
	}

	wg.Wait()

	fmt.Print("Winner is : #")
	fmt.Print(<-winner)
}
