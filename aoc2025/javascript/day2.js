import { readFile } from './utilities.js';

class ProductIDRange {
  constructor(min, max) {
    this.min = min;
    this.max = max;
  }
}

function processDay2() {
  const row = readFile('day2.part1')[0];
  const ranges = row.split(',');
  const productIDRanges = [];

  for (const value of ranges) {
    const idRange = value.split('-');
    const min = parseInt(idRange[0]);
    const max = parseInt(idRange[1]);
    productIDRanges.push(new ProductIDRange(min, max));
  }

  return productIDRanges;
}

function calculatePart1(productIDRanges) {
  let result = 0;
  for (const productIDRange of productIDRanges) {
    for (let i = productIDRange.min; i <= productIDRange.max; i++) {
      const iString = i.toString();
      if (iString.length % 2 !== 0) {
        continue;
      }

      const half = Math.floor(iString.length / 2);
      if (iString.substring(0, half) === iString.substring(half)) {
        result += i;
      }
    }
  }

  return result;
}

// My implementation
function calculatePart2(productIDRanges) {
  let result = 0;

  for (const productIDRange of productIDRanges) {
    for (let i = productIDRange.min; i <= productIDRange.max; i++) {
      const iString = i.toString();
      for (let j = 0; j <= iString.length - 2; j++) {
        const toCheck = iString.substring(0, j + 1);
        const expectedCount = Math.ceil(iString.length / toCheck.length);
        const count = (iString.match(new RegExp(toCheck, 'g')) || []).length;
        if (count >= 2 && count === expectedCount) {
          result += i;
          break;
        }
      }
    }
  }

  return result;
}

// Combo Bob and Derek implementation
function calculatePart2_1(productIDRanges) {
  let result = 0;

  for (const productIDRange of productIDRanges) {
    for (let value = productIDRange.min; value <= productIDRange.max; value++) {
      const iString = value.toString();
      for (let j = 1; j <= Math.floor(iString.length / 2); j++) {
        if (iString.substring(j).replaceAll(iString.substring(0, j), '') === '') {
          result += value;
          break;
        }
      }
    }
  }

  return result;
}

// Claude implementation
function calculatePart2_2(productIDRanges) {
  let result = 0;

  for (const productIDRange of productIDRanges) {
    for (let value = productIDRange.min; value <= productIDRange.max; value++) {
      const iString = value.toString();
      const strLen = iString.length;

      for (let patternLen = 1; patternLen <= Math.floor(strLen / 2); patternLen++) {
        if (strLen % patternLen !== 0) {
          continue;
        }

        const pattern = iString.substring(0, patternLen);
        let isRepeating = true;

        for (let i = patternLen; i < strLen; i += patternLen) {
          if (iString.substring(i, i + patternLen) !== pattern) {
            isRepeating = false;
            break;
          }
        }

        if (isRepeating) {
          result += value;
          break;
        }
      }
    }
  }

  return result;
}

// Gemini's implementation
function calculatePart2_3(productIDRanges) {
  let result = 0;

  for (const productIDRange of productIDRanges) {
    for (let value = productIDRange.min; value <= productIDRange.max; value++) {
      const s = value.toString();
      const n = s.length;
      if (n < 2) {
        continue;
      }

      // This is a known trick to find the smallest period of a string.
      // We create a new string by concatenating s with itself, then
      // search for s within this new string, starting from the second character.
      // The index of the first match + 1 gives the length of the repeating pattern (period).
      const period = (s + s).substring(1).indexOf(s) + 1;
      if (period > 0 && period < n) {
        result += value;
      }
    }
  }

  return result;
}

function day2() {
  const productIDRanges = processDay2();

  let start = performance.now();
  const part1Count = calculatePart1(productIDRanges);
  const part1Time = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const part2Count = calculatePart2(productIDRanges);
  const part2Time = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const part2_1Count = calculatePart2_1(productIDRanges);
  const part2_1Time = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const part2_2Count = calculatePart2_2(productIDRanges);
  const part2_2Time = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const part2_3Count = calculatePart2_3(productIDRanges);
  const part2_3Time = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 2:`);
  console.log(`  Part 1: ${part1Count} (${part1Time}s)`);
  console.log(`  Part 2: ${part2Count} (${part2Time}s)`);
  console.log(`  Part 2.1: ${part2_1Count} (${part2_1Time}s)`);
  console.log(`  Part 2.2: ${part2_2Count} (${part2_2Time}s)`);
  console.log(`  Part 2.3: ${part2_3Count} (${part2_3Time}s)`);
}

day2();
