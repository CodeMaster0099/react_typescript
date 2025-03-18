//// [tests/cases/compiler/validUseOfThisInSuper.ts] ////

//// [validUseOfThisInSuper.ts]
class Base {
    constructor(public b: Base) {
    }
}
class Super extends Base {
    constructor() {
        super((() => this)()); // ok since this is not the case: The constructor declares parameter properties or the containing class declares instance member variables with initializers.
    }
}

//// [validUseOfThisInSuper.js]
class Base {
    b;
    constructor(b) {
        this.b = b;
    }
}
class Super extends Base {
    constructor() {
        super((() => this)());
    }
}
