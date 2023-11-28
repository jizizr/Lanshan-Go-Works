package main

import (
	"fmt"
	"sync"
	"time"
)

const noOfPhilosophers = 5

// SyncChannel is used for signaling.
type SyncChannel struct {
	cond *sync.Cond
}

func NewSyncChannel() *SyncChannel {
	return &SyncChannel{cond: sync.NewCond(&sync.Mutex{})}
}

func (sc *SyncChannel) Wait() {
	sc.cond.L.Lock()
	sc.cond.Wait()
	sc.cond.L.Unlock()
}

func (sc *SyncChannel) NotifyAll() {
	sc.cond.Broadcast()
}

type TableSetup struct {
	done    bool
	channel *SyncChannel
}

type Fork struct {
	id      int
	owner   int
	dirty   bool
	mutex   sync.Mutex
	channel *SyncChannel
}

func NewFork(forkId, ownerId int) *Fork {
	return &Fork{id: forkId, owner: ownerId, dirty: true, channel: NewSyncChannel()}
}

func (f *Fork) request(ownerId int) {
	for f.owner != ownerId {
		if f.dirty {
			f.mutex.Lock()

			f.dirty = false
			f.owner = ownerId

			f.mutex.Unlock()
		} else {
			f.channel.Wait()
		}
	}
}

func (f *Fork) doneUsing() {
	f.dirty = true
	f.channel.NotifyAll()
}

type Philosopher struct {
	id        int
	name      string
	setup     *TableSetup
	leftFork  *Fork
	rightFork *Fork
}

func NewPhilosopher(id int, name string, setup *TableSetup, leftFork, rightFork *Fork) *Philosopher {
	return &Philosopher{id: id, name: name, setup: setup, leftFork: leftFork, rightFork: rightFork}
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for !p.setup.done {
		p.think()
		p.eat()
	}
}

func (p *Philosopher) print(text string) {
	fmt.Printf("%-10s %s\n", p.name, text)
}

func (p *Philosopher) eat() {
	p.leftFork.request(p.id)
	p.rightFork.request(p.id)

	// Lock both forks at the same time to avoid deadlock.
	lock := sync.Mutex{}
	lock.Lock()
	p.leftFork.mutex.Lock()

	p.rightFork.mutex.Lock()
	lock.Unlock() // Unlock the temporary lock after both forks are locked

	p.print("started eating")
	time.Sleep(time.Second) // Simulate eating
	p.print("finished eating")

	p.leftFork.doneUsing()
	p.rightFork.doneUsing()

	p.leftFork.mutex.Unlock()
	p.rightFork.mutex.Unlock()
}

func (p *Philosopher) think() {
	p.print("is thinking")
	time.Sleep(time.Second) // Simulate thinking
}

type Table struct {
	setup  TableSetup
	forks  []*Fork
	philos []*Philosopher
}

func NewTable() *Table {
	setup := TableSetup{done: false, channel: NewSyncChannel()}
	forks := make([]*Fork, noOfPhilosophers)
	philos := make([]*Philosopher, noOfPhilosophers)

	for i := 0; i < noOfPhilosophers; i++ {
		forks[i] = NewFork(i+1, (i+1)%noOfPhilosophers)
	}

	for i := 0; i < noOfPhilosophers; i++ {
		philos[i] = NewPhilosopher(i+1, fmt.Sprintf("Philosopher %d", i+1), &setup, forks[i], forks[(i+1)%noOfPhilosophers])
	}

	return &Table{
		setup:  setup,
		forks:  forks,
		philos: philos,
	}
}

func (t *Table) start() {
	var wg sync.WaitGroup
	for _, ph := range t.philos {
		wg.Add(1)
		go ph.dine(&wg)
	}

	t.setup.channel.NotifyAll() // Let all philosophers start dining
	wg.Wait()                   // Wait for all dining to be done
}

func (t *Table) stop() {
	t.setup.done = true
}

func main() {
	fmt.Println("Dinner started!")

	table := NewTable()
	go func() {
		time.Sleep(60 * time.Second) // Dine for 60 seconds
		table.stop()
	}()

	table.start()

	fmt.Println("Dinner done!")
}
