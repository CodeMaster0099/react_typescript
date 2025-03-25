package ls

import (
	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/astnav"
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/lsp/lsproto"
)

func (l *LanguageService) ProvideCompletion(fileName string, position int) *lsproto.CompletionList {
	program, file := l.getProgramAndFile(fileName)
	node := astnav.GetTouchingPropertyName(file, position)
	if node.Kind == ast.KindSourceFile {
		return nil
	}
	return l.getCompletionsAtPosition(program, file, position)
}

func (l *LanguageService) getCompletionsAtPosition(program *compiler.Program, file *ast.SourceFile, position int) *lsproto.CompletionList {
	// !!! trigger kind

	// !!! see if incomplete completion list

	// !!! string literal completions

	// !!! get completion data

	// !!! transform data into completion list

	return nil // !!!
}

type completionData struct {
	// !!!
}

func getCompletionData(program *compiler.Program, file *ast.SourceFile, position int) *completionData {
	typeChecker := program.GetTypeChecker()

	// !!! why do we use this? just for detecting comments/jsdoc scenario?
	currentToken := astnav.GetTokenAtPosition(file, position)

	// !!! comment, jsdoc stuff

	// !!! get relevant tokens

	// !!! adjust context token

	// !!! see if jsx tag or dot completion

	// !!! global completions

	return nil
}

// !!! Document this
// `contextToken` and `previousToken` can both be nil if we are at the beginning of the file.
func getRelevantTokens(position int, file *ast.SourceFile) (contextToken *ast.Node, previousToken *ast.Node) {

}
