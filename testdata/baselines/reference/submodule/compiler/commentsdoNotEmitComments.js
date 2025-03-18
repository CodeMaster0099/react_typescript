//// [tests/cases/compiler/commentsdoNotEmitComments.ts] ////

//// [commentsdoNotEmitComments.ts]
/** Variable comments*/
var myVariable = 10;

/** function comments*/
function foo(/** parameter comment*/p: number) {
}

/** variable with function type comment*/
var fooVar: () => void;
foo(50);
fooVar();

/**class comment*/
class c {
    /** constructor comment*/
    constructor() {
    }

    /** property comment */
    public b = 10;

    /** function comment */
    public myFoo() {
        return this.b;
    }

    /** getter comment*/
    public get prop1() {
        return this.b;
    }

    /** setter comment*/
    public set prop1(val: number) {
        this.b = val;
    }

    /** overload signature1*/
    public foo1(a: number): string;
    /** Overload signature 2*/
    public foo1(b: string): string;
    /** overload implementation signature*/
    public foo1(aOrb) {
        return aOrb.toString();
    }
}

/**instance comment*/
var i = new c();

/** interface comments*/
interface i1 {
    /** caller comments*/
    (a: number): number;

    /** new comments*/
    new (b: string);

    /**indexer property*/
    [a: number]: string;

    /** function property;*/
    myFoo(/*param prop*/a: number): string;

    /** prop*/
    prop: string;
}

/**interface instance comments*/
var i1_i: i1;

/** this is module comment*/
module m1 {
    /** class b */
    export class b {
        constructor(public x: number) {
           
        }
    }

    /// module m2
    export module m2 {
    }
}

/// this is x
declare var x;


/** const enum member value comment (generated by TS) */
const enum color { red, green, blue }
var shade: color = color.green;


//// [commentsdoNotEmitComments.js]
var myVariable = 10;
function foo(p) {
}
var fooVar;
foo(50);
fooVar();
class c {
    constructor() {
    }
    b = 10;
    myFoo() {
        return this.b;
    }
    get prop1() {
        return this.b;
    }
    set prop1(val) {
        this.b = val;
    }
    foo1(aOrb) {
        return aOrb.toString();
    }
}
var i = new c();
var i1_i;
var m1;
(function (m1) {
    class b {
        x;
        constructor(x) {
            this.x = x;
        }
    }
    m1.b = b;
})(m1 || (m1 = {}));
var color;
(function (color) {
    color[color["red"] = 0] = "red";
    color[color["green"] = 1] = "green";
    color[color["blue"] = 2] = "blue";
})(color || (color = {}));
var shade = color.green;
