import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string
import simplifile

pub type Worksheet {
  Worksheet(values: List(Int), cephalopod_values: List(Int), operator: String)
}

pub fn solve() {
  let assert Ok(content) = simplifile.read("input/day6.part1")
  let worksheets = parse_worksheets(content)

  let #(part1, part2) = calculate_worksheets(worksheets)

  io.println("Day 6:")
  io.println("  Part 1: " <> int.to_string(part1))
  io.println("  Part 2: " <> int.to_string(part2))
}

fn parse_worksheets(content: String) -> Dict(Int, Worksheet) {
  let lines =
    content
    |> string.trim
    |> string.split("\n")
    |> list.filter(fn(line) { string.length(line) > 0 })

  let assert Ok(operator_line) = list.last(lines)
  let value_lines = list.take(lines, list.length(lines) - 1)

  let operators =
    operator_line
    |> string.split(" ")
    |> list.filter(fn(s) { string.length(s) > 0 })

  let worksheet_map = build_operators(operators)

  worksheet_map
  |> process_values(value_lines)
  |> process_cephalopod_values(value_lines)
}

fn build_operators(operators: List(String)) -> Dict(Int, Worksheet) {
  operators
  |> list.index_map(fn(op, idx) {
    #(idx, Worksheet(values: [], cephalopod_values: [], operator: op))
  })
  |> dict.from_list
}

fn process_values(
  worksheet_map: Dict(Int, Worksheet),
  lines: List(String),
) -> Dict(Int, Worksheet) {
  list.fold(lines, worksheet_map, fn(map, line) {
    let values =
      line
      |> string.split(" ")
      |> list.filter(fn(s) { string.length(s) > 0 })
      |> list.map(fn(s) {
        let assert Ok(val) = int.parse(s)
        val
      })

    list.index_fold(values, map, fn(m, value, idx) {
      case dict.get(m, idx) {
        Ok(ws) ->
          dict.insert(
            m,
            idx,
            Worksheet(
              values: list.append(ws.values, [value]),
              cephalopod_values: ws.cephalopod_values,
              operator: ws.operator,
            ),
          )
        Error(_) ->
          dict.insert(
            m,
            idx,
            Worksheet(values: [value], cephalopod_values: [], operator: ""),
          )
      }
    })
  })
}

fn process_cephalopod_values(
  worksheet_map: Dict(Int, Worksheet),
  lines: List(String),
) -> Dict(Int, Worksheet) {
  let temp_array =
    lines
    |> list.map(string.to_graphemes)

  let assert Ok(first_line) = list.first(temp_array)
  let length = list.length(first_line)
  let map_size = dict.size(worksheet_map)

  process_columns(temp_array, length - 1, map_size - 1, worksheet_map)
}

fn process_columns(
  temp_array: List(List(String)),
  col_idx: Int,
  map_idx: Int,
  worksheet_map: Dict(Int, Worksheet),
) -> Dict(Int, Worksheet) {
  case col_idx < 0 {
    True -> worksheet_map
    False -> {
      let column = extract_column(temp_array, col_idx)
      let value_str =
        column
        |> string.join("")
        |> string.trim

      case value_str {
        "" ->
          process_columns(temp_array, col_idx - 1, map_idx - 1, worksheet_map)
        _ -> {
          case int.parse(value_str) {
            Error(_) ->
              process_columns(temp_array, col_idx - 1, map_idx, worksheet_map)
            Ok(value) -> {
          let updated_map = case dict.get(worksheet_map, map_idx) {
            Ok(ws) ->
              dict.insert(
                worksheet_map,
                map_idx,
                Worksheet(
                  values: ws.values,
                  cephalopod_values: list.append(ws.cephalopod_values, [value]),
                  operator: ws.operator,
                ),
              )
            Error(_) ->
              dict.insert(
                worksheet_map,
                map_idx,
                Worksheet(values: [], cephalopod_values: [value], operator: ""),
              )
          }
          process_columns(temp_array, col_idx - 1, map_idx, updated_map)
            }
          }
        }
      }
    }
  }
}

fn extract_column(array: List(List(String)), col_idx: Int) -> List(String) {
  list.map(array, fn(row) {
    case list_get(row, col_idx) {
      Ok(val) -> val
      Error(_) -> ""
    }
  })
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
    [first, ..rest] ->
      case index == 0 {
        True -> Ok(first)
        False -> list_get_helper(rest, index - 1)
      }
  }
}

fn calculate_worksheets(worksheets: Dict(Int, Worksheet)) -> #(Int, Int) {
  let part1_total =
    dict.fold(worksheets, 0, fn(acc, _key, ws) {
      acc + calculate_worksheet(ws.values, ws.operator)
    })

  let part2_total =
    dict.fold(worksheets, 0, fn(acc, _key, ws) {
      acc + calculate_worksheet(ws.cephalopod_values, ws.operator)
    })

  #(part1_total, part2_total)
}

fn calculate_worksheet(values: List(Int), operator: String) -> Int {
  case operator {
    "+" -> list.fold(values, 0, fn(acc, val) { acc + val })
    "*" ->
      list.fold(values, 1, fn(acc, val) {
        case acc == 0 {
          True -> val
          False -> acc * val
        }
      })
    _ -> 0
  }
}
