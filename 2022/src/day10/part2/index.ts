import { readFileSync } from 'fs';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input).split('\n');

let spritePosition = 1;
let cycle = 0;
const position = {
  col: 0,
  row: 0,
};

const delay = delayInms => {
  return new Promise(resolve => setTimeout(resolve, delayInms));
};

const screen = new Array(6).fill('').map(() => new Array(40).fill('#'));

function printScreen() {
  console.clear();
  screen.forEach(row => {
    console.log(row.join(''));
  });
}

async function increaseCycle() {
  cycle++;

  if (
    [spritePosition - 1, spritePosition, spritePosition + 1].includes(
      position.col,
    )
  ) {
    screen[position.row][position.col] = '#';
  } else {
    screen[position.row][position.col] = '.';
  }
  printScreen();
  console.log(position.row, position.col);
  console.log(spritePosition);
  await delay(50); // delay only to see the animation

  if (position.col === 39) {
    position.col = 0;
    if (position.row === 5) {
      position.row = 0;
    } else {
      position.row++;
    }
  } else {
    position.col++;
  }
}

(async () => {
  for (const instruction of parsedInput) {
    // first cycle
    await increaseCycle();
    const [command, args] = instruction.split(' ');

    if (command === 'noop') {
      continue;
    }

    // second cycle
    await increaseCycle();

    spritePosition += Number(args);
  }
})();
