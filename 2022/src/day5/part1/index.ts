import { readFileSync } from 'fs';
import { parse } from './parser';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = parse(input);

for (const move of parsedInput.moves || []) {
  const { count, from, to } = move;
  for (let i = 0; i < count; i++) {
    parsedInput.stackMap[to - 1]?.unshift(
      String(parsedInput.stackMap[from - 1]?.shift()),
    );
  }
}

console.log(parsedInput.stackMap.map(col => col[0]).join(''));
