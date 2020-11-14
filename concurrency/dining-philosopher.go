package main2

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func getPhilosopherHosts() []string {
	return []string{
		"140.158.131.59:12000",
		"140.158.130.53:12000",
		"140.158.129.240:12000",
		"140.158.130.72:12000",
		"140.158.129.243:12000",
	}
}

func getForkHosts() []string {
	return []string{
		"140.158.131.59:12030",
		"140.158.130.53:12030",
		"140.158.129.240:12030",
		"140.158.130.72:12030",
		"140.158.129.243:12030",
	}
}

const displayStatusHost = "140.158.130.73:12000"

const maxPhilosophers = 5 // There are five philosophers around the table
const maxForks = 5        // There are five chopticks on the table
const maxTimeToEat = 3    // philosophers can eat max 3 times

// Fork represents a fork along with a mechanism to lock it
type Fork struct{ sync.Mutex }

// Philosopher allows to handle the process of eating for a philosopher, he has :
type Philosopher struct {
	id                  int
	countEating         int
	status              string
	state               chan string
	leftFork, rightFork *Fork
	feedbackChannel     chan bool
}

// Request is used by the philosophers to send messages to the Host :
type Request struct {
	command     string
	philosopher Philosopher
}

// Below are the allowed command for the Request struct
const wantToEat = "wantToEat"
const finishedEating = "finishedEating"

// Bellow are the status
const thinking = "thinking"
const waiting = "waiting"
const eating = "eating"

// eat function allows to start the process of eating for a philosopher
func (philosopher *Philosopher) eat(requestChan chan Request, wg *sync.WaitGroup, changeStatus chan string) {
	philosopher.countEating = 0

	for philosopher.countEating < maxTimeToEat {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		changeStatus <- waiting
		philosopher.status = waiting
		requestChan <- Request{command: wantToEat, philosopher: *philosopher}
		isPhilosopherAllowedToEat := <-philosopher.feedbackChannel

		if isPhilosopherAllowedToEat {

			philosopher.leftFork.Lock()
			philosopher.rightFork.Lock()
			philosopher.status = eating
			changeStatus <- eating
			time.Sleep(time.Duration((rand.Intn(500) + 50)) * time.Millisecond)
			philosopher.rightFork.Unlock()
			philosopher.leftFork.Unlock()

			philosopher.countEating++

			philosopher.status = thinking
			changeStatus <- thinking
			requestChan <- Request{command: finishedEating, philosopher: *philosopher}
			defer wg.Done()
		}
	}
	close(philosopher.feedbackChannel)
}

// Host receives requests to eat from the philosophers, the host decide to accept or reject each request and ensures that :
func Host(requestChan chan Request) {
	var philosophersEating = make(map[int]Philosopher)

	for request := range requestChan {

		switch request.command {
		case wantToEat:
			if len(philosophersEating) == 0 {
				philosophersEating[request.philosopher.id] = request.philosopher
				AcceptRequestToEat(&request.philosopher)
			} else if len(philosophersEating) == 1 {
				var keys []int
				for k := range philosophersEating {
					keys = append(keys, k)
				}
				var philosopherCurrentlyEating = keys[0]
				var philosopherAskingToEat = request.philosopher.id
				// Neighborhoods ?
				if philosopherCurrentlyEating == 0 && philosopherAskingToEat == maxPhilosophers {
					RejectRequestToEat(&request.philosopher, "Neighborhood 0-")
				} else if philosopherCurrentlyEating == maxPhilosophers && philosopherAskingToEat == 0 {
					RejectRequestToEat(&request.philosopher, "Neighborhood -0")
				} else if math.Abs(float64(philosopherAskingToEat-philosopherCurrentlyEating)) == 1.0 {
					RejectRequestToEat(&request.philosopher, "Neighborhood")
				} else if philosopherAskingToEat == philosopherCurrentlyEating {
					RejectRequestToEat(&request.philosopher, "Philosopher already eating")
				} else {
					AcceptRequestToEat(&request.philosopher)
				}
			} else {
				RejectRequestToEat(&request.philosopher, "All allowed philoshopers are already eating")
			}
		case finishedEating:
			delete(philosophersEating, request.philosopher.id)
		}
	}
}

// RejectRequestToEat sends a message back to the philosopher denying him to eat
func RejectRequestToEat(philosopher *Philosopher, rejectReason string) {
	philosopher.feedbackChannel <- false
}

// AcceptRequestToEat sends a message back to the philosopher allowing him to eat
func AcceptRequestToEat(philosopher *Philosopher) {
	philosopher.feedbackChannel <- true
}

func main() {
	// Creating the Forks
	var changeStatus = make(chan string)
	var forks = make([]*Fork, maxForks)
	for fork := 0; fork < maxForks; fork++ {
		forks[fork] = new(Fork)
	}

	// Creating the Philosophers
	var philosophers = make([]*Philosopher, maxPhilosophers)
	for philosopher := 0; philosopher < maxPhilosophers; philosopher++ {
		var leftForkID = philosopher
		var rightForkID = (philosopher + 1) % maxPhilosophers
		philosophers[philosopher] = &Philosopher{
			id:              philosopher,
			countEating:     0,
			status:          thinking,
			state:           make(chan string),
			leftFork:        forks[leftForkID],
			rightFork:       forks[rightForkID],
			feedbackChannel: make(chan bool)}
	}

	// A wait group to allow the main program to wait for all the philosophers to eat 3 times
	var wg sync.WaitGroup
	wg.Add(maxPhilosophers * maxTimeToEat)

	var requestChan = make(chan Request)

	go Host(requestChan)

	// Create and start the goroutines for the philosophers
	for _, philosopher := range philosophers {
		go philosopher.eat(requestChan, &wg, changeStatus)
	}

	for range changeStatus {
		toPrint := time.Now().Format("01-02 15:04:05") + " : ---\t---"
		for _, philosopher := range philosophers {
			toPrint += philosopher.status + "-\t-- "
		}
		fmt.Println(toPrint)
	}

	wg.Wait()
	close(requestChan)
	close(changeStatus)

	fmt.Println("All philosophers have finished eating, good bye")
}
