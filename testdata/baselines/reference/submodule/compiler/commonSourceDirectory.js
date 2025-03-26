//// [tests/cases/compiler/commonSourceDirectory.ts] ////

//// [index.ts]
export const x = 0;

//// [bar.d.ts]
declare module "bar" {
    export const y = 0;
}

//// [index.ts]
/// <reference path="../types/bar.d.ts" preserve="true" />
import { x } from "foo";
import { y } from "bar";
x + y;


//// [/app/bin/node_modules/foo/index.js]
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.x = void 0;
exports.x = 0;
//// [/app/bin/app/index.js]
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
/// <reference path="../types/bar.d.ts" preserve="true" />
const foo_1 = require("foo");
const bar_1 = require("bar");
foo_1.x + bar_1.y;
