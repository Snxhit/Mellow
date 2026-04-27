package audio

import (
	"math"
	"sync"

	"github.com/ebitengine/oto/v3"
)

const sampleRate = 44100

type Engine struct {
	ctx    *oto.Context
	player *oto.Player
	mu     sync.Mutex
	notes  []Note
}

type Note struct {
	freq  float64
	phase float64
	life  int
}

func New() (*Engine, error) {
	op := &oto.NewContextOptions{
		SampleRate:   sampleRate,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	}

	ctx, ready, err := oto.NewContext(op)
	if err != nil {
		return nil, err
	}
	<-ready

	e := &Engine{
		ctx: ctx,
	}

	p := ctx.NewPlayer(e)
	e.player = p
	p.Play()

	return e, nil
}

func (e *Engine) Read(buf []byte) (int, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	samples := len(buf) / 4

	for i := 0; i < samples; i++ {
		var sample float64

		alive := e.notes[:0]

		for j := range e.notes {
			n := &e.notes[j]

			sample += math.Sin(2 * math.Pi * n.phase)
			n.phase += n.freq / sampleRate
			n.life--

			if n.life > 0 {
				alive = append(alive, *n)
			}
		}

		e.notes = alive

		if sample > 1 {
			sample = 1
		}
		if sample < -1 {
			sample = -1
		}

		v := int16(sample * 3000)

		idx := i * 4

		buf[idx] = byte(v)
		buf[idx+1] = byte(v >> 8)
		buf[idx+2] = byte(v)
		buf[idx+3] = byte(v >> 8)
	}

	return len(buf), nil
}

func (e *Engine) Play(freq float64) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.notes = append(e.notes, Note{
		freq:  freq,
		life:  sampleRate / 4,
		phase: 0,
	})
}
