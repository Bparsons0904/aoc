import { readFile } from './utilities.js';

class Worksheet {
  constructor(operator) {
    this.values = [];
    this.cephalopodValues = [];
    this.operator = operator;
  }
}

function buildOperators(lines) {
  const worksheetMap = {};
  const operators = lines[lines.length - 1].split(/\s+/);

  for (let i = 0; i < operators.length; i++) {
    worksheetMap[i] = new Worksheet(operators[i]);
  }

  return worksheetMap;
}

function processValues(worksheetMap, lines) {
  for (const row of lines) {
    const values = row.split(/\s+/);
    for (let j = 0; j < values.length; j++) {
      const valueInt = parseInt(values[j]);
      worksheetMap[j].values.push(valueInt);
    }
  }
}

function processCephalopodValues(worksheetMap, lines) {
  const tempArray = [];
  for (const row of lines) {
    tempArray.push(row.split(''));
  }

  const length = tempArray[0].length;
  let mapIndex = Object.keys(worksheetMap).length - 1;

  for (let j = length - 1; j >= 0; j--) {
    const columnChars = [];
    for (let i = 0; i < tempArray.length; i++) {
      columnChars.push(tempArray[i][j]);
    }

    const value = columnChars.join('').trim();
    if (value === '') {
      mapIndex--;
      continue;
    }

    const valueInt = parseInt(value);
    worksheetMap[mapIndex].cephalopodValues.push(valueInt);
  }
}

function processDay6File() {
  const lines = readFile('day6.part1');
  const worksheetMap = buildOperators(lines);
  const dataLines = lines.slice(0, -1);
  processValues(worksheetMap, dataLines);
  processCephalopodValues(worksheetMap, dataLines);

  return worksheetMap;
}

function calculateWorksheets(worksheetMap) {
  let total = 0;
  for (const ws of Object.values(worksheetMap)) {
    let subtotal = 0;
    for (const value of ws.values) {
      switch (ws.operator) {
        case '+':
          subtotal += value;
          break;
        case '*':
          if (subtotal === 0) {
            subtotal = 1;
          }
          subtotal *= value;
          break;
      }
    }
    total += subtotal;
  }

  let cephalopodTotal = 0;
  for (const ws of Object.values(worksheetMap)) {
    let cephalopodSubtotal = 0;
    for (const cephalopodValue of ws.cephalopodValues) {
      switch (ws.operator) {
        case '+':
          cephalopodSubtotal += cephalopodValue;
          break;
        case '*':
          if (cephalopodSubtotal === 0) {
            cephalopodSubtotal = 1;
          }
          cephalopodSubtotal *= cephalopodValue;
          break;
      }
    }
    cephalopodTotal += cephalopodSubtotal;
  }

  return [total, cephalopodTotal];
}

function day6() {
  const worksheet = processDay6File();

  const start = performance.now();
  const [part1Total, part2Total] = calculateWorksheets(worksheet);
  const elapsed = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 6:`);
  console.log(`  Part 1: ${part1Total}, Part 2: ${part2Total} (${elapsed}s)`);
}

day6();
