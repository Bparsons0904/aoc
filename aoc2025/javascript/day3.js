import { readFile } from './utilities.js';

function processDay3File() {
  const lines = readFile('day3.part1');
  const batteryPacks = [];

  for (const line of lines) {
    const batteryPack = [];
    for (const batteryChar of line) {
      const battery = parseInt(batteryChar);
      batteryPack.push(battery);
    }
    batteryPacks.push(batteryPack);
  }

  return batteryPacks;
}

function iterateBatteryPackCheck(battery, startingIndex, currentBatterySize, batterySize) {
  let maxFoundJoltage = 0;
  let endIndex = startingIndex;
  const stop = battery.length + currentBatterySize - (batterySize - 1);

  for (let i = startingIndex; i < stop; i++) {
    if (battery[i] > maxFoundJoltage) {
      maxFoundJoltage = battery[i];
      endIndex = i;
    }
  }

  return [maxFoundJoltage, endIndex];
}

function getLargestPackJoltage(batteryPack, batterySize) {
  const maxJoltage = [];

  let index = 0;
  let iteration = 0;

  while (maxJoltage.length < batterySize) {
    const [maxFoundJoltage, endIndex] = iterateBatteryPackCheck(
      batteryPack,
      index,
      iteration,
      batterySize
    );
    maxJoltage.push(maxFoundJoltage);
    index = endIndex + 1;
    iteration++;
  }

  return parseInt(maxJoltage.join(''));
}

function calculate12PackJoltage(batteryPacks, batterySize) {
  let maxJoltage = 0;
  for (const batteryPack of batteryPacks) {
    maxJoltage += getLargestPackJoltage(batteryPack, batterySize);
  }
  return maxJoltage;
}

function day3() {
  const batteryPacks = processDay3File();

  let start = performance.now();
  const part1Result = calculate12PackJoltage(batteryPacks, 2);
  const part1Time = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const part2Result = calculate12PackJoltage(batteryPacks, 12);
  const part2Time = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 3:`);
  console.log(`  Part 1: ${part1Result} (${part1Time}s)`);
  console.log(`  Part 2: ${part2Result} (${part2Time}s)`);
}

day3();
