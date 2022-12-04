interface RPSGame {
  yourChoice: 'rock' | 'papper' | 'scissors';
  opponentChoice: 'rock' | 'papper' | 'scissors';
}

const parseOptions = {
  A: 'rock',
  B: 'papper',
  C: 'scissors',
  Y: 'papper',
  X: 'rock',
  Z: 'scissors',
};

export function parse(input: string): Array<RPSGame> {
  const game = input.split('\n');
  return game.map(line => {
    const [opponentChoice, yourChoice] = line.split(' ');
    return {
      yourChoice: parseOptions[yourChoice as string],
      opponentChoice: parseOptions[opponentChoice as string],
    };
  });
}
