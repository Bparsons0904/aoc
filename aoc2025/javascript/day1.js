import { readFile } from './utilities.js';

class DialInstruction {
  constructor(direction, step) {
    this.direction = direction;
    this.step = step;
  }
}

class Dial {
  constructor(value, next = null, prev = null) {
    this.value = value;
    this.next = next;
    this.prev = prev;
  }
}

class DialList {
  constructor() {
    this.head = null;
    this.tail = null;
  }

  next() {
    this.head = this.head.next;
  }

  prev() {
    this.head = this.head.prev;
  }

  buildLinkedList() {
    let prevNode = null;
    for (let i = 0; i < 100; i++) {
      const node = new Dial(i, null, prevNode);

      if (this.head === null) {
        this.head = node;
      }

      if (prevNode !== null) {
        prevNode.next = node;
      }

      this.tail = node;
      prevNode = node;
    }

    this.tail.next = this.head;
    this.head.prev = this.tail;
  }
}

function parseInstructions(lines) {
  const instructions = [];
  for (const line of lines) {
    const direction = line[0];
    const step = parseInt(line.substring(1));
    instructions.push(new DialInstruction(direction, step));
  }
  return instructions;
}

function day1() {
  const start = performance.now();

  const files = readFile('day1.part1');
  const dialInstructions = parseInstructions(files);

  const dialList = new DialList();
  dialList.buildLinkedList();

  while (dialList.head.value !== 50) {
    dialList.next();
  }

  let day1part1Count = 0;
  let day1part2Count = 0;

  for (const dialInstruction of dialInstructions) {
    switch (dialInstruction.direction) {
      case 'R':
        for (let j = 0; j < dialInstruction.step; j++) {
          dialList.next();
          if (dialList.head.value === 0) {
            day1part2Count++;
          }
        }
        break;
      case 'L':
        for (let j = 0; j < dialInstruction.step; j++) {
          dialList.prev();
          if (dialList.head.value === 0) {
            day1part2Count++;
          }
        }
        break;
    }
    if (dialList.head.value === 0) {
      day1part1Count++;
    }
  }

  const elapsed = ((performance.now() - start) / 1000).toFixed(4);
  console.log(`Day 1:`);
  console.log(`  Part 1: ${day1part1Count}, Part 2: ${day1part2Count} (${elapsed}s)`);
}

function day1_1() {
  const start = performance.now();

  const files = readFile('day1.part1');
  const dialInstructions = parseInstructions(files);

  let step1Count = 0;
  let step2Count = 0;
  let currentValue = 50;

  for (const dialInstruction of dialInstructions) {
    switch (dialInstruction.direction) {
      case 'R': {
        const newValue = currentValue + dialInstruction.step;
        if (newValue >= 100) {
          if (currentValue === 0) {
            step2Count += Math.floor(dialInstruction.step / 100);
          } else {
            step2Count += Math.floor((dialInstruction.step + currentValue) / 100);
          }
        }
        currentValue = newValue % 100;
        break;
      }
      case 'L': {
        const newValue = currentValue - dialInstruction.step;
        if (currentValue === 0) {
          step2Count += Math.floor(dialInstruction.step / 100);
        } else if (dialInstruction.step >= currentValue) {
          step2Count += Math.floor((dialInstruction.step - currentValue) / 100) + 1;
        }
        currentValue = ((newValue % 100) + 100) % 100;
        break;
      }
    }
    if (currentValue === 0) {
      step1Count++;
    }
  }

  const elapsed = ((performance.now() - start) / 1000).toFixed(4);
  console.log(`\nDay 1 Optimized:`);
  console.log(`  Part 1: ${step1Count}, Part 2: ${step2Count} (${elapsed}s)`);
}

day1();
day1_1();
