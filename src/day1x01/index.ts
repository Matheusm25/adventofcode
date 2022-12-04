import { readFileSync } from 'fs';
import { parse } from './parser';

const testSample = readFileSync(__dirname + '/test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/input.txt', 'utf-8');

const parsedInput = parse(input);

let maxCalories = 0;

for (const foods of parsedInput) {
  const sumOfCalories = foods.reduce((acc, food) => acc + food, 0);
  if (sumOfCalories > maxCalories) {
    maxCalories = sumOfCalories;
  }
}

console.log(maxCalories);
