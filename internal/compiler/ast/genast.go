//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
)

var output = flag.String("output", "", "output file")

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	flag.Parse()

	var buf bytes.Buffer
	g := &generator{w: &buf}
	g.declare()
	g.generate()

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
		formatted = buf.Bytes()
	}

	f := os.Stdout
	if *output != "" {
		var err error
		f, err = os.Create(*output)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		defer f.Close()
	}

	f.Write(formatted)
	return exitCode
}

type generator struct {
	w io.Writer

	syntaxKinds []*syntaxKind
	unions      []*nodeUnion
	markers     []*marker

	firstToken *syntaxKind
}

func (g *generator) print(args ...any) {
	fmt.Fprint(g.w, args...)
}

func (g *generator) println(args ...any) {
	fmt.Fprintln(g.w, args...)
}

func (g *generator) printf(format string, args ...any) {
	fmt.Fprintf(g.w, format, args...)
}

func (g *generator) declare() {
	firstToken := g.declareToken("Unknown")
	g.firstToken = firstToken
	g.declareToken("EndOfFile")
	g.declareToken("ConflictMarkerTrivia")
	g.declareToken("NonTextFileMarkerTrivia")

	firstLiteralToken := g.declareNode("NumericLiteral", &nodeOpts{})
	g.declareNode("BigIntLiteral", &nodeOpts{})
	g.declareNode("StringLiteral", &nodeOpts{})
	g.declareNode("JsxText", &nodeOpts{})
	g.declareToken("JsxTextAllWhiteSpaces")
	g.declareNode("RegularExpressionLiteral", &nodeOpts{})
	lastLiteralToken := g.declareNode("NoSubstitutionTemplateLiteral", &nodeOpts{})

	// Pseudo-literals
	firstTemplateToken := lastLiteralToken
	g.declareNode("TemplateHead", &nodeOpts{})
	g.declareNode("TemplateMiddle", &nodeOpts{})
	lastTemplateToken := g.declareNode("TemplateTail", &nodeOpts{})

	firstPunctuation := g.declareToken("OpenBraceToken")
	g.declareToken("CloseBraceToken")
	g.declareToken("OpenParenToken")
	g.declareToken("CloseParenToken")
	g.declareToken("OpenBracketToken")
	g.declareToken("CloseBracketToken")
	g.declareToken("DotToken")
	g.declareToken("DotDotDotToken")
	g.declareToken("SemicolonToken")
	g.declareToken("CommaToken")
	g.declareToken("QuestionDotToken")
	firstBinaryOperator := g.declareToken("LessThanToken")
	g.declareToken("LessThanSlashToken")
	g.declareToken("GreaterThanToken")
	g.declareToken("LessThanEqualsToken")
	g.declareToken("GreaterThanEqualsToken")
	g.declareToken("EqualsEqualsToken")
	g.declareToken("ExclamationEqualsToken")
	g.declareToken("EqualsEqualsEqualsToken")
	g.declareToken("ExclamationEqualsEqualsToken")
	g.declareToken("EqualsGreaterThanToken")
	g.declareToken("PlusToken")
	g.declareToken("MinusToken")
	g.declareToken("AsteriskToken")
	g.declareToken("AsteriskAsteriskToken")
	g.declareToken("SlashToken")
	g.declareToken("PercentToken")
	g.declareToken("PlusPlusToken")
	g.declareToken("MinusMinusToken")
	g.declareToken("LessThanLessThanToken")
	g.declareToken("GreaterThanGreaterThanToken")
	g.declareToken("GreaterThanGreaterThanGreaterThanToken")
	g.declareToken("AmpersandToken")
	g.declareToken("BarToken")
	g.declareToken("CaretToken")
	g.declareToken("ExclamationToken")
	g.declareToken("TildeToken")
	g.declareToken("AmpersandAmpersandToken")
	g.declareToken("BarBarToken")
	g.declareToken("QuestionToken")
	g.declareToken("ColonToken")
	g.declareToken("AtToken")
	g.declareToken("QuestionQuestionToken")
	g.declareToken("BacktickToken")
	g.declareToken("HashToken")
	firstAssignment := g.declareToken("EqualsToken")
	firstCompoundAssignment := g.declareToken("PlusEqualsToken")
	g.declareToken("MinusEqualsToken")
	g.declareToken("AsteriskEqualsToken")
	g.declareToken("AsteriskAsteriskEqualsToken")
	g.declareToken("SlashEqualsToken")
	g.declareToken("PercentEqualsToken")
	g.declareToken("LessThanLessThanEqualsToken")
	g.declareToken("GreaterThanGreaterThanEqualsToken")
	g.declareToken("GreaterThanGreaterThanGreaterThanEqualsToken")
	g.declareToken("AmpersandEqualsToken")
	g.declareToken("BarEqualsToken")
	g.declareToken("BarBarEqualsToken")
	g.declareToken("AmpersandAmpersandEqualsToken")
	g.declareToken("QuestionQuestionEqualsToken")
	lastPunctuation := g.declareToken("CaretEqualsToken")
	lastAssignment := lastPunctuation
	lastCompoundAssignment := lastPunctuation
	lastBinaryOperator := lastPunctuation

	g.declareNode("Identifier", &nodeOpts{
		poolAllocate: true,
	})
	g.declareNode("PrivateIdentifier", &nodeOpts{})
	g.declareToken("JSDocCommentTextToken") // TODO: why is this here??

	firstReservedWord := g.declareToken("BreakKeyword")
	// Reserved words
	firstKeyword := firstReservedWord
	g.declareToken("CaseKeyword")
	g.declareToken("CatchKeyword")
	g.declareToken("ClassKeyword")
	g.declareToken("ConstKeyword")
	g.declareToken("ContinueKeyword")
	g.declareToken("DebuggerKeyword")
	g.declareToken("DefaultKeyword")
	g.declareToken("DeleteKeyword")
	g.declareToken("DoKeyword")
	g.declareToken("ElseKeyword")
	g.declareToken("EnumKeyword")
	g.declareToken("ExportKeyword")
	g.declareToken("ExtendsKeyword")
	g.declareToken("FalseKeyword")
	g.declareToken("FinallyKeyword")
	g.declareToken("ForKeyword")
	g.declareToken("FunctionKeyword")
	g.declareToken("IfKeyword")
	g.declareToken("ImportKeyword")
	g.declareToken("InKeyword")
	g.declareToken("InstanceOfKeyword")
	g.declareToken("NewKeyword")
	g.declareToken("NullKeyword")
	g.declareToken("ReturnKeyword")
	g.declareToken("SuperKeyword")
	g.declareToken("SwitchKeyword")
	g.declareToken("ThisKeyword")
	g.declareToken("ThrowKeyword")
	g.declareToken("TrueKeyword")
	g.declareToken("TryKeyword")
	g.declareToken("TypeOfKeyword")
	g.declareToken("VarKeyword")
	g.declareToken("VoidKeyword")
	g.declareToken("WhileKeyword")
	lastReservedWord := g.declareToken("WithKeyword")

	// Strict mode reserved words
	firstFutureReservedWord := g.declareToken("ImplementsKeyword")
	g.declareToken("InterfaceKeyword")
	g.declareToken("LetKeyword")
	g.declareToken("PackageKeyword")
	g.declareToken("PrivateKeyword")
	g.declareToken("ProtectedKeyword")
	g.declareToken("PublicKeyword")
	g.declareToken("StaticKeyword")
	lastFutureReservedWord := g.declareToken("YieldKeyword")

	// Contextual keywords
	firstContextualKeyword := g.declareToken("AbstractKeyword")
	g.declareToken("AccessorKeyword")
	g.declareToken("AsKeyword")
	g.declareToken("AssertsKeyword")
	g.declareToken("AssertKeyword")
	g.declareToken("AnyKeyword")
	g.declareToken("AsyncKeyword")
	g.declareToken("AwaitKeyword")
	g.declareToken("BooleanKeyword")
	g.declareToken("ConstructorKeyword")
	g.declareToken("DeclareKeyword")
	g.declareToken("GetKeyword")
	g.declareToken("ImmediateKeyword")
	g.declareToken("InferKeyword")
	g.declareToken("IntrinsicKeyword")
	g.declareToken("IsKeyword")
	g.declareToken("KeyOfKeyword")
	g.declareToken("ModuleKeyword")
	g.declareToken("NamespaceKeyword")
	g.declareToken("NeverKeyword")
	g.declareToken("OutKeyword")
	g.declareToken("ReadonlyKeyword")
	g.declareToken("RequireKeyword")
	g.declareToken("NumberKeyword")
	g.declareToken("ObjectKeyword")
	g.declareToken("SatisfiesKeyword")
	g.declareToken("SetKeyword")
	g.declareToken("StringKeyword")
	g.declareToken("SymbolKeyword")
	g.declareToken("TypeKeyword")
	g.declareToken("UndefinedKeyword")
	g.declareToken("UniqueKeyword")
	g.declareToken("UnknownKeyword")
	g.declareToken("UsingKeyword")
	g.declareToken("FromKeyword")
	g.declareToken("GlobalKeyword")
	g.declareToken("BigIntKeyword")
	g.declareToken("OverrideKeyword")
	lastKeyword := g.declareToken("OfKeyword")
	lastToken := lastKeyword
	lastContextualKeyword := lastKeyword

	// Parse tree nodes
	// Names
	firstNode := g.declareNode("QualifiedName", &nodeOpts{})
	g.declareNode("ComputedPropertyName", &nodeOpts{})

	// Lists
	g.declareNode("ModifierList", &nodeOpts{})
	g.declareNode("TypeParameterList", &nodeOpts{})
	g.declareNode("TypeArgumentList", &nodeOpts{})

	// Signature elements
	g.declareNode("TypeParameter", &nodeOpts{})
	g.declareNode("Parameter", &nodeOpts{})
	g.declareNode("Decorator", &nodeOpts{})

	// TypeMember
	g.declareNode("PropertySignature", &nodeOpts{})
	g.declareNode("PropertyDeclaration", &nodeOpts{})
	g.declareNode("MethodSignature", &nodeOpts{})
	g.declareNode("MethodDeclaration", &nodeOpts{})
	g.declareNode("ClassStaticBlockDeclaration", &nodeOpts{})
	g.declareNode("Constructor", &nodeOpts{})
	g.declareNode("GetAccessor", &nodeOpts{})
	g.declareNode("SetAccessor", &nodeOpts{})
	g.declareNode("CallSignature", &nodeOpts{})
	g.declareNode("ConstructSignature", &nodeOpts{})
	g.declareNode("IndexSignature", &nodeOpts{})

	// Type
	firstTypeNode := g.declareNode("TypePredicate", &nodeOpts{})
	g.declareNode("TypeReference", &nodeOpts{})
	g.declareNode("FunctionType", &nodeOpts{})
	g.declareNode("ConstructorType", &nodeOpts{})
	g.declareNode("TypeQuery", &nodeOpts{})
	g.declareNode("TypeLiteral", &nodeOpts{})
	g.declareNode("ArrayType", &nodeOpts{})
	g.declareNode("TupleType", &nodeOpts{})
	g.declareNode("OptionalType", &nodeOpts{})
	g.declareNode("RestType", &nodeOpts{})
	g.declareNode("UnionType", &nodeOpts{})
	g.declareNode("IntersectionType", &nodeOpts{})
	g.declareNode("ConditionalType", &nodeOpts{})
	g.declareNode("InferType", &nodeOpts{})
	g.declareNode("ParenthesizedType", &nodeOpts{})
	g.declareNode("ThisType", &nodeOpts{})
	g.declareNode("TypeOperator", &nodeOpts{})
	g.declareNode("IndexedAccessType", &nodeOpts{})
	g.declareNode("MappedType", &nodeOpts{})
	g.declareNode("LiteralType", &nodeOpts{})
	g.declareNode("NamedTupleMember", &nodeOpts{})
	g.declareNode("TemplateLiteralType", &nodeOpts{})
	g.declareNode("TemplateLiteralTypeSpan", &nodeOpts{})
	lastTypeNode := g.declareNode("ImportType", &nodeOpts{})

	// Binding patterns
	g.declareNode("ObjectBindingPattern", &nodeOpts{})
	g.declareNode("ArrayBindingPattern", &nodeOpts{})
	g.declareNode("BindingElement", &nodeOpts{})

	// Expression
	g.declareNode("ArrayLiteralExpression", &nodeOpts{})
	g.declareNode("ObjectLiteralExpression", &nodeOpts{})
	propertyAccessExpression := g.declareNode("PropertyAccessExpression", &nodeOpts{})
	elementAccessExpression := g.declareNode("ElementAccessExpression", &nodeOpts{})
	g.declareNode("CallExpression", &nodeOpts{})
	g.declareNode("NewExpression", &nodeOpts{})
	g.declareNode("TaggedTemplateExpression", &nodeOpts{})
	g.declareNode("TypeAssertionExpression", &nodeOpts{})
	g.declareNode("ParenthesizedExpression", &nodeOpts{})
	g.declareNode("FunctionExpression", &nodeOpts{})
	g.declareNode("ArrowFunction", &nodeOpts{})
	g.declareNode("DeleteExpression", &nodeOpts{})
	g.declareNode("TypeOfExpression", &nodeOpts{})
	g.declareNode("VoidExpression", &nodeOpts{})
	g.declareNode("AwaitExpression", &nodeOpts{})
	g.declareNode("PrefixUnaryExpression", &nodeOpts{})
	g.declareNode("PostfixUnaryExpression", &nodeOpts{})
	g.declareNode("BinaryExpression", &nodeOpts{})
	g.declareNode("ConditionalExpression", &nodeOpts{})
	g.declareNode("TemplateExpression", &nodeOpts{})
	g.declareNode("YieldExpression", &nodeOpts{})
	g.declareNode("SpreadElement", &nodeOpts{})
	g.declareNode("ClassExpression", &nodeOpts{})
	g.declareNode("OmittedExpression", &nodeOpts{})
	g.declareNode("ExpressionWithTypeArguments", &nodeOpts{})
	g.declareNode("AsExpression", &nodeOpts{})
	g.declareNode("NonNullExpression", &nodeOpts{})
	g.declareNode("MetaProperty", &nodeOpts{})
	g.declareNode("SyntheticExpression", &nodeOpts{})
	g.declareNode("SatisfiesExpression", &nodeOpts{})

	// Misc
	g.declareNode("TemplateSpan", &nodeOpts{})
	g.declareNode("SemicolonClassElement", &nodeOpts{})

	// Element
	g.declareNode("Block", &nodeOpts{})
	g.declareNode("EmptyStatement", &nodeOpts{})
	g.declareNode("VariableStatement", &nodeOpts{})
	g.declareNode("ExpressionStatement", &nodeOpts{})
	g.declareNode("IfStatement", &nodeOpts{})
	g.declareNode("DoStatement", &nodeOpts{})
	g.declareNode("WhileStatement", &nodeOpts{})
	g.declareNode("ForStatement", &nodeOpts{})
	g.declareNode("ForInStatement", &nodeOpts{})
	g.declareNode("ForOfStatement", &nodeOpts{})
	g.declareNode("ContinueStatement", &nodeOpts{})
	g.declareNode("BreakStatement", &nodeOpts{})
	g.declareNode("ReturnStatement", &nodeOpts{})
	g.declareNode("WithStatement", &nodeOpts{})
	g.declareNode("SwitchStatement", &nodeOpts{})
	g.declareNode("LabeledStatement", &nodeOpts{})
	g.declareNode("ThrowStatement", &nodeOpts{})
	g.declareNode("TryStatement", &nodeOpts{})
	g.declareNode("DebuggerStatement", &nodeOpts{})
	g.declareNode("VariableDeclaration", &nodeOpts{})
	g.declareNode("VariableDeclarationList", &nodeOpts{})
	g.declareNode("FunctionDeclaration", &nodeOpts{})
	g.declareNode("ClassDeclaration", &nodeOpts{})
	g.declareNode("InterfaceDeclaration", &nodeOpts{})
	g.declareNode("TypeAliasDeclaration", &nodeOpts{})
	g.declareNode("EnumDeclaration", &nodeOpts{})
	g.declareNode("ModuleDeclaration", &nodeOpts{})
	g.declareNode("ModuleBlock", &nodeOpts{})
	g.declareNode("CaseBlock", &nodeOpts{})
	g.declareNode("NamespaceExportDeclaration", &nodeOpts{})
	g.declareNode("ImportEqualsDeclaration", &nodeOpts{})
	g.declareNode("ImportDeclaration", &nodeOpts{})
	g.declareNode("ImportClause", &nodeOpts{})
	g.declareNode("NamespaceImport", &nodeOpts{})
	g.declareNode("NamedImports", &nodeOpts{})
	g.declareNode("ImportSpecifier", &nodeOpts{})
	g.declareNode("ExportAssignment", &nodeOpts{})
	g.declareNode("ExportDeclaration", &nodeOpts{})
	g.declareNode("NamedExports", &nodeOpts{})
	g.declareNode("NamespaceExport", &nodeOpts{})
	g.declareNode("ExportSpecifier", &nodeOpts{})
	g.declareNode("MissingDeclaration", &nodeOpts{})

	// Module references
	g.declareNode("ExternalModuleReference", &nodeOpts{})

	// JSX
	g.declareNode("JsxElement", &nodeOpts{})
	g.declareNode("JsxSelfClosingElement", &nodeOpts{})
	g.declareNode("JsxOpeningElement", &nodeOpts{})
	g.declareNode("JsxClosingElement", &nodeOpts{})
	g.declareNode("JsxFragment", &nodeOpts{})
	g.declareNode("JsxOpeningFragment", &nodeOpts{})
	g.declareNode("JsxClosingFragment", &nodeOpts{})
	g.declareNode("JsxAttribute", &nodeOpts{})
	g.declareNode("JsxAttributes", &nodeOpts{})
	g.declareNode("JsxSpreadAttribute", &nodeOpts{})
	g.declareNode("JsxExpression", &nodeOpts{})
	g.declareNode("JsxNamespacedName", &nodeOpts{})

	// Clauses
	g.declareNode("CaseClause", &nodeOpts{})
	g.declareNode("DefaultClause", &nodeOpts{})
	g.declareNode("HeritageClause", &nodeOpts{})
	g.declareNode("CatchClause", &nodeOpts{})

	// Import attributes
	g.declareNode("ImportAttributes", &nodeOpts{})
	g.declareNode("ImportAttribute", &nodeOpts{})

	// Property assignments
	g.declareNode("PropertyAssignment", &nodeOpts{})
	g.declareNode("ShorthandPropertyAssignment", &nodeOpts{})
	g.declareNode("SpreadAssignment", &nodeOpts{})

	// Enum
	g.declareNode("EnumMember", &nodeOpts{})

	// Top-level nodes
	g.declareNode("SourceFile", &nodeOpts{})
	g.declareNode("Bundle", &nodeOpts{})

	// JSDoc nodes
	firstJSDocNode := g.declareNode("JSDocTypeExpression", &nodeOpts{})
	g.declareNode("JSDocNameReference", &nodeOpts{})
	g.declareNode("JSDocMemberName", &nodeOpts{})  // C#p
	g.declareNode("JSDocAllType", &nodeOpts{})     // The * type
	g.declareNode("JSDocUnknownType", &nodeOpts{}) // The ? type
	g.declareNode("JSDocNullableType", &nodeOpts{})
	g.declareNode("JSDocNonNullableType", &nodeOpts{})
	g.declareNode("JSDocOptionalType", &nodeOpts{})
	g.declareNode("JSDocFunctionType", &nodeOpts{})
	g.declareNode("JSDocVariadicType", &nodeOpts{})
	g.declareNode("JSDocNamepathType", &nodeOpts{}) // https://jsdoc.app/about-namepaths.html
	g.declareNode("JSDoc", &nodeOpts{})
	g.declareNode("JSDocText", &nodeOpts{})
	g.declareNode("JSDocTypeLiteral", &nodeOpts{})
	g.declareNode("JSDocSignature", &nodeOpts{})
	g.declareNode("JSDocLink", &nodeOpts{})
	g.declareNode("JSDocLinkCode", &nodeOpts{})
	g.declareNode("JSDocLinkPlain", &nodeOpts{})
	firstJSDocTagNode := g.declareNode("JSDocTag", &nodeOpts{})
	g.declareNode("JSDocAugmentsTag", &nodeOpts{})
	g.declareNode("JSDocImplementsTag", &nodeOpts{})
	g.declareNode("JSDocAuthorTag", &nodeOpts{})
	g.declareNode("JSDocDeprecatedTag", &nodeOpts{})
	g.declareNode("JSDocImmediateTag", &nodeOpts{})
	g.declareNode("JSDocClassTag", &nodeOpts{})
	g.declareNode("JSDocPublicTag", &nodeOpts{})
	g.declareNode("JSDocPrivateTag", &nodeOpts{})
	g.declareNode("JSDocProtectedTag", &nodeOpts{})
	g.declareNode("JSDocReadonlyTag", &nodeOpts{})
	g.declareNode("JSDocOverrideTag", &nodeOpts{})
	g.declareNode("JSDocCallbackTag", &nodeOpts{})
	g.declareNode("JSDocOverloadTag", &nodeOpts{})
	g.declareNode("JSDocEnumTag", &nodeOpts{})
	g.declareNode("JSDocParameterTag", &nodeOpts{})
	g.declareNode("JSDocReturnTag", &nodeOpts{})
	g.declareNode("JSDocThisTag", &nodeOpts{})
	g.declareNode("JSDocTypeTag", &nodeOpts{})
	g.declareNode("JSDocTemplateTag", &nodeOpts{})
	g.declareNode("JSDocTypedefTag", &nodeOpts{})
	g.declareNode("JSDocSeeTag", &nodeOpts{})
	g.declareNode("JSDocPropertyTag", &nodeOpts{})
	g.declareNode("JSDocThrowsTag", &nodeOpts{})
	g.declareNode("JSDocSatisfiesTag", &nodeOpts{})
	lastJSDocNode := g.declareNode("JSDocImportTag", &nodeOpts{})
	lastJSDocTagNode := lastJSDocNode

	// Synthesized list
	g.declareNode("SyntaxList", &nodeOpts{})

	// Transformation nodes
	g.declareNode("NotEmittedStatement", &nodeOpts{})
	g.declareNode("PartiallyEmittedExpression", &nodeOpts{})
	g.declareNode("CommaListExpression", &nodeOpts{})
	g.declareNode("SyntheticReferenceExpression", &nodeOpts{})

	// Markers

	g.newMarker("Assignment", firstAssignment, lastAssignment)
	g.newMarker("CompoundAssignment", firstCompoundAssignment, lastCompoundAssignment)
	g.newMarker("ReservedWord", firstReservedWord, lastReservedWord)
	g.newMarker("Keyword", firstKeyword, lastKeyword)
	g.newMarker("FutureReservedWord", firstFutureReservedWord, lastFutureReservedWord)
	g.newMarker("TypeNode", firstTypeNode, lastTypeNode)
	g.newMarker("Punctuation", firstPunctuation, lastPunctuation)
	g.newMarker("Token", firstToken, lastToken)
	g.newMarker("LiteralToken", firstLiteralToken, lastLiteralToken)
	g.newMarker("TemplateToken", firstTemplateToken, lastTemplateToken)
	g.newMarker("BinaryOperator", firstBinaryOperator, lastBinaryOperator)
	g.newMarker("Node", firstNode, g.syntaxKinds[len(g.syntaxKinds)-1])
	g.newMarker("JSDocNode", firstJSDocNode, lastJSDocNode)
	g.newMarker("JSDocTagNode", firstJSDocTagNode, lastJSDocTagNode)
	g.newMarker("ContextualKeyword", firstContextualKeyword, lastContextualKeyword)

	g.newNodeUnion("AccessExpression", propertyAccessExpression, elementAccessExpression)
}

type marker struct {
	name  string
	start *syntaxKind
	end   *syntaxKind
}

func (g *generator) newMarker(name string, start, end *syntaxKind) {
	g.markers = append(g.markers, &marker{name: name, start: start, end: end})
}

type nodeOpts struct {
	poolAllocate bool

	embedded any

	children []field
}

type field struct {
	name string
	typ  genType
}

func (g *generator) declareToken(name string) *syntaxKind {
	return g.declareNode(name, nil)
}

func (g *generator) declareNode(name string, opts *nodeOpts) *syntaxKind {
	n := &syntaxKind{
		name: name,
		opts: opts,
	}
	g.syntaxKinds = append(g.syntaxKinds, n)
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
	anyType    = &goType{name: "any"}
	boolType   = &goType{name: "bool"}
)

type nodeUnion struct {
	name  string
	types []*syntaxKind
}

var _ genType = (*nodeUnion)(nil)

func (u *nodeUnion) Name() string { return u.name }

func (g *generator) newNodeUnion(name string, types ...*syntaxKind) *nodeUnion {
	u := &nodeUnion{name: name, types: types}
	g.unions = append(g.unions, u)
	return u
}

type syntaxKind struct {
	name   string
	opts   *nodeOpts // nil if a token
	fields []field
}

var _ genType = (*syntaxKind)(nil)

func (n *syntaxKind) Name() string { return n.name }

// Code generation

func (g *generator) generate() {
	g.println("package ast")
	g.println()
	g.println("import \"fmt\"")
	g.println()
	g.println("type SyntaxKind int16")
	g.println()
	g.println("const (")

	for i, n := range g.syntaxKinds {
		if i == 0 {
			g.printf("\tSyntaxKind%s SyntaxKind = iota\n", n.name)
		} else {
			g.printf("\tSyntaxKind%s\n", n.name)
		}
	}

	g.println()
	g.println("\tSyntaxKindCount")
	g.println()

	for _, marker := range g.markers {
		g.printf("\t%sStart = SyntaxKind%s\n", marker.name, marker.start.name)
		g.printf("\t%sEnd = SyntaxKind%s\n", marker.name, marker.end.name)
	}

	g.println(")")
	g.println()

	g.println("var syntaxKindNames = [SyntaxKindCount+1]string{")
	for _, n := range g.syntaxKinds {
		g.printf("\tSyntaxKind%s: %q,\n", n.name, n.name)
	}
	g.println("\tSyntaxKindCount: \"SyntaxKindCount\",")
	g.println("}")
	g.println()

	g.println("func (k SyntaxKind) String() string {")
	g.println("\tif k < 0 || k >= SyntaxKindCount {")
	g.printf("%s", "\t\treturn fmt.Sprintf(\"SyntaxKind(%d)\", k)")
	g.println("\t}")
	g.println("\treturn syntaxKindNames[k]")
	g.println("}")
	g.println()

	g.println("type NodeFlags uint32")
	g.println("type NodeID uint32")
	g.println()

	g.println("type TextRange struct {")
	g.println("\tpos int32")
	g.println("\tend int32")
	g.println("}")
	g.println()
	g.println("func NewTextRange(pos, end int32) TextRange { return TextRange{pos: pos, end: end} }")
	g.println()
	g.println("func (r TextRange) Pos() int32 { return r.pos }")
	g.println("func (r TextRange) End() int32 { return r.end }")
	g.println("func (r TextRange) Len() int32 { return r.end - r.pos }")
	g.println("func (r TextRange) ContainsInclusive(pos int32) bool { return r.pos <= pos && pos <= r.end }")
	g.println()

	// TODO
	g.println("type NodeData interface{")
	g.println("\tAsNode() *Node")
	g.println("}")
	g.println()

	g.println("type Node struct {")
	g.println("\tkind SyntaxKind")
	g.println("\tflags  NodeFlags")
	g.println("\tloc TextRange")
	g.println("\tid NodeID")
	g.println("\tparent *Node")
	g.println("\tdata NodeData")
	g.println("}")
	g.println()

	g.println("func (n *Node) Pos() int32 { return n.loc.Pos() }")

	// TODO
	g.println("type NodeDefault struct {\nNode\n}")
	g.println()
	g.println("func (n *NodeDefault) AsNode() *Node { return &n.Node }")
	g.println()

	// TODO
	g.println("type NodeBase struct {\nNodeDefault\n} ")
	g.println()

	g.println("type Factory struct { // TODO")
	for _, n := range g.syntaxKinds {
		token := n == g.firstToken
		if token || n.opts != nil && n.opts.poolAllocate {
			name := n.name
			if token {
				name = "Token"
			}
			g.printf("\tpool%s pool[%s]\n", name, name)
		}
	}
	g.println("}")
	g.println()

	for _, n := range g.syntaxKinds {
		g.generateNode(n)
	}

	for _, u := range g.unions {
		g.generateNodeUnion(u)
	}
}

func (g *generator) generateNodeUnion(u *nodeUnion) {
	g.printf("type %s = Node // ", u.name)
	for i, t := range u.types {
		if i > 0 {
			g.print(" | ")
		}
		g.printf("%s", t.name)
	}
	g.println()
	g.println()

	g.printf("var is%sTable = [SyntaxKindCount]bool{\n", u.name)
	for _, t := range u.types {
		g.printf("\tSyntaxKind%s: true,\n", t.name)
	}
	g.println("}")
	g.println()

	g.printf("func is%sKind(kind SyntaxKind) bool { return is%sTable[kind] }\n", u.name, u.name)
	g.printf("func is%s(n *Node) bool { return is%sTable[n.kind] }\n", u.name, u.name)
	g.printf("func assertIs%s(n *Node) {\n", u.name)
	g.printf("if !is%s(n) {\n", u.name)
	g.printf("panic(\"expected %s, got \" + n.kind.String())\n", u.name)
	g.printf("}\n")
	g.printf("}\n")
	g.println()
}

func (g *generator) generateNode(n *syntaxKind) {
	token := n == g.firstToken
	name := n.name
	if token {
		name = "Token"
	} else if n.opts == nil {
		return
	}

	g.println()
	g.printf("type %s struct {\n", name)
	g.println("\tNodeBase")
	// TODO: children
	g.println("}")
	g.println()

	g.printf("func (n *Node) As%s() *%s { return n.data.(*%s) }\n", name, name, name)
	if !token {
		g.printf("func (n *Node) Is%s() bool { return n.kind == SyntaxKind%s }\n", name, name)
	}
	g.println()

	printNewParams := func() {
		// TODO
		if token {
			g.print("kind SyntaxKind")
		}
	}
	printNewArgs := func() {
		// TODO
		if token {
			g.print("kind")
		}
	}

	// TODO: accept and set children
	g.printf("func (n *%s) set(", name)
	printNewParams()
	g.println(") {")
	// TODO: generate optimal assignment
	g.printf("\t*n = %s{}\n", name)
	if token {
		g.println("\tn.kind = kind")
	} else {
		g.println("\tn.kind = SyntaxKind" + name)
	}
	g.println("\tn.data = n")
	g.println("}")
	g.println()

	if !token {
		g.printf("func (n *%s) Kind() SyntaxKind { return SyntaxKind%s }\n", name, name)
		g.println()
	}

	g.printf("func New%s(", name)
	printNewParams()
	g.println(") *Node {")
	g.printf("\tn := &%s{}\n", name)
	g.printf("\tn.set(")
	printNewArgs()
	g.println(")")
	g.println("\treturn n.AsNode()")
	g.println("}")
	g.println()

	// TODO: params
	// TODO: return *Node, call a helper func that converts to *Node and sets kind
	g.printf("func (f *Factory) New%s(", name)
	printNewParams()
	g.println(") *Node {")
	if token || n.opts != nil && n.opts.poolAllocate {
		g.printf("\tn := f.pool%s.allocate()\n", name)
		g.printf("\tn.set(")
		printNewArgs()
		g.println(")")
		g.println("\treturn n.AsNode()")
	} else {
		g.printf("\treturn New%s(", name)
		printNewArgs()
		g.println(")")
	}
	g.println("}")

	g.println()
}
