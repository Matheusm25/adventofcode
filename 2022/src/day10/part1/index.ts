import { readFileSync } from 'fs';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input).split('\n');

const cyclesToCheck = [20, 60, 100, 140, 180, 220];
const cycleValues: Array<number> = [];

let x = 1;
let cycle = 0;

function increaseCycle() {
  cycle++;
  if (cyclesToCheck.includes(cycle)) {
    cycleValues.push(cycle * x);
  }
}

for (const instruction of parsedInput) {
  increaseCycle();
  const [command, args] = instruction.split(' ');

  if (command === 'noop') {
    continue;
  }

  increaseCycle();

  x += Number(args);
}

console.log(cycleValues.reduce((a, b) => a + b, 0));
