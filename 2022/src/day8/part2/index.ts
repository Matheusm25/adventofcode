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

const betterTreeHouse = { tree: -1, treeCount: 0, j: -1, i: -1 };

for (let i = 1; i < parsedInput.length - 1; i++) {
  for (let j = 1; j < parsedInput[i].length - 1; j++) {
    const check = parsedInput[i][j];
    const upsies = columns[j].split('').slice(0, i);
    const downsies = columns[j].split('').slice(i + 1);
    const lefties = rows[i].split('').slice(0, j);
    const righties = rows[i].split('').slice(j + 1);

    const mappedTrees = new Array(4).fill(0);

    upsies.reverse();
    lefties.reverse();

    let index = 0;
    for (const side of [upsies, lefties, downsies, righties]) {
      for (const treeSize of side) {
        mappedTrees[index]++;
        if (treeSize >= check) {
          break;
        }
      }
      index++;
    }

    if (mappedTrees.reduce((p, c) => p * c) > betterTreeHouse.treeCount) {
      betterTreeHouse.tree = check;
      betterTreeHouse.treeCount = mappedTrees.reduce((p, c) => p * c);
      betterTreeHouse.i = i;
      betterTreeHouse.j = j;
    }
  }
}

console.log(betterTreeHouse);
