package scheduler

import (
	"sync"
	"time"

	"Mellow/audio"
)

type Scheduler struct {
	loops []Loop
	stop  chan struct{}
	once  sync.Once
	audio *audio.Engine
}

type Loop struct {
	Note     string
	Interval time.Duration
}

func New() (*Scheduler, error) {
	a, err := audio.New()
	if err != nil {
		return nil, err
	}

	return &Scheduler{
		stop:  make(chan struct{}),
		audio: a,
	}, nil
}

func (s *Scheduler) AddLoop(note string, intervalMS int) {
	if intervalMS <= 0 {
		return
	}
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
	if l.Interval <= 0 {
		return
	}

	ticker := time.NewTicker(l.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if freq, ok := audio.Notes[l.Note]; ok {
				s.audio.Play(freq)
			}
		case <-s.stop:
			return
		}
	}
}
