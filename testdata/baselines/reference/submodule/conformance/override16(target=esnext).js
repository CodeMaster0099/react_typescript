//// [tests/cases/conformance/override/override16.ts] ////

//// [override16.ts]
class A {
    foo?: string;
}

class B extends A {
    override foo = "string";
}


//// [override16.js]
class A {
    foo;
}
class B extends A {
    foo = "string";
}
