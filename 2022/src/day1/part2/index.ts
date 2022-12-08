import { readFileSync } from 'fs';
import { parse } from './parser';

const testSample = readFileSync(__dirname + '/test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/input.txt', 'utf-8');

const parsedInput = parse(input);
const orderedInput = parsedInput.sort((a, b) => b - a);

console.log(
  Number(orderedInput[0]) + Number(orderedInput[1]) + Number(orderedInput[2]),
);
