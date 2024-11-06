package ast

import "fmt"

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
	SyntaxKindQualifiedName
	SyntaxKindComputedPropertyName
	SyntaxKindModifierList
	SyntaxKindTypeParameterList
	SyntaxKindTypeArgumentList
	SyntaxKindTypeParameter
	SyntaxKindParameter
	SyntaxKindDecorator
	SyntaxKindPropertySignature
	SyntaxKindPropertyDeclaration
	SyntaxKindMethodSignature
	SyntaxKindMethodDeclaration
	SyntaxKindClassStaticBlockDeclaration
	SyntaxKindConstructor
	SyntaxKindGetAccessor
	SyntaxKindSetAccessor
	SyntaxKindCallSignature
	SyntaxKindConstructSignature
	SyntaxKindIndexSignature
	SyntaxKindTypePredicate
	SyntaxKindTypeReference
	SyntaxKindFunctionType
	SyntaxKindConstructorType
	SyntaxKindTypeQuery
	SyntaxKindTypeLiteral
	SyntaxKindArrayType
	SyntaxKindTupleType
	SyntaxKindOptionalType
	SyntaxKindRestType
	SyntaxKindUnionType
	SyntaxKindIntersectionType
	SyntaxKindConditionalType
	SyntaxKindInferType
	SyntaxKindParenthesizedType
	SyntaxKindThisType
	SyntaxKindTypeOperator
	SyntaxKindIndexedAccessType
	SyntaxKindMappedType
	SyntaxKindLiteralType
	SyntaxKindNamedTupleMember
	SyntaxKindTemplateLiteralType
	SyntaxKindTemplateLiteralTypeSpan
	SyntaxKindImportType
	SyntaxKindObjectBindingPattern
	SyntaxKindArrayBindingPattern
	SyntaxKindBindingElement
	SyntaxKindArrayLiteralExpression
	SyntaxKindObjectLiteralExpression
	SyntaxKindPropertyAccessExpression
	SyntaxKindElementAccessExpression
	SyntaxKindCallExpression
	SyntaxKindNewExpression
	SyntaxKindTaggedTemplateExpression
	SyntaxKindTypeAssertionExpression
	SyntaxKindParenthesizedExpression
	SyntaxKindFunctionExpression
	SyntaxKindArrowFunction
	SyntaxKindDeleteExpression
	SyntaxKindTypeOfExpression
	SyntaxKindVoidExpression
	SyntaxKindAwaitExpression
	SyntaxKindPrefixUnaryExpression
	SyntaxKindPostfixUnaryExpression
	SyntaxKindBinaryExpression
	SyntaxKindConditionalExpression
	SyntaxKindTemplateExpression
	SyntaxKindYieldExpression
	SyntaxKindSpreadElement
	SyntaxKindClassExpression
	SyntaxKindOmittedExpression
	SyntaxKindExpressionWithTypeArguments
	SyntaxKindAsExpression
	SyntaxKindNonNullExpression
	SyntaxKindMetaProperty
	SyntaxKindSyntheticExpression
	SyntaxKindSatisfiesExpression
	SyntaxKindTemplateSpan
	SyntaxKindSemicolonClassElement
	SyntaxKindBlock
	SyntaxKindEmptyStatement
	SyntaxKindVariableStatement
	SyntaxKindExpressionStatement
	SyntaxKindIfStatement
	SyntaxKindDoStatement
	SyntaxKindWhileStatement
	SyntaxKindForStatement
	SyntaxKindForInStatement
	SyntaxKindForOfStatement
	SyntaxKindContinueStatement
	SyntaxKindBreakStatement
	SyntaxKindReturnStatement
	SyntaxKindWithStatement
	SyntaxKindSwitchStatement
	SyntaxKindLabeledStatement
	SyntaxKindThrowStatement
	SyntaxKindTryStatement
	SyntaxKindDebuggerStatement
	SyntaxKindVariableDeclaration
	SyntaxKindVariableDeclarationList
	SyntaxKindFunctionDeclaration
	SyntaxKindClassDeclaration
	SyntaxKindInterfaceDeclaration
	SyntaxKindTypeAliasDeclaration
	SyntaxKindEnumDeclaration
	SyntaxKindModuleDeclaration
	SyntaxKindModuleBlock
	SyntaxKindCaseBlock
	SyntaxKindNamespaceExportDeclaration
	SyntaxKindImportEqualsDeclaration
	SyntaxKindImportDeclaration
	SyntaxKindImportClause
	SyntaxKindNamespaceImport
	SyntaxKindNamedImports
	SyntaxKindImportSpecifier
	SyntaxKindExportAssignment
	SyntaxKindExportDeclaration
	SyntaxKindNamedExports
	SyntaxKindNamespaceExport
	SyntaxKindExportSpecifier
	SyntaxKindMissingDeclaration
	SyntaxKindExternalModuleReference
	SyntaxKindJsxElement
	SyntaxKindJsxSelfClosingElement
	SyntaxKindJsxOpeningElement
	SyntaxKindJsxClosingElement
	SyntaxKindJsxFragment
	SyntaxKindJsxOpeningFragment
	SyntaxKindJsxClosingFragment
	SyntaxKindJsxAttribute
	SyntaxKindJsxAttributes
	SyntaxKindJsxSpreadAttribute
	SyntaxKindJsxExpression
	SyntaxKindJsxNamespacedName
	SyntaxKindCaseClause
	SyntaxKindDefaultClause
	SyntaxKindHeritageClause
	SyntaxKindCatchClause
	SyntaxKindImportAttributes
	SyntaxKindImportAttribute
	SyntaxKindPropertyAssignment
	SyntaxKindShorthandPropertyAssignment
	SyntaxKindSpreadAssignment
	SyntaxKindEnumMember
	SyntaxKindSourceFile
	SyntaxKindBundle
	SyntaxKindJSDocTypeExpression
	SyntaxKindJSDocNameReference
	SyntaxKindJSDocMemberName
	SyntaxKindJSDocAllType
	SyntaxKindJSDocUnknownType
	SyntaxKindJSDocNullableType
	SyntaxKindJSDocNonNullableType
	SyntaxKindJSDocOptionalType
	SyntaxKindJSDocFunctionType
	SyntaxKindJSDocVariadicType
	SyntaxKindJSDocNamepathType
	SyntaxKindJSDoc
	SyntaxKindJSDocText
	SyntaxKindJSDocTypeLiteral
	SyntaxKindJSDocSignature
	SyntaxKindJSDocLink
	SyntaxKindJSDocLinkCode
	SyntaxKindJSDocLinkPlain
	SyntaxKindJSDocTag
	SyntaxKindJSDocAugmentsTag
	SyntaxKindJSDocImplementsTag
	SyntaxKindJSDocAuthorTag
	SyntaxKindJSDocDeprecatedTag
	SyntaxKindJSDocImmediateTag
	SyntaxKindJSDocClassTag
	SyntaxKindJSDocPublicTag
	SyntaxKindJSDocPrivateTag
	SyntaxKindJSDocProtectedTag
	SyntaxKindJSDocReadonlyTag
	SyntaxKindJSDocOverrideTag
	SyntaxKindJSDocCallbackTag
	SyntaxKindJSDocOverloadTag
	SyntaxKindJSDocEnumTag
	SyntaxKindJSDocParameterTag
	SyntaxKindJSDocReturnTag
	SyntaxKindJSDocThisTag
	SyntaxKindJSDocTypeTag
	SyntaxKindJSDocTemplateTag
	SyntaxKindJSDocTypedefTag
	SyntaxKindJSDocSeeTag
	SyntaxKindJSDocPropertyTag
	SyntaxKindJSDocThrowsTag
	SyntaxKindJSDocSatisfiesTag
	SyntaxKindJSDocImportTag
	SyntaxKindSyntaxList
	SyntaxKindNotEmittedStatement
	SyntaxKindPartiallyEmittedExpression
	SyntaxKindCommaListExpression
	SyntaxKindSyntheticReferenceExpression

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
	TypeNodeStart           = SyntaxKindTypePredicate
	TypeNodeEnd             = SyntaxKindImportType
	PunctuationStart        = SyntaxKindOpenBraceToken
	PunctuationEnd          = SyntaxKindCaretEqualsToken
	TokenStart              = SyntaxKindUnknown
	TokenEnd                = SyntaxKindOfKeyword
	LiteralTokenStart       = SyntaxKindNumericLiteral
	LiteralTokenEnd         = SyntaxKindNoSubstitutionTemplateLiteral
	TemplateTokenStart      = SyntaxKindNoSubstitutionTemplateLiteral
	TemplateTokenEnd        = SyntaxKindTemplateTail
	BinaryOperatorStart     = SyntaxKindLessThanToken
	BinaryOperatorEnd       = SyntaxKindCaretEqualsToken
	NodeStart               = SyntaxKindQualifiedName
	NodeEnd                 = SyntaxKindSyntheticReferenceExpression
	JSDocNodeStart          = SyntaxKindJSDocTypeExpression
	JSDocNodeEnd            = SyntaxKindJSDocImportTag
	JSDocTagNodeStart       = SyntaxKindJSDocTag
	JSDocTagNodeEnd         = SyntaxKindJSDocImportTag
	ContextualKeywordStart  = SyntaxKindAbstractKeyword
	ContextualKeywordEnd    = SyntaxKindOfKeyword
)

var syntaxKindNames = [SyntaxKindCount + 1]string{
	SyntaxKindUnknown:                                      "Unknown",
	SyntaxKindEndOfFile:                                    "EndOfFile",
	SyntaxKindConflictMarkerTrivia:                         "ConflictMarkerTrivia",
	SyntaxKindNonTextFileMarkerTrivia:                      "NonTextFileMarkerTrivia",
	SyntaxKindNumericLiteral:                               "NumericLiteral",
	SyntaxKindBigIntLiteral:                                "BigIntLiteral",
	SyntaxKindStringLiteral:                                "StringLiteral",
	SyntaxKindJsxText:                                      "JsxText",
	SyntaxKindJsxTextAllWhiteSpaces:                        "JsxTextAllWhiteSpaces",
	SyntaxKindRegularExpressionLiteral:                     "RegularExpressionLiteral",
	SyntaxKindNoSubstitutionTemplateLiteral:                "NoSubstitutionTemplateLiteral",
	SyntaxKindTemplateHead:                                 "TemplateHead",
	SyntaxKindTemplateMiddle:                               "TemplateMiddle",
	SyntaxKindTemplateTail:                                 "TemplateTail",
	SyntaxKindOpenBraceToken:                               "OpenBraceToken",
	SyntaxKindCloseBraceToken:                              "CloseBraceToken",
	SyntaxKindOpenParenToken:                               "OpenParenToken",
	SyntaxKindCloseParenToken:                              "CloseParenToken",
	SyntaxKindOpenBracketToken:                             "OpenBracketToken",
	SyntaxKindCloseBracketToken:                            "CloseBracketToken",
	SyntaxKindDotToken:                                     "DotToken",
	SyntaxKindDotDotDotToken:                               "DotDotDotToken",
	SyntaxKindSemicolonToken:                               "SemicolonToken",
	SyntaxKindCommaToken:                                   "CommaToken",
	SyntaxKindQuestionDotToken:                             "QuestionDotToken",
	SyntaxKindLessThanToken:                                "LessThanToken",
	SyntaxKindLessThanSlashToken:                           "LessThanSlashToken",
	SyntaxKindGreaterThanToken:                             "GreaterThanToken",
	SyntaxKindLessThanEqualsToken:                          "LessThanEqualsToken",
	SyntaxKindGreaterThanEqualsToken:                       "GreaterThanEqualsToken",
	SyntaxKindEqualsEqualsToken:                            "EqualsEqualsToken",
	SyntaxKindExclamationEqualsToken:                       "ExclamationEqualsToken",
	SyntaxKindEqualsEqualsEqualsToken:                      "EqualsEqualsEqualsToken",
	SyntaxKindExclamationEqualsEqualsToken:                 "ExclamationEqualsEqualsToken",
	SyntaxKindEqualsGreaterThanToken:                       "EqualsGreaterThanToken",
	SyntaxKindPlusToken:                                    "PlusToken",
	SyntaxKindMinusToken:                                   "MinusToken",
	SyntaxKindAsteriskToken:                                "AsteriskToken",
	SyntaxKindAsteriskAsteriskToken:                        "AsteriskAsteriskToken",
	SyntaxKindSlashToken:                                   "SlashToken",
	SyntaxKindPercentToken:                                 "PercentToken",
	SyntaxKindPlusPlusToken:                                "PlusPlusToken",
	SyntaxKindMinusMinusToken:                              "MinusMinusToken",
	SyntaxKindLessThanLessThanToken:                        "LessThanLessThanToken",
	SyntaxKindGreaterThanGreaterThanToken:                  "GreaterThanGreaterThanToken",
	SyntaxKindGreaterThanGreaterThanGreaterThanToken:       "GreaterThanGreaterThanGreaterThanToken",
	SyntaxKindAmpersandToken:                               "AmpersandToken",
	SyntaxKindBarToken:                                     "BarToken",
	SyntaxKindCaretToken:                                   "CaretToken",
	SyntaxKindExclamationToken:                             "ExclamationToken",
	SyntaxKindTildeToken:                                   "TildeToken",
	SyntaxKindAmpersandAmpersandToken:                      "AmpersandAmpersandToken",
	SyntaxKindBarBarToken:                                  "BarBarToken",
	SyntaxKindQuestionToken:                                "QuestionToken",
	SyntaxKindColonToken:                                   "ColonToken",
	SyntaxKindAtToken:                                      "AtToken",
	SyntaxKindQuestionQuestionToken:                        "QuestionQuestionToken",
	SyntaxKindBacktickToken:                                "BacktickToken",
	SyntaxKindHashToken:                                    "HashToken",
	SyntaxKindEqualsToken:                                  "EqualsToken",
	SyntaxKindPlusEqualsToken:                              "PlusEqualsToken",
	SyntaxKindMinusEqualsToken:                             "MinusEqualsToken",
	SyntaxKindAsteriskEqualsToken:                          "AsteriskEqualsToken",
	SyntaxKindAsteriskAsteriskEqualsToken:                  "AsteriskAsteriskEqualsToken",
	SyntaxKindSlashEqualsToken:                             "SlashEqualsToken",
	SyntaxKindPercentEqualsToken:                           "PercentEqualsToken",
	SyntaxKindLessThanLessThanEqualsToken:                  "LessThanLessThanEqualsToken",
	SyntaxKindGreaterThanGreaterThanEqualsToken:            "GreaterThanGreaterThanEqualsToken",
	SyntaxKindGreaterThanGreaterThanGreaterThanEqualsToken: "GreaterThanGreaterThanGreaterThanEqualsToken",
	SyntaxKindAmpersandEqualsToken:                         "AmpersandEqualsToken",
	SyntaxKindBarEqualsToken:                               "BarEqualsToken",
	SyntaxKindBarBarEqualsToken:                            "BarBarEqualsToken",
	SyntaxKindAmpersandAmpersandEqualsToken:                "AmpersandAmpersandEqualsToken",
	SyntaxKindQuestionQuestionEqualsToken:                  "QuestionQuestionEqualsToken",
	SyntaxKindCaretEqualsToken:                             "CaretEqualsToken",
	SyntaxKindIdentifier:                                   "Identifier",
	SyntaxKindPrivateIdentifier:                            "PrivateIdentifier",
	SyntaxKindJSDocCommentTextToken:                        "JSDocCommentTextToken",
	SyntaxKindBreakKeyword:                                 "BreakKeyword",
	SyntaxKindCaseKeyword:                                  "CaseKeyword",
	SyntaxKindCatchKeyword:                                 "CatchKeyword",
	SyntaxKindClassKeyword:                                 "ClassKeyword",
	SyntaxKindConstKeyword:                                 "ConstKeyword",
	SyntaxKindContinueKeyword:                              "ContinueKeyword",
	SyntaxKindDebuggerKeyword:                              "DebuggerKeyword",
	SyntaxKindDefaultKeyword:                               "DefaultKeyword",
	SyntaxKindDeleteKeyword:                                "DeleteKeyword",
	SyntaxKindDoKeyword:                                    "DoKeyword",
	SyntaxKindElseKeyword:                                  "ElseKeyword",
	SyntaxKindEnumKeyword:                                  "EnumKeyword",
	SyntaxKindExportKeyword:                                "ExportKeyword",
	SyntaxKindExtendsKeyword:                               "ExtendsKeyword",
	SyntaxKindFalseKeyword:                                 "FalseKeyword",
	SyntaxKindFinallyKeyword:                               "FinallyKeyword",
	SyntaxKindForKeyword:                                   "ForKeyword",
	SyntaxKindFunctionKeyword:                              "FunctionKeyword",
	SyntaxKindIfKeyword:                                    "IfKeyword",
	SyntaxKindImportKeyword:                                "ImportKeyword",
	SyntaxKindInKeyword:                                    "InKeyword",
	SyntaxKindInstanceOfKeyword:                            "InstanceOfKeyword",
	SyntaxKindNewKeyword:                                   "NewKeyword",
	SyntaxKindNullKeyword:                                  "NullKeyword",
	SyntaxKindReturnKeyword:                                "ReturnKeyword",
	SyntaxKindSuperKeyword:                                 "SuperKeyword",
	SyntaxKindSwitchKeyword:                                "SwitchKeyword",
	SyntaxKindThisKeyword:                                  "ThisKeyword",
	SyntaxKindThrowKeyword:                                 "ThrowKeyword",
	SyntaxKindTrueKeyword:                                  "TrueKeyword",
	SyntaxKindTryKeyword:                                   "TryKeyword",
	SyntaxKindTypeOfKeyword:                                "TypeOfKeyword",
	SyntaxKindVarKeyword:                                   "VarKeyword",
	SyntaxKindVoidKeyword:                                  "VoidKeyword",
	SyntaxKindWhileKeyword:                                 "WhileKeyword",
	SyntaxKindWithKeyword:                                  "WithKeyword",
	SyntaxKindImplementsKeyword:                            "ImplementsKeyword",
	SyntaxKindInterfaceKeyword:                             "InterfaceKeyword",
	SyntaxKindLetKeyword:                                   "LetKeyword",
	SyntaxKindPackageKeyword:                               "PackageKeyword",
	SyntaxKindPrivateKeyword:                               "PrivateKeyword",
	SyntaxKindProtectedKeyword:                             "ProtectedKeyword",
	SyntaxKindPublicKeyword:                                "PublicKeyword",
	SyntaxKindStaticKeyword:                                "StaticKeyword",
	SyntaxKindYieldKeyword:                                 "YieldKeyword",
	SyntaxKindAbstractKeyword:                              "AbstractKeyword",
	SyntaxKindAccessorKeyword:                              "AccessorKeyword",
	SyntaxKindAsKeyword:                                    "AsKeyword",
	SyntaxKindAssertsKeyword:                               "AssertsKeyword",
	SyntaxKindAssertKeyword:                                "AssertKeyword",
	SyntaxKindAnyKeyword:                                   "AnyKeyword",
	SyntaxKindAsyncKeyword:                                 "AsyncKeyword",
	SyntaxKindAwaitKeyword:                                 "AwaitKeyword",
	SyntaxKindBooleanKeyword:                               "BooleanKeyword",
	SyntaxKindConstructorKeyword:                           "ConstructorKeyword",
	SyntaxKindDeclareKeyword:                               "DeclareKeyword",
	SyntaxKindGetKeyword:                                   "GetKeyword",
	SyntaxKindImmediateKeyword:                             "ImmediateKeyword",
	SyntaxKindInferKeyword:                                 "InferKeyword",
	SyntaxKindIntrinsicKeyword:                             "IntrinsicKeyword",
	SyntaxKindIsKeyword:                                    "IsKeyword",
	SyntaxKindKeyOfKeyword:                                 "KeyOfKeyword",
	SyntaxKindModuleKeyword:                                "ModuleKeyword",
	SyntaxKindNamespaceKeyword:                             "NamespaceKeyword",
	SyntaxKindNeverKeyword:                                 "NeverKeyword",
	SyntaxKindOutKeyword:                                   "OutKeyword",
	SyntaxKindReadonlyKeyword:                              "ReadonlyKeyword",
	SyntaxKindRequireKeyword:                               "RequireKeyword",
	SyntaxKindNumberKeyword:                                "NumberKeyword",
	SyntaxKindObjectKeyword:                                "ObjectKeyword",
	SyntaxKindSatisfiesKeyword:                             "SatisfiesKeyword",
	SyntaxKindSetKeyword:                                   "SetKeyword",
	SyntaxKindStringKeyword:                                "StringKeyword",
	SyntaxKindSymbolKeyword:                                "SymbolKeyword",
	SyntaxKindTypeKeyword:                                  "TypeKeyword",
	SyntaxKindUndefinedKeyword:                             "UndefinedKeyword",
	SyntaxKindUniqueKeyword:                                "UniqueKeyword",
	SyntaxKindUnknownKeyword:                               "UnknownKeyword",
	SyntaxKindUsingKeyword:                                 "UsingKeyword",
	SyntaxKindFromKeyword:                                  "FromKeyword",
	SyntaxKindGlobalKeyword:                                "GlobalKeyword",
	SyntaxKindBigIntKeyword:                                "BigIntKeyword",
	SyntaxKindOverrideKeyword:                              "OverrideKeyword",
	SyntaxKindOfKeyword:                                    "OfKeyword",
	SyntaxKindQualifiedName:                                "QualifiedName",
	SyntaxKindComputedPropertyName:                         "ComputedPropertyName",
	SyntaxKindModifierList:                                 "ModifierList",
	SyntaxKindTypeParameterList:                            "TypeParameterList",
	SyntaxKindTypeArgumentList:                             "TypeArgumentList",
	SyntaxKindTypeParameter:                                "TypeParameter",
	SyntaxKindParameter:                                    "Parameter",
	SyntaxKindDecorator:                                    "Decorator",
	SyntaxKindPropertySignature:                            "PropertySignature",
	SyntaxKindPropertyDeclaration:                          "PropertyDeclaration",
	SyntaxKindMethodSignature:                              "MethodSignature",
	SyntaxKindMethodDeclaration:                            "MethodDeclaration",
	SyntaxKindClassStaticBlockDeclaration:                  "ClassStaticBlockDeclaration",
	SyntaxKindConstructor:                                  "Constructor",
	SyntaxKindGetAccessor:                                  "GetAccessor",
	SyntaxKindSetAccessor:                                  "SetAccessor",
	SyntaxKindCallSignature:                                "CallSignature",
	SyntaxKindConstructSignature:                           "ConstructSignature",
	SyntaxKindIndexSignature:                               "IndexSignature",
	SyntaxKindTypePredicate:                                "TypePredicate",
	SyntaxKindTypeReference:                                "TypeReference",
	SyntaxKindFunctionType:                                 "FunctionType",
	SyntaxKindConstructorType:                              "ConstructorType",
	SyntaxKindTypeQuery:                                    "TypeQuery",
	SyntaxKindTypeLiteral:                                  "TypeLiteral",
	SyntaxKindArrayType:                                    "ArrayType",
	SyntaxKindTupleType:                                    "TupleType",
	SyntaxKindOptionalType:                                 "OptionalType",
	SyntaxKindRestType:                                     "RestType",
	SyntaxKindUnionType:                                    "UnionType",
	SyntaxKindIntersectionType:                             "IntersectionType",
	SyntaxKindConditionalType:                              "ConditionalType",
	SyntaxKindInferType:                                    "InferType",
	SyntaxKindParenthesizedType:                            "ParenthesizedType",
	SyntaxKindThisType:                                     "ThisType",
	SyntaxKindTypeOperator:                                 "TypeOperator",
	SyntaxKindIndexedAccessType:                            "IndexedAccessType",
	SyntaxKindMappedType:                                   "MappedType",
	SyntaxKindLiteralType:                                  "LiteralType",
	SyntaxKindNamedTupleMember:                             "NamedTupleMember",
	SyntaxKindTemplateLiteralType:                          "TemplateLiteralType",
	SyntaxKindTemplateLiteralTypeSpan:                      "TemplateLiteralTypeSpan",
	SyntaxKindImportType:                                   "ImportType",
	SyntaxKindObjectBindingPattern:                         "ObjectBindingPattern",
	SyntaxKindArrayBindingPattern:                          "ArrayBindingPattern",
	SyntaxKindBindingElement:                               "BindingElement",
	SyntaxKindArrayLiteralExpression:                       "ArrayLiteralExpression",
	SyntaxKindObjectLiteralExpression:                      "ObjectLiteralExpression",
	SyntaxKindPropertyAccessExpression:                     "PropertyAccessExpression",
	SyntaxKindElementAccessExpression:                      "ElementAccessExpression",
	SyntaxKindCallExpression:                               "CallExpression",
	SyntaxKindNewExpression:                                "NewExpression",
	SyntaxKindTaggedTemplateExpression:                     "TaggedTemplateExpression",
	SyntaxKindTypeAssertionExpression:                      "TypeAssertionExpression",
	SyntaxKindParenthesizedExpression:                      "ParenthesizedExpression",
	SyntaxKindFunctionExpression:                           "FunctionExpression",
	SyntaxKindArrowFunction:                                "ArrowFunction",
	SyntaxKindDeleteExpression:                             "DeleteExpression",
	SyntaxKindTypeOfExpression:                             "TypeOfExpression",
	SyntaxKindVoidExpression:                               "VoidExpression",
	SyntaxKindAwaitExpression:                              "AwaitExpression",
	SyntaxKindPrefixUnaryExpression:                        "PrefixUnaryExpression",
	SyntaxKindPostfixUnaryExpression:                       "PostfixUnaryExpression",
	SyntaxKindBinaryExpression:                             "BinaryExpression",
	SyntaxKindConditionalExpression:                        "ConditionalExpression",
	SyntaxKindTemplateExpression:                           "TemplateExpression",
	SyntaxKindYieldExpression:                              "YieldExpression",
	SyntaxKindSpreadElement:                                "SpreadElement",
	SyntaxKindClassExpression:                              "ClassExpression",
	SyntaxKindOmittedExpression:                            "OmittedExpression",
	SyntaxKindExpressionWithTypeArguments:                  "ExpressionWithTypeArguments",
	SyntaxKindAsExpression:                                 "AsExpression",
	SyntaxKindNonNullExpression:                            "NonNullExpression",
	SyntaxKindMetaProperty:                                 "MetaProperty",
	SyntaxKindSyntheticExpression:                          "SyntheticExpression",
	SyntaxKindSatisfiesExpression:                          "SatisfiesExpression",
	SyntaxKindTemplateSpan:                                 "TemplateSpan",
	SyntaxKindSemicolonClassElement:                        "SemicolonClassElement",
	SyntaxKindBlock:                                        "Block",
	SyntaxKindEmptyStatement:                               "EmptyStatement",
	SyntaxKindVariableStatement:                            "VariableStatement",
	SyntaxKindExpressionStatement:                          "ExpressionStatement",
	SyntaxKindIfStatement:                                  "IfStatement",
	SyntaxKindDoStatement:                                  "DoStatement",
	SyntaxKindWhileStatement:                               "WhileStatement",
	SyntaxKindForStatement:                                 "ForStatement",
	SyntaxKindForInStatement:                               "ForInStatement",
	SyntaxKindForOfStatement:                               "ForOfStatement",
	SyntaxKindContinueStatement:                            "ContinueStatement",
	SyntaxKindBreakStatement:                               "BreakStatement",
	SyntaxKindReturnStatement:                              "ReturnStatement",
	SyntaxKindWithStatement:                                "WithStatement",
	SyntaxKindSwitchStatement:                              "SwitchStatement",
	SyntaxKindLabeledStatement:                             "LabeledStatement",
	SyntaxKindThrowStatement:                               "ThrowStatement",
	SyntaxKindTryStatement:                                 "TryStatement",
	SyntaxKindDebuggerStatement:                            "DebuggerStatement",
	SyntaxKindVariableDeclaration:                          "VariableDeclaration",
	SyntaxKindVariableDeclarationList:                      "VariableDeclarationList",
	SyntaxKindFunctionDeclaration:                          "FunctionDeclaration",
	SyntaxKindClassDeclaration:                             "ClassDeclaration",
	SyntaxKindInterfaceDeclaration:                         "InterfaceDeclaration",
	SyntaxKindTypeAliasDeclaration:                         "TypeAliasDeclaration",
	SyntaxKindEnumDeclaration:                              "EnumDeclaration",
	SyntaxKindModuleDeclaration:                            "ModuleDeclaration",
	SyntaxKindModuleBlock:                                  "ModuleBlock",
	SyntaxKindCaseBlock:                                    "CaseBlock",
	SyntaxKindNamespaceExportDeclaration:                   "NamespaceExportDeclaration",
	SyntaxKindImportEqualsDeclaration:                      "ImportEqualsDeclaration",
	SyntaxKindImportDeclaration:                            "ImportDeclaration",
	SyntaxKindImportClause:                                 "ImportClause",
	SyntaxKindNamespaceImport:                              "NamespaceImport",
	SyntaxKindNamedImports:                                 "NamedImports",
	SyntaxKindImportSpecifier:                              "ImportSpecifier",
	SyntaxKindExportAssignment:                             "ExportAssignment",
	SyntaxKindExportDeclaration:                            "ExportDeclaration",
	SyntaxKindNamedExports:                                 "NamedExports",
	SyntaxKindNamespaceExport:                              "NamespaceExport",
	SyntaxKindExportSpecifier:                              "ExportSpecifier",
	SyntaxKindMissingDeclaration:                           "MissingDeclaration",
	SyntaxKindExternalModuleReference:                      "ExternalModuleReference",
	SyntaxKindJsxElement:                                   "JsxElement",
	SyntaxKindJsxSelfClosingElement:                        "JsxSelfClosingElement",
	SyntaxKindJsxOpeningElement:                            "JsxOpeningElement",
	SyntaxKindJsxClosingElement:                            "JsxClosingElement",
	SyntaxKindJsxFragment:                                  "JsxFragment",
	SyntaxKindJsxOpeningFragment:                           "JsxOpeningFragment",
	SyntaxKindJsxClosingFragment:                           "JsxClosingFragment",
	SyntaxKindJsxAttribute:                                 "JsxAttribute",
	SyntaxKindJsxAttributes:                                "JsxAttributes",
	SyntaxKindJsxSpreadAttribute:                           "JsxSpreadAttribute",
	SyntaxKindJsxExpression:                                "JsxExpression",
	SyntaxKindJsxNamespacedName:                            "JsxNamespacedName",
	SyntaxKindCaseClause:                                   "CaseClause",
	SyntaxKindDefaultClause:                                "DefaultClause",
	SyntaxKindHeritageClause:                               "HeritageClause",
	SyntaxKindCatchClause:                                  "CatchClause",
	SyntaxKindImportAttributes:                             "ImportAttributes",
	SyntaxKindImportAttribute:                              "ImportAttribute",
	SyntaxKindPropertyAssignment:                           "PropertyAssignment",
	SyntaxKindShorthandPropertyAssignment:                  "ShorthandPropertyAssignment",
	SyntaxKindSpreadAssignment:                             "SpreadAssignment",
	SyntaxKindEnumMember:                                   "EnumMember",
	SyntaxKindSourceFile:                                   "SourceFile",
	SyntaxKindBundle:                                       "Bundle",
	SyntaxKindJSDocTypeExpression:                          "JSDocTypeExpression",
	SyntaxKindJSDocNameReference:                           "JSDocNameReference",
	SyntaxKindJSDocMemberName:                              "JSDocMemberName",
	SyntaxKindJSDocAllType:                                 "JSDocAllType",
	SyntaxKindJSDocUnknownType:                             "JSDocUnknownType",
	SyntaxKindJSDocNullableType:                            "JSDocNullableType",
	SyntaxKindJSDocNonNullableType:                         "JSDocNonNullableType",
	SyntaxKindJSDocOptionalType:                            "JSDocOptionalType",
	SyntaxKindJSDocFunctionType:                            "JSDocFunctionType",
	SyntaxKindJSDocVariadicType:                            "JSDocVariadicType",
	SyntaxKindJSDocNamepathType:                            "JSDocNamepathType",
	SyntaxKindJSDoc:                                        "JSDoc",
	SyntaxKindJSDocText:                                    "JSDocText",
	SyntaxKindJSDocTypeLiteral:                             "JSDocTypeLiteral",
	SyntaxKindJSDocSignature:                               "JSDocSignature",
	SyntaxKindJSDocLink:                                    "JSDocLink",
	SyntaxKindJSDocLinkCode:                                "JSDocLinkCode",
	SyntaxKindJSDocLinkPlain:                               "JSDocLinkPlain",
	SyntaxKindJSDocTag:                                     "JSDocTag",
	SyntaxKindJSDocAugmentsTag:                             "JSDocAugmentsTag",
	SyntaxKindJSDocImplementsTag:                           "JSDocImplementsTag",
	SyntaxKindJSDocAuthorTag:                               "JSDocAuthorTag",
	SyntaxKindJSDocDeprecatedTag:                           "JSDocDeprecatedTag",
	SyntaxKindJSDocImmediateTag:                            "JSDocImmediateTag",
	SyntaxKindJSDocClassTag:                                "JSDocClassTag",
	SyntaxKindJSDocPublicTag:                               "JSDocPublicTag",
	SyntaxKindJSDocPrivateTag:                              "JSDocPrivateTag",
	SyntaxKindJSDocProtectedTag:                            "JSDocProtectedTag",
	SyntaxKindJSDocReadonlyTag:                             "JSDocReadonlyTag",
	SyntaxKindJSDocOverrideTag:                             "JSDocOverrideTag",
	SyntaxKindJSDocCallbackTag:                             "JSDocCallbackTag",
	SyntaxKindJSDocOverloadTag:                             "JSDocOverloadTag",
	SyntaxKindJSDocEnumTag:                                 "JSDocEnumTag",
	SyntaxKindJSDocParameterTag:                            "JSDocParameterTag",
	SyntaxKindJSDocReturnTag:                               "JSDocReturnTag",
	SyntaxKindJSDocThisTag:                                 "JSDocThisTag",
	SyntaxKindJSDocTypeTag:                                 "JSDocTypeTag",
	SyntaxKindJSDocTemplateTag:                             "JSDocTemplateTag",
	SyntaxKindJSDocTypedefTag:                              "JSDocTypedefTag",
	SyntaxKindJSDocSeeTag:                                  "JSDocSeeTag",
	SyntaxKindJSDocPropertyTag:                             "JSDocPropertyTag",
	SyntaxKindJSDocThrowsTag:                               "JSDocThrowsTag",
	SyntaxKindJSDocSatisfiesTag:                            "JSDocSatisfiesTag",
	SyntaxKindJSDocImportTag:                               "JSDocImportTag",
	SyntaxKindSyntaxList:                                   "SyntaxList",
	SyntaxKindNotEmittedStatement:                          "NotEmittedStatement",
	SyntaxKindPartiallyEmittedExpression:                   "PartiallyEmittedExpression",
	SyntaxKindCommaListExpression:                          "CommaListExpression",
	SyntaxKindSyntheticReferenceExpression:                 "SyntheticReferenceExpression",
	SyntaxKindCount:                                        "SyntaxKindCount",
}

func (k SyntaxKind) String() string {
	if k < 0 || k >= SyntaxKindCount {
		return fmt.Sprintf("SyntaxKind(%d)", k)
	}
	return syntaxKindNames[k]
}

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

type NodeData interface{} // TODO

type Factory struct { // TODO
	_IdentifierPool pool[Identifier]
}

type NodeBase struct{} // TODO

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
	v := f._IdentifierPool.allocate()
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

type QualifiedName struct {
	NodeBase
}

func (n *Node) AsQualifiedName() *QualifiedName { return n.data.(*QualifiedName) }

func (n *QualifiedName) reset() {
	*n = QualifiedName{}
}

func (n *QualifiedName) Kind() SyntaxKind { return SyntaxKindQualifiedName }

func NewQualifiedName() *QualifiedName {
	v := &QualifiedName{}
	v.reset()
	return v
}

func (f *Factory) NewQualifiedName() *QualifiedName {
	v := &QualifiedName{}
	v.reset()
	return v
}

type ComputedPropertyName struct {
	NodeBase
}

func (n *Node) AsComputedPropertyName() *ComputedPropertyName { return n.data.(*ComputedPropertyName) }

func (n *ComputedPropertyName) reset() {
	*n = ComputedPropertyName{}
}

func (n *ComputedPropertyName) Kind() SyntaxKind { return SyntaxKindComputedPropertyName }

func NewComputedPropertyName() *ComputedPropertyName {
	v := &ComputedPropertyName{}
	v.reset()
	return v
}

func (f *Factory) NewComputedPropertyName() *ComputedPropertyName {
	v := &ComputedPropertyName{}
	v.reset()
	return v
}

type ModifierList struct {
	NodeBase
}

func (n *Node) AsModifierList() *ModifierList { return n.data.(*ModifierList) }

func (n *ModifierList) reset() {
	*n = ModifierList{}
}

func (n *ModifierList) Kind() SyntaxKind { return SyntaxKindModifierList }

func NewModifierList() *ModifierList {
	v := &ModifierList{}
	v.reset()
	return v
}

func (f *Factory) NewModifierList() *ModifierList {
	v := &ModifierList{}
	v.reset()
	return v
}

type TypeParameterList struct {
	NodeBase
}

func (n *Node) AsTypeParameterList() *TypeParameterList { return n.data.(*TypeParameterList) }

func (n *TypeParameterList) reset() {
	*n = TypeParameterList{}
}

func (n *TypeParameterList) Kind() SyntaxKind { return SyntaxKindTypeParameterList }

func NewTypeParameterList() *TypeParameterList {
	v := &TypeParameterList{}
	v.reset()
	return v
}

func (f *Factory) NewTypeParameterList() *TypeParameterList {
	v := &TypeParameterList{}
	v.reset()
	return v
}

type TypeArgumentList struct {
	NodeBase
}

func (n *Node) AsTypeArgumentList() *TypeArgumentList { return n.data.(*TypeArgumentList) }

func (n *TypeArgumentList) reset() {
	*n = TypeArgumentList{}
}

func (n *TypeArgumentList) Kind() SyntaxKind { return SyntaxKindTypeArgumentList }

func NewTypeArgumentList() *TypeArgumentList {
	v := &TypeArgumentList{}
	v.reset()
	return v
}

func (f *Factory) NewTypeArgumentList() *TypeArgumentList {
	v := &TypeArgumentList{}
	v.reset()
	return v
}

type TypeParameter struct {
	NodeBase
}

func (n *Node) AsTypeParameter() *TypeParameter { return n.data.(*TypeParameter) }

func (n *TypeParameter) reset() {
	*n = TypeParameter{}
}

func (n *TypeParameter) Kind() SyntaxKind { return SyntaxKindTypeParameter }

func NewTypeParameter() *TypeParameter {
	v := &TypeParameter{}
	v.reset()
	return v
}

func (f *Factory) NewTypeParameter() *TypeParameter {
	v := &TypeParameter{}
	v.reset()
	return v
}

type Parameter struct {
	NodeBase
}

func (n *Node) AsParameter() *Parameter { return n.data.(*Parameter) }

func (n *Parameter) reset() {
	*n = Parameter{}
}

func (n *Parameter) Kind() SyntaxKind { return SyntaxKindParameter }

func NewParameter() *Parameter {
	v := &Parameter{}
	v.reset()
	return v
}

func (f *Factory) NewParameter() *Parameter {
	v := &Parameter{}
	v.reset()
	return v
}

type Decorator struct {
	NodeBase
}

func (n *Node) AsDecorator() *Decorator { return n.data.(*Decorator) }

func (n *Decorator) reset() {
	*n = Decorator{}
}

func (n *Decorator) Kind() SyntaxKind { return SyntaxKindDecorator }

func NewDecorator() *Decorator {
	v := &Decorator{}
	v.reset()
	return v
}

func (f *Factory) NewDecorator() *Decorator {
	v := &Decorator{}
	v.reset()
	return v
}

type PropertySignature struct {
	NodeBase
}

func (n *Node) AsPropertySignature() *PropertySignature { return n.data.(*PropertySignature) }

func (n *PropertySignature) reset() {
	*n = PropertySignature{}
}

func (n *PropertySignature) Kind() SyntaxKind { return SyntaxKindPropertySignature }

func NewPropertySignature() *PropertySignature {
	v := &PropertySignature{}
	v.reset()
	return v
}

func (f *Factory) NewPropertySignature() *PropertySignature {
	v := &PropertySignature{}
	v.reset()
	return v
}

type PropertyDeclaration struct {
	NodeBase
}

func (n *Node) AsPropertyDeclaration() *PropertyDeclaration { return n.data.(*PropertyDeclaration) }

func (n *PropertyDeclaration) reset() {
	*n = PropertyDeclaration{}
}

func (n *PropertyDeclaration) Kind() SyntaxKind { return SyntaxKindPropertyDeclaration }

func NewPropertyDeclaration() *PropertyDeclaration {
	v := &PropertyDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewPropertyDeclaration() *PropertyDeclaration {
	v := &PropertyDeclaration{}
	v.reset()
	return v
}

type MethodSignature struct {
	NodeBase
}

func (n *Node) AsMethodSignature() *MethodSignature { return n.data.(*MethodSignature) }

func (n *MethodSignature) reset() {
	*n = MethodSignature{}
}

func (n *MethodSignature) Kind() SyntaxKind { return SyntaxKindMethodSignature }

func NewMethodSignature() *MethodSignature {
	v := &MethodSignature{}
	v.reset()
	return v
}

func (f *Factory) NewMethodSignature() *MethodSignature {
	v := &MethodSignature{}
	v.reset()
	return v
}

type MethodDeclaration struct {
	NodeBase
}

func (n *Node) AsMethodDeclaration() *MethodDeclaration { return n.data.(*MethodDeclaration) }

func (n *MethodDeclaration) reset() {
	*n = MethodDeclaration{}
}

func (n *MethodDeclaration) Kind() SyntaxKind { return SyntaxKindMethodDeclaration }

func NewMethodDeclaration() *MethodDeclaration {
	v := &MethodDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewMethodDeclaration() *MethodDeclaration {
	v := &MethodDeclaration{}
	v.reset()
	return v
}

type ClassStaticBlockDeclaration struct {
	NodeBase
}

func (n *Node) AsClassStaticBlockDeclaration() *ClassStaticBlockDeclaration {
	return n.data.(*ClassStaticBlockDeclaration)
}

func (n *ClassStaticBlockDeclaration) reset() {
	*n = ClassStaticBlockDeclaration{}
}

func (n *ClassStaticBlockDeclaration) Kind() SyntaxKind { return SyntaxKindClassStaticBlockDeclaration }

func NewClassStaticBlockDeclaration() *ClassStaticBlockDeclaration {
	v := &ClassStaticBlockDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewClassStaticBlockDeclaration() *ClassStaticBlockDeclaration {
	v := &ClassStaticBlockDeclaration{}
	v.reset()
	return v
}

type Constructor struct {
	NodeBase
}

func (n *Node) AsConstructor() *Constructor { return n.data.(*Constructor) }

func (n *Constructor) reset() {
	*n = Constructor{}
}

func (n *Constructor) Kind() SyntaxKind { return SyntaxKindConstructor }

func NewConstructor() *Constructor {
	v := &Constructor{}
	v.reset()
	return v
}

func (f *Factory) NewConstructor() *Constructor {
	v := &Constructor{}
	v.reset()
	return v
}

type GetAccessor struct {
	NodeBase
}

func (n *Node) AsGetAccessor() *GetAccessor { return n.data.(*GetAccessor) }

func (n *GetAccessor) reset() {
	*n = GetAccessor{}
}

func (n *GetAccessor) Kind() SyntaxKind { return SyntaxKindGetAccessor }

func NewGetAccessor() *GetAccessor {
	v := &GetAccessor{}
	v.reset()
	return v
}

func (f *Factory) NewGetAccessor() *GetAccessor {
	v := &GetAccessor{}
	v.reset()
	return v
}

type SetAccessor struct {
	NodeBase
}

func (n *Node) AsSetAccessor() *SetAccessor { return n.data.(*SetAccessor) }

func (n *SetAccessor) reset() {
	*n = SetAccessor{}
}

func (n *SetAccessor) Kind() SyntaxKind { return SyntaxKindSetAccessor }

func NewSetAccessor() *SetAccessor {
	v := &SetAccessor{}
	v.reset()
	return v
}

func (f *Factory) NewSetAccessor() *SetAccessor {
	v := &SetAccessor{}
	v.reset()
	return v
}

type CallSignature struct {
	NodeBase
}

func (n *Node) AsCallSignature() *CallSignature { return n.data.(*CallSignature) }

func (n *CallSignature) reset() {
	*n = CallSignature{}
}

func (n *CallSignature) Kind() SyntaxKind { return SyntaxKindCallSignature }

func NewCallSignature() *CallSignature {
	v := &CallSignature{}
	v.reset()
	return v
}

func (f *Factory) NewCallSignature() *CallSignature {
	v := &CallSignature{}
	v.reset()
	return v
}

type ConstructSignature struct {
	NodeBase
}

func (n *Node) AsConstructSignature() *ConstructSignature { return n.data.(*ConstructSignature) }

func (n *ConstructSignature) reset() {
	*n = ConstructSignature{}
}

func (n *ConstructSignature) Kind() SyntaxKind { return SyntaxKindConstructSignature }

func NewConstructSignature() *ConstructSignature {
	v := &ConstructSignature{}
	v.reset()
	return v
}

func (f *Factory) NewConstructSignature() *ConstructSignature {
	v := &ConstructSignature{}
	v.reset()
	return v
}

type IndexSignature struct {
	NodeBase
}

func (n *Node) AsIndexSignature() *IndexSignature { return n.data.(*IndexSignature) }

func (n *IndexSignature) reset() {
	*n = IndexSignature{}
}

func (n *IndexSignature) Kind() SyntaxKind { return SyntaxKindIndexSignature }

func NewIndexSignature() *IndexSignature {
	v := &IndexSignature{}
	v.reset()
	return v
}

func (f *Factory) NewIndexSignature() *IndexSignature {
	v := &IndexSignature{}
	v.reset()
	return v
}

type TypePredicate struct {
	NodeBase
}

func (n *Node) AsTypePredicate() *TypePredicate { return n.data.(*TypePredicate) }

func (n *TypePredicate) reset() {
	*n = TypePredicate{}
}

func (n *TypePredicate) Kind() SyntaxKind { return SyntaxKindTypePredicate }

func NewTypePredicate() *TypePredicate {
	v := &TypePredicate{}
	v.reset()
	return v
}

func (f *Factory) NewTypePredicate() *TypePredicate {
	v := &TypePredicate{}
	v.reset()
	return v
}

type TypeReference struct {
	NodeBase
}

func (n *Node) AsTypeReference() *TypeReference { return n.data.(*TypeReference) }

func (n *TypeReference) reset() {
	*n = TypeReference{}
}

func (n *TypeReference) Kind() SyntaxKind { return SyntaxKindTypeReference }

func NewTypeReference() *TypeReference {
	v := &TypeReference{}
	v.reset()
	return v
}

func (f *Factory) NewTypeReference() *TypeReference {
	v := &TypeReference{}
	v.reset()
	return v
}

type FunctionType struct {
	NodeBase
}

func (n *Node) AsFunctionType() *FunctionType { return n.data.(*FunctionType) }

func (n *FunctionType) reset() {
	*n = FunctionType{}
}

func (n *FunctionType) Kind() SyntaxKind { return SyntaxKindFunctionType }

func NewFunctionType() *FunctionType {
	v := &FunctionType{}
	v.reset()
	return v
}

func (f *Factory) NewFunctionType() *FunctionType {
	v := &FunctionType{}
	v.reset()
	return v
}

type ConstructorType struct {
	NodeBase
}

func (n *Node) AsConstructorType() *ConstructorType { return n.data.(*ConstructorType) }

func (n *ConstructorType) reset() {
	*n = ConstructorType{}
}

func (n *ConstructorType) Kind() SyntaxKind { return SyntaxKindConstructorType }

func NewConstructorType() *ConstructorType {
	v := &ConstructorType{}
	v.reset()
	return v
}

func (f *Factory) NewConstructorType() *ConstructorType {
	v := &ConstructorType{}
	v.reset()
	return v
}

type TypeQuery struct {
	NodeBase
}

func (n *Node) AsTypeQuery() *TypeQuery { return n.data.(*TypeQuery) }

func (n *TypeQuery) reset() {
	*n = TypeQuery{}
}

func (n *TypeQuery) Kind() SyntaxKind { return SyntaxKindTypeQuery }

func NewTypeQuery() *TypeQuery {
	v := &TypeQuery{}
	v.reset()
	return v
}

func (f *Factory) NewTypeQuery() *TypeQuery {
	v := &TypeQuery{}
	v.reset()
	return v
}

type TypeLiteral struct {
	NodeBase
}

func (n *Node) AsTypeLiteral() *TypeLiteral { return n.data.(*TypeLiteral) }

func (n *TypeLiteral) reset() {
	*n = TypeLiteral{}
}

func (n *TypeLiteral) Kind() SyntaxKind { return SyntaxKindTypeLiteral }

func NewTypeLiteral() *TypeLiteral {
	v := &TypeLiteral{}
	v.reset()
	return v
}

func (f *Factory) NewTypeLiteral() *TypeLiteral {
	v := &TypeLiteral{}
	v.reset()
	return v
}

type ArrayType struct {
	NodeBase
}

func (n *Node) AsArrayType() *ArrayType { return n.data.(*ArrayType) }

func (n *ArrayType) reset() {
	*n = ArrayType{}
}

func (n *ArrayType) Kind() SyntaxKind { return SyntaxKindArrayType }

func NewArrayType() *ArrayType {
	v := &ArrayType{}
	v.reset()
	return v
}

func (f *Factory) NewArrayType() *ArrayType {
	v := &ArrayType{}
	v.reset()
	return v
}

type TupleType struct {
	NodeBase
}

func (n *Node) AsTupleType() *TupleType { return n.data.(*TupleType) }

func (n *TupleType) reset() {
	*n = TupleType{}
}

func (n *TupleType) Kind() SyntaxKind { return SyntaxKindTupleType }

func NewTupleType() *TupleType {
	v := &TupleType{}
	v.reset()
	return v
}

func (f *Factory) NewTupleType() *TupleType {
	v := &TupleType{}
	v.reset()
	return v
}

type OptionalType struct {
	NodeBase
}

func (n *Node) AsOptionalType() *OptionalType { return n.data.(*OptionalType) }

func (n *OptionalType) reset() {
	*n = OptionalType{}
}

func (n *OptionalType) Kind() SyntaxKind { return SyntaxKindOptionalType }

func NewOptionalType() *OptionalType {
	v := &OptionalType{}
	v.reset()
	return v
}

func (f *Factory) NewOptionalType() *OptionalType {
	v := &OptionalType{}
	v.reset()
	return v
}

type RestType struct {
	NodeBase
}

func (n *Node) AsRestType() *RestType { return n.data.(*RestType) }

func (n *RestType) reset() {
	*n = RestType{}
}

func (n *RestType) Kind() SyntaxKind { return SyntaxKindRestType }

func NewRestType() *RestType {
	v := &RestType{}
	v.reset()
	return v
}

func (f *Factory) NewRestType() *RestType {
	v := &RestType{}
	v.reset()
	return v
}

type UnionType struct {
	NodeBase
}

func (n *Node) AsUnionType() *UnionType { return n.data.(*UnionType) }

func (n *UnionType) reset() {
	*n = UnionType{}
}

func (n *UnionType) Kind() SyntaxKind { return SyntaxKindUnionType }

func NewUnionType() *UnionType {
	v := &UnionType{}
	v.reset()
	return v
}

func (f *Factory) NewUnionType() *UnionType {
	v := &UnionType{}
	v.reset()
	return v
}

type IntersectionType struct {
	NodeBase
}

func (n *Node) AsIntersectionType() *IntersectionType { return n.data.(*IntersectionType) }

func (n *IntersectionType) reset() {
	*n = IntersectionType{}
}

func (n *IntersectionType) Kind() SyntaxKind { return SyntaxKindIntersectionType }

func NewIntersectionType() *IntersectionType {
	v := &IntersectionType{}
	v.reset()
	return v
}

func (f *Factory) NewIntersectionType() *IntersectionType {
	v := &IntersectionType{}
	v.reset()
	return v
}

type ConditionalType struct {
	NodeBase
}

func (n *Node) AsConditionalType() *ConditionalType { return n.data.(*ConditionalType) }

func (n *ConditionalType) reset() {
	*n = ConditionalType{}
}

func (n *ConditionalType) Kind() SyntaxKind { return SyntaxKindConditionalType }

func NewConditionalType() *ConditionalType {
	v := &ConditionalType{}
	v.reset()
	return v
}

func (f *Factory) NewConditionalType() *ConditionalType {
	v := &ConditionalType{}
	v.reset()
	return v
}

type InferType struct {
	NodeBase
}

func (n *Node) AsInferType() *InferType { return n.data.(*InferType) }

func (n *InferType) reset() {
	*n = InferType{}
}

func (n *InferType) Kind() SyntaxKind { return SyntaxKindInferType }

func NewInferType() *InferType {
	v := &InferType{}
	v.reset()
	return v
}

func (f *Factory) NewInferType() *InferType {
	v := &InferType{}
	v.reset()
	return v
}

type ParenthesizedType struct {
	NodeBase
}

func (n *Node) AsParenthesizedType() *ParenthesizedType { return n.data.(*ParenthesizedType) }

func (n *ParenthesizedType) reset() {
	*n = ParenthesizedType{}
}

func (n *ParenthesizedType) Kind() SyntaxKind { return SyntaxKindParenthesizedType }

func NewParenthesizedType() *ParenthesizedType {
	v := &ParenthesizedType{}
	v.reset()
	return v
}

func (f *Factory) NewParenthesizedType() *ParenthesizedType {
	v := &ParenthesizedType{}
	v.reset()
	return v
}

type ThisType struct {
	NodeBase
}

func (n *Node) AsThisType() *ThisType { return n.data.(*ThisType) }

func (n *ThisType) reset() {
	*n = ThisType{}
}

func (n *ThisType) Kind() SyntaxKind { return SyntaxKindThisType }

func NewThisType() *ThisType {
	v := &ThisType{}
	v.reset()
	return v
}

func (f *Factory) NewThisType() *ThisType {
	v := &ThisType{}
	v.reset()
	return v
}

type TypeOperator struct {
	NodeBase
}

func (n *Node) AsTypeOperator() *TypeOperator { return n.data.(*TypeOperator) }

func (n *TypeOperator) reset() {
	*n = TypeOperator{}
}

func (n *TypeOperator) Kind() SyntaxKind { return SyntaxKindTypeOperator }

func NewTypeOperator() *TypeOperator {
	v := &TypeOperator{}
	v.reset()
	return v
}

func (f *Factory) NewTypeOperator() *TypeOperator {
	v := &TypeOperator{}
	v.reset()
	return v
}

type IndexedAccessType struct {
	NodeBase
}

func (n *Node) AsIndexedAccessType() *IndexedAccessType { return n.data.(*IndexedAccessType) }

func (n *IndexedAccessType) reset() {
	*n = IndexedAccessType{}
}

func (n *IndexedAccessType) Kind() SyntaxKind { return SyntaxKindIndexedAccessType }

func NewIndexedAccessType() *IndexedAccessType {
	v := &IndexedAccessType{}
	v.reset()
	return v
}

func (f *Factory) NewIndexedAccessType() *IndexedAccessType {
	v := &IndexedAccessType{}
	v.reset()
	return v
}

type MappedType struct {
	NodeBase
}

func (n *Node) AsMappedType() *MappedType { return n.data.(*MappedType) }

func (n *MappedType) reset() {
	*n = MappedType{}
}

func (n *MappedType) Kind() SyntaxKind { return SyntaxKindMappedType }

func NewMappedType() *MappedType {
	v := &MappedType{}
	v.reset()
	return v
}

func (f *Factory) NewMappedType() *MappedType {
	v := &MappedType{}
	v.reset()
	return v
}

type LiteralType struct {
	NodeBase
}

func (n *Node) AsLiteralType() *LiteralType { return n.data.(*LiteralType) }

func (n *LiteralType) reset() {
	*n = LiteralType{}
}

func (n *LiteralType) Kind() SyntaxKind { return SyntaxKindLiteralType }

func NewLiteralType() *LiteralType {
	v := &LiteralType{}
	v.reset()
	return v
}

func (f *Factory) NewLiteralType() *LiteralType {
	v := &LiteralType{}
	v.reset()
	return v
}

type NamedTupleMember struct {
	NodeBase
}

func (n *Node) AsNamedTupleMember() *NamedTupleMember { return n.data.(*NamedTupleMember) }

func (n *NamedTupleMember) reset() {
	*n = NamedTupleMember{}
}

func (n *NamedTupleMember) Kind() SyntaxKind { return SyntaxKindNamedTupleMember }

func NewNamedTupleMember() *NamedTupleMember {
	v := &NamedTupleMember{}
	v.reset()
	return v
}

func (f *Factory) NewNamedTupleMember() *NamedTupleMember {
	v := &NamedTupleMember{}
	v.reset()
	return v
}

type TemplateLiteralType struct {
	NodeBase
}

func (n *Node) AsTemplateLiteralType() *TemplateLiteralType { return n.data.(*TemplateLiteralType) }

func (n *TemplateLiteralType) reset() {
	*n = TemplateLiteralType{}
}

func (n *TemplateLiteralType) Kind() SyntaxKind { return SyntaxKindTemplateLiteralType }

func NewTemplateLiteralType() *TemplateLiteralType {
	v := &TemplateLiteralType{}
	v.reset()
	return v
}

func (f *Factory) NewTemplateLiteralType() *TemplateLiteralType {
	v := &TemplateLiteralType{}
	v.reset()
	return v
}

type TemplateLiteralTypeSpan struct {
	NodeBase
}

func (n *Node) AsTemplateLiteralTypeSpan() *TemplateLiteralTypeSpan {
	return n.data.(*TemplateLiteralTypeSpan)
}

func (n *TemplateLiteralTypeSpan) reset() {
	*n = TemplateLiteralTypeSpan{}
}

func (n *TemplateLiteralTypeSpan) Kind() SyntaxKind { return SyntaxKindTemplateLiteralTypeSpan }

func NewTemplateLiteralTypeSpan() *TemplateLiteralTypeSpan {
	v := &TemplateLiteralTypeSpan{}
	v.reset()
	return v
}

func (f *Factory) NewTemplateLiteralTypeSpan() *TemplateLiteralTypeSpan {
	v := &TemplateLiteralTypeSpan{}
	v.reset()
	return v
}

type ImportType struct {
	NodeBase
}

func (n *Node) AsImportType() *ImportType { return n.data.(*ImportType) }

func (n *ImportType) reset() {
	*n = ImportType{}
}

func (n *ImportType) Kind() SyntaxKind { return SyntaxKindImportType }

func NewImportType() *ImportType {
	v := &ImportType{}
	v.reset()
	return v
}

func (f *Factory) NewImportType() *ImportType {
	v := &ImportType{}
	v.reset()
	return v
}

type ObjectBindingPattern struct {
	NodeBase
}

func (n *Node) AsObjectBindingPattern() *ObjectBindingPattern { return n.data.(*ObjectBindingPattern) }

func (n *ObjectBindingPattern) reset() {
	*n = ObjectBindingPattern{}
}

func (n *ObjectBindingPattern) Kind() SyntaxKind { return SyntaxKindObjectBindingPattern }

func NewObjectBindingPattern() *ObjectBindingPattern {
	v := &ObjectBindingPattern{}
	v.reset()
	return v
}

func (f *Factory) NewObjectBindingPattern() *ObjectBindingPattern {
	v := &ObjectBindingPattern{}
	v.reset()
	return v
}

type ArrayBindingPattern struct {
	NodeBase
}

func (n *Node) AsArrayBindingPattern() *ArrayBindingPattern { return n.data.(*ArrayBindingPattern) }

func (n *ArrayBindingPattern) reset() {
	*n = ArrayBindingPattern{}
}

func (n *ArrayBindingPattern) Kind() SyntaxKind { return SyntaxKindArrayBindingPattern }

func NewArrayBindingPattern() *ArrayBindingPattern {
	v := &ArrayBindingPattern{}
	v.reset()
	return v
}

func (f *Factory) NewArrayBindingPattern() *ArrayBindingPattern {
	v := &ArrayBindingPattern{}
	v.reset()
	return v
}

type BindingElement struct {
	NodeBase
}

func (n *Node) AsBindingElement() *BindingElement { return n.data.(*BindingElement) }

func (n *BindingElement) reset() {
	*n = BindingElement{}
}

func (n *BindingElement) Kind() SyntaxKind { return SyntaxKindBindingElement }

func NewBindingElement() *BindingElement {
	v := &BindingElement{}
	v.reset()
	return v
}

func (f *Factory) NewBindingElement() *BindingElement {
	v := &BindingElement{}
	v.reset()
	return v
}

type ArrayLiteralExpression struct {
	NodeBase
}

func (n *Node) AsArrayLiteralExpression() *ArrayLiteralExpression {
	return n.data.(*ArrayLiteralExpression)
}

func (n *ArrayLiteralExpression) reset() {
	*n = ArrayLiteralExpression{}
}

func (n *ArrayLiteralExpression) Kind() SyntaxKind { return SyntaxKindArrayLiteralExpression }

func NewArrayLiteralExpression() *ArrayLiteralExpression {
	v := &ArrayLiteralExpression{}
	v.reset()
	return v
}

func (f *Factory) NewArrayLiteralExpression() *ArrayLiteralExpression {
	v := &ArrayLiteralExpression{}
	v.reset()
	return v
}

type ObjectLiteralExpression struct {
	NodeBase
}

func (n *Node) AsObjectLiteralExpression() *ObjectLiteralExpression {
	return n.data.(*ObjectLiteralExpression)
}

func (n *ObjectLiteralExpression) reset() {
	*n = ObjectLiteralExpression{}
}

func (n *ObjectLiteralExpression) Kind() SyntaxKind { return SyntaxKindObjectLiteralExpression }

func NewObjectLiteralExpression() *ObjectLiteralExpression {
	v := &ObjectLiteralExpression{}
	v.reset()
	return v
}

func (f *Factory) NewObjectLiteralExpression() *ObjectLiteralExpression {
	v := &ObjectLiteralExpression{}
	v.reset()
	return v
}

type PropertyAccessExpression struct {
	NodeBase
}

func (n *Node) AsPropertyAccessExpression() *PropertyAccessExpression {
	return n.data.(*PropertyAccessExpression)
}

func (n *PropertyAccessExpression) reset() {
	*n = PropertyAccessExpression{}
}

func (n *PropertyAccessExpression) Kind() SyntaxKind { return SyntaxKindPropertyAccessExpression }

func NewPropertyAccessExpression() *PropertyAccessExpression {
	v := &PropertyAccessExpression{}
	v.reset()
	return v
}

func (f *Factory) NewPropertyAccessExpression() *PropertyAccessExpression {
	v := &PropertyAccessExpression{}
	v.reset()
	return v
}

type ElementAccessExpression struct {
	NodeBase
}

func (n *Node) AsElementAccessExpression() *ElementAccessExpression {
	return n.data.(*ElementAccessExpression)
}

func (n *ElementAccessExpression) reset() {
	*n = ElementAccessExpression{}
}

func (n *ElementAccessExpression) Kind() SyntaxKind { return SyntaxKindElementAccessExpression }

func NewElementAccessExpression() *ElementAccessExpression {
	v := &ElementAccessExpression{}
	v.reset()
	return v
}

func (f *Factory) NewElementAccessExpression() *ElementAccessExpression {
	v := &ElementAccessExpression{}
	v.reset()
	return v
}

type CallExpression struct {
	NodeBase
}

func (n *Node) AsCallExpression() *CallExpression { return n.data.(*CallExpression) }

func (n *CallExpression) reset() {
	*n = CallExpression{}
}

func (n *CallExpression) Kind() SyntaxKind { return SyntaxKindCallExpression }

func NewCallExpression() *CallExpression {
	v := &CallExpression{}
	v.reset()
	return v
}

func (f *Factory) NewCallExpression() *CallExpression {
	v := &CallExpression{}
	v.reset()
	return v
}

type NewExpression struct {
	NodeBase
}

func (n *Node) AsNewExpression() *NewExpression { return n.data.(*NewExpression) }

func (n *NewExpression) reset() {
	*n = NewExpression{}
}

func (n *NewExpression) Kind() SyntaxKind { return SyntaxKindNewExpression }

func NewNewExpression() *NewExpression {
	v := &NewExpression{}
	v.reset()
	return v
}

func (f *Factory) NewNewExpression() *NewExpression {
	v := &NewExpression{}
	v.reset()
	return v
}

type TaggedTemplateExpression struct {
	NodeBase
}

func (n *Node) AsTaggedTemplateExpression() *TaggedTemplateExpression {
	return n.data.(*TaggedTemplateExpression)
}

func (n *TaggedTemplateExpression) reset() {
	*n = TaggedTemplateExpression{}
}

func (n *TaggedTemplateExpression) Kind() SyntaxKind { return SyntaxKindTaggedTemplateExpression }

func NewTaggedTemplateExpression() *TaggedTemplateExpression {
	v := &TaggedTemplateExpression{}
	v.reset()
	return v
}

func (f *Factory) NewTaggedTemplateExpression() *TaggedTemplateExpression {
	v := &TaggedTemplateExpression{}
	v.reset()
	return v
}

type TypeAssertionExpression struct {
	NodeBase
}

func (n *Node) AsTypeAssertionExpression() *TypeAssertionExpression {
	return n.data.(*TypeAssertionExpression)
}

func (n *TypeAssertionExpression) reset() {
	*n = TypeAssertionExpression{}
}

func (n *TypeAssertionExpression) Kind() SyntaxKind { return SyntaxKindTypeAssertionExpression }

func NewTypeAssertionExpression() *TypeAssertionExpression {
	v := &TypeAssertionExpression{}
	v.reset()
	return v
}

func (f *Factory) NewTypeAssertionExpression() *TypeAssertionExpression {
	v := &TypeAssertionExpression{}
	v.reset()
	return v
}

type ParenthesizedExpression struct {
	NodeBase
}

func (n *Node) AsParenthesizedExpression() *ParenthesizedExpression {
	return n.data.(*ParenthesizedExpression)
}

func (n *ParenthesizedExpression) reset() {
	*n = ParenthesizedExpression{}
}

func (n *ParenthesizedExpression) Kind() SyntaxKind { return SyntaxKindParenthesizedExpression }

func NewParenthesizedExpression() *ParenthesizedExpression {
	v := &ParenthesizedExpression{}
	v.reset()
	return v
}

func (f *Factory) NewParenthesizedExpression() *ParenthesizedExpression {
	v := &ParenthesizedExpression{}
	v.reset()
	return v
}

type FunctionExpression struct {
	NodeBase
}

func (n *Node) AsFunctionExpression() *FunctionExpression { return n.data.(*FunctionExpression) }

func (n *FunctionExpression) reset() {
	*n = FunctionExpression{}
}

func (n *FunctionExpression) Kind() SyntaxKind { return SyntaxKindFunctionExpression }

func NewFunctionExpression() *FunctionExpression {
	v := &FunctionExpression{}
	v.reset()
	return v
}

func (f *Factory) NewFunctionExpression() *FunctionExpression {
	v := &FunctionExpression{}
	v.reset()
	return v
}

type ArrowFunction struct {
	NodeBase
}

func (n *Node) AsArrowFunction() *ArrowFunction { return n.data.(*ArrowFunction) }

func (n *ArrowFunction) reset() {
	*n = ArrowFunction{}
}

func (n *ArrowFunction) Kind() SyntaxKind { return SyntaxKindArrowFunction }

func NewArrowFunction() *ArrowFunction {
	v := &ArrowFunction{}
	v.reset()
	return v
}

func (f *Factory) NewArrowFunction() *ArrowFunction {
	v := &ArrowFunction{}
	v.reset()
	return v
}

type DeleteExpression struct {
	NodeBase
}

func (n *Node) AsDeleteExpression() *DeleteExpression { return n.data.(*DeleteExpression) }

func (n *DeleteExpression) reset() {
	*n = DeleteExpression{}
}

func (n *DeleteExpression) Kind() SyntaxKind { return SyntaxKindDeleteExpression }

func NewDeleteExpression() *DeleteExpression {
	v := &DeleteExpression{}
	v.reset()
	return v
}

func (f *Factory) NewDeleteExpression() *DeleteExpression {
	v := &DeleteExpression{}
	v.reset()
	return v
}

type TypeOfExpression struct {
	NodeBase
}

func (n *Node) AsTypeOfExpression() *TypeOfExpression { return n.data.(*TypeOfExpression) }

func (n *TypeOfExpression) reset() {
	*n = TypeOfExpression{}
}

func (n *TypeOfExpression) Kind() SyntaxKind { return SyntaxKindTypeOfExpression }

func NewTypeOfExpression() *TypeOfExpression {
	v := &TypeOfExpression{}
	v.reset()
	return v
}

func (f *Factory) NewTypeOfExpression() *TypeOfExpression {
	v := &TypeOfExpression{}
	v.reset()
	return v
}

type VoidExpression struct {
	NodeBase
}

func (n *Node) AsVoidExpression() *VoidExpression { return n.data.(*VoidExpression) }

func (n *VoidExpression) reset() {
	*n = VoidExpression{}
}

func (n *VoidExpression) Kind() SyntaxKind { return SyntaxKindVoidExpression }

func NewVoidExpression() *VoidExpression {
	v := &VoidExpression{}
	v.reset()
	return v
}

func (f *Factory) NewVoidExpression() *VoidExpression {
	v := &VoidExpression{}
	v.reset()
	return v
}

type AwaitExpression struct {
	NodeBase
}

func (n *Node) AsAwaitExpression() *AwaitExpression { return n.data.(*AwaitExpression) }

func (n *AwaitExpression) reset() {
	*n = AwaitExpression{}
}

func (n *AwaitExpression) Kind() SyntaxKind { return SyntaxKindAwaitExpression }

func NewAwaitExpression() *AwaitExpression {
	v := &AwaitExpression{}
	v.reset()
	return v
}

func (f *Factory) NewAwaitExpression() *AwaitExpression {
	v := &AwaitExpression{}
	v.reset()
	return v
}

type PrefixUnaryExpression struct {
	NodeBase
}

func (n *Node) AsPrefixUnaryExpression() *PrefixUnaryExpression {
	return n.data.(*PrefixUnaryExpression)
}

func (n *PrefixUnaryExpression) reset() {
	*n = PrefixUnaryExpression{}
}

func (n *PrefixUnaryExpression) Kind() SyntaxKind { return SyntaxKindPrefixUnaryExpression }

func NewPrefixUnaryExpression() *PrefixUnaryExpression {
	v := &PrefixUnaryExpression{}
	v.reset()
	return v
}

func (f *Factory) NewPrefixUnaryExpression() *PrefixUnaryExpression {
	v := &PrefixUnaryExpression{}
	v.reset()
	return v
}

type PostfixUnaryExpression struct {
	NodeBase
}

func (n *Node) AsPostfixUnaryExpression() *PostfixUnaryExpression {
	return n.data.(*PostfixUnaryExpression)
}

func (n *PostfixUnaryExpression) reset() {
	*n = PostfixUnaryExpression{}
}

func (n *PostfixUnaryExpression) Kind() SyntaxKind { return SyntaxKindPostfixUnaryExpression }

func NewPostfixUnaryExpression() *PostfixUnaryExpression {
	v := &PostfixUnaryExpression{}
	v.reset()
	return v
}

func (f *Factory) NewPostfixUnaryExpression() *PostfixUnaryExpression {
	v := &PostfixUnaryExpression{}
	v.reset()
	return v
}

type BinaryExpression struct {
	NodeBase
}

func (n *Node) AsBinaryExpression() *BinaryExpression { return n.data.(*BinaryExpression) }

func (n *BinaryExpression) reset() {
	*n = BinaryExpression{}
}

func (n *BinaryExpression) Kind() SyntaxKind { return SyntaxKindBinaryExpression }

func NewBinaryExpression() *BinaryExpression {
	v := &BinaryExpression{}
	v.reset()
	return v
}

func (f *Factory) NewBinaryExpression() *BinaryExpression {
	v := &BinaryExpression{}
	v.reset()
	return v
}

type ConditionalExpression struct {
	NodeBase
}

func (n *Node) AsConditionalExpression() *ConditionalExpression {
	return n.data.(*ConditionalExpression)
}

func (n *ConditionalExpression) reset() {
	*n = ConditionalExpression{}
}

func (n *ConditionalExpression) Kind() SyntaxKind { return SyntaxKindConditionalExpression }

func NewConditionalExpression() *ConditionalExpression {
	v := &ConditionalExpression{}
	v.reset()
	return v
}

func (f *Factory) NewConditionalExpression() *ConditionalExpression {
	v := &ConditionalExpression{}
	v.reset()
	return v
}

type TemplateExpression struct {
	NodeBase
}

func (n *Node) AsTemplateExpression() *TemplateExpression { return n.data.(*TemplateExpression) }

func (n *TemplateExpression) reset() {
	*n = TemplateExpression{}
}

func (n *TemplateExpression) Kind() SyntaxKind { return SyntaxKindTemplateExpression }

func NewTemplateExpression() *TemplateExpression {
	v := &TemplateExpression{}
	v.reset()
	return v
}

func (f *Factory) NewTemplateExpression() *TemplateExpression {
	v := &TemplateExpression{}
	v.reset()
	return v
}

type YieldExpression struct {
	NodeBase
}

func (n *Node) AsYieldExpression() *YieldExpression { return n.data.(*YieldExpression) }

func (n *YieldExpression) reset() {
	*n = YieldExpression{}
}

func (n *YieldExpression) Kind() SyntaxKind { return SyntaxKindYieldExpression }

func NewYieldExpression() *YieldExpression {
	v := &YieldExpression{}
	v.reset()
	return v
}

func (f *Factory) NewYieldExpression() *YieldExpression {
	v := &YieldExpression{}
	v.reset()
	return v
}

type SpreadElement struct {
	NodeBase
}

func (n *Node) AsSpreadElement() *SpreadElement { return n.data.(*SpreadElement) }

func (n *SpreadElement) reset() {
	*n = SpreadElement{}
}

func (n *SpreadElement) Kind() SyntaxKind { return SyntaxKindSpreadElement }

func NewSpreadElement() *SpreadElement {
	v := &SpreadElement{}
	v.reset()
	return v
}

func (f *Factory) NewSpreadElement() *SpreadElement {
	v := &SpreadElement{}
	v.reset()
	return v
}

type ClassExpression struct {
	NodeBase
}

func (n *Node) AsClassExpression() *ClassExpression { return n.data.(*ClassExpression) }

func (n *ClassExpression) reset() {
	*n = ClassExpression{}
}

func (n *ClassExpression) Kind() SyntaxKind { return SyntaxKindClassExpression }

func NewClassExpression() *ClassExpression {
	v := &ClassExpression{}
	v.reset()
	return v
}

func (f *Factory) NewClassExpression() *ClassExpression {
	v := &ClassExpression{}
	v.reset()
	return v
}

type OmittedExpression struct {
	NodeBase
}

func (n *Node) AsOmittedExpression() *OmittedExpression { return n.data.(*OmittedExpression) }

func (n *OmittedExpression) reset() {
	*n = OmittedExpression{}
}

func (n *OmittedExpression) Kind() SyntaxKind { return SyntaxKindOmittedExpression }

func NewOmittedExpression() *OmittedExpression {
	v := &OmittedExpression{}
	v.reset()
	return v
}

func (f *Factory) NewOmittedExpression() *OmittedExpression {
	v := &OmittedExpression{}
	v.reset()
	return v
}

type ExpressionWithTypeArguments struct {
	NodeBase
}

func (n *Node) AsExpressionWithTypeArguments() *ExpressionWithTypeArguments {
	return n.data.(*ExpressionWithTypeArguments)
}

func (n *ExpressionWithTypeArguments) reset() {
	*n = ExpressionWithTypeArguments{}
}

func (n *ExpressionWithTypeArguments) Kind() SyntaxKind { return SyntaxKindExpressionWithTypeArguments }

func NewExpressionWithTypeArguments() *ExpressionWithTypeArguments {
	v := &ExpressionWithTypeArguments{}
	v.reset()
	return v
}

func (f *Factory) NewExpressionWithTypeArguments() *ExpressionWithTypeArguments {
	v := &ExpressionWithTypeArguments{}
	v.reset()
	return v
}

type AsExpression struct {
	NodeBase
}

func (n *Node) AsAsExpression() *AsExpression { return n.data.(*AsExpression) }

func (n *AsExpression) reset() {
	*n = AsExpression{}
}

func (n *AsExpression) Kind() SyntaxKind { return SyntaxKindAsExpression }

func NewAsExpression() *AsExpression {
	v := &AsExpression{}
	v.reset()
	return v
}

func (f *Factory) NewAsExpression() *AsExpression {
	v := &AsExpression{}
	v.reset()
	return v
}

type NonNullExpression struct {
	NodeBase
}

func (n *Node) AsNonNullExpression() *NonNullExpression { return n.data.(*NonNullExpression) }

func (n *NonNullExpression) reset() {
	*n = NonNullExpression{}
}

func (n *NonNullExpression) Kind() SyntaxKind { return SyntaxKindNonNullExpression }

func NewNonNullExpression() *NonNullExpression {
	v := &NonNullExpression{}
	v.reset()
	return v
}

func (f *Factory) NewNonNullExpression() *NonNullExpression {
	v := &NonNullExpression{}
	v.reset()
	return v
}

type MetaProperty struct {
	NodeBase
}

func (n *Node) AsMetaProperty() *MetaProperty { return n.data.(*MetaProperty) }

func (n *MetaProperty) reset() {
	*n = MetaProperty{}
}

func (n *MetaProperty) Kind() SyntaxKind { return SyntaxKindMetaProperty }

func NewMetaProperty() *MetaProperty {
	v := &MetaProperty{}
	v.reset()
	return v
}

func (f *Factory) NewMetaProperty() *MetaProperty {
	v := &MetaProperty{}
	v.reset()
	return v
}

type SyntheticExpression struct {
	NodeBase
}

func (n *Node) AsSyntheticExpression() *SyntheticExpression { return n.data.(*SyntheticExpression) }

func (n *SyntheticExpression) reset() {
	*n = SyntheticExpression{}
}

func (n *SyntheticExpression) Kind() SyntaxKind { return SyntaxKindSyntheticExpression }

func NewSyntheticExpression() *SyntheticExpression {
	v := &SyntheticExpression{}
	v.reset()
	return v
}

func (f *Factory) NewSyntheticExpression() *SyntheticExpression {
	v := &SyntheticExpression{}
	v.reset()
	return v
}

type SatisfiesExpression struct {
	NodeBase
}

func (n *Node) AsSatisfiesExpression() *SatisfiesExpression { return n.data.(*SatisfiesExpression) }

func (n *SatisfiesExpression) reset() {
	*n = SatisfiesExpression{}
}

func (n *SatisfiesExpression) Kind() SyntaxKind { return SyntaxKindSatisfiesExpression }

func NewSatisfiesExpression() *SatisfiesExpression {
	v := &SatisfiesExpression{}
	v.reset()
	return v
}

func (f *Factory) NewSatisfiesExpression() *SatisfiesExpression {
	v := &SatisfiesExpression{}
	v.reset()
	return v
}

type TemplateSpan struct {
	NodeBase
}

func (n *Node) AsTemplateSpan() *TemplateSpan { return n.data.(*TemplateSpan) }

func (n *TemplateSpan) reset() {
	*n = TemplateSpan{}
}

func (n *TemplateSpan) Kind() SyntaxKind { return SyntaxKindTemplateSpan }

func NewTemplateSpan() *TemplateSpan {
	v := &TemplateSpan{}
	v.reset()
	return v
}

func (f *Factory) NewTemplateSpan() *TemplateSpan {
	v := &TemplateSpan{}
	v.reset()
	return v
}

type SemicolonClassElement struct {
	NodeBase
}

func (n *Node) AsSemicolonClassElement() *SemicolonClassElement {
	return n.data.(*SemicolonClassElement)
}

func (n *SemicolonClassElement) reset() {
	*n = SemicolonClassElement{}
}

func (n *SemicolonClassElement) Kind() SyntaxKind { return SyntaxKindSemicolonClassElement }

func NewSemicolonClassElement() *SemicolonClassElement {
	v := &SemicolonClassElement{}
	v.reset()
	return v
}

func (f *Factory) NewSemicolonClassElement() *SemicolonClassElement {
	v := &SemicolonClassElement{}
	v.reset()
	return v
}

type Block struct {
	NodeBase
}

func (n *Node) AsBlock() *Block { return n.data.(*Block) }

func (n *Block) reset() {
	*n = Block{}
}

func (n *Block) Kind() SyntaxKind { return SyntaxKindBlock }

func NewBlock() *Block {
	v := &Block{}
	v.reset()
	return v
}

func (f *Factory) NewBlock() *Block {
	v := &Block{}
	v.reset()
	return v
}

type EmptyStatement struct {
	NodeBase
}

func (n *Node) AsEmptyStatement() *EmptyStatement { return n.data.(*EmptyStatement) }

func (n *EmptyStatement) reset() {
	*n = EmptyStatement{}
}

func (n *EmptyStatement) Kind() SyntaxKind { return SyntaxKindEmptyStatement }

func NewEmptyStatement() *EmptyStatement {
	v := &EmptyStatement{}
	v.reset()
	return v
}

func (f *Factory) NewEmptyStatement() *EmptyStatement {
	v := &EmptyStatement{}
	v.reset()
	return v
}

type VariableStatement struct {
	NodeBase
}

func (n *Node) AsVariableStatement() *VariableStatement { return n.data.(*VariableStatement) }

func (n *VariableStatement) reset() {
	*n = VariableStatement{}
}

func (n *VariableStatement) Kind() SyntaxKind { return SyntaxKindVariableStatement }

func NewVariableStatement() *VariableStatement {
	v := &VariableStatement{}
	v.reset()
	return v
}

func (f *Factory) NewVariableStatement() *VariableStatement {
	v := &VariableStatement{}
	v.reset()
	return v
}

type ExpressionStatement struct {
	NodeBase
}

func (n *Node) AsExpressionStatement() *ExpressionStatement { return n.data.(*ExpressionStatement) }

func (n *ExpressionStatement) reset() {
	*n = ExpressionStatement{}
}

func (n *ExpressionStatement) Kind() SyntaxKind { return SyntaxKindExpressionStatement }

func NewExpressionStatement() *ExpressionStatement {
	v := &ExpressionStatement{}
	v.reset()
	return v
}

func (f *Factory) NewExpressionStatement() *ExpressionStatement {
	v := &ExpressionStatement{}
	v.reset()
	return v
}

type IfStatement struct {
	NodeBase
}

func (n *Node) AsIfStatement() *IfStatement { return n.data.(*IfStatement) }

func (n *IfStatement) reset() {
	*n = IfStatement{}
}

func (n *IfStatement) Kind() SyntaxKind { return SyntaxKindIfStatement }

func NewIfStatement() *IfStatement {
	v := &IfStatement{}
	v.reset()
	return v
}

func (f *Factory) NewIfStatement() *IfStatement {
	v := &IfStatement{}
	v.reset()
	return v
}

type DoStatement struct {
	NodeBase
}

func (n *Node) AsDoStatement() *DoStatement { return n.data.(*DoStatement) }

func (n *DoStatement) reset() {
	*n = DoStatement{}
}

func (n *DoStatement) Kind() SyntaxKind { return SyntaxKindDoStatement }

func NewDoStatement() *DoStatement {
	v := &DoStatement{}
	v.reset()
	return v
}

func (f *Factory) NewDoStatement() *DoStatement {
	v := &DoStatement{}
	v.reset()
	return v
}

type WhileStatement struct {
	NodeBase
}

func (n *Node) AsWhileStatement() *WhileStatement { return n.data.(*WhileStatement) }

func (n *WhileStatement) reset() {
	*n = WhileStatement{}
}

func (n *WhileStatement) Kind() SyntaxKind { return SyntaxKindWhileStatement }

func NewWhileStatement() *WhileStatement {
	v := &WhileStatement{}
	v.reset()
	return v
}

func (f *Factory) NewWhileStatement() *WhileStatement {
	v := &WhileStatement{}
	v.reset()
	return v
}

type ForStatement struct {
	NodeBase
}

func (n *Node) AsForStatement() *ForStatement { return n.data.(*ForStatement) }

func (n *ForStatement) reset() {
	*n = ForStatement{}
}

func (n *ForStatement) Kind() SyntaxKind { return SyntaxKindForStatement }

func NewForStatement() *ForStatement {
	v := &ForStatement{}
	v.reset()
	return v
}

func (f *Factory) NewForStatement() *ForStatement {
	v := &ForStatement{}
	v.reset()
	return v
}

type ForInStatement struct {
	NodeBase
}

func (n *Node) AsForInStatement() *ForInStatement { return n.data.(*ForInStatement) }

func (n *ForInStatement) reset() {
	*n = ForInStatement{}
}

func (n *ForInStatement) Kind() SyntaxKind { return SyntaxKindForInStatement }

func NewForInStatement() *ForInStatement {
	v := &ForInStatement{}
	v.reset()
	return v
}

func (f *Factory) NewForInStatement() *ForInStatement {
	v := &ForInStatement{}
	v.reset()
	return v
}

type ForOfStatement struct {
	NodeBase
}

func (n *Node) AsForOfStatement() *ForOfStatement { return n.data.(*ForOfStatement) }

func (n *ForOfStatement) reset() {
	*n = ForOfStatement{}
}

func (n *ForOfStatement) Kind() SyntaxKind { return SyntaxKindForOfStatement }

func NewForOfStatement() *ForOfStatement {
	v := &ForOfStatement{}
	v.reset()
	return v
}

func (f *Factory) NewForOfStatement() *ForOfStatement {
	v := &ForOfStatement{}
	v.reset()
	return v
}

type ContinueStatement struct {
	NodeBase
}

func (n *Node) AsContinueStatement() *ContinueStatement { return n.data.(*ContinueStatement) }

func (n *ContinueStatement) reset() {
	*n = ContinueStatement{}
}

func (n *ContinueStatement) Kind() SyntaxKind { return SyntaxKindContinueStatement }

func NewContinueStatement() *ContinueStatement {
	v := &ContinueStatement{}
	v.reset()
	return v
}

func (f *Factory) NewContinueStatement() *ContinueStatement {
	v := &ContinueStatement{}
	v.reset()
	return v
}

type BreakStatement struct {
	NodeBase
}

func (n *Node) AsBreakStatement() *BreakStatement { return n.data.(*BreakStatement) }

func (n *BreakStatement) reset() {
	*n = BreakStatement{}
}

func (n *BreakStatement) Kind() SyntaxKind { return SyntaxKindBreakStatement }

func NewBreakStatement() *BreakStatement {
	v := &BreakStatement{}
	v.reset()
	return v
}

func (f *Factory) NewBreakStatement() *BreakStatement {
	v := &BreakStatement{}
	v.reset()
	return v
}

type ReturnStatement struct {
	NodeBase
}

func (n *Node) AsReturnStatement() *ReturnStatement { return n.data.(*ReturnStatement) }

func (n *ReturnStatement) reset() {
	*n = ReturnStatement{}
}

func (n *ReturnStatement) Kind() SyntaxKind { return SyntaxKindReturnStatement }

func NewReturnStatement() *ReturnStatement {
	v := &ReturnStatement{}
	v.reset()
	return v
}

func (f *Factory) NewReturnStatement() *ReturnStatement {
	v := &ReturnStatement{}
	v.reset()
	return v
}

type WithStatement struct {
	NodeBase
}

func (n *Node) AsWithStatement() *WithStatement { return n.data.(*WithStatement) }

func (n *WithStatement) reset() {
	*n = WithStatement{}
}

func (n *WithStatement) Kind() SyntaxKind { return SyntaxKindWithStatement }

func NewWithStatement() *WithStatement {
	v := &WithStatement{}
	v.reset()
	return v
}

func (f *Factory) NewWithStatement() *WithStatement {
	v := &WithStatement{}
	v.reset()
	return v
}

type SwitchStatement struct {
	NodeBase
}

func (n *Node) AsSwitchStatement() *SwitchStatement { return n.data.(*SwitchStatement) }

func (n *SwitchStatement) reset() {
	*n = SwitchStatement{}
}

func (n *SwitchStatement) Kind() SyntaxKind { return SyntaxKindSwitchStatement }

func NewSwitchStatement() *SwitchStatement {
	v := &SwitchStatement{}
	v.reset()
	return v
}

func (f *Factory) NewSwitchStatement() *SwitchStatement {
	v := &SwitchStatement{}
	v.reset()
	return v
}

type LabeledStatement struct {
	NodeBase
}

func (n *Node) AsLabeledStatement() *LabeledStatement { return n.data.(*LabeledStatement) }

func (n *LabeledStatement) reset() {
	*n = LabeledStatement{}
}

func (n *LabeledStatement) Kind() SyntaxKind { return SyntaxKindLabeledStatement }

func NewLabeledStatement() *LabeledStatement {
	v := &LabeledStatement{}
	v.reset()
	return v
}

func (f *Factory) NewLabeledStatement() *LabeledStatement {
	v := &LabeledStatement{}
	v.reset()
	return v
}

type ThrowStatement struct {
	NodeBase
}

func (n *Node) AsThrowStatement() *ThrowStatement { return n.data.(*ThrowStatement) }

func (n *ThrowStatement) reset() {
	*n = ThrowStatement{}
}

func (n *ThrowStatement) Kind() SyntaxKind { return SyntaxKindThrowStatement }

func NewThrowStatement() *ThrowStatement {
	v := &ThrowStatement{}
	v.reset()
	return v
}

func (f *Factory) NewThrowStatement() *ThrowStatement {
	v := &ThrowStatement{}
	v.reset()
	return v
}

type TryStatement struct {
	NodeBase
}

func (n *Node) AsTryStatement() *TryStatement { return n.data.(*TryStatement) }

func (n *TryStatement) reset() {
	*n = TryStatement{}
}

func (n *TryStatement) Kind() SyntaxKind { return SyntaxKindTryStatement }

func NewTryStatement() *TryStatement {
	v := &TryStatement{}
	v.reset()
	return v
}

func (f *Factory) NewTryStatement() *TryStatement {
	v := &TryStatement{}
	v.reset()
	return v
}

type DebuggerStatement struct {
	NodeBase
}

func (n *Node) AsDebuggerStatement() *DebuggerStatement { return n.data.(*DebuggerStatement) }

func (n *DebuggerStatement) reset() {
	*n = DebuggerStatement{}
}

func (n *DebuggerStatement) Kind() SyntaxKind { return SyntaxKindDebuggerStatement }

func NewDebuggerStatement() *DebuggerStatement {
	v := &DebuggerStatement{}
	v.reset()
	return v
}

func (f *Factory) NewDebuggerStatement() *DebuggerStatement {
	v := &DebuggerStatement{}
	v.reset()
	return v
}

type VariableDeclaration struct {
	NodeBase
}

func (n *Node) AsVariableDeclaration() *VariableDeclaration { return n.data.(*VariableDeclaration) }

func (n *VariableDeclaration) reset() {
	*n = VariableDeclaration{}
}

func (n *VariableDeclaration) Kind() SyntaxKind { return SyntaxKindVariableDeclaration }

func NewVariableDeclaration() *VariableDeclaration {
	v := &VariableDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewVariableDeclaration() *VariableDeclaration {
	v := &VariableDeclaration{}
	v.reset()
	return v
}

type VariableDeclarationList struct {
	NodeBase
}

func (n *Node) AsVariableDeclarationList() *VariableDeclarationList {
	return n.data.(*VariableDeclarationList)
}

func (n *VariableDeclarationList) reset() {
	*n = VariableDeclarationList{}
}

func (n *VariableDeclarationList) Kind() SyntaxKind { return SyntaxKindVariableDeclarationList }

func NewVariableDeclarationList() *VariableDeclarationList {
	v := &VariableDeclarationList{}
	v.reset()
	return v
}

func (f *Factory) NewVariableDeclarationList() *VariableDeclarationList {
	v := &VariableDeclarationList{}
	v.reset()
	return v
}

type FunctionDeclaration struct {
	NodeBase
}

func (n *Node) AsFunctionDeclaration() *FunctionDeclaration { return n.data.(*FunctionDeclaration) }

func (n *FunctionDeclaration) reset() {
	*n = FunctionDeclaration{}
}

func (n *FunctionDeclaration) Kind() SyntaxKind { return SyntaxKindFunctionDeclaration }

func NewFunctionDeclaration() *FunctionDeclaration {
	v := &FunctionDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewFunctionDeclaration() *FunctionDeclaration {
	v := &FunctionDeclaration{}
	v.reset()
	return v
}

type ClassDeclaration struct {
	NodeBase
}

func (n *Node) AsClassDeclaration() *ClassDeclaration { return n.data.(*ClassDeclaration) }

func (n *ClassDeclaration) reset() {
	*n = ClassDeclaration{}
}

func (n *ClassDeclaration) Kind() SyntaxKind { return SyntaxKindClassDeclaration }

func NewClassDeclaration() *ClassDeclaration {
	v := &ClassDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewClassDeclaration() *ClassDeclaration {
	v := &ClassDeclaration{}
	v.reset()
	return v
}

type InterfaceDeclaration struct {
	NodeBase
}

func (n *Node) AsInterfaceDeclaration() *InterfaceDeclaration { return n.data.(*InterfaceDeclaration) }

func (n *InterfaceDeclaration) reset() {
	*n = InterfaceDeclaration{}
}

func (n *InterfaceDeclaration) Kind() SyntaxKind { return SyntaxKindInterfaceDeclaration }

func NewInterfaceDeclaration() *InterfaceDeclaration {
	v := &InterfaceDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewInterfaceDeclaration() *InterfaceDeclaration {
	v := &InterfaceDeclaration{}
	v.reset()
	return v
}

type TypeAliasDeclaration struct {
	NodeBase
}

func (n *Node) AsTypeAliasDeclaration() *TypeAliasDeclaration { return n.data.(*TypeAliasDeclaration) }

func (n *TypeAliasDeclaration) reset() {
	*n = TypeAliasDeclaration{}
}

func (n *TypeAliasDeclaration) Kind() SyntaxKind { return SyntaxKindTypeAliasDeclaration }

func NewTypeAliasDeclaration() *TypeAliasDeclaration {
	v := &TypeAliasDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewTypeAliasDeclaration() *TypeAliasDeclaration {
	v := &TypeAliasDeclaration{}
	v.reset()
	return v
}

type EnumDeclaration struct {
	NodeBase
}

func (n *Node) AsEnumDeclaration() *EnumDeclaration { return n.data.(*EnumDeclaration) }

func (n *EnumDeclaration) reset() {
	*n = EnumDeclaration{}
}

func (n *EnumDeclaration) Kind() SyntaxKind { return SyntaxKindEnumDeclaration }

func NewEnumDeclaration() *EnumDeclaration {
	v := &EnumDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewEnumDeclaration() *EnumDeclaration {
	v := &EnumDeclaration{}
	v.reset()
	return v
}

type ModuleDeclaration struct {
	NodeBase
}

func (n *Node) AsModuleDeclaration() *ModuleDeclaration { return n.data.(*ModuleDeclaration) }

func (n *ModuleDeclaration) reset() {
	*n = ModuleDeclaration{}
}

func (n *ModuleDeclaration) Kind() SyntaxKind { return SyntaxKindModuleDeclaration }

func NewModuleDeclaration() *ModuleDeclaration {
	v := &ModuleDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewModuleDeclaration() *ModuleDeclaration {
	v := &ModuleDeclaration{}
	v.reset()
	return v
}

type ModuleBlock struct {
	NodeBase
}

func (n *Node) AsModuleBlock() *ModuleBlock { return n.data.(*ModuleBlock) }

func (n *ModuleBlock) reset() {
	*n = ModuleBlock{}
}

func (n *ModuleBlock) Kind() SyntaxKind { return SyntaxKindModuleBlock }

func NewModuleBlock() *ModuleBlock {
	v := &ModuleBlock{}
	v.reset()
	return v
}

func (f *Factory) NewModuleBlock() *ModuleBlock {
	v := &ModuleBlock{}
	v.reset()
	return v
}

type CaseBlock struct {
	NodeBase
}

func (n *Node) AsCaseBlock() *CaseBlock { return n.data.(*CaseBlock) }

func (n *CaseBlock) reset() {
	*n = CaseBlock{}
}

func (n *CaseBlock) Kind() SyntaxKind { return SyntaxKindCaseBlock }

func NewCaseBlock() *CaseBlock {
	v := &CaseBlock{}
	v.reset()
	return v
}

func (f *Factory) NewCaseBlock() *CaseBlock {
	v := &CaseBlock{}
	v.reset()
	return v
}

type NamespaceExportDeclaration struct {
	NodeBase
}

func (n *Node) AsNamespaceExportDeclaration() *NamespaceExportDeclaration {
	return n.data.(*NamespaceExportDeclaration)
}

func (n *NamespaceExportDeclaration) reset() {
	*n = NamespaceExportDeclaration{}
}

func (n *NamespaceExportDeclaration) Kind() SyntaxKind { return SyntaxKindNamespaceExportDeclaration }

func NewNamespaceExportDeclaration() *NamespaceExportDeclaration {
	v := &NamespaceExportDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewNamespaceExportDeclaration() *NamespaceExportDeclaration {
	v := &NamespaceExportDeclaration{}
	v.reset()
	return v
}

type ImportEqualsDeclaration struct {
	NodeBase
}

func (n *Node) AsImportEqualsDeclaration() *ImportEqualsDeclaration {
	return n.data.(*ImportEqualsDeclaration)
}

func (n *ImportEqualsDeclaration) reset() {
	*n = ImportEqualsDeclaration{}
}

func (n *ImportEqualsDeclaration) Kind() SyntaxKind { return SyntaxKindImportEqualsDeclaration }

func NewImportEqualsDeclaration() *ImportEqualsDeclaration {
	v := &ImportEqualsDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewImportEqualsDeclaration() *ImportEqualsDeclaration {
	v := &ImportEqualsDeclaration{}
	v.reset()
	return v
}

type ImportDeclaration struct {
	NodeBase
}

func (n *Node) AsImportDeclaration() *ImportDeclaration { return n.data.(*ImportDeclaration) }

func (n *ImportDeclaration) reset() {
	*n = ImportDeclaration{}
}

func (n *ImportDeclaration) Kind() SyntaxKind { return SyntaxKindImportDeclaration }

func NewImportDeclaration() *ImportDeclaration {
	v := &ImportDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewImportDeclaration() *ImportDeclaration {
	v := &ImportDeclaration{}
	v.reset()
	return v
}

type ImportClause struct {
	NodeBase
}

func (n *Node) AsImportClause() *ImportClause { return n.data.(*ImportClause) }

func (n *ImportClause) reset() {
	*n = ImportClause{}
}

func (n *ImportClause) Kind() SyntaxKind { return SyntaxKindImportClause }

func NewImportClause() *ImportClause {
	v := &ImportClause{}
	v.reset()
	return v
}

func (f *Factory) NewImportClause() *ImportClause {
	v := &ImportClause{}
	v.reset()
	return v
}

type NamespaceImport struct {
	NodeBase
}

func (n *Node) AsNamespaceImport() *NamespaceImport { return n.data.(*NamespaceImport) }

func (n *NamespaceImport) reset() {
	*n = NamespaceImport{}
}

func (n *NamespaceImport) Kind() SyntaxKind { return SyntaxKindNamespaceImport }

func NewNamespaceImport() *NamespaceImport {
	v := &NamespaceImport{}
	v.reset()
	return v
}

func (f *Factory) NewNamespaceImport() *NamespaceImport {
	v := &NamespaceImport{}
	v.reset()
	return v
}

type NamedImports struct {
	NodeBase
}

func (n *Node) AsNamedImports() *NamedImports { return n.data.(*NamedImports) }

func (n *NamedImports) reset() {
	*n = NamedImports{}
}

func (n *NamedImports) Kind() SyntaxKind { return SyntaxKindNamedImports }

func NewNamedImports() *NamedImports {
	v := &NamedImports{}
	v.reset()
	return v
}

func (f *Factory) NewNamedImports() *NamedImports {
	v := &NamedImports{}
	v.reset()
	return v
}

type ImportSpecifier struct {
	NodeBase
}

func (n *Node) AsImportSpecifier() *ImportSpecifier { return n.data.(*ImportSpecifier) }

func (n *ImportSpecifier) reset() {
	*n = ImportSpecifier{}
}

func (n *ImportSpecifier) Kind() SyntaxKind { return SyntaxKindImportSpecifier }

func NewImportSpecifier() *ImportSpecifier {
	v := &ImportSpecifier{}
	v.reset()
	return v
}

func (f *Factory) NewImportSpecifier() *ImportSpecifier {
	v := &ImportSpecifier{}
	v.reset()
	return v
}

type ExportAssignment struct {
	NodeBase
}

func (n *Node) AsExportAssignment() *ExportAssignment { return n.data.(*ExportAssignment) }

func (n *ExportAssignment) reset() {
	*n = ExportAssignment{}
}

func (n *ExportAssignment) Kind() SyntaxKind { return SyntaxKindExportAssignment }

func NewExportAssignment() *ExportAssignment {
	v := &ExportAssignment{}
	v.reset()
	return v
}

func (f *Factory) NewExportAssignment() *ExportAssignment {
	v := &ExportAssignment{}
	v.reset()
	return v
}

type ExportDeclaration struct {
	NodeBase
}

func (n *Node) AsExportDeclaration() *ExportDeclaration { return n.data.(*ExportDeclaration) }

func (n *ExportDeclaration) reset() {
	*n = ExportDeclaration{}
}

func (n *ExportDeclaration) Kind() SyntaxKind { return SyntaxKindExportDeclaration }

func NewExportDeclaration() *ExportDeclaration {
	v := &ExportDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewExportDeclaration() *ExportDeclaration {
	v := &ExportDeclaration{}
	v.reset()
	return v
}

type NamedExports struct {
	NodeBase
}

func (n *Node) AsNamedExports() *NamedExports { return n.data.(*NamedExports) }

func (n *NamedExports) reset() {
	*n = NamedExports{}
}

func (n *NamedExports) Kind() SyntaxKind { return SyntaxKindNamedExports }

func NewNamedExports() *NamedExports {
	v := &NamedExports{}
	v.reset()
	return v
}

func (f *Factory) NewNamedExports() *NamedExports {
	v := &NamedExports{}
	v.reset()
	return v
}

type NamespaceExport struct {
	NodeBase
}

func (n *Node) AsNamespaceExport() *NamespaceExport { return n.data.(*NamespaceExport) }

func (n *NamespaceExport) reset() {
	*n = NamespaceExport{}
}

func (n *NamespaceExport) Kind() SyntaxKind { return SyntaxKindNamespaceExport }

func NewNamespaceExport() *NamespaceExport {
	v := &NamespaceExport{}
	v.reset()
	return v
}

func (f *Factory) NewNamespaceExport() *NamespaceExport {
	v := &NamespaceExport{}
	v.reset()
	return v
}

type ExportSpecifier struct {
	NodeBase
}

func (n *Node) AsExportSpecifier() *ExportSpecifier { return n.data.(*ExportSpecifier) }

func (n *ExportSpecifier) reset() {
	*n = ExportSpecifier{}
}

func (n *ExportSpecifier) Kind() SyntaxKind { return SyntaxKindExportSpecifier }

func NewExportSpecifier() *ExportSpecifier {
	v := &ExportSpecifier{}
	v.reset()
	return v
}

func (f *Factory) NewExportSpecifier() *ExportSpecifier {
	v := &ExportSpecifier{}
	v.reset()
	return v
}

type MissingDeclaration struct {
	NodeBase
}

func (n *Node) AsMissingDeclaration() *MissingDeclaration { return n.data.(*MissingDeclaration) }

func (n *MissingDeclaration) reset() {
	*n = MissingDeclaration{}
}

func (n *MissingDeclaration) Kind() SyntaxKind { return SyntaxKindMissingDeclaration }

func NewMissingDeclaration() *MissingDeclaration {
	v := &MissingDeclaration{}
	v.reset()
	return v
}

func (f *Factory) NewMissingDeclaration() *MissingDeclaration {
	v := &MissingDeclaration{}
	v.reset()
	return v
}

type ExternalModuleReference struct {
	NodeBase
}

func (n *Node) AsExternalModuleReference() *ExternalModuleReference {
	return n.data.(*ExternalModuleReference)
}

func (n *ExternalModuleReference) reset() {
	*n = ExternalModuleReference{}
}

func (n *ExternalModuleReference) Kind() SyntaxKind { return SyntaxKindExternalModuleReference }

func NewExternalModuleReference() *ExternalModuleReference {
	v := &ExternalModuleReference{}
	v.reset()
	return v
}

func (f *Factory) NewExternalModuleReference() *ExternalModuleReference {
	v := &ExternalModuleReference{}
	v.reset()
	return v
}

type JsxElement struct {
	NodeBase
}

func (n *Node) AsJsxElement() *JsxElement { return n.data.(*JsxElement) }

func (n *JsxElement) reset() {
	*n = JsxElement{}
}

func (n *JsxElement) Kind() SyntaxKind { return SyntaxKindJsxElement }

func NewJsxElement() *JsxElement {
	v := &JsxElement{}
	v.reset()
	return v
}

func (f *Factory) NewJsxElement() *JsxElement {
	v := &JsxElement{}
	v.reset()
	return v
}

type JsxSelfClosingElement struct {
	NodeBase
}

func (n *Node) AsJsxSelfClosingElement() *JsxSelfClosingElement {
	return n.data.(*JsxSelfClosingElement)
}

func (n *JsxSelfClosingElement) reset() {
	*n = JsxSelfClosingElement{}
}

func (n *JsxSelfClosingElement) Kind() SyntaxKind { return SyntaxKindJsxSelfClosingElement }

func NewJsxSelfClosingElement() *JsxSelfClosingElement {
	v := &JsxSelfClosingElement{}
	v.reset()
	return v
}

func (f *Factory) NewJsxSelfClosingElement() *JsxSelfClosingElement {
	v := &JsxSelfClosingElement{}
	v.reset()
	return v
}

type JsxOpeningElement struct {
	NodeBase
}

func (n *Node) AsJsxOpeningElement() *JsxOpeningElement { return n.data.(*JsxOpeningElement) }

func (n *JsxOpeningElement) reset() {
	*n = JsxOpeningElement{}
}

func (n *JsxOpeningElement) Kind() SyntaxKind { return SyntaxKindJsxOpeningElement }

func NewJsxOpeningElement() *JsxOpeningElement {
	v := &JsxOpeningElement{}
	v.reset()
	return v
}

func (f *Factory) NewJsxOpeningElement() *JsxOpeningElement {
	v := &JsxOpeningElement{}
	v.reset()
	return v
}

type JsxClosingElement struct {
	NodeBase
}

func (n *Node) AsJsxClosingElement() *JsxClosingElement { return n.data.(*JsxClosingElement) }

func (n *JsxClosingElement) reset() {
	*n = JsxClosingElement{}
}

func (n *JsxClosingElement) Kind() SyntaxKind { return SyntaxKindJsxClosingElement }

func NewJsxClosingElement() *JsxClosingElement {
	v := &JsxClosingElement{}
	v.reset()
	return v
}

func (f *Factory) NewJsxClosingElement() *JsxClosingElement {
	v := &JsxClosingElement{}
	v.reset()
	return v
}

type JsxFragment struct {
	NodeBase
}

func (n *Node) AsJsxFragment() *JsxFragment { return n.data.(*JsxFragment) }

func (n *JsxFragment) reset() {
	*n = JsxFragment{}
}

func (n *JsxFragment) Kind() SyntaxKind { return SyntaxKindJsxFragment }

func NewJsxFragment() *JsxFragment {
	v := &JsxFragment{}
	v.reset()
	return v
}

func (f *Factory) NewJsxFragment() *JsxFragment {
	v := &JsxFragment{}
	v.reset()
	return v
}

type JsxOpeningFragment struct {
	NodeBase
}

func (n *Node) AsJsxOpeningFragment() *JsxOpeningFragment { return n.data.(*JsxOpeningFragment) }

func (n *JsxOpeningFragment) reset() {
	*n = JsxOpeningFragment{}
}

func (n *JsxOpeningFragment) Kind() SyntaxKind { return SyntaxKindJsxOpeningFragment }

func NewJsxOpeningFragment() *JsxOpeningFragment {
	v := &JsxOpeningFragment{}
	v.reset()
	return v
}

func (f *Factory) NewJsxOpeningFragment() *JsxOpeningFragment {
	v := &JsxOpeningFragment{}
	v.reset()
	return v
}

type JsxClosingFragment struct {
	NodeBase
}

func (n *Node) AsJsxClosingFragment() *JsxClosingFragment { return n.data.(*JsxClosingFragment) }

func (n *JsxClosingFragment) reset() {
	*n = JsxClosingFragment{}
}

func (n *JsxClosingFragment) Kind() SyntaxKind { return SyntaxKindJsxClosingFragment }

func NewJsxClosingFragment() *JsxClosingFragment {
	v := &JsxClosingFragment{}
	v.reset()
	return v
}

func (f *Factory) NewJsxClosingFragment() *JsxClosingFragment {
	v := &JsxClosingFragment{}
	v.reset()
	return v
}

type JsxAttribute struct {
	NodeBase
}

func (n *Node) AsJsxAttribute() *JsxAttribute { return n.data.(*JsxAttribute) }

func (n *JsxAttribute) reset() {
	*n = JsxAttribute{}
}

func (n *JsxAttribute) Kind() SyntaxKind { return SyntaxKindJsxAttribute }

func NewJsxAttribute() *JsxAttribute {
	v := &JsxAttribute{}
	v.reset()
	return v
}

func (f *Factory) NewJsxAttribute() *JsxAttribute {
	v := &JsxAttribute{}
	v.reset()
	return v
}

type JsxAttributes struct {
	NodeBase
}

func (n *Node) AsJsxAttributes() *JsxAttributes { return n.data.(*JsxAttributes) }

func (n *JsxAttributes) reset() {
	*n = JsxAttributes{}
}

func (n *JsxAttributes) Kind() SyntaxKind { return SyntaxKindJsxAttributes }

func NewJsxAttributes() *JsxAttributes {
	v := &JsxAttributes{}
	v.reset()
	return v
}

func (f *Factory) NewJsxAttributes() *JsxAttributes {
	v := &JsxAttributes{}
	v.reset()
	return v
}

type JsxSpreadAttribute struct {
	NodeBase
}

func (n *Node) AsJsxSpreadAttribute() *JsxSpreadAttribute { return n.data.(*JsxSpreadAttribute) }

func (n *JsxSpreadAttribute) reset() {
	*n = JsxSpreadAttribute{}
}

func (n *JsxSpreadAttribute) Kind() SyntaxKind { return SyntaxKindJsxSpreadAttribute }

func NewJsxSpreadAttribute() *JsxSpreadAttribute {
	v := &JsxSpreadAttribute{}
	v.reset()
	return v
}

func (f *Factory) NewJsxSpreadAttribute() *JsxSpreadAttribute {
	v := &JsxSpreadAttribute{}
	v.reset()
	return v
}

type JsxExpression struct {
	NodeBase
}

func (n *Node) AsJsxExpression() *JsxExpression { return n.data.(*JsxExpression) }

func (n *JsxExpression) reset() {
	*n = JsxExpression{}
}

func (n *JsxExpression) Kind() SyntaxKind { return SyntaxKindJsxExpression }

func NewJsxExpression() *JsxExpression {
	v := &JsxExpression{}
	v.reset()
	return v
}

func (f *Factory) NewJsxExpression() *JsxExpression {
	v := &JsxExpression{}
	v.reset()
	return v
}

type JsxNamespacedName struct {
	NodeBase
}

func (n *Node) AsJsxNamespacedName() *JsxNamespacedName { return n.data.(*JsxNamespacedName) }

func (n *JsxNamespacedName) reset() {
	*n = JsxNamespacedName{}
}

func (n *JsxNamespacedName) Kind() SyntaxKind { return SyntaxKindJsxNamespacedName }

func NewJsxNamespacedName() *JsxNamespacedName {
	v := &JsxNamespacedName{}
	v.reset()
	return v
}

func (f *Factory) NewJsxNamespacedName() *JsxNamespacedName {
	v := &JsxNamespacedName{}
	v.reset()
	return v
}

type CaseClause struct {
	NodeBase
}

func (n *Node) AsCaseClause() *CaseClause { return n.data.(*CaseClause) }

func (n *CaseClause) reset() {
	*n = CaseClause{}
}

func (n *CaseClause) Kind() SyntaxKind { return SyntaxKindCaseClause }

func NewCaseClause() *CaseClause {
	v := &CaseClause{}
	v.reset()
	return v
}

func (f *Factory) NewCaseClause() *CaseClause {
	v := &CaseClause{}
	v.reset()
	return v
}

type DefaultClause struct {
	NodeBase
}

func (n *Node) AsDefaultClause() *DefaultClause { return n.data.(*DefaultClause) }

func (n *DefaultClause) reset() {
	*n = DefaultClause{}
}

func (n *DefaultClause) Kind() SyntaxKind { return SyntaxKindDefaultClause }

func NewDefaultClause() *DefaultClause {
	v := &DefaultClause{}
	v.reset()
	return v
}

func (f *Factory) NewDefaultClause() *DefaultClause {
	v := &DefaultClause{}
	v.reset()
	return v
}

type HeritageClause struct {
	NodeBase
}

func (n *Node) AsHeritageClause() *HeritageClause { return n.data.(*HeritageClause) }

func (n *HeritageClause) reset() {
	*n = HeritageClause{}
}

func (n *HeritageClause) Kind() SyntaxKind { return SyntaxKindHeritageClause }

func NewHeritageClause() *HeritageClause {
	v := &HeritageClause{}
	v.reset()
	return v
}

func (f *Factory) NewHeritageClause() *HeritageClause {
	v := &HeritageClause{}
	v.reset()
	return v
}

type CatchClause struct {
	NodeBase
}

func (n *Node) AsCatchClause() *CatchClause { return n.data.(*CatchClause) }

func (n *CatchClause) reset() {
	*n = CatchClause{}
}

func (n *CatchClause) Kind() SyntaxKind { return SyntaxKindCatchClause }

func NewCatchClause() *CatchClause {
	v := &CatchClause{}
	v.reset()
	return v
}

func (f *Factory) NewCatchClause() *CatchClause {
	v := &CatchClause{}
	v.reset()
	return v
}

type ImportAttributes struct {
	NodeBase
}

func (n *Node) AsImportAttributes() *ImportAttributes { return n.data.(*ImportAttributes) }

func (n *ImportAttributes) reset() {
	*n = ImportAttributes{}
}

func (n *ImportAttributes) Kind() SyntaxKind { return SyntaxKindImportAttributes }

func NewImportAttributes() *ImportAttributes {
	v := &ImportAttributes{}
	v.reset()
	return v
}

func (f *Factory) NewImportAttributes() *ImportAttributes {
	v := &ImportAttributes{}
	v.reset()
	return v
}

type ImportAttribute struct {
	NodeBase
}

func (n *Node) AsImportAttribute() *ImportAttribute { return n.data.(*ImportAttribute) }

func (n *ImportAttribute) reset() {
	*n = ImportAttribute{}
}

func (n *ImportAttribute) Kind() SyntaxKind { return SyntaxKindImportAttribute }

func NewImportAttribute() *ImportAttribute {
	v := &ImportAttribute{}
	v.reset()
	return v
}

func (f *Factory) NewImportAttribute() *ImportAttribute {
	v := &ImportAttribute{}
	v.reset()
	return v
}

type PropertyAssignment struct {
	NodeBase
}

func (n *Node) AsPropertyAssignment() *PropertyAssignment { return n.data.(*PropertyAssignment) }

func (n *PropertyAssignment) reset() {
	*n = PropertyAssignment{}
}

func (n *PropertyAssignment) Kind() SyntaxKind { return SyntaxKindPropertyAssignment }

func NewPropertyAssignment() *PropertyAssignment {
	v := &PropertyAssignment{}
	v.reset()
	return v
}

func (f *Factory) NewPropertyAssignment() *PropertyAssignment {
	v := &PropertyAssignment{}
	v.reset()
	return v
}

type ShorthandPropertyAssignment struct {
	NodeBase
}

func (n *Node) AsShorthandPropertyAssignment() *ShorthandPropertyAssignment {
	return n.data.(*ShorthandPropertyAssignment)
}

func (n *ShorthandPropertyAssignment) reset() {
	*n = ShorthandPropertyAssignment{}
}

func (n *ShorthandPropertyAssignment) Kind() SyntaxKind { return SyntaxKindShorthandPropertyAssignment }

func NewShorthandPropertyAssignment() *ShorthandPropertyAssignment {
	v := &ShorthandPropertyAssignment{}
	v.reset()
	return v
}

func (f *Factory) NewShorthandPropertyAssignment() *ShorthandPropertyAssignment {
	v := &ShorthandPropertyAssignment{}
	v.reset()
	return v
}

type SpreadAssignment struct {
	NodeBase
}

func (n *Node) AsSpreadAssignment() *SpreadAssignment { return n.data.(*SpreadAssignment) }

func (n *SpreadAssignment) reset() {
	*n = SpreadAssignment{}
}

func (n *SpreadAssignment) Kind() SyntaxKind { return SyntaxKindSpreadAssignment }

func NewSpreadAssignment() *SpreadAssignment {
	v := &SpreadAssignment{}
	v.reset()
	return v
}

func (f *Factory) NewSpreadAssignment() *SpreadAssignment {
	v := &SpreadAssignment{}
	v.reset()
	return v
}

type EnumMember struct {
	NodeBase
}

func (n *Node) AsEnumMember() *EnumMember { return n.data.(*EnumMember) }

func (n *EnumMember) reset() {
	*n = EnumMember{}
}

func (n *EnumMember) Kind() SyntaxKind { return SyntaxKindEnumMember }

func NewEnumMember() *EnumMember {
	v := &EnumMember{}
	v.reset()
	return v
}

func (f *Factory) NewEnumMember() *EnumMember {
	v := &EnumMember{}
	v.reset()
	return v
}

type SourceFile struct {
	NodeBase
}

func (n *Node) AsSourceFile() *SourceFile { return n.data.(*SourceFile) }

func (n *SourceFile) reset() {
	*n = SourceFile{}
}

func (n *SourceFile) Kind() SyntaxKind { return SyntaxKindSourceFile }

func NewSourceFile() *SourceFile {
	v := &SourceFile{}
	v.reset()
	return v
}

func (f *Factory) NewSourceFile() *SourceFile {
	v := &SourceFile{}
	v.reset()
	return v
}

type Bundle struct {
	NodeBase
}

func (n *Node) AsBundle() *Bundle { return n.data.(*Bundle) }

func (n *Bundle) reset() {
	*n = Bundle{}
}

func (n *Bundle) Kind() SyntaxKind { return SyntaxKindBundle }

func NewBundle() *Bundle {
	v := &Bundle{}
	v.reset()
	return v
}

func (f *Factory) NewBundle() *Bundle {
	v := &Bundle{}
	v.reset()
	return v
}

type JSDocTypeExpression struct {
	NodeBase
}

func (n *Node) AsJSDocTypeExpression() *JSDocTypeExpression { return n.data.(*JSDocTypeExpression) }

func (n *JSDocTypeExpression) reset() {
	*n = JSDocTypeExpression{}
}

func (n *JSDocTypeExpression) Kind() SyntaxKind { return SyntaxKindJSDocTypeExpression }

func NewJSDocTypeExpression() *JSDocTypeExpression {
	v := &JSDocTypeExpression{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocTypeExpression() *JSDocTypeExpression {
	v := &JSDocTypeExpression{}
	v.reset()
	return v
}

type JSDocNameReference struct {
	NodeBase
}

func (n *Node) AsJSDocNameReference() *JSDocNameReference { return n.data.(*JSDocNameReference) }

func (n *JSDocNameReference) reset() {
	*n = JSDocNameReference{}
}

func (n *JSDocNameReference) Kind() SyntaxKind { return SyntaxKindJSDocNameReference }

func NewJSDocNameReference() *JSDocNameReference {
	v := &JSDocNameReference{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocNameReference() *JSDocNameReference {
	v := &JSDocNameReference{}
	v.reset()
	return v
}

type JSDocMemberName struct {
	NodeBase
}

func (n *Node) AsJSDocMemberName() *JSDocMemberName { return n.data.(*JSDocMemberName) }

func (n *JSDocMemberName) reset() {
	*n = JSDocMemberName{}
}

func (n *JSDocMemberName) Kind() SyntaxKind { return SyntaxKindJSDocMemberName }

func NewJSDocMemberName() *JSDocMemberName {
	v := &JSDocMemberName{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocMemberName() *JSDocMemberName {
	v := &JSDocMemberName{}
	v.reset()
	return v
}

type JSDocAllType struct {
	NodeBase
}

func (n *Node) AsJSDocAllType() *JSDocAllType { return n.data.(*JSDocAllType) }

func (n *JSDocAllType) reset() {
	*n = JSDocAllType{}
}

func (n *JSDocAllType) Kind() SyntaxKind { return SyntaxKindJSDocAllType }

func NewJSDocAllType() *JSDocAllType {
	v := &JSDocAllType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocAllType() *JSDocAllType {
	v := &JSDocAllType{}
	v.reset()
	return v
}

type JSDocUnknownType struct {
	NodeBase
}

func (n *Node) AsJSDocUnknownType() *JSDocUnknownType { return n.data.(*JSDocUnknownType) }

func (n *JSDocUnknownType) reset() {
	*n = JSDocUnknownType{}
}

func (n *JSDocUnknownType) Kind() SyntaxKind { return SyntaxKindJSDocUnknownType }

func NewJSDocUnknownType() *JSDocUnknownType {
	v := &JSDocUnknownType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocUnknownType() *JSDocUnknownType {
	v := &JSDocUnknownType{}
	v.reset()
	return v
}

type JSDocNullableType struct {
	NodeBase
}

func (n *Node) AsJSDocNullableType() *JSDocNullableType { return n.data.(*JSDocNullableType) }

func (n *JSDocNullableType) reset() {
	*n = JSDocNullableType{}
}

func (n *JSDocNullableType) Kind() SyntaxKind { return SyntaxKindJSDocNullableType }

func NewJSDocNullableType() *JSDocNullableType {
	v := &JSDocNullableType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocNullableType() *JSDocNullableType {
	v := &JSDocNullableType{}
	v.reset()
	return v
}

type JSDocNonNullableType struct {
	NodeBase
}

func (n *Node) AsJSDocNonNullableType() *JSDocNonNullableType { return n.data.(*JSDocNonNullableType) }

func (n *JSDocNonNullableType) reset() {
	*n = JSDocNonNullableType{}
}

func (n *JSDocNonNullableType) Kind() SyntaxKind { return SyntaxKindJSDocNonNullableType }

func NewJSDocNonNullableType() *JSDocNonNullableType {
	v := &JSDocNonNullableType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocNonNullableType() *JSDocNonNullableType {
	v := &JSDocNonNullableType{}
	v.reset()
	return v
}

type JSDocOptionalType struct {
	NodeBase
}

func (n *Node) AsJSDocOptionalType() *JSDocOptionalType { return n.data.(*JSDocOptionalType) }

func (n *JSDocOptionalType) reset() {
	*n = JSDocOptionalType{}
}

func (n *JSDocOptionalType) Kind() SyntaxKind { return SyntaxKindJSDocOptionalType }

func NewJSDocOptionalType() *JSDocOptionalType {
	v := &JSDocOptionalType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocOptionalType() *JSDocOptionalType {
	v := &JSDocOptionalType{}
	v.reset()
	return v
}

type JSDocFunctionType struct {
	NodeBase
}

func (n *Node) AsJSDocFunctionType() *JSDocFunctionType { return n.data.(*JSDocFunctionType) }

func (n *JSDocFunctionType) reset() {
	*n = JSDocFunctionType{}
}

func (n *JSDocFunctionType) Kind() SyntaxKind { return SyntaxKindJSDocFunctionType }

func NewJSDocFunctionType() *JSDocFunctionType {
	v := &JSDocFunctionType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocFunctionType() *JSDocFunctionType {
	v := &JSDocFunctionType{}
	v.reset()
	return v
}

type JSDocVariadicType struct {
	NodeBase
}

func (n *Node) AsJSDocVariadicType() *JSDocVariadicType { return n.data.(*JSDocVariadicType) }

func (n *JSDocVariadicType) reset() {
	*n = JSDocVariadicType{}
}

func (n *JSDocVariadicType) Kind() SyntaxKind { return SyntaxKindJSDocVariadicType }

func NewJSDocVariadicType() *JSDocVariadicType {
	v := &JSDocVariadicType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocVariadicType() *JSDocVariadicType {
	v := &JSDocVariadicType{}
	v.reset()
	return v
}

type JSDocNamepathType struct {
	NodeBase
}

func (n *Node) AsJSDocNamepathType() *JSDocNamepathType { return n.data.(*JSDocNamepathType) }

func (n *JSDocNamepathType) reset() {
	*n = JSDocNamepathType{}
}

func (n *JSDocNamepathType) Kind() SyntaxKind { return SyntaxKindJSDocNamepathType }

func NewJSDocNamepathType() *JSDocNamepathType {
	v := &JSDocNamepathType{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocNamepathType() *JSDocNamepathType {
	v := &JSDocNamepathType{}
	v.reset()
	return v
}

type JSDoc struct {
	NodeBase
}

func (n *Node) AsJSDoc() *JSDoc { return n.data.(*JSDoc) }

func (n *JSDoc) reset() {
	*n = JSDoc{}
}

func (n *JSDoc) Kind() SyntaxKind { return SyntaxKindJSDoc }

func NewJSDoc() *JSDoc {
	v := &JSDoc{}
	v.reset()
	return v
}

func (f *Factory) NewJSDoc() *JSDoc {
	v := &JSDoc{}
	v.reset()
	return v
}

type JSDocText struct {
	NodeBase
}

func (n *Node) AsJSDocText() *JSDocText { return n.data.(*JSDocText) }

func (n *JSDocText) reset() {
	*n = JSDocText{}
}

func (n *JSDocText) Kind() SyntaxKind { return SyntaxKindJSDocText }

func NewJSDocText() *JSDocText {
	v := &JSDocText{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocText() *JSDocText {
	v := &JSDocText{}
	v.reset()
	return v
}

type JSDocTypeLiteral struct {
	NodeBase
}

func (n *Node) AsJSDocTypeLiteral() *JSDocTypeLiteral { return n.data.(*JSDocTypeLiteral) }

func (n *JSDocTypeLiteral) reset() {
	*n = JSDocTypeLiteral{}
}

func (n *JSDocTypeLiteral) Kind() SyntaxKind { return SyntaxKindJSDocTypeLiteral }

func NewJSDocTypeLiteral() *JSDocTypeLiteral {
	v := &JSDocTypeLiteral{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocTypeLiteral() *JSDocTypeLiteral {
	v := &JSDocTypeLiteral{}
	v.reset()
	return v
}

type JSDocSignature struct {
	NodeBase
}

func (n *Node) AsJSDocSignature() *JSDocSignature { return n.data.(*JSDocSignature) }

func (n *JSDocSignature) reset() {
	*n = JSDocSignature{}
}

func (n *JSDocSignature) Kind() SyntaxKind { return SyntaxKindJSDocSignature }

func NewJSDocSignature() *JSDocSignature {
	v := &JSDocSignature{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocSignature() *JSDocSignature {
	v := &JSDocSignature{}
	v.reset()
	return v
}

type JSDocLink struct {
	NodeBase
}

func (n *Node) AsJSDocLink() *JSDocLink { return n.data.(*JSDocLink) }

func (n *JSDocLink) reset() {
	*n = JSDocLink{}
}

func (n *JSDocLink) Kind() SyntaxKind { return SyntaxKindJSDocLink }

func NewJSDocLink() *JSDocLink {
	v := &JSDocLink{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocLink() *JSDocLink {
	v := &JSDocLink{}
	v.reset()
	return v
}

type JSDocLinkCode struct {
	NodeBase
}

func (n *Node) AsJSDocLinkCode() *JSDocLinkCode { return n.data.(*JSDocLinkCode) }

func (n *JSDocLinkCode) reset() {
	*n = JSDocLinkCode{}
}

func (n *JSDocLinkCode) Kind() SyntaxKind { return SyntaxKindJSDocLinkCode }

func NewJSDocLinkCode() *JSDocLinkCode {
	v := &JSDocLinkCode{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocLinkCode() *JSDocLinkCode {
	v := &JSDocLinkCode{}
	v.reset()
	return v
}

type JSDocLinkPlain struct {
	NodeBase
}

func (n *Node) AsJSDocLinkPlain() *JSDocLinkPlain { return n.data.(*JSDocLinkPlain) }

func (n *JSDocLinkPlain) reset() {
	*n = JSDocLinkPlain{}
}

func (n *JSDocLinkPlain) Kind() SyntaxKind { return SyntaxKindJSDocLinkPlain }

func NewJSDocLinkPlain() *JSDocLinkPlain {
	v := &JSDocLinkPlain{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocLinkPlain() *JSDocLinkPlain {
	v := &JSDocLinkPlain{}
	v.reset()
	return v
}

type JSDocTag struct {
	NodeBase
}

func (n *Node) AsJSDocTag() *JSDocTag { return n.data.(*JSDocTag) }

func (n *JSDocTag) reset() {
	*n = JSDocTag{}
}

func (n *JSDocTag) Kind() SyntaxKind { return SyntaxKindJSDocTag }

func NewJSDocTag() *JSDocTag {
	v := &JSDocTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocTag() *JSDocTag {
	v := &JSDocTag{}
	v.reset()
	return v
}

type JSDocAugmentsTag struct {
	NodeBase
}

func (n *Node) AsJSDocAugmentsTag() *JSDocAugmentsTag { return n.data.(*JSDocAugmentsTag) }

func (n *JSDocAugmentsTag) reset() {
	*n = JSDocAugmentsTag{}
}

func (n *JSDocAugmentsTag) Kind() SyntaxKind { return SyntaxKindJSDocAugmentsTag }

func NewJSDocAugmentsTag() *JSDocAugmentsTag {
	v := &JSDocAugmentsTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocAugmentsTag() *JSDocAugmentsTag {
	v := &JSDocAugmentsTag{}
	v.reset()
	return v
}

type JSDocImplementsTag struct {
	NodeBase
}

func (n *Node) AsJSDocImplementsTag() *JSDocImplementsTag { return n.data.(*JSDocImplementsTag) }

func (n *JSDocImplementsTag) reset() {
	*n = JSDocImplementsTag{}
}

func (n *JSDocImplementsTag) Kind() SyntaxKind { return SyntaxKindJSDocImplementsTag }

func NewJSDocImplementsTag() *JSDocImplementsTag {
	v := &JSDocImplementsTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocImplementsTag() *JSDocImplementsTag {
	v := &JSDocImplementsTag{}
	v.reset()
	return v
}

type JSDocAuthorTag struct {
	NodeBase
}

func (n *Node) AsJSDocAuthorTag() *JSDocAuthorTag { return n.data.(*JSDocAuthorTag) }

func (n *JSDocAuthorTag) reset() {
	*n = JSDocAuthorTag{}
}

func (n *JSDocAuthorTag) Kind() SyntaxKind { return SyntaxKindJSDocAuthorTag }

func NewJSDocAuthorTag() *JSDocAuthorTag {
	v := &JSDocAuthorTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocAuthorTag() *JSDocAuthorTag {
	v := &JSDocAuthorTag{}
	v.reset()
	return v
}

type JSDocDeprecatedTag struct {
	NodeBase
}

func (n *Node) AsJSDocDeprecatedTag() *JSDocDeprecatedTag { return n.data.(*JSDocDeprecatedTag) }

func (n *JSDocDeprecatedTag) reset() {
	*n = JSDocDeprecatedTag{}
}

func (n *JSDocDeprecatedTag) Kind() SyntaxKind { return SyntaxKindJSDocDeprecatedTag }

func NewJSDocDeprecatedTag() *JSDocDeprecatedTag {
	v := &JSDocDeprecatedTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocDeprecatedTag() *JSDocDeprecatedTag {
	v := &JSDocDeprecatedTag{}
	v.reset()
	return v
}

type JSDocImmediateTag struct {
	NodeBase
}

func (n *Node) AsJSDocImmediateTag() *JSDocImmediateTag { return n.data.(*JSDocImmediateTag) }

func (n *JSDocImmediateTag) reset() {
	*n = JSDocImmediateTag{}
}

func (n *JSDocImmediateTag) Kind() SyntaxKind { return SyntaxKindJSDocImmediateTag }

func NewJSDocImmediateTag() *JSDocImmediateTag {
	v := &JSDocImmediateTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocImmediateTag() *JSDocImmediateTag {
	v := &JSDocImmediateTag{}
	v.reset()
	return v
}

type JSDocClassTag struct {
	NodeBase
}

func (n *Node) AsJSDocClassTag() *JSDocClassTag { return n.data.(*JSDocClassTag) }

func (n *JSDocClassTag) reset() {
	*n = JSDocClassTag{}
}

func (n *JSDocClassTag) Kind() SyntaxKind { return SyntaxKindJSDocClassTag }

func NewJSDocClassTag() *JSDocClassTag {
	v := &JSDocClassTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocClassTag() *JSDocClassTag {
	v := &JSDocClassTag{}
	v.reset()
	return v
}

type JSDocPublicTag struct {
	NodeBase
}

func (n *Node) AsJSDocPublicTag() *JSDocPublicTag { return n.data.(*JSDocPublicTag) }

func (n *JSDocPublicTag) reset() {
	*n = JSDocPublicTag{}
}

func (n *JSDocPublicTag) Kind() SyntaxKind { return SyntaxKindJSDocPublicTag }

func NewJSDocPublicTag() *JSDocPublicTag {
	v := &JSDocPublicTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocPublicTag() *JSDocPublicTag {
	v := &JSDocPublicTag{}
	v.reset()
	return v
}

type JSDocPrivateTag struct {
	NodeBase
}

func (n *Node) AsJSDocPrivateTag() *JSDocPrivateTag { return n.data.(*JSDocPrivateTag) }

func (n *JSDocPrivateTag) reset() {
	*n = JSDocPrivateTag{}
}

func (n *JSDocPrivateTag) Kind() SyntaxKind { return SyntaxKindJSDocPrivateTag }

func NewJSDocPrivateTag() *JSDocPrivateTag {
	v := &JSDocPrivateTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocPrivateTag() *JSDocPrivateTag {
	v := &JSDocPrivateTag{}
	v.reset()
	return v
}

type JSDocProtectedTag struct {
	NodeBase
}

func (n *Node) AsJSDocProtectedTag() *JSDocProtectedTag { return n.data.(*JSDocProtectedTag) }

func (n *JSDocProtectedTag) reset() {
	*n = JSDocProtectedTag{}
}

func (n *JSDocProtectedTag) Kind() SyntaxKind { return SyntaxKindJSDocProtectedTag }

func NewJSDocProtectedTag() *JSDocProtectedTag {
	v := &JSDocProtectedTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocProtectedTag() *JSDocProtectedTag {
	v := &JSDocProtectedTag{}
	v.reset()
	return v
}

type JSDocReadonlyTag struct {
	NodeBase
}

func (n *Node) AsJSDocReadonlyTag() *JSDocReadonlyTag { return n.data.(*JSDocReadonlyTag) }

func (n *JSDocReadonlyTag) reset() {
	*n = JSDocReadonlyTag{}
}

func (n *JSDocReadonlyTag) Kind() SyntaxKind { return SyntaxKindJSDocReadonlyTag }

func NewJSDocReadonlyTag() *JSDocReadonlyTag {
	v := &JSDocReadonlyTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocReadonlyTag() *JSDocReadonlyTag {
	v := &JSDocReadonlyTag{}
	v.reset()
	return v
}

type JSDocOverrideTag struct {
	NodeBase
}

func (n *Node) AsJSDocOverrideTag() *JSDocOverrideTag { return n.data.(*JSDocOverrideTag) }

func (n *JSDocOverrideTag) reset() {
	*n = JSDocOverrideTag{}
}

func (n *JSDocOverrideTag) Kind() SyntaxKind { return SyntaxKindJSDocOverrideTag }

func NewJSDocOverrideTag() *JSDocOverrideTag {
	v := &JSDocOverrideTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocOverrideTag() *JSDocOverrideTag {
	v := &JSDocOverrideTag{}
	v.reset()
	return v
}

type JSDocCallbackTag struct {
	NodeBase
}

func (n *Node) AsJSDocCallbackTag() *JSDocCallbackTag { return n.data.(*JSDocCallbackTag) }

func (n *JSDocCallbackTag) reset() {
	*n = JSDocCallbackTag{}
}

func (n *JSDocCallbackTag) Kind() SyntaxKind { return SyntaxKindJSDocCallbackTag }

func NewJSDocCallbackTag() *JSDocCallbackTag {
	v := &JSDocCallbackTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocCallbackTag() *JSDocCallbackTag {
	v := &JSDocCallbackTag{}
	v.reset()
	return v
}

type JSDocOverloadTag struct {
	NodeBase
}

func (n *Node) AsJSDocOverloadTag() *JSDocOverloadTag { return n.data.(*JSDocOverloadTag) }

func (n *JSDocOverloadTag) reset() {
	*n = JSDocOverloadTag{}
}

func (n *JSDocOverloadTag) Kind() SyntaxKind { return SyntaxKindJSDocOverloadTag }

func NewJSDocOverloadTag() *JSDocOverloadTag {
	v := &JSDocOverloadTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocOverloadTag() *JSDocOverloadTag {
	v := &JSDocOverloadTag{}
	v.reset()
	return v
}

type JSDocEnumTag struct {
	NodeBase
}

func (n *Node) AsJSDocEnumTag() *JSDocEnumTag { return n.data.(*JSDocEnumTag) }

func (n *JSDocEnumTag) reset() {
	*n = JSDocEnumTag{}
}

func (n *JSDocEnumTag) Kind() SyntaxKind { return SyntaxKindJSDocEnumTag }

func NewJSDocEnumTag() *JSDocEnumTag {
	v := &JSDocEnumTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocEnumTag() *JSDocEnumTag {
	v := &JSDocEnumTag{}
	v.reset()
	return v
}

type JSDocParameterTag struct {
	NodeBase
}

func (n *Node) AsJSDocParameterTag() *JSDocParameterTag { return n.data.(*JSDocParameterTag) }

func (n *JSDocParameterTag) reset() {
	*n = JSDocParameterTag{}
}

func (n *JSDocParameterTag) Kind() SyntaxKind { return SyntaxKindJSDocParameterTag }

func NewJSDocParameterTag() *JSDocParameterTag {
	v := &JSDocParameterTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocParameterTag() *JSDocParameterTag {
	v := &JSDocParameterTag{}
	v.reset()
	return v
}

type JSDocReturnTag struct {
	NodeBase
}

func (n *Node) AsJSDocReturnTag() *JSDocReturnTag { return n.data.(*JSDocReturnTag) }

func (n *JSDocReturnTag) reset() {
	*n = JSDocReturnTag{}
}

func (n *JSDocReturnTag) Kind() SyntaxKind { return SyntaxKindJSDocReturnTag }

func NewJSDocReturnTag() *JSDocReturnTag {
	v := &JSDocReturnTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocReturnTag() *JSDocReturnTag {
	v := &JSDocReturnTag{}
	v.reset()
	return v
}

type JSDocThisTag struct {
	NodeBase
}

func (n *Node) AsJSDocThisTag() *JSDocThisTag { return n.data.(*JSDocThisTag) }

func (n *JSDocThisTag) reset() {
	*n = JSDocThisTag{}
}

func (n *JSDocThisTag) Kind() SyntaxKind { return SyntaxKindJSDocThisTag }

func NewJSDocThisTag() *JSDocThisTag {
	v := &JSDocThisTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocThisTag() *JSDocThisTag {
	v := &JSDocThisTag{}
	v.reset()
	return v
}

type JSDocTypeTag struct {
	NodeBase
}

func (n *Node) AsJSDocTypeTag() *JSDocTypeTag { return n.data.(*JSDocTypeTag) }

func (n *JSDocTypeTag) reset() {
	*n = JSDocTypeTag{}
}

func (n *JSDocTypeTag) Kind() SyntaxKind { return SyntaxKindJSDocTypeTag }

func NewJSDocTypeTag() *JSDocTypeTag {
	v := &JSDocTypeTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocTypeTag() *JSDocTypeTag {
	v := &JSDocTypeTag{}
	v.reset()
	return v
}

type JSDocTemplateTag struct {
	NodeBase
}

func (n *Node) AsJSDocTemplateTag() *JSDocTemplateTag { return n.data.(*JSDocTemplateTag) }

func (n *JSDocTemplateTag) reset() {
	*n = JSDocTemplateTag{}
}

func (n *JSDocTemplateTag) Kind() SyntaxKind { return SyntaxKindJSDocTemplateTag }

func NewJSDocTemplateTag() *JSDocTemplateTag {
	v := &JSDocTemplateTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocTemplateTag() *JSDocTemplateTag {
	v := &JSDocTemplateTag{}
	v.reset()
	return v
}

type JSDocTypedefTag struct {
	NodeBase
}

func (n *Node) AsJSDocTypedefTag() *JSDocTypedefTag { return n.data.(*JSDocTypedefTag) }

func (n *JSDocTypedefTag) reset() {
	*n = JSDocTypedefTag{}
}

func (n *JSDocTypedefTag) Kind() SyntaxKind { return SyntaxKindJSDocTypedefTag }

func NewJSDocTypedefTag() *JSDocTypedefTag {
	v := &JSDocTypedefTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocTypedefTag() *JSDocTypedefTag {
	v := &JSDocTypedefTag{}
	v.reset()
	return v
}

type JSDocSeeTag struct {
	NodeBase
}

func (n *Node) AsJSDocSeeTag() *JSDocSeeTag { return n.data.(*JSDocSeeTag) }

func (n *JSDocSeeTag) reset() {
	*n = JSDocSeeTag{}
}

func (n *JSDocSeeTag) Kind() SyntaxKind { return SyntaxKindJSDocSeeTag }

func NewJSDocSeeTag() *JSDocSeeTag {
	v := &JSDocSeeTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocSeeTag() *JSDocSeeTag {
	v := &JSDocSeeTag{}
	v.reset()
	return v
}

type JSDocPropertyTag struct {
	NodeBase
}

func (n *Node) AsJSDocPropertyTag() *JSDocPropertyTag { return n.data.(*JSDocPropertyTag) }

func (n *JSDocPropertyTag) reset() {
	*n = JSDocPropertyTag{}
}

func (n *JSDocPropertyTag) Kind() SyntaxKind { return SyntaxKindJSDocPropertyTag }

func NewJSDocPropertyTag() *JSDocPropertyTag {
	v := &JSDocPropertyTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocPropertyTag() *JSDocPropertyTag {
	v := &JSDocPropertyTag{}
	v.reset()
	return v
}

type JSDocThrowsTag struct {
	NodeBase
}

func (n *Node) AsJSDocThrowsTag() *JSDocThrowsTag { return n.data.(*JSDocThrowsTag) }

func (n *JSDocThrowsTag) reset() {
	*n = JSDocThrowsTag{}
}

func (n *JSDocThrowsTag) Kind() SyntaxKind { return SyntaxKindJSDocThrowsTag }

func NewJSDocThrowsTag() *JSDocThrowsTag {
	v := &JSDocThrowsTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocThrowsTag() *JSDocThrowsTag {
	v := &JSDocThrowsTag{}
	v.reset()
	return v
}

type JSDocSatisfiesTag struct {
	NodeBase
}

func (n *Node) AsJSDocSatisfiesTag() *JSDocSatisfiesTag { return n.data.(*JSDocSatisfiesTag) }

func (n *JSDocSatisfiesTag) reset() {
	*n = JSDocSatisfiesTag{}
}

func (n *JSDocSatisfiesTag) Kind() SyntaxKind { return SyntaxKindJSDocSatisfiesTag }

func NewJSDocSatisfiesTag() *JSDocSatisfiesTag {
	v := &JSDocSatisfiesTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocSatisfiesTag() *JSDocSatisfiesTag {
	v := &JSDocSatisfiesTag{}
	v.reset()
	return v
}

type JSDocImportTag struct {
	NodeBase
}

func (n *Node) AsJSDocImportTag() *JSDocImportTag { return n.data.(*JSDocImportTag) }

func (n *JSDocImportTag) reset() {
	*n = JSDocImportTag{}
}

func (n *JSDocImportTag) Kind() SyntaxKind { return SyntaxKindJSDocImportTag }

func NewJSDocImportTag() *JSDocImportTag {
	v := &JSDocImportTag{}
	v.reset()
	return v
}

func (f *Factory) NewJSDocImportTag() *JSDocImportTag {
	v := &JSDocImportTag{}
	v.reset()
	return v
}

type SyntaxList struct {
	NodeBase
}

func (n *Node) AsSyntaxList() *SyntaxList { return n.data.(*SyntaxList) }

func (n *SyntaxList) reset() {
	*n = SyntaxList{}
}

func (n *SyntaxList) Kind() SyntaxKind { return SyntaxKindSyntaxList }

func NewSyntaxList() *SyntaxList {
	v := &SyntaxList{}
	v.reset()
	return v
}

func (f *Factory) NewSyntaxList() *SyntaxList {
	v := &SyntaxList{}
	v.reset()
	return v
}

type NotEmittedStatement struct {
	NodeBase
}

func (n *Node) AsNotEmittedStatement() *NotEmittedStatement { return n.data.(*NotEmittedStatement) }

func (n *NotEmittedStatement) reset() {
	*n = NotEmittedStatement{}
}

func (n *NotEmittedStatement) Kind() SyntaxKind { return SyntaxKindNotEmittedStatement }

func NewNotEmittedStatement() *NotEmittedStatement {
	v := &NotEmittedStatement{}
	v.reset()
	return v
}

func (f *Factory) NewNotEmittedStatement() *NotEmittedStatement {
	v := &NotEmittedStatement{}
	v.reset()
	return v
}

type PartiallyEmittedExpression struct {
	NodeBase
}

func (n *Node) AsPartiallyEmittedExpression() *PartiallyEmittedExpression {
	return n.data.(*PartiallyEmittedExpression)
}

func (n *PartiallyEmittedExpression) reset() {
	*n = PartiallyEmittedExpression{}
}

func (n *PartiallyEmittedExpression) Kind() SyntaxKind { return SyntaxKindPartiallyEmittedExpression }

func NewPartiallyEmittedExpression() *PartiallyEmittedExpression {
	v := &PartiallyEmittedExpression{}
	v.reset()
	return v
}

func (f *Factory) NewPartiallyEmittedExpression() *PartiallyEmittedExpression {
	v := &PartiallyEmittedExpression{}
	v.reset()
	return v
}

type CommaListExpression struct {
	NodeBase
}

func (n *Node) AsCommaListExpression() *CommaListExpression { return n.data.(*CommaListExpression) }

func (n *CommaListExpression) reset() {
	*n = CommaListExpression{}
}

func (n *CommaListExpression) Kind() SyntaxKind { return SyntaxKindCommaListExpression }

func NewCommaListExpression() *CommaListExpression {
	v := &CommaListExpression{}
	v.reset()
	return v
}

func (f *Factory) NewCommaListExpression() *CommaListExpression {
	v := &CommaListExpression{}
	v.reset()
	return v
}

type SyntheticReferenceExpression struct {
	NodeBase
}

func (n *Node) AsSyntheticReferenceExpression() *SyntheticReferenceExpression {
	return n.data.(*SyntheticReferenceExpression)
}

func (n *SyntheticReferenceExpression) reset() {
	*n = SyntheticReferenceExpression{}
}

func (n *SyntheticReferenceExpression) Kind() SyntaxKind {
	return SyntaxKindSyntheticReferenceExpression
}

func NewSyntheticReferenceExpression() *SyntheticReferenceExpression {
	v := &SyntheticReferenceExpression{}
	v.reset()
	return v
}

func (f *Factory) NewSyntheticReferenceExpression() *SyntheticReferenceExpression {
	v := &SyntheticReferenceExpression{}
	v.reset()
	return v
}

type AccessExpression = Node // PropertyAccessExpression | ElementAccessExpression

var isAccessExpressionTable = [SyntaxKindCount]bool{
	SyntaxKindPropertyAccessExpression: true,
	SyntaxKindElementAccessExpression:  true,
}

func isAccessExpressionKind(kind SyntaxKind) bool { return isAccessExpressionTable[kind] }
func isAccessExpression(n *Node) bool             { return isAccessExpressionTable[n.kind] }
func assertIsAccessExpression(n *Node) {
	if !isAccessExpression(n) {
		panic("expected AccessExpression, got " + n.kind.String())
	}
}
