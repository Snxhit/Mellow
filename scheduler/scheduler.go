package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	loops []Loop
	stop  chan struct{}
	once  sync.Once
}

type Loop struct {
	Note     string
	Interval time.Duration
}

func New() *Scheduler {
	return &Scheduler{
		stop: make(chan struct{}),
	}
}

func (s *Scheduler) AddLoop(note string, intervalMS int) {
	s.loops = append(s.loops, Loop{
		Note:     note,
		Interval: time.Duration(intervalMS) * time.Millisecond,
	})
}

func (s *Scheduler) Run() {
	var wg sync.WaitGroup

	for _, loop := range s.loops {
		wg.Add(1)
		go func(l Loop) {
			defer wg.Done()
			s.runLoop(l)
		}(loop)
	}

	<-s.stop
	wg.Wait()
}

func (s *Scheduler) Stop() {
	s.once.Do(func() {
		close(s.stop)
	})
}

func (s *Scheduler) runLoop(l Loop) {
	ticker := time.NewTicker(l.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("PLAY:", l.Note)
		case <-s.stop:
			return
		}
	}
}

