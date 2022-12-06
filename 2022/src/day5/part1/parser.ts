export function parse(input: string): {
  moves?: Array<{ count: number; from: number; to: number }>;
  stackMap: Array<Array<string>>;
} {
  const [stacks, moves] = input.split('\n\n');

  const stackMap: any = [];

  stacks?.split('\n').forEach((line, idx) => {
    if (idx !== stacks.split('\n').length - 1) {
      for (let i = 0; i < line.length; i += 4) {
        stackMap[i / 4] = stackMap[i / 4] || [];
        if (line.slice(i, i + 3).trim().length) {
          stackMap[i / 4].push(line.slice(i, i + 3));
        }
      }
    }
  });

  return {
    moves: moves?.split('\n').map(move => {
      const [, count, , from, , to] = move.split(' ');
      return { count: Number(count), from: Number(from), to: Number(to) };
    }),
    stackMap,
  };
}
