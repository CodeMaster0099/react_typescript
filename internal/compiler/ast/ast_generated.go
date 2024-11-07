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
	pos int32
	end int32
}

func NewTextRange(pos, end int32) TextRange { return TextRange{pos: pos, end: end} }

func (r TextRange) Pos() int32                       { return r.pos }
func (r TextRange) End() int32                       { return r.end }
func (r TextRange) Len() int32                       { return r.end - r.pos }
func (r TextRange) ContainsInclusive(pos int32) bool { return r.pos <= pos && pos <= r.end }

type NodeData interface {
	AsNode() *Node
}

type Node struct {
	kind   SyntaxKind
	flags  NodeFlags
	loc    TextRange
	id     NodeID
	parent *Node
	data   NodeData
}

func (n *Node) Pos() int32 { return n.loc.Pos() }

type NodeDefault struct {
	Node
}

func (n *NodeDefault) AsNode() *Node { return &n.Node }

type NodeBase struct {
	NodeDefault
}

type Factory struct { // TODO
	poolToken      pool[Token]
	poolIdentifier pool[Identifier]
}

type Token struct {
	NodeBase
}

func (n *Node) AsToken() *Token { return n.data.(*Token) }

func (n *Token) set(kind SyntaxKind) {
	*n = Token{}
	n.kind = kind
	n.data = n
}

func NewToken(kind SyntaxKind) *Node {
	n := &Token{}
	n.set(kind)
	return n.AsNode()
}

func (f *Factory) NewToken(kind SyntaxKind) *Node {
	n := f.poolToken.allocate()
	n.set(kind)
	return n.AsNode()
}

type NumericLiteral struct {
	NodeBase
}

func (n *Node) AsNumericLiteral() *NumericLiteral { return n.data.(*NumericLiteral) }
func (n *Node) IsNumericLiteral() bool            { return n.kind == SyntaxKindNumericLiteral }

func (n *NumericLiteral) set() {
	*n = NumericLiteral{}
	n.kind = SyntaxKindNumericLiteral
	n.data = n
}

func (n *NumericLiteral) Kind() SyntaxKind { return SyntaxKindNumericLiteral }

func NewNumericLiteral() *Node {
	n := &NumericLiteral{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNumericLiteral() *Node {
	return NewNumericLiteral()
}

type BigIntLiteral struct {
	NodeBase
}

func (n *Node) AsBigIntLiteral() *BigIntLiteral { return n.data.(*BigIntLiteral) }
func (n *Node) IsBigIntLiteral() bool           { return n.kind == SyntaxKindBigIntLiteral }

func (n *BigIntLiteral) set() {
	*n = BigIntLiteral{}
	n.kind = SyntaxKindBigIntLiteral
	n.data = n
}

func (n *BigIntLiteral) Kind() SyntaxKind { return SyntaxKindBigIntLiteral }

func NewBigIntLiteral() *Node {
	n := &BigIntLiteral{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewBigIntLiteral() *Node {
	return NewBigIntLiteral()
}

type StringLiteral struct {
	NodeBase
}

func (n *Node) AsStringLiteral() *StringLiteral { return n.data.(*StringLiteral) }
func (n *Node) IsStringLiteral() bool           { return n.kind == SyntaxKindStringLiteral }

func (n *StringLiteral) set() {
	*n = StringLiteral{}
	n.kind = SyntaxKindStringLiteral
	n.data = n
}

func (n *StringLiteral) Kind() SyntaxKind { return SyntaxKindStringLiteral }

func NewStringLiteral() *Node {
	n := &StringLiteral{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewStringLiteral() *Node {
	return NewStringLiteral()
}

type JsxText struct {
	NodeBase
}

func (n *Node) AsJsxText() *JsxText { return n.data.(*JsxText) }
func (n *Node) IsJsxText() bool     { return n.kind == SyntaxKindJsxText }

func (n *JsxText) set() {
	*n = JsxText{}
	n.kind = SyntaxKindJsxText
	n.data = n
}

func (n *JsxText) Kind() SyntaxKind { return SyntaxKindJsxText }

func NewJsxText() *Node {
	n := &JsxText{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxText() *Node {
	return NewJsxText()
}

type RegularExpressionLiteral struct {
	NodeBase
}

func (n *Node) AsRegularExpressionLiteral() *RegularExpressionLiteral {
	return n.data.(*RegularExpressionLiteral)
}
func (n *Node) IsRegularExpressionLiteral() bool { return n.kind == SyntaxKindRegularExpressionLiteral }

func (n *RegularExpressionLiteral) set() {
	*n = RegularExpressionLiteral{}
	n.kind = SyntaxKindRegularExpressionLiteral
	n.data = n
}

func (n *RegularExpressionLiteral) Kind() SyntaxKind { return SyntaxKindRegularExpressionLiteral }

func NewRegularExpressionLiteral() *Node {
	n := &RegularExpressionLiteral{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewRegularExpressionLiteral() *Node {
	return NewRegularExpressionLiteral()
}

type NoSubstitutionTemplateLiteral struct {
	NodeBase
}

func (n *Node) AsNoSubstitutionTemplateLiteral() *NoSubstitutionTemplateLiteral {
	return n.data.(*NoSubstitutionTemplateLiteral)
}
func (n *Node) IsNoSubstitutionTemplateLiteral() bool {
	return n.kind == SyntaxKindNoSubstitutionTemplateLiteral
}

func (n *NoSubstitutionTemplateLiteral) set() {
	*n = NoSubstitutionTemplateLiteral{}
	n.kind = SyntaxKindNoSubstitutionTemplateLiteral
	n.data = n
}

func (n *NoSubstitutionTemplateLiteral) Kind() SyntaxKind {
	return SyntaxKindNoSubstitutionTemplateLiteral
}

func NewNoSubstitutionTemplateLiteral() *Node {
	n := &NoSubstitutionTemplateLiteral{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNoSubstitutionTemplateLiteral() *Node {
	return NewNoSubstitutionTemplateLiteral()
}

type TemplateHead struct {
	NodeBase
}

func (n *Node) AsTemplateHead() *TemplateHead { return n.data.(*TemplateHead) }
func (n *Node) IsTemplateHead() bool          { return n.kind == SyntaxKindTemplateHead }

func (n *TemplateHead) set() {
	*n = TemplateHead{}
	n.kind = SyntaxKindTemplateHead
	n.data = n
}

func (n *TemplateHead) Kind() SyntaxKind { return SyntaxKindTemplateHead }

func NewTemplateHead() *Node {
	n := &TemplateHead{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTemplateHead() *Node {
	return NewTemplateHead()
}

type TemplateMiddle struct {
	NodeBase
}

func (n *Node) AsTemplateMiddle() *TemplateMiddle { return n.data.(*TemplateMiddle) }
func (n *Node) IsTemplateMiddle() bool            { return n.kind == SyntaxKindTemplateMiddle }

func (n *TemplateMiddle) set() {
	*n = TemplateMiddle{}
	n.kind = SyntaxKindTemplateMiddle
	n.data = n
}

func (n *TemplateMiddle) Kind() SyntaxKind { return SyntaxKindTemplateMiddle }

func NewTemplateMiddle() *Node {
	n := &TemplateMiddle{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTemplateMiddle() *Node {
	return NewTemplateMiddle()
}

type TemplateTail struct {
	NodeBase
}

func (n *Node) AsTemplateTail() *TemplateTail { return n.data.(*TemplateTail) }
func (n *Node) IsTemplateTail() bool          { return n.kind == SyntaxKindTemplateTail }

func (n *TemplateTail) set() {
	*n = TemplateTail{}
	n.kind = SyntaxKindTemplateTail
	n.data = n
}

func (n *TemplateTail) Kind() SyntaxKind { return SyntaxKindTemplateTail }

func NewTemplateTail() *Node {
	n := &TemplateTail{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTemplateTail() *Node {
	return NewTemplateTail()
}

type Identifier struct {
	NodeBase
}

func (n *Node) AsIdentifier() *Identifier { return n.data.(*Identifier) }
func (n *Node) IsIdentifier() bool        { return n.kind == SyntaxKindIdentifier }

func (n *Identifier) set() {
	*n = Identifier{}
	n.kind = SyntaxKindIdentifier
	n.data = n
}

func (n *Identifier) Kind() SyntaxKind { return SyntaxKindIdentifier }

func NewIdentifier() *Node {
	n := &Identifier{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewIdentifier() *Node {
	n := f.poolIdentifier.allocate()
	n.set()
	return n.AsNode()
}

type PrivateIdentifier struct {
	NodeBase
}

func (n *Node) AsPrivateIdentifier() *PrivateIdentifier { return n.data.(*PrivateIdentifier) }
func (n *Node) IsPrivateIdentifier() bool               { return n.kind == SyntaxKindPrivateIdentifier }

func (n *PrivateIdentifier) set() {
	*n = PrivateIdentifier{}
	n.kind = SyntaxKindPrivateIdentifier
	n.data = n
}

func (n *PrivateIdentifier) Kind() SyntaxKind { return SyntaxKindPrivateIdentifier }

func NewPrivateIdentifier() *Node {
	n := &PrivateIdentifier{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPrivateIdentifier() *Node {
	return NewPrivateIdentifier()
}

type QualifiedName struct {
	NodeBase
}

func (n *Node) AsQualifiedName() *QualifiedName { return n.data.(*QualifiedName) }
func (n *Node) IsQualifiedName() bool           { return n.kind == SyntaxKindQualifiedName }

func (n *QualifiedName) set() {
	*n = QualifiedName{}
	n.kind = SyntaxKindQualifiedName
	n.data = n
}

func (n *QualifiedName) Kind() SyntaxKind { return SyntaxKindQualifiedName }

func NewQualifiedName() *Node {
	n := &QualifiedName{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewQualifiedName() *Node {
	return NewQualifiedName()
}

type ComputedPropertyName struct {
	NodeBase
}

func (n *Node) AsComputedPropertyName() *ComputedPropertyName { return n.data.(*ComputedPropertyName) }
func (n *Node) IsComputedPropertyName() bool                  { return n.kind == SyntaxKindComputedPropertyName }

func (n *ComputedPropertyName) set() {
	*n = ComputedPropertyName{}
	n.kind = SyntaxKindComputedPropertyName
	n.data = n
}

func (n *ComputedPropertyName) Kind() SyntaxKind { return SyntaxKindComputedPropertyName }

func NewComputedPropertyName() *Node {
	n := &ComputedPropertyName{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewComputedPropertyName() *Node {
	return NewComputedPropertyName()
}

type ModifierList struct {
	NodeBase
}

func (n *Node) AsModifierList() *ModifierList { return n.data.(*ModifierList) }
func (n *Node) IsModifierList() bool          { return n.kind == SyntaxKindModifierList }

func (n *ModifierList) set() {
	*n = ModifierList{}
	n.kind = SyntaxKindModifierList
	n.data = n
}

func (n *ModifierList) Kind() SyntaxKind { return SyntaxKindModifierList }

func NewModifierList() *Node {
	n := &ModifierList{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewModifierList() *Node {
	return NewModifierList()
}

type TypeParameterList struct {
	NodeBase
}

func (n *Node) AsTypeParameterList() *TypeParameterList { return n.data.(*TypeParameterList) }
func (n *Node) IsTypeParameterList() bool               { return n.kind == SyntaxKindTypeParameterList }

func (n *TypeParameterList) set() {
	*n = TypeParameterList{}
	n.kind = SyntaxKindTypeParameterList
	n.data = n
}

func (n *TypeParameterList) Kind() SyntaxKind { return SyntaxKindTypeParameterList }

func NewTypeParameterList() *Node {
	n := &TypeParameterList{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeParameterList() *Node {
	return NewTypeParameterList()
}

type TypeArgumentList struct {
	NodeBase
}

func (n *Node) AsTypeArgumentList() *TypeArgumentList { return n.data.(*TypeArgumentList) }
func (n *Node) IsTypeArgumentList() bool              { return n.kind == SyntaxKindTypeArgumentList }

func (n *TypeArgumentList) set() {
	*n = TypeArgumentList{}
	n.kind = SyntaxKindTypeArgumentList
	n.data = n
}

func (n *TypeArgumentList) Kind() SyntaxKind { return SyntaxKindTypeArgumentList }

func NewTypeArgumentList() *Node {
	n := &TypeArgumentList{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeArgumentList() *Node {
	return NewTypeArgumentList()
}

type TypeParameter struct {
	NodeBase
}

func (n *Node) AsTypeParameter() *TypeParameter { return n.data.(*TypeParameter) }
func (n *Node) IsTypeParameter() bool           { return n.kind == SyntaxKindTypeParameter }

func (n *TypeParameter) set() {
	*n = TypeParameter{}
	n.kind = SyntaxKindTypeParameter
	n.data = n
}

func (n *TypeParameter) Kind() SyntaxKind { return SyntaxKindTypeParameter }

func NewTypeParameter() *Node {
	n := &TypeParameter{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeParameter() *Node {
	return NewTypeParameter()
}

type Parameter struct {
	NodeBase
}

func (n *Node) AsParameter() *Parameter { return n.data.(*Parameter) }
func (n *Node) IsParameter() bool       { return n.kind == SyntaxKindParameter }

func (n *Parameter) set() {
	*n = Parameter{}
	n.kind = SyntaxKindParameter
	n.data = n
}

func (n *Parameter) Kind() SyntaxKind { return SyntaxKindParameter }

func NewParameter() *Node {
	n := &Parameter{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewParameter() *Node {
	return NewParameter()
}

type Decorator struct {
	NodeBase
}

func (n *Node) AsDecorator() *Decorator { return n.data.(*Decorator) }
func (n *Node) IsDecorator() bool       { return n.kind == SyntaxKindDecorator }

func (n *Decorator) set() {
	*n = Decorator{}
	n.kind = SyntaxKindDecorator
	n.data = n
}

func (n *Decorator) Kind() SyntaxKind { return SyntaxKindDecorator }

func NewDecorator() *Node {
	n := &Decorator{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewDecorator() *Node {
	return NewDecorator()
}

type PropertySignature struct {
	NodeBase
}

func (n *Node) AsPropertySignature() *PropertySignature { return n.data.(*PropertySignature) }
func (n *Node) IsPropertySignature() bool               { return n.kind == SyntaxKindPropertySignature }

func (n *PropertySignature) set() {
	*n = PropertySignature{}
	n.kind = SyntaxKindPropertySignature
	n.data = n
}

func (n *PropertySignature) Kind() SyntaxKind { return SyntaxKindPropertySignature }

func NewPropertySignature() *Node {
	n := &PropertySignature{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPropertySignature() *Node {
	return NewPropertySignature()
}

type PropertyDeclaration struct {
	NodeBase
}

func (n *Node) AsPropertyDeclaration() *PropertyDeclaration { return n.data.(*PropertyDeclaration) }
func (n *Node) IsPropertyDeclaration() bool                 { return n.kind == SyntaxKindPropertyDeclaration }

func (n *PropertyDeclaration) set() {
	*n = PropertyDeclaration{}
	n.kind = SyntaxKindPropertyDeclaration
	n.data = n
}

func (n *PropertyDeclaration) Kind() SyntaxKind { return SyntaxKindPropertyDeclaration }

func NewPropertyDeclaration() *Node {
	n := &PropertyDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPropertyDeclaration() *Node {
	return NewPropertyDeclaration()
}

type MethodSignature struct {
	NodeBase
}

func (n *Node) AsMethodSignature() *MethodSignature { return n.data.(*MethodSignature) }
func (n *Node) IsMethodSignature() bool             { return n.kind == SyntaxKindMethodSignature }

func (n *MethodSignature) set() {
	*n = MethodSignature{}
	n.kind = SyntaxKindMethodSignature
	n.data = n
}

func (n *MethodSignature) Kind() SyntaxKind { return SyntaxKindMethodSignature }

func NewMethodSignature() *Node {
	n := &MethodSignature{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewMethodSignature() *Node {
	return NewMethodSignature()
}

type MethodDeclaration struct {
	NodeBase
}

func (n *Node) AsMethodDeclaration() *MethodDeclaration { return n.data.(*MethodDeclaration) }
func (n *Node) IsMethodDeclaration() bool               { return n.kind == SyntaxKindMethodDeclaration }

func (n *MethodDeclaration) set() {
	*n = MethodDeclaration{}
	n.kind = SyntaxKindMethodDeclaration
	n.data = n
}

func (n *MethodDeclaration) Kind() SyntaxKind { return SyntaxKindMethodDeclaration }

func NewMethodDeclaration() *Node {
	n := &MethodDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewMethodDeclaration() *Node {
	return NewMethodDeclaration()
}

type ClassStaticBlockDeclaration struct {
	NodeBase
}

func (n *Node) AsClassStaticBlockDeclaration() *ClassStaticBlockDeclaration {
	return n.data.(*ClassStaticBlockDeclaration)
}
func (n *Node) IsClassStaticBlockDeclaration() bool {
	return n.kind == SyntaxKindClassStaticBlockDeclaration
}

func (n *ClassStaticBlockDeclaration) set() {
	*n = ClassStaticBlockDeclaration{}
	n.kind = SyntaxKindClassStaticBlockDeclaration
	n.data = n
}

func (n *ClassStaticBlockDeclaration) Kind() SyntaxKind { return SyntaxKindClassStaticBlockDeclaration }

func NewClassStaticBlockDeclaration() *Node {
	n := &ClassStaticBlockDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewClassStaticBlockDeclaration() *Node {
	return NewClassStaticBlockDeclaration()
}

type Constructor struct {
	NodeBase
}

func (n *Node) AsConstructor() *Constructor { return n.data.(*Constructor) }
func (n *Node) IsConstructor() bool         { return n.kind == SyntaxKindConstructor }

func (n *Constructor) set() {
	*n = Constructor{}
	n.kind = SyntaxKindConstructor
	n.data = n
}

func (n *Constructor) Kind() SyntaxKind { return SyntaxKindConstructor }

func NewConstructor() *Node {
	n := &Constructor{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewConstructor() *Node {
	return NewConstructor()
}

type GetAccessor struct {
	NodeBase
}

func (n *Node) AsGetAccessor() *GetAccessor { return n.data.(*GetAccessor) }
func (n *Node) IsGetAccessor() bool         { return n.kind == SyntaxKindGetAccessor }

func (n *GetAccessor) set() {
	*n = GetAccessor{}
	n.kind = SyntaxKindGetAccessor
	n.data = n
}

func (n *GetAccessor) Kind() SyntaxKind { return SyntaxKindGetAccessor }

func NewGetAccessor() *Node {
	n := &GetAccessor{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewGetAccessor() *Node {
	return NewGetAccessor()
}

type SetAccessor struct {
	NodeBase
}

func (n *Node) AsSetAccessor() *SetAccessor { return n.data.(*SetAccessor) }
func (n *Node) IsSetAccessor() bool         { return n.kind == SyntaxKindSetAccessor }

func (n *SetAccessor) set() {
	*n = SetAccessor{}
	n.kind = SyntaxKindSetAccessor
	n.data = n
}

func (n *SetAccessor) Kind() SyntaxKind { return SyntaxKindSetAccessor }

func NewSetAccessor() *Node {
	n := &SetAccessor{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSetAccessor() *Node {
	return NewSetAccessor()
}

type CallSignature struct {
	NodeBase
}

func (n *Node) AsCallSignature() *CallSignature { return n.data.(*CallSignature) }
func (n *Node) IsCallSignature() bool           { return n.kind == SyntaxKindCallSignature }

func (n *CallSignature) set() {
	*n = CallSignature{}
	n.kind = SyntaxKindCallSignature
	n.data = n
}

func (n *CallSignature) Kind() SyntaxKind { return SyntaxKindCallSignature }

func NewCallSignature() *Node {
	n := &CallSignature{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewCallSignature() *Node {
	return NewCallSignature()
}

type ConstructSignature struct {
	NodeBase
}

func (n *Node) AsConstructSignature() *ConstructSignature { return n.data.(*ConstructSignature) }
func (n *Node) IsConstructSignature() bool                { return n.kind == SyntaxKindConstructSignature }

func (n *ConstructSignature) set() {
	*n = ConstructSignature{}
	n.kind = SyntaxKindConstructSignature
	n.data = n
}

func (n *ConstructSignature) Kind() SyntaxKind { return SyntaxKindConstructSignature }

func NewConstructSignature() *Node {
	n := &ConstructSignature{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewConstructSignature() *Node {
	return NewConstructSignature()
}

type IndexSignature struct {
	NodeBase
}

func (n *Node) AsIndexSignature() *IndexSignature { return n.data.(*IndexSignature) }
func (n *Node) IsIndexSignature() bool            { return n.kind == SyntaxKindIndexSignature }

func (n *IndexSignature) set() {
	*n = IndexSignature{}
	n.kind = SyntaxKindIndexSignature
	n.data = n
}

func (n *IndexSignature) Kind() SyntaxKind { return SyntaxKindIndexSignature }

func NewIndexSignature() *Node {
	n := &IndexSignature{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewIndexSignature() *Node {
	return NewIndexSignature()
}

type TypePredicate struct {
	NodeBase
}

func (n *Node) AsTypePredicate() *TypePredicate { return n.data.(*TypePredicate) }
func (n *Node) IsTypePredicate() bool           { return n.kind == SyntaxKindTypePredicate }

func (n *TypePredicate) set() {
	*n = TypePredicate{}
	n.kind = SyntaxKindTypePredicate
	n.data = n
}

func (n *TypePredicate) Kind() SyntaxKind { return SyntaxKindTypePredicate }

func NewTypePredicate() *Node {
	n := &TypePredicate{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypePredicate() *Node {
	return NewTypePredicate()
}

type TypeReference struct {
	NodeBase
}

func (n *Node) AsTypeReference() *TypeReference { return n.data.(*TypeReference) }
func (n *Node) IsTypeReference() bool           { return n.kind == SyntaxKindTypeReference }

func (n *TypeReference) set() {
	*n = TypeReference{}
	n.kind = SyntaxKindTypeReference
	n.data = n
}

func (n *TypeReference) Kind() SyntaxKind { return SyntaxKindTypeReference }

func NewTypeReference() *Node {
	n := &TypeReference{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeReference() *Node {
	return NewTypeReference()
}

type FunctionType struct {
	NodeBase
}

func (n *Node) AsFunctionType() *FunctionType { return n.data.(*FunctionType) }
func (n *Node) IsFunctionType() bool          { return n.kind == SyntaxKindFunctionType }

func (n *FunctionType) set() {
	*n = FunctionType{}
	n.kind = SyntaxKindFunctionType
	n.data = n
}

func (n *FunctionType) Kind() SyntaxKind { return SyntaxKindFunctionType }

func NewFunctionType() *Node {
	n := &FunctionType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewFunctionType() *Node {
	return NewFunctionType()
}

type ConstructorType struct {
	NodeBase
}

func (n *Node) AsConstructorType() *ConstructorType { return n.data.(*ConstructorType) }
func (n *Node) IsConstructorType() bool             { return n.kind == SyntaxKindConstructorType }

func (n *ConstructorType) set() {
	*n = ConstructorType{}
	n.kind = SyntaxKindConstructorType
	n.data = n
}

func (n *ConstructorType) Kind() SyntaxKind { return SyntaxKindConstructorType }

func NewConstructorType() *Node {
	n := &ConstructorType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewConstructorType() *Node {
	return NewConstructorType()
}

type TypeQuery struct {
	NodeBase
}

func (n *Node) AsTypeQuery() *TypeQuery { return n.data.(*TypeQuery) }
func (n *Node) IsTypeQuery() bool       { return n.kind == SyntaxKindTypeQuery }

func (n *TypeQuery) set() {
	*n = TypeQuery{}
	n.kind = SyntaxKindTypeQuery
	n.data = n
}

func (n *TypeQuery) Kind() SyntaxKind { return SyntaxKindTypeQuery }

func NewTypeQuery() *Node {
	n := &TypeQuery{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeQuery() *Node {
	return NewTypeQuery()
}

type TypeLiteral struct {
	NodeBase
}

func (n *Node) AsTypeLiteral() *TypeLiteral { return n.data.(*TypeLiteral) }
func (n *Node) IsTypeLiteral() bool         { return n.kind == SyntaxKindTypeLiteral }

func (n *TypeLiteral) set() {
	*n = TypeLiteral{}
	n.kind = SyntaxKindTypeLiteral
	n.data = n
}

func (n *TypeLiteral) Kind() SyntaxKind { return SyntaxKindTypeLiteral }

func NewTypeLiteral() *Node {
	n := &TypeLiteral{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeLiteral() *Node {
	return NewTypeLiteral()
}

type ArrayType struct {
	NodeBase
}

func (n *Node) AsArrayType() *ArrayType { return n.data.(*ArrayType) }
func (n *Node) IsArrayType() bool       { return n.kind == SyntaxKindArrayType }

func (n *ArrayType) set() {
	*n = ArrayType{}
	n.kind = SyntaxKindArrayType
	n.data = n
}

func (n *ArrayType) Kind() SyntaxKind { return SyntaxKindArrayType }

func NewArrayType() *Node {
	n := &ArrayType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewArrayType() *Node {
	return NewArrayType()
}

type TupleType struct {
	NodeBase
}

func (n *Node) AsTupleType() *TupleType { return n.data.(*TupleType) }
func (n *Node) IsTupleType() bool       { return n.kind == SyntaxKindTupleType }

func (n *TupleType) set() {
	*n = TupleType{}
	n.kind = SyntaxKindTupleType
	n.data = n
}

func (n *TupleType) Kind() SyntaxKind { return SyntaxKindTupleType }

func NewTupleType() *Node {
	n := &TupleType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTupleType() *Node {
	return NewTupleType()
}

type OptionalType struct {
	NodeBase
}

func (n *Node) AsOptionalType() *OptionalType { return n.data.(*OptionalType) }
func (n *Node) IsOptionalType() bool          { return n.kind == SyntaxKindOptionalType }

func (n *OptionalType) set() {
	*n = OptionalType{}
	n.kind = SyntaxKindOptionalType
	n.data = n
}

func (n *OptionalType) Kind() SyntaxKind { return SyntaxKindOptionalType }

func NewOptionalType() *Node {
	n := &OptionalType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewOptionalType() *Node {
	return NewOptionalType()
}

type RestType struct {
	NodeBase
}

func (n *Node) AsRestType() *RestType { return n.data.(*RestType) }
func (n *Node) IsRestType() bool      { return n.kind == SyntaxKindRestType }

func (n *RestType) set() {
	*n = RestType{}
	n.kind = SyntaxKindRestType
	n.data = n
}

func (n *RestType) Kind() SyntaxKind { return SyntaxKindRestType }

func NewRestType() *Node {
	n := &RestType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewRestType() *Node {
	return NewRestType()
}

type UnionType struct {
	NodeBase
}

func (n *Node) AsUnionType() *UnionType { return n.data.(*UnionType) }
func (n *Node) IsUnionType() bool       { return n.kind == SyntaxKindUnionType }

func (n *UnionType) set() {
	*n = UnionType{}
	n.kind = SyntaxKindUnionType
	n.data = n
}

func (n *UnionType) Kind() SyntaxKind { return SyntaxKindUnionType }

func NewUnionType() *Node {
	n := &UnionType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewUnionType() *Node {
	return NewUnionType()
}

type IntersectionType struct {
	NodeBase
}

func (n *Node) AsIntersectionType() *IntersectionType { return n.data.(*IntersectionType) }
func (n *Node) IsIntersectionType() bool              { return n.kind == SyntaxKindIntersectionType }

func (n *IntersectionType) set() {
	*n = IntersectionType{}
	n.kind = SyntaxKindIntersectionType
	n.data = n
}

func (n *IntersectionType) Kind() SyntaxKind { return SyntaxKindIntersectionType }

func NewIntersectionType() *Node {
	n := &IntersectionType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewIntersectionType() *Node {
	return NewIntersectionType()
}

type ConditionalType struct {
	NodeBase
}

func (n *Node) AsConditionalType() *ConditionalType { return n.data.(*ConditionalType) }
func (n *Node) IsConditionalType() bool             { return n.kind == SyntaxKindConditionalType }

func (n *ConditionalType) set() {
	*n = ConditionalType{}
	n.kind = SyntaxKindConditionalType
	n.data = n
}

func (n *ConditionalType) Kind() SyntaxKind { return SyntaxKindConditionalType }

func NewConditionalType() *Node {
	n := &ConditionalType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewConditionalType() *Node {
	return NewConditionalType()
}

type InferType struct {
	NodeBase
}

func (n *Node) AsInferType() *InferType { return n.data.(*InferType) }
func (n *Node) IsInferType() bool       { return n.kind == SyntaxKindInferType }

func (n *InferType) set() {
	*n = InferType{}
	n.kind = SyntaxKindInferType
	n.data = n
}

func (n *InferType) Kind() SyntaxKind { return SyntaxKindInferType }

func NewInferType() *Node {
	n := &InferType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewInferType() *Node {
	return NewInferType()
}

type ParenthesizedType struct {
	NodeBase
}

func (n *Node) AsParenthesizedType() *ParenthesizedType { return n.data.(*ParenthesizedType) }
func (n *Node) IsParenthesizedType() bool               { return n.kind == SyntaxKindParenthesizedType }

func (n *ParenthesizedType) set() {
	*n = ParenthesizedType{}
	n.kind = SyntaxKindParenthesizedType
	n.data = n
}

func (n *ParenthesizedType) Kind() SyntaxKind { return SyntaxKindParenthesizedType }

func NewParenthesizedType() *Node {
	n := &ParenthesizedType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewParenthesizedType() *Node {
	return NewParenthesizedType()
}

type ThisType struct {
	NodeBase
}

func (n *Node) AsThisType() *ThisType { return n.data.(*ThisType) }
func (n *Node) IsThisType() bool      { return n.kind == SyntaxKindThisType }

func (n *ThisType) set() {
	*n = ThisType{}
	n.kind = SyntaxKindThisType
	n.data = n
}

func (n *ThisType) Kind() SyntaxKind { return SyntaxKindThisType }

func NewThisType() *Node {
	n := &ThisType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewThisType() *Node {
	return NewThisType()
}

type TypeOperator struct {
	NodeBase
}

func (n *Node) AsTypeOperator() *TypeOperator { return n.data.(*TypeOperator) }
func (n *Node) IsTypeOperator() bool          { return n.kind == SyntaxKindTypeOperator }

func (n *TypeOperator) set() {
	*n = TypeOperator{}
	n.kind = SyntaxKindTypeOperator
	n.data = n
}

func (n *TypeOperator) Kind() SyntaxKind { return SyntaxKindTypeOperator }

func NewTypeOperator() *Node {
	n := &TypeOperator{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeOperator() *Node {
	return NewTypeOperator()
}

type IndexedAccessType struct {
	NodeBase
}

func (n *Node) AsIndexedAccessType() *IndexedAccessType { return n.data.(*IndexedAccessType) }
func (n *Node) IsIndexedAccessType() bool               { return n.kind == SyntaxKindIndexedAccessType }

func (n *IndexedAccessType) set() {
	*n = IndexedAccessType{}
	n.kind = SyntaxKindIndexedAccessType
	n.data = n
}

func (n *IndexedAccessType) Kind() SyntaxKind { return SyntaxKindIndexedAccessType }

func NewIndexedAccessType() *Node {
	n := &IndexedAccessType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewIndexedAccessType() *Node {
	return NewIndexedAccessType()
}

type MappedType struct {
	NodeBase
}

func (n *Node) AsMappedType() *MappedType { return n.data.(*MappedType) }
func (n *Node) IsMappedType() bool        { return n.kind == SyntaxKindMappedType }

func (n *MappedType) set() {
	*n = MappedType{}
	n.kind = SyntaxKindMappedType
	n.data = n
}

func (n *MappedType) Kind() SyntaxKind { return SyntaxKindMappedType }

func NewMappedType() *Node {
	n := &MappedType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewMappedType() *Node {
	return NewMappedType()
}

type LiteralType struct {
	NodeBase
}

func (n *Node) AsLiteralType() *LiteralType { return n.data.(*LiteralType) }
func (n *Node) IsLiteralType() bool         { return n.kind == SyntaxKindLiteralType }

func (n *LiteralType) set() {
	*n = LiteralType{}
	n.kind = SyntaxKindLiteralType
	n.data = n
}

func (n *LiteralType) Kind() SyntaxKind { return SyntaxKindLiteralType }

func NewLiteralType() *Node {
	n := &LiteralType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewLiteralType() *Node {
	return NewLiteralType()
}

type NamedTupleMember struct {
	NodeBase
}

func (n *Node) AsNamedTupleMember() *NamedTupleMember { return n.data.(*NamedTupleMember) }
func (n *Node) IsNamedTupleMember() bool              { return n.kind == SyntaxKindNamedTupleMember }

func (n *NamedTupleMember) set() {
	*n = NamedTupleMember{}
	n.kind = SyntaxKindNamedTupleMember
	n.data = n
}

func (n *NamedTupleMember) Kind() SyntaxKind { return SyntaxKindNamedTupleMember }

func NewNamedTupleMember() *Node {
	n := &NamedTupleMember{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNamedTupleMember() *Node {
	return NewNamedTupleMember()
}

type TemplateLiteralType struct {
	NodeBase
}

func (n *Node) AsTemplateLiteralType() *TemplateLiteralType { return n.data.(*TemplateLiteralType) }
func (n *Node) IsTemplateLiteralType() bool                 { return n.kind == SyntaxKindTemplateLiteralType }

func (n *TemplateLiteralType) set() {
	*n = TemplateLiteralType{}
	n.kind = SyntaxKindTemplateLiteralType
	n.data = n
}

func (n *TemplateLiteralType) Kind() SyntaxKind { return SyntaxKindTemplateLiteralType }

func NewTemplateLiteralType() *Node {
	n := &TemplateLiteralType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTemplateLiteralType() *Node {
	return NewTemplateLiteralType()
}

type TemplateLiteralTypeSpan struct {
	NodeBase
}

func (n *Node) AsTemplateLiteralTypeSpan() *TemplateLiteralTypeSpan {
	return n.data.(*TemplateLiteralTypeSpan)
}
func (n *Node) IsTemplateLiteralTypeSpan() bool { return n.kind == SyntaxKindTemplateLiteralTypeSpan }

func (n *TemplateLiteralTypeSpan) set() {
	*n = TemplateLiteralTypeSpan{}
	n.kind = SyntaxKindTemplateLiteralTypeSpan
	n.data = n
}

func (n *TemplateLiteralTypeSpan) Kind() SyntaxKind { return SyntaxKindTemplateLiteralTypeSpan }

func NewTemplateLiteralTypeSpan() *Node {
	n := &TemplateLiteralTypeSpan{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTemplateLiteralTypeSpan() *Node {
	return NewTemplateLiteralTypeSpan()
}

type ImportType struct {
	NodeBase
}

func (n *Node) AsImportType() *ImportType { return n.data.(*ImportType) }
func (n *Node) IsImportType() bool        { return n.kind == SyntaxKindImportType }

func (n *ImportType) set() {
	*n = ImportType{}
	n.kind = SyntaxKindImportType
	n.data = n
}

func (n *ImportType) Kind() SyntaxKind { return SyntaxKindImportType }

func NewImportType() *Node {
	n := &ImportType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewImportType() *Node {
	return NewImportType()
}

type ObjectBindingPattern struct {
	NodeBase
}

func (n *Node) AsObjectBindingPattern() *ObjectBindingPattern { return n.data.(*ObjectBindingPattern) }
func (n *Node) IsObjectBindingPattern() bool                  { return n.kind == SyntaxKindObjectBindingPattern }

func (n *ObjectBindingPattern) set() {
	*n = ObjectBindingPattern{}
	n.kind = SyntaxKindObjectBindingPattern
	n.data = n
}

func (n *ObjectBindingPattern) Kind() SyntaxKind { return SyntaxKindObjectBindingPattern }

func NewObjectBindingPattern() *Node {
	n := &ObjectBindingPattern{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewObjectBindingPattern() *Node {
	return NewObjectBindingPattern()
}

type ArrayBindingPattern struct {
	NodeBase
}

func (n *Node) AsArrayBindingPattern() *ArrayBindingPattern { return n.data.(*ArrayBindingPattern) }
func (n *Node) IsArrayBindingPattern() bool                 { return n.kind == SyntaxKindArrayBindingPattern }

func (n *ArrayBindingPattern) set() {
	*n = ArrayBindingPattern{}
	n.kind = SyntaxKindArrayBindingPattern
	n.data = n
}

func (n *ArrayBindingPattern) Kind() SyntaxKind { return SyntaxKindArrayBindingPattern }

func NewArrayBindingPattern() *Node {
	n := &ArrayBindingPattern{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewArrayBindingPattern() *Node {
	return NewArrayBindingPattern()
}

type BindingElement struct {
	NodeBase
}

func (n *Node) AsBindingElement() *BindingElement { return n.data.(*BindingElement) }
func (n *Node) IsBindingElement() bool            { return n.kind == SyntaxKindBindingElement }

func (n *BindingElement) set() {
	*n = BindingElement{}
	n.kind = SyntaxKindBindingElement
	n.data = n
}

func (n *BindingElement) Kind() SyntaxKind { return SyntaxKindBindingElement }

func NewBindingElement() *Node {
	n := &BindingElement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewBindingElement() *Node {
	return NewBindingElement()
}

type ArrayLiteralExpression struct {
	NodeBase
}

func (n *Node) AsArrayLiteralExpression() *ArrayLiteralExpression {
	return n.data.(*ArrayLiteralExpression)
}
func (n *Node) IsArrayLiteralExpression() bool { return n.kind == SyntaxKindArrayLiteralExpression }

func (n *ArrayLiteralExpression) set() {
	*n = ArrayLiteralExpression{}
	n.kind = SyntaxKindArrayLiteralExpression
	n.data = n
}

func (n *ArrayLiteralExpression) Kind() SyntaxKind { return SyntaxKindArrayLiteralExpression }

func NewArrayLiteralExpression() *Node {
	n := &ArrayLiteralExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewArrayLiteralExpression() *Node {
	return NewArrayLiteralExpression()
}

type ObjectLiteralExpression struct {
	NodeBase
}

func (n *Node) AsObjectLiteralExpression() *ObjectLiteralExpression {
	return n.data.(*ObjectLiteralExpression)
}
func (n *Node) IsObjectLiteralExpression() bool { return n.kind == SyntaxKindObjectLiteralExpression }

func (n *ObjectLiteralExpression) set() {
	*n = ObjectLiteralExpression{}
	n.kind = SyntaxKindObjectLiteralExpression
	n.data = n
}

func (n *ObjectLiteralExpression) Kind() SyntaxKind { return SyntaxKindObjectLiteralExpression }

func NewObjectLiteralExpression() *Node {
	n := &ObjectLiteralExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewObjectLiteralExpression() *Node {
	return NewObjectLiteralExpression()
}

type PropertyAccessExpression struct {
	NodeBase
}

func (n *Node) AsPropertyAccessExpression() *PropertyAccessExpression {
	return n.data.(*PropertyAccessExpression)
}
func (n *Node) IsPropertyAccessExpression() bool { return n.kind == SyntaxKindPropertyAccessExpression }

func (n *PropertyAccessExpression) set() {
	*n = PropertyAccessExpression{}
	n.kind = SyntaxKindPropertyAccessExpression
	n.data = n
}

func (n *PropertyAccessExpression) Kind() SyntaxKind { return SyntaxKindPropertyAccessExpression }

func NewPropertyAccessExpression() *Node {
	n := &PropertyAccessExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPropertyAccessExpression() *Node {
	return NewPropertyAccessExpression()
}

type ElementAccessExpression struct {
	NodeBase
}

func (n *Node) AsElementAccessExpression() *ElementAccessExpression {
	return n.data.(*ElementAccessExpression)
}
func (n *Node) IsElementAccessExpression() bool { return n.kind == SyntaxKindElementAccessExpression }

func (n *ElementAccessExpression) set() {
	*n = ElementAccessExpression{}
	n.kind = SyntaxKindElementAccessExpression
	n.data = n
}

func (n *ElementAccessExpression) Kind() SyntaxKind { return SyntaxKindElementAccessExpression }

func NewElementAccessExpression() *Node {
	n := &ElementAccessExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewElementAccessExpression() *Node {
	return NewElementAccessExpression()
}

type CallExpression struct {
	NodeBase
}

func (n *Node) AsCallExpression() *CallExpression { return n.data.(*CallExpression) }
func (n *Node) IsCallExpression() bool            { return n.kind == SyntaxKindCallExpression }

func (n *CallExpression) set() {
	*n = CallExpression{}
	n.kind = SyntaxKindCallExpression
	n.data = n
}

func (n *CallExpression) Kind() SyntaxKind { return SyntaxKindCallExpression }

func NewCallExpression() *Node {
	n := &CallExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewCallExpression() *Node {
	return NewCallExpression()
}

type NewExpression struct {
	NodeBase
}

func (n *Node) AsNewExpression() *NewExpression { return n.data.(*NewExpression) }
func (n *Node) IsNewExpression() bool           { return n.kind == SyntaxKindNewExpression }

func (n *NewExpression) set() {
	*n = NewExpression{}
	n.kind = SyntaxKindNewExpression
	n.data = n
}

func (n *NewExpression) Kind() SyntaxKind { return SyntaxKindNewExpression }

func NewNewExpression() *Node {
	n := &NewExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNewExpression() *Node {
	return NewNewExpression()
}

type TaggedTemplateExpression struct {
	NodeBase
}

func (n *Node) AsTaggedTemplateExpression() *TaggedTemplateExpression {
	return n.data.(*TaggedTemplateExpression)
}
func (n *Node) IsTaggedTemplateExpression() bool { return n.kind == SyntaxKindTaggedTemplateExpression }

func (n *TaggedTemplateExpression) set() {
	*n = TaggedTemplateExpression{}
	n.kind = SyntaxKindTaggedTemplateExpression
	n.data = n
}

func (n *TaggedTemplateExpression) Kind() SyntaxKind { return SyntaxKindTaggedTemplateExpression }

func NewTaggedTemplateExpression() *Node {
	n := &TaggedTemplateExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTaggedTemplateExpression() *Node {
	return NewTaggedTemplateExpression()
}

type TypeAssertionExpression struct {
	NodeBase
}

func (n *Node) AsTypeAssertionExpression() *TypeAssertionExpression {
	return n.data.(*TypeAssertionExpression)
}
func (n *Node) IsTypeAssertionExpression() bool { return n.kind == SyntaxKindTypeAssertionExpression }

func (n *TypeAssertionExpression) set() {
	*n = TypeAssertionExpression{}
	n.kind = SyntaxKindTypeAssertionExpression
	n.data = n
}

func (n *TypeAssertionExpression) Kind() SyntaxKind { return SyntaxKindTypeAssertionExpression }

func NewTypeAssertionExpression() *Node {
	n := &TypeAssertionExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeAssertionExpression() *Node {
	return NewTypeAssertionExpression()
}

type ParenthesizedExpression struct {
	NodeBase
}

func (n *Node) AsParenthesizedExpression() *ParenthesizedExpression {
	return n.data.(*ParenthesizedExpression)
}
func (n *Node) IsParenthesizedExpression() bool { return n.kind == SyntaxKindParenthesizedExpression }

func (n *ParenthesizedExpression) set() {
	*n = ParenthesizedExpression{}
	n.kind = SyntaxKindParenthesizedExpression
	n.data = n
}

func (n *ParenthesizedExpression) Kind() SyntaxKind { return SyntaxKindParenthesizedExpression }

func NewParenthesizedExpression() *Node {
	n := &ParenthesizedExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewParenthesizedExpression() *Node {
	return NewParenthesizedExpression()
}

type FunctionExpression struct {
	NodeBase
}

func (n *Node) AsFunctionExpression() *FunctionExpression { return n.data.(*FunctionExpression) }
func (n *Node) IsFunctionExpression() bool                { return n.kind == SyntaxKindFunctionExpression }

func (n *FunctionExpression) set() {
	*n = FunctionExpression{}
	n.kind = SyntaxKindFunctionExpression
	n.data = n
}

func (n *FunctionExpression) Kind() SyntaxKind { return SyntaxKindFunctionExpression }

func NewFunctionExpression() *Node {
	n := &FunctionExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewFunctionExpression() *Node {
	return NewFunctionExpression()
}

type ArrowFunction struct {
	NodeBase
}

func (n *Node) AsArrowFunction() *ArrowFunction { return n.data.(*ArrowFunction) }
func (n *Node) IsArrowFunction() bool           { return n.kind == SyntaxKindArrowFunction }

func (n *ArrowFunction) set() {
	*n = ArrowFunction{}
	n.kind = SyntaxKindArrowFunction
	n.data = n
}

func (n *ArrowFunction) Kind() SyntaxKind { return SyntaxKindArrowFunction }

func NewArrowFunction() *Node {
	n := &ArrowFunction{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewArrowFunction() *Node {
	return NewArrowFunction()
}

type DeleteExpression struct {
	NodeBase
}

func (n *Node) AsDeleteExpression() *DeleteExpression { return n.data.(*DeleteExpression) }
func (n *Node) IsDeleteExpression() bool              { return n.kind == SyntaxKindDeleteExpression }

func (n *DeleteExpression) set() {
	*n = DeleteExpression{}
	n.kind = SyntaxKindDeleteExpression
	n.data = n
}

func (n *DeleteExpression) Kind() SyntaxKind { return SyntaxKindDeleteExpression }

func NewDeleteExpression() *Node {
	n := &DeleteExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewDeleteExpression() *Node {
	return NewDeleteExpression()
}

type TypeOfExpression struct {
	NodeBase
}

func (n *Node) AsTypeOfExpression() *TypeOfExpression { return n.data.(*TypeOfExpression) }
func (n *Node) IsTypeOfExpression() bool              { return n.kind == SyntaxKindTypeOfExpression }

func (n *TypeOfExpression) set() {
	*n = TypeOfExpression{}
	n.kind = SyntaxKindTypeOfExpression
	n.data = n
}

func (n *TypeOfExpression) Kind() SyntaxKind { return SyntaxKindTypeOfExpression }

func NewTypeOfExpression() *Node {
	n := &TypeOfExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeOfExpression() *Node {
	return NewTypeOfExpression()
}

type VoidExpression struct {
	NodeBase
}

func (n *Node) AsVoidExpression() *VoidExpression { return n.data.(*VoidExpression) }
func (n *Node) IsVoidExpression() bool            { return n.kind == SyntaxKindVoidExpression }

func (n *VoidExpression) set() {
	*n = VoidExpression{}
	n.kind = SyntaxKindVoidExpression
	n.data = n
}

func (n *VoidExpression) Kind() SyntaxKind { return SyntaxKindVoidExpression }

func NewVoidExpression() *Node {
	n := &VoidExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewVoidExpression() *Node {
	return NewVoidExpression()
}

type AwaitExpression struct {
	NodeBase
}

func (n *Node) AsAwaitExpression() *AwaitExpression { return n.data.(*AwaitExpression) }
func (n *Node) IsAwaitExpression() bool             { return n.kind == SyntaxKindAwaitExpression }

func (n *AwaitExpression) set() {
	*n = AwaitExpression{}
	n.kind = SyntaxKindAwaitExpression
	n.data = n
}

func (n *AwaitExpression) Kind() SyntaxKind { return SyntaxKindAwaitExpression }

func NewAwaitExpression() *Node {
	n := &AwaitExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewAwaitExpression() *Node {
	return NewAwaitExpression()
}

type PrefixUnaryExpression struct {
	NodeBase
}

func (n *Node) AsPrefixUnaryExpression() *PrefixUnaryExpression {
	return n.data.(*PrefixUnaryExpression)
}
func (n *Node) IsPrefixUnaryExpression() bool { return n.kind == SyntaxKindPrefixUnaryExpression }

func (n *PrefixUnaryExpression) set() {
	*n = PrefixUnaryExpression{}
	n.kind = SyntaxKindPrefixUnaryExpression
	n.data = n
}

func (n *PrefixUnaryExpression) Kind() SyntaxKind { return SyntaxKindPrefixUnaryExpression }

func NewPrefixUnaryExpression() *Node {
	n := &PrefixUnaryExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPrefixUnaryExpression() *Node {
	return NewPrefixUnaryExpression()
}

type PostfixUnaryExpression struct {
	NodeBase
}

func (n *Node) AsPostfixUnaryExpression() *PostfixUnaryExpression {
	return n.data.(*PostfixUnaryExpression)
}
func (n *Node) IsPostfixUnaryExpression() bool { return n.kind == SyntaxKindPostfixUnaryExpression }

func (n *PostfixUnaryExpression) set() {
	*n = PostfixUnaryExpression{}
	n.kind = SyntaxKindPostfixUnaryExpression
	n.data = n
}

func (n *PostfixUnaryExpression) Kind() SyntaxKind { return SyntaxKindPostfixUnaryExpression }

func NewPostfixUnaryExpression() *Node {
	n := &PostfixUnaryExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPostfixUnaryExpression() *Node {
	return NewPostfixUnaryExpression()
}

type BinaryExpression struct {
	NodeBase
}

func (n *Node) AsBinaryExpression() *BinaryExpression { return n.data.(*BinaryExpression) }
func (n *Node) IsBinaryExpression() bool              { return n.kind == SyntaxKindBinaryExpression }

func (n *BinaryExpression) set() {
	*n = BinaryExpression{}
	n.kind = SyntaxKindBinaryExpression
	n.data = n
}

func (n *BinaryExpression) Kind() SyntaxKind { return SyntaxKindBinaryExpression }

func NewBinaryExpression() *Node {
	n := &BinaryExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewBinaryExpression() *Node {
	return NewBinaryExpression()
}

type ConditionalExpression struct {
	NodeBase
}

func (n *Node) AsConditionalExpression() *ConditionalExpression {
	return n.data.(*ConditionalExpression)
}
func (n *Node) IsConditionalExpression() bool { return n.kind == SyntaxKindConditionalExpression }

func (n *ConditionalExpression) set() {
	*n = ConditionalExpression{}
	n.kind = SyntaxKindConditionalExpression
	n.data = n
}

func (n *ConditionalExpression) Kind() SyntaxKind { return SyntaxKindConditionalExpression }

func NewConditionalExpression() *Node {
	n := &ConditionalExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewConditionalExpression() *Node {
	return NewConditionalExpression()
}

type TemplateExpression struct {
	NodeBase
}

func (n *Node) AsTemplateExpression() *TemplateExpression { return n.data.(*TemplateExpression) }
func (n *Node) IsTemplateExpression() bool                { return n.kind == SyntaxKindTemplateExpression }

func (n *TemplateExpression) set() {
	*n = TemplateExpression{}
	n.kind = SyntaxKindTemplateExpression
	n.data = n
}

func (n *TemplateExpression) Kind() SyntaxKind { return SyntaxKindTemplateExpression }

func NewTemplateExpression() *Node {
	n := &TemplateExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTemplateExpression() *Node {
	return NewTemplateExpression()
}

type YieldExpression struct {
	NodeBase
}

func (n *Node) AsYieldExpression() *YieldExpression { return n.data.(*YieldExpression) }
func (n *Node) IsYieldExpression() bool             { return n.kind == SyntaxKindYieldExpression }

func (n *YieldExpression) set() {
	*n = YieldExpression{}
	n.kind = SyntaxKindYieldExpression
	n.data = n
}

func (n *YieldExpression) Kind() SyntaxKind { return SyntaxKindYieldExpression }

func NewYieldExpression() *Node {
	n := &YieldExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewYieldExpression() *Node {
	return NewYieldExpression()
}

type SpreadElement struct {
	NodeBase
}

func (n *Node) AsSpreadElement() *SpreadElement { return n.data.(*SpreadElement) }
func (n *Node) IsSpreadElement() bool           { return n.kind == SyntaxKindSpreadElement }

func (n *SpreadElement) set() {
	*n = SpreadElement{}
	n.kind = SyntaxKindSpreadElement
	n.data = n
}

func (n *SpreadElement) Kind() SyntaxKind { return SyntaxKindSpreadElement }

func NewSpreadElement() *Node {
	n := &SpreadElement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSpreadElement() *Node {
	return NewSpreadElement()
}

type ClassExpression struct {
	NodeBase
}

func (n *Node) AsClassExpression() *ClassExpression { return n.data.(*ClassExpression) }
func (n *Node) IsClassExpression() bool             { return n.kind == SyntaxKindClassExpression }

func (n *ClassExpression) set() {
	*n = ClassExpression{}
	n.kind = SyntaxKindClassExpression
	n.data = n
}

func (n *ClassExpression) Kind() SyntaxKind { return SyntaxKindClassExpression }

func NewClassExpression() *Node {
	n := &ClassExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewClassExpression() *Node {
	return NewClassExpression()
}

type OmittedExpression struct {
	NodeBase
}

func (n *Node) AsOmittedExpression() *OmittedExpression { return n.data.(*OmittedExpression) }
func (n *Node) IsOmittedExpression() bool               { return n.kind == SyntaxKindOmittedExpression }

func (n *OmittedExpression) set() {
	*n = OmittedExpression{}
	n.kind = SyntaxKindOmittedExpression
	n.data = n
}

func (n *OmittedExpression) Kind() SyntaxKind { return SyntaxKindOmittedExpression }

func NewOmittedExpression() *Node {
	n := &OmittedExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewOmittedExpression() *Node {
	return NewOmittedExpression()
}

type ExpressionWithTypeArguments struct {
	NodeBase
}

func (n *Node) AsExpressionWithTypeArguments() *ExpressionWithTypeArguments {
	return n.data.(*ExpressionWithTypeArguments)
}
func (n *Node) IsExpressionWithTypeArguments() bool {
	return n.kind == SyntaxKindExpressionWithTypeArguments
}

func (n *ExpressionWithTypeArguments) set() {
	*n = ExpressionWithTypeArguments{}
	n.kind = SyntaxKindExpressionWithTypeArguments
	n.data = n
}

func (n *ExpressionWithTypeArguments) Kind() SyntaxKind { return SyntaxKindExpressionWithTypeArguments }

func NewExpressionWithTypeArguments() *Node {
	n := &ExpressionWithTypeArguments{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewExpressionWithTypeArguments() *Node {
	return NewExpressionWithTypeArguments()
}

type AsExpression struct {
	NodeBase
}

func (n *Node) AsAsExpression() *AsExpression { return n.data.(*AsExpression) }
func (n *Node) IsAsExpression() bool          { return n.kind == SyntaxKindAsExpression }

func (n *AsExpression) set() {
	*n = AsExpression{}
	n.kind = SyntaxKindAsExpression
	n.data = n
}

func (n *AsExpression) Kind() SyntaxKind { return SyntaxKindAsExpression }

func NewAsExpression() *Node {
	n := &AsExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewAsExpression() *Node {
	return NewAsExpression()
}

type NonNullExpression struct {
	NodeBase
}

func (n *Node) AsNonNullExpression() *NonNullExpression { return n.data.(*NonNullExpression) }
func (n *Node) IsNonNullExpression() bool               { return n.kind == SyntaxKindNonNullExpression }

func (n *NonNullExpression) set() {
	*n = NonNullExpression{}
	n.kind = SyntaxKindNonNullExpression
	n.data = n
}

func (n *NonNullExpression) Kind() SyntaxKind { return SyntaxKindNonNullExpression }

func NewNonNullExpression() *Node {
	n := &NonNullExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNonNullExpression() *Node {
	return NewNonNullExpression()
}

type MetaProperty struct {
	NodeBase
}

func (n *Node) AsMetaProperty() *MetaProperty { return n.data.(*MetaProperty) }
func (n *Node) IsMetaProperty() bool          { return n.kind == SyntaxKindMetaProperty }

func (n *MetaProperty) set() {
	*n = MetaProperty{}
	n.kind = SyntaxKindMetaProperty
	n.data = n
}

func (n *MetaProperty) Kind() SyntaxKind { return SyntaxKindMetaProperty }

func NewMetaProperty() *Node {
	n := &MetaProperty{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewMetaProperty() *Node {
	return NewMetaProperty()
}

type SyntheticExpression struct {
	NodeBase
}

func (n *Node) AsSyntheticExpression() *SyntheticExpression { return n.data.(*SyntheticExpression) }
func (n *Node) IsSyntheticExpression() bool                 { return n.kind == SyntaxKindSyntheticExpression }

func (n *SyntheticExpression) set() {
	*n = SyntheticExpression{}
	n.kind = SyntaxKindSyntheticExpression
	n.data = n
}

func (n *SyntheticExpression) Kind() SyntaxKind { return SyntaxKindSyntheticExpression }

func NewSyntheticExpression() *Node {
	n := &SyntheticExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSyntheticExpression() *Node {
	return NewSyntheticExpression()
}

type SatisfiesExpression struct {
	NodeBase
}

func (n *Node) AsSatisfiesExpression() *SatisfiesExpression { return n.data.(*SatisfiesExpression) }
func (n *Node) IsSatisfiesExpression() bool                 { return n.kind == SyntaxKindSatisfiesExpression }

func (n *SatisfiesExpression) set() {
	*n = SatisfiesExpression{}
	n.kind = SyntaxKindSatisfiesExpression
	n.data = n
}

func (n *SatisfiesExpression) Kind() SyntaxKind { return SyntaxKindSatisfiesExpression }

func NewSatisfiesExpression() *Node {
	n := &SatisfiesExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSatisfiesExpression() *Node {
	return NewSatisfiesExpression()
}

type TemplateSpan struct {
	NodeBase
}

func (n *Node) AsTemplateSpan() *TemplateSpan { return n.data.(*TemplateSpan) }
func (n *Node) IsTemplateSpan() bool          { return n.kind == SyntaxKindTemplateSpan }

func (n *TemplateSpan) set() {
	*n = TemplateSpan{}
	n.kind = SyntaxKindTemplateSpan
	n.data = n
}

func (n *TemplateSpan) Kind() SyntaxKind { return SyntaxKindTemplateSpan }

func NewTemplateSpan() *Node {
	n := &TemplateSpan{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTemplateSpan() *Node {
	return NewTemplateSpan()
}

type SemicolonClassElement struct {
	NodeBase
}

func (n *Node) AsSemicolonClassElement() *SemicolonClassElement {
	return n.data.(*SemicolonClassElement)
}
func (n *Node) IsSemicolonClassElement() bool { return n.kind == SyntaxKindSemicolonClassElement }

func (n *SemicolonClassElement) set() {
	*n = SemicolonClassElement{}
	n.kind = SyntaxKindSemicolonClassElement
	n.data = n
}

func (n *SemicolonClassElement) Kind() SyntaxKind { return SyntaxKindSemicolonClassElement }

func NewSemicolonClassElement() *Node {
	n := &SemicolonClassElement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSemicolonClassElement() *Node {
	return NewSemicolonClassElement()
}

type Block struct {
	NodeBase
}

func (n *Node) AsBlock() *Block { return n.data.(*Block) }
func (n *Node) IsBlock() bool   { return n.kind == SyntaxKindBlock }

func (n *Block) set() {
	*n = Block{}
	n.kind = SyntaxKindBlock
	n.data = n
}

func (n *Block) Kind() SyntaxKind { return SyntaxKindBlock }

func NewBlock() *Node {
	n := &Block{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewBlock() *Node {
	return NewBlock()
}

type EmptyStatement struct {
	NodeBase
}

func (n *Node) AsEmptyStatement() *EmptyStatement { return n.data.(*EmptyStatement) }
func (n *Node) IsEmptyStatement() bool            { return n.kind == SyntaxKindEmptyStatement }

func (n *EmptyStatement) set() {
	*n = EmptyStatement{}
	n.kind = SyntaxKindEmptyStatement
	n.data = n
}

func (n *EmptyStatement) Kind() SyntaxKind { return SyntaxKindEmptyStatement }

func NewEmptyStatement() *Node {
	n := &EmptyStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewEmptyStatement() *Node {
	return NewEmptyStatement()
}

type VariableStatement struct {
	NodeBase
}

func (n *Node) AsVariableStatement() *VariableStatement { return n.data.(*VariableStatement) }
func (n *Node) IsVariableStatement() bool               { return n.kind == SyntaxKindVariableStatement }

func (n *VariableStatement) set() {
	*n = VariableStatement{}
	n.kind = SyntaxKindVariableStatement
	n.data = n
}

func (n *VariableStatement) Kind() SyntaxKind { return SyntaxKindVariableStatement }

func NewVariableStatement() *Node {
	n := &VariableStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewVariableStatement() *Node {
	return NewVariableStatement()
}

type ExpressionStatement struct {
	NodeBase
}

func (n *Node) AsExpressionStatement() *ExpressionStatement { return n.data.(*ExpressionStatement) }
func (n *Node) IsExpressionStatement() bool                 { return n.kind == SyntaxKindExpressionStatement }

func (n *ExpressionStatement) set() {
	*n = ExpressionStatement{}
	n.kind = SyntaxKindExpressionStatement
	n.data = n
}

func (n *ExpressionStatement) Kind() SyntaxKind { return SyntaxKindExpressionStatement }

func NewExpressionStatement() *Node {
	n := &ExpressionStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewExpressionStatement() *Node {
	return NewExpressionStatement()
}

type IfStatement struct {
	NodeBase
}

func (n *Node) AsIfStatement() *IfStatement { return n.data.(*IfStatement) }
func (n *Node) IsIfStatement() bool         { return n.kind == SyntaxKindIfStatement }

func (n *IfStatement) set() {
	*n = IfStatement{}
	n.kind = SyntaxKindIfStatement
	n.data = n
}

func (n *IfStatement) Kind() SyntaxKind { return SyntaxKindIfStatement }

func NewIfStatement() *Node {
	n := &IfStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewIfStatement() *Node {
	return NewIfStatement()
}

type DoStatement struct {
	NodeBase
}

func (n *Node) AsDoStatement() *DoStatement { return n.data.(*DoStatement) }
func (n *Node) IsDoStatement() bool         { return n.kind == SyntaxKindDoStatement }

func (n *DoStatement) set() {
	*n = DoStatement{}
	n.kind = SyntaxKindDoStatement
	n.data = n
}

func (n *DoStatement) Kind() SyntaxKind { return SyntaxKindDoStatement }

func NewDoStatement() *Node {
	n := &DoStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewDoStatement() *Node {
	return NewDoStatement()
}

type WhileStatement struct {
	NodeBase
}

func (n *Node) AsWhileStatement() *WhileStatement { return n.data.(*WhileStatement) }
func (n *Node) IsWhileStatement() bool            { return n.kind == SyntaxKindWhileStatement }

func (n *WhileStatement) set() {
	*n = WhileStatement{}
	n.kind = SyntaxKindWhileStatement
	n.data = n
}

func (n *WhileStatement) Kind() SyntaxKind { return SyntaxKindWhileStatement }

func NewWhileStatement() *Node {
	n := &WhileStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewWhileStatement() *Node {
	return NewWhileStatement()
}

type ForStatement struct {
	NodeBase
}

func (n *Node) AsForStatement() *ForStatement { return n.data.(*ForStatement) }
func (n *Node) IsForStatement() bool          { return n.kind == SyntaxKindForStatement }

func (n *ForStatement) set() {
	*n = ForStatement{}
	n.kind = SyntaxKindForStatement
	n.data = n
}

func (n *ForStatement) Kind() SyntaxKind { return SyntaxKindForStatement }

func NewForStatement() *Node {
	n := &ForStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewForStatement() *Node {
	return NewForStatement()
}

type ForInStatement struct {
	NodeBase
}

func (n *Node) AsForInStatement() *ForInStatement { return n.data.(*ForInStatement) }
func (n *Node) IsForInStatement() bool            { return n.kind == SyntaxKindForInStatement }

func (n *ForInStatement) set() {
	*n = ForInStatement{}
	n.kind = SyntaxKindForInStatement
	n.data = n
}

func (n *ForInStatement) Kind() SyntaxKind { return SyntaxKindForInStatement }

func NewForInStatement() *Node {
	n := &ForInStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewForInStatement() *Node {
	return NewForInStatement()
}

type ForOfStatement struct {
	NodeBase
}

func (n *Node) AsForOfStatement() *ForOfStatement { return n.data.(*ForOfStatement) }
func (n *Node) IsForOfStatement() bool            { return n.kind == SyntaxKindForOfStatement }

func (n *ForOfStatement) set() {
	*n = ForOfStatement{}
	n.kind = SyntaxKindForOfStatement
	n.data = n
}

func (n *ForOfStatement) Kind() SyntaxKind { return SyntaxKindForOfStatement }

func NewForOfStatement() *Node {
	n := &ForOfStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewForOfStatement() *Node {
	return NewForOfStatement()
}

type ContinueStatement struct {
	NodeBase
}

func (n *Node) AsContinueStatement() *ContinueStatement { return n.data.(*ContinueStatement) }
func (n *Node) IsContinueStatement() bool               { return n.kind == SyntaxKindContinueStatement }

func (n *ContinueStatement) set() {
	*n = ContinueStatement{}
	n.kind = SyntaxKindContinueStatement
	n.data = n
}

func (n *ContinueStatement) Kind() SyntaxKind { return SyntaxKindContinueStatement }

func NewContinueStatement() *Node {
	n := &ContinueStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewContinueStatement() *Node {
	return NewContinueStatement()
}

type BreakStatement struct {
	NodeBase
}

func (n *Node) AsBreakStatement() *BreakStatement { return n.data.(*BreakStatement) }
func (n *Node) IsBreakStatement() bool            { return n.kind == SyntaxKindBreakStatement }

func (n *BreakStatement) set() {
	*n = BreakStatement{}
	n.kind = SyntaxKindBreakStatement
	n.data = n
}

func (n *BreakStatement) Kind() SyntaxKind { return SyntaxKindBreakStatement }

func NewBreakStatement() *Node {
	n := &BreakStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewBreakStatement() *Node {
	return NewBreakStatement()
}

type ReturnStatement struct {
	NodeBase
}

func (n *Node) AsReturnStatement() *ReturnStatement { return n.data.(*ReturnStatement) }
func (n *Node) IsReturnStatement() bool             { return n.kind == SyntaxKindReturnStatement }

func (n *ReturnStatement) set() {
	*n = ReturnStatement{}
	n.kind = SyntaxKindReturnStatement
	n.data = n
}

func (n *ReturnStatement) Kind() SyntaxKind { return SyntaxKindReturnStatement }

func NewReturnStatement() *Node {
	n := &ReturnStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewReturnStatement() *Node {
	return NewReturnStatement()
}

type WithStatement struct {
	NodeBase
}

func (n *Node) AsWithStatement() *WithStatement { return n.data.(*WithStatement) }
func (n *Node) IsWithStatement() bool           { return n.kind == SyntaxKindWithStatement }

func (n *WithStatement) set() {
	*n = WithStatement{}
	n.kind = SyntaxKindWithStatement
	n.data = n
}

func (n *WithStatement) Kind() SyntaxKind { return SyntaxKindWithStatement }

func NewWithStatement() *Node {
	n := &WithStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewWithStatement() *Node {
	return NewWithStatement()
}

type SwitchStatement struct {
	NodeBase
}

func (n *Node) AsSwitchStatement() *SwitchStatement { return n.data.(*SwitchStatement) }
func (n *Node) IsSwitchStatement() bool             { return n.kind == SyntaxKindSwitchStatement }

func (n *SwitchStatement) set() {
	*n = SwitchStatement{}
	n.kind = SyntaxKindSwitchStatement
	n.data = n
}

func (n *SwitchStatement) Kind() SyntaxKind { return SyntaxKindSwitchStatement }

func NewSwitchStatement() *Node {
	n := &SwitchStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSwitchStatement() *Node {
	return NewSwitchStatement()
}

type LabeledStatement struct {
	NodeBase
}

func (n *Node) AsLabeledStatement() *LabeledStatement { return n.data.(*LabeledStatement) }
func (n *Node) IsLabeledStatement() bool              { return n.kind == SyntaxKindLabeledStatement }

func (n *LabeledStatement) set() {
	*n = LabeledStatement{}
	n.kind = SyntaxKindLabeledStatement
	n.data = n
}

func (n *LabeledStatement) Kind() SyntaxKind { return SyntaxKindLabeledStatement }

func NewLabeledStatement() *Node {
	n := &LabeledStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewLabeledStatement() *Node {
	return NewLabeledStatement()
}

type ThrowStatement struct {
	NodeBase
}

func (n *Node) AsThrowStatement() *ThrowStatement { return n.data.(*ThrowStatement) }
func (n *Node) IsThrowStatement() bool            { return n.kind == SyntaxKindThrowStatement }

func (n *ThrowStatement) set() {
	*n = ThrowStatement{}
	n.kind = SyntaxKindThrowStatement
	n.data = n
}

func (n *ThrowStatement) Kind() SyntaxKind { return SyntaxKindThrowStatement }

func NewThrowStatement() *Node {
	n := &ThrowStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewThrowStatement() *Node {
	return NewThrowStatement()
}

type TryStatement struct {
	NodeBase
}

func (n *Node) AsTryStatement() *TryStatement { return n.data.(*TryStatement) }
func (n *Node) IsTryStatement() bool          { return n.kind == SyntaxKindTryStatement }

func (n *TryStatement) set() {
	*n = TryStatement{}
	n.kind = SyntaxKindTryStatement
	n.data = n
}

func (n *TryStatement) Kind() SyntaxKind { return SyntaxKindTryStatement }

func NewTryStatement() *Node {
	n := &TryStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTryStatement() *Node {
	return NewTryStatement()
}

type DebuggerStatement struct {
	NodeBase
}

func (n *Node) AsDebuggerStatement() *DebuggerStatement { return n.data.(*DebuggerStatement) }
func (n *Node) IsDebuggerStatement() bool               { return n.kind == SyntaxKindDebuggerStatement }

func (n *DebuggerStatement) set() {
	*n = DebuggerStatement{}
	n.kind = SyntaxKindDebuggerStatement
	n.data = n
}

func (n *DebuggerStatement) Kind() SyntaxKind { return SyntaxKindDebuggerStatement }

func NewDebuggerStatement() *Node {
	n := &DebuggerStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewDebuggerStatement() *Node {
	return NewDebuggerStatement()
}

type VariableDeclaration struct {
	NodeBase
}

func (n *Node) AsVariableDeclaration() *VariableDeclaration { return n.data.(*VariableDeclaration) }
func (n *Node) IsVariableDeclaration() bool                 { return n.kind == SyntaxKindVariableDeclaration }

func (n *VariableDeclaration) set() {
	*n = VariableDeclaration{}
	n.kind = SyntaxKindVariableDeclaration
	n.data = n
}

func (n *VariableDeclaration) Kind() SyntaxKind { return SyntaxKindVariableDeclaration }

func NewVariableDeclaration() *Node {
	n := &VariableDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewVariableDeclaration() *Node {
	return NewVariableDeclaration()
}

type VariableDeclarationList struct {
	NodeBase
}

func (n *Node) AsVariableDeclarationList() *VariableDeclarationList {
	return n.data.(*VariableDeclarationList)
}
func (n *Node) IsVariableDeclarationList() bool { return n.kind == SyntaxKindVariableDeclarationList }

func (n *VariableDeclarationList) set() {
	*n = VariableDeclarationList{}
	n.kind = SyntaxKindVariableDeclarationList
	n.data = n
}

func (n *VariableDeclarationList) Kind() SyntaxKind { return SyntaxKindVariableDeclarationList }

func NewVariableDeclarationList() *Node {
	n := &VariableDeclarationList{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewVariableDeclarationList() *Node {
	return NewVariableDeclarationList()
}

type FunctionDeclaration struct {
	NodeBase
}

func (n *Node) AsFunctionDeclaration() *FunctionDeclaration { return n.data.(*FunctionDeclaration) }
func (n *Node) IsFunctionDeclaration() bool                 { return n.kind == SyntaxKindFunctionDeclaration }

func (n *FunctionDeclaration) set() {
	*n = FunctionDeclaration{}
	n.kind = SyntaxKindFunctionDeclaration
	n.data = n
}

func (n *FunctionDeclaration) Kind() SyntaxKind { return SyntaxKindFunctionDeclaration }

func NewFunctionDeclaration() *Node {
	n := &FunctionDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewFunctionDeclaration() *Node {
	return NewFunctionDeclaration()
}

type ClassDeclaration struct {
	NodeBase
}

func (n *Node) AsClassDeclaration() *ClassDeclaration { return n.data.(*ClassDeclaration) }
func (n *Node) IsClassDeclaration() bool              { return n.kind == SyntaxKindClassDeclaration }

func (n *ClassDeclaration) set() {
	*n = ClassDeclaration{}
	n.kind = SyntaxKindClassDeclaration
	n.data = n
}

func (n *ClassDeclaration) Kind() SyntaxKind { return SyntaxKindClassDeclaration }

func NewClassDeclaration() *Node {
	n := &ClassDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewClassDeclaration() *Node {
	return NewClassDeclaration()
}

type InterfaceDeclaration struct {
	NodeBase
}

func (n *Node) AsInterfaceDeclaration() *InterfaceDeclaration { return n.data.(*InterfaceDeclaration) }
func (n *Node) IsInterfaceDeclaration() bool                  { return n.kind == SyntaxKindInterfaceDeclaration }

func (n *InterfaceDeclaration) set() {
	*n = InterfaceDeclaration{}
	n.kind = SyntaxKindInterfaceDeclaration
	n.data = n
}

func (n *InterfaceDeclaration) Kind() SyntaxKind { return SyntaxKindInterfaceDeclaration }

func NewInterfaceDeclaration() *Node {
	n := &InterfaceDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewInterfaceDeclaration() *Node {
	return NewInterfaceDeclaration()
}

type TypeAliasDeclaration struct {
	NodeBase
}

func (n *Node) AsTypeAliasDeclaration() *TypeAliasDeclaration { return n.data.(*TypeAliasDeclaration) }
func (n *Node) IsTypeAliasDeclaration() bool                  { return n.kind == SyntaxKindTypeAliasDeclaration }

func (n *TypeAliasDeclaration) set() {
	*n = TypeAliasDeclaration{}
	n.kind = SyntaxKindTypeAliasDeclaration
	n.data = n
}

func (n *TypeAliasDeclaration) Kind() SyntaxKind { return SyntaxKindTypeAliasDeclaration }

func NewTypeAliasDeclaration() *Node {
	n := &TypeAliasDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewTypeAliasDeclaration() *Node {
	return NewTypeAliasDeclaration()
}

type EnumDeclaration struct {
	NodeBase
}

func (n *Node) AsEnumDeclaration() *EnumDeclaration { return n.data.(*EnumDeclaration) }
func (n *Node) IsEnumDeclaration() bool             { return n.kind == SyntaxKindEnumDeclaration }

func (n *EnumDeclaration) set() {
	*n = EnumDeclaration{}
	n.kind = SyntaxKindEnumDeclaration
	n.data = n
}

func (n *EnumDeclaration) Kind() SyntaxKind { return SyntaxKindEnumDeclaration }

func NewEnumDeclaration() *Node {
	n := &EnumDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewEnumDeclaration() *Node {
	return NewEnumDeclaration()
}

type ModuleDeclaration struct {
	NodeBase
}

func (n *Node) AsModuleDeclaration() *ModuleDeclaration { return n.data.(*ModuleDeclaration) }
func (n *Node) IsModuleDeclaration() bool               { return n.kind == SyntaxKindModuleDeclaration }

func (n *ModuleDeclaration) set() {
	*n = ModuleDeclaration{}
	n.kind = SyntaxKindModuleDeclaration
	n.data = n
}

func (n *ModuleDeclaration) Kind() SyntaxKind { return SyntaxKindModuleDeclaration }

func NewModuleDeclaration() *Node {
	n := &ModuleDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewModuleDeclaration() *Node {
	return NewModuleDeclaration()
}

type ModuleBlock struct {
	NodeBase
}

func (n *Node) AsModuleBlock() *ModuleBlock { return n.data.(*ModuleBlock) }
func (n *Node) IsModuleBlock() bool         { return n.kind == SyntaxKindModuleBlock }

func (n *ModuleBlock) set() {
	*n = ModuleBlock{}
	n.kind = SyntaxKindModuleBlock
	n.data = n
}

func (n *ModuleBlock) Kind() SyntaxKind { return SyntaxKindModuleBlock }

func NewModuleBlock() *Node {
	n := &ModuleBlock{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewModuleBlock() *Node {
	return NewModuleBlock()
}

type CaseBlock struct {
	NodeBase
}

func (n *Node) AsCaseBlock() *CaseBlock { return n.data.(*CaseBlock) }
func (n *Node) IsCaseBlock() bool       { return n.kind == SyntaxKindCaseBlock }

func (n *CaseBlock) set() {
	*n = CaseBlock{}
	n.kind = SyntaxKindCaseBlock
	n.data = n
}

func (n *CaseBlock) Kind() SyntaxKind { return SyntaxKindCaseBlock }

func NewCaseBlock() *Node {
	n := &CaseBlock{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewCaseBlock() *Node {
	return NewCaseBlock()
}

type NamespaceExportDeclaration struct {
	NodeBase
}

func (n *Node) AsNamespaceExportDeclaration() *NamespaceExportDeclaration {
	return n.data.(*NamespaceExportDeclaration)
}
func (n *Node) IsNamespaceExportDeclaration() bool {
	return n.kind == SyntaxKindNamespaceExportDeclaration
}

func (n *NamespaceExportDeclaration) set() {
	*n = NamespaceExportDeclaration{}
	n.kind = SyntaxKindNamespaceExportDeclaration
	n.data = n
}

func (n *NamespaceExportDeclaration) Kind() SyntaxKind { return SyntaxKindNamespaceExportDeclaration }

func NewNamespaceExportDeclaration() *Node {
	n := &NamespaceExportDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNamespaceExportDeclaration() *Node {
	return NewNamespaceExportDeclaration()
}

type ImportEqualsDeclaration struct {
	NodeBase
}

func (n *Node) AsImportEqualsDeclaration() *ImportEqualsDeclaration {
	return n.data.(*ImportEqualsDeclaration)
}
func (n *Node) IsImportEqualsDeclaration() bool { return n.kind == SyntaxKindImportEqualsDeclaration }

func (n *ImportEqualsDeclaration) set() {
	*n = ImportEqualsDeclaration{}
	n.kind = SyntaxKindImportEqualsDeclaration
	n.data = n
}

func (n *ImportEqualsDeclaration) Kind() SyntaxKind { return SyntaxKindImportEqualsDeclaration }

func NewImportEqualsDeclaration() *Node {
	n := &ImportEqualsDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewImportEqualsDeclaration() *Node {
	return NewImportEqualsDeclaration()
}

type ImportDeclaration struct {
	NodeBase
}

func (n *Node) AsImportDeclaration() *ImportDeclaration { return n.data.(*ImportDeclaration) }
func (n *Node) IsImportDeclaration() bool               { return n.kind == SyntaxKindImportDeclaration }

func (n *ImportDeclaration) set() {
	*n = ImportDeclaration{}
	n.kind = SyntaxKindImportDeclaration
	n.data = n
}

func (n *ImportDeclaration) Kind() SyntaxKind { return SyntaxKindImportDeclaration }

func NewImportDeclaration() *Node {
	n := &ImportDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewImportDeclaration() *Node {
	return NewImportDeclaration()
}

type ImportClause struct {
	NodeBase
}

func (n *Node) AsImportClause() *ImportClause { return n.data.(*ImportClause) }
func (n *Node) IsImportClause() bool          { return n.kind == SyntaxKindImportClause }

func (n *ImportClause) set() {
	*n = ImportClause{}
	n.kind = SyntaxKindImportClause
	n.data = n
}

func (n *ImportClause) Kind() SyntaxKind { return SyntaxKindImportClause }

func NewImportClause() *Node {
	n := &ImportClause{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewImportClause() *Node {
	return NewImportClause()
}

type NamespaceImport struct {
	NodeBase
}

func (n *Node) AsNamespaceImport() *NamespaceImport { return n.data.(*NamespaceImport) }
func (n *Node) IsNamespaceImport() bool             { return n.kind == SyntaxKindNamespaceImport }

func (n *NamespaceImport) set() {
	*n = NamespaceImport{}
	n.kind = SyntaxKindNamespaceImport
	n.data = n
}

func (n *NamespaceImport) Kind() SyntaxKind { return SyntaxKindNamespaceImport }

func NewNamespaceImport() *Node {
	n := &NamespaceImport{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNamespaceImport() *Node {
	return NewNamespaceImport()
}

type NamedImports struct {
	NodeBase
}

func (n *Node) AsNamedImports() *NamedImports { return n.data.(*NamedImports) }
func (n *Node) IsNamedImports() bool          { return n.kind == SyntaxKindNamedImports }

func (n *NamedImports) set() {
	*n = NamedImports{}
	n.kind = SyntaxKindNamedImports
	n.data = n
}

func (n *NamedImports) Kind() SyntaxKind { return SyntaxKindNamedImports }

func NewNamedImports() *Node {
	n := &NamedImports{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNamedImports() *Node {
	return NewNamedImports()
}

type ImportSpecifier struct {
	NodeBase
}

func (n *Node) AsImportSpecifier() *ImportSpecifier { return n.data.(*ImportSpecifier) }
func (n *Node) IsImportSpecifier() bool             { return n.kind == SyntaxKindImportSpecifier }

func (n *ImportSpecifier) set() {
	*n = ImportSpecifier{}
	n.kind = SyntaxKindImportSpecifier
	n.data = n
}

func (n *ImportSpecifier) Kind() SyntaxKind { return SyntaxKindImportSpecifier }

func NewImportSpecifier() *Node {
	n := &ImportSpecifier{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewImportSpecifier() *Node {
	return NewImportSpecifier()
}

type ExportAssignment struct {
	NodeBase
}

func (n *Node) AsExportAssignment() *ExportAssignment { return n.data.(*ExportAssignment) }
func (n *Node) IsExportAssignment() bool              { return n.kind == SyntaxKindExportAssignment }

func (n *ExportAssignment) set() {
	*n = ExportAssignment{}
	n.kind = SyntaxKindExportAssignment
	n.data = n
}

func (n *ExportAssignment) Kind() SyntaxKind { return SyntaxKindExportAssignment }

func NewExportAssignment() *Node {
	n := &ExportAssignment{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewExportAssignment() *Node {
	return NewExportAssignment()
}

type ExportDeclaration struct {
	NodeBase
}

func (n *Node) AsExportDeclaration() *ExportDeclaration { return n.data.(*ExportDeclaration) }
func (n *Node) IsExportDeclaration() bool               { return n.kind == SyntaxKindExportDeclaration }

func (n *ExportDeclaration) set() {
	*n = ExportDeclaration{}
	n.kind = SyntaxKindExportDeclaration
	n.data = n
}

func (n *ExportDeclaration) Kind() SyntaxKind { return SyntaxKindExportDeclaration }

func NewExportDeclaration() *Node {
	n := &ExportDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewExportDeclaration() *Node {
	return NewExportDeclaration()
}

type NamedExports struct {
	NodeBase
}

func (n *Node) AsNamedExports() *NamedExports { return n.data.(*NamedExports) }
func (n *Node) IsNamedExports() bool          { return n.kind == SyntaxKindNamedExports }

func (n *NamedExports) set() {
	*n = NamedExports{}
	n.kind = SyntaxKindNamedExports
	n.data = n
}

func (n *NamedExports) Kind() SyntaxKind { return SyntaxKindNamedExports }

func NewNamedExports() *Node {
	n := &NamedExports{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNamedExports() *Node {
	return NewNamedExports()
}

type NamespaceExport struct {
	NodeBase
}

func (n *Node) AsNamespaceExport() *NamespaceExport { return n.data.(*NamespaceExport) }
func (n *Node) IsNamespaceExport() bool             { return n.kind == SyntaxKindNamespaceExport }

func (n *NamespaceExport) set() {
	*n = NamespaceExport{}
	n.kind = SyntaxKindNamespaceExport
	n.data = n
}

func (n *NamespaceExport) Kind() SyntaxKind { return SyntaxKindNamespaceExport }

func NewNamespaceExport() *Node {
	n := &NamespaceExport{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNamespaceExport() *Node {
	return NewNamespaceExport()
}

type ExportSpecifier struct {
	NodeBase
}

func (n *Node) AsExportSpecifier() *ExportSpecifier { return n.data.(*ExportSpecifier) }
func (n *Node) IsExportSpecifier() bool             { return n.kind == SyntaxKindExportSpecifier }

func (n *ExportSpecifier) set() {
	*n = ExportSpecifier{}
	n.kind = SyntaxKindExportSpecifier
	n.data = n
}

func (n *ExportSpecifier) Kind() SyntaxKind { return SyntaxKindExportSpecifier }

func NewExportSpecifier() *Node {
	n := &ExportSpecifier{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewExportSpecifier() *Node {
	return NewExportSpecifier()
}

type MissingDeclaration struct {
	NodeBase
}

func (n *Node) AsMissingDeclaration() *MissingDeclaration { return n.data.(*MissingDeclaration) }
func (n *Node) IsMissingDeclaration() bool                { return n.kind == SyntaxKindMissingDeclaration }

func (n *MissingDeclaration) set() {
	*n = MissingDeclaration{}
	n.kind = SyntaxKindMissingDeclaration
	n.data = n
}

func (n *MissingDeclaration) Kind() SyntaxKind { return SyntaxKindMissingDeclaration }

func NewMissingDeclaration() *Node {
	n := &MissingDeclaration{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewMissingDeclaration() *Node {
	return NewMissingDeclaration()
}

type ExternalModuleReference struct {
	NodeBase
}

func (n *Node) AsExternalModuleReference() *ExternalModuleReference {
	return n.data.(*ExternalModuleReference)
}
func (n *Node) IsExternalModuleReference() bool { return n.kind == SyntaxKindExternalModuleReference }

func (n *ExternalModuleReference) set() {
	*n = ExternalModuleReference{}
	n.kind = SyntaxKindExternalModuleReference
	n.data = n
}

func (n *ExternalModuleReference) Kind() SyntaxKind { return SyntaxKindExternalModuleReference }

func NewExternalModuleReference() *Node {
	n := &ExternalModuleReference{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewExternalModuleReference() *Node {
	return NewExternalModuleReference()
}

type JsxElement struct {
	NodeBase
}

func (n *Node) AsJsxElement() *JsxElement { return n.data.(*JsxElement) }
func (n *Node) IsJsxElement() bool        { return n.kind == SyntaxKindJsxElement }

func (n *JsxElement) set() {
	*n = JsxElement{}
	n.kind = SyntaxKindJsxElement
	n.data = n
}

func (n *JsxElement) Kind() SyntaxKind { return SyntaxKindJsxElement }

func NewJsxElement() *Node {
	n := &JsxElement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxElement() *Node {
	return NewJsxElement()
}

type JsxSelfClosingElement struct {
	NodeBase
}

func (n *Node) AsJsxSelfClosingElement() *JsxSelfClosingElement {
	return n.data.(*JsxSelfClosingElement)
}
func (n *Node) IsJsxSelfClosingElement() bool { return n.kind == SyntaxKindJsxSelfClosingElement }

func (n *JsxSelfClosingElement) set() {
	*n = JsxSelfClosingElement{}
	n.kind = SyntaxKindJsxSelfClosingElement
	n.data = n
}

func (n *JsxSelfClosingElement) Kind() SyntaxKind { return SyntaxKindJsxSelfClosingElement }

func NewJsxSelfClosingElement() *Node {
	n := &JsxSelfClosingElement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxSelfClosingElement() *Node {
	return NewJsxSelfClosingElement()
}

type JsxOpeningElement struct {
	NodeBase
}

func (n *Node) AsJsxOpeningElement() *JsxOpeningElement { return n.data.(*JsxOpeningElement) }
func (n *Node) IsJsxOpeningElement() bool               { return n.kind == SyntaxKindJsxOpeningElement }

func (n *JsxOpeningElement) set() {
	*n = JsxOpeningElement{}
	n.kind = SyntaxKindJsxOpeningElement
	n.data = n
}

func (n *JsxOpeningElement) Kind() SyntaxKind { return SyntaxKindJsxOpeningElement }

func NewJsxOpeningElement() *Node {
	n := &JsxOpeningElement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxOpeningElement() *Node {
	return NewJsxOpeningElement()
}

type JsxClosingElement struct {
	NodeBase
}

func (n *Node) AsJsxClosingElement() *JsxClosingElement { return n.data.(*JsxClosingElement) }
func (n *Node) IsJsxClosingElement() bool               { return n.kind == SyntaxKindJsxClosingElement }

func (n *JsxClosingElement) set() {
	*n = JsxClosingElement{}
	n.kind = SyntaxKindJsxClosingElement
	n.data = n
}

func (n *JsxClosingElement) Kind() SyntaxKind { return SyntaxKindJsxClosingElement }

func NewJsxClosingElement() *Node {
	n := &JsxClosingElement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxClosingElement() *Node {
	return NewJsxClosingElement()
}

type JsxFragment struct {
	NodeBase
}

func (n *Node) AsJsxFragment() *JsxFragment { return n.data.(*JsxFragment) }
func (n *Node) IsJsxFragment() bool         { return n.kind == SyntaxKindJsxFragment }

func (n *JsxFragment) set() {
	*n = JsxFragment{}
	n.kind = SyntaxKindJsxFragment
	n.data = n
}

func (n *JsxFragment) Kind() SyntaxKind { return SyntaxKindJsxFragment }

func NewJsxFragment() *Node {
	n := &JsxFragment{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxFragment() *Node {
	return NewJsxFragment()
}

type JsxOpeningFragment struct {
	NodeBase
}

func (n *Node) AsJsxOpeningFragment() *JsxOpeningFragment { return n.data.(*JsxOpeningFragment) }
func (n *Node) IsJsxOpeningFragment() bool                { return n.kind == SyntaxKindJsxOpeningFragment }

func (n *JsxOpeningFragment) set() {
	*n = JsxOpeningFragment{}
	n.kind = SyntaxKindJsxOpeningFragment
	n.data = n
}

func (n *JsxOpeningFragment) Kind() SyntaxKind { return SyntaxKindJsxOpeningFragment }

func NewJsxOpeningFragment() *Node {
	n := &JsxOpeningFragment{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxOpeningFragment() *Node {
	return NewJsxOpeningFragment()
}

type JsxClosingFragment struct {
	NodeBase
}

func (n *Node) AsJsxClosingFragment() *JsxClosingFragment { return n.data.(*JsxClosingFragment) }
func (n *Node) IsJsxClosingFragment() bool                { return n.kind == SyntaxKindJsxClosingFragment }

func (n *JsxClosingFragment) set() {
	*n = JsxClosingFragment{}
	n.kind = SyntaxKindJsxClosingFragment
	n.data = n
}

func (n *JsxClosingFragment) Kind() SyntaxKind { return SyntaxKindJsxClosingFragment }

func NewJsxClosingFragment() *Node {
	n := &JsxClosingFragment{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxClosingFragment() *Node {
	return NewJsxClosingFragment()
}

type JsxAttribute struct {
	NodeBase
}

func (n *Node) AsJsxAttribute() *JsxAttribute { return n.data.(*JsxAttribute) }
func (n *Node) IsJsxAttribute() bool          { return n.kind == SyntaxKindJsxAttribute }

func (n *JsxAttribute) set() {
	*n = JsxAttribute{}
	n.kind = SyntaxKindJsxAttribute
	n.data = n
}

func (n *JsxAttribute) Kind() SyntaxKind { return SyntaxKindJsxAttribute }

func NewJsxAttribute() *Node {
	n := &JsxAttribute{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxAttribute() *Node {
	return NewJsxAttribute()
}

type JsxAttributes struct {
	NodeBase
}

func (n *Node) AsJsxAttributes() *JsxAttributes { return n.data.(*JsxAttributes) }
func (n *Node) IsJsxAttributes() bool           { return n.kind == SyntaxKindJsxAttributes }

func (n *JsxAttributes) set() {
	*n = JsxAttributes{}
	n.kind = SyntaxKindJsxAttributes
	n.data = n
}

func (n *JsxAttributes) Kind() SyntaxKind { return SyntaxKindJsxAttributes }

func NewJsxAttributes() *Node {
	n := &JsxAttributes{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxAttributes() *Node {
	return NewJsxAttributes()
}

type JsxSpreadAttribute struct {
	NodeBase
}

func (n *Node) AsJsxSpreadAttribute() *JsxSpreadAttribute { return n.data.(*JsxSpreadAttribute) }
func (n *Node) IsJsxSpreadAttribute() bool                { return n.kind == SyntaxKindJsxSpreadAttribute }

func (n *JsxSpreadAttribute) set() {
	*n = JsxSpreadAttribute{}
	n.kind = SyntaxKindJsxSpreadAttribute
	n.data = n
}

func (n *JsxSpreadAttribute) Kind() SyntaxKind { return SyntaxKindJsxSpreadAttribute }

func NewJsxSpreadAttribute() *Node {
	n := &JsxSpreadAttribute{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxSpreadAttribute() *Node {
	return NewJsxSpreadAttribute()
}

type JsxExpression struct {
	NodeBase
}

func (n *Node) AsJsxExpression() *JsxExpression { return n.data.(*JsxExpression) }
func (n *Node) IsJsxExpression() bool           { return n.kind == SyntaxKindJsxExpression }

func (n *JsxExpression) set() {
	*n = JsxExpression{}
	n.kind = SyntaxKindJsxExpression
	n.data = n
}

func (n *JsxExpression) Kind() SyntaxKind { return SyntaxKindJsxExpression }

func NewJsxExpression() *Node {
	n := &JsxExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxExpression() *Node {
	return NewJsxExpression()
}

type JsxNamespacedName struct {
	NodeBase
}

func (n *Node) AsJsxNamespacedName() *JsxNamespacedName { return n.data.(*JsxNamespacedName) }
func (n *Node) IsJsxNamespacedName() bool               { return n.kind == SyntaxKindJsxNamespacedName }

func (n *JsxNamespacedName) set() {
	*n = JsxNamespacedName{}
	n.kind = SyntaxKindJsxNamespacedName
	n.data = n
}

func (n *JsxNamespacedName) Kind() SyntaxKind { return SyntaxKindJsxNamespacedName }

func NewJsxNamespacedName() *Node {
	n := &JsxNamespacedName{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJsxNamespacedName() *Node {
	return NewJsxNamespacedName()
}

type CaseClause struct {
	NodeBase
}

func (n *Node) AsCaseClause() *CaseClause { return n.data.(*CaseClause) }
func (n *Node) IsCaseClause() bool        { return n.kind == SyntaxKindCaseClause }

func (n *CaseClause) set() {
	*n = CaseClause{}
	n.kind = SyntaxKindCaseClause
	n.data = n
}

func (n *CaseClause) Kind() SyntaxKind { return SyntaxKindCaseClause }

func NewCaseClause() *Node {
	n := &CaseClause{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewCaseClause() *Node {
	return NewCaseClause()
}

type DefaultClause struct {
	NodeBase
}

func (n *Node) AsDefaultClause() *DefaultClause { return n.data.(*DefaultClause) }
func (n *Node) IsDefaultClause() bool           { return n.kind == SyntaxKindDefaultClause }

func (n *DefaultClause) set() {
	*n = DefaultClause{}
	n.kind = SyntaxKindDefaultClause
	n.data = n
}

func (n *DefaultClause) Kind() SyntaxKind { return SyntaxKindDefaultClause }

func NewDefaultClause() *Node {
	n := &DefaultClause{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewDefaultClause() *Node {
	return NewDefaultClause()
}

type HeritageClause struct {
	NodeBase
}

func (n *Node) AsHeritageClause() *HeritageClause { return n.data.(*HeritageClause) }
func (n *Node) IsHeritageClause() bool            { return n.kind == SyntaxKindHeritageClause }

func (n *HeritageClause) set() {
	*n = HeritageClause{}
	n.kind = SyntaxKindHeritageClause
	n.data = n
}

func (n *HeritageClause) Kind() SyntaxKind { return SyntaxKindHeritageClause }

func NewHeritageClause() *Node {
	n := &HeritageClause{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewHeritageClause() *Node {
	return NewHeritageClause()
}

type CatchClause struct {
	NodeBase
}

func (n *Node) AsCatchClause() *CatchClause { return n.data.(*CatchClause) }
func (n *Node) IsCatchClause() bool         { return n.kind == SyntaxKindCatchClause }

func (n *CatchClause) set() {
	*n = CatchClause{}
	n.kind = SyntaxKindCatchClause
	n.data = n
}

func (n *CatchClause) Kind() SyntaxKind { return SyntaxKindCatchClause }

func NewCatchClause() *Node {
	n := &CatchClause{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewCatchClause() *Node {
	return NewCatchClause()
}

type ImportAttributes struct {
	NodeBase
}

func (n *Node) AsImportAttributes() *ImportAttributes { return n.data.(*ImportAttributes) }
func (n *Node) IsImportAttributes() bool              { return n.kind == SyntaxKindImportAttributes }

func (n *ImportAttributes) set() {
	*n = ImportAttributes{}
	n.kind = SyntaxKindImportAttributes
	n.data = n
}

func (n *ImportAttributes) Kind() SyntaxKind { return SyntaxKindImportAttributes }

func NewImportAttributes() *Node {
	n := &ImportAttributes{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewImportAttributes() *Node {
	return NewImportAttributes()
}

type ImportAttribute struct {
	NodeBase
}

func (n *Node) AsImportAttribute() *ImportAttribute { return n.data.(*ImportAttribute) }
func (n *Node) IsImportAttribute() bool             { return n.kind == SyntaxKindImportAttribute }

func (n *ImportAttribute) set() {
	*n = ImportAttribute{}
	n.kind = SyntaxKindImportAttribute
	n.data = n
}

func (n *ImportAttribute) Kind() SyntaxKind { return SyntaxKindImportAttribute }

func NewImportAttribute() *Node {
	n := &ImportAttribute{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewImportAttribute() *Node {
	return NewImportAttribute()
}

type PropertyAssignment struct {
	NodeBase
}

func (n *Node) AsPropertyAssignment() *PropertyAssignment { return n.data.(*PropertyAssignment) }
func (n *Node) IsPropertyAssignment() bool                { return n.kind == SyntaxKindPropertyAssignment }

func (n *PropertyAssignment) set() {
	*n = PropertyAssignment{}
	n.kind = SyntaxKindPropertyAssignment
	n.data = n
}

func (n *PropertyAssignment) Kind() SyntaxKind { return SyntaxKindPropertyAssignment }

func NewPropertyAssignment() *Node {
	n := &PropertyAssignment{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPropertyAssignment() *Node {
	return NewPropertyAssignment()
}

type ShorthandPropertyAssignment struct {
	NodeBase
}

func (n *Node) AsShorthandPropertyAssignment() *ShorthandPropertyAssignment {
	return n.data.(*ShorthandPropertyAssignment)
}
func (n *Node) IsShorthandPropertyAssignment() bool {
	return n.kind == SyntaxKindShorthandPropertyAssignment
}

func (n *ShorthandPropertyAssignment) set() {
	*n = ShorthandPropertyAssignment{}
	n.kind = SyntaxKindShorthandPropertyAssignment
	n.data = n
}

func (n *ShorthandPropertyAssignment) Kind() SyntaxKind { return SyntaxKindShorthandPropertyAssignment }

func NewShorthandPropertyAssignment() *Node {
	n := &ShorthandPropertyAssignment{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewShorthandPropertyAssignment() *Node {
	return NewShorthandPropertyAssignment()
}

type SpreadAssignment struct {
	NodeBase
}

func (n *Node) AsSpreadAssignment() *SpreadAssignment { return n.data.(*SpreadAssignment) }
func (n *Node) IsSpreadAssignment() bool              { return n.kind == SyntaxKindSpreadAssignment }

func (n *SpreadAssignment) set() {
	*n = SpreadAssignment{}
	n.kind = SyntaxKindSpreadAssignment
	n.data = n
}

func (n *SpreadAssignment) Kind() SyntaxKind { return SyntaxKindSpreadAssignment }

func NewSpreadAssignment() *Node {
	n := &SpreadAssignment{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSpreadAssignment() *Node {
	return NewSpreadAssignment()
}

type EnumMember struct {
	NodeBase
}

func (n *Node) AsEnumMember() *EnumMember { return n.data.(*EnumMember) }
func (n *Node) IsEnumMember() bool        { return n.kind == SyntaxKindEnumMember }

func (n *EnumMember) set() {
	*n = EnumMember{}
	n.kind = SyntaxKindEnumMember
	n.data = n
}

func (n *EnumMember) Kind() SyntaxKind { return SyntaxKindEnumMember }

func NewEnumMember() *Node {
	n := &EnumMember{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewEnumMember() *Node {
	return NewEnumMember()
}

type SourceFile struct {
	NodeBase
}

func (n *Node) AsSourceFile() *SourceFile { return n.data.(*SourceFile) }
func (n *Node) IsSourceFile() bool        { return n.kind == SyntaxKindSourceFile }

func (n *SourceFile) set() {
	*n = SourceFile{}
	n.kind = SyntaxKindSourceFile
	n.data = n
}

func (n *SourceFile) Kind() SyntaxKind { return SyntaxKindSourceFile }

func NewSourceFile() *Node {
	n := &SourceFile{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSourceFile() *Node {
	return NewSourceFile()
}

type Bundle struct {
	NodeBase
}

func (n *Node) AsBundle() *Bundle { return n.data.(*Bundle) }
func (n *Node) IsBundle() bool    { return n.kind == SyntaxKindBundle }

func (n *Bundle) set() {
	*n = Bundle{}
	n.kind = SyntaxKindBundle
	n.data = n
}

func (n *Bundle) Kind() SyntaxKind { return SyntaxKindBundle }

func NewBundle() *Node {
	n := &Bundle{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewBundle() *Node {
	return NewBundle()
}

type JSDocTypeExpression struct {
	NodeBase
}

func (n *Node) AsJSDocTypeExpression() *JSDocTypeExpression { return n.data.(*JSDocTypeExpression) }
func (n *Node) IsJSDocTypeExpression() bool                 { return n.kind == SyntaxKindJSDocTypeExpression }

func (n *JSDocTypeExpression) set() {
	*n = JSDocTypeExpression{}
	n.kind = SyntaxKindJSDocTypeExpression
	n.data = n
}

func (n *JSDocTypeExpression) Kind() SyntaxKind { return SyntaxKindJSDocTypeExpression }

func NewJSDocTypeExpression() *Node {
	n := &JSDocTypeExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocTypeExpression() *Node {
	return NewJSDocTypeExpression()
}

type JSDocNameReference struct {
	NodeBase
}

func (n *Node) AsJSDocNameReference() *JSDocNameReference { return n.data.(*JSDocNameReference) }
func (n *Node) IsJSDocNameReference() bool                { return n.kind == SyntaxKindJSDocNameReference }

func (n *JSDocNameReference) set() {
	*n = JSDocNameReference{}
	n.kind = SyntaxKindJSDocNameReference
	n.data = n
}

func (n *JSDocNameReference) Kind() SyntaxKind { return SyntaxKindJSDocNameReference }

func NewJSDocNameReference() *Node {
	n := &JSDocNameReference{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocNameReference() *Node {
	return NewJSDocNameReference()
}

type JSDocMemberName struct {
	NodeBase
}

func (n *Node) AsJSDocMemberName() *JSDocMemberName { return n.data.(*JSDocMemberName) }
func (n *Node) IsJSDocMemberName() bool             { return n.kind == SyntaxKindJSDocMemberName }

func (n *JSDocMemberName) set() {
	*n = JSDocMemberName{}
	n.kind = SyntaxKindJSDocMemberName
	n.data = n
}

func (n *JSDocMemberName) Kind() SyntaxKind { return SyntaxKindJSDocMemberName }

func NewJSDocMemberName() *Node {
	n := &JSDocMemberName{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocMemberName() *Node {
	return NewJSDocMemberName()
}

type JSDocAllType struct {
	NodeBase
}

func (n *Node) AsJSDocAllType() *JSDocAllType { return n.data.(*JSDocAllType) }
func (n *Node) IsJSDocAllType() bool          { return n.kind == SyntaxKindJSDocAllType }

func (n *JSDocAllType) set() {
	*n = JSDocAllType{}
	n.kind = SyntaxKindJSDocAllType
	n.data = n
}

func (n *JSDocAllType) Kind() SyntaxKind { return SyntaxKindJSDocAllType }

func NewJSDocAllType() *Node {
	n := &JSDocAllType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocAllType() *Node {
	return NewJSDocAllType()
}

type JSDocUnknownType struct {
	NodeBase
}

func (n *Node) AsJSDocUnknownType() *JSDocUnknownType { return n.data.(*JSDocUnknownType) }
func (n *Node) IsJSDocUnknownType() bool              { return n.kind == SyntaxKindJSDocUnknownType }

func (n *JSDocUnknownType) set() {
	*n = JSDocUnknownType{}
	n.kind = SyntaxKindJSDocUnknownType
	n.data = n
}

func (n *JSDocUnknownType) Kind() SyntaxKind { return SyntaxKindJSDocUnknownType }

func NewJSDocUnknownType() *Node {
	n := &JSDocUnknownType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocUnknownType() *Node {
	return NewJSDocUnknownType()
}

type JSDocNullableType struct {
	NodeBase
}

func (n *Node) AsJSDocNullableType() *JSDocNullableType { return n.data.(*JSDocNullableType) }
func (n *Node) IsJSDocNullableType() bool               { return n.kind == SyntaxKindJSDocNullableType }

func (n *JSDocNullableType) set() {
	*n = JSDocNullableType{}
	n.kind = SyntaxKindJSDocNullableType
	n.data = n
}

func (n *JSDocNullableType) Kind() SyntaxKind { return SyntaxKindJSDocNullableType }

func NewJSDocNullableType() *Node {
	n := &JSDocNullableType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocNullableType() *Node {
	return NewJSDocNullableType()
}

type JSDocNonNullableType struct {
	NodeBase
}

func (n *Node) AsJSDocNonNullableType() *JSDocNonNullableType { return n.data.(*JSDocNonNullableType) }
func (n *Node) IsJSDocNonNullableType() bool                  { return n.kind == SyntaxKindJSDocNonNullableType }

func (n *JSDocNonNullableType) set() {
	*n = JSDocNonNullableType{}
	n.kind = SyntaxKindJSDocNonNullableType
	n.data = n
}

func (n *JSDocNonNullableType) Kind() SyntaxKind { return SyntaxKindJSDocNonNullableType }

func NewJSDocNonNullableType() *Node {
	n := &JSDocNonNullableType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocNonNullableType() *Node {
	return NewJSDocNonNullableType()
}

type JSDocOptionalType struct {
	NodeBase
}

func (n *Node) AsJSDocOptionalType() *JSDocOptionalType { return n.data.(*JSDocOptionalType) }
func (n *Node) IsJSDocOptionalType() bool               { return n.kind == SyntaxKindJSDocOptionalType }

func (n *JSDocOptionalType) set() {
	*n = JSDocOptionalType{}
	n.kind = SyntaxKindJSDocOptionalType
	n.data = n
}

func (n *JSDocOptionalType) Kind() SyntaxKind { return SyntaxKindJSDocOptionalType }

func NewJSDocOptionalType() *Node {
	n := &JSDocOptionalType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocOptionalType() *Node {
	return NewJSDocOptionalType()
}

type JSDocFunctionType struct {
	NodeBase
}

func (n *Node) AsJSDocFunctionType() *JSDocFunctionType { return n.data.(*JSDocFunctionType) }
func (n *Node) IsJSDocFunctionType() bool               { return n.kind == SyntaxKindJSDocFunctionType }

func (n *JSDocFunctionType) set() {
	*n = JSDocFunctionType{}
	n.kind = SyntaxKindJSDocFunctionType
	n.data = n
}

func (n *JSDocFunctionType) Kind() SyntaxKind { return SyntaxKindJSDocFunctionType }

func NewJSDocFunctionType() *Node {
	n := &JSDocFunctionType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocFunctionType() *Node {
	return NewJSDocFunctionType()
}

type JSDocVariadicType struct {
	NodeBase
}

func (n *Node) AsJSDocVariadicType() *JSDocVariadicType { return n.data.(*JSDocVariadicType) }
func (n *Node) IsJSDocVariadicType() bool               { return n.kind == SyntaxKindJSDocVariadicType }

func (n *JSDocVariadicType) set() {
	*n = JSDocVariadicType{}
	n.kind = SyntaxKindJSDocVariadicType
	n.data = n
}

func (n *JSDocVariadicType) Kind() SyntaxKind { return SyntaxKindJSDocVariadicType }

func NewJSDocVariadicType() *Node {
	n := &JSDocVariadicType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocVariadicType() *Node {
	return NewJSDocVariadicType()
}

type JSDocNamepathType struct {
	NodeBase
}

func (n *Node) AsJSDocNamepathType() *JSDocNamepathType { return n.data.(*JSDocNamepathType) }
func (n *Node) IsJSDocNamepathType() bool               { return n.kind == SyntaxKindJSDocNamepathType }

func (n *JSDocNamepathType) set() {
	*n = JSDocNamepathType{}
	n.kind = SyntaxKindJSDocNamepathType
	n.data = n
}

func (n *JSDocNamepathType) Kind() SyntaxKind { return SyntaxKindJSDocNamepathType }

func NewJSDocNamepathType() *Node {
	n := &JSDocNamepathType{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocNamepathType() *Node {
	return NewJSDocNamepathType()
}

type JSDoc struct {
	NodeBase
}

func (n *Node) AsJSDoc() *JSDoc { return n.data.(*JSDoc) }
func (n *Node) IsJSDoc() bool   { return n.kind == SyntaxKindJSDoc }

func (n *JSDoc) set() {
	*n = JSDoc{}
	n.kind = SyntaxKindJSDoc
	n.data = n
}

func (n *JSDoc) Kind() SyntaxKind { return SyntaxKindJSDoc }

func NewJSDoc() *Node {
	n := &JSDoc{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDoc() *Node {
	return NewJSDoc()
}

type JSDocText struct {
	NodeBase
}

func (n *Node) AsJSDocText() *JSDocText { return n.data.(*JSDocText) }
func (n *Node) IsJSDocText() bool       { return n.kind == SyntaxKindJSDocText }

func (n *JSDocText) set() {
	*n = JSDocText{}
	n.kind = SyntaxKindJSDocText
	n.data = n
}

func (n *JSDocText) Kind() SyntaxKind { return SyntaxKindJSDocText }

func NewJSDocText() *Node {
	n := &JSDocText{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocText() *Node {
	return NewJSDocText()
}

type JSDocTypeLiteral struct {
	NodeBase
}

func (n *Node) AsJSDocTypeLiteral() *JSDocTypeLiteral { return n.data.(*JSDocTypeLiteral) }
func (n *Node) IsJSDocTypeLiteral() bool              { return n.kind == SyntaxKindJSDocTypeLiteral }

func (n *JSDocTypeLiteral) set() {
	*n = JSDocTypeLiteral{}
	n.kind = SyntaxKindJSDocTypeLiteral
	n.data = n
}

func (n *JSDocTypeLiteral) Kind() SyntaxKind { return SyntaxKindJSDocTypeLiteral }

func NewJSDocTypeLiteral() *Node {
	n := &JSDocTypeLiteral{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocTypeLiteral() *Node {
	return NewJSDocTypeLiteral()
}

type JSDocSignature struct {
	NodeBase
}

func (n *Node) AsJSDocSignature() *JSDocSignature { return n.data.(*JSDocSignature) }
func (n *Node) IsJSDocSignature() bool            { return n.kind == SyntaxKindJSDocSignature }

func (n *JSDocSignature) set() {
	*n = JSDocSignature{}
	n.kind = SyntaxKindJSDocSignature
	n.data = n
}

func (n *JSDocSignature) Kind() SyntaxKind { return SyntaxKindJSDocSignature }

func NewJSDocSignature() *Node {
	n := &JSDocSignature{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocSignature() *Node {
	return NewJSDocSignature()
}

type JSDocLink struct {
	NodeBase
}

func (n *Node) AsJSDocLink() *JSDocLink { return n.data.(*JSDocLink) }
func (n *Node) IsJSDocLink() bool       { return n.kind == SyntaxKindJSDocLink }

func (n *JSDocLink) set() {
	*n = JSDocLink{}
	n.kind = SyntaxKindJSDocLink
	n.data = n
}

func (n *JSDocLink) Kind() SyntaxKind { return SyntaxKindJSDocLink }

func NewJSDocLink() *Node {
	n := &JSDocLink{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocLink() *Node {
	return NewJSDocLink()
}

type JSDocLinkCode struct {
	NodeBase
}

func (n *Node) AsJSDocLinkCode() *JSDocLinkCode { return n.data.(*JSDocLinkCode) }
func (n *Node) IsJSDocLinkCode() bool           { return n.kind == SyntaxKindJSDocLinkCode }

func (n *JSDocLinkCode) set() {
	*n = JSDocLinkCode{}
	n.kind = SyntaxKindJSDocLinkCode
	n.data = n
}

func (n *JSDocLinkCode) Kind() SyntaxKind { return SyntaxKindJSDocLinkCode }

func NewJSDocLinkCode() *Node {
	n := &JSDocLinkCode{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocLinkCode() *Node {
	return NewJSDocLinkCode()
}

type JSDocLinkPlain struct {
	NodeBase
}

func (n *Node) AsJSDocLinkPlain() *JSDocLinkPlain { return n.data.(*JSDocLinkPlain) }
func (n *Node) IsJSDocLinkPlain() bool            { return n.kind == SyntaxKindJSDocLinkPlain }

func (n *JSDocLinkPlain) set() {
	*n = JSDocLinkPlain{}
	n.kind = SyntaxKindJSDocLinkPlain
	n.data = n
}

func (n *JSDocLinkPlain) Kind() SyntaxKind { return SyntaxKindJSDocLinkPlain }

func NewJSDocLinkPlain() *Node {
	n := &JSDocLinkPlain{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocLinkPlain() *Node {
	return NewJSDocLinkPlain()
}

type JSDocTag struct {
	NodeBase
}

func (n *Node) AsJSDocTag() *JSDocTag { return n.data.(*JSDocTag) }
func (n *Node) IsJSDocTag() bool      { return n.kind == SyntaxKindJSDocTag }

func (n *JSDocTag) set() {
	*n = JSDocTag{}
	n.kind = SyntaxKindJSDocTag
	n.data = n
}

func (n *JSDocTag) Kind() SyntaxKind { return SyntaxKindJSDocTag }

func NewJSDocTag() *Node {
	n := &JSDocTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocTag() *Node {
	return NewJSDocTag()
}

type JSDocAugmentsTag struct {
	NodeBase
}

func (n *Node) AsJSDocAugmentsTag() *JSDocAugmentsTag { return n.data.(*JSDocAugmentsTag) }
func (n *Node) IsJSDocAugmentsTag() bool              { return n.kind == SyntaxKindJSDocAugmentsTag }

func (n *JSDocAugmentsTag) set() {
	*n = JSDocAugmentsTag{}
	n.kind = SyntaxKindJSDocAugmentsTag
	n.data = n
}

func (n *JSDocAugmentsTag) Kind() SyntaxKind { return SyntaxKindJSDocAugmentsTag }

func NewJSDocAugmentsTag() *Node {
	n := &JSDocAugmentsTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocAugmentsTag() *Node {
	return NewJSDocAugmentsTag()
}

type JSDocImplementsTag struct {
	NodeBase
}

func (n *Node) AsJSDocImplementsTag() *JSDocImplementsTag { return n.data.(*JSDocImplementsTag) }
func (n *Node) IsJSDocImplementsTag() bool                { return n.kind == SyntaxKindJSDocImplementsTag }

func (n *JSDocImplementsTag) set() {
	*n = JSDocImplementsTag{}
	n.kind = SyntaxKindJSDocImplementsTag
	n.data = n
}

func (n *JSDocImplementsTag) Kind() SyntaxKind { return SyntaxKindJSDocImplementsTag }

func NewJSDocImplementsTag() *Node {
	n := &JSDocImplementsTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocImplementsTag() *Node {
	return NewJSDocImplementsTag()
}

type JSDocAuthorTag struct {
	NodeBase
}

func (n *Node) AsJSDocAuthorTag() *JSDocAuthorTag { return n.data.(*JSDocAuthorTag) }
func (n *Node) IsJSDocAuthorTag() bool            { return n.kind == SyntaxKindJSDocAuthorTag }

func (n *JSDocAuthorTag) set() {
	*n = JSDocAuthorTag{}
	n.kind = SyntaxKindJSDocAuthorTag
	n.data = n
}

func (n *JSDocAuthorTag) Kind() SyntaxKind { return SyntaxKindJSDocAuthorTag }

func NewJSDocAuthorTag() *Node {
	n := &JSDocAuthorTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocAuthorTag() *Node {
	return NewJSDocAuthorTag()
}

type JSDocDeprecatedTag struct {
	NodeBase
}

func (n *Node) AsJSDocDeprecatedTag() *JSDocDeprecatedTag { return n.data.(*JSDocDeprecatedTag) }
func (n *Node) IsJSDocDeprecatedTag() bool                { return n.kind == SyntaxKindJSDocDeprecatedTag }

func (n *JSDocDeprecatedTag) set() {
	*n = JSDocDeprecatedTag{}
	n.kind = SyntaxKindJSDocDeprecatedTag
	n.data = n
}

func (n *JSDocDeprecatedTag) Kind() SyntaxKind { return SyntaxKindJSDocDeprecatedTag }

func NewJSDocDeprecatedTag() *Node {
	n := &JSDocDeprecatedTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocDeprecatedTag() *Node {
	return NewJSDocDeprecatedTag()
}

type JSDocImmediateTag struct {
	NodeBase
}

func (n *Node) AsJSDocImmediateTag() *JSDocImmediateTag { return n.data.(*JSDocImmediateTag) }
func (n *Node) IsJSDocImmediateTag() bool               { return n.kind == SyntaxKindJSDocImmediateTag }

func (n *JSDocImmediateTag) set() {
	*n = JSDocImmediateTag{}
	n.kind = SyntaxKindJSDocImmediateTag
	n.data = n
}

func (n *JSDocImmediateTag) Kind() SyntaxKind { return SyntaxKindJSDocImmediateTag }

func NewJSDocImmediateTag() *Node {
	n := &JSDocImmediateTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocImmediateTag() *Node {
	return NewJSDocImmediateTag()
}

type JSDocClassTag struct {
	NodeBase
}

func (n *Node) AsJSDocClassTag() *JSDocClassTag { return n.data.(*JSDocClassTag) }
func (n *Node) IsJSDocClassTag() bool           { return n.kind == SyntaxKindJSDocClassTag }

func (n *JSDocClassTag) set() {
	*n = JSDocClassTag{}
	n.kind = SyntaxKindJSDocClassTag
	n.data = n
}

func (n *JSDocClassTag) Kind() SyntaxKind { return SyntaxKindJSDocClassTag }

func NewJSDocClassTag() *Node {
	n := &JSDocClassTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocClassTag() *Node {
	return NewJSDocClassTag()
}

type JSDocPublicTag struct {
	NodeBase
}

func (n *Node) AsJSDocPublicTag() *JSDocPublicTag { return n.data.(*JSDocPublicTag) }
func (n *Node) IsJSDocPublicTag() bool            { return n.kind == SyntaxKindJSDocPublicTag }

func (n *JSDocPublicTag) set() {
	*n = JSDocPublicTag{}
	n.kind = SyntaxKindJSDocPublicTag
	n.data = n
}

func (n *JSDocPublicTag) Kind() SyntaxKind { return SyntaxKindJSDocPublicTag }

func NewJSDocPublicTag() *Node {
	n := &JSDocPublicTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocPublicTag() *Node {
	return NewJSDocPublicTag()
}

type JSDocPrivateTag struct {
	NodeBase
}

func (n *Node) AsJSDocPrivateTag() *JSDocPrivateTag { return n.data.(*JSDocPrivateTag) }
func (n *Node) IsJSDocPrivateTag() bool             { return n.kind == SyntaxKindJSDocPrivateTag }

func (n *JSDocPrivateTag) set() {
	*n = JSDocPrivateTag{}
	n.kind = SyntaxKindJSDocPrivateTag
	n.data = n
}

func (n *JSDocPrivateTag) Kind() SyntaxKind { return SyntaxKindJSDocPrivateTag }

func NewJSDocPrivateTag() *Node {
	n := &JSDocPrivateTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocPrivateTag() *Node {
	return NewJSDocPrivateTag()
}

type JSDocProtectedTag struct {
	NodeBase
}

func (n *Node) AsJSDocProtectedTag() *JSDocProtectedTag { return n.data.(*JSDocProtectedTag) }
func (n *Node) IsJSDocProtectedTag() bool               { return n.kind == SyntaxKindJSDocProtectedTag }

func (n *JSDocProtectedTag) set() {
	*n = JSDocProtectedTag{}
	n.kind = SyntaxKindJSDocProtectedTag
	n.data = n
}

func (n *JSDocProtectedTag) Kind() SyntaxKind { return SyntaxKindJSDocProtectedTag }

func NewJSDocProtectedTag() *Node {
	n := &JSDocProtectedTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocProtectedTag() *Node {
	return NewJSDocProtectedTag()
}

type JSDocReadonlyTag struct {
	NodeBase
}

func (n *Node) AsJSDocReadonlyTag() *JSDocReadonlyTag { return n.data.(*JSDocReadonlyTag) }
func (n *Node) IsJSDocReadonlyTag() bool              { return n.kind == SyntaxKindJSDocReadonlyTag }

func (n *JSDocReadonlyTag) set() {
	*n = JSDocReadonlyTag{}
	n.kind = SyntaxKindJSDocReadonlyTag
	n.data = n
}

func (n *JSDocReadonlyTag) Kind() SyntaxKind { return SyntaxKindJSDocReadonlyTag }

func NewJSDocReadonlyTag() *Node {
	n := &JSDocReadonlyTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocReadonlyTag() *Node {
	return NewJSDocReadonlyTag()
}

type JSDocOverrideTag struct {
	NodeBase
}

func (n *Node) AsJSDocOverrideTag() *JSDocOverrideTag { return n.data.(*JSDocOverrideTag) }
func (n *Node) IsJSDocOverrideTag() bool              { return n.kind == SyntaxKindJSDocOverrideTag }

func (n *JSDocOverrideTag) set() {
	*n = JSDocOverrideTag{}
	n.kind = SyntaxKindJSDocOverrideTag
	n.data = n
}

func (n *JSDocOverrideTag) Kind() SyntaxKind { return SyntaxKindJSDocOverrideTag }

func NewJSDocOverrideTag() *Node {
	n := &JSDocOverrideTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocOverrideTag() *Node {
	return NewJSDocOverrideTag()
}

type JSDocCallbackTag struct {
	NodeBase
}

func (n *Node) AsJSDocCallbackTag() *JSDocCallbackTag { return n.data.(*JSDocCallbackTag) }
func (n *Node) IsJSDocCallbackTag() bool              { return n.kind == SyntaxKindJSDocCallbackTag }

func (n *JSDocCallbackTag) set() {
	*n = JSDocCallbackTag{}
	n.kind = SyntaxKindJSDocCallbackTag
	n.data = n
}

func (n *JSDocCallbackTag) Kind() SyntaxKind { return SyntaxKindJSDocCallbackTag }

func NewJSDocCallbackTag() *Node {
	n := &JSDocCallbackTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocCallbackTag() *Node {
	return NewJSDocCallbackTag()
}

type JSDocOverloadTag struct {
	NodeBase
}

func (n *Node) AsJSDocOverloadTag() *JSDocOverloadTag { return n.data.(*JSDocOverloadTag) }
func (n *Node) IsJSDocOverloadTag() bool              { return n.kind == SyntaxKindJSDocOverloadTag }

func (n *JSDocOverloadTag) set() {
	*n = JSDocOverloadTag{}
	n.kind = SyntaxKindJSDocOverloadTag
	n.data = n
}

func (n *JSDocOverloadTag) Kind() SyntaxKind { return SyntaxKindJSDocOverloadTag }

func NewJSDocOverloadTag() *Node {
	n := &JSDocOverloadTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocOverloadTag() *Node {
	return NewJSDocOverloadTag()
}

type JSDocEnumTag struct {
	NodeBase
}

func (n *Node) AsJSDocEnumTag() *JSDocEnumTag { return n.data.(*JSDocEnumTag) }
func (n *Node) IsJSDocEnumTag() bool          { return n.kind == SyntaxKindJSDocEnumTag }

func (n *JSDocEnumTag) set() {
	*n = JSDocEnumTag{}
	n.kind = SyntaxKindJSDocEnumTag
	n.data = n
}

func (n *JSDocEnumTag) Kind() SyntaxKind { return SyntaxKindJSDocEnumTag }

func NewJSDocEnumTag() *Node {
	n := &JSDocEnumTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocEnumTag() *Node {
	return NewJSDocEnumTag()
}

type JSDocParameterTag struct {
	NodeBase
}

func (n *Node) AsJSDocParameterTag() *JSDocParameterTag { return n.data.(*JSDocParameterTag) }
func (n *Node) IsJSDocParameterTag() bool               { return n.kind == SyntaxKindJSDocParameterTag }

func (n *JSDocParameterTag) set() {
	*n = JSDocParameterTag{}
	n.kind = SyntaxKindJSDocParameterTag
	n.data = n
}

func (n *JSDocParameterTag) Kind() SyntaxKind { return SyntaxKindJSDocParameterTag }

func NewJSDocParameterTag() *Node {
	n := &JSDocParameterTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocParameterTag() *Node {
	return NewJSDocParameterTag()
}

type JSDocReturnTag struct {
	NodeBase
}

func (n *Node) AsJSDocReturnTag() *JSDocReturnTag { return n.data.(*JSDocReturnTag) }
func (n *Node) IsJSDocReturnTag() bool            { return n.kind == SyntaxKindJSDocReturnTag }

func (n *JSDocReturnTag) set() {
	*n = JSDocReturnTag{}
	n.kind = SyntaxKindJSDocReturnTag
	n.data = n
}

func (n *JSDocReturnTag) Kind() SyntaxKind { return SyntaxKindJSDocReturnTag }

func NewJSDocReturnTag() *Node {
	n := &JSDocReturnTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocReturnTag() *Node {
	return NewJSDocReturnTag()
}

type JSDocThisTag struct {
	NodeBase
}

func (n *Node) AsJSDocThisTag() *JSDocThisTag { return n.data.(*JSDocThisTag) }
func (n *Node) IsJSDocThisTag() bool          { return n.kind == SyntaxKindJSDocThisTag }

func (n *JSDocThisTag) set() {
	*n = JSDocThisTag{}
	n.kind = SyntaxKindJSDocThisTag
	n.data = n
}

func (n *JSDocThisTag) Kind() SyntaxKind { return SyntaxKindJSDocThisTag }

func NewJSDocThisTag() *Node {
	n := &JSDocThisTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocThisTag() *Node {
	return NewJSDocThisTag()
}

type JSDocTypeTag struct {
	NodeBase
}

func (n *Node) AsJSDocTypeTag() *JSDocTypeTag { return n.data.(*JSDocTypeTag) }
func (n *Node) IsJSDocTypeTag() bool          { return n.kind == SyntaxKindJSDocTypeTag }

func (n *JSDocTypeTag) set() {
	*n = JSDocTypeTag{}
	n.kind = SyntaxKindJSDocTypeTag
	n.data = n
}

func (n *JSDocTypeTag) Kind() SyntaxKind { return SyntaxKindJSDocTypeTag }

func NewJSDocTypeTag() *Node {
	n := &JSDocTypeTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocTypeTag() *Node {
	return NewJSDocTypeTag()
}

type JSDocTemplateTag struct {
	NodeBase
}

func (n *Node) AsJSDocTemplateTag() *JSDocTemplateTag { return n.data.(*JSDocTemplateTag) }
func (n *Node) IsJSDocTemplateTag() bool              { return n.kind == SyntaxKindJSDocTemplateTag }

func (n *JSDocTemplateTag) set() {
	*n = JSDocTemplateTag{}
	n.kind = SyntaxKindJSDocTemplateTag
	n.data = n
}

func (n *JSDocTemplateTag) Kind() SyntaxKind { return SyntaxKindJSDocTemplateTag }

func NewJSDocTemplateTag() *Node {
	n := &JSDocTemplateTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocTemplateTag() *Node {
	return NewJSDocTemplateTag()
}

type JSDocTypedefTag struct {
	NodeBase
}

func (n *Node) AsJSDocTypedefTag() *JSDocTypedefTag { return n.data.(*JSDocTypedefTag) }
func (n *Node) IsJSDocTypedefTag() bool             { return n.kind == SyntaxKindJSDocTypedefTag }

func (n *JSDocTypedefTag) set() {
	*n = JSDocTypedefTag{}
	n.kind = SyntaxKindJSDocTypedefTag
	n.data = n
}

func (n *JSDocTypedefTag) Kind() SyntaxKind { return SyntaxKindJSDocTypedefTag }

func NewJSDocTypedefTag() *Node {
	n := &JSDocTypedefTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocTypedefTag() *Node {
	return NewJSDocTypedefTag()
}

type JSDocSeeTag struct {
	NodeBase
}

func (n *Node) AsJSDocSeeTag() *JSDocSeeTag { return n.data.(*JSDocSeeTag) }
func (n *Node) IsJSDocSeeTag() bool         { return n.kind == SyntaxKindJSDocSeeTag }

func (n *JSDocSeeTag) set() {
	*n = JSDocSeeTag{}
	n.kind = SyntaxKindJSDocSeeTag
	n.data = n
}

func (n *JSDocSeeTag) Kind() SyntaxKind { return SyntaxKindJSDocSeeTag }

func NewJSDocSeeTag() *Node {
	n := &JSDocSeeTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocSeeTag() *Node {
	return NewJSDocSeeTag()
}

type JSDocPropertyTag struct {
	NodeBase
}

func (n *Node) AsJSDocPropertyTag() *JSDocPropertyTag { return n.data.(*JSDocPropertyTag) }
func (n *Node) IsJSDocPropertyTag() bool              { return n.kind == SyntaxKindJSDocPropertyTag }

func (n *JSDocPropertyTag) set() {
	*n = JSDocPropertyTag{}
	n.kind = SyntaxKindJSDocPropertyTag
	n.data = n
}

func (n *JSDocPropertyTag) Kind() SyntaxKind { return SyntaxKindJSDocPropertyTag }

func NewJSDocPropertyTag() *Node {
	n := &JSDocPropertyTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocPropertyTag() *Node {
	return NewJSDocPropertyTag()
}

type JSDocThrowsTag struct {
	NodeBase
}

func (n *Node) AsJSDocThrowsTag() *JSDocThrowsTag { return n.data.(*JSDocThrowsTag) }
func (n *Node) IsJSDocThrowsTag() bool            { return n.kind == SyntaxKindJSDocThrowsTag }

func (n *JSDocThrowsTag) set() {
	*n = JSDocThrowsTag{}
	n.kind = SyntaxKindJSDocThrowsTag
	n.data = n
}

func (n *JSDocThrowsTag) Kind() SyntaxKind { return SyntaxKindJSDocThrowsTag }

func NewJSDocThrowsTag() *Node {
	n := &JSDocThrowsTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocThrowsTag() *Node {
	return NewJSDocThrowsTag()
}

type JSDocSatisfiesTag struct {
	NodeBase
}

func (n *Node) AsJSDocSatisfiesTag() *JSDocSatisfiesTag { return n.data.(*JSDocSatisfiesTag) }
func (n *Node) IsJSDocSatisfiesTag() bool               { return n.kind == SyntaxKindJSDocSatisfiesTag }

func (n *JSDocSatisfiesTag) set() {
	*n = JSDocSatisfiesTag{}
	n.kind = SyntaxKindJSDocSatisfiesTag
	n.data = n
}

func (n *JSDocSatisfiesTag) Kind() SyntaxKind { return SyntaxKindJSDocSatisfiesTag }

func NewJSDocSatisfiesTag() *Node {
	n := &JSDocSatisfiesTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocSatisfiesTag() *Node {
	return NewJSDocSatisfiesTag()
}

type JSDocImportTag struct {
	NodeBase
}

func (n *Node) AsJSDocImportTag() *JSDocImportTag { return n.data.(*JSDocImportTag) }
func (n *Node) IsJSDocImportTag() bool            { return n.kind == SyntaxKindJSDocImportTag }

func (n *JSDocImportTag) set() {
	*n = JSDocImportTag{}
	n.kind = SyntaxKindJSDocImportTag
	n.data = n
}

func (n *JSDocImportTag) Kind() SyntaxKind { return SyntaxKindJSDocImportTag }

func NewJSDocImportTag() *Node {
	n := &JSDocImportTag{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewJSDocImportTag() *Node {
	return NewJSDocImportTag()
}

type SyntaxList struct {
	NodeBase
}

func (n *Node) AsSyntaxList() *SyntaxList { return n.data.(*SyntaxList) }
func (n *Node) IsSyntaxList() bool        { return n.kind == SyntaxKindSyntaxList }

func (n *SyntaxList) set() {
	*n = SyntaxList{}
	n.kind = SyntaxKindSyntaxList
	n.data = n
}

func (n *SyntaxList) Kind() SyntaxKind { return SyntaxKindSyntaxList }

func NewSyntaxList() *Node {
	n := &SyntaxList{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSyntaxList() *Node {
	return NewSyntaxList()
}

type NotEmittedStatement struct {
	NodeBase
}

func (n *Node) AsNotEmittedStatement() *NotEmittedStatement { return n.data.(*NotEmittedStatement) }
func (n *Node) IsNotEmittedStatement() bool                 { return n.kind == SyntaxKindNotEmittedStatement }

func (n *NotEmittedStatement) set() {
	*n = NotEmittedStatement{}
	n.kind = SyntaxKindNotEmittedStatement
	n.data = n
}

func (n *NotEmittedStatement) Kind() SyntaxKind { return SyntaxKindNotEmittedStatement }

func NewNotEmittedStatement() *Node {
	n := &NotEmittedStatement{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewNotEmittedStatement() *Node {
	return NewNotEmittedStatement()
}

type PartiallyEmittedExpression struct {
	NodeBase
}

func (n *Node) AsPartiallyEmittedExpression() *PartiallyEmittedExpression {
	return n.data.(*PartiallyEmittedExpression)
}
func (n *Node) IsPartiallyEmittedExpression() bool {
	return n.kind == SyntaxKindPartiallyEmittedExpression
}

func (n *PartiallyEmittedExpression) set() {
	*n = PartiallyEmittedExpression{}
	n.kind = SyntaxKindPartiallyEmittedExpression
	n.data = n
}

func (n *PartiallyEmittedExpression) Kind() SyntaxKind { return SyntaxKindPartiallyEmittedExpression }

func NewPartiallyEmittedExpression() *Node {
	n := &PartiallyEmittedExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewPartiallyEmittedExpression() *Node {
	return NewPartiallyEmittedExpression()
}

type CommaListExpression struct {
	NodeBase
}

func (n *Node) AsCommaListExpression() *CommaListExpression { return n.data.(*CommaListExpression) }
func (n *Node) IsCommaListExpression() bool                 { return n.kind == SyntaxKindCommaListExpression }

func (n *CommaListExpression) set() {
	*n = CommaListExpression{}
	n.kind = SyntaxKindCommaListExpression
	n.data = n
}

func (n *CommaListExpression) Kind() SyntaxKind { return SyntaxKindCommaListExpression }

func NewCommaListExpression() *Node {
	n := &CommaListExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewCommaListExpression() *Node {
	return NewCommaListExpression()
}

type SyntheticReferenceExpression struct {
	NodeBase
}

func (n *Node) AsSyntheticReferenceExpression() *SyntheticReferenceExpression {
	return n.data.(*SyntheticReferenceExpression)
}
func (n *Node) IsSyntheticReferenceExpression() bool {
	return n.kind == SyntaxKindSyntheticReferenceExpression
}

func (n *SyntheticReferenceExpression) set() {
	*n = SyntheticReferenceExpression{}
	n.kind = SyntaxKindSyntheticReferenceExpression
	n.data = n
}

func (n *SyntheticReferenceExpression) Kind() SyntaxKind {
	return SyntaxKindSyntheticReferenceExpression
}

func NewSyntheticReferenceExpression() *Node {
	n := &SyntheticReferenceExpression{}
	n.set()
	return n.AsNode()
}

func (f *Factory) NewSyntheticReferenceExpression() *Node {
	return NewSyntheticReferenceExpression()
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
