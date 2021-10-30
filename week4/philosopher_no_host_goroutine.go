package main

import (
	"fmt"
	"sync"
	"time"
)

// DEFAULT_HUNGER controls how many times a Philosopher Eats
const DEFAULT_HUNGER = 3

// Host a host has a channel (bounded) to allow concurrent eating of Philosopher
type Host struct {
	allowedConcurrency chan bool
}

// CanEat removes from the allowedConcurrency channel, when allowedConcurrency is 0 the host forbids dining
func (h *Host) CanEat() {
	h.allowedConcurrency <- true
}

// CanEat push to allowedConcurrency channel, when allowedConcurrency is > 0 the host allows dining
func (h *Host) DoneEat() {
	<-h.allowedConcurrency
}

// NewTwoSeatsHost is a Host that allows at most 2 concurrent Philosopher dining
func NewTwoSeatsHost() Host {
	allowedCh := make(chan bool, 2)
	return Host{
		allowedConcurrency: allowedCh,
	}
}

// Chopstick is the chopstick in the Dining Philosopher problem
type Chopstick struct {
	sync.Mutex
}

// Philosopher is a philosopher in the dining problem. Holds 2 chopsticks, knows the Host and eats till is Hungry
type Philosopher struct {
	Id             int
	Hunger         int
	Host           Host
	RightChopstick *Chopstick
	LeftChopstick  *Chopstick
}

//Eat is the way a Philosopher consumes food
func (ph *Philosopher) Eat(wg *sync.WaitGroup) {
	defer wg.Done()     // signal this goroutine is done
	if ph.Hunger <= 0 { // if not hungry stop eating
		return
	}
	defer func() { wg.Add(1); ph.Eat(wg) }() // if still hungry schedule another dinner after this one

	//leave the chopstick, as stated in the problem
	defer ph.RightChopstick.Unlock()
	defer ph.LeftChopstick.Unlock()

	// a Philosopher thinks
	fmt.Printf("thinking (%d)...\n", ph.Id)
	time.Sleep(1 * time.Second)

	//a Philosopher gets the chopsticks to eat
	ph.LeftChopstick.Lock()
	ph.RightChopstick.Lock()

	//a Philosopher Checks if can eat, if not it blocks for teh Host's permission
	ph.Host.CanEat()

	//a Philosopher eats for a bit
	fmt.Printf("starting to Eat %d\n", ph.Id)
	time.Sleep(2 * time.Second)

	//a Philosopher decreases its hunger and signals to the Host that he's done with food
	ph.Hunger--
	ph.Host.DoneEat()
	fmt.Printf("finishing eating %d\n", ph.Id)
}

func main() {
	wg := &sync.WaitGroup{}
	twoSeatsHost := NewTwoSeatsHost() //Host

	chopsticks := [5]*Chopstick{} //Chopsticks
	for i := 0; i < 5; i++ {
		chopsticks[i] = &Chopstick{}
	}

	phs := make([]Philosopher, 5) //Philosophers
	for i := 0; i < 5; i++ {
		phs[i] = Philosopher{
			Id:             i + 1,
			Hunger:         DEFAULT_HUNGER,
			LeftChopstick:  chopsticks[i],
			RightChopstick: chopsticks[(i+1)%5],
			Host:           twoSeatsHost,
		}
	}

	// Spawn 5 go routine equal to the number of Philosophers,
	// then each Philosopher will eventually create a new goroutine
	// given its current hunger level
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go phs[i].Eat(wg)
	}
	wg.Wait() //wait for Philosopher to have done with dining, and terminate

}
