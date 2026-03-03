package ast

type Program struct {
	Statements []Statement
}

type Statement interface {
	isStatement()
}

type PlayLoop struct {
	Note       string
	IntervalMS int
}

func (PlayLoop) isStatement() {}
