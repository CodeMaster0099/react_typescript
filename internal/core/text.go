package core

// TextPos

type TextPos int32

// TextRange

type TextRange struct {
	pos TextPos
	end TextPos
}

func NewTextRange(pos int, end int) TextRange {
	return TextRange{pos: TextPos(pos), end: TextPos(end)}
}

func (t TextRange) Pos() TextPos {
	return t.pos
}

func (t TextRange) End() TextPos {
	return t.end
}

func (t TextRange) Len() int {
	return int(t.end - t.pos)
}

func (t TextRange) ContainsInclusive(pos int) bool {
	return pos >= int(t.pos) && pos <= int(t.end)
}
