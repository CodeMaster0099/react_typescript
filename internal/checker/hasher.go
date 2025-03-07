package checker

import (
	"slices"
	"strings"

	"github.com/microsoft/typescript-go/internal/ast"
)

type hasher struct {
	strings.Builder
}

var base64chars = []byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V',
	'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
	'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '$', '%',
}

func (b *hasher) WriteInt(value int) {
	for value != 0 {
		b.WriteByte(base64chars[value&0x3F])
		value >>= 6
	}
}

func (b *hasher) WriteSymbol(s *ast.Symbol) {
	b.WriteInt(int(ast.GetSymbolId(s)))
}

func (b *hasher) WriteType(t *Type) {
	b.WriteInt(int(t.id))
}

func (b *hasher) WriteTypes(types []*Type) {
	i := 0
	var tail bool
	for i < len(types) {
		startId := types[i].id
		count := 1
		for i+count < len(types) && types[i+count].id == startId+TypeId(count) {
			count++
		}
		if tail {
			b.WriteByte(',')
		}
		b.WriteInt(int(startId))
		if count > 1 {
			b.WriteByte(':')
			b.WriteInt(count)
		}
		i += count
		tail = true
	}
}

func (b *hasher) WriteAlias(alias *TypeAlias) {
	if alias != nil {
		b.WriteByte('@')
		b.WriteSymbol(alias.symbol)
		if len(alias.typeArguments) != 0 {
			b.WriteByte(':')
			b.WriteTypes(alias.typeArguments)
		}
	}
}

func (b *hasher) WriteGenericTypeReferences(source *Type, target *Type, ignoreConstraints bool) bool {
	var constrained bool
	typeParameters := make([]*Type, 0, 8)
	var writeTypeReference func(*Type, int)
	// writeTypeReference(A<T, number, U>) writes "111=0-12=1"
	// where A.id=111 and number.id=12
	writeTypeReference = func(ref *Type, depth int) {
		b.WriteType(ref.Target())
		for _, t := range ref.AsTypeReference().resolvedTypeArguments {
			if t.flags&TypeFlagsTypeParameter != 0 {
				if ignoreConstraints || t.checker.getConstraintOfTypeParameter(t) == nil {
					index := slices.Index(typeParameters, t)
					if index < 0 {
						index = len(typeParameters)
						typeParameters = append(typeParameters, t)
					}
					b.WriteByte('=')
					b.WriteInt(index)
					continue
				}
				constrained = true
			} else if depth < 4 && isTypeReferenceWithGenericArguments(t) {
				b.WriteByte('<')
				writeTypeReference(t, depth+1)
				b.WriteByte('>')
				continue
			}
			b.WriteByte('-')
			b.WriteType(t)
		}
	}
	writeTypeReference(source, 0)
	b.WriteByte(',')
	writeTypeReference(target, 0)
	return constrained
}

func (b *hasher) WriteNode(node *ast.Node) {
	if node != nil {
		b.WriteInt(int(ast.GetNodeId(node)))
	}
}

func getTypeListKey(types []*Type) string {
	var b hasher
	b.WriteTypes(types)
	return b.String()
}

func getAliasKey(alias *TypeAlias) string {
	var b hasher
	b.WriteAlias(alias)
	return b.String()
}

func getUnionKey(types []*Type, origin *Type, alias *TypeAlias) string {
	var b hasher
	switch {
	case origin == nil:
		b.WriteTypes(types)
	case origin.flags&TypeFlagsUnion != 0:
		b.WriteByte('|')
		b.WriteTypes(origin.Types())
	case origin.flags&TypeFlagsIntersection != 0:
		b.WriteByte('&')
		b.WriteTypes(origin.Types())
	case origin.flags&TypeFlagsIndex != 0:
		// origin type id alone is insufficient, as `keyof x` may resolve to multiple WIP values while `x` is still resolving
		b.WriteByte('#')
		b.WriteType(origin)
		b.WriteByte('|')
		b.WriteTypes(types)
	default:
		panic("Unhandled case in getUnionKey")
	}
	b.WriteAlias(alias)
	return b.String()
}

func getIntersectionKey(types []*Type, flags IntersectionFlags, alias *TypeAlias) string {
	var b hasher
	b.WriteTypes(types)
	if flags&IntersectionFlagsNoConstraintReduction == 0 {
		b.WriteAlias(alias)
	} else {
		b.WriteByte('*')
	}
	return b.String()
}

func getTupleKey(elementInfos []TupleElementInfo, readonly bool) string {
	var b hasher
	for _, e := range elementInfos {
		switch {
		case e.flags&ElementFlagsRequired != 0:
			b.WriteByte('#')
		case e.flags&ElementFlagsOptional != 0:
			b.WriteByte('?')
		case e.flags&ElementFlagsRest != 0:
			b.WriteByte('.')
		default:
			b.WriteByte('*')
		}
		if e.labeledDeclaration != nil {
			b.WriteInt(int(ast.GetNodeId(e.labeledDeclaration)))
		}
	}
	if readonly {
		b.WriteByte('!')
	}
	return b.String()
}

func getTypeAliasInstantiationKey(typeArguments []*Type, alias *TypeAlias) string {
	return getTypeInstantiationKey(typeArguments, alias, false)
}

func getTypeInstantiationKey(typeArguments []*Type, alias *TypeAlias, singleSignature bool) string {
	var b hasher
	b.WriteTypes(typeArguments)
	b.WriteAlias(alias)
	if singleSignature {
		b.WriteByte('!')
	}
	return b.String()
}

func getIndexedAccessKey(objectType *Type, indexType *Type, accessFlags AccessFlags, alias *TypeAlias) string {
	var b hasher
	b.WriteType(objectType)
	b.WriteByte(',')
	b.WriteType(indexType)
	b.WriteByte(',')
	b.WriteInt(int(accessFlags))
	b.WriteAlias(alias)
	return b.String()
}

func getTemplateTypeKey(texts []string, types []*Type) string {
	var b hasher
	b.WriteTypes(types)
	b.WriteByte('|')
	for i, s := range texts {
		if i != 0 {
			b.WriteByte(',')
		}
		b.WriteInt(len(s))
	}
	b.WriteByte('|')
	for _, s := range texts {
		b.WriteString(s)
	}
	return b.String()
}

func getConditionalTypeKey(typeArguments []*Type, alias *TypeAlias, forConstraint bool) string {
	var b hasher
	b.WriteTypes(typeArguments)
	b.WriteAlias(alias)
	if forConstraint {
		b.WriteByte('!')
	}
	return b.String()
}

func getRelationKey(source *Type, target *Type, intersectionState IntersectionState, isIdentity bool, ignoreConstraints bool) string {
	if isIdentity && source.id > target.id {
		source, target = target, source
	}
	var b hasher
	var constrained bool
	if isTypeReferenceWithGenericArguments(source) && isTypeReferenceWithGenericArguments(target) {
		constrained = b.WriteGenericTypeReferences(source, target, ignoreConstraints)
	} else {
		b.WriteType(source)
		b.WriteByte(',')
		b.WriteType(target)
	}
	if intersectionState != IntersectionStateNone {
		b.WriteByte(':')
		b.WriteInt(int(intersectionState))
	}
	if constrained {
		// We mark keys with type references that reference constrained type parameters such that we know
		// to obtain and look for a "broadest equivalent key" in the cache.
		b.WriteByte('*')
	}
	return b.String()
}

func getNodeListKey(nodes []*ast.Node) string {
	var b hasher
	for i, n := range nodes {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteNode(n)
	}
	return b.String()
}
