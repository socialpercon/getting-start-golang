package main

import (
	"fmt"

	"sync"

	"time"
)

type Stat struct {
	sync.Mutex

	counters map[string]int64
}

type StatAverage Stat

func (s *Stat) increaseCounter(key string) {

	s.Lock()

	defer s.Unlock()

	if _, exists := s.counters[key]; !exists {

		s.counters[key] = 1

	} else {
		for i := 0; i < 1000000; i++ {
			s.counters[key] += 1
		}
	}

}

func main() {

	s := &Stat{

		counters: make(map[string]int64),
	}

	for i := 0; i < 100; i++ {
		go s.increaseCounter("test")

		go s.increaseCounter("test")
	}

	s.increaseCounter("test")

	for {
		time.Sleep(time.Millisecond)

		fmt.Printf("%#v\n", s.counters)
	}

}
