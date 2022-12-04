export function parse(input: string): Array<Array<Array<number>>> {
  return input.split('\n').map(line => {
    return line.split(',').map(range => {
      const [start, end] = range.split('-');

      return Array(Number(end) - Number(start) + 1)
        .fill(0)
        .map((_, idx) => Number(start) + idx);
    });
  });
}
