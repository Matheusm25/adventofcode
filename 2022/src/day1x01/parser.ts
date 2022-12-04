export function parse(input: string) {
  const elves = input.split('\n\n');
  return elves.map(elf => {
    return elf
      .trim()
      .split('\n')
      .map(line => Number(line));
  });
}
