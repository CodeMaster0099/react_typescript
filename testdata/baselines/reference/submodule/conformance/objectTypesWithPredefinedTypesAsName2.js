//// [tests/cases/conformance/types/specifyingTypes/predefinedTypes/objectTypesWithPredefinedTypesAsName2.ts] ////

//// [objectTypesWithPredefinedTypesAsName2.ts]
// it is an error to use a predefined type as a type name

class void {} // parse error unlike the others

//// [objectTypesWithPredefinedTypesAsName2.js]
class {
}
void {};
