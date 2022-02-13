const rawInput = require('./rawInput');

const filteredInput = rawInput
  .filter(i => i)
  .map(i =>
    i
      .split('  ')
      .join(' ')
      .split(' ')
      .filter(inner => inner),
  );

let count = 0;
let currentIndex = 0;
const parsedInput = [[]];
filteredInput.forEach(line => {
  if (count < 5) {
    parsedInput[currentIndex].push(line);
    count++;
  } else {
    currentIndex++;
    parsedInput.push([]);
    parsedInput[currentIndex].push(line);
    count = 1;
  }
});

console.log(parsedInput);
