//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var d declarer

	firstToken := d.newWithoutNode("Unknown")
	d.newWithoutNode("EndOfFile")
	d.newWithoutNode("ConflictMarkerTrivia")
	d.newWithoutNode("NonTextFileMarkerTrivia")

	firstLiteralToken := d.newNode("NumericLiteral", &nodeOptions{})
	d.newNode("BigIntLiteral", &nodeOptions{})
	d.newNode("StringLiteral", &nodeOptions{})
	d.newNode("JsxText", &nodeOptions{})
	d.newWithoutNode("JsxTextAllWhiteSpaces")
	d.newNode("RegularExpressionLiteral", &nodeOptions{})
	lastLiteralToken := d.newNode("NoSubstitutionTemplateLiteral", &nodeOptions{})

	// Pseudo-literals
	firstTemplateToken := lastLiteralToken
	d.newNode("TemplateHead", &nodeOptions{})
	d.newNode("TemplateMiddle", &nodeOptions{})
	lastTemplateToken := d.newNode("TemplateTail", &nodeOptions{})

	firstPunctuation := d.newWithoutNode("OpenBraceToken")
	d.newWithoutNode("CloseBraceToken")
	d.newWithoutNode("OpenParenToken")
	d.newWithoutNode("CloseParenToken")
	d.newWithoutNode("OpenBracketToken")
	d.newWithoutNode("CloseBracketToken")
	d.newWithoutNode("DotToken")
	d.newWithoutNode("DotDotDotToken")
	d.newWithoutNode("SemicolonToken")
	d.newWithoutNode("CommaToken")
	d.newWithoutNode("QuestionDotToken")
	d.newWithoutNode("LessThanToken")
	d.newWithoutNode("LessThanSlashToken")
	d.newWithoutNode("GreaterThanToken")
	d.newWithoutNode("LessThanEqualsToken")
	d.newWithoutNode("GreaterThanEqualsToken")
	d.newWithoutNode("EqualsEqualsToken")
	d.newWithoutNode("ExclamationEqualsToken")
	d.newWithoutNode("EqualsEqualsEqualsToken")
	d.newWithoutNode("ExclamationEqualsEqualsToken")
	d.newWithoutNode("EqualsGreaterThanToken")
	d.newWithoutNode("PlusToken")
	d.newWithoutNode("MinusToken")
	d.newWithoutNode("AsteriskToken")
	d.newWithoutNode("AsteriskAsteriskToken")
	d.newWithoutNode("SlashToken")
	d.newWithoutNode("PercentToken")
	d.newWithoutNode("PlusPlusToken")
	d.newWithoutNode("MinusMinusToken")
	d.newWithoutNode("LessThanLessThanToken")
	d.newWithoutNode("GreaterThanGreaterThanToken")
	d.newWithoutNode("GreaterThanGreaterThanGreaterThanToken")
	d.newWithoutNode("AmpersandToken")
	d.newWithoutNode("BarToken")
	d.newWithoutNode("CaretToken")
	d.newWithoutNode("ExclamationToken")
	d.newWithoutNode("TildeToken")
	d.newWithoutNode("AmpersandAmpersandToken")
	d.newWithoutNode("BarBarToken")
	d.newWithoutNode("QuestionToken")
	d.newWithoutNode("ColonToken")
	d.newWithoutNode("AtToken")
	d.newWithoutNode("QuestionQuestionToken")
	d.newWithoutNode("BacktickToken")
	d.newWithoutNode("HashToken")
	firstAssignment := d.newWithoutNode("EqualsToken")
	firstCompoundAssignment := d.newWithoutNode("PlusEqualsToken")
	d.newWithoutNode("MinusEqualsToken")
	d.newWithoutNode("AsteriskEqualsToken")
	d.newWithoutNode("AsteriskAsteriskEqualsToken")
	d.newWithoutNode("SlashEqualsToken")
	d.newWithoutNode("PercentEqualsToken")
	d.newWithoutNode("LessThanLessThanEqualsToken")
	d.newWithoutNode("GreaterThanGreaterThanEqualsToken")
	d.newWithoutNode("GreaterThanGreaterThanGreaterThanEqualsToken")
	d.newWithoutNode("AmpersandEqualsToken")
	d.newWithoutNode("BarEqualsToken")
	d.newWithoutNode("BarBarEqualsToken")
	d.newWithoutNode("AmpersandAmpersandEqualsToken")
	d.newWithoutNode("QuestionQuestionEqualsToken")
	lastPunctuation := d.newWithoutNode("CaretEqualsToken")
	lastAssignment := lastPunctuation
	lastCompoundAssignment := lastAssignment

	d.newNode("Identifier", &nodeOptions{})
	d.newNode("PrivateIdentifier", &nodeOptions{})
	d.newWithoutNode("JSDocCommentTextToken") // TODO: why is this here??

	firstReservedWord := d.newWithoutNode("BreakKeyword")
	// Reserved words
	firstKeyword := firstReservedWord
	d.newWithoutNode("CaseKeyword")
	d.newWithoutNode("CatchKeyword")
	d.newWithoutNode("ClassKeyword")
	d.newWithoutNode("ConstKeyword")
	d.newWithoutNode("ContinueKeyword")
	d.newWithoutNode("DebuggerKeyword")
	d.newWithoutNode("DefaultKeyword")
	d.newWithoutNode("DeleteKeyword")
	d.newWithoutNode("DoKeyword")
	d.newWithoutNode("ElseKeyword")
	d.newWithoutNode("EnumKeyword")
	d.newWithoutNode("ExportKeyword")
	d.newWithoutNode("ExtendsKeyword")
	d.newWithoutNode("FalseKeyword")
	d.newWithoutNode("FinallyKeyword")
	d.newWithoutNode("ForKeyword")
	d.newWithoutNode("FunctionKeyword")
	d.newWithoutNode("IfKeyword")
	d.newWithoutNode("ImportKeyword")
	d.newWithoutNode("InKeyword")
	d.newWithoutNode("InstanceOfKeyword")
	d.newWithoutNode("NewKeyword")
	d.newWithoutNode("NullKeyword")
	d.newWithoutNode("ReturnKeyword")
	d.newWithoutNode("SuperKeyword")
	d.newWithoutNode("SwitchKeyword")
	d.newWithoutNode("ThisKeyword")
	d.newWithoutNode("ThrowKeyword")
	d.newWithoutNode("TrueKeyword")
	d.newWithoutNode("TryKeyword")
	d.newWithoutNode("TypeOfKeyword")
	d.newWithoutNode("VarKeyword")
	d.newWithoutNode("VoidKeyword")
	d.newWithoutNode("WhileKeyword")
	lastReservedWord := d.newWithoutNode("WithKeyword")

	// Strict mode reserved words
	firstFutureReservedWord := d.newWithoutNode("ImplementsKeyword")
	d.newWithoutNode("InterfaceKeyword")
	d.newWithoutNode("LetKeyword")
	d.newWithoutNode("PackageKeyword")
	d.newWithoutNode("PrivateKeyword")
	d.newWithoutNode("ProtectedKeyword")
	d.newWithoutNode("PublicKeyword")
	d.newWithoutNode("StaticKeyword")
	lastFutureReservedWord := d.newWithoutNode("YieldKeyword")

	// Contextual keywords
	d.newWithoutNode("AbstractKeyword")
	d.newWithoutNode("AccessorKeyword")
	d.newWithoutNode("AsKeyword")
	d.newWithoutNode("AssertsKeyword")
	d.newWithoutNode("AssertKeyword")
	d.newWithoutNode("AnyKeyword")
	d.newWithoutNode("AsyncKeyword")
	d.newWithoutNode("AwaitKeyword")
	d.newWithoutNode("BooleanKeyword")
	d.newWithoutNode("ConstructorKeyword")
	d.newWithoutNode("DeclareKeyword")
	d.newWithoutNode("GetKeyword")
	d.newWithoutNode("ImmediateKeyword")
	d.newWithoutNode("InferKeyword")
	d.newWithoutNode("IntrinsicKeyword")
	d.newWithoutNode("IsKeyword")
	d.newWithoutNode("KeyOfKeyword")
	d.newWithoutNode("ModuleKeyword")
	d.newWithoutNode("NamespaceKeyword")
	d.newWithoutNode("NeverKeyword")
	d.newWithoutNode("OutKeyword")
	d.newWithoutNode("ReadonlyKeyword")
	d.newWithoutNode("RequireKeyword")
	d.newWithoutNode("NumberKeyword")
	d.newWithoutNode("ObjectKeyword")
	d.newWithoutNode("SatisfiesKeyword")
	d.newWithoutNode("SetKeyword")
	d.newWithoutNode("StringKeyword")
	d.newWithoutNode("SymbolKeyword")
	d.newWithoutNode("TypeKeyword")
	d.newWithoutNode("UndefinedKeyword")
	d.newWithoutNode("UniqueKeyword")
	d.newWithoutNode("UnknownKeyword")
	d.newWithoutNode("UsingKeyword")
	d.newWithoutNode("FromKeyword")
	d.newWithoutNode("GlobalKeyword")
	d.newWithoutNode("BigIntKeyword")
	d.newWithoutNode("OverrideKeyword")
	lastKeyword := d.newWithoutNode("OfKeyword")
	lastToken := lastKeyword

	// Parse tree nodes
	// Names

	d.createGroup("Assignment", firstAssignment, lastAssignment)
	d.createGroup("CompoundAssignment", firstCompoundAssignment, lastCompoundAssignment)
	d.createGroup("ReservedWord", firstReservedWord, lastReservedWord)
	d.createGroup("Keyword", firstKeyword, lastKeyword)
	d.createGroup("FutureReservedWord", firstFutureReservedWord, lastFutureReservedWord)
	// d.createGroup("TypeNode", firstTypeNode, lastTypeNode)
	d.createGroup("Punctuation", firstPunctuation, lastPunctuation)
	d.createGroup("Token", firstToken, lastToken)
	d.createGroup("LiteralToken", firstLiteralToken, lastLiteralToken)
	d.createGroup("TemplateToken", firstTemplateToken, lastTemplateToken)
	// d.createGroup("BinaryOperator", firstBinaryOperator, lastBinaryOperator)
	// d.createGroup("Node", firstNode, d.nodes[len(d.nodes)-1])
	// d.createGroup("JSDocNode", firstJSDocNode, lastJSDocNode)
	// d.createGroup("JSDocTagNode", firstJSDocTagNode, lastJSDocTagNode)
	// d.createGroup("ContextualKeyword", firstContextualKeyword, lastContextualKeyword)

	fmt.Println("type SyntaxKind int64")
	fmt.Println()
	fmt.Println("const (")

	for _, n := range d.nodes {
		if n.first {
			fmt.Printf("\tSyntaxKind%s SyntaxKind = iota\n", n.name)
		} else {
			fmt.Printf("\tSyntaxKind%s\n", n.name)
		}
	}

	fmt.Println()
	fmt.Println("\tSyntaxKindCount")
	fmt.Println()

	for _, g := range d.groups {
		fmt.Printf("\t%sStart = SyntaxKind%s\n", g.name, g.start.name)
		fmt.Printf("\t%sEnd = SyntaxKind%s\n", g.name, g.end.name)
	}

	fmt.Println(")")

	// TODO: Count

}

type declarer struct {
	nodes       []*SyntaxKind
	nodesByName map[string]*SyntaxKind
	groups      []*group
}

type group struct {
	name  string
	start *SyntaxKind
	end   *SyntaxKind
}

func (d *declarer) createGroup(name string, start, end *SyntaxKind) {
	d.groups = append(d.groups, &group{name: name, start: start, end: end})
}

type nodeOptions struct {
}

func (d *declarer) newWithoutNode(name string) *SyntaxKind {
	return d.newNode(name, nil)
}

func (d *declarer) newNode(name string, opts *nodeOptions) *SyntaxKind {
	if d.nodesByName == nil {
		d.nodesByName = make(map[string]*SyntaxKind)
	}

	n := &SyntaxKind{
		first: len(d.nodes) == 0,
		name:  name,
	}
	d.nodes = append(d.nodes, n)
	d.nodesByName[name] = n
	return n
}

type Type interface {
	Name() string
}

type SyntaxKind struct {
	first    bool
	name     string
	children []Child
}

func (n *SyntaxKind) Name() string { return n.name }

type Child struct {
	name string
	typ  Type
}
