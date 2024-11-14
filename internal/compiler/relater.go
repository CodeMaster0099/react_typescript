package compiler

import (
	"iter"
	"slices"
	"strconv"
	"strings"

	"github.com/microsoft/typescript-go/internal/ast"
	"github.com/microsoft/typescript-go/internal/compiler/diagnostics"
	"github.com/microsoft/typescript-go/internal/core"
)

type SignatureCheckMode uint32

const (
	SignatureCheckModeNone               SignatureCheckMode = 0
	SignatureCheckModeBivariantCallback  SignatureCheckMode = 1 << 0
	SignatureCheckModeStrictCallback     SignatureCheckMode = 1 << 1
	SignatureCheckModeIgnoreReturnTypes  SignatureCheckMode = 1 << 2
	SignatureCheckModeStrictArity        SignatureCheckMode = 1 << 3
	SignatureCheckModeStrictTopSignature SignatureCheckMode = 1 << 4
	SignatureCheckModeCallback           SignatureCheckMode = SignatureCheckModeBivariantCallback | SignatureCheckModeStrictCallback
)

type MinArgumentCountFlags uint32

const (
	MinArgumentCountFlagsNone                    MinArgumentCountFlags = 0
	MinArgumentCountFlagsStrongArityForUntypedJS MinArgumentCountFlags = 1 << 0
	MinArgumentCountFlagsVoidIsNonOptional       MinArgumentCountFlags = 1 << 1
)

type IntersectionState uint32

const (
	IntersectionStateNone   IntersectionState = 0
	IntersectionStateSource IntersectionState = 1 << 0 // Source type is a constituent of an outer intersection
	IntersectionStateTarget IntersectionState = 1 << 1 // Target type is a constituent of an outer intersection
)

type RecursionFlags uint32

const (
	RecursionFlagsNone   RecursionFlags = 0
	RecursionFlagsSource RecursionFlags = 1 << 0
	RecursionFlagsTarget RecursionFlags = 1 << 1
	RecursionFlagsBoth                  = RecursionFlagsSource | RecursionFlagsTarget
)

type ExpandingFlags uint8

const (
	ExpandingFlagsNone   ExpandingFlags = 0
	ExpandingFlagsSource ExpandingFlags = 1
	ExpandingFlagsTarget ExpandingFlags = 1 << 1
	ExpandingFlagsBoth                  = ExpandingFlagsSource | ExpandingFlagsTarget
)

type RelationComparisonResult uint32

const (
	RelationComparisonResultNone                RelationComparisonResult = 0
	RelationComparisonResultSucceeded           RelationComparisonResult = 1 << 0
	RelationComparisonResultFailed              RelationComparisonResult = 1 << 1
	RelationComparisonResultReportsUnmeasurable RelationComparisonResult = 1 << 3
	RelationComparisonResultReportsUnreliable   RelationComparisonResult = 1 << 4
	RelationComparisonResultComplexityOverflow  RelationComparisonResult = 1 << 5
	RelationComparisonResultStackDepthOverflow  RelationComparisonResult = 1 << 6
	RelationComparisonResultReportsMask                                  = RelationComparisonResultReportsUnmeasurable | RelationComparisonResultReportsUnreliable
	RelationComparisonResultOverflow                                     = RelationComparisonResultComplexityOverflow | RelationComparisonResultStackDepthOverflow
)

type DiagnosticAndArguments struct {
	message   *diagnostics.Message
	arguments []any
}

type ErrorOutputContainer struct {
	errors      []*ast.Diagnostic
	skipLogging bool
}

type ErrorReporter func(message *diagnostics.Message, args ...any)

type RecursionIdKind uint32

const (
	RecursionIdKindNode RecursionIdKind = iota
	RecursionIdKindSymbol
	RecursionIdKindType
)

type RecursionId struct {
	kind RecursionIdKind
	id   uint32
}

type Relation struct {
	results map[string]RelationComparisonResult
}

func (r *Relation) get(key string) RelationComparisonResult {
	return r.results[key]
}

func (r *Relation) set(key string, result RelationComparisonResult) {
	if r.results == nil {
		r.results = make(map[string]RelationComparisonResult)
	}
	r.results[key] = result
}

func (r *Relation) size() int {
	return len(r.results)
}

func (c *Checker) compareTypesIdentical(source *Type, target *Type) Ternary {
	if c.isTypeRelatedTo(source, target, c.identityRelation) {
		return TernaryTrue
	}
	return TernaryFalse
}

func (c *Checker) isTypeAssignableTo(source *Type, target *Type) bool {
	return c.isTypeRelatedTo(source, target, c.assignableRelation)
}

func (c *Checker) isTypeSubtypeOf(source *Type, target *Type) bool {
	return c.isTypeRelatedTo(source, target, c.subtypeRelation)
}

func (c *Checker) isTypeStrictSubtypeOf(source *Type, target *Type) bool {
	return c.isTypeRelatedTo(source, target, c.strictSubtypeRelation)
}

func (c *Checker) isTypeRelatedTo(source *Type, target *Type, relation *Relation) bool {
	if isFreshLiteralType(source) {
		source = source.AsLiteralType().regularType
	}
	if isFreshLiteralType(target) {
		target = target.AsLiteralType().regularType
	}
	if source == target {
		return true
	}
	if relation != c.identityRelation {
		if relation == c.comparableRelation && target.flags&TypeFlagsNever == 0 && c.isSimpleTypeRelatedTo(target, source, relation, nil) || c.isSimpleTypeRelatedTo(source, target, relation, nil) {
			return true
		}
	} else if !((source.flags|target.flags)&(TypeFlagsUnionOrIntersection|TypeFlagsIndexedAccess|TypeFlagsConditional|TypeFlagsSubstitution) != 0) {
		// We have excluded types that may simplify to other forms, so types must have identical flags
		if source.flags != target.flags {
			return false
		}
		if source.flags&TypeFlagsSingleton != 0 {
			return true
		}
	}
	if source.flags&TypeFlagsObject != 0 && target.flags&TypeFlagsObject != 0 {
		related := relation.get(getRelationKey(source, target, IntersectionStateNone, relation == c.identityRelation, false))
		if related != RelationComparisonResultNone {
			return related&RelationComparisonResultSucceeded != 0
		}
	}
	if source.flags&TypeFlagsStructuredOrInstantiable != 0 || target.flags&TypeFlagsStructuredOrInstantiable != 0 {
		return c.checkTypeRelatedTo(source, target, relation, nil /*errorNode*/)
	}
	return false
}

func (c *Checker) isSimpleTypeRelatedTo(source *Type, target *Type, relation *Relation, errorReporter ErrorReporter) bool {
	s := source.flags
	t := target.flags
	if t&TypeFlagsAny != 0 || s&TypeFlagsNever != 0 || source == c.wildcardType {
		return true
	}
	if t&TypeFlagsUnknown != 0 && !(relation == c.strictSubtypeRelation && s&TypeFlagsAny != 0) {
		return true
	}
	if t&TypeFlagsNever != 0 {
		return false
	}
	if s&TypeFlagsStringLike != 0 && t&TypeFlagsString != 0 {
		return true
	}
	if s&TypeFlagsStringLiteral != 0 && s&TypeFlagsEnumLiteral != 0 && t&TypeFlagsStringLiteral != 0 && t&TypeFlagsEnumLiteral == 0 && source.AsLiteralType().value == target.AsLiteralType().value {
		return true
	}
	if s&TypeFlagsNumberLike != 0 && t&TypeFlagsNumber != 0 {
		return true
	}
	if s&TypeFlagsNumberLiteral != 0 && s&TypeFlagsEnumLiteral != 0 && t&TypeFlagsNumberLiteral != 0 && t&TypeFlagsEnumLiteral == 0 && source.AsLiteralType().value == target.AsLiteralType().value {
		return true
	}
	if s&TypeFlagsBigIntLike != 0 && t&TypeFlagsBigInt != 0 {
		return true
	}
	if s&TypeFlagsBooleanLike != 0 && t&TypeFlagsBoolean != 0 {
		return true
	}
	if s&TypeFlagsESSymbolLike != 0 && t&TypeFlagsESSymbol != 0 {
		return true
	}
	if s&TypeFlagsEnum != 0 && t&TypeFlagsEnum != 0 && source.symbol.Name == target.symbol.Name && c.isEnumTypeRelatedTo(source.symbol, target.symbol, errorReporter) {
		return true
	}
	if s&TypeFlagsEnumLiteral != 0 && t&TypeFlagsEnumLiteral != 0 {
		if s&TypeFlagsUnion != 0 && t&TypeFlagsUnion != 0 && c.isEnumTypeRelatedTo(source.symbol, target.symbol, errorReporter) {
			return true
		}
		if s&TypeFlagsLiteral != 0 && t&TypeFlagsLiteral != 0 && source.AsLiteralType().value == target.AsLiteralType().value && c.isEnumTypeRelatedTo(source.symbol, target.symbol, errorReporter) {
			return true
		}
	}
	// In non-strictNullChecks mode, `undefined` and `null` are assignable to anything except `never`.
	// Since unions and intersections may reduce to `never`, we exclude them here.
	if s&TypeFlagsUndefined != 0 && (!c.strictNullChecks && t&TypeFlagsUnionOrIntersection == 0 || t&(TypeFlagsUndefined|TypeFlagsVoid) != 0) {
		return true
	}
	if s&TypeFlagsNull != 0 && (!c.strictNullChecks && t&TypeFlagsUnionOrIntersection == 0 || t&TypeFlagsNull != 0) {
		return true
	}
	if s&TypeFlagsObject != 0 && t&TypeFlagsNonPrimitive != 0 && !(relation == c.strictSubtypeRelation && c.isEmptyAnonymousObjectType(source) && source.objectFlags&ObjectFlagsFreshLiteral == 0) {
		return true
	}
	if relation == c.assignableRelation || relation == c.comparableRelation {
		if s&TypeFlagsAny != 0 {
			return true
		}
		// Type number is assignable to any computed numeric enum type or any numeric enum literal type, and
		// a numeric literal type is assignable any computed numeric enum type or any numeric enum literal type
		// with a matching value. These rules exist such that enums can be used for bit-flag purposes.
		if s&TypeFlagsNumber != 0 && (t&TypeFlagsEnum != 0 || t&TypeFlagsNumberLiteral != 0 && t&TypeFlagsEnumLiteral != 0) {
			return true
		}
		if s&TypeFlagsNumberLiteral != 0 && s&TypeFlagsEnumLiteral == 0 && (t&TypeFlagsEnum != 0 || t&TypeFlagsNumberLiteral != 0 && t&TypeFlagsEnumLiteral != 0 && source.AsLiteralType().value == target.AsLiteralType().value) {
			return true
		}
		// Anything is assignable to a union containing undefined, null, and {}
		if c.isUnknownLikeUnionType(target) {
			return true
		}
	}
	return false
}

func (c *Checker) isEnumTypeRelatedTo(source *ast.Symbol, target *ast.Symbol, errorReporter ErrorReporter) bool {
	return source == target // !!!
}

func (c *Checker) checkTypeAssignableTo(source *Type, target *Type, errorNode *ast.Node, headMessage *diagnostics.Message, containingMessageChain func() *ast.MessageChain, errorOutputObject *ErrorOutputContainer) bool {
	return c.checkTypeRelatedToEx(source, target, c.assignableRelation, errorNode, headMessage, containingMessageChain, errorOutputObject)
}

func (c *Checker) checkTypeRelatedTo(source *Type, target *Type, relation *Relation, errorNode *ast.Node) bool {
	return c.checkTypeRelatedToEx(source, target, relation, errorNode, nil, nil, nil)
}

func (c *Checker) checkTypeRelatedToEx(
	source *Type,
	target *Type,
	relation *Relation,
	errorNode *ast.Node,
	headMessage *diagnostics.Message,
	containingMessageChain func() *ast.MessageChain,
	errorOutputContainer *ErrorOutputContainer,
) bool {
	r := Relater{}
	r.c = c
	r.relation = relation
	r.errorNode = errorNode
	r.relationCount = (16_000_000 - relation.size()) / 8
	result := r.isRelatedToEx(source, target, RecursionFlagsBoth, errorNode != nil /*reportErrors*/, headMessage, IntersectionStateNone)
	if r.overflow {
		// Record this relation as having failed such that we don't attempt the overflowing operation again.
		id := getRelationKey(source, target, IntersectionStateNone, relation == c.identityRelation, false /*ignoreConstraints*/)
		relation.set(id, RelationComparisonResultFailed|ifElse(r.relationCount <= 0, RelationComparisonResultComplexityOverflow, RelationComparisonResultStackDepthOverflow))
		message := ifElse(r.relationCount <= 0, diagnostics.Excessive_complexity_comparing_types_0_and_1, diagnostics.Excessive_stack_depth_comparing_types_0_and_1)
		if errorNode == nil {
			errorNode = c.currentNode
		}
		diag := c.error(errorNode, message, c.typeToString(source), c.typeToString(target))
		if errorOutputContainer != nil {
			errorOutputContainer.errors = append(errorOutputContainer.errors, diag)
		}
	} else if r.errorChain != nil {
		messageChain := createMessageChainFromErrorChain(r.errorChain)
		if containingMessageChain != nil {
			chain := containingMessageChain()
			if chain != nil {
				concatenateDiagnosticMessageChains(chain, messageChain)
			}
		}
		// !!!
		// var relatedInformation []*Diagnostic
		// // Check if we should issue an extra diagnostic to produce a quickfix for a slightly incorrect import statement
		// if headMessage != nil && errorNode != nil && result == TernaryFalse && source.symbol != nil {
		// 	links := c.getSymbolLinks(source.symbol)
		// 	if links.originatingImport && !isImportCall(links.originatingImport) {
		// 		helpfulRetry := c.checkTypeRelatedTo(c.getTypeOfSymbol(links.target), target, relation /*errorNode*/, nil)
		// 		if helpfulRetry {
		// 			// Likely an incorrect import. Issue a helpful diagnostic to produce a quickfix to change the import
		// 			diag := createDiagnosticForNode(links.originatingImport, Diagnostics.Type_originates_at_this_import_A_namespace_style_import_cannot_be_called_or_constructed_and_will_cause_a_failure_at_runtime_Consider_using_a_default_import_or_import_require_here_instead)
		// 			relatedInformation = append(relatedInformation, diag)
		// 			// Cause the error to appear with the error that triggered it
		// 		}
		// 	}
		// }
		diag := NewDiagnosticForNodeFromMessageChain(r.errorNode, messageChain).SetRelatedInfo(r.relatedInfo)
		if errorOutputContainer != nil {
			errorOutputContainer.errors = append(errorOutputContainer.errors, diag)
		}
		if errorOutputContainer == nil || !errorOutputContainer.skipLogging {
			c.diagnostics.add(diag)
		}
	}
	// !!!
	// if errorNode != nil && errorOutputContainer != nil && errorOutputContainer.skipLogging && result == TernaryFalse {
	// 	Debug.assert(!!errorOutputContainer.errors, "missed opportunity to interact with error.")
	// }
	return result != TernaryFalse
}

func createMessageChainFromErrorChain(chain *ErrorChain) *ast.MessageChain {
	for chain != nil && chain.message.ElidedInCompatabilityPyramid() {
		chain = chain.next
	}
	if chain == nil {
		return nil
	}
	return chainDiagnosticMessages(createMessageChainFromErrorChain(chain.next), chain.message, chain.args...)
}

func (c *Checker) checkTypeAssignableToAndOptionallyElaborate(source *Type, target *Type, errorNode *ast.Node, expr *ast.Node, headMessage *diagnostics.Message, containingMessageChain func() *ast.MessageChain) bool {
	return c.checkTypeRelatedToAndOptionallyElaborate(source, target, c.assignableRelation, errorNode, expr, headMessage, containingMessageChain /*errorOutputContainer*/, nil)
}

func (c *Checker) checkTypeRelatedToAndOptionallyElaborate(source *Type, target *Type, relation *Relation, errorNode *ast.Node, expr *ast.Node, headMessage *diagnostics.Message, containingMessageChain func() *ast.MessageChain, errorOutputContainer *ErrorOutputContainer) bool {
	if c.isTypeRelatedTo(source, target, relation) {
		return true
	}
	if errorNode == nil || !c.elaborateError(expr, source, target, relation, headMessage, containingMessageChain, errorOutputContainer) {
		return c.checkTypeRelatedToEx(source, target, relation, errorNode, headMessage, containingMessageChain, errorOutputContainer)
	}
	return false
}

func (c *Checker) elaborateError(node *ast.Node, source *Type, target *Type, relation *Relation, headMessage *diagnostics.Message, containingMessageChain func() *ast.MessageChain, errorOutputContainer *ErrorOutputContainer) bool {
	return false // !!!
}

// A type is 'weak' if it is an object type with at least one optional property
// and no required properties, call/construct signatures or index signatures
func (c *Checker) isWeakType(t *Type) bool {
	if t.flags&TypeFlagsObject != 0 {
		resolved := c.resolveStructuredTypeMembers(t)
		return len(resolved.signatures) == 0 && len(resolved.indexInfos) == 0 && len(resolved.properties) > 0 && core.Every(resolved.properties, func(p *ast.Symbol) bool {
			return p.Flags&ast.SymbolFlagsOptional != 0
		})
	}
	if t.flags&TypeFlagsSubstitution != 0 {
		return c.isWeakType(t.AsSubstitutionType().baseType)
	}
	if t.flags&TypeFlagsIntersection != 0 {
		return core.Every(t.Types(), c.isWeakType)
	}
	return false
}

func (c *Checker) hasCommonProperties(source *Type, target *Type, isComparingJsxAttributes bool) bool {
	for _, prop := range c.getPropertiesOfType(source) {
		if c.isKnownProperty(target, prop.Name, isComparingJsxAttributes) {
			return true
		}
	}
	return false
}

/**
 * Check if a property with the given name is known anywhere in the given type. In an object type, a property
 * is considered known if
 * 1. the object type is empty and the check is for assignability, or
 * 2. if the object type has index signatures, or
 * 3. if the property is actually declared in the object type
 *    (this means that 'toString', for example, is not usually a known property).
 * 4. In a union or intersection type,
 *    a property is considered known if it is known in any constituent type.
 * @param targetType a type to search a given name in
 * @param name a property name to search
 * @param isComparingJsxAttributes a boolean flag indicating whether we are searching in JsxAttributesType
 */
func (c *Checker) isKnownProperty(targetType *Type, name string, isComparingJsxAttributes bool) bool {
	if targetType.flags&TypeFlagsObject != 0 {
		// For backwards compatibility a symbol-named property is satisfied by a string index signature. This
		// is incorrect and inconsistent with element access expressions, where it is an error, so eventually
		// we should remove this exception.
		if c.getPropertyOfObjectType(targetType, name) != nil ||
			c.getApplicableIndexInfoForName(targetType, name) != nil ||
			isLateBoundName(name) && c.getIndexInfoOfType(targetType, c.stringType) != nil ||
			isComparingJsxAttributes && isHyphenatedJsxName(name) {
			// For JSXAttributes, if the attribute has a hyphenated name, consider that the attribute to be known.
			return true
		}
	}
	if targetType.flags&TypeFlagsSubstitution != 0 {
		return c.isKnownProperty(targetType.AsSubstitutionType().baseType, name, isComparingJsxAttributes)
	}
	if targetType.flags&TypeFlagsUnionOrIntersection != 0 && isExcessPropertyCheckTarget(targetType) {
		for _, t := range targetType.Types() {
			if c.isKnownProperty(t, name, isComparingJsxAttributes) {
				return true
			}
		}
	}
	return false
}

func isHyphenatedJsxName(name string) bool {
	return strings.Contains(name, "-")
}

func isExcessPropertyCheckTarget(t *Type) bool {
	return t.flags&TypeFlagsObject != 0 && t.objectFlags&ObjectFlagsObjectLiteralPatternWithComputedProperties == 0 ||
		t.flags&TypeFlagsNonPrimitive != 0 ||
		t.flags&TypeFlagsSubstitution != 0 && isExcessPropertyCheckTarget(t.AsSubstitutionType().baseType) ||
		t.flags&TypeFlagsUnion != 0 && core.Some(t.Types(), isExcessPropertyCheckTarget) ||
		t.flags&TypeFlagsIntersection != 0 && core.Every(t.Types(), isExcessPropertyCheckTarget)
}

// Return true if the given type is deeply nested. We consider this to be the case when the given stack contains
// maxDepth or more occurrences of types with the same recursion identity as the given type. The recursion identity
// provides a shared identity for type instantiations that repeat in some (possibly infinite) pattern. For example,
// in `type Deep<T> = { next: Deep<Deep<T>> }`, repeatedly referencing the `next` property leads to an infinite
// sequence of ever deeper instantiations with the same recursion identity (in this case the symbol associated with
// the object type literal).
// A homomorphic mapped type is considered deeply nested if its target type is deeply nested, and an intersection is
// considered deeply nested if any constituent of the intersection is deeply nested.
// It is possible, though highly unlikely, for the deeply nested check to be true in a situation where a chain of
// instantiations is not infinitely expanding. Effectively, we will generate a false positive when two types are
// structurally equal to at least maxDepth levels, but unequal at some level beyond that.
func (c *Checker) isDeeplyNestedType(t *Type, stack []*Type, maxDepth int) bool {
	if len(stack) >= maxDepth {
		if t.objectFlags&ObjectFlagsInstantiatedMapped == ObjectFlagsInstantiatedMapped {
			t = c.getMappedTargetWithSymbol(t)
		}
		if t.flags&TypeFlagsIntersection != 0 {
			for _, t := range t.Types() {
				if c.isDeeplyNestedType(t, stack, maxDepth) {
					return true
				}
			}
		}
		identity := getRecursionIdentity(t)
		count := 0
		lastTypeId := TypeId(0)
		for _, t := range stack {
			if c.hasMatchingRecursionIdentity(t, identity) {
				// We only count occurrences with a higher type id than the previous occurrence, since higher
				// type ids are an indicator of newer instantiations caused by recursion.
				if t.id >= lastTypeId {
					count++
					if count >= maxDepth {
						return true
					}
				}
				lastTypeId = t.id
			}
		}
	}
	return false
}

// Unwrap nested homomorphic mapped types and return the deepest target type that has a symbol. This better
// preserves unique type identities for mapped types applied to explicitly written object literals. For example
// in `Mapped<{ x: Mapped<{ x: Mapped<{ x: string }>}>}>`, each of the mapped type applications will have a
// unique recursion identity (that of their target object type literal) and thus avoid appearing deeply nested.
func (c *Checker) getMappedTargetWithSymbol(t *Type) *Type {
	for {
		if t.objectFlags&ObjectFlagsInstantiatedMapped == ObjectFlagsInstantiatedMapped {
			target := c.getModifiersTypeFromMappedType(t)
			if target != nil && (target.symbol != nil || target.flags&TypeFlagsIntersection != 0 &&
				core.Some(target.Types(), func(t *Type) bool { return t.symbol != nil })) {
				t = target
				continue
			}
		}
		return t
	}
}

func (c *Checker) hasMatchingRecursionIdentity(t *Type, identity RecursionId) bool {
	if t.objectFlags&ObjectFlagsInstantiatedMapped == ObjectFlagsInstantiatedMapped {
		t = c.getMappedTargetWithSymbol(t)
	}
	if t.flags&TypeFlagsIntersection != 0 {
		for _, t := range t.Types() {
			if c.hasMatchingRecursionIdentity(t, identity) {
				return true
			}
		}
	}
	return getRecursionIdentity(t) == identity
}

// The recursion identity of a type is an object identity that is shared among multiple instantiations of the type.
// We track recursion identities in order to identify deeply nested and possibly infinite type instantiations with
// the same origin. For example, when type parameters are in scope in an object type such as { x: T }, all
// instantiations of that type have the same recursion identity. The default recursion identity is the object
// identity of the type, meaning that every type is unique. Generally, types with constituents that could circularly
// reference the type have a recursion identity that differs from the object identity.
func getRecursionIdentity(t *Type) RecursionId {
	// Object and array literals are known not to contain recursive references and don't need a recursion identity.
	if t.flags&TypeFlagsObject != 0 && !isObjectOrArrayLiteralType(t) {
		if t.objectFlags&ObjectFlagsReference != 0 && t.AsTypeReference().node != nil {
			// Deferred type references are tracked through their associated AST node. This gives us finer
			// granularity than using their associated target because each manifest type reference has a
			// unique AST node.
			return RecursionId{kind: RecursionIdKindNode, id: uint32(getNodeId(t.AsTypeReference().node))}
		}
		if t.symbol != nil && !(t.objectFlags&ObjectFlagsAnonymous != 0 && t.symbol.Flags&ast.SymbolFlagsClass != 0) {
			// We track object types that have a symbol by that symbol (representing the origin of the type), but
			// exclude the static side of a class since it shares its symbol with the instance side.
			return RecursionId{kind: RecursionIdKindSymbol, id: uint32(getSymbolId(t.symbol))}
		}
		if isTupleType(t) {
			return RecursionId{kind: RecursionIdKindType, id: uint32(t.Target().id)}
		}
	}
	if t.flags&TypeFlagsTypeParameter != 0 && t.symbol != nil {
		// We use the symbol of the type parameter such that all "fresh" instantiations of that type parameter
		// have the same recursion identity.
		return RecursionId{kind: RecursionIdKindSymbol, id: uint32(getSymbolId(t.symbol))}
	}
	if t.flags&TypeFlagsIndexedAccess != 0 {
		// Identity is the leftmost object type in a chain of indexed accesses, eg, in A[P1][P2][P3] it is A.
		t = t.AsIndexedAccessType().objectType
		for t.flags&TypeFlagsIndexedAccess != 0 {
			t = t.AsIndexedAccessType().objectType
		}
		return RecursionId{kind: RecursionIdKindType, id: uint32(t.id)}
	}
	if t.flags&TypeFlagsConditional != 0 {
		// The root object represents the origin of the conditional type
		return RecursionId{kind: RecursionIdKindNode, id: uint32(getNodeId(t.AsConditionalType().root.node))}
	}
	return RecursionId{kind: RecursionIdKindType, id: uint32(t.id)}
}

func (c *Checker) getBestMatchingType(source *Type, target *Type, isRelatedTo func(source *Type, target *Type) Ternary) *Type {
	if t := c.findMatchingDiscriminantType(source, target, isRelatedTo); t != nil {
		return t
	}
	if t := c.findMatchingTypeReferenceOrTypeAliasReference(source, target); t != nil {
		return t
	}
	if t := c.findBestTypeForObjectLiteral(source, target); t != nil {
		return t
	}
	if t := c.findBestTypeForInvokable(source, target, SignatureKindCall); t != nil {
		return t
	}
	if t := c.findBestTypeForInvokable(source, target, SignatureKindConstruct); t != nil {
		return t
	}
	return c.findMostOverlappyType(source, target)
}

func (c *Checker) findMatchingTypeReferenceOrTypeAliasReference(source *Type, unionTarget *Type) *Type {
	sourceObjectFlags := source.objectFlags
	if sourceObjectFlags&(ObjectFlagsReference|ObjectFlagsAnonymous) != 0 && unionTarget.flags&TypeFlagsUnion != 0 {
		for _, target := range unionTarget.Types() {
			if target.flags&TypeFlagsObject != 0 {
				overlapObjFlags := sourceObjectFlags & target.objectFlags
				if overlapObjFlags&ObjectFlagsReference != 0 && source.Target() == target.Target() {
					return target
				}
				if overlapObjFlags&ObjectFlagsAnonymous != 0 && source.alias != nil && target.alias != nil && source.alias.symbol == target.alias.symbol {
					return target
				}
			}
		}
	}
	return nil
}

func (c *Checker) findBestTypeForInvokable(source *Type, unionTarget *Type, kind SignatureKind) *Type {
	if len(c.getSignaturesOfType(source, kind)) != 0 {
		return core.Find(unionTarget.Types(), func(t *Type) bool { return len(c.getSignaturesOfType(t, kind)) != 0 })
	}
	return nil
}

func (c *Checker) findMostOverlappyType(source *Type, unionTarget *Type) *Type {
	var bestMatch *Type
	if source.flags&(TypeFlagsPrimitive|TypeFlagsInstantiablePrimitive) == 0 {
		matchingCount := 0
		for _, target := range unionTarget.Types() {
			if target.flags&(TypeFlagsPrimitive|TypeFlagsInstantiablePrimitive) == 0 {
				overlap := c.getIntersectionType([]*Type{c.getIndexType(source), c.getIndexType(target)})
				if overlap.flags&TypeFlagsIndex != 0 {
					// perfect overlap of keys
					return target
				} else if isUnitType(overlap) || overlap.flags&TypeFlagsUnion != 0 {
					// We only want to account for literal types otherwise.
					// If we have a union of index types, it seems likely that we
					// needed to elaborate between two generic mapped types anyway.
					var length = 1
					if overlap.flags&TypeFlagsUnion != 0 {
						length = core.CountWhere(overlap.Types(), isUnitType)
					}
					if length >= matchingCount {
						bestMatch = target
						matchingCount = length
					}
				}
			}
		}
	}
	return bestMatch
}

func (c *Checker) findBestTypeForObjectLiteral(source *Type, unionTarget *Type) *Type {
	if source.objectFlags&ObjectFlagsObjectLiteral != 0 && someType(unionTarget, c.isArrayLikeType) {
		return core.Find(unionTarget.Types(), func(t *Type) bool { return !c.isArrayLikeType(t) })
	}
	return nil
}

func (c *Checker) shouldReportUnmatchedPropertyError(source *Type, target *Type) bool {
	typeCallSignatures := c.getSignaturesOfType(source, SignatureKindCall)
	typeConstructSignatures := c.getSignaturesOfType(source, SignatureKindConstruct)
	typeProperties := c.getPropertiesOfObjectType(source)
	if (len(typeCallSignatures) != 0 || len(typeConstructSignatures) != 0) && len(typeProperties) == 0 {
		if (len(c.getSignaturesOfType(target, SignatureKindCall)) != 0 && len(typeCallSignatures) != 0) ||
			len(c.getSignaturesOfType(target, SignatureKindConstruct)) != 0 && len(typeConstructSignatures) != 0 {
			// target has similar signature kinds to source, still focus on the unmatched property
			return true
		}
		return false
	}
	return true
}

func (c *Checker) getUnmatchedProperties(source *Type, target *Type, requireOptionalProperties bool, matchDiscriminantProperties bool) iter.Seq[*ast.Symbol] {
	return func(yield func(*ast.Symbol) bool) {
		properties := c.getPropertiesOfType(target)
		for _, targetProp := range properties {
			// TODO: remove this when we support static private identifier fields and find other solutions to get privateNamesAndStaticFields test to pass
			if isStaticPrivateIdentifierProperty(targetProp) {
				continue
			}
			if requireOptionalProperties || targetProp.Flags&ast.SymbolFlagsOptional == 0 && targetProp.CheckFlags&ast.CheckFlagsPartial == 0 {
				sourceProp := c.getPropertyOfType(source, targetProp.Name)
				if sourceProp == nil {
					if !yield(targetProp) {
						return
					}
				} else if matchDiscriminantProperties {
					targetType := c.getTypeOfSymbol(targetProp)
					if targetType.flags&TypeFlagsUnit != 0 {
						sourceType := c.getTypeOfSymbol(sourceProp)
						if !(sourceType.flags&TypeFlagsAny != 0 || c.getRegularTypeOfLiteralType(sourceType) == c.getRegularTypeOfLiteralType(targetType)) {
							if !yield(targetProp) {
								return
							}
						}
					}
				}
			}
		}
	}
}

func (c *Checker) getUnmatchedProperty(source *Type, target *Type, requireOptionalProperties bool, matchDiscriminantProperties bool) *ast.Symbol {
	return core.FirstOrNilSeq(c.getUnmatchedProperties(source, target, requireOptionalProperties, matchDiscriminantProperties))
}

func excludeProperties(properties []*ast.Symbol, excludedProperties core.Set[string]) []*ast.Symbol {
	if excludedProperties.Len() == 0 || len(properties) == 0 {
		return properties
	}
	var reduced []*ast.Symbol
	var excluded bool
	for i, prop := range properties {
		if !excludedProperties.Has(prop.Name) {
			if excluded {
				reduced = append(reduced, prop)
			}
		} else if !excluded {
			reduced = slices.Clip(properties[:i])
		}
	}
	if excluded {
		return reduced
	}
	return properties
}

type TypeDiscriminator struct {
	c           *Checker
	props       []*ast.Symbol
	isRelatedTo func(*Type, *Type) Ternary
}

func (d *TypeDiscriminator) len() int {
	return len(d.props)
}

func (d *TypeDiscriminator) name(index int) string {
	return d.props[index].Name
}

func (d *TypeDiscriminator) matches(index int, t *Type) bool {
	propType := d.c.getTypeOfSymbol(d.props[index])
	for _, s := range propType.Distributed() {
		if d.isRelatedTo(s, t) != TernaryFalse {
			return true
		}
	}
	return false
}

// Keep this up-to-date with the same logic within `getApparentTypeOfContextualType`, since they should behave similarly
func (c *Checker) findMatchingDiscriminantType(source *Type, target *Type, isRelatedTo func(source *Type, target *Type) Ternary) *Type {
	if target.flags&TypeFlagsUnion != 0 && source.flags&(TypeFlagsIntersection|TypeFlagsObject) != 0 {
		if match := c.getMatchingUnionConstituentForType(target, source); match != nil {
			return match
		}
		discriminantProperties := c.findDiscriminantProperties(c.getPropertiesOfType(source), target)
		discriminator := &TypeDiscriminator{c: c, props: discriminantProperties, isRelatedTo: isRelatedTo}
		discriminated := c.discriminateTypeByDiscriminableItems(target, discriminator)
		if discriminated != target {
			return discriminated
		}
	}
	return nil
}

func (c *Checker) findDiscriminantProperties(sourceProperties []*ast.Symbol, target *Type) []*ast.Symbol {
	var result []*ast.Symbol
	for _, sourceProperty := range sourceProperties {
		if c.isDiscriminantProperty(target, sourceProperty.Name) {
			result = append(result, sourceProperty)
		}
	}
	return result
}

func (c *Checker) isDiscriminantProperty(t *Type, name string) bool {
	if t != nil && t.flags&TypeFlagsUnion != 0 {
		prop := c.getUnionOrIntersectionProperty(t, name, false /*skipObjectFunctionPropertyAugment*/)
		if prop != nil && prop.CheckFlags&ast.CheckFlagsSyntheticProperty != 0 {
			if prop.CheckFlags&ast.CheckFlagsIsDiscriminantComputed == 0 {
				prop.CheckFlags |= ast.CheckFlagsIsDiscriminantComputed
				if prop.CheckFlags&ast.CheckFlagsNonUniformAndLiteral == ast.CheckFlagsNonUniformAndLiteral && !c.isGenericType(c.getTypeOfSymbol(prop)) {
					prop.CheckFlags |= ast.CheckFlagsIsDiscriminant
				}
			}
			return prop.CheckFlags&ast.CheckFlagsIsDiscriminant != 0
		}
	}
	return false
}

func (c *Checker) getMatchingUnionConstituentForType(unionType *Type, t *Type) *Type {
	keyPropertyName := c.getKeyPropertyName(unionType)
	if keyPropertyName == "" {
		return nil
	}
	propType := c.getTypeOfPropertyOfType(t, keyPropertyName)
	if propType == nil {
		return nil
	}
	return c.getConstituentTypeForKeyType(unionType, propType)
}

// Return the name of a discriminant property for which it was possible and feasible to construct a map of
// constituent types keyed by the literal types of the property by that name in each constituent type. Return
// an empty string if no such discriminant property exists.
func (c *Checker) getKeyPropertyName(t *Type) string {
	u := t.AsUnionType()
	if u.keyPropertyName == "" {
		u.keyPropertyName, u.constituentMap = c.computeKeyPropertyNameAndMap(t)
	}
	if u.keyPropertyName == InternalSymbolNameMissing {
		return ""
	}
	return u.keyPropertyName
}

// Given a union type for which getKeyPropertyName returned a non-empty string, return the constituent
// that corresponds to the given key type for that property name.
func (c *Checker) getConstituentTypeForKeyType(t *Type, keyType *Type) *Type {
	result := t.AsUnionType().constituentMap[c.getRegularTypeOfLiteralType(keyType)]
	if result != c.unknownType {
		return result
	}
	return nil
}

func (c *Checker) computeKeyPropertyNameAndMap(t *Type) (string, map[*Type]*Type) {
	types := t.Types()
	if len(types) < 10 || t.objectFlags&ObjectFlagsPrimitiveUnion != 0 || core.CountWhere(types, isObjectOrInstantiableNonPrimitive) < 10 {
		return InternalSymbolNameMissing, nil
	}
	keyPropertyName := c.getKeyPropertyCandidateName(types)
	if keyPropertyName == "" {
		return InternalSymbolNameMissing, nil
	}
	mapByKeyProperty := c.mapTypesByKeyProperty(types, keyPropertyName)
	if mapByKeyProperty == nil {
		return InternalSymbolNameMissing, nil
	}
	return keyPropertyName, mapByKeyProperty
}

func isObjectOrInstantiableNonPrimitive(t *Type) bool {
	return t.flags&(TypeFlagsObject|TypeFlagsInstantiableNonPrimitive) != 0
}

func (c *Checker) getKeyPropertyCandidateName(types []*Type) string {
	for _, t := range types {
		if t.flags&(TypeFlagsObject|TypeFlagsInstantiableNonPrimitive) != 0 {
			for _, p := range c.getPropertiesOfType(t) {
				if isUnitType(c.getTypeOfSymbol(p)) {
					return p.Name
				}
			}
		}
	}
	return ""
}

// Given a set of constituent types and a property name, create and return a map keyed by the literal
// types of the property by that name in each constituent type. No map is returned if some key property
// has a non-literal type or if less than 10 or less than 50% of the constituents have a unique key.
// Entries with duplicate keys have unknownType as the value.
func (c *Checker) mapTypesByKeyProperty(types []*Type, keyPropertyName string) map[*Type]*Type {
	typesByKey := make(map[*Type]*Type)
	count := 0
	for _, t := range types {
		if t.flags&(TypeFlagsObject|TypeFlagsIntersection|TypeFlagsInstantiableNonPrimitive) != 0 {
			discriminant := c.getTypeOfPropertyOfType(t, keyPropertyName)
			if discriminant != nil {
				if !isLiteralType(discriminant) {
					return nil
				}
				duplicate := false
				for _, d := range discriminant.Distributed() {
					key := c.getRegularTypeOfLiteralType(d)
					if existing := typesByKey[key]; existing == nil {
						typesByKey[key] = t
					} else if existing != c.unknownType {
						typesByKey[key] = c.unknownType
						duplicate = true
					}
				}
				if !duplicate {
					count++
				}
			}
		}
	}
	if count >= 10 && count*2 >= len(types) {
		return typesByKey
	}
	return nil
}

type Discriminator interface {
	len() int                        // Number of discriminant properties
	name(index int) string           // Property name of index-th discriminator
	matches(index int, t *Type) bool // True if index-th discriminator matches the given type
}

func (c *Checker) discriminateTypeByDiscriminableItems(target *Type, discriminator Discriminator) *Type {
	types := target.Types()
	include := make([]Ternary, len(types))
	for i, t := range types {
		if t.flags&TypeFlagsPrimitive == 0 {
			include[i] = TernaryTrue
		}
	}
	for n := range discriminator.len() {
		// If the remaining target types include at least one with a matching discriminant, eliminate those that
		// have non-matching discriminants. This ensures that we ignore erroneous discriminators and gradually
		// refine the target set without eliminating every constituent (which would lead to `never`).
		matched := false
		for i := range types {
			if include[i] != TernaryFalse {
				targetType := c.getTypeOfPropertyOrIndexSignatureOfType(types[i], discriminator.name(n))
				if targetType != nil && discriminator.matches(n, targetType) {
					matched = true
				} else {
					include[i] = TernaryMaybe
				}
			}
		}
		// Turn each Ternary.Maybe into Ternary.False if there was a match. Otherwise, revert to Ternary.True.
		for i := range types {
			if include[i] == TernaryMaybe {
				if matched {
					include[i] = TernaryFalse
				} else {
					include[i] = TernaryTrue
				}
			}
		}
	}
	if slices.Contains(include, TernaryFalse) {
		var filteredTypes []*Type
		for i, t := range types {
			if include[i] == TernaryTrue {
				filteredTypes = append(filteredTypes, t)
			}
		}
		filtered := c.getUnionTypeEx(filteredTypes, UnionReductionNone, nil, nil)
		if filtered.flags&TypeFlagsNever == 0 {
			return filtered
		}
	}
	return target
}

func (c *Checker) filterPrimitivesIfContainsNonPrimitive(unionType *Type) *Type {
	if c.maybeTypeOfKind(unionType, TypeFlagsNonPrimitive) {
		result := c.filterType(unionType, isNonPrimitiveType)
		if result.flags&TypeFlagsNever == 0 {
			return result
		}
	}
	return unionType
}

func isNonPrimitiveType(t *Type) bool {
	return t.flags&TypeFlagsPrimitive == 0
}

func (c *Checker) getTypeNamesForErrorDisplay(left *Type, right *Type) (string, string) {
	var leftStr string
	if c.symbolValueDeclarationIsContextSensitive(left.symbol) {
		leftStr = c.typeToStringEx(left, left.symbol.ValueDeclaration, TypeFormatFlagsNone)
	} else {
		leftStr = c.typeToString(left)
	}
	var rightStr string
	if c.symbolValueDeclarationIsContextSensitive(right.symbol) {
		rightStr = c.typeToStringEx(right, right.symbol.ValueDeclaration, TypeFormatFlagsNone)
	} else {
		rightStr = c.typeToString(right)
	}
	if leftStr == rightStr {
		leftStr = c.getTypeNameForErrorDisplay(left)
		rightStr = c.getTypeNameForErrorDisplay(right)
	}
	return leftStr, rightStr
}

func (c *Checker) getTypeNameForErrorDisplay(t *Type) string {
	return c.typeToStringEx(t, nil /*enclosingDeclaration*/, TypeFormatFlagsUseFullyQualifiedType)
}

func (c *Checker) symbolValueDeclarationIsContextSensitive(symbol *ast.Symbol) bool {
	return symbol != nil && symbol.ValueDeclaration != nil && isExpression(symbol.ValueDeclaration) && !c.isContextSensitive(symbol.ValueDeclaration)
}

func (c *Checker) typeCouldHaveTopLevelSingletonTypes(t *Type) bool {
	// Okay, yes, 'boolean' is a union of 'true | false', but that's not useful
	// in error reporting scenarios. If you need to use this function but that detail matters,
	// feel free to add a flag.
	if t.flags&TypeFlagsBoolean != 0 {
		return false
	}
	if t.flags&TypeFlagsUnionOrIntersection != 0 {
		return core.Some(t.Types(), c.typeCouldHaveTopLevelSingletonTypes)
	}
	if t.flags&TypeFlagsInstantiable != 0 {
		constraint := c.getConstraintOfType(t)
		if constraint != nil && constraint != t {
			return c.typeCouldHaveTopLevelSingletonTypes(constraint)
		}
	}
	return isUnitType(t) || t.flags&TypeFlagsTemplateLiteral != 0 || t.flags&TypeFlagsStringMapping != 0
}

func (c *Checker) getVariances(t *Type) []VarianceFlags {
	// Arrays and tuples are known to be covariant, no need to spend time computing this.
	if t == c.globalArrayType || t == c.globalReadonlyArrayType || t.objectFlags&ObjectFlagsTuple != 0 {
		return c.arrayVariances
	}
	return c.getVariancesWorker(t.symbol, t.AsInterfaceType().TypeParameters())
}

func (c *Checker) getAliasVariances(symbol *ast.Symbol) []VarianceFlags {
	return c.getVariancesWorker(symbol, c.typeAliasLinks.get(symbol).typeParameters)
}

// Return an array containing the variance of each type parameter. The variance is effectively
// a digest of the type comparisons that occur for each type argument when instantiations of the
// generic type are structurally compared. We infer the variance information by comparing
// instantiations of the generic type for type arguments with known relations. The function
// returns an empty slice when invoked recursively for the given generic type.
func (c *Checker) getVariancesWorker(symbol *ast.Symbol, typeParameters []*Type) []VarianceFlags {
	links := c.varianceLinks.get(symbol)
	if links.variances == nil {
		oldVarianceComputation := c.inVarianceComputation
		saveResolutionStart := c.resolutionStart
		if !c.inVarianceComputation {
			c.inVarianceComputation = true
			c.resolutionStart = len(c.typeResolutions)
		}
		links.variances = []VarianceFlags{}
		variances := make([]VarianceFlags, len(typeParameters))
		for i, tp := range typeParameters {
			modifiers := c.getTypeParameterModifiers(tp)
			var variance VarianceFlags
			switch {
			case modifiers&ast.ModifierFlagsOut != 0:
				if modifiers&ast.ModifierFlagsIn != 0 {
					variance = VarianceFlagsInvariant
				} else {
					variance = VarianceFlagsCovariant
				}
			case modifiers&ast.ModifierFlagsIn != 0:
				variance = VarianceFlagsContravariant
			default:
				unmeasurable := false
				unreliable := false
				oldHandler := c.outofbandVarianceMarkerHandler
				c.outofbandVarianceMarkerHandler = func(onlyUnreliable bool) {
					if onlyUnreliable {
						unreliable = true
					} else {
						unmeasurable = true
					}
				}
				// We first compare instantiations where the type parameter is replaced with
				// marker types that have a known subtype relationship. From this we can infer
				// invariance, covariance, contravariance or bivariance.
				typeWithSuper := c.createMarkerType(symbol, tp, c.markerSuperType)
				typeWithSub := c.createMarkerType(symbol, tp, c.markerSubType)
				variance = (ifElse(c.isTypeAssignableTo(typeWithSub, typeWithSuper), VarianceFlagsCovariant, 0)) |
					(ifElse(c.isTypeAssignableTo(typeWithSuper, typeWithSub), VarianceFlagsContravariant, 0))
				// If the instantiations appear to be related bivariantly it may be because the
				// type parameter is independent (i.e. it isn't witnessed anywhere in the generic
				// type). To determine this we compare instantiations where the type parameter is
				// replaced with marker types that are known to be unrelated.
				if variance == VarianceFlagsBivariant && c.isTypeAssignableTo(c.createMarkerType(symbol, tp, c.markerOtherType), typeWithSuper) {
					variance = VarianceFlagsIndependent
				}
				c.outofbandVarianceMarkerHandler = oldHandler
				if unmeasurable {
					variance |= VarianceFlagsUnmeasurable
				}
				if unreliable {
					variance |= VarianceFlagsUnreliable
				}
			}
			variances[i] = variance
		}
		if !oldVarianceComputation {
			c.inVarianceComputation = false
			c.resolutionStart = saveResolutionStart
		}
		links.variances = variances
	}
	return links.variances
}

func (c *Checker) createMarkerType(symbol *ast.Symbol, source *Type, target *Type) *Type {
	mapper := newSimpleTypeMapper(source, target)
	t := c.getDeclaredTypeOfSymbol(symbol)
	if c.isErrorType(t) {
		return t
	}
	var result *Type
	if symbol.Flags&ast.SymbolFlagsTypeAlias != 0 {
		result = c.getTypeAliasInstantiation(symbol, c.instantiateTypes(c.typeAliasLinks.get(symbol).typeParameters, mapper), nil)
	} else {
		result = c.createTypeReference(t, c.instantiateTypes(t.AsInterfaceType().TypeParameters(), mapper))
	}
	c.markerTypes.Add(result)
	return result
}

func (c *Checker) isMarkerType(t *Type) bool {
	return c.markerTypes.Has(t)
}

func (c *Checker) getTypeParameterModifiers(tp *Type) ast.ModifierFlags {
	var flags ast.ModifierFlags
	if tp.symbol != nil {
		for _, d := range tp.symbol.Declarations {
			flags |= getEffectiveModifierFlags(d)
		}
	}
	return flags & (ast.ModifierFlagsIn | ast.ModifierFlagsOut | ast.ModifierFlagsConst)
}

// Return true if the given type reference has a 'void' type argument for a covariant type parameter.
// See comment at call in recursiveTypeRelatedTo for when this case matters.
func (c *Checker) hasCovariantVoidArgument(typeArguments []*Type, variances []VarianceFlags) bool {
	for i, v := range variances {
		if v&VarianceFlagsVarianceMask == VarianceFlagsCovariant && typeArguments[i].flags&TypeFlagsVoid != 0 {
			return true
		}
	}
	return false
}

func (c *Checker) compareSignaturesRelated(source *Signature, target *Signature, checkMode SignatureCheckMode, reportErrors bool, errorReporter ErrorReporter, compareTypes TypeComparer, reportUnreliableMarkers *TypeMapper) Ternary {
	if source == target {
		return TernaryTrue
	}
	if !(checkMode&SignatureCheckModeStrictTopSignature != 0 && c.isTopSignature(source)) && c.isTopSignature(target) {
		return TernaryTrue
	}
	if checkMode&SignatureCheckModeStrictTopSignature != 0 && c.isTopSignature(source) && !c.isTopSignature(target) {
		return TernaryFalse
	}
	targetCount := c.getParameterCount(target)
	var sourceHasMoreParameters bool
	if !c.hasEffectiveRestParameter(target) {
		if checkMode&SignatureCheckModeStrictArity != 0 {
			sourceHasMoreParameters = c.hasEffectiveRestParameter(source) || c.getParameterCount(source) > targetCount
		} else {
			sourceHasMoreParameters = c.getMinArgumentCount(source) > targetCount
		}
	}
	if sourceHasMoreParameters {
		if reportErrors && (checkMode&SignatureCheckModeStrictArity == 0) {
			// the second condition should be redundant, because there is no error reporting when comparing signatures by strict arity
			// since it is only done for subtype reduction
			errorReporter(diagnostics.Target_signature_provides_too_few_arguments_Expected_0_or_more_but_got_1, c.getMinArgumentCount(source), targetCount)
		}
		return TernaryFalse
	}
	if source.typeParameters != nil && !core.Same(source.typeParameters, target.typeParameters) {
		target = c.getCanonicalSignature(target)
		source = c.instantiateSignatureInContextOf(source, target /*inferenceContext*/, nil, compareTypes)
	}
	sourceCount := c.getParameterCount(source)
	sourceRestType := c.getNonArrayRestType(source)
	targetRestType := c.getNonArrayRestType(target)
	if sourceRestType != nil || targetRestType != nil {
		c.instantiateType(ifElse(sourceRestType != nil, sourceRestType, targetRestType), reportUnreliableMarkers)
	}
	kind := ast.KindUnknown
	if target.declaration != nil {
		kind = target.declaration.Kind
	}
	strictVariance := checkMode&SignatureCheckModeCallback != 0 && c.strictFunctionTypes && kind != ast.KindMethodDeclaration && kind != ast.KindMethodSignature && kind != ast.KindConstructor
	result := TernaryTrue
	sourceThisType := c.getThisTypeOfSignature(source)
	if sourceThisType != nil && sourceThisType != c.voidType {
		targetThisType := c.getThisTypeOfSignature(target)
		if targetThisType != nil {
			// void sources are assignable to anything.
			var related Ternary
			if !strictVariance {
				related = compareTypes(sourceThisType, targetThisType, false /*reportErrors*/)
			}
			if related == TernaryFalse {
				related = compareTypes(targetThisType, sourceThisType, reportErrors)
			}
			if related == TernaryFalse {
				if reportErrors {
					errorReporter(diagnostics.The_this_types_of_each_signature_are_incompatible)
				}
				return TernaryFalse
			}
			result &= related
		}
	}
	var paramCount int
	if sourceRestType != nil || targetRestType != nil {
		paramCount = min(sourceCount, targetCount)
	} else {
		paramCount = max(sourceCount, targetCount)
	}
	var restIndex int
	if sourceRestType != nil || targetRestType != nil {
		restIndex = paramCount - 1
	} else {
		restIndex = -1
	}
	for i := range paramCount {
		var sourceType *Type
		if i == restIndex {
			sourceType = c.getRestOrAnyTypeAtPosition(source, i)
		} else {
			sourceType = c.tryGetTypeAtPosition(source, i)
		}
		var targetType *Type
		if i == restIndex {
			targetType = c.getRestOrAnyTypeAtPosition(target, i)
		} else {
			targetType = c.tryGetTypeAtPosition(target, i)
		}
		if sourceType != nil && targetType != nil && (sourceType != targetType || checkMode&SignatureCheckModeStrictArity != 0) {
			// In order to ensure that any generic type Foo<T> is at least co-variant with respect to T no matter
			// how Foo uses T, we need to relate parameters bi-variantly (given that parameters are input positions,
			// they naturally relate only contra-variantly). However, if the source and target parameters both have
			// function types with a single call signature, we know we are relating two callback parameters. In
			// that case it is sufficient to only relate the parameters of the signatures co-variantly because,
			// similar to return values, callback parameters are output positions. This means that a Promise<T>,
			// where T is used only in callback parameter positions, will be co-variant (as opposed to bi-variant)
			// with respect to T.
			var sourceSig *Signature
			var targetSig *Signature
			if checkMode&SignatureCheckModeCallback == 0 && !c.isInstantiatedGenericParameter(source, i) {
				sourceSig = c.getSingleCallSignature(c.getNonNullableType(sourceType))
			}
			if checkMode&SignatureCheckModeCallback == 0 && !c.isInstantiatedGenericParameter(target, i) {
				targetSig = c.getSingleCallSignature(c.getNonNullableType(targetType))
			}
			callbacks := sourceSig != nil && targetSig != nil && c.getTypePredicateOfSignature(sourceSig) == nil && c.getTypePredicateOfSignature(targetSig) == nil &&
				c.getTypeFacts(sourceType, TypeFactsIsUndefinedOrNull) == c.getTypeFacts(targetType, TypeFactsIsUndefinedOrNull)
			var related Ternary
			if callbacks {
				related = c.compareSignaturesRelated(targetSig, sourceSig, checkMode&SignatureCheckModeStrictArity|ifElse(strictVariance, SignatureCheckModeStrictCallback, SignatureCheckModeBivariantCallback), reportErrors, errorReporter, compareTypes, reportUnreliableMarkers)
			} else {
				if checkMode&SignatureCheckModeCallback != 0 && !strictVariance {
					related = compareTypes(sourceType, targetType, false /*reportErrors*/)
				}
				if related == TernaryFalse {
					related = compareTypes(targetType, sourceType, reportErrors)
				}
			}
			// With strict arity, (x: number | undefined) => void is a subtype of (x?: number | undefined) => void
			if related != TernaryFalse && checkMode&SignatureCheckModeStrictArity != 0 && i >= c.getMinArgumentCount(source) && i < c.getMinArgumentCount(target) && compareTypes(sourceType, targetType, false /*reportErrors*/) != TernaryFalse {
				related = TernaryFalse
			}
			if related == TernaryFalse {
				if reportErrors {
					errorReporter(diagnostics.Types_of_parameters_0_and_1_are_incompatible, c.getParameterNameAtPosition(source, i), c.getParameterNameAtPosition(target, i))
				}
				return TernaryFalse
			}
			result &= related
		}
	}
	if checkMode&SignatureCheckModeIgnoreReturnTypes == 0 {
		// If a signature resolution is already in-flight, skip issuing a circularity error
		// here and just use the `any` type directly
		targetReturnType := c.getNonCircularReturnTypeOfSignature(target)
		if targetReturnType == c.voidType || targetReturnType == c.anyType {
			return result
		}
		sourceReturnType := c.getNonCircularReturnTypeOfSignature(source)
		// The following block preserves behavior forbidding boolean returning functions from being assignable to type guard returning functions
		targetTypePredicate := c.getTypePredicateOfSignature(target)
		if targetTypePredicate != nil {
			sourceTypePredicate := c.getTypePredicateOfSignature(source)
			if sourceTypePredicate != nil {
				result &= c.compareTypePredicateRelatedTo(sourceTypePredicate, targetTypePredicate, reportErrors, errorReporter, compareTypes)
			} else if targetTypePredicate.kind == TypePredicateKindIdentifier || targetTypePredicate.kind == TypePredicateKindThis {
				if reportErrors {
					errorReporter(diagnostics.Signature_0_must_be_a_type_predicate, c.signatureToString(source))
				}
				return TernaryFalse
			}
		} else {
			// When relating callback signatures, we still need to relate return types bi-variantly as otherwise
			// the containing type wouldn't be co-variant. For example, interface Foo<T> { add(cb: () => T): void }
			// wouldn't be co-variant for T without this rule.
			var related Ternary
			if checkMode&SignatureCheckModeBivariantCallback != 0 {
				related = compareTypes(targetReturnType, sourceReturnType, false /*reportErrors*/)
			}
			if related == TernaryFalse {
				related = compareTypes(sourceReturnType, targetReturnType, reportErrors)
			}
			result &= related
			if result == TernaryFalse && reportErrors {
				// The errors reported here serve as markers that trigger error chain reduction in the (*Relater).reportError
				// method. The markers are elided in the final diagnostic chain and never actually reported.
				var message *diagnostics.Message
				if len(source.parameters) == 0 && len(target.parameters) == 0 {
					message = ifElse(source.flags&SignatureFlagsConstruct != 0,
						diagnostics.Construct_signatures_with_no_arguments_have_incompatible_return_types_0_and_1,
						diagnostics.Call_signatures_with_no_arguments_have_incompatible_return_types_0_and_1)
				} else {
					message = ifElse(source.flags&SignatureFlagsConstruct != 0,
						diagnostics.Construct_signature_return_types_0_and_1_are_incompatible,
						diagnostics.Call_signature_return_types_0_and_1_are_incompatible)
				}
				errorReporter(message, c.typeToString(sourceReturnType), c.typeToString(targetReturnType))
			}
		}
	}
	return result
}

func (c *Checker) compareTypePredicateRelatedTo(source *TypePredicate, target *TypePredicate, reportErrors bool, errorReporter ErrorReporter, compareTypes TypeComparer) Ternary {
	if source.kind != target.kind {
		if reportErrors {
			errorReporter(diagnostics.A_this_based_type_guard_is_not_compatible_with_a_parameter_based_type_guard)
			errorReporter(diagnostics.Type_predicate_0_is_not_assignable_to_1, c.typePredicateToString(source), c.typePredicateToString(target))
		}
		return TernaryFalse
	}
	if source.kind == TypePredicateKindIdentifier || source.kind == TypePredicateKindAssertsIdentifier {
		if source.parameterIndex != target.parameterIndex {
			if reportErrors {
				errorReporter(diagnostics.Parameter_0_is_not_in_the_same_position_as_parameter_1, source.parameterName, target.parameterName)
				errorReporter(diagnostics.Type_predicate_0_is_not_assignable_to_1, c.typePredicateToString(source), c.typePredicateToString(target))
			}
			return TernaryFalse
		}
	}
	var related Ternary
	switch {
	case source.t == target.t:
		related = TernaryTrue
	case source.t != nil && target.t != nil:
		related = compareTypes(source.t, target.t, reportErrors)
	default:
		related = TernaryFalse
	}
	if related == TernaryFalse && reportErrors {
		errorReporter(diagnostics.Type_predicate_0_is_not_assignable_to_1, c.typePredicateToString(source), c.typePredicateToString(target))
	}
	return related
}

// Returns true if `s` is `(...args: A) => R` where `A` is `any`, `any[]`, `never`, or `never[]`, and `R` is `any` or `unknown`.
func (c *Checker) isTopSignature(s *Signature) bool {
	if s.typeParameters == nil && (s.thisParameter == nil || isTypeAny(c.getTypeOfParameter(s.thisParameter))) && len(s.parameters) == 1 && signatureHasRestParameter(s) {
		paramType := c.getTypeOfParameter(s.parameters[0])
		var restType *Type
		if c.isArrayType(paramType) {
			restType = c.getTypeArguments(paramType)[0]
		} else {
			restType = paramType
		}
		return restType.flags&(TypeFlagsAny|TypeFlagsNever) != 0 && c.getReturnTypeOfSignature(s).flags&TypeFlagsAnyOrUnknown != 0
	}
	return false
}

// Return the number of parameters in a signature. The rest parameter, if present, counts as one
// parameter. For example, the parameter count of (x: number, y: number, ...z: string[]) is 3 and
// the parameter count of (x: number, ...args: [number, ...string[], boolean])) is also 3. In the
// latter example, the effective rest type is [...string[], boolean].
func (c *Checker) getParameterCount(signature *Signature) int {
	length := len(signature.parameters)
	if signatureHasRestParameter(signature) {
		restType := c.getTypeOfSymbol(signature.parameters[length-1])
		if isTupleType(restType) {
			return length + restType.TargetTupleType().fixedLength - ifElse(restType.TargetTupleType().combinedFlags&ElementFlagsVariable != 0, 0, 1)
		}
	}
	return length
}

func (c *Checker) getMinArgumentCount(signature *Signature) int {
	return c.getMinArgumentCountEx(signature, MinArgumentCountFlagsNone)
}

func (c *Checker) getMinArgumentCountEx(signature *Signature, flags MinArgumentCountFlags) int {
	strongArityForUntypedJS := flags & MinArgumentCountFlagsStrongArityForUntypedJS
	voidIsNonOptional := flags & MinArgumentCountFlagsVoidIsNonOptional
	if voidIsNonOptional != 0 || signature.resolvedMinArgumentCount == -1 {
		minArgumentCount := -1
		if signatureHasRestParameter(signature) {
			restType := c.getTypeOfSymbol(signature.parameters[len(signature.parameters)-1])
			if isTupleType(restType) {
				firstOptionalIndex := core.FindIndex(restType.TargetTupleType().elementInfos, func(info TupleElementInfo) bool {
					return info.flags&ElementFlagsRequired == 0
				})
				requiredCount := firstOptionalIndex
				if firstOptionalIndex < 0 {
					requiredCount = restType.TargetTupleType().fixedLength
				}
				if requiredCount > 0 {
					minArgumentCount = len(signature.parameters) - 1 + requiredCount
				}
			}
		}
		if minArgumentCount == -1 {
			if strongArityForUntypedJS == 0 && signature.flags&SignatureFlagsIsUntypedSignatureInJSFile != 0 {
				return 0
			}
			minArgumentCount = int(signature.minArgumentCount)
		}
		if voidIsNonOptional != 0 {
			return minArgumentCount
		}
		for i := minArgumentCount - 1; i >= 0; i-- {
			t := c.getTypeAtPosition(signature, i)
			if !someType(t, func(t *Type) bool { return t.flags&TypeFlagsVoid != 0 }) {
				break
			}
			minArgumentCount = i
		}
		signature.resolvedMinArgumentCount = int32(minArgumentCount)
	}
	return int(signature.resolvedMinArgumentCount)
}

func (c *Checker) hasEffectiveRestParameter(signature *Signature) bool {
	if signatureHasRestParameter(signature) {
		restType := c.getTypeOfSymbol(signature.parameters[len(signature.parameters)-1])
		return !isTupleType(restType) || restType.TargetTupleType().combinedFlags&ElementFlagsVariable != 0
	}
	return false
}

func (c *Checker) getTypeAtPosition(signature *Signature, pos int) *Type {
	t := c.tryGetTypeAtPosition(signature, pos)
	if t != nil {
		return t
	}
	return c.anyType
}

func (c *Checker) tryGetTypeAtPosition(signature *Signature, pos int) *Type {
	paramCount := len(signature.parameters) - ifElse(signatureHasRestParameter(signature), 1, 0)
	if pos < paramCount {
		return c.getTypeOfParameter(signature.parameters[pos])
	}
	if signatureHasRestParameter(signature) {
		// We want to return the value undefined for an out of bounds parameter position,
		// so we need to check bounds here before calling getIndexedAccessType (which
		// otherwise would return the type 'undefined').
		restType := c.getTypeOfSymbol(signature.parameters[paramCount])
		index := pos - paramCount
		if !isTupleType(restType) || restType.TargetTupleType().combinedFlags&ElementFlagsVariable != 0 || index < restType.TargetTupleType().fixedLength {
			return c.getIndexedAccessType(restType, c.getNumberLiteralType(float64(index)))
		}
	}
	return nil
}

// Return the rest type at the given position, transforming `any[]` into just `any`. We do this because
// in signatures we want `any[]` in a rest position to be compatible with anything, but `any[]` isn't
// assignable to tuple types with required elements.
func (c *Checker) getRestOrAnyTypeAtPosition(source *Signature, pos int) *Type {
	restType := c.getRestTypeAtPosition(source, pos, false)
	if restType != nil {
		if elementType := c.getElementTypeOfArrayType(restType); elementType != nil && isTypeAny(elementType) {
			return c.anyType
		}
	}
	return restType
}

func (c *Checker) getRestTypeAtPosition(source *Signature, pos int, readonly bool) *Type {
	parameterCount := c.getParameterCount(source)
	minArgumentCount := c.getMinArgumentCount(source)
	restType := c.getEffectiveRestType(source)
	if restType != nil && pos >= parameterCount-1 {
		if pos == parameterCount-1 {
			return restType
		} else {
			return c.createArrayType(c.getIndexedAccessType(restType, c.numberType))
		}
	}
	types := make([]*Type, parameterCount)
	infos := make([]TupleElementInfo, parameterCount)
	for i := range parameterCount {
		var flags ElementFlags
		if restType == nil || i < parameterCount-1 {
			types[i] = c.getTypeAtPosition(source, i)
			flags = ifElse(i < minArgumentCount, ElementFlagsRequired, ElementFlagsOptional)
		} else {
			types[i] = restType
			flags = ElementFlagsVariadic
		}
		infos[i] = TupleElementInfo{flags: flags, labeledDeclaration: c.getNameableDeclarationAtPosition(source, i)}
	}
	return c.createTupleTypeEx(types, infos, readonly)
}

func (c *Checker) getNameableDeclarationAtPosition(signature *Signature, pos int) *ast.Node {
	paramCount := len(signature.parameters) - ifElse(signatureHasRestParameter(signature), 1, 0)
	if pos < paramCount {
		decl := signature.parameters[pos].ValueDeclaration
		if decl != nil && c.isValidDeclarationForTupleLabel(decl) {
			return decl
		}
		return nil
	}
	if signatureHasRestParameter(signature) {
		restParameter := signature.parameters[paramCount]
		restType := c.getTypeOfSymbol(restParameter)
		if isTupleType(restType) {
			elementInfos := restType.TargetTupleType().elementInfos
			index := pos - paramCount
			if index < len(elementInfos) {
				return elementInfos[index].labeledDeclaration
			}
			return nil
		}
		if restParameter.ValueDeclaration != nil && c.isValidDeclarationForTupleLabel(restParameter.ValueDeclaration) {
			return restParameter.ValueDeclaration
		}
	}
	return nil
}

func (c *Checker) isValidDeclarationForTupleLabel(d *ast.Node) bool {
	return ast.IsNamedTupleMember(d) || ast.IsParameter(d) && d.Name() != nil && ast.IsIdentifier(d.Name())
}

func (c *Checker) getNonArrayRestType(signature *Signature) *Type {
	restType := c.getEffectiveRestType(signature)
	if restType != nil && !c.isArrayType(restType) && isTypeAny(restType) {
		return restType
	}
	return nil
}

func (c *Checker) getEffectiveRestType(signature *Signature) *Type {
	if signatureHasRestParameter(signature) {
		restType := c.getTypeOfSymbol(signature.parameters[len(signature.parameters)-1])
		if !isTupleType(restType) {
			if isTypeAny(restType) {
				return c.anyArrayType
			}
			return restType
		}
		if restType.TargetTupleType().combinedFlags&ElementFlagsVariable != 0 {
			return c.sliceTupleType(restType, restType.TargetTupleType().fixedLength, 0)
		}
	}
	return nil
}

func (c *Checker) sliceTupleType(t *Type, index int, endSkipCount int) *Type {
	target := t.TargetTupleType()
	endIndex := c.getTypeReferenceArity(t) - endSkipCount
	if index > target.fixedLength {
		if restArrayType := c.getRestArrayTypeOfTupleType(t); restArrayType != nil {
			return restArrayType
		}
		return c.createTupleType(nil)
	}
	return c.createTupleTypeEx(c.getTypeArguments(t)[index:endIndex], target.elementInfos[index:endIndex], false /*readonly*/)
}

func (c *Checker) getKnownKeysOfTupleType(t *Type) *Type {
	fixedLength := t.TargetTupleType().fixedLength
	keys := make([]*Type, fixedLength+1)
	for i := range fixedLength {
		keys[i] = c.getStringLiteralType(strconv.Itoa(i))
	}
	keys[fixedLength] = c.getIndexType(ifElse(t.TargetTupleType().readonly, c.globalReadonlyArrayType, c.globalArrayType))
	return c.getUnionType(keys)
}

func (c *Checker) getRestArrayTypeOfTupleType(t *Type) *Type {
	if restType := c.getRestTypeOfTupleType(t); restType != nil {
		return c.createArrayType(restType)
	}
	return nil
}

func (c *Checker) getThisTypeOfSignature(signature *Signature) *Type {
	if signature.thisParameter != nil {
		return c.getTypeOfSymbol(signature.thisParameter)
	}
	return nil
}

func (c *Checker) isInstantiatedGenericParameter(signature *Signature, pos int) bool {
	if signature.target == nil {
		return false
	}
	t := c.tryGetTypeAtPosition(signature.target, pos)
	return t != nil && c.isGenericType(t)
}

func (c *Checker) getParameterNameAtPosition(signature *Signature, pos int) string {
	paramCount := len(signature.parameters) - ifElse(signatureHasRestParameter(signature), 1, 0)
	if pos < paramCount {
		return signature.parameters[pos].Name
	}
	restParameter := signature.parameters[paramCount]
	restType := c.getTypeOfSymbol(restParameter)
	if isTupleType(restType) {
		index := pos - paramCount
		c.getTupleElementLabel(restType.TargetTupleType().elementInfos[index], restParameter, index)
	}
	return restParameter.Name
}

func (c *Checker) getTupleElementLabel(elementInfo TupleElementInfo, restSymbol *ast.Symbol, index int) string {
	if elementInfo.labeledDeclaration != nil {
		return elementInfo.labeledDeclaration.Name().Text()
	}
	if restSymbol != nil && restSymbol.ValueDeclaration != nil && ast.IsParameter(restSymbol.ValueDeclaration) {
		return c.getTupleElementLabelFromBindingElement(restSymbol.ValueDeclaration, index, elementInfo.flags)
	}
	var rootName string
	if restSymbol != nil {
		rootName = restSymbol.Name
	} else {
		rootName = "arg"
	}
	return rootName + "_" + strconv.Itoa(index)
}

func (c *Checker) getTupleElementLabelFromBindingElement(node *ast.Node, index int, elementFlags ElementFlags) string {
	// !!! Extract from parameter or binding element
	return "arg_" + strconv.Itoa(index)
}

func (c *Checker) getTypePredicateOfSignature(sig *Signature) *TypePredicate {
	if sig.resolvedTypePredicate == nil {
		switch {
		case sig.target != nil:
			targetTypePredicate := c.getTypePredicateOfSignature(sig.target)
			if targetTypePredicate != nil {
				sig.resolvedTypePredicate = c.instantiateTypePredicate(targetTypePredicate, sig.mapper)
			}
		case sig.composite != nil:
			sig.resolvedTypePredicate = c.getUnionOrIntersectionTypePredicate(sig.composite.signatures, sig.composite.flags)
		default:
			var typeNode *ast.TypeNode
			if sig.declaration != nil {
				typeNode = getEffectiveTypeAnnotationNode(sig.declaration)
				switch {
				case typeNode != nil && ast.IsTypePredicateNode(typeNode):
					sig.resolvedTypePredicate = c.createTypePredicateFromTypePredicateNode(typeNode, sig)
				case isFunctionLikeDeclaration(sig.declaration) && (sig.resolvedReturnType == nil || sig.resolvedReturnType.flags&TypeFlagsBoolean != 0) && c.getParameterCount(sig) > 0:
					sig.resolvedTypePredicate = c.noTypePredicate // avoid infinite loop
					sig.resolvedTypePredicate = c.getTypePredicateFromBody(sig.declaration)
				}
			}
		}
		if sig.resolvedTypePredicate == nil {
			sig.resolvedTypePredicate = c.noTypePredicate
		}
	}
	if sig.resolvedTypePredicate == c.noTypePredicate {
		return nil
	}
	return sig.resolvedTypePredicate
}

func (c *Checker) getUnionOrIntersectionTypePredicate(signatures []*Signature, flags TypeFlags) *TypePredicate {
	var last *TypePredicate
	var types []*Type
	for _, sig := range signatures {
		pred := c.getTypePredicateOfSignature(sig)
		if pred != nil {
			// Constituent type predicates must all have matching kinds. We don't create composite type predicates for assertions.
			if pred.kind != TypePredicateKindThis && pred.kind != TypePredicateKindIdentifier || last != nil && !c.typePredicateKindsMatch(last, pred) {
				return nil
			}
			last = pred
			types = append(types, pred.t)
		} else {
			// In composite union signatures we permit and ignore signatures with a return type `false`.
			var returnType *Type
			if flags&TypeFlagsUnion != 0 {
				returnType = c.getReturnTypeOfSignature(sig)
			}
			if returnType != c.falseType && returnType != c.regularFalseType {
				return nil
			}
		}
	}
	if last == nil {
		return nil
	}
	compositeType := c.getUnionOrIntersectionType(types, flags, UnionReductionLiteral)
	return c.newTypePredicate(last.kind, last.parameterName, last.parameterIndex, compositeType)
}

func (c *Checker) typePredicateKindsMatch(a *TypePredicate, b *TypePredicate) bool {
	return a.kind == b.kind && a.parameterIndex == b.parameterIndex
}

func (c *Checker) createTypePredicateFromTypePredicateNode(node *ast.Node, signature *Signature) *TypePredicate {
	predicateNode := node.AsTypePredicateNode()
	var t *Type
	if predicateNode.TypeNode != nil {
		t = c.getTypeFromTypeNode(predicateNode.TypeNode)
	}
	if ast.IsThisTypeNode(predicateNode.ParameterName) {
		kind := ifElse(predicateNode.AssertsModifier != nil, TypePredicateKindAssertsThis, TypePredicateKindThis)
		return c.newTypePredicate(kind, "" /*parameterName*/, 0 /*parameterIndex*/, t)
	}
	kind := ifElse(predicateNode.AssertsModifier != nil, TypePredicateKindAssertsIdentifier, TypePredicateKindIdentifier)
	name := predicateNode.ParameterName.Text()
	index := core.FindIndex(signature.parameters, func(p *ast.Symbol) bool { return p.Name == name })
	return c.newTypePredicate(kind, name, int32(index), t)
}

func (c *Checker) instantiateTypePredicate(predicate *TypePredicate, mapper *TypeMapper) *TypePredicate {
	t := c.instantiateType(predicate.t, mapper)
	if t == predicate.t {
		return predicate
	}
	return c.newTypePredicate(predicate.kind, predicate.parameterName, predicate.parameterIndex, t)
}

func (c *Checker) newTypePredicate(kind TypePredicateKind, parameterName string, parameterIndex int32, t *Type) *TypePredicate {
	return &TypePredicate{kind: kind, parameterIndex: parameterIndex, parameterName: parameterName, t: t}
}

func (c *Checker) isResolvingReturnTypeOfSignature(signature *Signature) bool {
	if signature.composite != nil && core.Some(signature.composite.signatures, c.isResolvingReturnTypeOfSignature) {
		return true
	}
	return signature.resolvedReturnType == nil && c.findResolutionCycleStartIndex(signature, TypeSystemPropertyNameResolvedReturnType) >= 0
}

/**
 * See signatureRelatedTo, compareSignaturesIdentical
 */
func (c *Checker) compareSignaturesIdentical(source *Signature, target *Signature, partialMatch bool, ignoreThisTypes bool, ignoreReturnTypes bool, compareTypes func(s *Type, t *Type) Ternary) Ternary {
	if source == target {
		return TernaryTrue
	}
	if !c.isMatchingSignature(source, target, partialMatch) {
		return TernaryFalse
	}
	// Check that the two signatures have the same number of type parameters.
	if len(source.typeParameters) != len(target.typeParameters) {
		return TernaryFalse
	}
	// Check that type parameter constraints and defaults match. If they do, instantiate the source
	// signature with the type parameters of the target signature and continue the comparison.
	if target.typeParameters != nil {
		mapper := newTypeMapper(source.typeParameters, target.typeParameters)
		for i := range len(target.typeParameters) {
			s := source.typeParameters[i]
			t := target.typeParameters[i]
			if !(s == t || compareTypes(c.instantiateType(c.getConstraintOrUnknownFromTypeParameter(s), mapper), c.getConstraintOrUnknownFromTypeParameter(t)) != TernaryFalse &&
				compareTypes(c.instantiateType(c.getDefaultOrUnknownFromTypeParameter(s), mapper), c.getDefaultOrUnknownFromTypeParameter(t)) != TernaryFalse) {
				return TernaryFalse
			}
		}
		source = c.instantiateSignatureEx(source, mapper, true /*eraseTypeParameters*/)
	}
	result := TernaryTrue
	if !ignoreThisTypes {
		sourceThisType := c.getThisTypeOfSignature(source)
		if sourceThisType != nil {
			targetThisType := c.getThisTypeOfSignature(target)
			if targetThisType != nil {
				related := compareTypes(sourceThisType, targetThisType)
				if related == TernaryFalse {
					return TernaryFalse
				}
				result &= related
			}
		}
	}
	for i := range c.getParameterCount(target) {
		s := c.getTypeAtPosition(source, i)
		t := c.getTypeAtPosition(target, i)
		related := compareTypes(t, s)
		if related == TernaryFalse {
			return TernaryFalse
		}
		result &= related
	}
	if !ignoreReturnTypes {
		sourceTypePredicate := c.getTypePredicateOfSignature(source)
		targetTypePredicate := c.getTypePredicateOfSignature(target)
		if sourceTypePredicate != nil || targetTypePredicate != nil {
			result &= c.compareTypePredicatesIdentical(sourceTypePredicate, targetTypePredicate, compareTypes)
		} else {
			result &= compareTypes(c.getReturnTypeOfSignature(source), c.getReturnTypeOfSignature(target))
		}
	}
	return result
}

func (c *Checker) isMatchingSignature(source *Signature, target *Signature, partialMatch bool) bool {
	sourceParameterCount := c.getParameterCount(source)
	targetParameterCount := c.getParameterCount(target)
	sourceMinArgumentCount := c.getMinArgumentCount(source)
	targetMinArgumentCount := c.getMinArgumentCount(target)
	sourceHasRestParameter := c.hasEffectiveRestParameter(source)
	targetHasRestParameter := c.hasEffectiveRestParameter(target)
	// A source signature matches a target signature if the two signatures have the same number of required,
	// optional, and rest parameters.
	if sourceParameterCount == targetParameterCount && sourceMinArgumentCount == targetMinArgumentCount && sourceHasRestParameter == targetHasRestParameter {
		return true
	}
	// A source signature partially matches a target signature if the target signature has no fewer required
	// parameters
	if partialMatch && sourceMinArgumentCount <= targetMinArgumentCount {
		return true
	}
	return false
}

func (c *Checker) compareTypePredicatesIdentical(source *TypePredicate, target *TypePredicate, compareTypes func(s *Type, t *Type) Ternary) Ternary {
	switch {
	case source == nil || target == nil || !c.typePredicateKindsMatch(source, target):
		return TernaryFalse
	case source.t == target.t:
		return TernaryTrue
	case source.t != nil && target.t != nil:
		return compareTypes(source.t, target.t)
	}
	return TernaryFalse
}

func (c *Checker) getEffectiveConstraintOfIntersection(types []*Type, targetIsUnion bool) *Type {
	var constraints []*Type
	hasDisjointDomainType := false
	for _, t := range types {
		if t.flags&TypeFlagsInstantiable != 0 {
			// We keep following constraints as long as we have an instantiable type that is known
			// not to be circular or infinite (hence we stop on index access types).
			constraint := c.getConstraintOfType(t)
			for constraint != nil && constraint.flags&(TypeFlagsTypeParameter|TypeFlagsIndex|TypeFlagsConditional) != 0 {
				constraint = c.getConstraintOfType(constraint)
			}
			if constraint != nil {
				constraints = append(constraints, constraint)
				if targetIsUnion {
					constraints = append(constraints, t)
				}
			}
		} else if t.flags&TypeFlagsDisjointDomains != 0 || c.isEmptyAnonymousObjectType(t) {
			hasDisjointDomainType = true
		}
	}
	// If the target is a union type or if we are intersecting with types belonging to one of the
	// disjoint domains, we may end up producing a constraint that hasn't been examined before.
	if constraints != nil && (targetIsUnion || hasDisjointDomainType) {
		if hasDisjointDomainType {
			// We add any types belong to one of the disjoint domains because they might cause the final
			// intersection operation to reduce the union constraints.
			for _, t := range types {
				if t.flags&TypeFlagsDisjointDomains != 0 || c.isEmptyAnonymousObjectType(t) {
					constraints = append(constraints, t)
				}
			}
		}
		// The source types were normalized; ensure the result is normalized too.
		return c.getNormalizedType(c.getIntersectionTypeEx(constraints, IntersectionFlagsNoConstraintReduction, nil), false /*writing*/)
	}
	return nil
}

func visibilityToString(flags ast.ModifierFlags) string {
	if flags == ast.ModifierFlagsPrivate {
		return "private"
	}
	if flags == ast.ModifierFlagsProtected {
		return "protected"
	}
	return "public"
}

type errorState struct {
	errorChain  *ErrorChain
	relatedInfo []*ast.Diagnostic
}

type ErrorChain struct {
	next    *ErrorChain
	message *diagnostics.Message
	args    []any
}

type Relater struct {
	c              *Checker
	relation       *Relation
	errorNode      *ast.Node
	errorChain     *ErrorChain
	relatedInfo    []*ast.Diagnostic
	maybeKeys      []string
	maybeKeysSet   core.Set[string]
	sourceStack    []*Type
	targetStack    []*Type
	maybeCount     int
	sourceDepth    int
	targetDepth    int
	expandingFlags ExpandingFlags
	overflow       bool
	relationCount  int
}

func (r *Relater) isRelatedToSimple(source *Type, target *Type) Ternary {
	return r.isRelatedToEx(source, target, RecursionFlagsNone, false /*reportErrors*/, nil /*headMessage*/, IntersectionStateNone)
}

func (r *Relater) isRelatedTo(source *Type, target *Type, recursionFlags RecursionFlags, reportErrors bool) Ternary {
	return r.isRelatedToEx(source, target, recursionFlags, reportErrors, nil, IntersectionStateNone)
}

func (r *Relater) isRelatedToEx(originalSource *Type, originalTarget *Type, recursionFlags RecursionFlags, reportErrors bool, headMessage *diagnostics.Message, intersectionState IntersectionState) Ternary {
	if originalSource == originalTarget {
		return TernaryTrue
	}
	// Before normalization: if `source` is type an object type, and `target` is primitive,
	// skip all the checks we don't need and just return `isSimpleTypeRelatedTo` result
	if originalSource.flags&TypeFlagsObject != 0 && originalTarget.flags&TypeFlagsPrimitive != 0 {
		if r.relation == r.c.comparableRelation && originalTarget.flags&TypeFlagsNever == 0 && r.c.isSimpleTypeRelatedTo(originalTarget, originalSource, r.relation, nil) ||
			r.c.isSimpleTypeRelatedTo(originalSource, originalTarget, r.relation, ifElse(reportErrors, r.reportError, nil)) {
			return TernaryTrue
		}
		if reportErrors {
			r.reportErrorResults(originalSource, originalTarget, originalSource, originalTarget, headMessage)
		}
		return TernaryFalse
	}
	// Normalize the source and target types: Turn fresh literal types into regular literal types,
	// turn deferred type references into regular type references, simplify indexed access and
	// conditional types, and resolve substitution types to either the substitution (on the source
	// side) or the type variable (on the target side).
	source := r.c.getNormalizedType(originalSource, false /*writing*/)
	target := r.c.getNormalizedType(originalTarget, true /*writing*/)
	if source == target {
		return TernaryTrue
	}
	if r.relation == r.c.identityRelation {
		if source.flags != target.flags {
			return TernaryFalse
		}
		if source.flags&TypeFlagsSingleton != 0 {
			return TernaryTrue
		}
		return r.recursiveTypeRelatedTo(source, target, false /*reportErrors*/, IntersectionStateNone, recursionFlags)
	}
	// We fastpath comparing a type parameter to exactly its constraint, as this is _super_ common,
	// and otherwise, for type parameters in large unions, causes us to need to compare the union to itself,
	// as we break down the _target_ union first, _then_ get the source constraint - so for every
	// member of the target, we attempt to find a match in the source. This avoids that in cases where
	// the target is exactly the constraint.
	if source.flags&TypeFlagsTypeParameter != 0 && r.c.getConstraintOfType(source) == target {
		return TernaryTrue
	}
	// See if we're relating a definitely non-nullable type to a union that includes null and/or undefined
	// plus a single non-nullable type. If so, remove null and/or undefined from the target type.
	if source.flags&TypeFlagsDefinitelyNonNullable != 0 && target.flags&TypeFlagsUnion != 0 {
		types := target.Types()
		var candidate *Type
		switch {
		case len(types) == 2 && types[0].flags&TypeFlagsNullable != 0:
			candidate = types[1]
		case len(types) == 3 && types[0].flags&TypeFlagsNullable != 0 && types[1].flags&TypeFlagsNullable != 0:
			candidate = types[2]
		}
		if candidate != nil && candidate.flags&TypeFlagsNullable == 0 {
			target = r.c.getNormalizedType(candidate /*writing*/, true)
			if source == target {
				return TernaryTrue
			}
		}
	}
	if r.relation == r.c.comparableRelation && target.flags&TypeFlagsNever == 0 && r.c.isSimpleTypeRelatedTo(target, source, r.relation, nil) ||
		r.c.isSimpleTypeRelatedTo(source, target, r.relation, ifElse(reportErrors, r.reportError, nil)) {
		return TernaryTrue
	}
	if source.flags&TypeFlagsStructuredOrInstantiable != 0 || target.flags&TypeFlagsStructuredOrInstantiable != 0 {
		isPerformingExcessPropertyChecks := intersectionState&IntersectionStateTarget == 0 && isObjectLiteralType(source) && source.objectFlags&ObjectFlagsFreshLiteral != 0
		if isPerformingExcessPropertyChecks {
			if r.hasExcessProperties(source, target, reportErrors) {
				if reportErrors {
					r.reportRelationError(headMessage, source, ifElse(originalTarget.alias != nil, originalTarget, target))
				}
				return TernaryFalse
			}
		}
		isPerformingCommonPropertyChecks := (r.relation != r.c.comparableRelation || isUnitType(source)) &&
			intersectionState&IntersectionStateTarget == 0 &&
			source.flags&(TypeFlagsPrimitive|TypeFlagsObject|TypeFlagsIntersection) != 0 && source != r.c.globalObjectType &&
			target.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0 && r.c.isWeakType(target) && (len(r.c.getPropertiesOfType(source)) > 0 || r.c.typeHasCallOrConstructSignatures(source))
		isComparingJsxAttributes := source.objectFlags&ObjectFlagsJsxAttributes != 0
		if isPerformingCommonPropertyChecks && !r.c.hasCommonProperties(source, target, isComparingJsxAttributes) {
			if reportErrors {
				sourceString := r.c.typeToString(ifElse(originalSource.alias != nil, originalSource, source))
				targetString := r.c.typeToString(ifElse(originalTarget.alias != nil, originalTarget, target))
				calls := r.c.getSignaturesOfType(source, SignatureKindCall)
				constructs := r.c.getSignaturesOfType(source, SignatureKindConstruct)
				if len(calls) > 0 && r.isRelatedTo(r.c.getReturnTypeOfSignature(calls[0]), target, RecursionFlagsSource, false /*reportErrors*/) != TernaryFalse ||
					len(constructs) > 0 && r.isRelatedTo(r.c.getReturnTypeOfSignature(constructs[0]), target, RecursionFlagsSource, false /*reportErrors*/) != TernaryFalse {
					r.reportError(diagnostics.Value_of_type_0_has_no_properties_in_common_with_type_1_Did_you_mean_to_call_it, sourceString, targetString)
				} else {
					r.reportError(diagnostics.Type_0_has_no_properties_in_common_with_type_1, sourceString, targetString)
				}
			}
			return TernaryFalse
		}
		skipCaching := source.flags&TypeFlagsUnion != 0 && len(source.Types()) < 4 && target.flags&TypeFlagsUnion == 0 ||
			target.flags&TypeFlagsUnion != 0 && len(target.Types()) < 4 && source.flags&TypeFlagsStructuredOrInstantiable == 0
		var result Ternary
		if skipCaching {
			result = r.unionOrIntersectionRelatedTo(source, target, reportErrors, intersectionState)
		} else {
			result = r.recursiveTypeRelatedTo(source, target, reportErrors, intersectionState, recursionFlags)
		}
		if result != TernaryFalse {
			return result
		}
	}
	if reportErrors {
		r.reportErrorResults(originalSource, originalTarget, source, target, headMessage)
	}
	return TernaryFalse
}

func (r *Relater) hasExcessProperties(source *Type, target *Type, reportErrors bool) bool {
	if !isExcessPropertyCheckTarget(target) || !r.c.noImplicitAny && target.objectFlags&ObjectFlagsJSLiteral != 0 {
		// Disable excess property checks on JS literals to simulate having an implicit "index signature" - but only outside of noImplicitAny
		return false
	}
	isComparingJsxAttributes := source.objectFlags&ObjectFlagsJsxAttributes != 0
	if (r.relation == r.c.assignableRelation || r.relation == r.c.comparableRelation) && (r.c.isTypeSubsetOf(r.c.globalObjectType, target) || (!isComparingJsxAttributes && r.c.isEmptyObjectType(target))) {
		return false
	}
	reducedTarget := target
	var checkTypes []*Type
	if target.flags&TypeFlagsUnion != 0 {
		reducedTarget = r.c.findMatchingDiscriminantType(source, target, r.isRelatedToSimple)
		if reducedTarget == nil {
			reducedTarget = r.c.filterPrimitivesIfContainsNonPrimitive(target)
		}
		checkTypes = reducedTarget.Distributed()
	}
	for _, prop := range r.c.getPropertiesOfType(source) {
		if shouldCheckAsExcessProperty(prop, source.symbol) && !isIgnoredJsxProperty(source, prop) {
			if !r.c.isKnownProperty(reducedTarget, prop.Name, isComparingJsxAttributes) {
				if reportErrors {
					// Report error in terms of object types in the target as those are the only ones
					// we check in isKnownProperty.
					errorTarget := r.c.filterType(reducedTarget, isExcessPropertyCheckTarget)
					// We know *exactly* where things went wrong when comparing the types.
					// Use this property as the error node as this will be more helpful in
					// reasoning about what went wrong.
					if r.errorNode == nil {
						panic("No errorNode in hasExcessProperties")
					}
					if ast.IsJsxAttributes(r.errorNode) || isJsxOpeningLikeElement(r.errorNode) || isJsxOpeningLikeElement(r.errorNode.Parent) {
						// !!!
						// // JsxAttributes has an object-literal flag and undergo same type-assignablity check as normal object-literal.
						// // However, using an object-literal error message will be very confusing to the users so we give different a message.
						// if prop.valueDeclaration && isJsxAttribute(prop.valueDeclaration) && getSourceFileOfNode(errorNode) == getSourceFileOfNode(prop.valueDeclaration.name) {
						// 	// Note that extraneous children (as in `<NoChild>extra</NoChild>`) don't pass this check,
						// 	// since `children` is a Kind.PropertySignature instead of a Kind.JsxAttribute.
						// 	errorNode = prop.valueDeclaration.name
						// }
						// propName := c.symbolToString(prop)
						// suggestionSymbol := c.getSuggestedSymbolForNonexistentJSXAttribute(propName, errorTarget)
						// var suggestion *string
						// if suggestionSymbol {
						// 	suggestion = c.symbolToString(suggestionSymbol)
						// } else {
						// 	suggestion = nil
						// }
						// if suggestion {
						// 	reportError(Diagnostics.Property_0_does_not_exist_on_type_1_Did_you_mean_2, propName, c.typeToString(errorTarget), suggestion)
						// } else {
						// 	reportError(Diagnostics.Property_0_does_not_exist_on_type_1, propName, c.typeToString(errorTarget))
						// }
					} else {
						// use the property's value declaration if the property is assigned inside the literal itself
						var objectLiteralDeclaration *ast.Node
						if source.symbol != nil {
							objectLiteralDeclaration = core.FirstOrNil(source.symbol.Declarations)
						}
						var suggestion string
						if prop.ValueDeclaration != nil && isObjectLiteralElementLike(prop.ValueDeclaration) &&
							findAncestor(prop.ValueDeclaration, func(d *ast.Node) bool { return d == objectLiteralDeclaration }) != nil &&
							getSourceFileOfNode(objectLiteralDeclaration) == getSourceFileOfNode(r.errorNode) {
							name := prop.ValueDeclaration.Name()
							r.errorNode = name
							if ast.IsIdentifier(name) {
								suggestion = r.c.getSuggestionForNonexistentProperty(name.Text(), errorTarget)
							}
						}
						if suggestion != "" {
							r.reportError(diagnostics.Object_literal_may_only_specify_known_properties_but_0_does_not_exist_in_type_1_Did_you_mean_to_write_2, r.c.symbolToString(prop), r.c.typeToString(errorTarget), suggestion)
						} else {
							r.reportError(diagnostics.Object_literal_may_only_specify_known_properties_and_0_does_not_exist_in_type_1, r.c.symbolToString(prop), r.c.typeToString(errorTarget))
						}
					}
				}
				return true
			}
			if checkTypes != nil && r.isRelatedTo(r.c.getTypeOfSymbol(prop), r.c.getTypeOfPropertyInTypes(checkTypes, prop.Name), RecursionFlagsBoth, reportErrors) == TernaryFalse {
				if reportErrors {
					r.reportError(diagnostics.Types_of_property_0_are_incompatible, r.c.symbolToString(prop))
				}
				return true
			}
		}
	}
	return false
}

func (c *Checker) getTypeOfPropertyInTypes(types []*Type, name string) *Type {
	var propTypes []*Type
	for _, t := range types {
		propTypes = append(propTypes, c.getTypeOfPropertyInType(t, name))
	}
	return c.getUnionType(propTypes)
}

func (c *Checker) getTypeOfPropertyInType(t *Type, name string) *Type {
	t = c.getApparentType(t)
	var prop *ast.Symbol
	if t.flags&TypeFlagsUnionOrIntersection != 0 {
		prop = c.getPropertyOfUnionOrIntersectionType(t, name, false)
	} else {
		prop = c.getPropertyOfObjectType(t, name)
	}
	if prop != nil {
		return c.getTypeOfSymbol(prop)
	}
	indexInfo := c.getApplicableIndexInfoForName(t, name)
	if indexInfo != nil {
		return indexInfo.valueType
	}
	return c.undefinedType
}

func shouldCheckAsExcessProperty(prop *ast.Symbol, container *ast.Symbol) bool {
	return prop.ValueDeclaration != nil && container.ValueDeclaration != nil && prop.ValueDeclaration.Parent == container.ValueDeclaration
}

func isIgnoredJsxProperty(source *Type, sourceProp *ast.Symbol) bool {
	return source.objectFlags&ObjectFlagsJsxAttributes != 0 && isHyphenatedJsxName(sourceProp.Name)
}

func (c *Checker) isTypeSubsetOf(source *Type, target *Type) bool {
	return source == target || source.flags&TypeFlagsNever != 0 || target.flags&TypeFlagsUnion != 0 && c.isTypeSubsetOfUnion(source, target)
}

func (c *Checker) isTypeSubsetOfUnion(source *Type, target *Type) bool {
	if source.flags&TypeFlagsUnion != 0 {
		for _, t := range source.AsUnionType().types {
			if containsType(target.Types(), t) {
				return false
			}
		}
		return true
	}
	if source.flags&TypeFlagsEnumLike != 0 && c.getBaseTypeOfEnumLikeType(source) == target {
		return true
	}
	return containsType(target.Types(), source)
}

func (r *Relater) unionOrIntersectionRelatedTo(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	// Note that these checks are specifically ordered to produce correct results. In particular,
	// we need to deconstruct unions before intersections (because unions are always at the top),
	// and we need to handle "each" relations before "some" relations for the same kind of type.
	if source.flags&TypeFlagsUnion != 0 {
		if target.flags&TypeFlagsUnion != 0 {
			// Intersections of union types are normalized into unions of intersection types, and such normalized
			// unions can get very large and expensive to relate. The following fast path checks if the source union
			// originated in an intersection. If so, and if that intersection contains the target type, then we know
			// the result to be true (for any two types A and B, A & B is related to both A and B).
			sourceOrigin := source.AsUnionType().origin
			if sourceOrigin != nil && sourceOrigin.flags&TypeFlagsIntersection != 0 && target.alias != nil && slices.Contains(sourceOrigin.Types(), target) {
				return TernaryTrue
			}
			// Similarly, in unions of unions the we preserve the original list of unions. This original list is often
			// much shorter than the normalized result, so we scan it in the following fast path.
			targetOrigin := target.AsUnionType().origin
			if targetOrigin != nil && targetOrigin.flags&TypeFlagsUnion != 0 && source.alias != nil && slices.Contains(targetOrigin.Types(), source) {
				return TernaryTrue
			}
		}
		if r.relation == r.c.comparableRelation {
			return r.someTypeRelatedToType(source, target, reportErrors && source.flags&TypeFlagsPrimitive == 0, intersectionState)
		}
		return r.eachTypeRelatedToType(source, target, reportErrors && source.flags&TypeFlagsPrimitive == 0, intersectionState)
	}
	if target.flags&TypeFlagsUnion != 0 {
		return r.typeRelatedToSomeType(r.c.getRegularTypeOfObjectLiteral(source), target, reportErrors && source.flags&TypeFlagsPrimitive == 0 && target.flags&TypeFlagsPrimitive == 0, intersectionState)
	}
	if target.flags&TypeFlagsIntersection != 0 {
		return r.typeRelatedToEachType(source, target, reportErrors, IntersectionStateTarget)
	}
	// Source is an intersection. For the comparable relation, if the target is a primitive type we hoist the
	// constraints of all non-primitive types in the source into a new intersection. We do this because the
	// intersection may further constrain the constraints of the non-primitive types. For example, given a type
	// parameter 'T extends 1 | 2', the intersection 'T & 1' should be reduced to '1' such that it doesn't
	// appear to be comparable to '2'.
	if r.relation == r.c.comparableRelation && target.flags&TypeFlagsPrimitive != 0 {
		constraints := core.SameMap(source.Types(), func(t *Type) *Type {
			if t.flags&TypeFlagsInstantiable != 0 {
				constraint := r.c.getBaseConstraintOfType(t)
				if constraint != nil {
					return constraint
				}
				return r.c.unknownType
			}
			return t
		})
		if !core.Same(constraints, source.Types()) {
			source = r.c.getIntersectionType(constraints)
			if source.flags&TypeFlagsNever != 0 {
				return TernaryFalse
			}
			if source.flags&TypeFlagsIntersection == 0 {
				result := r.isRelatedTo(source, target, RecursionFlagsSource, false /*reportErrors*/)
				if result != TernaryFalse {
					return result
				}
				return r.isRelatedTo(target, source, RecursionFlagsSource, false /*reportErrors*/)
			}
		}
	}
	// Check to see if any constituents of the intersection are immediately related to the target.
	// Don't report errors though. Elaborating on whether a source constituent is related to the target is
	// not actually useful and leads to some confusing error messages. Instead, we rely on the caller
	// checking whether the full intersection viewed as an object is related to the target.
	return r.someTypeRelatedToType(source, target, false /*reportErrors*/, IntersectionStateSource)
}

func (r *Relater) someTypeRelatedToType(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	sourceTypes := source.Types()
	if source.flags&TypeFlagsUnion != 0 && containsType(sourceTypes, target) {
		return TernaryTrue
	}
	for i, t := range sourceTypes {
		related := r.isRelatedToEx(t, target, RecursionFlagsSource, reportErrors && i == len(sourceTypes)-1, nil /*headMessage*/, intersectionState)
		if related != TernaryFalse {
			return related
		}
	}
	return TernaryFalse
}

func (r *Relater) eachTypeRelatedToType(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	result := TernaryTrue
	sourceTypes := source.Types()
	// We strip `undefined` from the target if the `source` trivially doesn't contain it for our correspondence-checking fastpath
	// since `undefined` is frequently added by optionality and would otherwise spoil a potentially useful correspondence
	strippedTarget := r.getUndefinedStrippedTargetIfNeeded(source, target)
	var strippedTypes []*Type
	if strippedTarget.flags&TypeFlagsUnion != 0 {
		strippedTypes = strippedTarget.Types()
	}
	for i, sourceType := range sourceTypes {
		if strippedTarget.flags&TypeFlagsUnion != 0 && len(sourceTypes) >= len(strippedTypes) && len(sourceTypes)%len(strippedTypes) == 0 {
			// many unions are mappings of one another; in such cases, simply comparing members at the same index can shortcut the comparison
			// such unions will have identical lengths, and their corresponding elements will match up. Another common scenario is where a large
			// union has a union of objects intersected with it. In such cases, if the input was, eg `("a" | "b" | "c") & (string | boolean | {} | {whatever})`,
			// the result will have the structure `"a" | "b" | "c" | "a" & {} | "b" & {} | "c" & {} | "a" & {whatever} | "b" & {whatever} | "c" & {whatever}`
			// - the resulting union has a length which is a multiple of the original union, and the elements correspond modulo the length of the original union
			related := r.isRelatedToEx(sourceType, strippedTypes[i%len(strippedTypes)], RecursionFlagsBoth, false /*reportErrors*/, nil /*headMessage*/, intersectionState)
			if related != TernaryFalse {
				result &= related
				continue
			}
		}
		related := r.isRelatedToEx(sourceType, target, RecursionFlagsSource, reportErrors, nil /*headMessage*/, intersectionState)
		if related == TernaryFalse {
			return TernaryFalse
		}
		result &= related
	}
	return result
}

func (r *Relater) getUndefinedStrippedTargetIfNeeded(source *Type, target *Type) *Type {
	if source.flags&TypeFlagsUnion != 0 && target.flags&TypeFlagsUnion != 0 && source.Types()[0].flags&TypeFlagsUndefined == 0 && target.Types()[0].flags&TypeFlagsUndefined != 0 {
		return r.c.extractTypesOfKind(target, ^TypeFlagsUndefined)
	}
	return target
}

func (r *Relater) typeRelatedToSomeType(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	targetTypes := target.Types()
	if target.flags&TypeFlagsUnion != 0 {
		if containsType(targetTypes, source) {
			return TernaryTrue
		}
		if r.relation != r.c.comparableRelation && target.objectFlags&ObjectFlagsPrimitiveUnion != 0 && source.flags&TypeFlagsEnumLiteral == 0 &&
			(source.flags&(TypeFlagsStringLiteral|TypeFlagsBooleanLiteral|TypeFlagsBigIntLiteral) != 0 ||
				(r.relation == r.c.subtypeRelation || r.relation == r.c.strictSubtypeRelation) && source.flags&TypeFlagsNumberLiteral != 0) {
			// When relating a literal type to a union of primitive types, we know the relation is false unless
			// the union contains the base primitive type or the literal type in one of its fresh/regular forms.
			// We exclude numeric literals for non-subtype relations because numeric literals are assignable to
			// numeric enum literals with the same value. Similarly, we exclude enum literal types because
			// identically named enum types are related (see isEnumTypeRelatedTo). We exclude the comparable
			// relation in entirety because it needs to be checked in both directions.
			var alternateForm *Type
			if source == source.AsLiteralType().regularType {
				alternateForm = source.AsLiteralType().freshType
			} else {
				alternateForm = source.AsLiteralType().regularType
			}
			var primitive *Type
			switch {
			case source.flags&TypeFlagsStringLiteral != 0:
				primitive = r.c.stringType
			case source.flags&TypeFlagsNumberLiteral != 0:
				primitive = r.c.numberType
			case source.flags&TypeFlagsBigIntLiteral != 0:
				primitive = r.c.bigintType
			}
			if primitive != nil && containsType(targetTypes, primitive) || alternateForm != nil && containsType(targetTypes, alternateForm) {
				return TernaryTrue
			}
			return TernaryFalse
		}
		match := r.c.getMatchingUnionConstituentForType(target, source)
		if match != nil {
			related := r.isRelatedToEx(source, match, RecursionFlagsTarget, false /*reportErrors*/, nil /*headMessage*/, intersectionState)
			if related != TernaryFalse {
				return related
			}
		}
	}
	for _, t := range targetTypes {
		related := r.isRelatedToEx(source, t, RecursionFlagsTarget, false /*reportErrors*/, nil /*headMessage*/, intersectionState)
		if related != TernaryFalse {
			return related
		}
	}
	if reportErrors {
		// Elaborate only if we can find a best matching type in the target union
		bestMatchingType := r.c.getBestMatchingType(source, target, r.isRelatedToSimple)
		if bestMatchingType != nil {
			r.isRelatedToEx(source, bestMatchingType, RecursionFlagsTarget, true /*reportErrors*/, nil /*headMessage*/, intersectionState)
		}
	}
	return TernaryFalse
}

func (r *Relater) typeRelatedToEachType(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	result := TernaryTrue
	targetTypes := target.Types()
	for _, targetType := range targetTypes {
		related := r.isRelatedToEx(source, targetType, RecursionFlagsTarget, reportErrors /*headMessage*/, nil, intersectionState)
		if related == TernaryFalse {
			return TernaryFalse
		}
		result &= related
	}
	return result
}

func (r *Relater) eachTypeRelatedToSomeType(source *Type, target *Type) Ternary {
	result := TernaryTrue
	sourceTypes := source.Types()
	for _, sourceType := range sourceTypes {
		related := r.typeRelatedToSomeType(sourceType, target, false /*reportErrors*/, IntersectionStateNone)
		if related == TernaryFalse {
			return TernaryFalse
		}
		result &= related
	}
	return result
}

// Determine if possibly recursive types are related. First, check if the result is already available in the global cache.
// Second, check if we have already started a comparison of the given two types in which case we assume the result to be true.
// Third, check if both types are part of deeply nested chains of generic type instantiations and if so assume the types are
// equal and infinitely expanding. Fourth, if we have reached a depth of 100 nested comparisons, assume we have runaway recursion
// and issue an error. Otherwise, actually compare the structure of the two types.
func (r *Relater) recursiveTypeRelatedTo(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState, recursionFlags RecursionFlags) Ternary {
	if r.overflow {
		return TernaryFalse
	}
	id := getRelationKey(source, target, intersectionState, r.relation == r.c.identityRelation, false /*ignoreConstraints*/)
	entry := r.relation.get(id)
	if entry != RelationComparisonResultNone {
		if reportErrors && entry&RelationComparisonResultFailed != 0 && entry&RelationComparisonResultOverflow == 0 {
			// We are elaborating errors and the cached result is a failure not due to a comparison overflow,
			// so we will do the comparison again to generate an error message.
		} else {
			if r.c.outofbandVarianceMarkerHandler != nil {
				// We're in the middle of variance checking - integrate any unmeasurable/unreliable flags from this cached component
				saved := entry & RelationComparisonResultReportsMask
				if saved&RelationComparisonResultReportsUnmeasurable != 0 {
					r.c.instantiateType(source, r.c.reportUnmeasurableMapper)
				}
				if saved&RelationComparisonResultReportsUnreliable != 0 {
					r.c.instantiateType(source, r.c.reportUnreliableMapper)
				}
			}
			if reportErrors && entry&RelationComparisonResultOverflow != 0 {
				message := ifElse(entry&RelationComparisonResultComplexityOverflow != 0,
					diagnostics.Excessive_complexity_comparing_types_0_and_1,
					diagnostics.Excessive_stack_depth_comparing_types_0_and_1)
				r.reportError(message, r.c.typeToString(source), r.c.typeToString(target))
			}
			if entry&RelationComparisonResultSucceeded != 0 {
				return TernaryTrue
			}
			return TernaryFalse
		}
	}
	if r.relationCount <= 0 {
		r.overflow = true
		return TernaryFalse
	}
	// If source and target are already being compared, consider them related with assumptions
	if r.maybeKeysSet.Has(id) {
		return TernaryMaybe
	}
	// A key that ends with "*" is an indication that we have type references that reference constrained
	// type parameters. For such keys we also check against the key we would have gotten if all type parameters
	// were unconstrained.
	if strings.HasSuffix(id, "*") {
		broadestEquivalentId := getRelationKey(source, target, intersectionState, r.relation == r.c.identityRelation, true /*ignoreConstraints*/)
		if r.maybeKeysSet.Has(broadestEquivalentId) {
			return TernaryMaybe
		}
	}
	if len(r.sourceStack) == 100 || len(r.targetStack) == 100 {
		r.overflow = true
		return TernaryFalse
	}
	maybeStart := len(r.maybeKeys)
	r.maybeKeys = append(r.maybeKeys, id)
	r.maybeKeysSet.Add(id)
	saveExpandingFlags := r.expandingFlags
	if recursionFlags&RecursionFlagsSource != 0 {
		r.sourceStack = append(r.sourceStack, source)
		if r.expandingFlags&ExpandingFlagsSource == 0 && r.c.isDeeplyNestedType(source, r.sourceStack, 3) {
			r.expandingFlags |= ExpandingFlagsSource
		}
	}
	if recursionFlags&RecursionFlagsTarget != 0 {
		r.targetStack = append(r.targetStack, target)
		if r.expandingFlags&ExpandingFlagsTarget == 0 && r.c.isDeeplyNestedType(target, r.targetStack, 3) {
			r.expandingFlags |= ExpandingFlagsTarget
		}
	}
	propagatingVarianceFlags := RelationComparisonResultNone
	var originalHandler func(bool)
	if r.c.outofbandVarianceMarkerHandler != nil {
		originalHandler = r.c.outofbandVarianceMarkerHandler
		r.c.outofbandVarianceMarkerHandler = func(onlyUnreliable bool) {
			if onlyUnreliable {
				propagatingVarianceFlags |= RelationComparisonResultReportsUnreliable
			} else {
				propagatingVarianceFlags |= RelationComparisonResultReportsUnmeasurable
			}
			originalHandler(onlyUnreliable)
		}
	}
	var result Ternary
	if r.expandingFlags == ExpandingFlagsBoth {
		result = TernaryMaybe
	} else {
		result = r.structuredTypeRelatedTo(source, target, reportErrors, intersectionState)
	}
	if r.c.outofbandVarianceMarkerHandler != nil {
		r.c.outofbandVarianceMarkerHandler = originalHandler
	}
	if recursionFlags&RecursionFlagsSource != 0 {
		r.sourceStack = r.sourceStack[:len(r.sourceStack)-1]
	}
	if recursionFlags&RecursionFlagsTarget != 0 {
		r.targetStack = r.targetStack[:len(r.targetStack)-1]
	}
	r.expandingFlags = saveExpandingFlags
	if result != TernaryFalse {
		if result == TernaryTrue || (len(r.sourceStack) == 0 && len(r.targetStack) == 0) {
			if result == TernaryTrue || result == TernaryMaybe {
				// If result is definitely true, record all maybe keys as having succeeded. Also, record Ternary.Maybe
				// results as having succeeded once we reach depth 0, but never record Ternary.Unknown results.
				r.resetMaybeStack(maybeStart, propagatingVarianceFlags, true)
			} else {
				r.resetMaybeStack(maybeStart, propagatingVarianceFlags, false)
			}
		}
		// Note: it's intentional that we don't reset in the else case;
		// we leave them on the stack such that when we hit depth zero
		// above, we can report all of them as successful.
	} else {
		// A false result goes straight into global cache (when something is false under
		// assumptions it will also be false without assumptions)
		r.relation.set(id, RelationComparisonResultFailed|propagatingVarianceFlags)
		r.relationCount--
		r.resetMaybeStack(maybeStart, propagatingVarianceFlags, false)
	}
	return result
}

func (r *Relater) resetMaybeStack(maybeStart int, propagatingVarianceFlags RelationComparisonResult, markAllAsSucceeded bool) {
	for i := maybeStart; i < len(r.maybeKeys); i++ {
		r.maybeKeysSet.Delete(r.maybeKeys[i])
		if markAllAsSucceeded {
			r.relation.set(r.maybeKeys[i], RelationComparisonResultSucceeded|propagatingVarianceFlags)
			r.relationCount--
		}
	}
	r.maybeKeys = r.maybeKeys[:maybeStart]
}

func (r *Relater) getErrorState() errorState {
	return errorState{
		errorChain:  r.errorChain,
		relatedInfo: r.relatedInfo,
	}
}

func (r *Relater) restoreErrorState(e errorState) {
	r.errorChain = e.errorChain
	r.relatedInfo = e.relatedInfo
}

func (r *Relater) structuredTypeRelatedTo(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	saveErrorState := r.getErrorState()
	result := r.structuredTypeRelatedToWorker(source, target, reportErrors, intersectionState)
	if r.relation != r.c.identityRelation {
		// The combined constraint of an intersection type is the intersection of the constraints of
		// the constituents. When an intersection type contains instantiable types with union type
		// constraints, there are situations where we need to examine the combined constraint. One is
		// when the target is a union type. Another is when the intersection contains types belonging
		// to one of the disjoint domains. For example, given type variables T and U, each with the
		// constraint 'string | number', the combined constraint of 'T & U' is 'string | number' and
		// we need to check this constraint against a union on the target side. Also, given a type
		// variable V constrained to 'string | number', 'V & number' has a combined constraint of
		// 'string & number | number & number' which reduces to just 'number'.
		// This also handles type parameters, as a type parameter with a union constraint compared against a union
		// needs to have its constraint hoisted into an intersection with said type parameter, this way
		// the type param can be compared with itself in the target (with the influence of its constraint to match other parts)
		// For example, if `T extends 1 | 2` and `U extends 2 | 3` and we compare `T & U` to `T & U & (1 | 2 | 3)`
		if result == TernaryFalse && (source.flags&TypeFlagsIntersection != 0 || source.flags&TypeFlagsTypeParameter != 0 && target.flags&TypeFlagsUnion != 0) {
			var sourceTypes []*Type
			if source.flags&TypeFlagsIntersection != 0 {
				sourceTypes = source.Types()
			} else {
				sourceTypes = []*Type{source}
			}
			constraint := r.c.getEffectiveConstraintOfIntersection(sourceTypes, target.flags&TypeFlagsUnion != 0)
			if constraint != nil && everyType(constraint, func(c *Type) bool { return c != source }) {
				// TODO: Stack errors so we get a pyramid for the "normal" comparison above, _and_ a second for this
				result = r.isRelatedToEx(constraint, target, RecursionFlagsSource, false /*reportErrors*/, nil /*headMessage*/, intersectionState)
			}
		}
		switch {
		// When the target is an intersection we need an extra property check in order to detect nested excess
		// properties and nested weak types. The following are motivating examples that all should be errors, but
		// aren't without this extra property check:
		//
		//   let obj: { a: { x: string } } & { c: number } = { a: { x: 'hello', y: 2 }, c: 5 };  // Nested excess property
		//
		//   declare let wrong: { a: { y: string } };
		//   let weak: { a?: { x?: number } } & { c?: string } = wrong;  // Nested weak object type
		//
		case result != TernaryFalse && intersectionState&IntersectionStateTarget == 0 && target.flags&TypeFlagsIntersection != 0 && !r.c.isGenericObjectType(target) && source.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0:
			result &= r.propertiesRelatedTo(source, target, reportErrors, core.Set[string]{} /*excludedProperties*/, false /*optionalsOnly*/, IntersectionStateNone)
			if result != 0 && isObjectLiteralType(source) && source.objectFlags&ObjectFlagsFreshLiteral != 0 {
				result &= r.indexSignaturesRelatedTo(source, target, false /*sourceIsPrimitive*/, reportErrors, IntersectionStateNone)
			}
		// When the source is an intersection we need an extra check of any optional properties in the target to
		// detect possible mismatched property types. For example:
		//
		//   function foo<T extends object>(x: { a?: string }, y: T & { a: boolean }) {
		//     x = y;  // Mismatched property in source intersection
		//   }
		//
		case result != 0 && r.c.isNonGenericObjectType(target) && !r.c.isArrayOrTupleType(target) && r.isSourceIntersectionNeedingExtraCheck(source, target):
			result &= r.propertiesRelatedTo(source, target, reportErrors, core.Set[string]{} /*excludedProperties*/, true /*optionalsOnly*/, intersectionState)
		}
	}
	if result != TernaryFalse {
		r.restoreErrorState(saveErrorState)
	}
	return result
}

func (r *Relater) isSourceIntersectionNeedingExtraCheck(source *Type, target *Type) bool {
	return source.flags&TypeFlagsIntersection != 0 && r.c.getApparentType(source).flags&TypeFlagsStructuredType != 0 &&
		!core.Some(source.Types(), func(t *Type) bool {
			return t == target || t.objectFlags&ObjectFlagsNonInferrableType != 0
		})
}

func (r *Relater) structuredTypeRelatedToWorker(source *Type, target *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	var result Ternary
	var varianceCheckFailed bool
	var originalErrorChain *ErrorChain
	saveErrorState := r.getErrorState()
	relateVariances := func(sourceTypeArguments []*Type, targetTypeArguments []*Type, variances []VarianceFlags, intersectionState IntersectionState) (Ternary, bool) {
		if result := r.typeArgumentsRelatedTo(sourceTypeArguments, targetTypeArguments, variances, reportErrors, intersectionState); result != TernaryFalse {
			return result, true
		}
		if core.Some(variances, func(v VarianceFlags) bool { return v&VarianceFlagsAllowsStructuralFallback != 0 }) {
			// If some type parameter was `Unmeasurable` or `Unreliable`, and we couldn't pass by assuming it was identical, then we
			// have to allow a structural fallback check
			// We elide the variance-based error elaborations, since those might not be too helpful, since we'll potentially
			// be assuming identity of the type parameter.
			originalErrorChain = nil
			r.restoreErrorState(saveErrorState)
			return TernaryFalse, false
		}
		allowStructuralFallback := r.c.hasCovariantVoidArgument(targetTypeArguments, variances)
		varianceCheckFailed = !allowStructuralFallback
		// The type arguments did not relate appropriately, but it may be because we have no variance
		// information (in which case typeArgumentsRelatedTo defaulted to covariance for all type
		// arguments). It might also be the case that the target type has a 'void' type argument for
		// a covariant type parameter that is only used in return positions within the generic type
		// (in which case any type argument is permitted on the source side). In those cases we proceed
		// with a structural comparison. Otherwise, we know for certain the instantiations aren't
		// related and we can return here.
		if len(variances) != 0 && !allowStructuralFallback {
			// In some cases generic types that are covariant in regular type checking mode become
			// invariant in --strictFunctionTypes mode because one or more type parameters are used in
			// both co- and contravariant positions. In order to make it easier to diagnose *why* such
			// types are invariant, if any of the type parameters are invariant we reset the reported
			// errors and instead force a structural comparison (which will include elaborations that
			// reveal the reason).
			// We can switch on `reportErrors` here, since varianceCheckFailed guarantees we return `False`,
			// we can return `False` early here to skip calculating the structural error message we don't need.
			if varianceCheckFailed && !(reportErrors && core.Some(variances, func(v VarianceFlags) bool { return (v & VarianceFlagsVarianceMask) == VarianceFlagsInvariant })) {
				return TernaryFalse, true
			}
			// We remember the original error information so we can restore it in case the structural
			// comparison unexpectedly succeeds. This can happen when the structural comparison result
			// is a Ternary.Maybe for example caused by the recursion depth limiter.
			originalErrorChain = r.errorChain
			r.restoreErrorState(saveErrorState)
		}
		return TernaryFalse, false
	}
	switch {
	case r.relation == r.c.identityRelation:
		// We've already checked that source.flags and target.flags are identical
		switch {
		case source.flags&TypeFlagsUnionOrIntersection != 0:
			result := r.eachTypeRelatedToSomeType(source, target)
			if result != TernaryFalse {
				result &= r.eachTypeRelatedToSomeType(target, source)
			}
			return result
		case source.flags&TypeFlagsIndex != 0:
			return r.isRelatedTo(source.Target(), target.Target(), RecursionFlagsBoth, false /*reportErrors*/)
		case source.flags&TypeFlagsIndexedAccess != 0:
			result = r.isRelatedTo(source.AsIndexedAccessType().objectType, target.AsIndexedAccessType().objectType, RecursionFlagsBoth, false /*reportErrors*/)
			if result != TernaryFalse {
				result &= r.isRelatedTo(source.AsIndexedAccessType().indexType, target.AsIndexedAccessType().indexType, RecursionFlagsBoth, false /*reportErrors*/)
				if result != TernaryFalse {
					return result
				}
			}
		case source.flags&TypeFlagsConditional != 0:
			if source.AsConditionalType().root.isDistributive == target.AsConditionalType().root.isDistributive {
				result = r.isRelatedTo(source.AsConditionalType().checkType, target.AsConditionalType().checkType, RecursionFlagsBoth, false /*reportErrors*/)
				if result != TernaryFalse {
					result &= r.isRelatedTo(source.AsConditionalType().extendsType, target.AsConditionalType().extendsType, RecursionFlagsBoth, false /*reportErrors*/)
					if result != TernaryFalse {
						result &= r.isRelatedTo(r.c.getTrueTypeFromConditionalType(source), r.c.getTrueTypeFromConditionalType(target), RecursionFlagsBoth, false /*reportErrors*/)
						if result != TernaryFalse {
							result &= r.isRelatedTo(r.c.getFalseTypeFromConditionalType(source), r.c.getFalseTypeFromConditionalType(target), RecursionFlagsBoth, false /*reportErrors*/)
							if result != TernaryFalse {
								return result
							}
						}
					}
				}
			}
		case source.flags&TypeFlagsSubstitution != 0:
			result = r.isRelatedTo(source.AsSubstitutionType().baseType, target.AsSubstitutionType().baseType, RecursionFlagsBoth, false /*reportErrors*/)
			if result != TernaryFalse {
				result &= r.isRelatedTo(source.AsSubstitutionType().constraint, target.AsSubstitutionType().constraint, RecursionFlagsBoth, false /*reportErrors*/)
				if result != TernaryFalse {
					return result
				}
			}
		}
		if source.flags&TypeFlagsObject == 0 {
			return TernaryFalse
		}
	case source.flags&TypeFlagsUnionOrIntersection != 0 || target.flags&TypeFlagsUnionOrIntersection != 0:
		result = r.unionOrIntersectionRelatedTo(source, target, reportErrors, intersectionState)
		if result != TernaryFalse {
			return result
		}
		// The ordered decomposition above doesn't handle all cases. Specifically, we also need to handle:
		// Source is instantiable (e.g. source has union or intersection constraint).
		// Source is an object, target is a union (e.g. { a, b: boolean } <=> { a, b: true } | { a, b: false }).
		// Source is an intersection, target is an object (e.g. { a } & { b } <=> { a, b }).
		// Source is an intersection, target is a union (e.g. { a } & { b: boolean } <=> { a, b: true } | { a, b: false }).
		// Source is an intersection, target instantiable (e.g. string & { tag } <=> T["a"] constrained to string & { tag }).
		if !(source.flags&TypeFlagsInstantiable != 0 ||
			source.flags&TypeFlagsObject != 0 && target.flags&TypeFlagsUnion != 0 ||
			source.flags&TypeFlagsIntersection != 0 && target.flags&(TypeFlagsObject|TypeFlagsUnion|TypeFlagsInstantiable) != 0) {
			return TernaryFalse
		}
	}
	// We limit alias variance probing to only object and conditional types since their alias behavior
	// is more predictable than other, interned types, which may or may not have an alias depending on
	// the order in which things were checked.
	if source.flags&(TypeFlagsObject|TypeFlagsConditional) != 0 && source.alias != nil && len(source.alias.typeArguments) != 0 &&
		target.alias != nil && source.alias.symbol == target.alias.symbol && !(r.c.isMarkerType(source) || r.c.isMarkerType(target)) {
		variances := r.c.getAliasVariances(source.alias.symbol)
		if len(variances) == 0 {
			return TernaryUnknown
		}
		params := r.c.typeAliasLinks.get(source.alias.symbol).typeParameters
		minParams := r.c.getMinTypeArgumentCount(params)
		sourceTypes := r.c.fillMissingTypeArguments(source.alias.typeArguments, params, minParams)
		targetTypes := r.c.fillMissingTypeArguments(target.alias.typeArguments, params, minParams)
		varianceResult, ok := relateVariances(sourceTypes, targetTypes, variances, intersectionState)
		if ok {
			return varianceResult
		}
	}
	// For a generic type T and a type U that is assignable to T, [...U] is assignable to T, U is assignable to readonly [...T],
	// and U is assignable to [...T] when U is constrained to a mutable array or tuple type.
	if isSingleElementGenericTupleType(source) && !source.TargetTupleType().readonly {
		result = r.isRelatedTo(r.c.getTypeArguments(source)[0], target, RecursionFlagsSource, false /*reportErrors*/)
		if result != TernaryFalse {
			return result
		}
	}
	if isSingleElementGenericTupleType(target) && (target.TargetTupleType().readonly || r.c.isMutableArrayOrTuple(r.c.getBaseConstraintOrType(source))) {
		result = r.isRelatedTo(source, r.c.getTypeArguments(target)[0], RecursionFlagsTarget, false /*reportErrors*/)
		if result != TernaryFalse {
			return result
		}
	}
	switch {
	case target.flags&TypeFlagsTypeParameter != 0:
		// A source type { [P in Q]: X } is related to a target type T if keyof T is related to Q and X is related to T[Q].
		if source.objectFlags&ObjectFlagsMapped != 0 && source.AsMappedType().declaration.NameType == nil && r.isRelatedTo(r.c.getIndexType(target), r.c.getConstraintTypeFromMappedType(source), RecursionFlagsBoth, false) != TernaryFalse {
			if getMappedTypeModifiers(source)&MappedTypeModifiersIncludeOptional == 0 {
				templateType := r.c.getTemplateTypeFromMappedType(source)
				indexedAccessType := r.c.getIndexedAccessType(target, r.c.getTypeParameterFromMappedType(source))
				result = r.isRelatedTo(templateType, indexedAccessType, RecursionFlagsBoth, reportErrors)
				if result != TernaryFalse {
					return result
				}
			}
		}
		if r.relation == r.c.comparableRelation && source.flags&TypeFlagsTypeParameter != 0 {
			// This is a carve-out in comparability to essentially forbid comparing a type parameter with another type parameter
			// unless one extends the other. (Remember: comparability is mostly bidirectional!)
			constraint := r.c.getConstraintOfTypeParameter(source)
			if constraint != nil {
				for constraint != nil && someType(constraint, func(c *Type) bool { return c.flags&TypeFlagsTypeParameter != 0 }) {
					result = r.isRelatedTo(constraint, target, RecursionFlagsSource, false /*reportErrors*/)
					if result != TernaryFalse {
						return result
					}
					constraint = r.c.getConstraintOfTypeParameter(constraint)
				}
			}
			return TernaryFalse
		}
	case target.flags&TypeFlagsIndexedAccess != 0:
		if source.flags&TypeFlagsIndexedAccess != 0 {
			// Relate components directly before falling back to constraint relationships
			// A type S[K] is related to a type T[J] if S is related to T and K is related to J.
			result = r.isRelatedTo(source.AsIndexedAccessType().objectType, target.AsIndexedAccessType().objectType, RecursionFlagsBoth, reportErrors)
			if result != TernaryFalse {
				result &= r.isRelatedTo(source.AsIndexedAccessType().indexType, target.AsIndexedAccessType().indexType, RecursionFlagsBoth, reportErrors)
			}
			if result != TernaryFalse {
				return result
			}
			if reportErrors {
				originalErrorChain = r.errorChain
			}
		}
		// A type S is related to a type T[K] if S is related to C, where C is the base
		// constraint of T[K] for writing.
		if r.relation == r.c.assignableRelation || r.relation == r.c.comparableRelation {
			objectType := target.AsIndexedAccessType().objectType
			indexType := target.AsIndexedAccessType().indexType
			baseObjectType := r.c.getBaseConstraintOrType(objectType)
			baseIndexType := r.c.getBaseConstraintOrType(indexType)
			if !r.c.isGenericObjectType(baseObjectType) && !r.c.isGenericIndexType(baseIndexType) {
				accessFlags := AccessFlagsWriting | (ifElse(baseObjectType != objectType, AccessFlagsNoIndexSignatures, 0))
				constraint := r.c.getIndexedAccessTypeOrUndefined(baseObjectType, baseIndexType, accessFlags, nil, nil)
				if constraint != nil {
					if reportErrors && originalErrorChain != nil {
						// create a new chain for the constraint error
						r.restoreErrorState(saveErrorState)
					}
					result = r.isRelatedToEx(source, constraint, RecursionFlagsTarget, reportErrors, nil /*headMessage*/, intersectionState)
					if result != TernaryFalse {
						return result
					}
					// prefer the shorter chain of the constraint comparison chain, and the direct comparison chain
					if reportErrors && originalErrorChain != nil && r.errorChain != nil {
						if chainDepth(originalErrorChain) <= chainDepth(r.errorChain) {
							r.errorChain = originalErrorChain
						}
					}
				}
			}
		}
		if reportErrors {
			originalErrorChain = nil
		}
	case target.flags&TypeFlagsIndex != 0:
		targetType := target.AsIndexType().target
		// A keyof S is related to a keyof T if T is related to S.
		if source.flags&TypeFlagsIndex != 0 {
			result = r.isRelatedTo(targetType, source.AsIndexType().target, RecursionFlagsBoth, false /*reportErrors*/)
			if result != TernaryFalse {
				return result
			}
		}
		if isTupleType(targetType) {
			// An index type can have a tuple type target when the tuple type contains variadic elements.
			// Check if the source is related to the known keys of the tuple type.
			result = r.isRelatedTo(source, r.c.getKnownKeysOfTupleType(targetType), RecursionFlagsTarget, reportErrors)
			if result != TernaryFalse {
				return result
			}
		} else {
			// A type S is assignable to keyof T if S is assignable to keyof C, where C is the
			// simplified form of T or, if T doesn't simplify, the constraint of T.
			constraint := r.c.getSimplifiedTypeOrConstraint(targetType)
			if constraint != nil {
				// We require Ternary.True here such that circular constraints don't cause
				// false positives. For example, given 'T extends { [K in keyof T]: string }',
				// 'keyof T' has itself as its constraint and produces a Ternary.Maybe when
				// related to other types.
				if r.isRelatedTo(source, r.c.getIndexTypeEx(constraint, target.AsIndexType().indexFlags|IndexFlagsNoReducibleCheck), RecursionFlagsTarget, reportErrors) == TernaryTrue {
					return TernaryTrue
				}
			} else if r.c.isGenericMappedType(targetType) {
				// generic mapped types that don't simplify or have a constraint still have a very simple set of keys we can compare against
				// - their nameType or constraintType.
				// In many ways, this comparison is a deferred version of what `getIndexTypeForMappedType` does to actually resolve the keys for _non_-generic types
				nameType := r.c.getNameTypeFromMappedType(targetType)
				constraintType := r.c.getConstraintTypeFromMappedType(targetType)
				var targetKeys *Type
				if nameType != nil && r.c.isMappedTypeWithKeyofConstraintDeclaration(targetType) {
					// we need to get the apparent mappings and union them with the generic mappings, since some properties may be
					// missing from the `constraintType` which will otherwise be mapped in the object
					mappedKeys := r.c.getApparentMappedTypeKeys(nameType, targetType)
					// We still need to include the non-apparent (and thus still generic) keys in the target side of the comparison (in case they're in the source side)
					targetKeys = r.c.getUnionType([]*Type{mappedKeys, nameType})
				} else if nameType != nil {
					targetKeys = nameType
				} else {
					targetKeys = constraintType
				}
				if r.isRelatedTo(source, targetKeys, RecursionFlagsTarget, reportErrors) == TernaryTrue {
					return TernaryTrue
				}
			}
		}
	case target.flags&TypeFlagsConditional != 0:
		// !!!
	case target.flags&TypeFlagsTemplateLiteral != 0:
		// !!!
	case target.flags&TypeFlagsStringMapping != 0:
		// !!!
	case r.c.isGenericMappedType(target) && r.relation != r.c.identityRelation:
		// !!!
	}
	switch {
	case source.flags&TypeFlagsTypeVariable != 0:
		// IndexedAccess comparisons are handled above in the `target.flags&TypeFlagsIndexedAccess` branch
		if source.flags&TypeFlagsIndexedAccess == 0 || target.flags&TypeFlagsIndexedAccess == 0 {
			constraint := r.c.getConstraintOfType(source)
			if constraint == nil {
				constraint = r.c.unknownType
			}
			// hi-speed no-this-instantiation check (less accurate, but avoids costly `this`-instantiation when the constraint will suffice), see #28231 for report on why this is needed
			result = r.isRelatedToEx(constraint, target, RecursionFlagsSource, false /*reportErrors*/, nil /*headMessage*/, intersectionState)
			if result != TernaryFalse {
				return result
			}
			constraintWithThis := r.c.getTypeWithThisArgument(constraint, source, false /*needApparentType*/)
			result = r.isRelatedToEx(constraintWithThis, target, RecursionFlagsSource, reportErrors && constraint != r.c.unknownType && target.flags&source.flags&TypeFlagsTypeParameter == 0, nil /*headMessage*/, intersectionState)
			if result != TernaryFalse {
				return result
			}
			if r.c.isMappedTypeGenericIndexedAccess(source) {
				// For an indexed access type { [P in K]: E}[X], above we have already explored an instantiation of E with X
				// substituted for P. We also want to explore type { [P in K]: E }[C], where C is the constraint of X.
				indexConstraint := r.c.getConstraintOfType(source.AsIndexedAccessType().indexType)
				if indexConstraint != nil {
					result = r.isRelatedTo(r.c.getIndexedAccessType(source.AsIndexedAccessType().objectType, indexConstraint), target, RecursionFlagsSource, reportErrors)
					if result != TernaryFalse {
						return result
					}
				}
			}
		}
	case source.flags&TypeFlagsIndex != 0:
		isDeferredMappedIndex := r.c.shouldDeferIndexType(source.AsIndexType().target, source.AsIndexType().indexFlags) && source.AsIndexType().target.objectFlags&ObjectFlagsMapped != 0
		result = r.isRelatedTo(r.c.stringNumberSymbolType, target, RecursionFlagsSource, reportErrors && !isDeferredMappedIndex)
		if result != TernaryFalse {
			return result
		}
		if isDeferredMappedIndex {
			mappedType := source.AsIndexType().target
			nameType := r.c.getNameTypeFromMappedType(mappedType)
			// Unlike on the target side, on the source side we do *not* include the generic part of the `nameType`, since that comes from a
			// (potentially anonymous) mapped type local type parameter, so that'd never assign outside the mapped type body, but we still want to
			// allow assignments of index types of identical (or similar enough) mapped types.
			// eg, `keyof {[X in keyof A]: Obj[X]}` should be assignable to `keyof {[Y in keyof A]: Tup[Y]}` because both map over the same set of keys (`keyof A`).
			// Without this source-side breakdown, a `keyof {[X in keyof A]: Obj[X]}` style type won't be assignable to anything except itself, which is much too strict.
			var sourceMappedKeys *Type
			if nameType != nil && r.c.isMappedTypeWithKeyofConstraintDeclaration(mappedType) {
				sourceMappedKeys = r.c.getApparentMappedTypeKeys(nameType, mappedType)
			} else if nameType != nil {
				sourceMappedKeys = nameType
			} else {
				sourceMappedKeys = r.c.getConstraintTypeFromMappedType(mappedType)
			}
			result = r.isRelatedTo(sourceMappedKeys, target, RecursionFlagsSource, reportErrors)
			if result != TernaryFalse {
				return result
			}
		}
	case source.flags&TypeFlagsConditional != 0:
		// !!!
	case source.flags&TypeFlagsTemplateLiteral != 0 && target.flags&TypeFlagsObject == 0:
		if target.flags&TypeFlagsTemplateLiteral == 0 {
			constraint := r.c.getBaseConstraintOfType(source)
			if constraint != nil && constraint != source {
				result = r.isRelatedTo(constraint, target, RecursionFlagsSource, reportErrors)
				if result != TernaryFalse {
					return result
				}
			}
		}
	case source.flags&TypeFlagsStringMapping != 0:
		// !!!
	default:
		// An empty object type is related to any mapped type that includes a '?' modifier.
		if r.relation != r.c.subtypeRelation && r.relation != r.c.strictSubtypeRelation && isPartialMappedType(target) && r.c.isEmptyObjectType(source) {
			return TernaryTrue
		}
		if r.c.isGenericMappedType(target) {
			if r.c.isGenericMappedType(source) {
				result = r.mappedTypeRelatedTo(source, target, reportErrors)
				if result != TernaryFalse {
					return result
				}
			}
			return TernaryFalse
		}
		sourceIsPrimitive := source.flags&TypeFlagsPrimitive != 0
		if r.relation != r.c.identityRelation {
			source = r.c.getApparentType(source)
		} else if r.c.isGenericMappedType(source) {
			return TernaryFalse
		}
		switch {
		case source.objectFlags&ObjectFlagsReference != 0 && target.objectFlags&ObjectFlagsReference != 0 && source.Target() == target.Target() && !isTupleType(source) && !r.c.isMarkerType(source) && !r.c.isMarkerType(target):
			// When strictNullChecks is disabled, the element type of the empty array literal is undefinedWideningType,
			// and an empty array literal wouldn't be assignable to a `never[]` without this check.
			if r.c.isEmptyArrayLiteralType(source) {
				return TernaryTrue
			}
			// We have type references to the same generic type, and the type references are not marker
			// type references (which are intended by be compared structurally). Obtain the variance
			// information for the type parameters and relate the type arguments accordingly.
			variances := r.c.getVariances(source.Target())
			// We return Ternary.Maybe for a recursive invocation of getVariances (signalled by emptyArray). This
			// effectively means we measure variance only from type parameter occurrences that aren't nested in
			// recursive instantiations of the generic type.
			if len(variances) == 0 {
				return TernaryUnknown
			}
			varianceResult, ok := relateVariances(r.c.getTypeArguments(source), r.c.getTypeArguments(target), variances, intersectionState)
			if ok {
				return varianceResult
			}
		case r.c.isArrayType(target) && (r.c.isReadonlyArrayType(target) && everyType(source, r.c.isArrayOrTupleType) || everyType(source, isMutableTupleType)):
			if r.relation != r.c.identityRelation {
				return r.isRelatedTo(r.c.getIndexTypeOfTypeEx(source, r.c.numberType, r.c.anyType), r.c.getIndexTypeOfTypeEx(target, r.c.numberType, r.c.anyType), RecursionFlagsBoth, reportErrors)
			}
			// By flags alone, we know that the `target` is a readonly array while the source is a normal array or tuple
			// or `target` is an array and source is a tuple - in both cases the types cannot be identical, by construction
			return TernaryFalse
		case r.c.isGenericTupleType(source) && isTupleType(target) && !r.c.isGenericTupleType(target):
			constraint := r.c.getBaseConstraintOrType(source)
			if constraint != source {
				return r.isRelatedTo(constraint, target, RecursionFlagsSource, reportErrors)
			}
		case (r.relation == r.c.subtypeRelation || r.relation == r.c.strictSubtypeRelation) && r.c.isEmptyObjectType(target) && target.objectFlags&ObjectFlagsFreshLiteral != 0 && !r.c.isEmptyObjectType(source):
			return TernaryFalse
		}
		// Even if relationship doesn't hold for unions, intersections, or generic type references,
		// it may hold in a structural comparison.
		// In a check of the form X = A & B, we will have previously checked if A relates to X or B relates
		// to X. Failing both of those we want to check if the aggregation of A and B's members structurally
		// relates to X. Thus, we include intersection types on the source side here.
		if source.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0 && target.flags&TypeFlagsObject != 0 {
			// Report structural errors only if we haven't reported any errors yet
			reportStructuralErrors := reportErrors && r.errorChain == saveErrorState.errorChain && !sourceIsPrimitive
			result = r.propertiesRelatedTo(source, target, reportStructuralErrors, core.Set[string]{} /*excludedProperties*/, false /*optionalsOnly*/, intersectionState)
			if result != TernaryFalse {
				result &= r.signaturesRelatedTo(source, target, SignatureKindCall, reportStructuralErrors, intersectionState)
				if result != TernaryFalse {
					result &= r.signaturesRelatedTo(source, target, SignatureKindConstruct, reportStructuralErrors, intersectionState)
					if result != TernaryFalse {
						result &= r.indexSignaturesRelatedTo(source, target, sourceIsPrimitive, reportStructuralErrors, intersectionState)
					}
				}
			}
			if result != TernaryFalse {
				if !varianceCheckFailed {
					return result
				}
				if originalErrorChain != nil {
					r.errorChain = originalErrorChain
				} else if r.errorChain == nil {
					r.errorChain = saveErrorState.errorChain
				}
				// Use variance error (there is no structural one) and return false
			}
		}
		// If S is an object type and T is a discriminated union, S may be related to T if
		// there exists a constituent of T for every combination of the discriminants of S
		// with respect to T. We do not report errors here, as we will use the existing
		// error result from checking each constituent of the union.
		if source.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0 && target.flags&TypeFlagsUnion != 0 {
			objectOnlyTarget := r.c.extractTypesOfKind(target, TypeFlagsObject|TypeFlagsIntersection|TypeFlagsSubstitution)
			if objectOnlyTarget.flags&TypeFlagsUnion != 0 {
				result := r.typeRelatedToDiscriminatedType(source, objectOnlyTarget)
				if result != TernaryFalse {
					return result
				}
			}
		}
	}
	return TernaryFalse
}

func (r *Relater) typeArgumentsRelatedTo(sources []*Type, targets []*Type, variances []VarianceFlags, reportErrors bool, intersectionState IntersectionState) Ternary {
	if len(sources) != len(targets) && r.relation == r.c.identityRelation {
		return TernaryFalse
	}
	length := min(len(sources), len(targets))
	result := TernaryTrue
	for i := range length {
		// When variance information isn't available we default to covariance. This happens
		// in the process of computing variance information for recursive types and when
		// comparing 'this' type arguments.
		varianceFlags := VarianceFlagsCovariant
		if i < len(variances) {
			varianceFlags = variances[i]
		}
		variance := varianceFlags & VarianceFlagsVarianceMask
		// We ignore arguments for independent type parameters (because they're never witnessed).
		if variance != VarianceFlagsIndependent {
			s := sources[i]
			t := targets[i]
			var related Ternary
			if varianceFlags&VarianceFlagsUnmeasurable != 0 {
				// Even an `Unmeasurable` variance works out without a structural check if the source and target are _identical_.
				// We can't simply assume invariance, because `Unmeasurable` marks nonlinear relations, for example, a relation tained by
				// the `-?` modifier in a mapped type (where, no matter how the inputs are related, the outputs still might not be)
				if r.relation == r.c.identityRelation {
					related = r.isRelatedTo(s, t, RecursionFlagsBoth, false /*reportErrors*/)
				} else {
					related = r.c.compareTypesIdentical(s, t)
				}
			} else if variance == VarianceFlagsCovariant {
				related = r.isRelatedToEx(s, t, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
			} else if variance == VarianceFlagsContravariant {
				related = r.isRelatedToEx(t, s, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
			} else if variance == VarianceFlagsBivariant {
				// In the bivariant case we first compare contravariantly without reporting
				// errors. Then, if that doesn't succeed, we compare covariantly with error
				// reporting. Thus, error elaboration will be based on the the covariant check,
				// which is generally easier to reason about.
				related = r.isRelatedTo(t, s, RecursionFlagsBoth, false /*reportErrors*/)
				if related == TernaryFalse {
					related = r.isRelatedToEx(s, t, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
				}
			} else {
				// In the invariant case we first compare covariantly, and only when that
				// succeeds do we proceed to compare contravariantly. Thus, error elaboration
				// will typically be based on the covariant check.
				related = r.isRelatedToEx(s, t, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
				if related != TernaryFalse {
					related &= r.isRelatedToEx(t, s, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
				}
			}
			if related == TernaryFalse {
				return TernaryFalse
			}
			result &= related
		}
	}
	return result
}

// A type [P in S]: X is related to a type [Q in T]: Y if T is related to S and X' is
// related to Y, where X' is an instantiation of X in which P is replaced with Q. Notice
// that S and T are contra-variant whereas X and Y are co-variant.
func (r *Relater) mappedTypeRelatedTo(source *Type, target *Type, reportErrors bool) Ternary {
	// !!!
	// modifiersRelated := relation == c.comparableRelation || (ifelse(relation == c.identityRelation, c.getMappedTypeModifiers(source) == c.getMappedTypeModifiers(target), c.getCombinedMappedTypeOptionality(source) <= c.getCombinedMappedTypeOptionality(target)))
	// if modifiersRelated {
	// 	var result Ternary
	// 	targetConstraint := c.getConstraintTypeFromMappedType(target)
	// 	sourceConstraint := c.instantiateType(c.getConstraintTypeFromMappedType(source), ifelse(c.getCombinedMappedTypeOptionality(source) < 0, c.reportUnmeasurableMapper, c.reportUnreliableMapper))
	// 	if /* TODO(TS-TO-GO) EqualsToken BinaryExpression: result = isRelatedTo(targetConstraint, sourceConstraint, RecursionFlags.Both, reportErrors) */ TODO {
	// 		mapper := c.createTypeMapper([]TypeParameter{c.getTypeParameterFromMappedType(source)}, []TypeParameter{c.getTypeParameterFromMappedType(target)})
	// 		if c.instantiateType(c.getNameTypeFromMappedType(source), mapper) == c.instantiateType(c.getNameTypeFromMappedType(target), mapper) {
	// 			return result & isRelatedTo(c.instantiateType(c.getTemplateTypeFromMappedType(source), mapper), c.getTemplateTypeFromMappedType(target), RecursionFlagsBoth, reportErrors)
	// 		}
	// 	}
	// }
	return TernaryFalse
}

func (r *Relater) typeRelatedToDiscriminatedType(source *Type, target *Type) Ternary {
	// 1. Generate the combinations of discriminant properties & types 'source' can satisfy.
	//    a. If the number of combinations is above a set limit, the comparison is too complex.
	// 2. Filter 'target' to the subset of types whose discriminants exist in the matrix.
	//    a. If 'target' does not satisfy all discriminants in the matrix, 'source' is not related.
	// 3. For each type in the filtered 'target', determine if all non-discriminant properties of
	//    'target' are related to a property in 'source'.
	//
	// NOTE: See ~/tests/cases/conformance/types/typeRelationships/assignmentCompatibility/assignmentCompatWithDiscriminatedUnion.ts
	//       for examples.
	sourceProperties := r.c.getPropertiesOfType(source)
	sourcePropertiesFiltered := r.c.findDiscriminantProperties(sourceProperties, target)
	if len(sourcePropertiesFiltered) == 0 {
		return TernaryFalse
	}
	// Though we could compute the number of combinations as we generate
	// the matrix, this would incur additional memory overhead due to
	// array allocations. To reduce this overhead, we first compute
	// the number of combinations to ensure we will not surpass our
	// fixed limit before incurring the cost of any allocations:
	numCombinations := 1
	for _, sourceProperty := range sourcePropertiesFiltered {
		numCombinations *= countTypes(r.c.getNonMissingTypeOfSymbol(sourceProperty))
		if numCombinations == 0 || numCombinations > 25 {
			return TernaryFalse
		}
	}
	// Compute the set of types for each discriminant property.
	sourceDiscriminantTypes := make([][]*Type, len(sourcePropertiesFiltered))
	var excludedProperties core.Set[string]
	for i, sourceProperty := range sourcePropertiesFiltered {
		sourcePropertyType := r.c.getNonMissingTypeOfSymbol(sourceProperty)
		sourceDiscriminantTypes[i] = sourcePropertyType.Distributed()
		excludedProperties.Add(sourceProperty.Name)
	}
	// Build the cartesian product
	discriminantCombinations := make([][]*Type, numCombinations)
	for i := range numCombinations {
		combination := make([]*Type, len(sourceDiscriminantTypes))
		n := i
		for j := len(sourceDiscriminantTypes) - 1; j >= 0; j-- {
			sourceTypes := sourceDiscriminantTypes[j]
			length := len(sourceTypes)
			combination[j] = sourceTypes[n%length]
			n = n / length
		}
		discriminantCombinations[i] = combination
	}
	// Match each combination of the cartesian product of discriminant properties to one or more
	// constituents of 'target'. If any combination does not have a match then 'source' is not relatable.
	var matchingTypes []*Type
	for _, combination := range discriminantCombinations {
		hasMatch := false
	outer:
		for _, t := range target.Types() {
			for i := range sourcePropertiesFiltered {
				sourceProperty := sourcePropertiesFiltered[i]
				targetProperty := r.c.getPropertyOfType(t, sourceProperty.Name)
				if targetProperty == nil {
					continue outer
				}
				if sourceProperty == targetProperty {
					continue
				}
				// We compare the source property to the target in the context of a single discriminant type.
				related := r.propertyRelatedTo(source, target, sourceProperty, targetProperty, func(*ast.Symbol) *Type { return combination[i] },
					false /*reportErrors*/, IntersectionStateNone, r.c.strictNullChecks || r.relation == r.c.comparableRelation /*skipOptional*/)
				// If the target property could not be found, or if the properties were not related,
				// then this constituent is not a match.
				if related == TernaryFalse {
					continue outer
				}
			}
			matchingTypes = core.AppendIfUnique(matchingTypes, t)
			hasMatch = true
		}
		if !hasMatch {
			// We failed to match any type for this combination.
			return TernaryFalse
		}
	}
	// Compare the remaining non-discriminant properties of each match.
	result := TernaryTrue
	for _, t := range matchingTypes {
		result &= r.propertiesRelatedTo(source, t /*reportErrors*/, false, excludedProperties /*optionalsOnly*/, false, IntersectionStateNone)
		if result != TernaryFalse {
			result &= r.signaturesRelatedTo(source, t, SignatureKindCall /*reportErrors*/, false, IntersectionStateNone)
			if result != TernaryFalse {
				result &= r.signaturesRelatedTo(source, t, SignatureKindConstruct /*reportErrors*/, false, IntersectionStateNone)
				if result != TernaryFalse && !(isTupleType(source) && isTupleType(t)) {
					// Comparing numeric index types when both `source` and `type` are tuples is unnecessary as the
					// element types should be sufficiently covered by `propertiesRelatedTo`. It also causes problems
					// with index type assignability as the types for the excluded discriminants are still included
					// in the index type.
					result &= r.indexSignaturesRelatedTo(source, t /*sourceIsPrimitive*/, false /*reportErrors*/, false, IntersectionStateNone)
				}
			}
		}
		if result == TernaryFalse {
			return result
		}
	}
	return result
}

func (r *Relater) propertiesRelatedTo(source *Type, target *Type, reportErrors bool, excludedProperties core.Set[string], optionalsOnly bool, intersectionState IntersectionState) Ternary {
	if r.relation == r.c.identityRelation {
		return r.propertiesIdenticalTo(source, target, excludedProperties)
	}
	result := TernaryTrue
	if isTupleType(target) {
		if r.c.isArrayOrTupleType(source) {
			if !target.TargetTupleType().readonly && (r.c.isReadonlyArrayType(source) || isTupleType(source) && source.TargetTupleType().readonly) {
				return TernaryFalse
			}
			sourceArity := r.c.getTypeReferenceArity(source)
			targetArity := r.c.getTypeReferenceArity(target)
			var sourceRest bool
			if isTupleType(source) {
				sourceRest = source.TargetTupleType().combinedFlags&ElementFlagsRest != 0
			} else {
				sourceRest = true
			}
			targetHasRestElement := target.TargetTupleType().combinedFlags&ElementFlagsVariable != 0
			var sourceMinLength int
			if isTupleType(source) {
				sourceMinLength = source.TargetTupleType().minLength
			} else {
				sourceMinLength = 0
			}
			targetMinLength := target.TargetTupleType().minLength
			if !sourceRest && sourceArity < targetMinLength {
				if reportErrors {
					r.reportError(diagnostics.Source_has_0_element_s_but_target_requires_1, sourceArity, targetMinLength)
				}
				return TernaryFalse
			}
			if !targetHasRestElement && targetArity < sourceMinLength {
				if reportErrors {
					r.reportError(diagnostics.Source_has_0_element_s_but_target_allows_only_1, sourceMinLength, targetArity)
				}
				return TernaryFalse
			}
			if !targetHasRestElement && (sourceRest || targetArity < sourceArity) {
				if reportErrors {
					if sourceMinLength < targetMinLength {
						r.reportError(diagnostics.Target_requires_0_element_s_but_source_may_have_fewer, targetMinLength)
					} else {
						r.reportError(diagnostics.Target_allows_only_0_element_s_but_source_may_have_more, targetArity)
					}
				}
				return TernaryFalse
			}
			sourceTypeArguments := r.c.getTypeArguments(source)
			targetTypeArguments := r.c.getTypeArguments(target)
			targetStartCount := getStartElementCount(target.TargetTupleType(), ElementFlagsNonRest)
			targetEndCount := getEndElementCount(target.TargetTupleType(), ElementFlagsNonRest)
			canExcludeDiscriminants := excludedProperties.Len() != 0
			for sourcePosition := range sourceArity {
				var sourceFlags ElementFlags
				if isTupleType(source) {
					sourceFlags = source.TargetTupleType().elementInfos[sourcePosition].flags
				} else {
					sourceFlags = ElementFlagsRest
				}
				sourcePositionFromEnd := sourceArity - 1 - sourcePosition
				var targetPosition int
				if targetHasRestElement && sourcePosition >= targetStartCount {
					targetPosition = targetArity - 1 - min(sourcePositionFromEnd, targetEndCount)
				} else {
					targetPosition = sourcePosition
				}
				targetFlags := target.TargetTupleType().elementInfos[targetPosition].flags
				if targetFlags&ElementFlagsVariadic != 0 && sourceFlags&ElementFlagsVariadic == 0 {
					if reportErrors {
						r.reportError(diagnostics.Source_provides_no_match_for_variadic_element_at_position_0_in_target, targetPosition)
					}
					return TernaryFalse
				}
				if sourceFlags&ElementFlagsVariadic != 0 && targetFlags&ElementFlagsVariable == 0 {
					if reportErrors {
						r.reportError(diagnostics.Variadic_element_at_position_0_in_source_does_not_match_element_at_position_1_in_target, sourcePosition, targetPosition)
					}
					return TernaryFalse
				}
				if targetFlags&ElementFlagsRequired != 0 && sourceFlags&ElementFlagsRequired == 0 {
					if reportErrors {
						r.reportError(diagnostics.Source_provides_no_match_for_required_element_at_position_0_in_target, targetPosition)
					}
					return TernaryFalse
				}
				// We can only exclude discriminant properties if we have not yet encountered a variable-length element.
				if canExcludeDiscriminants {
					if sourceFlags&ElementFlagsVariable != 0 || targetFlags&ElementFlagsVariable != 0 {
						canExcludeDiscriminants = false
					}
					if canExcludeDiscriminants && excludedProperties.Has(strconv.Itoa(sourcePosition)) {
						continue
					}
				}
				sourceType := r.c.removeMissingType(sourceTypeArguments[sourcePosition], sourceFlags&targetFlags&ElementFlagsOptional != 0)
				targetType := targetTypeArguments[targetPosition]
				var targetCheckType *Type
				if sourceFlags&ElementFlagsVariadic != 0 && targetFlags&ElementFlagsRest != 0 {
					targetCheckType = r.c.createArrayType(targetType)
				} else {
					targetCheckType = r.c.removeMissingType(targetType, targetFlags&ElementFlagsOptional != 0)
				}
				related := r.isRelatedToEx(sourceType, targetCheckType, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
				if related == TernaryFalse {
					if reportErrors && (targetArity > 1 || sourceArity > 1) {
						if targetHasRestElement && sourcePosition >= targetStartCount && sourcePositionFromEnd >= targetEndCount && targetStartCount != sourceArity-targetEndCount-1 {
							r.reportError(diagnostics.Type_at_positions_0_through_1_in_source_is_not_compatible_with_type_at_position_2_in_target, targetStartCount, sourceArity-targetEndCount-1, targetPosition)
						} else {
							r.reportError(diagnostics.Type_at_position_0_in_source_is_not_compatible_with_type_at_position_1_in_target, sourcePosition, targetPosition)
						}
					}
					return TernaryFalse
				}
				result &= related
			}
			return result
		}
		if target.TargetTupleType().combinedFlags&ElementFlagsVariable != 0 {
			return TernaryFalse
		}
	}
	requireOptionalProperties := (r.relation == r.c.subtypeRelation || r.relation == r.c.strictSubtypeRelation) && isObjectLiteralType(source) && !r.c.isEmptyArrayLiteralType(source) && !isTupleType(source)
	unmatchedProperty := r.c.getUnmatchedProperty(source, target, requireOptionalProperties, false /*matchDiscriminantProperties*/)
	if unmatchedProperty != nil {
		if reportErrors && r.c.shouldReportUnmatchedPropertyError(source, target) {
			r.reportUnmatchedProperty(source, target, unmatchedProperty, requireOptionalProperties)
		}
		return TernaryFalse
	}
	if isObjectLiteralType(target) {
		for _, sourceProp := range excludeProperties(r.c.getPropertiesOfType(source), excludedProperties) {
			if r.c.getPropertyOfObjectType(target, sourceProp.Name) == nil {
				sourceType := r.c.getTypeOfSymbol(sourceProp)
				if sourceType.flags&TypeFlagsUndefined == 0 {
					if reportErrors {
						r.reportError(diagnostics.Property_0_does_not_exist_on_type_1, r.c.symbolToString(sourceProp), r.c.typeToString(target))
					}
					return TernaryFalse
				}
			}
		}
	}
	// We only call this for union target types when we're attempting to do excess property checking - in those cases, we want to get _all possible props_
	// from the target union, across all members
	properties := r.c.getPropertiesOfType(target)
	numericNamesOnly := isTupleType(source) && isTupleType(target)
	for _, targetProp := range excludeProperties(properties, excludedProperties) {
		name := targetProp.Name
		if targetProp.Flags&ast.SymbolFlagsPrototype == 0 && (!numericNamesOnly || isNumericLiteralName(name) || name == "length") && (!optionalsOnly || targetProp.Flags&ast.SymbolFlagsOptional != 0) {
			sourceProp := r.c.getPropertyOfType(source, name)
			if sourceProp != nil && sourceProp != targetProp {
				related := r.propertyRelatedTo(source, target, sourceProp, targetProp, r.c.getNonMissingTypeOfSymbol, reportErrors, intersectionState, r.relation == r.c.comparableRelation)
				if related == TernaryFalse {
					return TernaryFalse
				}
				result &= related
			}
		}
	}
	return result
}

func (r *Relater) propertyRelatedTo(source *Type, target *Type, sourceProp *ast.Symbol, targetProp *ast.Symbol, getTypeOfSourceProperty func(sym *ast.Symbol) *Type, reportErrors bool, intersectionState IntersectionState, skipOptional bool) Ternary {
	sourcePropFlags := getDeclarationModifierFlagsFromSymbol(sourceProp)
	targetPropFlags := getDeclarationModifierFlagsFromSymbol(targetProp)
	switch {
	case sourcePropFlags&ast.ModifierFlagsPrivate != 0 || targetPropFlags&ast.ModifierFlagsPrivate != 0:
		if sourceProp.ValueDeclaration != targetProp.ValueDeclaration {
			if reportErrors {
				if sourcePropFlags&ast.ModifierFlagsPrivate != 0 && targetPropFlags&ast.ModifierFlagsPrivate != 0 {
					r.reportError(diagnostics.Types_have_separate_declarations_of_a_private_property_0, r.c.symbolToString(targetProp))
				} else {
					r.reportError(diagnostics.Property_0_is_private_in_type_1_but_not_in_type_2, r.c.symbolToString(targetProp), r.c.typeToString(ifElse(sourcePropFlags&ast.ModifierFlagsPrivate != 0, source, target)), r.c.typeToString(ifElse(sourcePropFlags&ast.ModifierFlagsPrivate != 0, target, source)))
				}
			}
			return TernaryFalse
		}
	case targetPropFlags&ast.ModifierFlagsProtected != 0:
		if !r.c.isValidOverrideOf(sourceProp, targetProp) {
			if reportErrors {
				sourceType := r.c.getDeclaringClass(sourceProp)
				if sourceType == nil {
					sourceType = source
				}
				targetType := r.c.getDeclaringClass(targetProp)
				if targetType == nil {
					targetType = target
				}
				r.reportError(diagnostics.Property_0_is_protected_but_type_1_is_not_a_class_derived_from_2, r.c.symbolToString(targetProp), r.c.typeToString(sourceType), r.c.typeToString(targetType))
			}
			return TernaryFalse
		}
	case sourcePropFlags&ast.ModifierFlagsProtected != 0:
		if reportErrors {
			r.reportError(diagnostics.Property_0_is_protected_in_type_1_but_public_in_type_2, r.c.symbolToString(targetProp), r.c.typeToString(source), r.c.typeToString(target))
		}
		return TernaryFalse
	}
	// Ensure {readonly a: whatever} is not a subtype of {a: whatever},
	// while {a: whatever} is a subtype of {readonly a: whatever}.
	// This ensures the subtype relationship is ordered, and preventing declaration order
	// from deciding which type "wins" in union subtype reduction.
	// They're still assignable to one another, since `readonly` doesn't affect assignability.
	// This is only applied during the strictSubtypeRelation -- currently used in subtype reduction
	if r.relation == r.c.strictSubtypeRelation && r.c.isReadonlySymbol(sourceProp) && !r.c.isReadonlySymbol(targetProp) {
		return TernaryFalse
	}
	// If the target comes from a partial union prop, allow `undefined` in the target type
	related := r.isPropertySymbolTypeRelated(sourceProp, targetProp, getTypeOfSourceProperty, reportErrors, intersectionState)
	if related == TernaryFalse {
		if reportErrors {
			r.reportError(diagnostics.Types_of_property_0_are_incompatible, r.c.symbolToString(targetProp))
		}
		return TernaryFalse
	}
	// When checking for comparability, be more lenient with optional properties.
	if !skipOptional && sourceProp.Flags&ast.SymbolFlagsOptional != 0 && targetProp.Flags&ast.SymbolFlagsClassMember != 0 && targetProp.Flags&ast.SymbolFlagsOptional == 0 {
		// TypeScript 1.0 spec (April 2014): 3.8.3
		// S is a subtype of a type T, and T is a supertype of S if ...
		// S' and T are object types and, for each member M in T..
		// M is a property and S' contains a property N where
		// if M is a required property, N is also a required property
		// (M - property in T)
		// (N - property in S)
		if reportErrors {
			r.reportError(diagnostics.Property_0_is_optional_in_type_1_but_required_in_type_2, r.c.symbolToString(targetProp), r.c.typeToString(source), r.c.typeToString(target))
		}
		return TernaryFalse
	}
	return related
}

func (r *Relater) isPropertySymbolTypeRelated(sourceProp *ast.Symbol, targetProp *ast.Symbol, getTypeOfSourceProperty func(sym *ast.Symbol) *Type, reportErrors bool, intersectionState IntersectionState) Ternary {
	targetIsOptional := r.c.strictNullChecks && targetProp.CheckFlags&ast.CheckFlagsPartial != 0
	effectiveTarget := r.c.addOptionalityEx(r.c.getNonMissingTypeOfSymbol(targetProp), false /*isProperty*/, targetIsOptional)
	effectiveSource := getTypeOfSourceProperty(sourceProp)
	return r.isRelatedToEx(effectiveSource, effectiveTarget, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
}

func (r *Relater) reportUnmatchedProperty(source *Type, target *Type, unmatchedProperty *ast.Symbol, requireOptionalProperties bool) {
	// give specific error in case where private names have the same description
	if unmatchedProperty.ValueDeclaration != nil &&
		unmatchedProperty.ValueDeclaration.Name() != nil &&
		ast.IsPrivateIdentifier(unmatchedProperty.ValueDeclaration.Name()) &&
		source.symbol != nil &&
		source.symbol.Flags&ast.SymbolFlagsClass != 0 {
		privateIdentifierDescription := unmatchedProperty.ValueDeclaration.Name().Text()
		symbolTableKey := getSymbolNameForPrivateIdentifier(source.symbol, privateIdentifierDescription)
		if r.c.getPropertyOfType(source, symbolTableKey) != nil {
			sourceName := declarationNameToString(getNameOfDeclaration(source.symbol.ValueDeclaration))
			targetName := declarationNameToString(getNameOfDeclaration(target.symbol.ValueDeclaration))
			r.reportError(diagnostics.Property_0_in_type_1_refers_to_a_different_member_that_cannot_be_accessed_from_within_type_2, privateIdentifierDescription, sourceName, targetName)
			return
		}
	}
	props := slices.Collect(r.c.getUnmatchedProperties(source, target, requireOptionalProperties, false /*matchDiscriminantProperties*/))
	if len(props) == 1 {
		propName := r.c.symbolToString(unmatchedProperty)
		r.reportError(diagnostics.Property_0_is_missing_in_type_1_but_required_in_type_2, unmatchedProperty.Name, r.c.typeToString(source), r.c.typeToString(target))
		if len(unmatchedProperty.Declarations) != 0 {
			r.relatedInfo = append(r.relatedInfo, createDiagnosticForNode(unmatchedProperty.Declarations[0], diagnostics.X_0_is_declared_here, propName))
		}
	} else if r.tryElaborateArrayLikeErrors(source, target, false /*reportErrors*/) {
		if len(props) > 5 {
			propNames := strings.Join(core.Map(props[:4], r.c.symbolToString), ", ")
			r.reportError(diagnostics.Type_0_is_missing_the_following_properties_from_type_1_Colon_2_and_3_more, r.c.typeToString(source), r.c.typeToString(target), propNames, len(props)-4)
		} else {
			propNames := strings.Join(core.Map(props, r.c.symbolToString), ", ")
			r.reportError(diagnostics.Type_0_is_missing_the_following_properties_from_type_1_Colon_2, r.c.typeToString(source), r.c.typeToString(target), propNames)
		}
	}
}

func (r *Relater) tryElaborateArrayLikeErrors(source *Type, target *Type, reportErrors bool) bool {
	/**
	 * The spec for elaboration is:
	 * - If the source is a readonly tuple and the target is a mutable array or tuple, elaborate on mutability and skip property elaborations.
	 * - If the source is a tuple then skip property elaborations if the target is an array or tuple.
	 * - If the source is a readonly array and the target is a mutable array or tuple, elaborate on mutability and skip property elaborations.
	 * - If the source an array then skip property elaborations if the target is a tuple.
	 */
	if isTupleType(source) {
		if source.TargetTupleType().readonly && r.c.isMutableArrayOrTuple(target) {
			if reportErrors {
				r.reportError(diagnostics.The_type_0_is_readonly_and_cannot_be_assigned_to_the_mutable_type_1, r.c.typeToString(source), r.c.typeToString(target))
			}
			return false
		}
		return r.c.isArrayOrTupleType(target)
	}
	if r.c.isReadonlyArrayType(source) && r.c.isMutableArrayOrTuple(target) {
		if reportErrors {
			r.reportError(diagnostics.The_type_0_is_readonly_and_cannot_be_assigned_to_the_mutable_type_1, r.c.typeToString(source), r.c.typeToString(target))
		}
		return false
	}
	if isTupleType(target) {
		return r.c.isArrayType(source)
	}
	return true
}

func (r *Relater) tryElaborateErrorsForPrimitivesAndObjects(source *Type, target *Type) {
	if (source == r.c.globalStringType && target == r.c.stringType) ||
		(source == r.c.globalNumberType && target == r.c.numberType) ||
		(source == r.c.globalBooleanType && target == r.c.booleanType) ||
		(source == r.c.getGlobalESSymbolType() && target == r.c.esSymbolType) {
		r.reportError(diagnostics.X_0_is_a_primitive_but_1_is_a_wrapper_object_Prefer_using_0_when_possible, r.c.typeToString(target), r.c.typeToString(source))
	}
}

func (r *Relater) propertiesIdenticalTo(source *Type, target *Type, excludedProperties core.Set[string]) Ternary {
	if source.flags&TypeFlagsObject == 0 || target.flags&TypeFlagsObject == 0 {
		return TernaryFalse
	}
	sourceProperties := excludeProperties(r.c.getPropertiesOfObjectType(source), excludedProperties)
	targetProperties := excludeProperties(r.c.getPropertiesOfObjectType(target), excludedProperties)
	if len(sourceProperties) != len(targetProperties) {
		return TernaryFalse
	}
	result := TernaryTrue
	for _, sourceProp := range sourceProperties {
		targetProp := r.c.getPropertyOfObjectType(target, sourceProp.Name)
		if targetProp == nil {
			return TernaryFalse
		}
		related := r.c.compareProperties(sourceProp, targetProp, r.isRelatedToSimple)
		if related == TernaryFalse {
			return TernaryFalse
		}
		result &= related
	}
	return result
}

func (r *Relater) signaturesRelatedTo(source *Type, target *Type, kind SignatureKind, reportErrors bool, intersectionState IntersectionState) Ternary {
	if r.relation == r.c.identityRelation {
		return r.signaturesIdenticalTo(source, target, kind)
	}
	if target == r.c.anyFunctionType || source == r.c.anyFunctionType {
		return TernaryTrue
	}
	sourceSignatures := r.c.getSignaturesOfType(source, kind)
	targetSignatures := r.c.getSignaturesOfType(target, kind)
	if kind == SignatureKindConstruct && len(sourceSignatures) != 0 && len(targetSignatures) != 0 {
		sourceIsAbstract := sourceSignatures[0].flags&SignatureFlagsAbstract != 0
		targetIsAbstract := targetSignatures[0].flags&SignatureFlagsAbstract != 0
		if sourceIsAbstract && !targetIsAbstract {
			// An abstract constructor type is not assignable to a non-abstract constructor type
			// as it would otherwise be possible to new an abstract class. Note that the assignability
			// check we perform for an extends clause excludes construct signatures from the target,
			// so this check never proceeds.
			if reportErrors {
				r.reportError(diagnostics.Cannot_assign_an_abstract_constructor_type_to_a_non_abstract_constructor_type)
			}
			return TernaryFalse
		}
		if !r.constructorVisibilitiesAreCompatible(sourceSignatures[0], targetSignatures[0], reportErrors) {
			return TernaryFalse
		}
	}
	result := TernaryTrue
	switch {
	case source.objectFlags&ObjectFlagsInstantiated != 0 && target.objectFlags&ObjectFlagsInstantiated != 0 && source.symbol == target.symbol ||
		source.objectFlags&ObjectFlagsReference != 0 && target.objectFlags&ObjectFlagsReference != 0 && source.Target() == target.Target():
		// We have instantiations of the same anonymous type (which typically will be the type of a
		// method). Simply do a pairwise comparison of the signatures in the two signature lists instead
		// of the much more expensive N * M comparison matrix we explore below. We erase type parameters
		// as they are known to always be the same.
		for i := range targetSignatures {
			related := r.signatureRelatedTo(sourceSignatures[i], targetSignatures[i], true /*erase*/, reportErrors, intersectionState)
			if related == TernaryFalse {
				return TernaryFalse
			}
			result &= related
		}
	case len(sourceSignatures) == 1 && len(targetSignatures) == 1:
		// For simple functions (functions with a single signature) we only erase type parameters for
		// the comparable relation. Otherwise, if the source signature is generic, we instantiate it
		// in the context of the target signature before checking the relationship. Ideally we'd do
		// this regardless of the number of signatures, but the potential costs are prohibitive due
		// to the quadratic nature of the logic below.
		eraseGenerics := r.relation == r.c.comparableRelation
		result = r.signatureRelatedTo(sourceSignatures[0], targetSignatures[0], eraseGenerics, reportErrors, intersectionState)
	default:
	outer:
		for _, t := range targetSignatures {
			saveErrorState := r.getErrorState()
			// Only elaborate errors from the first failure
			shouldElaborateErrors := reportErrors
			for _, s := range sourceSignatures {
				related := r.signatureRelatedTo(s, t, true /*erase*/, shouldElaborateErrors, intersectionState)
				if related != TernaryFalse {
					result &= related
					r.restoreErrorState(saveErrorState)
					continue outer
				}
				shouldElaborateErrors = false
			}
			if shouldElaborateErrors {
				r.reportError(diagnostics.Type_0_provides_no_match_for_the_signature_1, r.c.typeToString(source), r.c.signatureToString(t))
			}
			return TernaryFalse
		}
	}
	return result
}

func (r *Relater) constructorVisibilitiesAreCompatible(sourceSignature *Signature, targetSignature *Signature, reportErrors bool) bool {
	if sourceSignature.declaration == nil || targetSignature.declaration == nil {
		return true
	}
	sourceAccessibility := getEffectiveModifierFlags(sourceSignature.declaration) & ast.ModifierFlagsNonPublicAccessibilityModifier
	targetAccessibility := getEffectiveModifierFlags(targetSignature.declaration) & ast.ModifierFlagsNonPublicAccessibilityModifier
	// A public, protected and private signature is assignable to a private signature.
	if targetAccessibility == ast.ModifierFlagsPrivate {
		return true
	}
	// A public and protected signature is assignable to a protected signature.
	if targetAccessibility == ast.ModifierFlagsProtected && sourceAccessibility != ast.ModifierFlagsPrivate {
		return true
	}
	// Only a public signature is assignable to public signature.
	if targetAccessibility != ast.ModifierFlagsProtected && sourceAccessibility == 0 {
		return true
	}
	if reportErrors {
		r.reportError(diagnostics.Cannot_assign_a_0_constructor_type_to_a_1_constructor_type, visibilityToString(sourceAccessibility), visibilityToString(targetAccessibility))
	}
	return false
}

// See signatureAssignableTo, compareSignaturesIdentical
func (r *Relater) signatureRelatedTo(source *Signature, target *Signature, erase bool, reportErrors bool, intersectionState IntersectionState) Ternary {
	checkMode := SignatureCheckModeNone
	switch {
	case r.relation == r.c.subtypeRelation:
		checkMode = SignatureCheckModeStrictTopSignature
	case r.relation == r.c.strictSubtypeRelation:
		checkMode = SignatureCheckModeStrictTopSignature | SignatureCheckModeStrictArity
	}
	if erase {
		source = r.c.getErasedSignature(source)
		target = r.c.getErasedSignature(target)
	}
	isRelatedToWorker := func(source *Type, target *Type, reportErrors bool) Ternary {
		return r.isRelatedToEx(source, target, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
	}
	return r.c.compareSignaturesRelated(source, target, checkMode, reportErrors, r.reportError, isRelatedToWorker, r.c.reportUnreliableMapper)
}

func (r *Relater) signaturesIdenticalTo(source *Type, target *Type, kind SignatureKind) Ternary {
	sourceSignatures := r.c.getSignaturesOfType(source, kind)
	targetSignatures := r.c.getSignaturesOfType(target, kind)
	if len(sourceSignatures) != len(targetSignatures) {
		return TernaryFalse
	}
	result := TernaryTrue
	for i := range sourceSignatures {
		related := r.c.compareSignaturesIdentical(sourceSignatures[i], targetSignatures[i], false /*partialMatch*/, false /*ignoreThisTypes*/, false /*ignoreReturnTypes*/, r.isRelatedToSimple)
		if related == 0 {
			return TernaryFalse
		}
		result &= related
	}
	return result
}

func (r *Relater) indexSignaturesRelatedTo(source *Type, target *Type, sourceIsPrimitive bool, reportErrors bool, intersectionState IntersectionState) Ternary {
	if r.relation == r.c.identityRelation {
		return r.indexSignaturesIdenticalTo(source, target)
	}
	indexInfos := r.c.getIndexInfosOfType(target)
	targetHasStringIndex := core.Some(indexInfos, func(info *IndexInfo) bool { return info.keyType == r.c.stringType })
	result := TernaryTrue
	for _, targetInfo := range indexInfos {
		var related Ternary
		switch {
		case r.relation != r.c.strictSubtypeRelation && !sourceIsPrimitive && targetHasStringIndex && targetInfo.valueType.flags&TypeFlagsAny != 0:
			related = TernaryTrue
		case r.c.isGenericMappedType(source) && targetHasStringIndex:
			related = r.isRelatedTo(r.c.getTemplateTypeFromMappedType(source), targetInfo.valueType, RecursionFlagsBoth, reportErrors)
		default:
			related = r.typeRelatedToIndexInfo(source, targetInfo, reportErrors, intersectionState)
		}
		if related == TernaryFalse {
			return TernaryFalse
		}
		result &= related
	}
	return result
}

func (r *Relater) typeRelatedToIndexInfo(source *Type, targetInfo *IndexInfo, reportErrors bool, intersectionState IntersectionState) Ternary {
	sourceInfo := r.c.getApplicableIndexInfo(source, targetInfo.keyType)
	if sourceInfo != nil {
		return r.indexInfoRelatedTo(sourceInfo, targetInfo, reportErrors, intersectionState)
	}
	// Intersection constituents are never considered to have an inferred index signature. Also, in the strict subtype relation,
	// only fresh object literals are considered to have inferred index signatures. This ensures { [x: string]: xxx } <: {} but
	// not vice-versa. Without this rule, those types would be mutual strict subtypes.
	if intersectionState&IntersectionStateSource == 0 && (r.relation != r.c.strictSubtypeRelation || source.objectFlags&ObjectFlagsFreshLiteral != 0) && r.c.isObjectTypeWithInferableIndex(source) {
		return r.membersRelatedToIndexInfo(source, targetInfo, reportErrors, intersectionState)
	}
	if reportErrors {
		r.reportError(diagnostics.Index_signature_for_type_0_is_missing_in_type_1, r.c.typeToString(targetInfo.keyType), r.c.typeToString(source))
	}
	return TernaryFalse
}

/**
 * Return true if type was inferred from an object literal, written as an object type literal, or is the shape of a module
 * with no call or construct signatures.
 */
func (c *Checker) isObjectTypeWithInferableIndex(t *Type) bool {
	if t.flags&TypeFlagsIntersection != 0 {
		return core.Every(t.Types(), c.isObjectTypeWithInferableIndex)
	}
	return t.symbol != nil && t.symbol.Flags&(ast.SymbolFlagsObjectLiteral|ast.SymbolFlagsTypeLiteral|ast.SymbolFlagsEnum|ast.SymbolFlagsValueModule) != 0 &&
		t.symbol.Flags&ast.SymbolFlagsClass == 0 && !c.typeHasCallOrConstructSignatures(t) ||
		t.objectFlags&ObjectFlagsObjectRestType != 0 ||
		t.objectFlags&ObjectFlagsReverseMapped != 0 && c.isObjectTypeWithInferableIndex(t.AsReverseMappedType().source)
}

func (r *Relater) membersRelatedToIndexInfo(source *Type, targetInfo *IndexInfo, reportErrors bool, intersectionState IntersectionState) Ternary {
	result := TernaryTrue
	keyType := targetInfo.keyType
	var props []*ast.Symbol
	if source.flags&TypeFlagsIntersection != 0 {
		props = r.c.getPropertiesOfUnionOrIntersectionType(source)
	} else {
		props = r.c.getPropertiesOfObjectType(source)
	}
	for _, prop := range props {
		// Skip over ignored JSX and symbol-named members
		if isIgnoredJsxProperty(source, prop) {
			continue
		}
		if r.c.isApplicableIndexType(r.c.getLiteralTypeFromProperty(prop, TypeFlagsStringOrNumberLiteralOrUnique, false), keyType) {
			propType := r.c.getNonMissingTypeOfSymbol(prop)
			var t *Type
			if r.c.exactOptionalPropertyTypes || propType.flags&TypeFlagsUndefined != 0 || keyType == r.c.numberType || prop.Flags&ast.SymbolFlagsOptional == 0 {
				t = propType
			} else {
				t = r.c.getTypeWithFacts(propType, TypeFactsNEUndefined)
			}
			related := r.isRelatedToEx(t, targetInfo.valueType, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
			if related == TernaryFalse {
				if reportErrors {
					r.reportError(diagnostics.Property_0_is_incompatible_with_index_signature, r.c.symbolToString(prop))
				}
				return TernaryFalse
			}
			result &= related
		}
	}
	for _, info := range r.c.getIndexInfosOfType(source) {
		if r.c.isApplicableIndexType(info.keyType, keyType) {
			related := r.indexInfoRelatedTo(info, targetInfo, reportErrors, intersectionState)
			if !(related != 0) {
				return TernaryFalse
			}
			result &= related
		}
	}
	return result
}

func (r *Relater) indexInfoRelatedTo(sourceInfo *IndexInfo, targetInfo *IndexInfo, reportErrors bool, intersectionState IntersectionState) Ternary {
	related := r.isRelatedToEx(sourceInfo.valueType, targetInfo.valueType, RecursionFlagsBoth, reportErrors, nil /*headMessage*/, intersectionState)
	if related == TernaryFalse && reportErrors {
		if sourceInfo.keyType == targetInfo.keyType {
			r.reportError(diagnostics.X_0_index_signatures_are_incompatible, r.c.typeToString(sourceInfo.keyType))
		} else {
			r.reportError(diagnostics.X_0_and_1_index_signatures_are_incompatible, r.c.typeToString(sourceInfo.keyType), r.c.typeToString(targetInfo.keyType))
		}
	}
	return related
}

func (r *Relater) indexSignaturesIdenticalTo(source *Type, target *Type) Ternary {
	sourceInfos := r.c.getIndexInfosOfType(source)
	targetInfos := r.c.getIndexInfosOfType(target)
	if len(sourceInfos) != len(targetInfos) {
		return TernaryFalse
	}
	for _, targetInfo := range targetInfos {
		sourceInfo := r.c.getIndexInfoOfType(source, targetInfo.keyType)
		if !(sourceInfo != nil && r.isRelatedTo(sourceInfo.valueType, targetInfo.valueType, RecursionFlagsBoth, false) != TernaryFalse && sourceInfo.isReadonly == targetInfo.isReadonly) {
			return TernaryFalse
		}
	}
	return TernaryTrue
}

func (r *Relater) reportErrorResults(originalSource *Type, originalTarget *Type, source *Type, target *Type, headMessage *diagnostics.Message) {
	sourceHasBase := r.c.getSingleBaseForNonAugmentingSubtype(originalSource) != nil
	targetHasBase := r.c.getSingleBaseForNonAugmentingSubtype(originalTarget) != nil
	if originalSource.alias != nil || sourceHasBase {
		source = originalSource
	}
	if originalTarget.alias != nil || targetHasBase {
		target = originalTarget
	}
	if source.flags&TypeFlagsObject != 0 && target.flags&TypeFlagsObject != 0 {
		r.tryElaborateArrayLikeErrors(source, target, true /*reportErrors*/)
	}
	switch {
	case source.flags&TypeFlagsObject != 0 && target.flags&TypeFlagsPrimitive != 0:
		r.tryElaborateErrorsForPrimitivesAndObjects(source, target)
	case source.symbol != nil && source.flags&TypeFlagsObject != 0 && r.c.globalObjectType == source:
		r.reportError(diagnostics.The_Object_type_is_assignable_to_very_few_other_types_Did_you_mean_to_use_the_any_type_instead)
	case source.objectFlags&ObjectFlagsJsxAttributes != 0 && target.flags&TypeFlagsIntersection != 0:
		// !!!
		// targetTypes := target.Types()
		// intrinsicAttributes := c.getJsxType(JsxNames.IntrinsicAttributes, errorNode)
		// intrinsicClassAttributes := c.getJsxType(JsxNames.IntrinsicClassAttributes, errorNode)
		// if !c.isErrorType(intrinsicAttributes) && !c.isErrorType(intrinsicClassAttributes) && (contains(targetTypes, intrinsicAttributes) || contains(targetTypes, intrinsicClassAttributes)) {
		// 	// do not report top error
		// 	return
		// }
	case originalTarget.flags&TypeFlagsIntersection != 0 && originalTarget.objectFlags&ObjectFlagsIsNeverIntersection != 0:
		message := diagnostics.The_intersection_0_was_reduced_to_never_because_property_1_has_conflicting_types_in_some_constituents
		prop := core.Find(r.c.getPropertiesOfUnionOrIntersectionType(originalTarget), r.c.isDiscriminantWithNeverType)
		if prop == nil {
			message = diagnostics.The_intersection_0_was_reduced_to_never_because_property_1_exists_in_multiple_constituents_and_is_private_in_some
			prop = core.Find(r.c.getPropertiesOfUnionOrIntersectionType(originalTarget), isConflictingPrivateProperty)
		}
		if prop != nil {
			r.reportError(message, r.c.typeToStringEx(originalTarget, nil /*enclosingDeclaration*/, TypeFormatFlagsNoTypeReduction), r.c.symbolToString(prop))
		}
	}
	// !!! Logic having to do with canonical diagnostics for deduplication purposes
	r.reportRelationError(headMessage, source, target)
	if source.flags&TypeFlagsTypeParameter != 0 && source.symbol != nil && len(source.symbol.Declarations) != 0 && r.c.getConstraintOfType(source) == nil {
		syntheticParam := r.c.cloneTypeParameter(source)
		syntheticParam.AsTypeParameter().constraint = r.c.instantiateType(target, newSimpleTypeMapper(source, syntheticParam))
		if r.c.hasNonCircularBaseConstraint(syntheticParam) {
			targetConstraintString := r.c.typeToString(target)
			r.relatedInfo = append(r.relatedInfo, NewDiagnosticForNode(source.symbol.Declarations[0], diagnostics.This_type_parameter_might_need_an_extends_0_constraint, targetConstraintString))
		}
	}
}

func (r *Relater) reportRelationError(message *diagnostics.Message, source *Type, target *Type) {
	sourceType, targetType := r.c.getTypeNamesForErrorDisplay(source, target)
	generalizedSource := source
	generalizedSourceType := sourceType
	// Don't generalize on 'never' - we really want the original type
	// to be displayed for use-cases like 'assertNever'.
	if target.flags&TypeFlagsNever == 0 && isLiteralType(source) && !r.c.typeCouldHaveTopLevelSingletonTypes(target) {
		generalizedSource = r.c.getBaseTypeOfLiteralType(source)
		// !!! Debug.assert(!c.isTypeAssignableTo(generalizedSource, target), "generalized source shouldn't be assignable")
		generalizedSourceType = r.c.getTypeNameForErrorDisplay(generalizedSource)
	}
	// If `target` is of indexed access type (and `source` it is not), we use the object type of `target` for better error reporting
	var targetFlags TypeFlags
	if target.flags&TypeFlagsIndexedAccess != 0 && source.flags&TypeFlagsIndexedAccess == 0 {
		targetFlags = target.AsIndexedAccessType().objectType.flags
	} else {
		targetFlags = target.flags
	}
	if targetFlags&TypeFlagsTypeParameter != 0 && target != r.c.markerSuperTypeForCheck && target != r.c.markerSubTypeForCheck {
		constraint := r.c.getBaseConstraintOfType(target)
		switch {
		case constraint != nil && r.c.isTypeAssignableTo(generalizedSource, constraint):
			r.reportError(diagnostics.X_0_is_assignable_to_the_constraint_of_type_1_but_1_could_be_instantiated_with_a_different_subtype_of_constraint_2, generalizedSourceType, targetType, r.c.typeToString(constraint))
		case constraint != nil && r.c.isTypeAssignableTo(source, constraint):
			r.reportError(diagnostics.X_0_is_assignable_to_the_constraint_of_type_1_but_1_could_be_instantiated_with_a_different_subtype_of_constraint_2, sourceType, targetType, r.c.typeToString(constraint))
		default:
			r.reportError(diagnostics.X_0_could_be_instantiated_with_an_arbitrary_type_which_could_be_unrelated_to_1, targetType, generalizedSourceType)
		}
	}
	if message == nil {
		if r.relation == r.c.comparableRelation {
			message = diagnostics.Type_0_is_not_comparable_to_type_1
		} else if sourceType == targetType {
			message = diagnostics.Type_0_is_not_assignable_to_type_1_Two_different_types_with_this_name_exist_but_they_are_unrelated
		} else {
			message = diagnostics.Type_0_is_not_assignable_to_type_1
		}
	}
	switch r.getChainMessage(0) {
	// Suppress if next message is an excess property error
	case diagnostics.Object_literal_may_only_specify_known_properties_and_0_does_not_exist_in_type_1,
		diagnostics.Object_literal_may_only_specify_known_properties_but_0_does_not_exist_in_type_1_Did_you_mean_to_write_2:
		return
	// Suppress if next message is an excessive complexity/stack depth message for source and target or a readonly
	// vs. mutable error for source and target
	case diagnostics.Excessive_complexity_comparing_types_0_and_1,
		diagnostics.Excessive_stack_depth_comparing_types_0_and_1,
		diagnostics.The_type_0_is_readonly_and_cannot_be_assigned_to_the_mutable_type_1:
		if r.chainArgsMatch(sourceType, targetType) {
			return
		}
	// Suppress if next message is a missing property message for source and target and we're not
	// reporting on interface implementation
	case diagnostics.Property_0_is_missing_in_type_1_but_required_in_type_2:
		if !isInterfaceImplementationMessage(message) && r.chainArgsMatch(nil, sourceType, targetType) {
			return
		}
	// Suppress if next message is a missing property message for source and target and we're not
	// reporting on interface implementation
	case diagnostics.Type_0_is_missing_the_following_properties_from_type_1_Colon_2_and_3_more,
		diagnostics.Type_0_is_missing_the_following_properties_from_type_1_Colon_2:
		if !isInterfaceImplementationMessage(message) && r.chainArgsMatch(sourceType, targetType) {
			return
		}
	}
	r.reportError(message, sourceType, targetType)
}

func (r *Relater) reportError(message *diagnostics.Message, args ...any) {
	if message == diagnostics.Types_of_property_0_are_incompatible {
		// Suppress if next message is an excess property error
		switch r.getChainMessage(0) {
		case diagnostics.Object_literal_may_only_specify_known_properties_and_0_does_not_exist_in_type_1,
			diagnostics.Object_literal_may_only_specify_known_properties_but_0_does_not_exist_in_type_1_Did_you_mean_to_write_2:
			return
		}
		// Transform a property incompatibility message for property 'x' followed by some elaboration message
		// followed by a signature return type incompatibility message into a single return type incompatiblity
		// message for 'x()' or 'x(...)'
		var arg string
		switch r.getChainMessage(1) {
		case diagnostics.Call_signatures_with_no_arguments_have_incompatible_return_types_0_and_1:
			arg = getPropertyNameArg(args[0]) + "()"
		case diagnostics.Construct_signatures_with_no_arguments_have_incompatible_return_types_0_and_1:
			arg = "new " + getPropertyNameArg(args[0]) + "()"
		case diagnostics.Call_signature_return_types_0_and_1_are_incompatible:
			arg = getPropertyNameArg(args[0]) + "(...)"
		case diagnostics.Construct_signature_return_types_0_and_1_are_incompatible:
			arg = "new " + getPropertyNameArg(args[0]) + "(...)"
		}
		if arg != "" {
			message = diagnostics.The_types_returned_by_0_are_incompatible_between_these_types
			args[0] = arg
			r.errorChain = r.errorChain.next.next
		}
		// Transform a property incompatibility message for property 'x' followed by some elaboration message
		// followed by a property incompatibility message for property 'y' into a single property incompatibility
		// message for 'x.y'
		switch r.getChainMessage(1) {
		case diagnostics.Types_of_property_0_are_incompatible,
			diagnostics.The_types_of_0_are_incompatible_between_these_types,
			diagnostics.The_types_returned_by_0_are_incompatible_between_these_types:
			head := getPropertyNameArg(args[0])
			tail := getPropertyNameArg(r.errorChain.next.args[0])
			var arg string
			if len(tail) != 0 && tail[0] == '[' {
				arg = head + tail
			} else {
				arg = head + "." + tail
			}
			r.errorChain = r.errorChain.next.next
			if message == diagnostics.Types_of_property_0_are_incompatible {
				message = diagnostics.The_types_of_0_are_incompatible_between_these_types
			}
			r.reportError(message, arg)
			return
		}
	}
	r.errorChain = &ErrorChain{next: r.errorChain, message: message, args: args}
}

func (r *Relater) getChainMessage(index int) *diagnostics.Message {
	e := r.errorChain
	for {
		if e == nil {
			return nil
		}
		if index == 0 {
			return e.message
		}
		e = e.next
		index--
	}
}

// Return true if the arguments of the first entry on the error chain match the
// given arguments (where nil acts as a wildcard).
func (r *Relater) chainArgsMatch(args ...any) bool {
	for i, a := range args {
		if a != nil && a != r.errorChain.args[i] {
			return false
		}
	}
	return true
}

func getPropertyNameArg(arg any) string {
	s := arg.(string)
	if len(s) != 0 && (s[0] == '"' || s[0] == '\'' || s[0] == '`') {
		return "[" + s + "]"
	}
	return s
}

func isInterfaceImplementationMessage(message *diagnostics.Message) bool {
	return message == diagnostics.Class_0_incorrectly_implements_interface_1 ||
		message == diagnostics.Class_0_incorrectly_implements_class_1_Did_you_mean_to_extend_1_and_inherit_its_members_as_a_subclass
}

func chainDepth(chain *ErrorChain) int {
	depth := 0
	for chain != nil {
		depth++
		chain = chain.next
	}
	return depth
}
