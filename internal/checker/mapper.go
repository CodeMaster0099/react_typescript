package checker

import "github.com/microsoft/typescript-go/internal/core"

// TypeMapperKind

type TypeMapperKind int32

const (
	TypeMapperKindUnknown TypeMapperKind = iota
	TypeMapperKindSimple
	TypeMapperKindArray
	TypeMapperKindMerged
)

// TypeMapper

type TypeMapper struct {
	data TypeMapperData
}

func (m *TypeMapper) Map(t *Type) *Type    { return m.data.Map(t) }
func (m *TypeMapper) Kind() TypeMapperKind { return m.data.Kind() }

// TypeMapperData

type TypeMapperData interface {
	Map(t *Type) *Type
	Kind() TypeMapperKind
}

// Factory functions

func (c *Checker) newTypeMapper(sources []*Type, targets []*Type) *TypeMapper {
	if len(sources) == 1 {
		return c.typeMapperFactory.newSimpleTypeMapper(sources[0], targets[0])
	}
	return c.typeMapperFactory.newArrayTypeMapper(sources, targets)
}

func (c *Checker) combineTypeMappers(m1 *TypeMapper, m2 *TypeMapper) *TypeMapper {
	if m1 != nil {
		return c.typeMapperFactory.newCompositeTypeMapper(c, m1, m2)
	}
	return m2
}

func (c *Checker) mergeTypeMappers(m1 *TypeMapper, m2 *TypeMapper) *TypeMapper {
	if m1 != nil {
		return c.typeMapperFactory.newMergedTypeMapper(m1, m2)
	}
	return m2
}

func (c *Checker) prependTypeMapping(source *Type, target *Type, mapper *TypeMapper) *TypeMapper {
	if mapper == nil {
		return c.typeMapperFactory.newSimpleTypeMapper(source, target)
	}
	return c.typeMapperFactory.newMergedTypeMapper(c.typeMapperFactory.newSimpleTypeMapper(source, target), mapper)
}

func (c *Checker) appendTypeMapping(mapper *TypeMapper, source *Type, target *Type) *TypeMapper {
	if mapper == nil {
		return c.typeMapperFactory.newSimpleTypeMapper(source, target)
	}
	return c.typeMapperFactory.newMergedTypeMapper(mapper, c.typeMapperFactory.newSimpleTypeMapper(source, target))
}

// Maps forward-references to later types parameters to the empty object type.
// This is used during inference when instantiating type parameter defaults.
func (c *Checker) newBackreferenceMapper(context *InferenceContext, index int) *TypeMapper {
	forwardInferences := context.inferences[index:]
	typeParameters := core.Map(forwardInferences, func(i *InferenceInfo) *Type {
		return i.typeParameter
	})
	return c.typeMapperFactory.newArrayToSingleTypeMapper(typeParameters, c.unknownType)
}

// TypeMapperBase

type TypeMapperBase struct {
	TypeMapper
}

func (m *TypeMapperBase) Map(t *Type) *Type    { return t }
func (m *TypeMapperBase) Kind() TypeMapperKind { return TypeMapperKindUnknown }

type TypeMapperFactory struct {
	simpleTypeMapperPool        core.Pool[SimpleTypeMapper]
	arrayTypeMapperPool         core.Pool[ArrayTypeMapper]
	arrayToSingleTypeMapperPool core.Pool[ArrayToSingleTypeMapper]
	deferredTypeMapperPool      core.Pool[DeferredTypeMapper]
	functionTypeMapperPool      core.Pool[FunctionTypeMapper]
	mergedTypeMapperPool        core.Pool[MergedTypeMapper]
	compositeTypeMapperPool     core.Pool[CompositeTypeMapper]
	inferenceTypeMapperPool     core.Pool[InferenceTypeMapper]
}

// SimpleTypeMapper

type SimpleTypeMapper struct {
	TypeMapperBase
	source *Type
	target *Type
}

func (f *TypeMapperFactory) newSimpleTypeMapper(source *Type, target *Type) *TypeMapper {
	m := f.simpleTypeMapperPool.New()
	m.data = m
	m.source = source
	m.target = target
	return &m.TypeMapper
}

func (m *SimpleTypeMapper) Map(t *Type) *Type {
	if t == m.source {
		return m.target
	}
	return t
}

func (m *SimpleTypeMapper) Kind() TypeMapperKind {
	return TypeMapperKindSimple
}

// ArrayTypeMapper

type ArrayTypeMapper struct {
	TypeMapperBase
	sources []*Type
	targets []*Type
}

func (f *TypeMapperFactory) newArrayTypeMapper(sources []*Type, targets []*Type) *TypeMapper {
	m := f.arrayTypeMapperPool.New()
	m.data = m
	m.sources = sources
	m.targets = targets
	return &m.TypeMapper
}

func (m *ArrayTypeMapper) Map(t *Type) *Type {
	for i, s := range m.sources {
		if t == s {
			return m.targets[i]
		}
	}
	return t
}

func (m *ArrayTypeMapper) Kind() TypeMapperKind {
	return TypeMapperKindArray
}

// ArrayToSingleTypeMapper

type ArrayToSingleTypeMapper struct {
	TypeMapperBase
	sources []*Type
	target  *Type
}

func (f *TypeMapperFactory) newArrayToSingleTypeMapper(sources []*Type, target *Type) *TypeMapper {
	m := f.arrayToSingleTypeMapperPool.New()
	m.data = m
	m.sources = sources
	m.target = target
	return &m.TypeMapper
}

func (m *ArrayToSingleTypeMapper) Map(t *Type) *Type {
	for _, s := range m.sources {
		if t == s {
			return m.target
		}
	}
	return t
}

// DeferredTypeMapper

type DeferredTypeMapper struct {
	TypeMapperBase
	sources []*Type
	targets []func() *Type
}

func (f *TypeMapperFactory) newDeferredTypeMapper(sources []*Type, targets []func() *Type) *TypeMapper {
	m := f.deferredTypeMapperPool.New()
	m.data = m
	m.sources = sources
	m.targets = targets
	return &m.TypeMapper
}

func (m *DeferredTypeMapper) Map(t *Type) *Type {
	for i, s := range m.sources {
		if t == s {
			return m.targets[i]()
		}
	}
	return t
}

// FunctionTypeMapper

type FunctionTypeMapper struct {
	TypeMapperBase
	fn func(*Type) *Type
}

func (f *TypeMapperFactory) newFunctionTypeMapper(fn func(*Type) *Type) *TypeMapper {
	m := f.functionTypeMapperPool.New()
	m.data = m
	m.fn = fn
	return &m.TypeMapper
}

func (m *FunctionTypeMapper) Map(t *Type) *Type {
	return m.fn(t)
}

// MergedTypeMapper

type MergedTypeMapper struct {
	TypeMapperBase
	m1 *TypeMapper
	m2 *TypeMapper
}

func (f *TypeMapperFactory) newMergedTypeMapper(m1 *TypeMapper, m2 *TypeMapper) *TypeMapper {
	m := f.mergedTypeMapperPool.New()
	m.data = m
	m.m1 = m1
	m.m2 = m2
	return &m.TypeMapper
}

func (m *MergedTypeMapper) Map(t *Type) *Type {
	return m.m2.Map(m.m1.Map(t))
}

func (m *MergedTypeMapper) Kind() TypeMapperKind {
	return TypeMapperKindMerged
}

// CompositeTypeMapper

type CompositeTypeMapper struct {
	TypeMapperBase
	c  *Checker
	m1 *TypeMapper
	m2 *TypeMapper
}

func (f *TypeMapperFactory) newCompositeTypeMapper(c *Checker, m1 *TypeMapper, m2 *TypeMapper) *TypeMapper {
	m := f.compositeTypeMapperPool.New()
	m.data = m
	m.c = c
	m.m1 = m1
	m.m2 = m2
	return &m.TypeMapper
}

func (m *CompositeTypeMapper) Map(t *Type) *Type {
	t1 := m.m1.Map(t)
	if t1 != t {
		return m.c.instantiateType(t1, m.m2)
	}
	return m.m2.Map(t)
}

// InferenceTypeMapper

type InferenceTypeMapper struct {
	TypeMapperBase
	c      *Checker
	n      *InferenceContext
	fixing bool
}

func (f *TypeMapperFactory) newInferenceTypeMapper(c *Checker, n *InferenceContext, fixing bool) *TypeMapper {
	m := f.inferenceTypeMapperPool.New()
	m.data = m
	m.c = c
	m.n = n
	m.fixing = fixing
	return &m.TypeMapper
}

func (m *InferenceTypeMapper) Map(t *Type) *Type {
	for i, inference := range m.n.inferences {
		if t == inference.typeParameter {
			if m.fixing && !inference.isFixed {
				// Before we commit to a particular inference (and thus lock out any further inferences),
				// we infer from any intra-expression inference sites we have collected.
				m.c.inferFromIntraExpressionSites(m.n)
				clearCachedInferences(m.n.inferences)
				inference.isFixed = true
			}
			return m.c.getInferredType(m.n, i)
		}
	}
	return t
}
