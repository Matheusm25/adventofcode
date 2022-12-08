import { readFileSync } from 'fs';
import { parse } from './parser';

const testSample = readFileSync(__dirname + '/test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/input.txt', 'utf-8');

const parsedInput = parse(input);

function getLetterPoint(letter: string) {
  if (letter === letter.toLowerCase()) {
    return letter.charCodeAt(0) - 96;
  } else {
    return letter.charCodeAt(0) - 38;
  }
}

let sum = 0;

for (const line of parsedInput) {
  for (const letter of String(line[0])) {
    if (line[1]?.includes(letter)) {
      sum += getLetterPoint(letter);
      break;
    }
  }
}

console.log({ sum });
