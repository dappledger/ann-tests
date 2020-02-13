package main

import (
	"fmt"
	`math/rand`
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

func randMS()time.Duration{
	s :=rand.NewSource(time.Now().UnixNano())
	r:= rand.New(s).Int31n(1000*1000)
	return time.Duration(r)*time.Microsecond
}

func (b *Bench) startWorkers(worker int) {
	time.Sleep(time.Second+randMS())
	sleep := b.calcSleep()
	start:= time.Now().UnixNano()
	if b.dbg {
		tmp:= int64(sleep/time.Millisecond)
		fmt.Printf("worker_%d start...\n\tsleep %d'millsec once call\n",worker,tmp)
	}
	i:=0
	for i = 0; i < b.total1Thread; i++ {
		if atomic.LoadInt32(&b.stop) == 1 {
			break
		}
		_ = b.Do(worker)
		time.Sleep(sleep)
	}
	if b.dbg {
		cost := time.Now().UnixNano() - start
		cost /= int64(time.Second)
		if cost <=0 {
			cost = 1
		}
		fmt.Printf("work(%d,%d) cost %d's, exit...\n",worker,i,cost)
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
	return time.Duration(1000 / b.Ntps)*time.Millisecond
}

func (b *Bench) stat() {

}
