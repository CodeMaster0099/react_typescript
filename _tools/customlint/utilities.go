package customlint

import (
	"go/ast"
	"go/token"
	"slices"
)

func findComment(file *ast.File, pRange posRange) (int, bool) {
	return slices.BinarySearchFunc(file.Comments, pRange, posRangeCmp)
}

type posRange struct {
	start, end token.Pos
}

func posRangeCmp(c *ast.CommentGroup, target posRange) int {
	if c.End() < target.start {
		return -1
	}
	if c.Pos() >= target.end {
		return 1
	}
	return 0
}
