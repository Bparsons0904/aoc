import { readFile } from './utilities.js';

// Constants
export const EMPTY = '.';
export const PAPER_ROLL = '@';
export const MOVED_UP = '^';
export const MOVED_DOWN = 'v';
export const MOVED_LEFT = '<';
export const MOVED_RIGHT = '>';
export const START = 'S';
export const TACHYON = MOVED_UP;

// Direction constants
export const RIGHT = { x: 1, y: 0 };
export const LEFT = { x: -1, y: 0 };
export const DOWN = { x: 0, y: 1 };
export const UP = { x: 0, y: -1 };
export const RIGHT_DOWN = { x: 1, y: 1 };
export const RIGHT_UP = { x: 1, y: -1 };
export const LEFT_DOWN = { x: -1, y: 1 };
export const LEFT_UP = { x: -1, y: -1 };

export const DIRECTIONS = [LEFT, RIGHT, UP, DOWN, LEFT_DOWN, LEFT_UP, RIGHT_DOWN, RIGHT_UP];

export class Point {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }

  equals(other) {
    return this.x === other.x && this.y === other.y;
  }

  toString() {
    return `Point(${this.x}, ${this.y})`;
  }
}

export class Visit {
  constructor(point, direction) {
    this.point = point;
    this.direction = direction;
  }
}

export class Grid {
  constructor(filename) {
    const lines = readFile(filename);
    this.height = lines.length;
    this.width = lines.length > 0 ? lines[0].length : 0;
    this.map = [];
    this.visited = [];
    this.start = null;
    this.current = null;
    this.tiles = {};

    for (let i = 0; i < lines.length; i++) {
      this.map[i] = [];
      for (let j = 0; j < lines[i].length; j++) {
        const char = lines[i][j];
        this.map[i][j] = char;
        if (char === START) {
          this.start = new Point(j, i);
          this.current = new Point(j, i);
          this.visited.push(new Visit(new Point(j, i), { x: 0, y: 0 }));
        }
      }
    }
  }

  static makeGridByPoints(sizePoints, ...tiles) {
    let maxX = 0;
    let maxY = 0;
    for (const point of sizePoints) {
      if (point.x > maxX) maxX = point.x;
      if (point.y > maxY) maxY = point.y;
    }

    const grid = new Grid.__emptyGrid();
    grid.width = maxX + 1;
    grid.height = maxY + 1;
    grid.map = Array.from({ length: maxY + 1 }, () => Array(maxX + 1).fill(EMPTY));
    grid.tiles = {};
    grid.visited = [];

    for (const tileDef of tiles) {
      grid.tiles[tileDef.char] = tileDef.color;
      for (const point of tileDef.points) {
        grid.map[point.y][point.x] = tileDef.char;
      }
    }

    return grid;
  }

  static __emptyGrid() {
    const grid = Object.create(Grid.prototype);
    return grid;
  }

  setStart(point) {
    this.start = point;
    this.current = point;
    this.visited = [new Visit(point, { x: 0, y: 0 })];
  }

  setObject(point, object) {
    this.map[point.y][point.x] = object;
  }

  positionContainsObject(point, object) {
    if (!this.pointWithinBounds(point)) {
      return false;
    }
    return this.map[point.y][point.x] === object;
  }

  pointWithinBounds(point) {
    return !(point.x < 0 || point.x >= this.width || point.y < 0 || point.y >= this.height);
  }

  canMoveRight() {
    return this.current.x < this.width - 1;
  }

  canMoveLeft() {
    return this.current.x > 0;
  }

  canMoveDown() {
    return this.current.y < this.height - 1;
  }

  canMoveUp() {
    return this.current.y > 0;
  }

  canMove(direction) {
    if (direction === RIGHT || (direction.x === 1 && direction.y === 0)) {
      return this.canMoveRight();
    }
    if (direction === LEFT || (direction.x === -1 && direction.y === 0)) {
      return this.canMoveLeft();
    }
    if (direction === DOWN || (direction.x === 0 && direction.y === 1)) {
      return this.canMoveDown();
    }
    if (direction === UP || (direction.x === 0 && direction.y === -1)) {
      return this.canMoveUp();
    }

    let canMoveX = true;
    let canMoveY = true;

    if (direction.x > 0) {
      canMoveX = this.canMoveRight();
    } else if (direction.x < 0) {
      canMoveX = this.canMoveLeft();
    }

    if (direction.y > 0) {
      canMoveY = this.canMoveDown();
    } else if (direction.y < 0) {
      canMoveY = this.canMoveUp();
    }

    return canMoveX && canMoveY;
  }

  move(direction) {
    const canMove = this.canMove(direction);

    if (!canMove) {
      return false;
    }

    const point = new Point(
      this.current.x + direction.x,
      this.current.y + direction.y
    );

    const visit = new Visit(point, direction);
    this.visited.push(visit);
    this.current = point;

    return true;
  }

  findLastObjectToRight(start, object) {
    for (let x = this.width - 1; x > start.x; x--) {
      if (this.map[start.y][x] === object) {
        return new Point(x, start.y);
      }
    }
    return new Point(0, 0);
  }

  findLastObjectToBottom(start, object) {
    for (let y = this.height - 1; y > start.y; y--) {
      if (this.map[y][start.x] === object) {
        return new Point(start.x, y);
      }
    }
    return new Point(0, 0);
  }

  print() {
    for (const row of this.map) {
      console.log(row.join(''));
    }
  }

  printVisited() {
    const gridCopy = this.map.map(row => [...row]);
    const visitMap = new Map();

    for (const visit of this.visited) {
      gridCopy[visit.point.y][visit.point.x] = this._directionToArrow(visit.direction);
      visitMap.set(`${visit.point.x},${visit.point.y}`, true);
    }

    for (let y = 0; y < gridCopy.length; y++) {
      console.log(gridCopy[y].join(''));
    }
  }

  _directionToArrow(direction) {
    if (direction.x === 0 && direction.y === 0) {
      return START;
    }
    if (direction.x === 0 && direction.y === -1) {
      return MOVED_UP;
    }
    if (direction.x === 0 && direction.y === 1) {
      return MOVED_DOWN;
    }
    if (direction.x === -1 && direction.y === 0) {
      return MOVED_LEFT;
    }
    if (direction.x === 1 && direction.y === 0) {
      return MOVED_RIGHT;
    }
    return 'X';
  }
}
