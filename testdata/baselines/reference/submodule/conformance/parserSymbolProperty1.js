//// [tests/cases/conformance/parser/ecmascript6/Symbols/parserSymbolProperty1.ts] ////

//// [parserSymbolProperty1.ts]
interface I {
    [Symbol.iterator]: string;
}

//// [parserSymbolProperty1.js]
