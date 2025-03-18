//// [tests/cases/conformance/dynamicImport/importCallExpressionAsyncES6System.ts] ////

//// [test.ts]
export async function fn() {
    const req = await import('./test') // ONE
}

export class cl1 {
    public async m() {
        const req = await import('./test') // TWO
    }
}

export const obj = {
    m: async () => {
        const req = await import('./test') // THREE
    }
}

export class cl2 {
    public p = {
        m: async () => {
            const req = await import('./test') // FOUR
        }
    }
}

export const l = async () => {
    const req = await import('./test') // FIVE
}


//// [test.js]
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.l = exports.cl2 = exports.obj = exports.cl1 = void 0;
exports.fn = fn;
async function fn() {
    const req = await Promise.resolve().then(() => require('./test'));
}
class cl1 {
    async m() {
        const req = await Promise.resolve().then(() => require('./test'));
    }
}
exports.cl1 = cl1;
exports.obj = {
    m: async () => {
        const req = await Promise.resolve().then(() => require('./test'));
    }
};
class cl2 {
    p = {
        m: async () => {
            const req = await Promise.resolve().then(() => require('./test'));
        }
    };
}
exports.cl2 = cl2;
const l = async () => {
    const req = await Promise.resolve().then(() => require('./test'));
};
exports.l = l;
