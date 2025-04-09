//// [tests/cases/conformance/classes/propertyMemberDeclarations/staticPropertyNameConflicts.ts] ////

//// [staticPropertyNameConflicts.ts]
const FunctionPropertyNames = {
    name: 'name',
    length: 'length',
    prototype: 'prototype',
    caller: 'caller',
    arguments: 'arguments',
} as const;

// name
class StaticName {
    static name: number; // error without useDefineForClassFields
    name: string; // ok
}

class StaticName2 {
    static [FunctionPropertyNames.name]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.name]: number; // ok
}

class StaticNameFn {
    static name() {} // error without useDefineForClassFields
    name() {} // ok
}

class StaticNameFn2 {
    static [FunctionPropertyNames.name]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.name]() {} // ok
}

// length
class StaticLength {
    static length: number; // error without useDefineForClassFields
    length: string; // ok
}

class StaticLength2 {
    static [FunctionPropertyNames.length]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.length]: number; // ok
}

class StaticLengthFn {
    static length() {} // error without useDefineForClassFields
    length() {} // ok
}

class StaticLengthFn2 {
    static [FunctionPropertyNames.length]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.length]() {} // ok
}

// prototype
class StaticPrototype {
    static prototype: number; // always an error
    prototype: string; // ok
}

class StaticPrototype2 {
    static [FunctionPropertyNames.prototype]: number; // always an error
    [FunctionPropertyNames.prototype]: string; // ok
}

class StaticPrototypeFn {
    static prototype() {} // always an error
    prototype() {} // ok
}

class StaticPrototypeFn2 {
    static [FunctionPropertyNames.prototype]() {} // always an error
    [FunctionPropertyNames.prototype]() {} // ok
}

// caller
class StaticCaller {
    static caller: number; // error without useDefineForClassFields
    caller: string; // ok
}

class StaticCaller2 {
    static [FunctionPropertyNames.caller]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.caller]: string; // ok
}

class StaticCallerFn {
    static caller() {} // error without useDefineForClassFields
    caller() {} // ok
}

class StaticCallerFn2 {
    static [FunctionPropertyNames.caller]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.caller]() {} // ok
}

// arguments
class StaticArguments {
    static arguments: number; // error without useDefineForClassFields
    arguments: string; // ok
}

class StaticArguments2 {
    static [FunctionPropertyNames.arguments]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]: string; // ok
}

class StaticArgumentsFn {
    static arguments() {} // error without useDefineForClassFields
    arguments() {} // ok
}

class StaticArgumentsFn2 {
    static [FunctionPropertyNames.arguments]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]() {} // ok
}


// === Static properties on anonymous classes ===

// name
var StaticName_Anonymous = class {
    static name: number; // error without useDefineForClassFields
    name: string; // ok
}

var StaticName_Anonymous2 = class {
    static [FunctionPropertyNames.name]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.name]: string; // ok
}

var StaticNameFn_Anonymous = class {
    static name() {} // error without useDefineForClassFields
    name() {} // ok
}

var StaticNameFn_Anonymous2 = class {
    static [FunctionPropertyNames.name]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.name]() {} // ok
}

// length
var StaticLength_Anonymous = class {
    static length: number; // error without useDefineForClassFields
    length: string; // ok
}

var StaticLength_Anonymous2 = class {
    static [FunctionPropertyNames.length]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.length]: string; // ok
}

var StaticLengthFn_Anonymous = class {
    static length() {} // error without useDefineForClassFields
    length() {} // ok
}

var StaticLengthFn_Anonymous2 = class {
    static [FunctionPropertyNames.length]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.length]() {} // ok
}

// prototype
var StaticPrototype_Anonymous = class {
    static prototype: number; // always an error
    prototype: string; // ok
}

var StaticPrototype_Anonymous2 = class {
    static [FunctionPropertyNames.prototype]: number; // always an error
    [FunctionPropertyNames.prototype]: string; // ok
}

var StaticPrototypeFn_Anonymous = class {
    static prototype() {} // always an error
    prototype() {} // ok
}

var StaticPrototypeFn_Anonymous2 = class {
    static [FunctionPropertyNames.prototype]() {} // always an error
    [FunctionPropertyNames.prototype]() {} // ok
}

// caller
var StaticCaller_Anonymous = class {
    static caller: number; // error without useDefineForClassFields
    caller: string; // ok
}

var StaticCaller_Anonymous2 = class {
    static [FunctionPropertyNames.caller]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.caller]: string; // ok
}

var StaticCallerFn_Anonymous = class {
    static caller() {} // error without useDefineForClassFields
    caller() {} // ok
}

var StaticCallerFn_Anonymous2 = class {
    static [FunctionPropertyNames.caller]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.caller]() {} // ok
}

// arguments
var StaticArguments_Anonymous = class {
    static arguments: number; // error without useDefineForClassFields
    arguments: string; // ok
}

var StaticArguments_Anonymous2 = class {
    static [FunctionPropertyNames.arguments]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]: string; // ok
}

var StaticArgumentsFn_Anonymous = class {
    static arguments() {} // error without useDefineForClassFields
    arguments() {} // ok
}

var StaticArgumentsFn_Anonymous2 = class {
    static [FunctionPropertyNames.arguments]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]() {} // ok
}


// === Static properties on default exported classes ===

// name
module TestOnDefaultExportedClass_1 {
    class StaticName {
        static name: number; // error without useDefineForClassFields
        name: string; // ok
    }
}

export class ExportedStaticName {
    static [FunctionPropertyNames.name]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.name]: string; // ok
}

module TestOnDefaultExportedClass_2 {
    class StaticNameFn {
        static name() {} // error without useDefineForClassFields
        name() {} // ok
    }
}

export class ExportedStaticNameFn {
    static [FunctionPropertyNames.name]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.name]() {} // ok
}

// length
module TestOnDefaultExportedClass_3 {
    export default class StaticLength {
        static length: number; // error without useDefineForClassFields
        length: string; // ok
    }
}

export class ExportedStaticLength {
    static [FunctionPropertyNames.length]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.length]: string; // ok
}

module TestOnDefaultExportedClass_4 {
    export default class StaticLengthFn {
        static length() {} // error without useDefineForClassFields
        length() {} // ok
    }
}

export class ExportedStaticLengthFn {
    static [FunctionPropertyNames.length]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.length]() {} // ok
}

// prototype
module TestOnDefaultExportedClass_5 {
    export default class StaticPrototype {
        static prototype: number; // always an error
        prototype: string; // ok
    }
}

export class ExportedStaticPrototype {
    static [FunctionPropertyNames.prototype]: number; // always an error
    [FunctionPropertyNames.prototype]: string; // ok
}

module TestOnDefaultExportedClass_6 {
    export default class StaticPrototypeFn {
        static prototype() {} // always an error
        prototype() {} // ok
    }
}

export class ExportedStaticPrototypeFn {
    static [FunctionPropertyNames.prototype]() {} // always an error
    [FunctionPropertyNames.prototype]() {} // ok
}

// caller
module TestOnDefaultExportedClass_7 {
    export default class StaticCaller {
        static caller: number; // error without useDefineForClassFields
        caller: string; // ok
    }
}

export class ExportedStaticCaller {
    static [FunctionPropertyNames.caller]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.caller]: string; // ok
}

module TestOnDefaultExportedClass_8 {
    export default class StaticCallerFn {
        static caller() {} // error without useDefineForClassFields
        caller() {} // ok
    }
}

export class ExportedStaticCallerFn {
    static [FunctionPropertyNames.caller]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.caller]() {} // ok
}

// arguments
module TestOnDefaultExportedClass_9 {
    export default class StaticArguments {
        static arguments: number; // error without useDefineForClassFields
        arguments: string; // ok
    }
}

export class ExportedStaticArguments {
    static [FunctionPropertyNames.arguments]: number; // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]: string; // ok
}

module TestOnDefaultExportedClass_10 {
    export default class StaticArgumentsFn {
        static arguments() {} // error without useDefineForClassFields
        arguments() {} // ok
    }
}

export class ExportedStaticArgumentsFn {
    static [FunctionPropertyNames.arguments]() {} // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]() {} // ok
}

//// [staticPropertyNameConflicts.js]
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExportedStaticArgumentsFn = exports.ExportedStaticArguments = exports.ExportedStaticCallerFn = exports.ExportedStaticCaller = exports.ExportedStaticPrototypeFn = exports.ExportedStaticPrototype = exports.ExportedStaticLengthFn = exports.ExportedStaticLength = exports.ExportedStaticNameFn = exports.ExportedStaticName = void 0;
const FunctionPropertyNames = {
    name: 'name',
    length: 'length',
    prototype: 'prototype',
    caller: 'caller',
    arguments: 'arguments',
};
// name
class StaticName {
    static name; // error without useDefineForClassFields
    name; // ok
}
class StaticName2 {
    static [FunctionPropertyNames.name]; // error without useDefineForClassFields
    [FunctionPropertyNames.name]; // ok
}
class StaticNameFn {
    static name() { } // error without useDefineForClassFields
    name() { } // ok
}
class StaticNameFn2 {
    static [FunctionPropertyNames.name]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.name]() { } // ok
}
// length
class StaticLength {
    static length; // error without useDefineForClassFields
    length; // ok
}
class StaticLength2 {
    static [FunctionPropertyNames.length]; // error without useDefineForClassFields
    [FunctionPropertyNames.length]; // ok
}
class StaticLengthFn {
    static length() { } // error without useDefineForClassFields
    length() { } // ok
}
class StaticLengthFn2 {
    static [FunctionPropertyNames.length]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.length]() { } // ok
}
// prototype
class StaticPrototype {
    static prototype; // always an error
    prototype; // ok
}
class StaticPrototype2 {
    static [FunctionPropertyNames.prototype]; // always an error
    [FunctionPropertyNames.prototype]; // ok
}
class StaticPrototypeFn {
    static prototype() { } // always an error
    prototype() { } // ok
}
class StaticPrototypeFn2 {
    static [FunctionPropertyNames.prototype]() { } // always an error
    [FunctionPropertyNames.prototype]() { } // ok
}
// caller
class StaticCaller {
    static caller; // error without useDefineForClassFields
    caller; // ok
}
class StaticCaller2 {
    static [FunctionPropertyNames.caller]; // error without useDefineForClassFields
    [FunctionPropertyNames.caller]; // ok
}
class StaticCallerFn {
    static caller() { } // error without useDefineForClassFields
    caller() { } // ok
}
class StaticCallerFn2 {
    static [FunctionPropertyNames.caller]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.caller]() { } // ok
}
// arguments
class StaticArguments {
    static arguments; // error without useDefineForClassFields
    arguments; // ok
}
class StaticArguments2 {
    static [FunctionPropertyNames.arguments]; // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]; // ok
}
class StaticArgumentsFn {
    static arguments() { } // error without useDefineForClassFields
    arguments() { } // ok
}
class StaticArgumentsFn2 {
    static [FunctionPropertyNames.arguments]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]() { } // ok
}
// === Static properties on anonymous classes ===
// name
var StaticName_Anonymous = class {
    static name; // error without useDefineForClassFields
    name; // ok
};
var StaticName_Anonymous2 = class {
    static [FunctionPropertyNames.name]; // error without useDefineForClassFields
    [FunctionPropertyNames.name]; // ok
};
var StaticNameFn_Anonymous = class {
    static name() { } // error without useDefineForClassFields
    name() { } // ok
};
var StaticNameFn_Anonymous2 = class {
    static [FunctionPropertyNames.name]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.name]() { } // ok
};
// length
var StaticLength_Anonymous = class {
    static length; // error without useDefineForClassFields
    length; // ok
};
var StaticLength_Anonymous2 = class {
    static [FunctionPropertyNames.length]; // error without useDefineForClassFields
    [FunctionPropertyNames.length]; // ok
};
var StaticLengthFn_Anonymous = class {
    static length() { } // error without useDefineForClassFields
    length() { } // ok
};
var StaticLengthFn_Anonymous2 = class {
    static [FunctionPropertyNames.length]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.length]() { } // ok
};
// prototype
var StaticPrototype_Anonymous = class {
    static prototype; // always an error
    prototype; // ok
};
var StaticPrototype_Anonymous2 = class {
    static [FunctionPropertyNames.prototype]; // always an error
    [FunctionPropertyNames.prototype]; // ok
};
var StaticPrototypeFn_Anonymous = class {
    static prototype() { } // always an error
    prototype() { } // ok
};
var StaticPrototypeFn_Anonymous2 = class {
    static [FunctionPropertyNames.prototype]() { } // always an error
    [FunctionPropertyNames.prototype]() { } // ok
};
// caller
var StaticCaller_Anonymous = class {
    static caller; // error without useDefineForClassFields
    caller; // ok
};
var StaticCaller_Anonymous2 = class {
    static [FunctionPropertyNames.caller]; // error without useDefineForClassFields
    [FunctionPropertyNames.caller]; // ok
};
var StaticCallerFn_Anonymous = class {
    static caller() { } // error without useDefineForClassFields
    caller() { } // ok
};
var StaticCallerFn_Anonymous2 = class {
    static [FunctionPropertyNames.caller]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.caller]() { } // ok
};
// arguments
var StaticArguments_Anonymous = class {
    static arguments; // error without useDefineForClassFields
    arguments; // ok
};
var StaticArguments_Anonymous2 = class {
    static [FunctionPropertyNames.arguments]; // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]; // ok
};
var StaticArgumentsFn_Anonymous = class {
    static arguments() { } // error without useDefineForClassFields
    arguments() { } // ok
};
var StaticArgumentsFn_Anonymous2 = class {
    static [FunctionPropertyNames.arguments]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]() { } // ok
};
// === Static properties on default exported classes ===
// name
var TestOnDefaultExportedClass_1;
(function (TestOnDefaultExportedClass_1) {
    class StaticName {
        static name; // error without useDefineForClassFields
        name; // ok
    }
})(TestOnDefaultExportedClass_1 || (TestOnDefaultExportedClass_1 = {}));
class ExportedStaticName {
    static [FunctionPropertyNames.name]; // error without useDefineForClassFields
    [FunctionPropertyNames.name]; // ok
}
exports.ExportedStaticName = ExportedStaticName;
var TestOnDefaultExportedClass_2;
(function (TestOnDefaultExportedClass_2) {
    class StaticNameFn {
        static name() { } // error without useDefineForClassFields
        name() { } // ok
    }
})(TestOnDefaultExportedClass_2 || (TestOnDefaultExportedClass_2 = {}));
class ExportedStaticNameFn {
    static [FunctionPropertyNames.name]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.name]() { } // ok
}
exports.ExportedStaticNameFn = ExportedStaticNameFn;
// length
var TestOnDefaultExportedClass_3;
(function (TestOnDefaultExportedClass_3) {
    class StaticLength {
        static length; // error without useDefineForClassFields
        length; // ok
    }
    TestOnDefaultExportedClass_3.StaticLength = StaticLength;
})(TestOnDefaultExportedClass_3 || (TestOnDefaultExportedClass_3 = {}));
class ExportedStaticLength {
    static [FunctionPropertyNames.length]; // error without useDefineForClassFields
    [FunctionPropertyNames.length]; // ok
}
exports.ExportedStaticLength = ExportedStaticLength;
var TestOnDefaultExportedClass_4;
(function (TestOnDefaultExportedClass_4) {
    class StaticLengthFn {
        static length() { } // error without useDefineForClassFields
        length() { } // ok
    }
    TestOnDefaultExportedClass_4.StaticLengthFn = StaticLengthFn;
})(TestOnDefaultExportedClass_4 || (TestOnDefaultExportedClass_4 = {}));
class ExportedStaticLengthFn {
    static [FunctionPropertyNames.length]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.length]() { } // ok
}
exports.ExportedStaticLengthFn = ExportedStaticLengthFn;
// prototype
var TestOnDefaultExportedClass_5;
(function (TestOnDefaultExportedClass_5) {
    class StaticPrototype {
        static prototype; // always an error
        prototype; // ok
    }
    TestOnDefaultExportedClass_5.StaticPrototype = StaticPrototype;
})(TestOnDefaultExportedClass_5 || (TestOnDefaultExportedClass_5 = {}));
class ExportedStaticPrototype {
    static [FunctionPropertyNames.prototype]; // always an error
    [FunctionPropertyNames.prototype]; // ok
}
exports.ExportedStaticPrototype = ExportedStaticPrototype;
var TestOnDefaultExportedClass_6;
(function (TestOnDefaultExportedClass_6) {
    class StaticPrototypeFn {
        static prototype() { } // always an error
        prototype() { } // ok
    }
    TestOnDefaultExportedClass_6.StaticPrototypeFn = StaticPrototypeFn;
})(TestOnDefaultExportedClass_6 || (TestOnDefaultExportedClass_6 = {}));
class ExportedStaticPrototypeFn {
    static [FunctionPropertyNames.prototype]() { } // always an error
    [FunctionPropertyNames.prototype]() { } // ok
}
exports.ExportedStaticPrototypeFn = ExportedStaticPrototypeFn;
// caller
var TestOnDefaultExportedClass_7;
(function (TestOnDefaultExportedClass_7) {
    class StaticCaller {
        static caller; // error without useDefineForClassFields
        caller; // ok
    }
    TestOnDefaultExportedClass_7.StaticCaller = StaticCaller;
})(TestOnDefaultExportedClass_7 || (TestOnDefaultExportedClass_7 = {}));
class ExportedStaticCaller {
    static [FunctionPropertyNames.caller]; // error without useDefineForClassFields
    [FunctionPropertyNames.caller]; // ok
}
exports.ExportedStaticCaller = ExportedStaticCaller;
var TestOnDefaultExportedClass_8;
(function (TestOnDefaultExportedClass_8) {
    class StaticCallerFn {
        static caller() { } // error without useDefineForClassFields
        caller() { } // ok
    }
    TestOnDefaultExportedClass_8.StaticCallerFn = StaticCallerFn;
})(TestOnDefaultExportedClass_8 || (TestOnDefaultExportedClass_8 = {}));
class ExportedStaticCallerFn {
    static [FunctionPropertyNames.caller]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.caller]() { } // ok
}
exports.ExportedStaticCallerFn = ExportedStaticCallerFn;
// arguments
var TestOnDefaultExportedClass_9;
(function (TestOnDefaultExportedClass_9) {
    class StaticArguments {
        static arguments; // error without useDefineForClassFields
        arguments; // ok
    }
    TestOnDefaultExportedClass_9.StaticArguments = StaticArguments;
})(TestOnDefaultExportedClass_9 || (TestOnDefaultExportedClass_9 = {}));
class ExportedStaticArguments {
    static [FunctionPropertyNames.arguments]; // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]; // ok
}
exports.ExportedStaticArguments = ExportedStaticArguments;
var TestOnDefaultExportedClass_10;
(function (TestOnDefaultExportedClass_10) {
    class StaticArgumentsFn {
        static arguments() { } // error without useDefineForClassFields
        arguments() { } // ok
    }
    TestOnDefaultExportedClass_10.StaticArgumentsFn = StaticArgumentsFn;
})(TestOnDefaultExportedClass_10 || (TestOnDefaultExportedClass_10 = {}));
class ExportedStaticArgumentsFn {
    static [FunctionPropertyNames.arguments]() { } // error without useDefineForClassFields
    [FunctionPropertyNames.arguments]() { } // ok
}
exports.ExportedStaticArgumentsFn = ExportedStaticArgumentsFn;
