import { readFileSync } from 'fs';
import { set } from 'lodash';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input).split('\n');

let currentPath = '/';

const commands = {
  cd: args => {
    if (args[0] === '..') {
      currentPath = currentPath.split('.').slice(0, -1).join('.');
    } else if (args[0] === '/') {
      currentPath = '/';
    } else {
      currentPath = `${currentPath}.${args[0]}`;
    }
  },
  ls: args => {
    // ignore
  },
  file: args => {
    if (Number.isInteger(Number(args[0]))) {
      set(
        storage,
        `${currentPath}.${args[1].replace('.', '-')}`,
        Number(args[0]),
      );
    }
  },
};

const storage = {};

for (const line of parsedInput) {
  let command;
  let args;
  if (line.startsWith('$')) {
    [, command, ...args] = line.split(' ');
  } else {
    command = 'file';
    args = line.split(' ');
  }
  commands[command](args);
}

let validSpace = 0;

function setDirSize(dir, limit) {
  let size = 0;
  for (const key in dir) {
    if (typeof dir[key] === 'number') {
      size += dir[key];
    } else {
      size += setDirSize(dir[key], limit);
    }
  }
  dir.size = size;
  if (dir.size <= limit) {
    validSpace += dir.size;
  }
  return size;
}

setDirSize(storage, 100000);
console.log(validSpace);
