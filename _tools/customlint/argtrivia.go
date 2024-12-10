package customlint

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var argTriviaAnalyzer = &analysis.Analyzer{
	Name: "argtrivia",
	Doc:  "requires argument name annotations",
	Run:  runArgTrivia,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func runArgTrivia(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.File)(nil),
		(*ast.CallExpr)(nil),
	}

	var file *ast.File

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			file = n
		case *ast.CallExpr:
			checkCallExpression(pass, file, n)
		}
	})

	return nil, nil
}

func checkCallExpression(pass *analysis.Pass, file *ast.File, call *ast.CallExpr) {
	if target, ok := call.Fun.(*ast.SelectorExpr); ok {
		obj := pass.TypesInfo.ObjectOf(target.Sel)
		pkg := obj.Pkg()
		if pkg == nil {
			return
		}

		pkgPath := pkg.Path()

		const (
			ourRepo      = "github.com/microsoft/TypeScript-go"
			ourRepoSlash = ourRepo + "/"
		)

		if pkgPath != ourRepo && !strings.HasPrefix(pkgPath, ourRepoSlash) {
			return
		}
	}

	funType := pass.TypesInfo.TypeOf(call.Fun)
	if t, ok := funType.(*types.Pointer); ok {
		funType = t.Elem()
	}
	sig, ok := funType.(*types.Signature)
	if !ok {
		return
	}
	paramTuple := sig.Params()

	end := len(call.Args) - 1

	for i, arg := range call.Args {
		if i >= paramTuple.Len() {
			continue
		}

		ident, ok := arg.(*ast.Ident)
		if !ok {
			continue
		}

		shouldComment := false
		switch ident.Name {
		case "nil", "true", "false":
			shouldComment = true
		}

		paramName := paramTuple.At(i).Name()
		expectedContents := "/*" + paramName + "*/"

		nextEndPos := call.Rparen
		if i != end {
			nextEndPos = call.Args[i+1].Pos()
		}
		// TODO: this is wrong, need to back up past the comma

		cIdx, found := findComment(file, posRange{ident.End(), nextEndPos})
		if found {
			comment := file.Comments[cIdx].List[0]
			if !shouldComment {
				pass.Report(analysis.Diagnostic{
					Pos:     comment.Pos(),
					End:     comment.End(),
					Message: "this argument should not be annotated",
				})
				continue
			}

			contents := comment.Text

			if contents != expectedContents {
				pass.Report(analysis.Diagnostic{
					Pos:     comment.Pos(),
					End:     comment.End(),
					Message: fmt.Sprintf("this argument should be annotated with %q", expectedContents),
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: "replace comment with expected",
							TextEdits: []analysis.TextEdit{
								{
									Pos:     comment.Pos(),
									End:     comment.End(),
									NewText: []byte(expectedContents),
								},
							},
						},
					},
				})
			}
			continue
		}

		if !shouldComment {
			continue
		}

		pass.Report(analysis.Diagnostic{
			Pos:     ident.Pos(),
			End:     ident.End(),
			Message: fmt.Sprintf("this argument should be annotated with %q", expectedContents),
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "add comment",
					TextEdits: []analysis.TextEdit{
						{
							Pos:     ident.End(),
							End:     ident.End(),
							NewText: []byte(" " + expectedContents),
						},
					},
				},
			},
		})
	}
}
