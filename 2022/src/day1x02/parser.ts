export function parse(input: string): Array<number> {
  const elves = input.split('\n\n');
  return elves.map(elf => {
    return elf
      .trim()
      .split('\n')
      .map(line => Number(line))
      .reduce((acc, food) => acc + food, 0);
  });
}
