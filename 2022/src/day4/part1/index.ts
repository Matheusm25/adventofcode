import { readFileSync } from 'fs';
import { parse } from './parser';

const testSample = readFileSync(__dirname + '/test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/input.txt', 'utf-8');

const parsedInput = parse(input);

let count = 0;

for (const line of parsedInput) {
  const [first, second] = line;

  const firstCoversSecond = first?.every(point => second?.includes(point));
  const secondCoversFirst = second?.every(point => first?.includes(point));

  if (firstCoversSecond || secondCoversFirst) {
    count++;
  }
}

console.log({ count });
