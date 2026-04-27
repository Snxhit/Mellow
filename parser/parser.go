package parser

import "Mellow/ast"
import "fmt"
import "strconv"
import "strings"

func Parse(src string) (*ast.Program, error) {
	st := []ast.Statement{}
	tokens := strings.Fields(src)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "loop":
			if i+4 >= len(tokens) {
				return nil, fmt.Errorf("invalid loop statement: expected 5 tokens starting at %q", tokens[i])
			}
			if tokens[i+1] != "note" || tokens[i+3] != "every" {
				return nil, fmt.Errorf("invalid loop statement: expected format 'loop note <NOTE> every <MS>'")
			}
			ims, err := strconv.Atoi(tokens[i+4])
			if err != nil {
				return nil, fmt.Errorf("invalid interval %q: %w", tokens[i+4], err)
			}
			if ims <= 0 {
				return nil, fmt.Errorf("invalid interval %d: must be > 0", ims)
			}
			st = append(st, ast.PlayLoop{
				Note:       tokens[i+2],
				IntervalMS: ims,
			})
			i += 4
		}
	}
	return &ast.Program{
		Statements: st,
	}, nil
}
