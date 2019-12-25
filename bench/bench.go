package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type BCRet int

const (
	BCErr     BCRet = -1
	BCSuccess BCRet = 0
	BCOK            = 1
)

type BenchCall func(workerIdx int) BCRet

type Bench struct {
	BenchBase
	stop int32
	dbg  bool
}

type BenchBase struct {
	Nthread      int
	Ntps         int
	total1Thread int //ntps*nsec
	Do           BenchCall
}

func (b *Bench) Running() {
	b.stop = 0
	//
	go b.doSignal()
	//
	var wg sync.WaitGroup
	wg.Add(b.Nthread)
	for i := 0; i < b.Nthread; i++ {
		go func(worker int) {
			b.startWorkers(worker)
			wg.Done()
		}(i)
	}
	wg.Wait()
	//
	b.stat()
}

func (b *Bench) startWorkers(worker int) {
	if b.dbg {
		fmt.Println("worker(", worker, ") start...")
	}
	sleep := b.calcSleep()
	for i := 0; i < b.total1Thread; i++ {
		if atomic.LoadInt32(&b.stop) == 1 {
			break
		}
		_ = b.Do(worker)
		time.Sleep(sleep)
	}
	if b.dbg {
		fmt.Println("worker(", worker, ") exit...")
	}
}

func (b *Bench) doSignal() {
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	v := <-sc
	if b.dbg {
		fmt.Println("recv sig=", v)
	}
	atomic.CompareAndSwapInt32(&b.stop, 0, 1)
}

func (b *Bench) calcSleep() time.Duration {
	if b.Ntps > 1000 {
		return time.Millisecond
	}
	return time.Duration(1000 / b.Ntps)
}

func (b *Bench) stat() {

}
