import { readFileSync } from 'fs';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input);

let processed;

for (let i = 0; i < parsedInput.length - 3; i++) {
  const chars = parsedInput.slice(i, i + 4).split('');
  if (new Set(chars).size === 4) {
    processed = i + 4;
    break;
  }
}

console.log({ processed });
