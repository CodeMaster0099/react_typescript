package ast

type SyntaxKind int16

const (
	SyntaxKindUnknown SyntaxKind = iota
	SyntaxKindEndOfFile
	SyntaxKindConflictMarkerTrivia
	SyntaxKindNonTextFileMarkerTrivia
	SyntaxKindNumericLiteral
	SyntaxKindBigIntLiteral
	SyntaxKindStringLiteral
	SyntaxKindJsxText
	SyntaxKindJsxTextAllWhiteSpaces
	SyntaxKindRegularExpressionLiteral
	SyntaxKindNoSubstitutionTemplateLiteral
	SyntaxKindTemplateHead
	SyntaxKindTemplateMiddle
	SyntaxKindTemplateTail
	SyntaxKindOpenBraceToken
	SyntaxKindCloseBraceToken
	SyntaxKindOpenParenToken
	SyntaxKindCloseParenToken
	SyntaxKindOpenBracketToken
	SyntaxKindCloseBracketToken
	SyntaxKindDotToken
	SyntaxKindDotDotDotToken
	SyntaxKindSemicolonToken
	SyntaxKindCommaToken
	SyntaxKindQuestionDotToken
	SyntaxKindLessThanToken
	SyntaxKindLessThanSlashToken
	SyntaxKindGreaterThanToken
	SyntaxKindLessThanEqualsToken
	SyntaxKindGreaterThanEqualsToken
	SyntaxKindEqualsEqualsToken
	SyntaxKindExclamationEqualsToken
	SyntaxKindEqualsEqualsEqualsToken
	SyntaxKindExclamationEqualsEqualsToken
	SyntaxKindEqualsGreaterThanToken
	SyntaxKindPlusToken
	SyntaxKindMinusToken
	SyntaxKindAsteriskToken
	SyntaxKindAsteriskAsteriskToken
	SyntaxKindSlashToken
	SyntaxKindPercentToken
	SyntaxKindPlusPlusToken
	SyntaxKindMinusMinusToken
	SyntaxKindLessThanLessThanToken
	SyntaxKindGreaterThanGreaterThanToken
	SyntaxKindGreaterThanGreaterThanGreaterThanToken
	SyntaxKindAmpersandToken
	SyntaxKindBarToken
	SyntaxKindCaretToken
	SyntaxKindExclamationToken
	SyntaxKindTildeToken
	SyntaxKindAmpersandAmpersandToken
	SyntaxKindBarBarToken
	SyntaxKindQuestionToken
	SyntaxKindColonToken
	SyntaxKindAtToken
	SyntaxKindQuestionQuestionToken
	SyntaxKindBacktickToken
	SyntaxKindHashToken
	SyntaxKindEqualsToken
	SyntaxKindPlusEqualsToken
	SyntaxKindMinusEqualsToken
	SyntaxKindAsteriskEqualsToken
	SyntaxKindAsteriskAsteriskEqualsToken
	SyntaxKindSlashEqualsToken
	SyntaxKindPercentEqualsToken
	SyntaxKindLessThanLessThanEqualsToken
	SyntaxKindGreaterThanGreaterThanEqualsToken
	SyntaxKindGreaterThanGreaterThanGreaterThanEqualsToken
	SyntaxKindAmpersandEqualsToken
	SyntaxKindBarEqualsToken
	SyntaxKindBarBarEqualsToken
	SyntaxKindAmpersandAmpersandEqualsToken
	SyntaxKindQuestionQuestionEqualsToken
	SyntaxKindCaretEqualsToken
	SyntaxKindIdentifier
	SyntaxKindPrivateIdentifier
	SyntaxKindJSDocCommentTextToken
	SyntaxKindBreakKeyword
	SyntaxKindCaseKeyword
	SyntaxKindCatchKeyword
	SyntaxKindClassKeyword
	SyntaxKindConstKeyword
	SyntaxKindContinueKeyword
	SyntaxKindDebuggerKeyword
	SyntaxKindDefaultKeyword
	SyntaxKindDeleteKeyword
	SyntaxKindDoKeyword
	SyntaxKindElseKeyword
	SyntaxKindEnumKeyword
	SyntaxKindExportKeyword
	SyntaxKindExtendsKeyword
	SyntaxKindFalseKeyword
	SyntaxKindFinallyKeyword
	SyntaxKindForKeyword
	SyntaxKindFunctionKeyword
	SyntaxKindIfKeyword
	SyntaxKindImportKeyword
	SyntaxKindInKeyword
	SyntaxKindInstanceOfKeyword
	SyntaxKindNewKeyword
	SyntaxKindNullKeyword
	SyntaxKindReturnKeyword
	SyntaxKindSuperKeyword
	SyntaxKindSwitchKeyword
	SyntaxKindThisKeyword
	SyntaxKindThrowKeyword
	SyntaxKindTrueKeyword
	SyntaxKindTryKeyword
	SyntaxKindTypeOfKeyword
	SyntaxKindVarKeyword
	SyntaxKindVoidKeyword
	SyntaxKindWhileKeyword
	SyntaxKindWithKeyword
	SyntaxKindImplementsKeyword
	SyntaxKindInterfaceKeyword
	SyntaxKindLetKeyword
	SyntaxKindPackageKeyword
	SyntaxKindPrivateKeyword
	SyntaxKindProtectedKeyword
	SyntaxKindPublicKeyword
	SyntaxKindStaticKeyword
	SyntaxKindYieldKeyword
	SyntaxKindAbstractKeyword
	SyntaxKindAccessorKeyword
	SyntaxKindAsKeyword
	SyntaxKindAssertsKeyword
	SyntaxKindAssertKeyword
	SyntaxKindAnyKeyword
	SyntaxKindAsyncKeyword
	SyntaxKindAwaitKeyword
	SyntaxKindBooleanKeyword
	SyntaxKindConstructorKeyword
	SyntaxKindDeclareKeyword
	SyntaxKindGetKeyword
	SyntaxKindImmediateKeyword
	SyntaxKindInferKeyword
	SyntaxKindIntrinsicKeyword
	SyntaxKindIsKeyword
	SyntaxKindKeyOfKeyword
	SyntaxKindModuleKeyword
	SyntaxKindNamespaceKeyword
	SyntaxKindNeverKeyword
	SyntaxKindOutKeyword
	SyntaxKindReadonlyKeyword
	SyntaxKindRequireKeyword
	SyntaxKindNumberKeyword
	SyntaxKindObjectKeyword
	SyntaxKindSatisfiesKeyword
	SyntaxKindSetKeyword
	SyntaxKindStringKeyword
	SyntaxKindSymbolKeyword
	SyntaxKindTypeKeyword
	SyntaxKindUndefinedKeyword
	SyntaxKindUniqueKeyword
	SyntaxKindUnknownKeyword
	SyntaxKindUsingKeyword
	SyntaxKindFromKeyword
	SyntaxKindGlobalKeyword
	SyntaxKindBigIntKeyword
	SyntaxKindOverrideKeyword
	SyntaxKindOfKeyword

	SyntaxKindCount

	AssignmentStart         = SyntaxKindEqualsToken
	AssignmentEnd           = SyntaxKindCaretEqualsToken
	CompoundAssignmentStart = SyntaxKindPlusEqualsToken
	CompoundAssignmentEnd   = SyntaxKindCaretEqualsToken
	ReservedWordStart       = SyntaxKindBreakKeyword
	ReservedWordEnd         = SyntaxKindWithKeyword
	KeywordStart            = SyntaxKindBreakKeyword
	KeywordEnd              = SyntaxKindOfKeyword
	FutureReservedWordStart = SyntaxKindImplementsKeyword
	FutureReservedWordEnd   = SyntaxKindYieldKeyword
	PunctuationStart        = SyntaxKindOpenBraceToken
	PunctuationEnd          = SyntaxKindCaretEqualsToken
	TokenStart              = SyntaxKindUnknown
	TokenEnd                = SyntaxKindOfKeyword
	LiteralTokenStart       = SyntaxKindNumericLiteral
	LiteralTokenEnd         = SyntaxKindNoSubstitutionTemplateLiteral
	TemplateTokenStart      = SyntaxKindNoSubstitutionTemplateLiteral
	TemplateTokenEnd        = SyntaxKindTemplateTail
)

type NodeFlags uint32
type NodeID uint32

type TextRange struct {
	Pos int32
	End int32
}

type Node struct {
	kind   SyntaxKind
	flags  NodeFlags
	loc    TextRange
	id     NodeID
	parent *Node
	data   NodeData
}

type NumericLiteral struct {
	NodeBase
}

func (n *Node) AsNumericLiteral() *NumericLiteral { return n.data.(*NumericLiteral) }

func (n *NumericLiteral) reset() {
	*n = NumericLiteral{}
}

func (n *NumericLiteral) Kind() SyntaxKind { return SyntaxKindNumericLiteral }

func NewNumericLiteral() *NumericLiteral {
	v := &NumericLiteral{}
	v.reset()
	return v
}

func (f *Factory) NewNumericLiteral() *NumericLiteral {
	v := &NumericLiteral{}
	v.reset()
	return v
}

type BigIntLiteral struct {
	NodeBase
}

func (n *Node) AsBigIntLiteral() *BigIntLiteral { return n.data.(*BigIntLiteral) }

func (n *BigIntLiteral) reset() {
	*n = BigIntLiteral{}
}

func (n *BigIntLiteral) Kind() SyntaxKind { return SyntaxKindBigIntLiteral }

func NewBigIntLiteral() *BigIntLiteral {
	v := &BigIntLiteral{}
	v.reset()
	return v
}

func (f *Factory) NewBigIntLiteral() *BigIntLiteral {
	v := &BigIntLiteral{}
	v.reset()
	return v
}

type StringLiteral struct {
	NodeBase
}

func (n *Node) AsStringLiteral() *StringLiteral { return n.data.(*StringLiteral) }

func (n *StringLiteral) reset() {
	*n = StringLiteral{}
}

func (n *StringLiteral) Kind() SyntaxKind { return SyntaxKindStringLiteral }

func NewStringLiteral() *StringLiteral {
	v := &StringLiteral{}
	v.reset()
	return v
}

func (f *Factory) NewStringLiteral() *StringLiteral {
	v := &StringLiteral{}
	v.reset()
	return v
}

type JsxText struct {
	NodeBase
}

func (n *Node) AsJsxText() *JsxText { return n.data.(*JsxText) }

func (n *JsxText) reset() {
	*n = JsxText{}
}

func (n *JsxText) Kind() SyntaxKind { return SyntaxKindJsxText }

func NewJsxText() *JsxText {
	v := &JsxText{}
	v.reset()
	return v
}

func (f *Factory) NewJsxText() *JsxText {
	v := &JsxText{}
	v.reset()
	return v
}

type RegularExpressionLiteral struct {
	NodeBase
}

func (n *Node) AsRegularExpressionLiteral() *RegularExpressionLiteral {
	return n.data.(*RegularExpressionLiteral)
}

func (n *RegularExpressionLiteral) reset() {
	*n = RegularExpressionLiteral{}
}

func (n *RegularExpressionLiteral) Kind() SyntaxKind { return SyntaxKindRegularExpressionLiteral }

func NewRegularExpressionLiteral() *RegularExpressionLiteral {
	v := &RegularExpressionLiteral{}
	v.reset()
	return v
}

func (f *Factory) NewRegularExpressionLiteral() *RegularExpressionLiteral {
	v := &RegularExpressionLiteral{}
	v.reset()
	return v
}

type NoSubstitutionTemplateLiteral struct {
	NodeBase
}

func (n *Node) AsNoSubstitutionTemplateLiteral() *NoSubstitutionTemplateLiteral {
	return n.data.(*NoSubstitutionTemplateLiteral)
}

func (n *NoSubstitutionTemplateLiteral) reset() {
	*n = NoSubstitutionTemplateLiteral{}
}

func (n *NoSubstitutionTemplateLiteral) Kind() SyntaxKind {
	return SyntaxKindNoSubstitutionTemplateLiteral
}

func NewNoSubstitutionTemplateLiteral() *NoSubstitutionTemplateLiteral {
	v := &NoSubstitutionTemplateLiteral{}
	v.reset()
	return v
}

func (f *Factory) NewNoSubstitutionTemplateLiteral() *NoSubstitutionTemplateLiteral {
	v := &NoSubstitutionTemplateLiteral{}
	v.reset()
	return v
}

type TemplateHead struct {
	NodeBase
}

func (n *Node) AsTemplateHead() *TemplateHead { return n.data.(*TemplateHead) }

func (n *TemplateHead) reset() {
	*n = TemplateHead{}
}

func (n *TemplateHead) Kind() SyntaxKind { return SyntaxKindTemplateHead }

func NewTemplateHead() *TemplateHead {
	v := &TemplateHead{}
	v.reset()
	return v
}

func (f *Factory) NewTemplateHead() *TemplateHead {
	v := &TemplateHead{}
	v.reset()
	return v
}

type TemplateMiddle struct {
	NodeBase
}

func (n *Node) AsTemplateMiddle() *TemplateMiddle { return n.data.(*TemplateMiddle) }

func (n *TemplateMiddle) reset() {
	*n = TemplateMiddle{}
}

func (n *TemplateMiddle) Kind() SyntaxKind { return SyntaxKindTemplateMiddle }

func NewTemplateMiddle() *TemplateMiddle {
	v := &TemplateMiddle{}
	v.reset()
	return v
}

func (f *Factory) NewTemplateMiddle() *TemplateMiddle {
	v := &TemplateMiddle{}
	v.reset()
	return v
}

type TemplateTail struct {
	NodeBase
}

func (n *Node) AsTemplateTail() *TemplateTail { return n.data.(*TemplateTail) }

func (n *TemplateTail) reset() {
	*n = TemplateTail{}
}

func (n *TemplateTail) Kind() SyntaxKind { return SyntaxKindTemplateTail }

func NewTemplateTail() *TemplateTail {
	v := &TemplateTail{}
	v.reset()
	return v
}

func (f *Factory) NewTemplateTail() *TemplateTail {
	v := &TemplateTail{}
	v.reset()
	return v
}

type Identifier struct {
	NodeBase
}

func (n *Node) AsIdentifier() *Identifier { return n.data.(*Identifier) }

func (n *Identifier) reset() {
	*n = Identifier{}
}

func (n *Identifier) Kind() SyntaxKind { return SyntaxKindIdentifier }

func NewIdentifier() *Identifier {
	v := &Identifier{}
	v.reset()
	return v
}

func (f *Factory) NewIdentifier() *Identifier {
	v := f.IdentifierPool.Get()
	v.reset()
	return v
}

type PrivateIdentifier struct {
	NodeBase
}

func (n *Node) AsPrivateIdentifier() *PrivateIdentifier { return n.data.(*PrivateIdentifier) }

func (n *PrivateIdentifier) reset() {
	*n = PrivateIdentifier{}
}

func (n *PrivateIdentifier) Kind() SyntaxKind { return SyntaxKindPrivateIdentifier }

func NewPrivateIdentifier() *PrivateIdentifier {
	v := &PrivateIdentifier{}
	v.reset()
	return v
}

func (f *Factory) NewPrivateIdentifier() *PrivateIdentifier {
	v := &PrivateIdentifier{}
	v.reset()
	return v
}
