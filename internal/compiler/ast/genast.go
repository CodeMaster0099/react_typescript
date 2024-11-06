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
	firstBinaryOperator := d.declareKindOnly("LessThanToken")
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
	lastCompoundAssignment := lastPunctuation
	lastBinaryOperator := lastPunctuation

	d.declare("Identifier", &declOptions{
		poolAllocate: true,
	})
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
	firstContextualKeyword := d.declareKindOnly("AbstractKeyword")
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
	lastContextualKeyword := lastKeyword

	// Parse tree nodes
	// Names
	firstNode := d.declare("QualifiedName", &declOptions{})
	d.declare("ComputedPropertyName", &declOptions{})

	// Lists
	d.declare("ModifierList", &declOptions{})
	d.declare("TypeParameterList", &declOptions{})
	d.declare("TypeArgumentList", &declOptions{})

	// Signature elements
	d.declare("TypeParameter", &declOptions{})
	d.declare("Parameter", &declOptions{})
	d.declare("Decorator", &declOptions{})

	// TypeMember
	d.declare("PropertySignature", &declOptions{})
	d.declare("PropertyDeclaration", &declOptions{})
	d.declare("MethodSignature", &declOptions{})
	d.declare("MethodDeclaration", &declOptions{})
	d.declare("ClassStaticBlockDeclaration", &declOptions{})
	d.declare("Constructor", &declOptions{})
	d.declare("GetAccessor", &declOptions{})
	d.declare("SetAccessor", &declOptions{})
	d.declare("CallSignature", &declOptions{})
	d.declare("ConstructSignature", &declOptions{})
	d.declare("IndexSignature", &declOptions{})

	// Type
	firstTypeNode := d.declare("TypePredicate", &declOptions{})
	d.declare("TypeReference", &declOptions{})
	d.declare("FunctionType", &declOptions{})
	d.declare("ConstructorType", &declOptions{})
	d.declare("TypeQuery", &declOptions{})
	d.declare("TypeLiteral", &declOptions{})
	d.declare("ArrayType", &declOptions{})
	d.declare("TupleType", &declOptions{})
	d.declare("OptionalType", &declOptions{})
	d.declare("RestType", &declOptions{})
	d.declare("UnionType", &declOptions{})
	d.declare("IntersectionType", &declOptions{})
	d.declare("ConditionalType", &declOptions{})
	d.declare("InferType", &declOptions{})
	d.declare("ParenthesizedType", &declOptions{})
	d.declare("ThisType", &declOptions{})
	d.declare("TypeOperator", &declOptions{})
	d.declare("IndexedAccessType", &declOptions{})
	d.declare("MappedType", &declOptions{})
	d.declare("LiteralType", &declOptions{})
	d.declare("NamedTupleMember", &declOptions{})
	d.declare("TemplateLiteralType", &declOptions{})
	d.declare("TemplateLiteralTypeSpan", &declOptions{})
	lastTypeNode := d.declare("ImportType", &declOptions{})

	// Binding patterns
	d.declare("ObjectBindingPattern", &declOptions{})
	d.declare("ArrayBindingPattern", &declOptions{})
	d.declare("BindingElement", &declOptions{})

	// Expression
	d.declare("ArrayLiteralExpression", &declOptions{})
	d.declare("ObjectLiteralExpression", &declOptions{})
	propertyAccessExpression := d.declare("PropertyAccessExpression", &declOptions{})
	elementAccessExpression := d.declare("ElementAccessExpression", &declOptions{})
	d.declare("CallExpression", &declOptions{})
	d.declare("NewExpression", &declOptions{})
	d.declare("TaggedTemplateExpression", &declOptions{})
	d.declare("TypeAssertionExpression", &declOptions{})
	d.declare("ParenthesizedExpression", &declOptions{})
	d.declare("FunctionExpression", &declOptions{})
	d.declare("ArrowFunction", &declOptions{})
	d.declare("DeleteExpression", &declOptions{})
	d.declare("TypeOfExpression", &declOptions{})
	d.declare("VoidExpression", &declOptions{})
	d.declare("AwaitExpression", &declOptions{})
	d.declare("PrefixUnaryExpression", &declOptions{})
	d.declare("PostfixUnaryExpression", &declOptions{})
	d.declare("BinaryExpression", &declOptions{})
	d.declare("ConditionalExpression", &declOptions{})
	d.declare("TemplateExpression", &declOptions{})
	d.declare("YieldExpression", &declOptions{})
	d.declare("SpreadElement", &declOptions{})
	d.declare("ClassExpression", &declOptions{})
	d.declare("OmittedExpression", &declOptions{})
	d.declare("ExpressionWithTypeArguments", &declOptions{})
	d.declare("AsExpression", &declOptions{})
	d.declare("NonNullExpression", &declOptions{})
	d.declare("MetaProperty", &declOptions{})
	d.declare("SyntheticExpression", &declOptions{})
	d.declare("SatisfiesExpression", &declOptions{})

	// Misc
	d.declare("TemplateSpan", &declOptions{})
	d.declare("SemicolonClassElement", &declOptions{})

	// Element
	d.declare("Block", &declOptions{})
	d.declare("EmptyStatement", &declOptions{})
	d.declare("VariableStatement", &declOptions{})
	d.declare("ExpressionStatement", &declOptions{})
	d.declare("IfStatement", &declOptions{})
	d.declare("DoStatement", &declOptions{})
	d.declare("WhileStatement", &declOptions{})
	d.declare("ForStatement", &declOptions{})
	d.declare("ForInStatement", &declOptions{})
	d.declare("ForOfStatement", &declOptions{})
	d.declare("ContinueStatement", &declOptions{})
	d.declare("BreakStatement", &declOptions{})
	d.declare("ReturnStatement", &declOptions{})
	d.declare("WithStatement", &declOptions{})
	d.declare("SwitchStatement", &declOptions{})
	d.declare("LabeledStatement", &declOptions{})
	d.declare("ThrowStatement", &declOptions{})
	d.declare("TryStatement", &declOptions{})
	d.declare("DebuggerStatement", &declOptions{})
	d.declare("VariableDeclaration", &declOptions{})
	d.declare("VariableDeclarationList", &declOptions{})
	d.declare("FunctionDeclaration", &declOptions{})
	d.declare("ClassDeclaration", &declOptions{})
	d.declare("InterfaceDeclaration", &declOptions{})
	d.declare("TypeAliasDeclaration", &declOptions{})
	d.declare("EnumDeclaration", &declOptions{})
	d.declare("ModuleDeclaration", &declOptions{})
	d.declare("ModuleBlock", &declOptions{})
	d.declare("CaseBlock", &declOptions{})
	d.declare("NamespaceExportDeclaration", &declOptions{})
	d.declare("ImportEqualsDeclaration", &declOptions{})
	d.declare("ImportDeclaration", &declOptions{})
	d.declare("ImportClause", &declOptions{})
	d.declare("NamespaceImport", &declOptions{})
	d.declare("NamedImports", &declOptions{})
	d.declare("ImportSpecifier", &declOptions{})
	d.declare("ExportAssignment", &declOptions{})
	d.declare("ExportDeclaration", &declOptions{})
	d.declare("NamedExports", &declOptions{})
	d.declare("NamespaceExport", &declOptions{})
	d.declare("ExportSpecifier", &declOptions{})
	d.declare("MissingDeclaration", &declOptions{})

	// Module references
	d.declare("ExternalModuleReference", &declOptions{})

	// JSX
	d.declare("JsxElement", &declOptions{})
	d.declare("JsxSelfClosingElement", &declOptions{})
	d.declare("JsxOpeningElement", &declOptions{})
	d.declare("JsxClosingElement", &declOptions{})
	d.declare("JsxFragment", &declOptions{})
	d.declare("JsxOpeningFragment", &declOptions{})
	d.declare("JsxClosingFragment", &declOptions{})
	d.declare("JsxAttribute", &declOptions{})
	d.declare("JsxAttributes", &declOptions{})
	d.declare("JsxSpreadAttribute", &declOptions{})
	d.declare("JsxExpression", &declOptions{})
	d.declare("JsxNamespacedName", &declOptions{})

	// Clauses
	d.declare("CaseClause", &declOptions{})
	d.declare("DefaultClause", &declOptions{})
	d.declare("HeritageClause", &declOptions{})
	d.declare("CatchClause", &declOptions{})

	// Import attributes
	d.declare("ImportAttributes", &declOptions{})
	d.declare("ImportAttribute", &declOptions{})

	// Property assignments
	d.declare("PropertyAssignment", &declOptions{})
	d.declare("ShorthandPropertyAssignment", &declOptions{})
	d.declare("SpreadAssignment", &declOptions{})

	// Enum
	d.declare("EnumMember", &declOptions{})

	// Top-level nodes
	d.declare("SourceFile", &declOptions{})
	d.declare("Bundle", &declOptions{})

	// JSDoc nodes
	firstJSDocNode := d.declare("JSDocTypeExpression", &declOptions{})
	d.declare("JSDocNameReference", &declOptions{})
	d.declare("JSDocMemberName", &declOptions{})  // C#p
	d.declare("JSDocAllType", &declOptions{})     // The * type
	d.declare("JSDocUnknownType", &declOptions{}) // The ? type
	d.declare("JSDocNullableType", &declOptions{})
	d.declare("JSDocNonNullableType", &declOptions{})
	d.declare("JSDocOptionalType", &declOptions{})
	d.declare("JSDocFunctionType", &declOptions{})
	d.declare("JSDocVariadicType", &declOptions{})
	d.declare("JSDocNamepathType", &declOptions{}) // https://jsdoc.app/about-namepaths.html
	d.declare("JSDoc", &declOptions{})
	d.declare("JSDocText", &declOptions{})
	d.declare("JSDocTypeLiteral", &declOptions{})
	d.declare("JSDocSignature", &declOptions{})
	d.declare("JSDocLink", &declOptions{})
	d.declare("JSDocLinkCode", &declOptions{})
	d.declare("JSDocLinkPlain", &declOptions{})
	firstJSDocTagNode := d.declare("JSDocTag", &declOptions{})
	d.declare("JSDocAugmentsTag", &declOptions{})
	d.declare("JSDocImplementsTag", &declOptions{})
	d.declare("JSDocAuthorTag", &declOptions{})
	d.declare("JSDocDeprecatedTag", &declOptions{})
	d.declare("JSDocImmediateTag", &declOptions{})
	d.declare("JSDocClassTag", &declOptions{})
	d.declare("JSDocPublicTag", &declOptions{})
	d.declare("JSDocPrivateTag", &declOptions{})
	d.declare("JSDocProtectedTag", &declOptions{})
	d.declare("JSDocReadonlyTag", &declOptions{})
	d.declare("JSDocOverrideTag", &declOptions{})
	d.declare("JSDocCallbackTag", &declOptions{})
	d.declare("JSDocOverloadTag", &declOptions{})
	d.declare("JSDocEnumTag", &declOptions{})
	d.declare("JSDocParameterTag", &declOptions{})
	d.declare("JSDocReturnTag", &declOptions{})
	d.declare("JSDocThisTag", &declOptions{})
	d.declare("JSDocTypeTag", &declOptions{})
	d.declare("JSDocTemplateTag", &declOptions{})
	d.declare("JSDocTypedefTag", &declOptions{})
	d.declare("JSDocSeeTag", &declOptions{})
	d.declare("JSDocPropertyTag", &declOptions{})
	d.declare("JSDocThrowsTag", &declOptions{})
	d.declare("JSDocSatisfiesTag", &declOptions{})
	lastJSDocNode := d.declare("JSDocImportTag", &declOptions{})
	lastJSDocTagNode := lastJSDocNode

	// Synthesized list
	d.declare("SyntaxList", &declOptions{})

	// Transformation nodes
	d.declare("NotEmittedStatement", &declOptions{})
	d.declare("PartiallyEmittedExpression", &declOptions{})
	d.declare("CommaListExpression", &declOptions{})
	d.declare("SyntheticReferenceExpression", &declOptions{})

	// Markers

	d.createGroup("Assignment", firstAssignment, lastAssignment)
	d.createGroup("CompoundAssignment", firstCompoundAssignment, lastCompoundAssignment)
	d.createGroup("ReservedWord", firstReservedWord, lastReservedWord)
	d.createGroup("Keyword", firstKeyword, lastKeyword)
	d.createGroup("FutureReservedWord", firstFutureReservedWord, lastFutureReservedWord)
	d.createGroup("TypeNode", firstTypeNode, lastTypeNode)
	d.createGroup("Punctuation", firstPunctuation, lastPunctuation)
	d.createGroup("Token", firstToken, lastToken)
	d.createGroup("LiteralToken", firstLiteralToken, lastLiteralToken)
	d.createGroup("TemplateToken", firstTemplateToken, lastTemplateToken)
	d.createGroup("BinaryOperator", firstBinaryOperator, lastBinaryOperator)
	d.createGroup("Node", firstNode, d.nodes[len(d.nodes)-1])
	d.createGroup("JSDocNode", firstJSDocNode, lastJSDocNode)
	d.createGroup("JSDocTagNode", firstJSDocTagNode, lastJSDocTagNode)
	d.createGroup("ContextualKeyword", firstContextualKeyword, lastContextualKeyword)

	fmt.Fprintln(w, "package ast")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "import \"fmt\"")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "type SyntaxKind int16")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "const (")

	for i, n := range d.nodes {
		if i == 0 {
			fmt.Fprintf(w, "\tSyntaxKind%s SyntaxKind = iota\n", n.name)
		} else {
			fmt.Fprintf(w, "\tSyntaxKind%s\n", n.name)
		}
	}

	fmt.Fprintln(w)
	fmt.Fprintln(w, "\tSyntaxKindCount")
	fmt.Fprintln(w)

	for _, g := range d.groups {
		fmt.Fprintf(w, "\t%sStart = SyntaxKind%s\n", g.name, g.start.name)
		fmt.Fprintf(w, "\t%sEnd = SyntaxKind%s\n", g.name, g.end.name)
	}

	fmt.Fprintln(w, ")")
	fmt.Fprintln(w)

	fmt.Fprintln(w, "var syntaxKindNames = [SyntaxKindCount+1]string{")
	for _, n := range d.nodes {
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
	for _, n := range d.nodes {
		if n.opts != nil && n.opts.poolAllocate {
			fmt.Fprintf(w, "\t_%sPool pool[%s]\n", n.name, n.name)
		}
	}
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w)

	for _, n := range d.nodes {
		n.Generate(w)
	}

	accessExpression := newNodeUnion("AccessExpression", propertyAccessExpression, elementAccessExpression)
	d.generateNodeUnion(w, accessExpression)
}

type declarer struct {
	nodes       []*syntaxKind
	nodesByName map[string]*syntaxKind
	groups      []*group
	count       int
}

type group struct {
	name  string
	start *syntaxKind
	end   *syntaxKind
}

func (d *declarer) createGroup(name string, start, end *syntaxKind) {
	if start.kind > end.kind {
		panic("start > end")
	}
	d.groups = append(d.groups, &group{name: name, start: start, end: end})
}

type declOptions struct {
	poolAllocate bool

	embedded any

	children []field
}

func (d *declarer) declareKindOnly(name string) *syntaxKind {
	return d.declare(name, nil)
}

func (d *declarer) declare(name string, opts *declOptions) *syntaxKind {
	if d.nodesByName == nil {
		d.nodesByName = make(map[string]*syntaxKind)
	}

	n := &syntaxKind{
		kind: d.count,
		name: name,
		opts: opts,
	}
	d.nodes = append(d.nodes, n)
	d.nodesByName[name] = n
	d.count++
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

func (d *declarer) generateNodeUnion(w io.Writer, u *nodeUnion) {
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
	kind   int
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
