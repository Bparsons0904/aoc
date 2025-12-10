import { readFile } from './utilities.js';
import { Point } from './grid.js';

class Interval {
  constructor(start, end) {
    this.start = start;
    this.end = end;
  }
}

class SegmentTree {
  constructor(maxWidthAtY) {
    this.yCoords = Object.keys(maxWidthAtY).map(Number).sort((a, b) => a - b);
    this.yIndex = {};
    for (let i = 0; i < this.yCoords.length; i++) {
      this.yIndex[this.yCoords[i]] = i;
    }
    this.size = this.yCoords.length;
    this.tree = new Array(4 * this.size).fill(0);

    if (this.size > 0) {
      this._build(0, 0, this.size - 1, maxWidthAtY);
    }
  }

  _build(arrIndex, start, end, maxWidthAtY) {
    if (start === end) {
      this.tree[arrIndex] = maxWidthAtY[this.yCoords[start]];
      return;
    }
    const mid = Math.floor((start + end) / 2);
    this._build(2 * arrIndex + 1, start, mid, maxWidthAtY);
    this._build(2 * arrIndex + 2, mid + 1, end, maxWidthAtY);
    this.tree[arrIndex] = Math.min(this.tree[2 * arrIndex + 1], this.tree[2 * arrIndex + 2]);
  }

  query(minY, maxY) {
    if (this.size === 0) {
      return 0;
    }

    let l, r;
    if (minY in this.yIndex) {
      l = this.yIndex[minY];
    } else {
      const idx = this._binarySearchLeft(this.yCoords, minY);
      if (idx >= this.yCoords.length) {
        return 0;
      }
      l = idx;
    }

    if (maxY in this.yIndex) {
      r = this.yIndex[maxY];
    } else {
      const idx = this._binarySearchLeft(this.yCoords, maxY);
      if (idx === 0) {
        return 0;
      }
      r = idx - 1;
    }

    if (l > r) {
      if (minY in this.yIndex && minY === maxY) {
        r = l;
      } else {
        return 0;
      }
    }

    return this._query(0, 0, this.size - 1, l, r);
  }

  _query(arrIndex, start, end, qStart, qEnd) {
    if (qStart > end || qEnd < start) {
      return Infinity;
    }
    if (qStart <= start && qEnd >= end) {
      return this.tree[arrIndex];
    }
    const mid = Math.floor((start + end) / 2);
    const leftQuery = this._query(2 * arrIndex + 1, start, mid, qStart, qEnd);
    const rightQuery = this._query(2 * arrIndex + 2, mid + 1, end, qStart, qEnd);
    return Math.min(leftQuery, rightQuery);
  }

  _binarySearchLeft(arr, target) {
    let left = 0;
    let right = arr.length;
    while (left < right) {
      const mid = Math.floor((left + right) / 2);
      if (arr[mid] < target) {
        left = mid + 1;
      } else {
        right = mid;
      }
    }
    return left;
  }
}

function getRedTiles(filename) {
  const lines = readFile(filename);
  const redTiles = [];

  for (const row of lines) {
    const coordinates = row.split(',');
    const x = parseInt(coordinates[0]);
    const y = parseInt(coordinates[1]);
    redTiles.push(new Point(x, y));
  }

  return redTiles;
}

function getLargestArea(redTiles) {
  let result = 0;

  for (const point1 of redTiles) {
    for (const point2 of redTiles) {
      const width = Math.abs(point1.x - point2.x) + 1;
      const height = Math.abs(point1.y - point2.y) + 1;
      const area = width * height;
      if (area > result) {
        result = area;
      }
    }
  }

  return result;
}

function mergeIntervals(intervals) {
  if (intervals.length === 0) {
    return [];
  }

  intervals.sort((a, b) => a.start - b.start);
  const merged = [intervals[0]];

  for (let i = 1; i < intervals.length; i++) {
    const last = merged[merged.length - 1];
    const curr = intervals[i];

    if (curr.start <= last.end + 1) {
      if (curr.end > last.end) {
        last.end = curr.end;
      }
    } else {
      merged.push(curr);
    }
  }

  return merged;
}

function getInsideIntervals(redTiles) {
  const verticalEdges = [];
  const horizontalEdges = [];

  for (let i = 0; i < redTiles.length; i++) {
    const current = redTiles[i];
    const next = redTiles[(i + 1) % redTiles.length];

    if (current.x === next.x) {
      const minY = Math.min(current.y, next.y);
      const maxY = Math.max(current.y, next.y);
      verticalEdges.push({ x: current.x, minY, maxY });
    } else {
      const minX = Math.min(current.x, next.x);
      const maxX = Math.max(current.x, next.x);
      horizontalEdges.push({ y: current.y, minX, maxX });
    }
  }

  let minY = redTiles[0].y;
  let maxY = redTiles[0].y;
  for (const p of redTiles) {
    if (p.y < minY) minY = p.y;
    if (p.y > maxY) maxY = p.y;
  }

  const intervals = {};

  for (let y = minY; y <= maxY; y++) {
    const crossings = [];

    for (const edge of verticalEdges) {
      if (y >= edge.minY && y < edge.maxY) {
        crossings.push(edge.x);
      }
    }

    crossings.sort((a, b) => a - b);

    const rowIntervals = [];
    for (let i = 0; i + 1 < crossings.length; i += 2) {
      rowIntervals.push(new Interval(crossings[i], crossings[i + 1]));
    }

    for (const edge of horizontalEdges) {
      if (edge.y === y) {
        rowIntervals.push(new Interval(edge.minX, edge.maxX));
      }
    }

    if (rowIntervals.length > 0) {
      intervals[y] = mergeIntervals(rowIntervals);
    }
  }

  return intervals;
}

function isRectangleInside(minX, maxX, minY, maxY, intervals) {
  for (let y = minY; y <= maxY; y++) {
    if (!(y in intervals)) {
      return false;
    }

    const rowIntervals = intervals[y];

    let lo = 0;
    let hi = rowIntervals.length;
    while (lo < hi) {
      const mid = Math.floor((lo + hi) / 2);
      if (rowIntervals[mid].start <= minX) {
        lo = mid + 1;
      } else {
        hi = mid;
      }
    }

    if (lo === 0 || rowIntervals[lo - 1].end < maxX) {
      return false;
    }
  }
  return true;
}

function getLargestAreaWithIntervals(redTiles, intervals) {
  const maxWidthAtY = {};
  for (const [y, ivs] of Object.entries(intervals)) {
    let maxW = 0;
    for (const iv of ivs) {
      const w = iv.end - iv.start + 1;
      if (w > maxW) {
        maxW = w;
      }
    }
    maxWidthAtY[y] = maxW;
  }

  const candidates = [];
  for (let i = 0; i < redTiles.length; i++) {
    for (let j = i + 1; j < redTiles.length; j++) {
      const p1 = redTiles[i];
      const p2 = redTiles[j];
      const minX = Math.min(p1.x, p2.x);
      const maxX = Math.max(p1.x, p2.x);
      const minY = Math.min(p1.y, p2.y);
      const maxY = Math.max(p1.y, p2.y);
      const width = maxX - minX + 1;
      const height = maxY - minY + 1;
      candidates.push({ minX, maxX, minY, maxY, area: width * height });
    }
  }

  candidates.sort((a, b) => b.area - a.area);

  for (const c of candidates) {
    const width = c.maxX - c.minX + 1;

    let canFit = true;
    for (let y = c.minY; y <= c.maxY; y++) {
      if (!(y in maxWidthAtY) || maxWidthAtY[y] < width) {
        canFit = false;
        break;
      }
    }

    if (!canFit) {
      continue;
    }

    if (isRectangleInside(c.minX, c.maxX, c.minY, c.maxY, intervals)) {
      return c.area;
    }
  }

  return 0;
}

function getLargestAreaWithIntervalsOptimized(redTiles, intervals) {
  const maxWidthAtY = {};
  for (const [y, ivs] of Object.entries(intervals)) {
    let maxW = 0;
    for (const iv of ivs) {
      const w = iv.end - iv.start + 1;
      if (w > maxW) {
        maxW = w;
      }
    }
    maxWidthAtY[y] = maxW;
  }

  const segTree = new SegmentTree(maxWidthAtY);

  const candidates = [];
  for (let i = 0; i < redTiles.length; i++) {
    for (let j = i + 1; j < redTiles.length; j++) {
      const p1 = redTiles[i];
      const p2 = redTiles[j];
      const minX = Math.min(p1.x, p2.x);
      const maxX = Math.max(p1.x, p2.x);
      const minY = Math.min(p1.y, p2.y);
      const maxY = Math.max(p1.y, p2.y);
      const width = maxX - minX + 1;

      if (p1.y in maxWidthAtY && maxWidthAtY[p1.y] < width) {
        continue;
      }
      if (p2.y in maxWidthAtY && maxWidthAtY[p2.y] < width) {
        continue;
      }

      const height = maxY - minY + 1;
      candidates.push({ minX, maxX, minY, maxY, area: width * height });
    }
  }

  candidates.sort((a, b) => b.area - a.area);

  for (const c of candidates) {
    const width = c.maxX - c.minX + 1;

    const minWidthInRange = segTree.query(c.minY, c.maxY);

    if (minWidthInRange < width) {
      continue;
    }

    if (isRectangleInside(c.minX, c.maxX, c.minY, c.maxY, intervals)) {
      return c.area;
    }
  }

  return 0;
}

function day9() {
  const filename = 'day9.part1';
  const redTiles = getRedTiles(filename);

  let start = performance.now();
  const largestAreaPart1 = getLargestArea(redTiles);
  const part1Time = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const intervals = getInsideIntervals(redTiles);
  const intervalTime = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const largestAreaPart2 = getLargestAreaWithIntervals(redTiles, intervals);
  const part2Time = ((performance.now() - start) / 1000).toFixed(4);

  start = performance.now();
  const largestAreaPart2Optimized = getLargestAreaWithIntervalsOptimized(redTiles, intervals);
  const part2OptimizedTime = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 9:`);
  console.log(`  Part 1: ${largestAreaPart1} (${part1Time}s)`);
  console.log(`  Intervals: (${intervalTime}s)`);
  console.log(`  Part 2: ${largestAreaPart2} (${part2Time}s)`);
  console.log(`  Part 2 Optimized: ${largestAreaPart2Optimized} (${part2OptimizedTime}s)`);
}

day9();
