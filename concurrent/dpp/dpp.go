package main

import (
	"fmt"
	"sync"
	"time"
)

var noOfPhilosophers = 5

//lockForks 尝试同时锁两个互斥锁
func lockForks(a, b *sync.Mutex) {
	for {
		a.Lock()
		if b.TryLock() {
			return
		}
		a.Unlock()
		a, b = b, a
	}
}

// 信号量实现
type semaphore struct {
	ch chan struct{}
}

func newSemaphore() *semaphore {
	return &semaphore{ch: make(chan struct{}, 1)} // 容量为1的信号量
}

func (s *semaphore) wait() {
	s.ch <- struct{}{} // 获取信号量，如果信号量已满，这里会阻塞
}

func (s *semaphore) signal() {
	<-s.ch // 释放信号量，允许其他协程获取
}

type tableSetup struct {
	done bool //atomic
}

func NewTableSetup() *tableSetup {
	return &tableSetup{done: false}
}

type fork struct {
	id      int
	owner   int
	Dirty   bool
	m       *sync.Mutex
	sem     *semaphore
}

func NewFork(id, owner int) *fork {
	return &fork{id: id, owner: owner, Dirty: true, m: &sync.Mutex{}, sem: newSemaphore()}
}

func (f *fork) requests(ownerId int) {
	f.sem.wait() // 使用信号量等待
	defer f.sem.signal() // 完成后释放信号量

	for ownerId != f.owner {
		if f.Dirty {
			f.m.Lock()
			f.Dirty = false
			f.owner = ownerId
			f.m.Unlock()
		}
	}
}

func (f *fork) done_using(ownerId int) {
	f.Dirty = true
}

func (f *fork) getmutex() *sync.Mutex {
	return f.m
}

type philosopher struct {
	id          int
	table_setup *tableSetup
	left_fork   *fork
	right_fork  *fork
}

func NewPhilosopher(id int, table_setup *tableSetup, left_fork *fork, right_fork *fork) *philosopher {
	return &philosopher{id: id, table_setup: table_setup, left_fork: left_fork, right_fork: right_fork}
}

func (p *philosopher) dine(setup *tableSetup) {
	for !setup.done {
		p.think()
		p.eat()
	}
}

func (p *philosopher) think() {
	fmt.Println(p.id, "is thinking")
	// time.Sleep(time.Second) // 增加思考时间
}

func (p *philosopher) eat() {
	p.left_fork.requests(p.id)
	p.right_fork.requests(p.id)
	lockForks(p.left_fork.getmutex(), p.right_fork.getmutex())
	defer p.left_fork.getmutex().Unlock()
	defer p.right_fork.getmutex().Unlock()
	fmt.Println(p.id, "is eating")
	// time.Sleep(time.Second) // 增加就餐时间
	fmt.Println(p.id, "is done eating")
	p.left_fork.done_using(p.id)
	p.right_fork.done_using(p.id)
}

type table struct {
	setup       *tableSetup
	forks       []*fork
	philosopher []*philosopher
}

func NewTable() *table {
	s := NewTableSetup()
	f := make([]*fork, noOfPhilosophers)
	p := make([]*philosopher, noOfPhilosophers)
	for i := 0; i < noOfPhilosophers; i++ {
		f[i] = NewFork(i, i)
	}
	for i := 0; i < noOfPhilosophers; i++ {
		p[i] = NewPhilosopher(i, s, f[i], f[(i+1)%noOfPhilosophers])
	}
	return &table{setup: s, forks: f, philosopher: p}
}

func (t *table) start() {
	for i := 0; i < noOfPhilosophers; i++ {
		go t.philosopher[i].dine(t.setup)
	}
}

func (t *table) stop() {
	t.setup.done = true
}

func dine() {
	fmt.Println("Dinner started!")
	t := NewTable()
	t.start()
	time.Sleep(60 * time.Second)
	t.stop()
	fmt.Println("Dinner done!")
}

func main() {
	dine()
}
