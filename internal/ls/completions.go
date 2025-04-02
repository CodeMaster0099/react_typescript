package ls

import (
	"slices"

	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/astnav"
	"github.com/microsoft/typescript-go/internal/checker"
	"github.com/microsoft/typescript-go/internal/compiler"
	"github.com/microsoft/typescript-go/internal/core"
	"github.com/microsoft/typescript-go/internal/lsp/lsproto"
	"github.com/microsoft/typescript-go/internal/tspath"
)

func (l *LanguageService) ProvideCompletion(fileName string, position int, context *lsproto.CompletionContext) *lsproto.CompletionList {
	program, file := l.getProgramAndFile(fileName)
	node := astnav.GetTouchingPropertyName(file, position)
	if node.Kind == ast.KindSourceFile {
		return nil
	}
	return l.getCompletionsAtPosition(program, file, position, context, nil /*preferences*/) // !!! get preferences
}

type completionData struct {
	// !!!
}

type importStatementCompletionInfo struct {
	// !!!
}

// If we're after the `=` sign but no identifier has been typed yet,
// value will be `true` but initializer will be `nil`.
type jsxInitializer struct {
	isInitializer bool
	initializer   *ast.Identifier
}

type KeywordCompletionFilters int

const (
	KeywordCompletionFiltersNone                         KeywordCompletionFilters = iota // No keywords
	KeywordCompletionFiltersAll                                                          // Every possible kewyord
	KeywordCompletionFiltersClassElementKeywords                                         // Keywords inside class body
	KeywordCompletionFiltersInterfaceElementKeywords                                     // Keywords inside interface body
	KeywordCompletionFiltersConstructorParameterKeywords                                 // Keywords at constructor parameter
	KeywordCompletionFiltersFunctionLikeBodyKeywords                                     // Keywords at function like body
	KeywordCompletionFiltersTypeAssertionKeywords
	KeywordCompletionFiltersTypeKeywords
	KeywordCompletionFiltersTypeKeyword // Literally just `type`
	KeywordCompletionFiltersLast        = KeywordCompletionFiltersTypeKeyword
)

type CompletionKind int

const (
	CompletionKindNone CompletionKind = iota
	CompletionKindObjectPropertyDeclaration
	CompletionKindGlobal
	CompletionKindPropertyAccess
	CompletionKindMemberLike
	CompletionKindString
)

// All commit characters, valid when `isNewIdentifierLocation` is false.
var allCommitCharacters = []string{".", ",", ";"}

// Commit characters valid at expression positions where we could be inside a parameter list.
var noCommaCommitCharacters = []string{".", ";"}

type sortText string

const (
	SortTextLocalDeclarationPriority         sortText = "10"
	sortTextLocationPriority                 sortText = "11"
	sortTextOptionalMember                   sortText = "12"
	sortTextMemberDeclaredBySpreadAssignment sortText = "13"
	sortTextSuggestedClassMembers            sortText = "14"
	sortTextGlobalsOrKeywords                sortText = "15"
	sortTextAutoImportSuggestions            sortText = "16"
	sortTextClassMemberSnippets              sortText = "17"
	sortTextJavascriptIdentifiers            sortText = "18"
)

// !!! sort text transformations

type symbolOriginInfoKind int

const (
	symbolOriginInfoKindThisType symbolOriginInfoKind = 1 << iota
	symbolOriginInfoKindSymbolMember
	symbolOriginInfoKindExport
	symbolOriginInfoKindPromise
	symbolOriginInfoKindNullable
	symbolOriginInfoKindResolvedExport
	symbolOriginInfoKindTypeOnlyAlias
	symbolOriginInfoKindObjectLiteralMethod
	symbolOriginInfoKindIgnore
	symbolOriginInfoKindComputedPropertyName

	symbolOriginInfoKindSymbolMemberNoExport symbolOriginInfoKind = symbolOriginInfoKindSymbolMember
	symbolOriginInfoKindSymbolMemberExport                        = symbolOriginInfoKindSymbolMember | symbolOriginInfoKindExport
)

type symbolOriginInfo struct {
	kind              symbolOriginInfoKind
	isDefaultExport   bool
	isFromPackageJson bool
	fileName          string
}

func (l *LanguageService) getCompletionsAtPosition(
	program *compiler.Program,
	file *ast.SourceFile,
	position int,
	context *lsproto.CompletionContext,
	preferences *UserPreferences) *lsproto.CompletionList {
	previousToken, _ := getRelevantTokens(position, file)
	if context.TriggerCharacter != nil && !isInString(file, position, previousToken) && !isValidTrigger(file, *context.TriggerCharacter, previousToken, position) {
		return nil
	}

	if *context.TriggerCharacter == " " {
		// `isValidTrigger` ensures we are at `import |`
		if preferences.includeCompletionsForImportStatements {
			// !!! isMemberCompletion
			return &lsproto.CompletionList{
				IsIncomplete: true,
				ItemDefaults: &lsproto.CompletionItemDefaults{ // !!! do we need this if no entries? also, check if client supports item defaults
					CommitCharacters: getDefaultCommitCharacters(true /*isNewIdentifierLocation*/),
				},
			}
		}
		return nil
	}

	compilerOptions := program.GetCompilerOptions()
	checker := program.GetTypeChecker()

	// !!! see if incomplete completion list and continue or clean

	// !!! string literal completions

	// !!! label completions

	completionData := getCompletionData(program, file, position, preferences)
	// !!! get completion data

	// !!! transform data into completion list

	return nil // !!!
}

func getCompletionData(program *compiler.Program, file *ast.SourceFile, position int, preferences *UserPreferences) *completionData {
	typeChecker := program.GetTypeChecker()
	inCheckedFile := isCheckedFile(file, program.GetCompilerOptions())

	currentToken := astnav.GetTokenAtPosition(file, position)

	insideComment := isInComment(file, position, currentToken)

	insideJsDocTagTypeExpression := false
	insideJsDocImportTag := false
	isInSnippetScope := false
	if insideComment != nil {
		// !!! jsdoc
	}

	// The decision to provide completion depends on the contextToken, which is determined through the previousToken.
	// Note: 'previousToken' (and thus 'contextToken') can be undefined if we are the beginning of the file
	isJSOnlyLocation := !insideJsDocTagTypeExpression && !insideJsDocImportTag && ast.IsSourceFileJs(file)
	previousToken, contextToken := getRelevantTokens(position, file)

	// Find the node where completion is requested on.
	// Also determine whether we are trying to complete with members of that node
	// or attributes of a JSX tag.
	node := currentToken
	var propertyAccessToConvert *ast.PropertyAccessExpression
	isRightOfDot := false
	isRightOfQuestionDot := false
	isRightOfOpenTag := false
	isStartingCloseTag := false
	var jsxInitializer jsxInitializer
	isJsxIdentifierExpected := false
	var importStatementCompletion *importStatementCompletionInfo
	location := astnav.GetTouchingPropertyName(file, position)
	keywordFilters := KeywordCompletionFiltersNone
	isNewIdentifierLocation := false
	// !!!
	// flags := CompletionInfoFlagsNone
	var defaultCommitCharacters []string

	if contextToken != nil {
		// !!! import completions

		parent := contextToken.Parent
		if contextToken.Kind == ast.KindDotToken || contextToken.Kind == ast.KindQuestionDotToken {
			isRightOfDot = contextToken.Kind == ast.KindDotToken
			isRightOfQuestionDot = contextToken.Kind == ast.KindQuestionDotToken
			switch parent.Kind {
			case ast.KindPropertyAccessExpression:
				propertyAccessToConvert = parent.AsPropertyAccessExpression()
				node = propertyAccessToConvert.Expression
				leftMostAccessExpression := ast.GetLeftmostAccessExpression(parent)
				if ast.NodeIsMissing(leftMostAccessExpression) ||
					((ast.IsCallExpression(node) || ast.IsFunctionLike(node)) &&
						node.End() == contextToken.Pos() &&
						getLastChild(node, file).Kind != ast.KindCloseParenToken) {
					// This is likely dot from incorrectly parsed expression and user is starting to write spread
					// eg: Math.min(./**/)
					// const x = function (./**/) {}
					// ({./**/})
					return nil
				}
			case ast.KindQualifiedName:
				node = parent.AsQualifiedName().Left
			case ast.KindModuleDeclaration:
				node = parent.Name()
			case ast.KindImportType:
				node = parent
			case ast.KindMetaProperty:
				node = getFirstToken(parent, file)
				if node.Kind != ast.KindImportKeyword || node.Kind != ast.KindNewKeyword {
					panic("Unexpected token kind: " + node.Kind.String())
				}
			default:
				// There is nothing that precedes the dot, so this likely just a stray character
				// or leading into a '...' token. Just bail out instead.
				return nil
			}
		} else { // !!! else if (!importStatementCompletion)
			// <UI.Test /* completion position */ />
			// If the tagname is a property access expression, we will then walk up to the top most of property access expression.
			// Then, try to get a JSX container and its associated attributes type.
			if parent != nil && parent.Kind == ast.KindPropertyAccessExpression {
				contextToken = parent
				parent = parent.Parent
			}

			// Fix location
			if parent == location {
				switch currentToken.Kind {
				case ast.KindGreaterThanToken:
					if parent.Kind == ast.KindJsxElement || parent.Kind == ast.KindJsxOpeningElement {
						location = currentToken
					}
				case ast.KindSlashToken:
					if parent.Kind == ast.KindJsxSelfClosingElement {
						location = currentToken
					}
				}
			}

			switch parent.Kind {
			case ast.KindJsxClosingElement:
				if contextToken.Kind == ast.KindSlashToken {
					isStartingCloseTag = true
					location = contextToken
				}
			case ast.KindBinaryExpression:
				if !binaryExpressionMayBeOpenTag(parent.AsBinaryExpression()) {
					break
				}
				fallthrough
			case ast.KindJsxSelfClosingElement, ast.KindJsxElement, ast.KindJsxOpeningElement:
				isJsxIdentifierExpected = true
				if contextToken.Kind == ast.KindLessThanToken {
					isRightOfOpenTag = true
					location = contextToken
				}
			case ast.KindJsxExpression, ast.KindJsxSpreadAttribute:
				// First case is for `<div foo={true} [||] />` or `<div foo={true} [||] ></div>`,
				// `parent` will be `{true}` and `previousToken` will be `}`.
				// Second case is for `<div foo={true} t[||] ></div>`.
				// Second case must not match for `<div foo={undefine[||]}></div>`.
				if previousToken.Kind == ast.KindCloseBraceToken ||
					previousToken.Kind == ast.KindIdentifier && previousToken.Parent.Kind == ast.KindJsxAttribute {
					isJsxIdentifierExpected = true
				}
			case ast.KindJsxAttribute:
				// For `<div className="x" [||] ></div>`, `parent` will be JsxAttribute and `previousToken` will be its initializer.
				if parent.Initializer() == previousToken && previousToken.End() < position {
					isJsxIdentifierExpected = true
				} else {
					switch previousToken.Kind {
					case ast.KindEqualsToken:
						jsxInitializer.isInitializer = true
					case ast.KindIdentifier:
						isJsxIdentifierExpected = true
						// For `<div x=[|f/**/|]`, `parent` will be `x` and `previousToken.parent` will be `f` (which is its own JsxAttribute).
						// Note for `<div someBool f>` we don't want to treat this as a jsx inializer, instead it's the attribute name.
						if parent != previousToken.Parent &&
							parent.Initializer() == nil &&
							findChildOfKind(parent, ast.KindEqualsToken, file) != nil {
							jsxInitializer.initializer = previousToken.AsIdentifier()
						}
					}
				}
			}
		}
	}

	completionKind := CompletionKindNone
	hasUnresolvedAutoImports := false
	// This also gets mutated in nested-functions after the return
	var symbols []*ast.Symbol
	var symbolToOriginInfoMap map[ast.SymbolId]symbolOriginInfo
	var symbolToSortTextMap map[ast.SymbolId]sortText
	var importSpecifierResolver any // !!!
	seenPropertySymbols := make(map[ast.SymbolId]struct{})
	isTypeOnlyLocation := insideJsDocTagTypeExpression || insideJsDocImportTag ||
		importStatementCompletion != nil && ast.IsTypeOnlyImportOrExportDeclaration(location.Parent) ||
		!isContextTokenValueLocation(contextToken) &&
			(isPossiblyTypeArgumentPosition(contextToken, file, typeChecker) ||
				ast.IsPartOfTypeNode(location) ||
				isContextTokenTypeLocation(contextToken))
	var getModuleSpecifierResolutionHost any // !!!

	addSymbolOriginInfo := func(symbol *ast.Symbol, insertQuestionDot bool, insertAwait bool) {
		symbolId := ast.GetSymbolId(symbol)
		if insertAwait && core.AddToSeen(seenPropertySymbols, symbolId, struct{}{}) {
			symbolToOriginInfoMap[symbolId] = symbolOriginInfo{kind: getNullableSymbolOriginInfoKind(symbolOriginInfoKindPromise, insertQuestionDot)}
		} else if insertQuestionDot {
			symbolToOriginInfoMap[symbolId] = symbolOriginInfo{kind: symbolOriginInfoKindNullable}
		}
	}

	addSymbolSortInfo := func(symbol *ast.Symbol) {
		symbolId := ast.GetSymbolId(symbol)
		if isStaticProperty(symbol) {
			symbolToSortTextMap[symbolId] = SortTextLocalDeclarationPriority
		}
	}

	addPropertySymbol := func(symbol *ast.Symbol, insertAwait bool, insertQuestionDot bool) {
		// For a computed property with an accessible name like `Symbol.iterator`,
		// we'll add a completion for the *name* `Symbol` instead of for the property.
		// If this is e.g. [Symbol.iterator], add a completion for `Symbol`.
		computedPropertyName := core.FirstNonNil(symbol.Declarations, func(decl *ast.Node) *ast.Node {
			name := ast.GetNameOfDeclaration(decl)
			if name != nil && name.Kind == ast.KindComputedPropertyName {
				return name
			}
			return nil
		})

		if computedPropertyName != nil {
			leftMostName := getLeftMostName(computedPropertyName.Expression()) // The completion is for `Symbol`, not `iterator`.
			var nameSymbol *ast.Symbol
			if leftMostName != nil {
				nameSymbol = typeChecker.GetSymbolAtLocation(leftMostName)
			}
			// If this is nested like for `namespace N { export const sym = Symbol(); }`, we'll add the completion for `N`.
			var firstAccessibleSymbol *ast.Symbol
			if nameSymbol != nil {
				firstAccessibleSymbol = getFirstSymbolInChain(nameSymbol, contextToken, typeChecker)
			}
			var firstAccessibleSymbolId ast.SymbolId
			if firstAccessibleSymbol != nil {
				firstAccessibleSymbolId = ast.GetSymbolId(firstAccessibleSymbol)
			}
			if firstAccessibleSymbolId != 0 && core.AddToSeen(seenPropertySymbols, firstAccessibleSymbolId, struct{}{}) {
				index := len(symbols)
				symbols = append(symbols, firstAccessibleSymbol)
				moduleSymbol := firstAccessibleSymbol.Parent
				if moduleSymbol == nil ||
					!checker.IsExternalModuleSymbol(moduleSymbol) ||
					typeChecker.TryGetMemberInModuleExportsAndProperties(firstAccessibleSymbol.Name, moduleSymbol) != firstAccessibleSymbol {
					symbolToOriginInfoMap[ast.GetSymbolId(symbol)] = symbolOriginInfo{kind: getNullableSymbolOriginInfoKind(symbolOriginInfoKindSymbolMemberNoExport, insertQuestionDot)}
				} else {
					var fileName string
					if tspath.IsExternalModuleNameRelative(core.StripQuotes(moduleSymbol.Name)) {
						fileName = ast.GetSourceFileOfModule(moduleSymbol).FileName()
					}
					if importSpecifierResolver == nil { // !!! verify if this is right, depending on the type of importSpecifierResolver
						// !!!
						// importSpecifierResolver ||= codefix.createImportSpecifierResolver(sourceFile, program, host, preferences))
					}
					// !!!
					// const { moduleSpecifier } = importSpecifier.getModuleSpecifierForBestExportInfo(
					// 	[{
					// 		exportKind: ExportKind.Named,
					// 		moduleFileName: fileName,
					// 		isFromPackageJson: false,
					// 		moduleSymbol,
					// 		symbol: firstAccessibleSymbol,
					// 		targetFlags: skipAlias(firstAccessibleSymbol, typeChecker).flags,
					// 	}],
					// 	position,
					// 	isValidTypeOnlyAliasUseSite(location),
					// ) || {};
					// if (moduleSpecifier) {
					// 	const origin: SymbolOriginInfoResolvedExport = {
					// 		kind: getNullableSymbolOriginInfoKind(SymbolOriginInfoKind.SymbolMemberExport),
					// 		moduleSymbol,
					// 		isDefaultExport: false,
					// 		symbolName: firstAccessibleSymbol.name,
					// 		exportName: firstAccessibleSymbol.name,
					// 		fileName,
					// 		moduleSpecifier,
					// 	};
					// 	symbolToOriginInfoMap[index] = origin;
					// }
				}
			} else if _, ok := seenPropertySymbols[firstAccessibleSymbolId]; firstAccessibleSymbolId == 0 || !ok {
				symbols = append(symbols, symbol)
				addSymbolOriginInfo(symbol, insertQuestionDot, insertAwait)
				addSymbolSortInfo(symbol)
			}
		} else {
			symbols = append(symbols, symbol)
			addSymbolOriginInfo(symbol, insertQuestionDot, insertAwait)
			addSymbolSortInfo(symbol)
		}
	}

	addTypeProperties := func(t *checker.Type, insertAwait bool, insertQuestionDot bool) {
		if typeChecker.GetStringIndexType(t) != nil {
			isNewIdentifierLocation = true
			defaultCommitCharacters = make([]string, 0)
		}
		if isRightOfQuestionDot && len(typeChecker.GetCallSignatures(t)) != 0 {
			isNewIdentifierLocation = true
			if len(defaultCommitCharacters) == 0 {
				defaultCommitCharacters = slices.Clone(allCommitCharacters) // Only invalid commit character here would be `(`.
			}
		}

		var propertyAccess *ast.Node
		if node.Kind == ast.KindImportType {
			propertyAccess = node
		} else {
			propertyAccess = node.Parent
		}

		if inCheckedFile {
			for _, symbol := range typeChecker.GetApparentProperties(t) {
				if typeChecker.IsValidPropertyAccessForCompletions(propertyAccess, t, symbol) {
					addPropertySymbol(symbol, false /*insertAwait*/, insertQuestionDot)
				}
			}
		} else {
			// In javascript files, for union types, we don't just get the members that
			// the individual types have in common, we also include all the members that
			// each individual type has. This is because we're going to add all identifiers
			// anyways. So we might as well elevate the members that were at least part
			// of the individual types to a higher status since we know what they are.
			for _, symbol := range getPropertiesForCompletion(t, typeChecker) {
				if typeChecker.IsValidPropertyAccessForCompletions(propertyAccess, t, symbol) {
					symbols = append(symbols, symbol)
				}
			}
		}

		if insertAwait {
			promiseType := typeChecker.GetPromisedTypeOfPromise(t)
			if promiseType != nil {
				for _, symbol := range typeChecker.GetApparentProperties(promiseType) {
					if typeChecker.IsValidPropertyAccessForCompletions(propertyAccess, promiseType, symbol) {
						addPropertySymbol(symbol, true /*insertAwait*/, insertQuestionDot)
					}
				}
			}
		}
	}

	getTypeScriptMemberSymbols := func() {
		// Right of dot member completion list
		completionKind = CompletionKindPropertyAccess

		// Since this is qualified name check it's a type node location
		isImportType := ast.IsLiteralImportTypeNode(node)
		isTypeLocation := (isImportType && !node.AsImportTypeNode().IsTypeOf) ||
			ast.IsPartOfTypeNode(node.Parent) ||
			isPossiblyTypeArgumentPosition(contextToken, file, typeChecker)
		isRhsOfImportDeclaration := isInRightSideOfInternalImportEqualsDeclaration(node)
		if ast.IsEntityName(node) || isImportType || ast.IsPropertyAccessExpression(node) {
			isNamespaceName := ast.IsModuleDeclaration(node.Parent)
			if isNamespaceName {
				isNewIdentifierLocation = true
				defaultCommitCharacters = make([]string, 0)
			}
			symbol := typeChecker.GetSymbolAtLocation(node)
			if symbol != nil {
				symbol := checker.SkipAlias(symbol, typeChecker)
				if symbol.Flags&(ast.SymbolFlagsModule|ast.SymbolFlagsEnum) != 0 {
					var valueAccessNode *ast.Node
					if isImportType {
						valueAccessNode = node
					} else {
						valueAccessNode = node.Parent
					}
					// Extract module or enum members
					exportedSymbols := typeChecker.GetExportsOfModule(symbol)
					for _, exportedSymbol := range exportedSymbols {
						if exportedSymbol == nil {
							panic("getExporsOfModule() should all be defined")
						}
						isValidValueAccess := func(s *ast.Symbol) bool {
							return typeChecker.IsValidPropertyAccess(valueAccessNode, s.Name)
						}
						isValidTypeAccess := func(s *ast.Symbol) bool {
							return symbolCanBeReferencedAtTypeLocation(s, typeChecker, nil /*seenModules*/)
						}
						var isValidAccess bool
						if isNamespaceName {
							// At `namespace N.M/**/`, if this is the only declaration of `M`, don't include `M` as a completion.
							isValidAccess = exportedSymbol.Flags&ast.SymbolFlagsNamespace != 0 &&
								!core.Every(exportedSymbol.Declarations, func(declaration *ast.Declaration) bool {
									return declaration.Parent == node.Parent
								})
						} else if isRhsOfImportDeclaration {
							// Any kind is allowed when dotting off namespace in internal import equals declaration
							isValidAccess = isValidTypeAccess(exportedSymbol) || isValidValueAccess(exportedSymbol)
						} else if isTypeLocation || insideJsDocTagTypeExpression {
							isValidAccess = isValidTypeAccess(exportedSymbol)
						} else {
							isValidAccess = isValidValueAccess(exportedSymbol)
						}
						if isValidAccess {
							symbols = append(symbols, exportedSymbol)
						}
					}

					// If the module is merged with a value, we must get the type of the class and add its properties (for inherited static methods).
					if !isTypeLocation && !insideJsDocTagTypeExpression &&
						core.Some(
							symbol.Declarations,
							func(decl *ast.Declaration) bool {
								return decl.Kind != ast.KindSourceFile && decl.Kind != ast.KindModuleDeclaration && decl.Kind != ast.KindEnumDeclaration
							}) {
						t := typeChecker.GetNonOptionalType(typeChecker.GetTypeOfSymbolAtLocation(symbol, node))
						insertQuestionDot := false
						if typeChecker.IsNullableType(t) {
							canCorrectToQuestionDot := isRightOfDot && !isRightOfQuestionDot &&
								preferences.includeAutomaticOptionalChainCompletions
							if canCorrectToQuestionDot || isRightOfQuestionDot {
								t = typeChecker.GetNonNullableType(t)
								if canCorrectToQuestionDot {
									insertQuestionDot = true
								}
							}
						}
						addTypeProperties(t, node.Flags&ast.NodeFlagsAwaitContext != 0, insertQuestionDot)
					}
				}
			}
		}
	}

	if isRightOfDot || isRightOfQuestionDot {
		getTypeScriptMemberSymbols()
	} else if isRightOfOpenTag {
		// !!! jsx completions
	} else if isStartingCloseTag {
		// !!! jsx completions
	} else {
		// For JavaScript or TypeScript, if we're not after a dot, then just try to get the
		// global symbols in scope.  These results should be valid for either language as
		// the set of symbols that can be referenced from this location.
		// !!! global completions
	}

	// !!! adjustments

	return nil
}

// In a scenarion such as `const x = 1 * |`, the context and previous tokens are both `*`.
// In `const x = 1 * o|`, the context token is *, and the previous token is `o`.
// `contextToken` and `previousToken` can both be nil if we are at the beginning of the file.
func getRelevantTokens(position int, file *ast.SourceFile) (contextToken *ast.Node, previousToken *ast.Node) {
	previousToken = astnav.FindPrecedingToken(file, position)
	if previousToken != nil && position <= previousToken.End() && (ast.IsMemberName(previousToken) || ast.IsKeywordKind(previousToken.Kind)) {
		contextToken := astnav.FindPrecedingToken(file, previousToken.Pos())
		return contextToken, previousToken
	}
	return previousToken, previousToken
}

// "." | '"' | "'" | "`" | "/" | "@" | "<" | "#" | " "
type CompletionsTriggerCharacter = string

func isValidTrigger(file *ast.SourceFile, triggerCharacter CompletionsTriggerCharacter, contextToken *ast.Node, position int) bool {
	switch triggerCharacter {
	case ".", "@":
		return true
	case "\"", "'", "`":
		// Only automatically bring up completions if this is an opening quote.
		return contextToken != nil &&
			isStringLiteralOrTemplate(contextToken) &&
			position == getStartOfNode(contextToken, file)+1
	case "#":
		return contextToken != nil &&
			ast.IsPrivateIdentifier(contextToken) &&
			ast.GetContainingClass(contextToken) != nil
	case "<":
		// Opening JSX tag
		return contextToken != nil &&
			contextToken.Kind == ast.KindLessThanToken &&
			(!ast.IsBinaryExpression(contextToken.Parent) || binaryExpressionMayBeOpenTag(contextToken.Parent.AsBinaryExpression()))
	case "/":
		if contextToken == nil {
			return false
		}
		if ast.IsStringLiteralLike(contextToken) {
			return tryGetImportFromModuleSpecifier(contextToken) != nil
		}
		return contextToken.Kind == ast.KindSlashToken && ast.IsJsxClosingElement(contextToken.Parent)
	case " ":
		return contextToken != nil && contextToken.Kind == ast.KindImportKeyword && contextToken.Parent.Kind == ast.KindSourceFile
	default:
		panic("Unknown trigger character: " + triggerCharacter)
	}
}

func isStringLiteralOrTemplate(node *ast.Node) bool {
	switch node.Kind {
	case ast.KindStringLiteral, ast.KindNoSubstitutionTemplateLiteral, ast.KindTemplateExpression,
		ast.KindTaggedTemplateExpression:
		return true
	}
	return false
}

func binaryExpressionMayBeOpenTag(binaryExpression *ast.BinaryExpression) bool {
	return ast.NodeIsMissing(binaryExpression.Left)
}

func getDefaultCommitCharacters(isNewIdentifierLocation bool) *[]string {
	// !!!
	return nil
}

func isCheckedFile(file *ast.SourceFile, compilerOptions *core.CompilerOptions) bool {
	return !ast.IsSourceFileJs(file) || ast.IsCheckJsEnabledForFile(file, compilerOptions)
}

func isContextTokenValueLocation(contextToken *ast.Node) bool {
	return contextToken != nil && ((contextToken.Kind == ast.KindTypeOfKeyword &&
		(contextToken.Parent.Kind == ast.KindTypeQuery || ast.IsTypeOfExpression(contextToken.Parent))) ||
		(contextToken.Kind == ast.KindAssertsKeyword && contextToken.Parent.Kind == ast.KindTypePredicate))
}

func isPossiblyTypeArgumentPosition(token *ast.Node, sourceFile *ast.SourceFile, typeChecker *checker.Checker) bool {
	info := getPossibleTypeArgumentsInfo(token, sourceFile)
	return info != nil && (ast.IsPartOfTypeNode(info.called) ||
		len(getPossibleGenericSignatures(info.called, info.nTypeArguments, typeChecker)) != 0 ||
		isPossiblyTypeArgumentPosition(info.called, sourceFile, typeChecker))
}

func isContextTokenTypeLocation(contextToken *ast.Node) bool {
	if contextToken != nil {
		parentKind := contextToken.Parent.Kind
		switch contextToken.Kind {
		case ast.KindColonToken:
			return parentKind == ast.KindPropertyDeclaration ||
				parentKind == ast.KindPropertySignature ||
				parentKind == ast.KindParameter ||
				parentKind == ast.KindVariableDeclaration ||
				ast.IsFunctionLikeKind(parentKind)
		case ast.KindEqualsToken:
			return parentKind == ast.KindTypeAliasDeclaration || parentKind == ast.KindTypeParameter
		case ast.KindAsKeyword:
			return parentKind == ast.KindAsExpression
		case ast.KindLessThanToken:
			return parentKind == ast.KindTypeReference || parentKind == ast.KindTypeAssertionExpression
		case ast.KindExtendsKeyword:
			return parentKind == ast.KindTypeParameter
		case ast.KindSatisfiesKeyword:
			return parentKind == ast.KindSatisfiesExpression
		}
	}
	return false
}

// True if symbol is a type or a module containing at least one type.
func symbolCanBeReferencedAtTypeLocation(symbol *ast.Symbol, typeChecker *checker.Checker, seenModules map[ast.SymbolId]struct{}) bool {
	if seenModules == nil {
		seenModules = make(map[ast.SymbolId]struct{})
	}
	// Since an alias can be merged with a local declaration, we need to test both the alias and its target.
	// This code used to just test the result of `skipAlias`, but that would ignore any locally introduced meanings.
	return nonAliasCanBeReferencedAtTypeLocation(symbol, typeChecker, seenModules) ||
		nonAliasCanBeReferencedAtTypeLocation(
			checker.SkipAlias(core.IfElse(symbol.ExportSymbol != nil, symbol.ExportSymbol, symbol), typeChecker),
			typeChecker,
			seenModules,
		)
}

func nonAliasCanBeReferencedAtTypeLocation(symbol *ast.Symbol, typeChecker *checker.Checker, seenModules map[ast.SymbolId]struct{}) bool {
	return symbol.Flags&ast.SymbolFlagsType != 0 || typeChecker.IsUnknownSymbol(symbol) ||
		symbol.Flags&ast.SymbolFlagsModule != 0 && core.AddToSeen(seenModules, ast.GetSymbolId(symbol), struct{}{}) &&
			core.Some(
				typeChecker.GetExportsOfModule(symbol),
				func(e *ast.Symbol) bool { return symbolCanBeReferencedAtTypeLocation(e, typeChecker, seenModules) })
}

// Gets all properties on a type, but if that type is a union of several types,
// excludes array-like types or callable/constructable types.
func getPropertiesForCompletion(t *checker.Type, typeChecker *checker.Checker) []*ast.Symbol {
	if typeChecker.IsUnion(t) {
		return core.CheckEachDefined(typeChecker.GetAllPossiblePropertiesOfTypes(t.Types()), "getAllPossiblePropertiesOfTypes() should all be defined.")
	} else {
		return core.CheckEachDefined(typeChecker.GetApparentProperties(t), "getApparentProperties() should all be defined.")
	}
}

// Given 'a.b.c', returns 'a'.
func getLeftMostName(e *ast.Expression) *ast.IdentifierNode {
	if ast.IsIdentifier(e) {
		return e
	} else if ast.IsPropertyAccessExpression(e) {
		return getLeftMostName(e.Expression())
	} else {
		return nil
	}
}

func getFirstSymbolInChain(symbol *ast.Symbol, enclosingDeclaration *ast.Node, typeChecker *checker.Checker) *ast.Symbol {
	chain := typeChecker.GetAccessibleSymbolChain(
		symbol,
		enclosingDeclaration,
		ast.SymbolFlagsAll, /*meaning*/
		false /*useOnlyExternalAliasing*/)
	if len(chain) > 0 {
		return chain[0]
	}
	if symbol.Parent != nil {
		if isModuleSymbol(symbol.Parent) {
			return symbol
		}
		return getFirstSymbolInChain(symbol.Parent, enclosingDeclaration, typeChecker)
	}
	return nil
}

func isModuleSymbol(symbol *ast.Symbol) bool {
	return core.Some(symbol.Declarations, func(decl *ast.Declaration) bool { return decl.Kind == ast.KindSourceFile })
}

func getNullableSymbolOriginInfoKind(kind symbolOriginInfoKind, insertQuestionDot bool) symbolOriginInfoKind {
	if insertQuestionDot {
		kind |= symbolOriginInfoKindNullable
	}
	return kind
}

func isStaticProperty(symbol *ast.Symbol) bool {
	return symbol.ValueDeclaration != nil &&
		checker.GetEffectiveModifierFlags(symbol.ValueDeclaration)&ast.ModifierFlagsStatic != 0 &&
		ast.IsClassLike(symbol.ValueDeclaration.Parent)
}
