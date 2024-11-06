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
	declare(&buf)

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

func declare(w io.Writer) {
	var g generator

	firstToken := g.declareKindOnly("Unknown")
	g.declareKindOnly("EndOfFile")
	g.declareKindOnly("ConflictMarkerTrivia")
	g.declareKindOnly("NonTextFileMarkerTrivia")

	firstLiteralToken := g.declare("NumericLiteral", &declOptions{})
	g.declare("BigIntLiteral", &declOptions{})
	g.declare("StringLiteral", &declOptions{})
	g.declare("JsxText", &declOptions{})
	g.declareKindOnly("JsxTextAllWhiteSpaces")
	g.declare("RegularExpressionLiteral", &declOptions{})
	lastLiteralToken := g.declare("NoSubstitutionTemplateLiteral", &declOptions{})

	// Pseudo-literals
	firstTemplateToken := lastLiteralToken
	g.declare("TemplateHead", &declOptions{})
	g.declare("TemplateMiddle", &declOptions{})
	lastTemplateToken := g.declare("TemplateTail", &declOptions{})

	firstPunctuation := g.declareKindOnly("OpenBraceToken")
	g.declareKindOnly("CloseBraceToken")
	g.declareKindOnly("OpenParenToken")
	g.declareKindOnly("CloseParenToken")
	g.declareKindOnly("OpenBracketToken")
	g.declareKindOnly("CloseBracketToken")
	g.declareKindOnly("DotToken")
	g.declareKindOnly("DotDotDotToken")
	g.declareKindOnly("SemicolonToken")
	g.declareKindOnly("CommaToken")
	g.declareKindOnly("QuestionDotToken")
	firstBinaryOperator := g.declareKindOnly("LessThanToken")
	g.declareKindOnly("LessThanSlashToken")
	g.declareKindOnly("GreaterThanToken")
	g.declareKindOnly("LessThanEqualsToken")
	g.declareKindOnly("GreaterThanEqualsToken")
	g.declareKindOnly("EqualsEqualsToken")
	g.declareKindOnly("ExclamationEqualsToken")
	g.declareKindOnly("EqualsEqualsEqualsToken")
	g.declareKindOnly("ExclamationEqualsEqualsToken")
	g.declareKindOnly("EqualsGreaterThanToken")
	g.declareKindOnly("PlusToken")
	g.declareKindOnly("MinusToken")
	g.declareKindOnly("AsteriskToken")
	g.declareKindOnly("AsteriskAsteriskToken")
	g.declareKindOnly("SlashToken")
	g.declareKindOnly("PercentToken")
	g.declareKindOnly("PlusPlusToken")
	g.declareKindOnly("MinusMinusToken")
	g.declareKindOnly("LessThanLessThanToken")
	g.declareKindOnly("GreaterThanGreaterThanToken")
	g.declareKindOnly("GreaterThanGreaterThanGreaterThanToken")
	g.declareKindOnly("AmpersandToken")
	g.declareKindOnly("BarToken")
	g.declareKindOnly("CaretToken")
	g.declareKindOnly("ExclamationToken")
	g.declareKindOnly("TildeToken")
	g.declareKindOnly("AmpersandAmpersandToken")
	g.declareKindOnly("BarBarToken")
	g.declareKindOnly("QuestionToken")
	g.declareKindOnly("ColonToken")
	g.declareKindOnly("AtToken")
	g.declareKindOnly("QuestionQuestionToken")
	g.declareKindOnly("BacktickToken")
	g.declareKindOnly("HashToken")
	firstAssignment := g.declareKindOnly("EqualsToken")
	firstCompoundAssignment := g.declareKindOnly("PlusEqualsToken")
	g.declareKindOnly("MinusEqualsToken")
	g.declareKindOnly("AsteriskEqualsToken")
	g.declareKindOnly("AsteriskAsteriskEqualsToken")
	g.declareKindOnly("SlashEqualsToken")
	g.declareKindOnly("PercentEqualsToken")
	g.declareKindOnly("LessThanLessThanEqualsToken")
	g.declareKindOnly("GreaterThanGreaterThanEqualsToken")
	g.declareKindOnly("GreaterThanGreaterThanGreaterThanEqualsToken")
	g.declareKindOnly("AmpersandEqualsToken")
	g.declareKindOnly("BarEqualsToken")
	g.declareKindOnly("BarBarEqualsToken")
	g.declareKindOnly("AmpersandAmpersandEqualsToken")
	g.declareKindOnly("QuestionQuestionEqualsToken")
	lastPunctuation := g.declareKindOnly("CaretEqualsToken")
	lastAssignment := lastPunctuation
	lastCompoundAssignment := lastPunctuation
	lastBinaryOperator := lastPunctuation

	g.declare("Identifier", &declOptions{
		poolAllocate: true,
	})
	g.declare("PrivateIdentifier", &declOptions{})
	g.declareKindOnly("JSDocCommentTextToken") // TODO: why is this here??

	firstReservedWord := g.declareKindOnly("BreakKeyword")
	// Reserved words
	firstKeyword := firstReservedWord
	g.declareKindOnly("CaseKeyword")
	g.declareKindOnly("CatchKeyword")
	g.declareKindOnly("ClassKeyword")
	g.declareKindOnly("ConstKeyword")
	g.declareKindOnly("ContinueKeyword")
	g.declareKindOnly("DebuggerKeyword")
	g.declareKindOnly("DefaultKeyword")
	g.declareKindOnly("DeleteKeyword")
	g.declareKindOnly("DoKeyword")
	g.declareKindOnly("ElseKeyword")
	g.declareKindOnly("EnumKeyword")
	g.declareKindOnly("ExportKeyword")
	g.declareKindOnly("ExtendsKeyword")
	g.declareKindOnly("FalseKeyword")
	g.declareKindOnly("FinallyKeyword")
	g.declareKindOnly("ForKeyword")
	g.declareKindOnly("FunctionKeyword")
	g.declareKindOnly("IfKeyword")
	g.declareKindOnly("ImportKeyword")
	g.declareKindOnly("InKeyword")
	g.declareKindOnly("InstanceOfKeyword")
	g.declareKindOnly("NewKeyword")
	g.declareKindOnly("NullKeyword")
	g.declareKindOnly("ReturnKeyword")
	g.declareKindOnly("SuperKeyword")
	g.declareKindOnly("SwitchKeyword")
	g.declareKindOnly("ThisKeyword")
	g.declareKindOnly("ThrowKeyword")
	g.declareKindOnly("TrueKeyword")
	g.declareKindOnly("TryKeyword")
	g.declareKindOnly("TypeOfKeyword")
	g.declareKindOnly("VarKeyword")
	g.declareKindOnly("VoidKeyword")
	g.declareKindOnly("WhileKeyword")
	lastReservedWord := g.declareKindOnly("WithKeyword")

	// Strict mode reserved words
	firstFutureReservedWord := g.declareKindOnly("ImplementsKeyword")
	g.declareKindOnly("InterfaceKeyword")
	g.declareKindOnly("LetKeyword")
	g.declareKindOnly("PackageKeyword")
	g.declareKindOnly("PrivateKeyword")
	g.declareKindOnly("ProtectedKeyword")
	g.declareKindOnly("PublicKeyword")
	g.declareKindOnly("StaticKeyword")
	lastFutureReservedWord := g.declareKindOnly("YieldKeyword")

	// Contextual keywords
	firstContextualKeyword := g.declareKindOnly("AbstractKeyword")
	g.declareKindOnly("AccessorKeyword")
	g.declareKindOnly("AsKeyword")
	g.declareKindOnly("AssertsKeyword")
	g.declareKindOnly("AssertKeyword")
	g.declareKindOnly("AnyKeyword")
	g.declareKindOnly("AsyncKeyword")
	g.declareKindOnly("AwaitKeyword")
	g.declareKindOnly("BooleanKeyword")
	g.declareKindOnly("ConstructorKeyword")
	g.declareKindOnly("DeclareKeyword")
	g.declareKindOnly("GetKeyword")
	g.declareKindOnly("ImmediateKeyword")
	g.declareKindOnly("InferKeyword")
	g.declareKindOnly("IntrinsicKeyword")
	g.declareKindOnly("IsKeyword")
	g.declareKindOnly("KeyOfKeyword")
	g.declareKindOnly("ModuleKeyword")
	g.declareKindOnly("NamespaceKeyword")
	g.declareKindOnly("NeverKeyword")
	g.declareKindOnly("OutKeyword")
	g.declareKindOnly("ReadonlyKeyword")
	g.declareKindOnly("RequireKeyword")
	g.declareKindOnly("NumberKeyword")
	g.declareKindOnly("ObjectKeyword")
	g.declareKindOnly("SatisfiesKeyword")
	g.declareKindOnly("SetKeyword")
	g.declareKindOnly("StringKeyword")
	g.declareKindOnly("SymbolKeyword")
	g.declareKindOnly("TypeKeyword")
	g.declareKindOnly("UndefinedKeyword")
	g.declareKindOnly("UniqueKeyword")
	g.declareKindOnly("UnknownKeyword")
	g.declareKindOnly("UsingKeyword")
	g.declareKindOnly("FromKeyword")
	g.declareKindOnly("GlobalKeyword")
	g.declareKindOnly("BigIntKeyword")
	g.declareKindOnly("OverrideKeyword")
	lastKeyword := g.declareKindOnly("OfKeyword")
	lastToken := lastKeyword
	lastContextualKeyword := lastKeyword

	// Parse tree nodes
	// Names
	firstNode := g.declare("QualifiedName", &declOptions{})
	g.declare("ComputedPropertyName", &declOptions{})

	// Lists
	g.declare("ModifierList", &declOptions{})
	g.declare("TypeParameterList", &declOptions{})
	g.declare("TypeArgumentList", &declOptions{})

	// Signature elements
	g.declare("TypeParameter", &declOptions{})
	g.declare("Parameter", &declOptions{})
	g.declare("Decorator", &declOptions{})

	// TypeMember
	g.declare("PropertySignature", &declOptions{})
	g.declare("PropertyDeclaration", &declOptions{})
	g.declare("MethodSignature", &declOptions{})
	g.declare("MethodDeclaration", &declOptions{})
	g.declare("ClassStaticBlockDeclaration", &declOptions{})
	g.declare("Constructor", &declOptions{})
	g.declare("GetAccessor", &declOptions{})
	g.declare("SetAccessor", &declOptions{})
	g.declare("CallSignature", &declOptions{})
	g.declare("ConstructSignature", &declOptions{})
	g.declare("IndexSignature", &declOptions{})

	// Type
	firstTypeNode := g.declare("TypePredicate", &declOptions{})
	g.declare("TypeReference", &declOptions{})
	g.declare("FunctionType", &declOptions{})
	g.declare("ConstructorType", &declOptions{})
	g.declare("TypeQuery", &declOptions{})
	g.declare("TypeLiteral", &declOptions{})
	g.declare("ArrayType", &declOptions{})
	g.declare("TupleType", &declOptions{})
	g.declare("OptionalType", &declOptions{})
	g.declare("RestType", &declOptions{})
	g.declare("UnionType", &declOptions{})
	g.declare("IntersectionType", &declOptions{})
	g.declare("ConditionalType", &declOptions{})
	g.declare("InferType", &declOptions{})
	g.declare("ParenthesizedType", &declOptions{})
	g.declare("ThisType", &declOptions{})
	g.declare("TypeOperator", &declOptions{})
	g.declare("IndexedAccessType", &declOptions{})
	g.declare("MappedType", &declOptions{})
	g.declare("LiteralType", &declOptions{})
	g.declare("NamedTupleMember", &declOptions{})
	g.declare("TemplateLiteralType", &declOptions{})
	g.declare("TemplateLiteralTypeSpan", &declOptions{})
	lastTypeNode := g.declare("ImportType", &declOptions{})

	// Binding patterns
	g.declare("ObjectBindingPattern", &declOptions{})
	g.declare("ArrayBindingPattern", &declOptions{})
	g.declare("BindingElement", &declOptions{})

	// Expression
	g.declare("ArrayLiteralExpression", &declOptions{})
	g.declare("ObjectLiteralExpression", &declOptions{})
	propertyAccessExpression := g.declare("PropertyAccessExpression", &declOptions{})
	elementAccessExpression := g.declare("ElementAccessExpression", &declOptions{})
	g.declare("CallExpression", &declOptions{})
	g.declare("NewExpression", &declOptions{})
	g.declare("TaggedTemplateExpression", &declOptions{})
	g.declare("TypeAssertionExpression", &declOptions{})
	g.declare("ParenthesizedExpression", &declOptions{})
	g.declare("FunctionExpression", &declOptions{})
	g.declare("ArrowFunction", &declOptions{})
	g.declare("DeleteExpression", &declOptions{})
	g.declare("TypeOfExpression", &declOptions{})
	g.declare("VoidExpression", &declOptions{})
	g.declare("AwaitExpression", &declOptions{})
	g.declare("PrefixUnaryExpression", &declOptions{})
	g.declare("PostfixUnaryExpression", &declOptions{})
	g.declare("BinaryExpression", &declOptions{})
	g.declare("ConditionalExpression", &declOptions{})
	g.declare("TemplateExpression", &declOptions{})
	g.declare("YieldExpression", &declOptions{})
	g.declare("SpreadElement", &declOptions{})
	g.declare("ClassExpression", &declOptions{})
	g.declare("OmittedExpression", &declOptions{})
	g.declare("ExpressionWithTypeArguments", &declOptions{})
	g.declare("AsExpression", &declOptions{})
	g.declare("NonNullExpression", &declOptions{})
	g.declare("MetaProperty", &declOptions{})
	g.declare("SyntheticExpression", &declOptions{})
	g.declare("SatisfiesExpression", &declOptions{})

	// Misc
	g.declare("TemplateSpan", &declOptions{})
	g.declare("SemicolonClassElement", &declOptions{})

	// Element
	g.declare("Block", &declOptions{})
	g.declare("EmptyStatement", &declOptions{})
	g.declare("VariableStatement", &declOptions{})
	g.declare("ExpressionStatement", &declOptions{})
	g.declare("IfStatement", &declOptions{})
	g.declare("DoStatement", &declOptions{})
	g.declare("WhileStatement", &declOptions{})
	g.declare("ForStatement", &declOptions{})
	g.declare("ForInStatement", &declOptions{})
	g.declare("ForOfStatement", &declOptions{})
	g.declare("ContinueStatement", &declOptions{})
	g.declare("BreakStatement", &declOptions{})
	g.declare("ReturnStatement", &declOptions{})
	g.declare("WithStatement", &declOptions{})
	g.declare("SwitchStatement", &declOptions{})
	g.declare("LabeledStatement", &declOptions{})
	g.declare("ThrowStatement", &declOptions{})
	g.declare("TryStatement", &declOptions{})
	g.declare("DebuggerStatement", &declOptions{})
	g.declare("VariableDeclaration", &declOptions{})
	g.declare("VariableDeclarationList", &declOptions{})
	g.declare("FunctionDeclaration", &declOptions{})
	g.declare("ClassDeclaration", &declOptions{})
	g.declare("InterfaceDeclaration", &declOptions{})
	g.declare("TypeAliasDeclaration", &declOptions{})
	g.declare("EnumDeclaration", &declOptions{})
	g.declare("ModuleDeclaration", &declOptions{})
	g.declare("ModuleBlock", &declOptions{})
	g.declare("CaseBlock", &declOptions{})
	g.declare("NamespaceExportDeclaration", &declOptions{})
	g.declare("ImportEqualsDeclaration", &declOptions{})
	g.declare("ImportDeclaration", &declOptions{})
	g.declare("ImportClause", &declOptions{})
	g.declare("NamespaceImport", &declOptions{})
	g.declare("NamedImports", &declOptions{})
	g.declare("ImportSpecifier", &declOptions{})
	g.declare("ExportAssignment", &declOptions{})
	g.declare("ExportDeclaration", &declOptions{})
	g.declare("NamedExports", &declOptions{})
	g.declare("NamespaceExport", &declOptions{})
	g.declare("ExportSpecifier", &declOptions{})
	g.declare("MissingDeclaration", &declOptions{})

	// Module references
	g.declare("ExternalModuleReference", &declOptions{})

	// JSX
	g.declare("JsxElement", &declOptions{})
	g.declare("JsxSelfClosingElement", &declOptions{})
	g.declare("JsxOpeningElement", &declOptions{})
	g.declare("JsxClosingElement", &declOptions{})
	g.declare("JsxFragment", &declOptions{})
	g.declare("JsxOpeningFragment", &declOptions{})
	g.declare("JsxClosingFragment", &declOptions{})
	g.declare("JsxAttribute", &declOptions{})
	g.declare("JsxAttributes", &declOptions{})
	g.declare("JsxSpreadAttribute", &declOptions{})
	g.declare("JsxExpression", &declOptions{})
	g.declare("JsxNamespacedName", &declOptions{})

	// Clauses
	g.declare("CaseClause", &declOptions{})
	g.declare("DefaultClause", &declOptions{})
	g.declare("HeritageClause", &declOptions{})
	g.declare("CatchClause", &declOptions{})

	// Import attributes
	g.declare("ImportAttributes", &declOptions{})
	g.declare("ImportAttribute", &declOptions{})

	// Property assignments
	g.declare("PropertyAssignment", &declOptions{})
	g.declare("ShorthandPropertyAssignment", &declOptions{})
	g.declare("SpreadAssignment", &declOptions{})

	// Enum
	g.declare("EnumMember", &declOptions{})

	// Top-level nodes
	g.declare("SourceFile", &declOptions{})
	g.declare("Bundle", &declOptions{})

	// JSDoc nodes
	firstJSDocNode := g.declare("JSDocTypeExpression", &declOptions{})
	g.declare("JSDocNameReference", &declOptions{})
	g.declare("JSDocMemberName", &declOptions{})  // C#p
	g.declare("JSDocAllType", &declOptions{})     // The * type
	g.declare("JSDocUnknownType", &declOptions{}) // The ? type
	g.declare("JSDocNullableType", &declOptions{})
	g.declare("JSDocNonNullableType", &declOptions{})
	g.declare("JSDocOptionalType", &declOptions{})
	g.declare("JSDocFunctionType", &declOptions{})
	g.declare("JSDocVariadicType", &declOptions{})
	g.declare("JSDocNamepathType", &declOptions{}) // https://jsdoc.app/about-namepaths.html
	g.declare("JSDoc", &declOptions{})
	g.declare("JSDocText", &declOptions{})
	g.declare("JSDocTypeLiteral", &declOptions{})
	g.declare("JSDocSignature", &declOptions{})
	g.declare("JSDocLink", &declOptions{})
	g.declare("JSDocLinkCode", &declOptions{})
	g.declare("JSDocLinkPlain", &declOptions{})
	firstJSDocTagNode := g.declare("JSDocTag", &declOptions{})
	g.declare("JSDocAugmentsTag", &declOptions{})
	g.declare("JSDocImplementsTag", &declOptions{})
	g.declare("JSDocAuthorTag", &declOptions{})
	g.declare("JSDocDeprecatedTag", &declOptions{})
	g.declare("JSDocImmediateTag", &declOptions{})
	g.declare("JSDocClassTag", &declOptions{})
	g.declare("JSDocPublicTag", &declOptions{})
	g.declare("JSDocPrivateTag", &declOptions{})
	g.declare("JSDocProtectedTag", &declOptions{})
	g.declare("JSDocReadonlyTag", &declOptions{})
	g.declare("JSDocOverrideTag", &declOptions{})
	g.declare("JSDocCallbackTag", &declOptions{})
	g.declare("JSDocOverloadTag", &declOptions{})
	g.declare("JSDocEnumTag", &declOptions{})
	g.declare("JSDocParameterTag", &declOptions{})
	g.declare("JSDocReturnTag", &declOptions{})
	g.declare("JSDocThisTag", &declOptions{})
	g.declare("JSDocTypeTag", &declOptions{})
	g.declare("JSDocTemplateTag", &declOptions{})
	g.declare("JSDocTypedefTag", &declOptions{})
	g.declare("JSDocSeeTag", &declOptions{})
	g.declare("JSDocPropertyTag", &declOptions{})
	g.declare("JSDocThrowsTag", &declOptions{})
	g.declare("JSDocSatisfiesTag", &declOptions{})
	lastJSDocNode := g.declare("JSDocImportTag", &declOptions{})
	lastJSDocTagNode := lastJSDocNode

	// Synthesized list
	g.declare("SyntaxList", &declOptions{})

	// Transformation nodes
	g.declare("NotEmittedStatement", &declOptions{})
	g.declare("PartiallyEmittedExpression", &declOptions{})
	g.declare("CommaListExpression", &declOptions{})
	g.declare("SyntheticReferenceExpression", &declOptions{})

	// Markers

	g.createGroup("Assignment", firstAssignment, lastAssignment)
	g.createGroup("CompoundAssignment", firstCompoundAssignment, lastCompoundAssignment)
	g.createGroup("ReservedWord", firstReservedWord, lastReservedWord)
	g.createGroup("Keyword", firstKeyword, lastKeyword)
	g.createGroup("FutureReservedWord", firstFutureReservedWord, lastFutureReservedWord)
	g.createGroup("TypeNode", firstTypeNode, lastTypeNode)
	g.createGroup("Punctuation", firstPunctuation, lastPunctuation)
	g.createGroup("Token", firstToken, lastToken)
	g.createGroup("LiteralToken", firstLiteralToken, lastLiteralToken)
	g.createGroup("TemplateToken", firstTemplateToken, lastTemplateToken)
	g.createGroup("BinaryOperator", firstBinaryOperator, lastBinaryOperator)
	g.createGroup("Node", firstNode, g.nodes[len(g.nodes)-1])
	g.createGroup("JSDocNode", firstJSDocNode, lastJSDocNode)
	g.createGroup("JSDocTagNode", firstJSDocTagNode, lastJSDocTagNode)
	g.createGroup("ContextualKeyword", firstContextualKeyword, lastContextualKeyword)

	fmt.Fprintln(w, "package ast")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "import \"fmt\"")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "type SyntaxKind int16")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "const (")

	for i, n := range g.nodes {
		if i == 0 {
			fmt.Fprintf(w, "\tSyntaxKind%s SyntaxKind = iota\n", n.name)
		} else {
			fmt.Fprintf(w, "\tSyntaxKind%s\n", n.name)
		}
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "\tSyntaxKindCount")
	fmt.Fprintln(w)

	for _, g := range g.groups {
		fmt.Fprintf(w, "\t%sStart = SyntaxKind%s\n", g.name, g.start.name)
		fmt.Fprintf(w, "\t%sEnd = SyntaxKind%s\n", g.name, g.end.name)
	}

	fmt.Fprintln(w, ")")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "var syntaxKindNames = [SyntaxKindCount+1]string{")
	for _, n := range g.nodes {
		fmt.Fprintf(w, "\tSyntaxKind%s: %q,\n", n.name, n.name)
	}
	fmt.Fprintln(w, "\tSyntaxKindCount: \"SyntaxKindCount\",")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "func (k SyntaxKind) String() string {")
	fmt.Fprintln(w, "\tif k < 0 || k >= SyntaxKindCount {")
	fmt.Fprintf(w, "%s", "\t\treturn fmt.Sprintf(\"SyntaxKind(%d)\", k)")
	fmt.Fprintln(w, "\t}")
	fmt.Fprintln(w, "\treturn syntaxKindNames[k]")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "type NodeFlags uint32")
	fmt.Fprintln(w, "type NodeID uint32")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "type TextRange struct {")
	fmt.Fprintln(w, "\tpos int32")
	fmt.Fprintln(w, "\tend int32")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "func NewTextRange(pos, end int32) TextRange { return TextRange{pos: pos, end: end} }")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "func (r TextRange) Pos() int32 { return r.pos }")
	fmt.Fprintln(w, "func (r TextRange) End() int32 { return r.end }")
	fmt.Fprintln(w, "func (r TextRange) Len() int32 { return r.end - r.pos }")
	fmt.Fprintln(w, "func (r TextRange) ContainsInclusive(pos int32) bool { return r.pos <= pos && pos <= r.end }")
	fmt.Fprintln(w)

	// TODO
	fmt.Fprintln(w, "type NodeData interface{")
	fmt.Fprintln(w, "\tAsNode() *Node")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "type Node struct {")
	fmt.Fprintln(w, "\tkind SyntaxKind")
	fmt.Fprintln(w, "\tflags  NodeFlags")
	fmt.Fprintln(w, "\tloc TextRange")
	fmt.Fprintln(w, "\tid NodeID")
	fmt.Fprintln(w, "\tparent *Node")
	fmt.Fprintln(w, "\tdata NodeData")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "func (n *Node) Pos() int32 { return n.loc.Pos() }")

	// TODO
	fmt.Fprintln(w, "type NodeDefault struct {\nNode\n}")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "func (n *NodeDefault) AsNode() *Node { return &n.Node }")
	fmt.Fprintln(w)

	// TODO
	fmt.Fprintln(w, "type NodeBase struct {\nNodeDefault\n} ")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "type Factory struct { // TODO")
	for _, n := range g.nodes {
		if n.opts != nil && n.opts.poolAllocate {
			fmt.Fprintf(w, "\t_%sPool pool[%s]\n", n.name, n.name)
		}
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	for _, n := range g.nodes {
		n.Generate(w)
	}

	accessExpression := newNodeUnion("AccessExpression", propertyAccessExpression, elementAccessExpression)
	g.generateNodeUnion(w, accessExpression)
}

type generator struct {
	nodes       []*syntaxKind
	nodesByName map[string]*syntaxKind
	groups      []*group
}

type group struct {
	name  string
	start *syntaxKind
	end   *syntaxKind
}

func (g *generator) createGroup(name string, start, end *syntaxKind) {
	g.groups = append(g.groups, &group{name: name, start: start, end: end})
}

type declOptions struct {
	poolAllocate bool

	embedded any

	children []field
}

func (g *generator) declareKindOnly(name string) *syntaxKind {
	return g.declare(name, nil)
}

func (g *generator) declare(name string, opts *declOptions) *syntaxKind {
	if g.nodesByName == nil {
		g.nodesByName = make(map[string]*syntaxKind)
	}

	n := &syntaxKind{
		name: name,
		opts: opts,
	}
	g.nodes = append(g.nodes, n)
	g.nodesByName[name] = n
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

func newNodeUnion(name string, types ...*syntaxKind) *nodeUnion {
	return &nodeUnion{name: name, types: types}
}

func (g *generator) generateNodeUnion(w io.Writer, u *nodeUnion) {
	fmt.Fprintf(w, "type %s = Node // ", u.name)
	for i, t := range u.types {
		if i > 0 {
			fmt.Fprint(w, " | ")
		}
		fmt.Fprintf(w, "%s", t.name)
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "var is%sTable = [SyntaxKindCount]bool{\n", u.name)
	for _, t := range u.types {
		fmt.Fprintf(w, "\tSyntaxKind%s: true,\n", t.name)
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	fmt.Fprintf(w, "func is%sKind(kind SyntaxKind) bool { return is%sTable[kind] }\n", u.name, u.name)
	fmt.Fprintf(w, "func is%s(n *Node) bool { return is%sTable[n.kind] }\n", u.name, u.name)
	fmt.Fprintf(w, "func assertIs%s(n *Node) {\n", u.name)
	fmt.Fprintf(w, "if !is%s(n) {\n", u.name)
	fmt.Fprintf(w, "panic(\"expected %s, got \" + n.kind.String())\n", u.name)
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintln(w)
}

type syntaxKind struct {
	name   string
	opts   *declOptions
	fields []field
}

var _ genType = (*syntaxKind)(nil)

type field struct {
	name string
	typ  genType
}

func (n *syntaxKind) Name() string { return n.name }

func (n *syntaxKind) Generate(w io.Writer) {
	if n.opts == nil {
		return // not a real node; no code
	}

	fmt.Fprintln(w)
	fmt.Fprintf(w, "type %s struct {\n", n.name)
	fmt.Fprintln(w, "\tNodeBase")
	// TODO: children
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	fmt.Fprintf(w, "func (n *Node) As%s() *%s { return n.data.(*%s) }\n", n.name, n.name, n.name)
	fmt.Fprintf(w, "func (n *Node) Is%s() bool { return n.kind == SyntaxKind%s }\n", n.name, n.name)
	fmt.Fprintln(w)

	printNewParams := func() {
		// TODO
	}
	printNewArgs := func() {
		// TODO
	}

	// TODO: accept and set children
	fmt.Fprintf(w, "func (n *%s) set(", n.name)
	printNewParams()
	fmt.Fprintln(w, ") {")
	// TODO: generate optimal assignment
	fmt.Fprintf(w, "\t*n = %s{}\n", n.name)
	fmt.Fprintln(w, "\tn.kind = SyntaxKind"+n.name)
	fmt.Fprintln(w, "\tn.data = n")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	fmt.Fprintf(w, "func (n *%s) Kind() SyntaxKind { return SyntaxKind%s }\n", n.name, n.name)
	fmt.Fprintln(w)

	fmt.Fprintf(w, "func New%s(", n.name)
	printNewParams()
	fmt.Fprintln(w, ") *Node {")
	fmt.Fprintf(w, "\tn := &%s{}\n", n.name)
	fmt.Fprintf(w, "\tn.set(")
	printNewArgs()
	fmt.Fprintln(w, ")")
	fmt.Fprintln(w, "\treturn n.AsNode()")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	// TODO: params
	// TODO: return *Node, call a helper func that converts to *Node and sets kind
	fmt.Fprintf(w, "func (f *Factory) New%s() *Node {\n", n.name)
	if n.opts.poolAllocate {
		fmt.Fprintf(w, "\tn := f._%sPool.allocate()\n", n.name)
		fmt.Fprintf(w, "\tn.set(")
		printNewArgs()
		fmt.Fprintln(w, ")")
		fmt.Fprintln(w, "\treturn n.AsNode()")
	} else {
		fmt.Fprintf(w, "\treturn New%s(", n.name)
		printNewArgs()
		fmt.Fprintln(w, ")")
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintln(w)
}
