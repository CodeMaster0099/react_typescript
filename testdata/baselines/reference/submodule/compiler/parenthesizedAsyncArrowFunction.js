//// [tests/cases/compiler/parenthesizedAsyncArrowFunction.ts] ////

//// [parenthesizedAsyncArrowFunction.ts]
// Repro from #20096

let foo = (async bar => bar);


//// [parenthesizedAsyncArrowFunction.js]
// Repro from #20096
let foo = (async bar => bar);
