import gleam/int
import gleam/io
import gleam/list
import gleam/order
import gleam/string
import simplifile

pub type IngredientRange {
  IngredientRange(min: Int, max: Int)
}

pub type Ingredients {
  Ingredients(
    fresh_ingredient_ranges: List(IngredientRange),
    ingredients: List(Int),
  )
}

pub fn solve() {
  let assert Ok(content) = simplifile.read("input/day5.part1")
  let ingredients = parse_ingredients(content)

  let #(part1, part2) = count_fresh_ingredients(ingredients)

  io.println("Day 5:")
  io.println("  Part 1: " <> int.to_string(part1))
  io.println("  Part 2: " <> int.to_string(part2))
}

fn parse_ingredients(content: String) -> Ingredients {
  let lines =
    content
    |> string.trim
    |> string.split("\n")

  let #(range_lines, ingredient_lines) = split_at_empty(lines, [], [])

  let ingredient_ranges =
    range_lines
    |> list.map(parse_ingredient_range)
    |> list.sort(fn(a, b) {
      case a.min < b.min {
        True -> order.Lt
        False ->
          case a.min == b.min {
            True -> order.Eq
            False -> order.Gt
          }
      }
    })

  let merged_ranges = merge_ranges(ingredient_ranges, [])

  let ingredients_list =
    ingredient_lines
    |> list.map(fn(line) {
      let assert Ok(val) = int.parse(line)
      val
    })

  Ingredients(fresh_ingredient_ranges: merged_ranges, ingredients: ingredients_list)
}

fn split_at_empty(
  lines: List(String),
  ranges: List(String),
  ingredients: List(String),
) -> #(List(String), List(String)) {
  case lines {
    [] -> #(list.reverse(ranges), list.reverse(ingredients))
    ["", ..rest] -> collect_ingredients(rest, list.reverse(ranges), [])
    [line, ..rest] -> split_at_empty(rest, [line, ..ranges], ingredients)
  }
}

fn collect_ingredients(
  lines: List(String),
  ranges: List(String),
  ingredients: List(String),
) -> #(List(String), List(String)) {
  case lines {
    [] -> #(ranges, list.reverse(ingredients))
    [line, ..rest] ->
      collect_ingredients(rest, ranges, [line, ..ingredients])
  }
}

fn parse_ingredient_range(line: String) -> IngredientRange {
  let assert [min_str, max_str] = string.split(line, "-")
  let assert Ok(min) = int.parse(min_str)
  let assert Ok(max) = int.parse(max_str)
  IngredientRange(min: min, max: max)
}

fn merge_ranges(
  ranges: List(IngredientRange),
  merged: List(IngredientRange),
) -> List(IngredientRange) {
  case ranges {
    [] -> list.reverse(merged)
    [range, ..rest] ->
      case merged {
        [] -> merge_ranges(rest, [range])
        [last, ..other_merged] ->
          case can_merge(range, last) {
            True -> {
              let new_range =
                IngredientRange(
                  min: int.min(last.min, range.min),
                  max: int.max(last.max, range.max),
                )
              merge_ranges(rest, [new_range, ..other_merged])
            }
            False -> merge_ranges(rest, [range, last, ..other_merged])
          }
      }
  }
}

fn can_merge(range1: IngredientRange, range2: IngredientRange) -> Bool {
  { range1.min >= range2.min && range1.min <= range2.max }
  || { range1.max >= range2.min && range1.max <= range2.max }
}

fn count_fresh_ingredients(ingredients: Ingredients) -> #(Int, Int) {
  let count_of_fresh =
    ingredients.ingredients
    |> list.filter(fn(ingredient) {
      is_in_ranges(ingredient, ingredients.fresh_ingredient_ranges)
    })
    |> list.length

  let count_of_potential =
    ingredients.fresh_ingredient_ranges
    |> list.map(fn(range) { range.max - range.min + 1 })
    |> list.fold(0, fn(acc, val) { acc + val })

  #(count_of_fresh, count_of_potential)
}

fn is_in_ranges(ingredient: Int, ranges: List(IngredientRange)) -> Bool {
  case ranges {
    [] -> False
    [range, ..rest] ->
      case ingredient >= range.min && ingredient <= range.max {
        True -> True
        False -> is_in_ranges(ingredient, rest)
      }
  }
}
