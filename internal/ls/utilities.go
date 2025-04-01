package ls

import (
	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/astnav"
	"github.com/microsoft/typescript-go/internal/checker"
)

// !!! Shared (placeholder)
func isInString(file *ast.SourceFile, position int, previousToken *ast.Node) bool {
	if previousToken == nil {
		previousToken = astnav.FindPrecedingToken(file, position)
	}

	if previousToken != nil && isStringTextContainingNode(previousToken) {
		// start := previousToken.
		// !!! HERE
	}

	return false
}

// !!! Shared (placeholder)
func isStringTextContainingNode(node *ast.Node) bool {
	return true
}

// !!! Shared (placeholder)
func getStartOfNode(node *ast.Node, file *ast.SourceFile) int {
	return node.Pos()
}

func tryGetImportFromModuleSpecifier(node *ast.StringLiteralLike) *ast.Node {
	switch node.Parent.Kind {
	case ast.KindImportDeclaration, ast.KindExportDeclaration, ast.KindJSDocImportTag:
		return node.Parent
	case ast.KindExternalModuleReference:
		return node.Parent.Parent
	case ast.KindCallExpression:
		if ast.IsImportCall(node.Parent) || ast.IsRequireCall(node.Parent, false /*requireStringLiteralLikeArgument*/) {
			return node.Parent
		}
		return nil
	case ast.KindLiteralType:
		if !ast.IsStringLiteral(node) {
			return nil
		}
		if ast.IsImportTypeNode(node.Parent.Parent) {
			return node.Parent.Parent
		}
		return nil
	}
	return nil
}

// !!! Shared (placeholder)
func isInComment(file *ast.SourceFile, position int, tokenAtPosition *ast.Node) *ast.CommentRange {
	return nil
}

// // Returns a non-nil comment range if the cursor at position in sourceFile is within a comment.
// // tokenAtPosition: must equal `getTokenAtPosition(sourceFile, position)`
// // predicate: additional predicate to test on the comment range.
// func isInComment(file *ast.SourceFile, position int, tokenAtPosition *ast.Node) *ast.CommentRange {
// 	return getRangeOfEnclosingComment(file, position, nil /*precedingToken*/, tokenAtPosition)
// }

// !!!
// Replaces last(node.getChildren(sourceFile))
func getLastChild(node *ast.Node, sourceFile *ast.SourceFile) *ast.Node {
	return nil
}

func getLastToken(node *ast.Node, sourceFile *ast.SourceFile) *ast.Node {
	assertHasRealPosition(node)

	lastChild := getLastChild(node, sourceFile)
	if lastChild == nil {
		return nil
	}

	if lastChild.Kind < ast.KindFirstNode {
		return lastChild
	} else {
		return getLastToken(lastChild, sourceFile)
	}
}

// !!!
func getFirstToken(node *ast.Node, sourceFile *ast.SourceFile) *ast.Node {
	return nil
}

func assertHasRealPosition(node *ast.Node) {
	if ast.PositionIsSynthesized(node.Pos()) || ast.PositionIsSynthesized(node.End()) {
		panic("Node must have a real position for this operation.")
	}
}

// !!!
func findChildOfKind(node *ast.Node, kind ast.Kind, sourceFile *ast.SourceFile) *ast.Node {
	return nil
}

// !!! Shared: placeholder
type PossibleTypeArgumentInfo struct {
	called         *ast.IdentifierNode
	nTypeArguments int
}

// !!! Shared: placeholder
func getPossibleTypeArgumentsInfo(tokenIn *ast.Node, sourceFile *ast.SourceFile) *PossibleTypeArgumentInfo {
	return nil
}

// !!! Shared: placeholder
func getPossibleGenericSignatures(called *ast.Expression, typeArgumentCount int, checker *checker.Checker) []*checker.Signature {
	return nil
}

func isInRightSideOfInternalImportEqualsDeclaration(node *ast.Node) bool {
	for node.Parent.Kind == ast.KindQualifiedName {
		node = node.Parent
	}

	return ast.IsInternalModuleImportEqualsDeclaration(node.Parent) && node.Parent.AsImportEqualsDeclaration().ModuleReference == node
}
