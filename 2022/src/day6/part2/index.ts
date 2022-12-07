import { readFileSync } from 'fs';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input);

let processed;

for (let i = 0; i < parsedInput.length - 13; i++) {
  const chars = parsedInput.slice(i, i + 14).split('');
  if (new Set(chars).size === 14) {
    processed = i + 14;
    break;
  }
}

console.log({ processed });
