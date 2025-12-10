import { readFile } from './utilities.js';

class IngredientRange {
  constructor(min, max) {
    this.min = min;
    this.max = max;
  }
}

class Ingredients {
  constructor() {
    this.freshIngredientRanges = [];
    this.ingredients = [];
  }

  countFreshIngredients() {
    let countOfFreshIngredients = 0;
    for (const ingredient of this.ingredients) {
      for (const freshIngredientRange of this.freshIngredientRanges) {
        if (freshIngredientRange.min <= ingredient && freshIngredientRange.max >= ingredient) {
          countOfFreshIngredients++;
          break;
        }
      }
    }

    let countOfPotentialFreshIngredients = 0;
    for (const freshIngredientRange of this.freshIngredientRanges) {
      countOfPotentialFreshIngredients += (freshIngredientRange.max - freshIngredientRange.min) + 1;
    }

    return [countOfFreshIngredients, countOfPotentialFreshIngredients];
  }

  parseIngredients(lines) {
    const ingredientRanges = [];
    let isFirstHalf = true;

    for (const row of lines) {
      if (row === '') {
        isFirstHalf = false;
        continue;
      }

      if (isFirstHalf) {
        const parts = row.split('-');
        const min = parseInt(parts[0]);
        const max = parseInt(parts[1]);
        ingredientRanges.push(new IngredientRange(min, max));
      } else {
        const ingredientID = parseInt(row);
        this.ingredients.push(ingredientID);
      }
    }

    ingredientRanges.sort((a, b) => a.min - b.min);

    for (const ingredientRange of ingredientRanges) {
      if (this.freshIngredientRanges.length === 0) {
        this.freshIngredientRanges.push(ingredientRange);
        continue;
      }

      let found = false;
      for (let i = 0; i < this.freshIngredientRanges.length; i++) {
        const freshIngredientRange = this.freshIngredientRanges[i];
        if ((ingredientRange.min >= freshIngredientRange.min && ingredientRange.min <= freshIngredientRange.max) ||
            (ingredientRange.max >= freshIngredientRange.min && ingredientRange.max <= freshIngredientRange.max)) {
          this.freshIngredientRanges[i].min = Math.min(
            this.freshIngredientRanges[i].min,
            ingredientRange.min
          );
          this.freshIngredientRanges[i].max = Math.max(
            this.freshIngredientRanges[i].max,
            ingredientRange.max
          );
          found = true;
        }
      }

      if (!found) {
        this.freshIngredientRanges.push(ingredientRange);
      }
    }
  }
}

function day5() {
  const lines = readFile('day5.part1');
  const ingredients = new Ingredients();
  ingredients.parseIngredients(lines);

  const start = performance.now();
  const [part1Count, part2Count] = ingredients.countFreshIngredients();
  const elapsed = ((performance.now() - start) / 1000).toFixed(4);

  console.log(`\nDay 5:`);
  console.log(`  Part 1: ${part1Count}, Part 2: ${part2Count} (${elapsed}s)`);
}

day5();
