use std::cmp::{max};
use std::fs::File;
use std::io::{self, BufRead};
use std::time::Instant;

#[derive(Debug, Clone, Copy, Eq, PartialEq)]
struct IngredientRange {
    min: i64,
    max: i64,
}

struct Ingredients {
    fresh_ingredient_ranges: Vec<IngredientRange>,
    ingredients: Vec<i64>,
}

fn main() {
    let ingredients = parse_ingredients("files/day5.part1").unwrap();
    let now = Instant::now();
    let (part1_count, part2_count) = count_fresh_ingredients(&ingredients);
    let elapsed = now.elapsed();

    println!("Day 5, Part 1: {}", part1_count);
    println!("Day 5, Part 2: {}", part2_count);
    println!("Part 1 & 2 took: {:.2?}", elapsed);
}

fn count_fresh_ingredients(ingredients: &Ingredients) -> (i64, i64) {
    let mut count_of_fresh_ingredients = 0;
    for &ingredient in &ingredients.ingredients {
        for range in &ingredients.fresh_ingredient_ranges {
            if range.min <= ingredient && range.max >= ingredient {
                count_of_fresh_ingredients += 1;
                break;
            }
        }
    }

    let mut count_of_potential_fresh_ingredients = 0;
    for range in &ingredients.fresh_ingredient_ranges {
        count_of_potential_fresh_ingredients += (range.max - range.min) + 1;
    }

    (count_of_fresh_ingredients, count_of_potential_fresh_ingredients)
}

fn parse_ingredients(filename: &str) -> io::Result<Ingredients> {
    let file = File::open(filename)?;
    let reader = io::BufReader::new(file);

    let mut ingredient_ranges = Vec::new();
    let mut ingredients = Vec::new();
    let mut is_first_half = true;

    for line in reader.lines() {
        let line = line?;
        if line.is_empty() {
            is_first_half = false;
            continue;
        }

        if is_first_half {
            let parts: Vec<&str> = line.split('-').collect();
            if parts.len() == 2 {
                let min = parts[0].parse::<i64>().unwrap();
                let max = parts[1].parse::<i64>().unwrap();
                ingredient_ranges.push(IngredientRange { min, max });
            }
        } else {
            let ingredient_id = line.parse::<i64>().unwrap();
            ingredients.push(ingredient_id);
        }
    }

    ingredient_ranges.sort_by_key(|r| r.min);

    let mut fresh_ingredient_ranges = Vec::new();
    if !ingredient_ranges.is_empty() {
        let mut current_range = ingredient_ranges[0];
        for &next_range in &ingredient_ranges[1..] {
            if next_range.min <= current_range.max + 1 {
                current_range.max = max(current_range.max, next_range.max);
            } else {
                fresh_ingredient_ranges.push(current_range);
                current_range = next_range;
            }
        }
        fresh_ingredient_ranges.push(current_range);
    }
    
    Ok(Ingredients { fresh_ingredient_ranges, ingredients })
}