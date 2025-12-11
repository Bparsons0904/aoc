import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile

pub type ProductIDRange {
  ProductIDRange(min: Int, max: Int)
}

pub fn solve() {
  let assert Ok(content) = simplifile.read("input/day2.part1")
  let ranges = parse_ranges(content)

  let part1 = calculate_part1(ranges)
  let part2 = calculate_part2(ranges)

  io.println("Day 2:")
  io.println("  Part 1: " <> int.to_string(part1))
  io.println("  Part 2: " <> int.to_string(part2))
}

fn parse_ranges(content: String) -> List(ProductIDRange) {
  content
  |> string.trim
  |> string.split(",")
  |> list.map(fn(range_str) {
    let assert [min_str, max_str] = string.split(range_str, "-")
    let assert Ok(min) = int.parse(min_str)
    let assert Ok(max) = int.parse(max_str)
    ProductIDRange(min, max)
  })
}

fn is_part1_match(i: Int) -> Bool {
  let s = int.to_string(i)
  let n = string.length(s)
  case n % 2 != 0 {
    True -> False
    False -> {
      let half = n / 2
      let first_half = string.slice(s, 0, half)
      let second_half = string.slice(s, half, n - half)
      first_half == second_half
    }
  }
}

fn calculate_part1(ranges: List(ProductIDRange)) -> Int {
  ranges
  |> list.flat_map(fn(range) { list.range(range.min, range.max) })
  |> list.filter(is_part1_match)
  |> list.fold(0, fn(acc, i) { acc + i })
}

fn is_part2_match(i: Int) -> Bool {
  let s = int.to_string(i)
  let n = string.length(s)
  case n < 2 {
    True -> False
    False -> {
      let doubled_s_slice = string.slice(s <> s, 1, n * 2 - 1)
      case string.contains(doubled_s_slice, s) {
        True -> {
          let split_list = string.split(doubled_s_slice, s)
          case split_list {
            [prefix, _] -> {
              let index = string.length(prefix)
              let period = index + 1
              period > 0 && period < n
            }
            _ -> False
          }
        }
        False -> False
      }
    }
  }
}

fn calculate_part2(ranges: List(ProductIDRange)) -> Int {
  ranges
  |> list.flat_map(fn(range) { list.range(range.min, range.max) })
  |> list.filter(is_part2_match)
  |> list.fold(0, fn(acc, i) { acc + i })
}