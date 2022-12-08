export function parse(input: string): Array<Array<string>> {
  return input.split('\n').map(line => {
    return [line.slice(0, line.length / 2), line.slice(line.length / 2)];
  });
}
