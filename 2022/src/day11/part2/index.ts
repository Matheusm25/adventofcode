import { readFileSync } from 'fs';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input).split('\n\n');

const monkeys: Array<any> = [];

for (const monkey of parsedInput) {
  const infos = monkey.split('\n');
  const monkeyData = {
    items: infos[1]
      .split(':')[1]
      .trim()
      .split(', ')
      .map(item => Number(item)),
    operation: old => {
      const operation = infos[2].split('=')[1].trim();
      return eval(operation.replace(/old/g, old));
    },
    test: value => value % Number(infos[3].split('by')[1].trim()) === 0,
    testModule: Number(infos[3].split('by')[1].trim()),
    ifTrue: Number(infos[4].split('monkey')[1].trim()),
    ifFalse: Number(infos[5].split('monkey')[1].trim()),
    itemsChecked: 0,
  };

  monkeys.push(monkeyData);
}

const superModule = monkeys.reduce((mOld, mNew) => mOld * mNew.testModule, 1);

for (let i = 0; i < 10000; i++) {
  for (const monkey of monkeys) {
    for (const item of monkey.items) {
      const worry = monkey.operation(item);
      if (monkey.test(worry)) {
        monkeys[monkey.ifTrue].items.push(worry % superModule);
      } else {
        monkeys[monkey.ifFalse].items.push(worry % superModule);
      }
    }
    monkey.itemsChecked += monkey.items.length;
    monkey.items = [];
  }
}

monkeys.sort((a, b) => b.itemsChecked - a.itemsChecked);
console.log(monkeys[0].itemsChecked * monkeys[1].itemsChecked);
