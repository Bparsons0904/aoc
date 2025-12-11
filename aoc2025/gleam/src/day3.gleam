import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile

pub type BatteryPack =
  List(Int)

pub fn solve() {
  let assert Ok(content) = simplifile.read("input/day3.part1")
  let battery_packs = parse_battery_packs(content)

  let part1 = calculate_pack_joltage(battery_packs, 2)
  let part2 = calculate_pack_joltage(battery_packs, 12)

  io.println("Day 3:")
  io.println("  Part 1: " <> int.to_string(part1))
  io.println("  Part 2: " <> int.to_string(part2))
}

fn parse_battery_packs(content: String) -> List(BatteryPack) {
  content
  |> string.trim
  |> string.split("\n")
  |> list.filter(fn(line) { string.length(line) > 0 })
  |> list.map(parse_battery_pack)
}

fn parse_battery_pack(line: String) -> BatteryPack {
  line
  |> string.to_graphemes
  |> list.map(fn(c) {
    let assert Ok(val) = int.parse(c)
    val
  })
}

fn calculate_pack_joltage(battery_packs: List(BatteryPack), battery_size: Int) -> Int {
  battery_packs
  |> list.map(fn(pack) { get_largest_pack_joltage(pack, battery_size) })
  |> list.fold(0, fn(acc, val) { acc + val })
}

fn get_largest_pack_joltage(battery_pack: BatteryPack, battery_size: Int) -> Int {
  let max_joltage = find_max_digits(battery_pack, battery_size, 0, [])

  max_joltage
  |> list.map(int.to_string)
  |> string.join("")
  |> int.parse
  |> fn(result) {
    case result {
      Ok(val) -> val
      Error(_) -> 0
    }
  }
}

fn find_max_digits(
  battery_pack: BatteryPack,
  battery_size: Int,
  iteration: Int,
  acc: List(Int),
) -> List(Int) {
  case list.length(acc) == battery_size {
    True -> acc
    False -> {
      let start_index = iteration
      let #(max_val, _end_index) =
        iterate_battery_check(battery_pack, start_index, iteration, battery_size)
      find_max_digits(
        battery_pack,
        battery_size,
        iteration + 1,
        list.append(acc, [max_val]),
      )
    }
  }
}

fn iterate_battery_check(
  battery: BatteryPack,
  starting_index: Int,
  current_battery_size: Int,
  battery_size: Int,
) -> #(Int, Int) {
  let stop = list.length(battery) + current_battery_size - { battery_size - 1 }

  iterate_check_helper(battery, starting_index, stop, 0, starting_index)
}

fn iterate_check_helper(
  battery: BatteryPack,
  current_index: Int,
  stop: Int,
  max_found: Int,
  end_index: Int,
) -> #(Int, Int) {
  case current_index >= stop {
    True -> #(max_found, end_index)
    False -> {
      let assert Ok(val) = list_get(battery, current_index)
      case val > max_found {
        True -> iterate_check_helper(battery, current_index + 1, stop, val, current_index)
        False -> iterate_check_helper(battery, current_index + 1, stop, max_found, end_index)
      }
    }
  }
}

fn list_get(lst: List(a), index: Int) -> Result(a, Nil) {
  case index < 0 {
    True -> Error(Nil)
    False -> list_get_helper(lst, index)
  }
}

fn list_get_helper(lst: List(a), index: Int) -> Result(a, Nil) {
  case lst {
    [] -> Error(Nil)
    [first, ..rest] -> case index == 0 {
      True -> Ok(first)
      False -> list_get_helper(rest, index - 1)
    }
  }
}
