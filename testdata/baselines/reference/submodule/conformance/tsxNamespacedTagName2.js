//// [tests/cases/conformance/jsx/tsxNamespacedTagName2.tsx] ////

//// [a.tsx]
declare var React: any;

const a = <svg:path></svg:path>;
const b = <svg : path></svg : path>;
const c = <A:foo></A:foo>;
const d = <a:foo></a:foo>;


//// [a.js]
const a = <svg:path></svg:path>;
const b = <svg:path></svg:path>;
const c = <A:foo></A:foo>;
const d = <a:foo></a:foo>;
