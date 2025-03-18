//// [tests/cases/conformance/jsx/tsxReactEmitWhitespace2.tsx] ////

//// [file.tsx]
declare module JSX {
	interface Element { }
	interface IntrinsicElements {
		[s: string]: any;
	}
}
declare var React: any;

// Emit ' word' in the last string
<div>word <code>code</code> word</div>;
// Same here
<div><code>code</code> word</div>;
// And here
<div><code /> word</div>;



//// [file.js]
<div>word <code>code</code> word</div>;
<div><code>code</code> word</div>;
<div><code /> word</div>;
