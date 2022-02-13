// const input = require('./input');
const input = require('./secondInput');

const sevenDigitDisplayMapper = {
  acedgfb: 8,
  cdfbe: 5,
  gcdfa: 2,
  fbcad: 3,
  dab: 7,
  cefabd: 9,
  cdfgeb: 6,
  eafb: 4,
  cagedb: 0,
  ab: 1,
};

const mappedWireInput = input.map(wireSignal => {
  return {
    unique: wireSignal.split('|')[0].trim().split(' '),
    fourDigitOutput: wireSignal.split('|')[1].trim().split(' '),
  };
});

let sum = 0;

mappedWireInput.forEach(wireSignal => {
  const outputDecimalDigits = wireSignal.fourDigitOutput.map(output => {
    const encodedDigits = Object.keys(sevenDigitDisplayMapper);
    encodedDigits.forEach(encodedDigit => {
      if (
        output.length === encodedDigit.length &&
        encodedDigit.split('').every(digit => output.split('').includes(digit))
      ) {
        console.log('inside', sevenDigitDisplayMapper[encodedDigit]);
        return sevenDigitDisplayMapper[encodedDigit];
      }
    });
  });

  sum += Number(outputDecimalDigits.join(''));
});

console.log(sum);
