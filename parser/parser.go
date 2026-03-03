package parser

import "Mellow/ast"

func Parse(src string) (*ast.Program, error) {
	return &ast.Program{
		Statements: []ast.Statement{
			ast.PlayLoop{
				Note:       "C4",
				IntervalMS: 500,
			},
		},
	}, nil
}
