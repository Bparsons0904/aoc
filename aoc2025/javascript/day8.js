import { readFile } from './utilities.js';

class JunctionBox {
  constructor(x, y, z) {
    this.x = x;
    this.y = y;
    this.z = z;
  }

  toString() {
    return `${this.x},${this.y},${this.z}`;
  }
}

class JunctionBoxGroup {
  constructor(connections = []) {
    this.connections = connections;
  }
}

class ShortestConnection {
  constructor(distance, from, to) {
    this.distance = distance;
    this.from = from;
    this.to = to;
  }
}

function buildJunctionBoxes(filename) {
  const lines = readFile(filename);
  const junctionBoxes = [];
  const junctionBoxConnectionIndex = new Map();

  for (const row of lines) {
    const coordinates = row.split(',');
    const x = parseInt(coordinates[0]);
    const y = parseInt(coordinates[1]);
    const z = parseInt(coordinates[2]);

    const newJunctionBox = new JunctionBox(x, y, z);
    junctionBoxes.push(newJunctionBox);
    junctionBoxConnectionIndex.set(newJunctionBox.toString(), -1);
  }

  return [junctionBoxes, junctionBoxConnectionIndex];
}

function getSortedConnections(junctionBoxes) {
  const shortestConnections = [];

  for (let i = 0; i < junctionBoxes.length; i++) {
    for (let k = i + 1; k < junctionBoxes.length; k++) {
      const junctionBox = junctionBoxes[i];
      const otherJunctionBox = junctionBoxes[k];

      const dx = junctionBox.x - otherJunctionBox.x;
      const dy = junctionBox.y - otherJunctionBox.y;
      const dz = junctionBox.z - otherJunctionBox.z;

      const distance = Math.floor(Math.sqrt(dx * dx + dy * dy + dz * dz));

      shortestConnections.push(
        new ShortestConnection(distance, junctionBox, otherJunctionBox)
      );
    }
  }

  shortestConnections.sort((a, b) => a.distance - b.distance);
  return shortestConnections;
}

function buildJunctionBoxConnectionsPart1(
  shortestConnections,
  junctionBoxConnectionIndex,
  junctionBoxes,
  limit
) {
  const connections = [];
  let connectionsMade = 0;

  for (const connection of shortestConnections) {
    if (connectionsMade >= limit) {
      break;
    }

    const fromIndex = junctionBoxConnectionIndex.get(connection.from.toString());
    const toIndex = junctionBoxConnectionIndex.get(connection.to.toString());

    connectionsMade++;
    if (fromIndex !== -1 && fromIndex === toIndex) {
      continue;
    }

    if (fromIndex === -1 && toIndex === -1) {
      connections.push(new JunctionBoxGroup([connection.from, connection.to]));
      junctionBoxConnectionIndex.set(connection.from.toString(), connections.length - 1);
      junctionBoxConnectionIndex.set(connection.to.toString(), connections.length - 1);
      continue;
    }

    if (fromIndex === -1) {
      connections[toIndex].connections.push(connection.from);
      junctionBoxConnectionIndex.set(connection.from.toString(), toIndex);
      continue;
    }

    if (toIndex === -1) {
      connections[fromIndex].connections.push(connection.to);
      junctionBoxConnectionIndex.set(connection.to.toString(), fromIndex);
      continue;
    }

    let smallerIdx = fromIndex;
    let largerIdx = toIndex;
    if (connections[fromIndex].connections.length > connections[toIndex].connections.length) {
      smallerIdx = toIndex;
      largerIdx = fromIndex;
    }

    for (const box of connections[smallerIdx].connections) {
      junctionBoxConnectionIndex.set(box.toString(), largerIdx);
    }

    connections[largerIdx].connections.push(...connections[smallerIdx].connections);
  }

  for (const junctionBox of junctionBoxes) {
    if (junctionBoxConnectionIndex.get(junctionBox.toString()) === -1) {
      connections.push(new JunctionBoxGroup([junctionBox]));
      junctionBoxConnectionIndex.set(junctionBox.toString(), connections.length - 1);
    }
  }

  connections.sort((a, b) => b.connections.length - a.connections.length);

  let result = 0;
  for (let i = 0; i < Math.min(3, connections.length); i++) {
    if (result === 0) {
      result = connections[i].connections.length;
    } else {
      result *= connections[i].connections.length;
    }
  }

  return result;
}

function buildJunctionBoxConnectionsPart2(
  shortestConnections,
  junctionBoxConnectionIndex,
  junctionBoxes
) {
  const connections = [];
  let connectionsMade = 0;

  for (let i = 0; i < shortestConnections.length; i++) {
    const connection = shortestConnections[i];

    if (connections.length === 1 && connections[0].connections.length === junctionBoxes.length) {
      return shortestConnections[i - 1].from.x * shortestConnections[i - 1].to.x;
    }

    const fromIndex = junctionBoxConnectionIndex.get(connection.from.toString());
    const toIndex = junctionBoxConnectionIndex.get(connection.to.toString());

    connectionsMade++;
    if (fromIndex !== -1 && fromIndex === toIndex) {
      continue;
    }

    if (fromIndex === -1 && toIndex === -1) {
      connections.push(new JunctionBoxGroup([connection.from, connection.to]));
      junctionBoxConnectionIndex.set(connection.from.toString(), connections.length - 1);
      junctionBoxConnectionIndex.set(connection.to.toString(), connections.length - 1);
      continue;
    }

    if (fromIndex === -1) {
      connections[toIndex].connections.push(connection.from);
      junctionBoxConnectionIndex.set(connection.from.toString(), toIndex);
      continue;
    }

    if (toIndex === -1) {
      connections[fromIndex].connections.push(connection.to);
      junctionBoxConnectionIndex.set(connection.to.toString(), fromIndex);
      continue;
    }

    let smallerIdx = fromIndex;
    let largerIdx = toIndex;
    if (connections[fromIndex].connections.length > connections[toIndex].connections.length) {
      smallerIdx = toIndex;
      largerIdx = fromIndex;
    }

    for (const box of connections[smallerIdx].connections) {
      junctionBoxConnectionIndex.set(box.toString(), largerIdx);
    }

    connections[largerIdx].connections.push(...connections[smallerIdx].connections);

    connections.splice(smallerIdx, 1);

    for (const [key, idx] of junctionBoxConnectionIndex.entries()) {
      if (idx > smallerIdx) {
        junctionBoxConnectionIndex.set(key, idx - 1);
      }
    }
  }

  return 0;
}

function day8() {
  let [junctionBoxes, junctionBoxConnectionIndex] = buildJunctionBoxes('day8.part1');
  const connections = getSortedConnections(junctionBoxes);

  let start = performance.now();
  const part1Results = buildJunctionBoxConnectionsPart1(
    connections,
    junctionBoxConnectionIndex,
    junctionBoxes,
    1000
  );
  const part1Time = ((performance.now() - start) / 1000).toFixed(4);

  [junctionBoxes, junctionBoxConnectionIndex] = buildJunctionBoxes('day8.part1');
  start = performance.now();
  const part2Results = buildJunctionBoxConnectionsPart2(
    connections,
    junctionBoxConnectionIndex,
    junctionBoxes
  );
  const part2Time = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 8:`);
  console.log(`  Part 1: ${part1Results} (${part1Time}s)`);
  console.log(`  Part 2: ${part2Results} (${part2Time}s)`);
}

day8();
