package receivernamelen

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "receivernamelen checks the length of the receiver variable names"

// Analyzer is receivernamelen analyzer
var Analyzer = &analysis.Analyzer{
	Name: "receivernamelen",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if n.Recv != nil {
				for _, field := range n.Recv.List {
					for _, name := range field.Names {
						if len([]rune(name.String())) >= 3 {
							pass.Reportf(name.Pos(), "receiver variable names must be one or two letters in length")
						}
					}
				}
			}
		}
	})

	return nil, nil
}
