//// [tests/cases/compiler/unusedTypeParameters6.ts] ////

//// [a.ts]
class C<T> { }

//// [b.ts]
interface C<T> { a: T; }

//// [a.js]
class C {
}
//// [b.js]
