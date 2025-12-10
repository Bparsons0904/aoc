import { Grid, TACHYON, Point } from './grid.js';

function locateTachyon(tachyonGraph, tachyonGrid, point) {
  for (let i = point.y; i < tachyonGrid.map.length; i++) {
    const key = `${point.x},${i}`;
    if (tachyonGraph.has(key)) {
      return tachyonGraph.get(key);
    }
  }
  return 1;
}

function processTachyonBeamRoutesCounter(tachyonGrid) {
  const tachyonGraph = new Map();

  const rowLength = tachyonGrid.map[0].length;
  for (let i = tachyonGrid.map.length - 1; i >= 0; i--) {
    for (let j = 0; j < rowLength; j++) {
      if (tachyonGrid.map[i][j] === TACHYON) {
        let tachyonPathCount = 0;

        if (j - 1 >= 0) {
          tachyonPathCount = locateTachyon(
            tachyonGraph,
            tachyonGrid,
            new Point(j - 1, i)
          );
        }

        if (j + 1 < rowLength) {
          tachyonPathCount += locateTachyon(
            tachyonGraph,
            tachyonGrid,
            new Point(j + 1, i)
          );
        }

        const key = `${j},${i}`;
        tachyonGraph.set(key, tachyonPathCount);
      }
    }
  }

  const total = locateTachyon(
    tachyonGraph,
    tachyonGrid,
    new Point(tachyonGrid.current.x, tachyonGrid.current.y)
  );

  return total;
}

function processTachyonBeamSplitCounter(tachyonGrid) {
  let tachyonSplitCounter = 0;
  const tachyonCurrentLines = new Map();
  tachyonCurrentLines.set(tachyonGrid.current.x, true);

  for (const row of tachyonGrid.map) {
    const newTachyonCurrentLines = new Map();
    for (let x = 0; x < row.length; x++) {
      const space = row[x];
      if (space === TACHYON && tachyonCurrentLines.has(x)) {
        tachyonSplitCounter++;
        newTachyonCurrentLines.set(x - 1, true);
        newTachyonCurrentLines.set(x + 1, true);
        tachyonCurrentLines.delete(x);
      }
    }

    for (const x of newTachyonCurrentLines.keys()) {
      if (x >= 0 && x < tachyonGrid.map[0].length) {
        tachyonCurrentLines.set(x, true);
      }
    }
  }

  return tachyonSplitCounter;
}

function day7() {
  const tachyonGrid = new Grid('day7.part1');

  let start = performance.now();
  const part1Count = processTachyonBeamSplitCounter(tachyonGrid);
  const part1Time = ((performance.now() - start) / 1000).toFixed(4);

  const tachyonGrid2 = new Grid('day7.part1');
  start = performance.now();
  const part2Count = processTachyonBeamRoutesCounter(tachyonGrid2);
  const part2Time = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 7:`);
  console.log(`  Part 1: ${part1Count} (${part1Time}s)`);
  console.log(`  Part 2: ${part2Count} (${part2Time}s)`);
}

day7();
