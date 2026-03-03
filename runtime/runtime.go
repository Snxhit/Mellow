package runtime

import (
	"Mellow/ast"
	"Mellow/scheduler"
)

type Runtime struct {
	sch    *scheduler.Scheduler
	reload chan *ast.Program
}

func New() *Runtime {
	return &Runtime{
		reload: make(chan *ast.Program, 1),
	}
}

func (r *Runtime) Load(p *ast.Program) {
	select {
	case r.reload <- p:
	default:
		<-r.reload
		r.reload <- p
	}
}

func (r *Runtime) Run() {
	for {
		prog := <-r.reload

		if r.sch != nil {
			r.sch.Stop()
		}

		sch := scheduler.New()

		for _, st := range prog.Statements {
			switch s := st.(type) {
			case ast.PlayLoop:
				sch.AddLoop(s.Note, s.IntervalMS)
			}
		}

		r.sch = sch
		go r.sch.Run()
	}
}

