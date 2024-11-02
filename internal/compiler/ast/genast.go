//go:build ignore

package main

import (
	"fmt"
)

func main() {
	var d declarer

	firstToken := d.declareKindOnly("Unknown")
	d.declareKindOnly("EndOfFile")
	d.declareKindOnly("ConflictMarkerTrivia")
	d.declareKindOnly("NonTextFileMarkerTrivia")

	firstLiteralToken := d.declare("NumericLiteral", &declOptions{})
	d.declare("BigIntLiteral", &declOptions{})
	d.declare("StringLiteral", &declOptions{})
	d.declare("JsxText", &declOptions{})
	d.declareKindOnly("JsxTextAllWhiteSpaces")
	d.declare("RegularExpressionLiteral", &declOptions{})
	lastLiteralToken := d.declare("NoSubstitutionTemplateLiteral", &declOptions{})

	// Pseudo-literals
	firstTemplateToken := lastLiteralToken
	d.declare("TemplateHead", &declOptions{})
	d.declare("TemplateMiddle", &declOptions{})
	lastTemplateToken := d.declare("TemplateTail", &declOptions{})

	firstPunctuation := d.declareKindOnly("OpenBraceToken")
	d.declareKindOnly("CloseBraceToken")
	d.declareKindOnly("OpenParenToken")
	d.declareKindOnly("CloseParenToken")
	d.declareKindOnly("OpenBracketToken")
	d.declareKindOnly("CloseBracketToken")
	d.declareKindOnly("DotToken")
	d.declareKindOnly("DotDotDotToken")
	d.declareKindOnly("SemicolonToken")
	d.declareKindOnly("CommaToken")
	d.declareKindOnly("QuestionDotToken")
	d.declareKindOnly("LessThanToken")
	d.declareKindOnly("LessThanSlashToken")
	d.declareKindOnly("GreaterThanToken")
	d.declareKindOnly("LessThanEqualsToken")
	d.declareKindOnly("GreaterThanEqualsToken")
	d.declareKindOnly("EqualsEqualsToken")
	d.declareKindOnly("ExclamationEqualsToken")
	d.declareKindOnly("EqualsEqualsEqualsToken")
	d.declareKindOnly("ExclamationEqualsEqualsToken")
	d.declareKindOnly("EqualsGreaterThanToken")
	d.declareKindOnly("PlusToken")
	d.declareKindOnly("MinusToken")
	d.declareKindOnly("AsteriskToken")
	d.declareKindOnly("AsteriskAsteriskToken")
	d.declareKindOnly("SlashToken")
	d.declareKindOnly("PercentToken")
	d.declareKindOnly("PlusPlusToken")
	d.declareKindOnly("MinusMinusToken")
	d.declareKindOnly("LessThanLessThanToken")
	d.declareKindOnly("GreaterThanGreaterThanToken")
	d.declareKindOnly("GreaterThanGreaterThanGreaterThanToken")
	d.declareKindOnly("AmpersandToken")
	d.declareKindOnly("BarToken")
	d.declareKindOnly("CaretToken")
	d.declareKindOnly("ExclamationToken")
	d.declareKindOnly("TildeToken")
	d.declareKindOnly("AmpersandAmpersandToken")
	d.declareKindOnly("BarBarToken")
	d.declareKindOnly("QuestionToken")
	d.declareKindOnly("ColonToken")
	d.declareKindOnly("AtToken")
	d.declareKindOnly("QuestionQuestionToken")
	d.declareKindOnly("BacktickToken")
	d.declareKindOnly("HashToken")
	firstAssignment := d.declareKindOnly("EqualsToken")
	firstCompoundAssignment := d.declareKindOnly("PlusEqualsToken")
	d.declareKindOnly("MinusEqualsToken")
	d.declareKindOnly("AsteriskEqualsToken")
	d.declareKindOnly("AsteriskAsteriskEqualsToken")
	d.declareKindOnly("SlashEqualsToken")
	d.declareKindOnly("PercentEqualsToken")
	d.declareKindOnly("LessThanLessThanEqualsToken")
	d.declareKindOnly("GreaterThanGreaterThanEqualsToken")
	d.declareKindOnly("GreaterThanGreaterThanGreaterThanEqualsToken")
	d.declareKindOnly("AmpersandEqualsToken")
	d.declareKindOnly("BarEqualsToken")
	d.declareKindOnly("BarBarEqualsToken")
	d.declareKindOnly("AmpersandAmpersandEqualsToken")
	d.declareKindOnly("QuestionQuestionEqualsToken")
	lastPunctuation := d.declareKindOnly("CaretEqualsToken")
	lastAssignment := lastPunctuation
	lastCompoundAssignment := lastAssignment

	d.declare("Identifier", &declOptions{})
	d.declare("PrivateIdentifier", &declOptions{})
	d.declareKindOnly("JSDocCommentTextToken") // TODO: why is this here??

	firstReservedWord := d.declareKindOnly("BreakKeyword")
	// Reserved words
	firstKeyword := firstReservedWord
	d.declareKindOnly("CaseKeyword")
	d.declareKindOnly("CatchKeyword")
	d.declareKindOnly("ClassKeyword")
	d.declareKindOnly("ConstKeyword")
	d.declareKindOnly("ContinueKeyword")
	d.declareKindOnly("DebuggerKeyword")
	d.declareKindOnly("DefaultKeyword")
	d.declareKindOnly("DeleteKeyword")
	d.declareKindOnly("DoKeyword")
	d.declareKindOnly("ElseKeyword")
	d.declareKindOnly("EnumKeyword")
	d.declareKindOnly("ExportKeyword")
	d.declareKindOnly("ExtendsKeyword")
	d.declareKindOnly("FalseKeyword")
	d.declareKindOnly("FinallyKeyword")
	d.declareKindOnly("ForKeyword")
	d.declareKindOnly("FunctionKeyword")
	d.declareKindOnly("IfKeyword")
	d.declareKindOnly("ImportKeyword")
	d.declareKindOnly("InKeyword")
	d.declareKindOnly("InstanceOfKeyword")
	d.declareKindOnly("NewKeyword")
	d.declareKindOnly("NullKeyword")
	d.declareKindOnly("ReturnKeyword")
	d.declareKindOnly("SuperKeyword")
	d.declareKindOnly("SwitchKeyword")
	d.declareKindOnly("ThisKeyword")
	d.declareKindOnly("ThrowKeyword")
	d.declareKindOnly("TrueKeyword")
	d.declareKindOnly("TryKeyword")
	d.declareKindOnly("TypeOfKeyword")
	d.declareKindOnly("VarKeyword")
	d.declareKindOnly("VoidKeyword")
	d.declareKindOnly("WhileKeyword")
	lastReservedWord := d.declareKindOnly("WithKeyword")

	// Strict mode reserved words
	firstFutureReservedWord := d.declareKindOnly("ImplementsKeyword")
	d.declareKindOnly("InterfaceKeyword")
	d.declareKindOnly("LetKeyword")
	d.declareKindOnly("PackageKeyword")
	d.declareKindOnly("PrivateKeyword")
	d.declareKindOnly("ProtectedKeyword")
	d.declareKindOnly("PublicKeyword")
	d.declareKindOnly("StaticKeyword")
	lastFutureReservedWord := d.declareKindOnly("YieldKeyword")

	// Contextual keywords
	d.declareKindOnly("AbstractKeyword")
	d.declareKindOnly("AccessorKeyword")
	d.declareKindOnly("AsKeyword")
	d.declareKindOnly("AssertsKeyword")
	d.declareKindOnly("AssertKeyword")
	d.declareKindOnly("AnyKeyword")
	d.declareKindOnly("AsyncKeyword")
	d.declareKindOnly("AwaitKeyword")
	d.declareKindOnly("BooleanKeyword")
	d.declareKindOnly("ConstructorKeyword")
	d.declareKindOnly("DeclareKeyword")
	d.declareKindOnly("GetKeyword")
	d.declareKindOnly("ImmediateKeyword")
	d.declareKindOnly("InferKeyword")
	d.declareKindOnly("IntrinsicKeyword")
	d.declareKindOnly("IsKeyword")
	d.declareKindOnly("KeyOfKeyword")
	d.declareKindOnly("ModuleKeyword")
	d.declareKindOnly("NamespaceKeyword")
	d.declareKindOnly("NeverKeyword")
	d.declareKindOnly("OutKeyword")
	d.declareKindOnly("ReadonlyKeyword")
	d.declareKindOnly("RequireKeyword")
	d.declareKindOnly("NumberKeyword")
	d.declareKindOnly("ObjectKeyword")
	d.declareKindOnly("SatisfiesKeyword")
	d.declareKindOnly("SetKeyword")
	d.declareKindOnly("StringKeyword")
	d.declareKindOnly("SymbolKeyword")
	d.declareKindOnly("TypeKeyword")
	d.declareKindOnly("UndefinedKeyword")
	d.declareKindOnly("UniqueKeyword")
	d.declareKindOnly("UnknownKeyword")
	d.declareKindOnly("UsingKeyword")
	d.declareKindOnly("FromKeyword")
	d.declareKindOnly("GlobalKeyword")
	d.declareKindOnly("BigIntKeyword")
	d.declareKindOnly("OverrideKeyword")
	lastKeyword := d.declareKindOnly("OfKeyword")
	lastToken := lastKeyword

	// Parse tree nodes
	// Names

	// Markers

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

	for i, n := range d.nodes {
		if i == 0 {
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

	for _, n := range d.nodes {
		n.Generate()
	}
}

type declarer struct {
	nodes       []*syntaxKind
	nodesByName map[string]*syntaxKind
	groups      []*group
}

type group struct {
	name  string
	start *syntaxKind
	end   *syntaxKind
}

func (d *declarer) createGroup(name string, start, end *syntaxKind) {
	d.groups = append(d.groups, &group{name: name, start: start, end: end})
}

type declOptions struct {
	poolAllocate bool
}

func (d *declarer) declareKindOnly(name string) *syntaxKind {
	return d.declare(name, nil)
}

func (d *declarer) declare(name string, opts *declOptions) *syntaxKind {
	if d.nodesByName == nil {
		d.nodesByName = make(map[string]*syntaxKind)
	}

	n := &syntaxKind{
		name: name,
		opts: opts,
	}
	d.nodes = append(d.nodes, n)
	d.nodesByName[name] = n
	return n
}

type genType interface {
	Name() string
}

type goType struct {
	name string
}

func (t *goType) Name() string { return t.name }

var (
	stringType = &goType{name: "string"}
)

type syntaxKind struct {
	name   string
	opts   *declOptions
	fields []field
}

type field struct {
	name string
	typ  genType
}

func (n *syntaxKind) Name() string { return n.name }

func (n *syntaxKind) Generate() {
	// TODO: stream to output?

	if n.opts == nil {
		return // not a real node; no code
	}

	fmt.Println()
	fmt.Printf("type %s struct {\n", n.name)
	fmt.Println("\tNodeBase")
	// TODO: children
	fmt.Println("}")
	fmt.Println()

	printNewParams := func() {
		// TODO
	}
	printNewArgs := func() {
		// TODO
	}

	// TODO: accept and set children
	fmt.Printf("func (n *%s) reset(", n.name)
	printNewParams()
	fmt.Println(") {")
	// TODO: generate optimal assignment
	fmt.Printf("\t*n = %s{}\n", n.name)
	fmt.Println("}")
	fmt.Println()

	fmt.Printf("func (n *%s) Kind() SyntaxKind { return SyntaxKind%s }\n", n.name, n.name)
	fmt.Println()

	fmt.Printf("func New%s(", n.name)
	printNewParams()
	fmt.Printf(") *%s {\n", n.name)
	fmt.Printf("\tv := &%s{}\n", n.name)
	fmt.Printf("\tv.reset(")
	printNewArgs()
	fmt.Println(")")
	fmt.Println("\treturn v")
	fmt.Println("}")
	fmt.Println()

	// TODO: params
	fmt.Printf("func (f *Factory) New%s() *%s {\n", n.name, n.name)
	if n.opts.poolAllocate {
		fmt.Printf("\tv := f.%sPool.Get()\n", n.name)
	} else {
		fmt.Printf("\tv := &%s{}\n", n.name)
	}
	fmt.Printf("\tv.reset(")
	printNewArgs()
	fmt.Println(")")
	fmt.Println("\treturn v")
	fmt.Println("}")
	fmt.Println()
}
