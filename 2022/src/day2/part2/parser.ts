interface RPSGame {
  yourChoice?: 'rock' | 'papper' | 'scissors';
  opponentChoice: 'rock' | 'papper' | 'scissors';
  expectedResult?: 'win' | 'lose' | 'draw';
}

const parseOptions = {
  A: 'rock',
  B: 'papper',
  C: 'scissors',
  Y: 'draw',
  X: 'lose',
  Z: 'win',
};

export function parse(input: string): Array<RPSGame> {
  const game = input.split('\n');
  return game.map(line => {
    const [opponentChoice, expectedResult] = line.split(' ');
    return {
      expectedResult: parseOptions[expectedResult as string],
      opponentChoice: parseOptions[opponentChoice as string],
    };
  });
}
