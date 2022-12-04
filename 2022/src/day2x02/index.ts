import { readFileSync } from 'fs';
import { parse } from './parser';

const testSample = readFileSync(__dirname + '/test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/input.txt', 'utf-8');

const parsedInput = parse(input);

const rules = {
  rock: {
    points: 1,
    rock: 'draw',
    papper: 'lose',
    scissors: 'win',
  },
  papper: {
    points: 2,
    rock: 'win',
    papper: 'draw',
    scissors: 'lose',
  },
  scissors: {
    points: 3,
    rock: 'lose',
    papper: 'win',
    scissors: 'draw',
  },
  win: 6,
  lose: 0,
  draw: 3,
};

const score = parsedInput.reduce((acc, game) => {
  const invertWin = game.expectedResult === 'win' ? 'lose' : 'win';
  const mustPlay = Object.keys(rules[game.opponentChoice]).find(
    key =>
      rules[game.opponentChoice][key] ===
      (game.expectedResult !== 'draw' ? invertWin : game.expectedResult),
  ) as string;
  return (
    acc + rules[rules[mustPlay][game.opponentChoice]] + rules[mustPlay].points
  );
}, 0);

console.log({ score });
