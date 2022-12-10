import { readFileSync } from 'fs';

const testSample = readFileSync(__dirname + '/../test-sample.txt', 'utf-8');
const input = readFileSync(__dirname + '/../input.txt', 'utf-8');

const parsedInput = String(input).split('\n');

const positions = {
  tail: { X: 0, Y: 0 },
  head: { X: 0, Y: 0 },
  xDistance: 0,
  yDistance: 0,
};

const tailTracking: Array<string> = [];

function pointsAreClose(point1, point2) {
  const xDistance = Math.abs(point1.X - point2.X);
  const yDistance = Math.abs(point1.Y - point2.Y);

  return xDistance <= 1 && yDistance <= 1;
}

for (const instruction of parsedInput) {
  const [direction, distance] = instruction.split(' ');
  const parsedDistance = Number(distance);

  for (let i = 0; i < parsedDistance; i++) {
    switch (direction) {
      case 'R':
        positions.head.X += 1;
        break;
      case 'U':
        positions.head.Y += 1;
        break;
      case 'L':
        positions.head.X -= 1;
        break;
      case 'D':
        positions.head.Y -= 1;
        break;
      default:
        console.log('entered default', direction, distance);
    }

    if (!pointsAreClose(positions.head, positions.tail)) {
      const xDistance = positions.head.X - positions.tail.X;
      const yDistance = positions.head.Y - positions.tail.Y;

      positions.xDistance = xDistance;
      positions.yDistance = yDistance;

      if (xDistance > 1 && yDistance === 0) {
        positions.tail.X += 1;
      } else if (xDistance < -1 && yDistance === 0) {
        positions.tail.X -= 1;
      } else if (yDistance > 1 && xDistance === 0) {
        positions.tail.Y += 1;
      } else if (yDistance < -1 && xDistance === 0) {
        positions.tail.Y -= 1;
      } else if (xDistance >= 1 && yDistance >= 1) {
        positions.tail.X += 1;
        positions.tail.Y += 1;
      } else if (xDistance >= 1 && yDistance <= -1) {
        positions.tail.X += 1;
        positions.tail.Y -= 1;
      } else if (yDistance >= 1 && xDistance <= -1) {
        positions.tail.X -= 1;
        positions.tail.Y += 1;
      } else if (yDistance >= 1 && xDistance >= 1) {
        positions.tail.X += 1;
        positions.tail.Y += 1;
      } else if (xDistance <= -1 && yDistance <= -1) {
        positions.tail.X -= 1;
        positions.tail.Y -= 1;
      }
    }
    tailTracking.push(JSON.stringify(positions.tail));
  }
}

console.log(new Set(tailTracking).size);
