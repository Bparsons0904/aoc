import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string
import simplifile

pub type DialInstruction {
  DialInstruction(direction: String, steps: Int)
}

pub fn solve() {
  let assert Ok(content) = simplifile.read("../input/day1.part1")
  let instructions = parse_instructions(content)

  let #(part1, part2) = process_instructions(instructions)

  io.println("Day 1:")
  io.println("  Part 1: " <> int.to_string(part1))
  io.println("  Part 2: " <> int.to_string(part2))
}

fn parse_instructions(content: String) -> List(DialInstruction) {
  content
  |> string.trim
  |> string.split("\n")
  |> list.filter(fn(line) { string.length(line) > 0 })
  |> list.map(parse_instruction)
}

fn parse_instruction(line: String) -> DialInstruction {
  let direction = string.slice(line, 0, 1)
  let steps_str = string.slice(line, 1, string.length(line) - 1)
  let assert Ok(steps) = int.parse(steps_str)
  DialInstruction(direction: direction, steps: steps)
}

fn process_instructions(instructions: List(DialInstruction)) -> #(Int, Int) {
  let initial_state = #(50, 0, 0)

  let #(_, part1_count, part2_count) =
    list.fold(instructions, initial_state, fn(state, instruction) {
      let #(current_value, step1_count, step2_count) = state
      process_instruction(current_value, step1_count, step2_count, instruction)
    })

  #(part1_count, part2_count)
}

fn process_instruction(
  current_value: Int,
  step1_count: Int,
  step2_count: Int,
  instruction: DialInstruction,
) -> #(Int, Int, Int) {
  let #(new_value, new_step2_count) = case instruction.direction {
    "R" -> {
      let new_val = current_value + instruction.steps
      let crosses = case new_val >= 100 {
        True ->
          case current_value == 0 {
            True -> instruction.steps / 100
            False -> { instruction.steps + current_value } / 100
          }
        False -> 0
      }
      #(new_val % 100, step2_count + crosses)
    }
    "L" -> {
      let new_val = current_value - instruction.steps
      let crosses = case current_value == 0 {
        True -> instruction.steps / 100
        False ->
          case instruction.steps >= current_value {
            True -> { instruction.steps - current_value } / 100 + 1
            False -> 0
          }
      }
      let wrapped_value = { { new_val % 100 } + 100 } % 100
      #(wrapped_value, step2_count + crosses)
    }
    _ -> #(current_value, step2_count)
  }

  let new_step1_count = case new_value == 0 {
    True -> step1_count + 1
    False -> step1_count
  }

  #(new_value, new_step1_count, new_step2_count)
}
