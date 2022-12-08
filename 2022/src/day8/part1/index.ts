import { readFileSync } from 'fs';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input)
  .split('\n')
  .map(line => line.split('').map(n => Number(n)));

const rows = parsedInput.map(line => line.join(''));
const columns: Array<any> = [];

for (let i = 0; i < parsedInput[0].length; i++) {
  columns.push(parsedInput.map(line => line[i]).join(''));
}

let countVisible = (parsedInput.length - 2) * 2 + parsedInput[0].length * 2;

for (let i = 1; i < parsedInput.length - 1; i++) {
  for (let j = 1; j < parsedInput[i].length - 1; j++) {
    const check = parsedInput[i][j];
    const upsies = columns[j].split('').slice(0, i);
    const downsies = columns[j].split('').slice(i + 1);
    const lefties = rows[i].split('').slice(0, j);
    const righties = rows[i].split('').slice(j + 1);

    if (
      [upsies, downsies, lefties, righties].some(
        list => !list.some(item => Number(item) >= Number(check)),
      )
    ) {
      countVisible++;
    }
  }
}

console.log({ countVisible });
