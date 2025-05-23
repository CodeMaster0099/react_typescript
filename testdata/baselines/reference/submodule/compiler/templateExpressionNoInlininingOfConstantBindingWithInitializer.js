//// [tests/cases/compiler/templateExpressionNoInlininingOfConstantBindingWithInitializer.ts] ////

//// [templateExpressionNoInlininingOfConstantBindingWithInitializer.ts]
type Params = {
  value?: string | number
}

function example(parameters: Params) {
  const { value = '123' } = parameters
  return `${value}` === '345'
}

function example2(parameters: Params) {
  const { value = '123' } = parameters
  const b = `${value}`;
  return b;
}


//// [templateExpressionNoInlininingOfConstantBindingWithInitializer.js]
function example(parameters) {
    const { value = '123' } = parameters;
    return `${value}` === '345';
}
function example2(parameters) {
    const { value = '123' } = parameters;
    const b = `${value}`;
    return b;
}
