import { Grid, PAPER_ROLL, EMPTY, DIRECTIONS, Point } from './grid.js';

function movePoint(point, direction) {
  return new Point(point.x + direction.x, point.y + direction.y);
}

function countPaperRollContacts(grid, point) {
  let count = 0;

  for (const direction of DIRECTIONS) {
    if (grid.positionContainsObject(movePoint(point, direction), PAPER_ROLL)) {
      count++;
    }
  }

  return count;
}

function calculatePaperRollsQueue(grid) {
  let part1Count = 0;
  let part2Count = 0;

  const stack = [];
  for (let y = 0; y < grid.height; y++) {
    for (let x = 0; x < grid.width; x++) {
      if (grid.map[y][x] === PAPER_ROLL) {
        const point = new Point(x, y);
        const connectedRolls = countPaperRollContacts(grid, point);
        if (connectedRolls < 4) {
          part1Count++;
          stack.push(point);
        }
      }
    }
  }

  while (stack.length > 0) {
    const point = stack.pop();

    if (grid.map[point.y][point.x] !== PAPER_ROLL) {
      continue;
    }

    grid.setObject(point, EMPTY);
    part2Count++;

    for (const dir of DIRECTIONS) {
      const neighbor = movePoint(point, dir);
      if (grid.positionContainsObject(neighbor, PAPER_ROLL)) {
        const connectedRolls = countPaperRollContacts(grid, neighbor);
        if (connectedRolls < 4) {
          stack.push(neighbor);
        }
      }
    }
  }

  return [part1Count, part2Count];
}

function calculatePaperRollsOptimized(grid) {
  let part1Count = 0;
  let part2Count = 0;

  const mappedPaperRolls = new Map();
  for (let y = 0; y < grid.height; y++) {
    for (let x = 0; x < grid.width; x++) {
      if (grid.map[y][x] === PAPER_ROLL) {
        const key = `${x},${y}`;
        mappedPaperRolls.set(key, new Point(x, y));
        const connectedRolls = countPaperRollContacts(grid, new Point(x, y));
        if (connectedRolls < 4) {
          part1Count++;
        }
      }
    }
  }

  let lastPassCount = -1;
  while (lastPassCount !== 0) {
    lastPassCount = 0;

    for (const [key, point] of Array.from(mappedPaperRolls)) {
      const connectedRolls = countPaperRollContacts(grid, point);
      if (connectedRolls < 4) {
        grid.setObject(point, EMPTY);
        mappedPaperRolls.delete(key);
        lastPassCount++;
        part2Count++;
      }
    }
  }

  return [part1Count, part2Count];
}

function calculatePaperRolls(grid) {
  let part1Count = 0;
  let part2Count = 0;

  let firstPass = true;
  let lastPassCount = -1;

  while (lastPassCount !== 0) {
    lastPassCount = 0;

    for (let y = 0; y < grid.height; y++) {
      for (let x = 0; x < grid.width; x++) {
        if (grid.map[y][x] !== PAPER_ROLL) {
          continue;
        }
        const connectedRolls = countPaperRollContacts(grid, new Point(x, y));
        if (connectedRolls < 4) {
          if (firstPass) {
            part1Count++;
          } else {
            grid.setObject(new Point(x, y), EMPTY);
            lastPassCount++;
            part2Count++;
          }
        }
      }
    }

    if (firstPass) {
      firstPass = false;
      lastPassCount = -1;
    }
  }

  return [part1Count, part2Count];
}

function day4() {
  const grid = new Grid('day4.part1');

  let start = performance.now();
  const [part1CountQueue, part2CountQueue] = calculatePaperRollsQueue(grid);
  const elapsed = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 4:`);
  console.log(`  Part 1: ${part1CountQueue}, Part 2: ${part2CountQueue} (${elapsed}s)`);
}

day4();
