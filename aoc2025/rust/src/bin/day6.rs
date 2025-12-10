use std::fs::File;
use std::io::{self, BufRead};
use std::collections::HashMap;
use std::time::Instant;

#[derive(Debug, Clone)]
struct Worksheet {
    values: Vec<i64>,
    cephalopod_values: Vec<i64>,
    operator: String,
}


impl Worksheet {
    fn new(operator: String) -> Self {
        Worksheet {
            values: Vec::new(),
            cephalopod_values: Vec::new(),
            operator,
        }
    }
}

type WorksheetMap = HashMap<usize, Worksheet>;

fn main() {
    let worksheet = process_day6_file("files/day6.part1").unwrap();
    let now = Instant::now();
    let (part1_total, part2_total) = calculate_worksheets(&worksheet);
    let elapsed = now.elapsed();

    println!("Day 6, Part 1: {}", part1_total);
    println!("Day 6, Part 2: {}", part2_total);
    println!("Part 1 & 2 took: {:.2?}", elapsed);
}

fn calculate_worksheets(worksheet: &WorksheetMap) -> (i64, i64) {
    let mut total: i64 = 0;
    for i in 0..worksheet.len() {
        if let Some(ws) = worksheet.get(&i) {
            let mut subtotal: i64 = 0;
            for &value in &ws.values {
                match ws.operator.as_str() {
                    "+" => subtotal += value,
                    "*" => {
                        if subtotal == 0 { subtotal = 1; }
                        subtotal *= value;
                    }
                    _ => {}
                }
            }
            total += subtotal;
        }
    }

    let mut cephalopod_total: i64 = 0;
    for i in 0..worksheet.len() {
         if let Some(ws) = worksheet.get(&i) {
            let mut cephalopod_subtotal: i64 = 0;
            for &cephalopod_value in &ws.cephalopod_values {
                match ws.operator.as_str() {
                    "+" => cephalopod_subtotal += cephalopod_value,
                    "*" => {
                        if cephalopod_subtotal == 0 { cephalopod_subtotal = 1; }
                        cephalopod_subtotal *= cephalopod_value;
                    }
                    _ => {}
                }
            }
            cephalopod_total += cephalopod_subtotal;
        }
    }

    (total, cephalopod_total)
}


fn process_day6_file(filename: &str) -> io::Result<WorksheetMap> {
    let file = File::open(filename)?;
    let lines: Vec<String> = io::BufReader::new(file).lines().map(|l| l.unwrap()).collect();
    
    let mut worksheet_map = build_operators(&lines);
    process_values(&mut worksheet_map, &lines[..lines.len()-1]);
    process_cephalopod_values(&mut worksheet_map, &lines[..lines.len()-1]);

    Ok(worksheet_map)
}

fn process_cephalopod_values(ws: &mut WorksheetMap, file: &[String]) {
    let temp_array: Vec<Vec<char>> = file.iter().map(|row| row.chars().collect()).collect();
    let length = temp_array.iter().map(|row| row.len()).max().unwrap_or(0);
    let mut map_index = ws.len() - 1;

    for j in (0..length).rev() {
        let mut sb = String::new();
        for i in 0..temp_array.len() {
            if j < temp_array[i].len() {
                sb.push(temp_array[i][j]);
            } else {
                sb.push(' ');
            }
        }

        let value_str = sb.trim();
        if value_str.is_empty() {
            if map_index > 0 {
                map_index -= 1;
            }
            continue;
        }
        
        if let Ok(value_int) = value_str.parse::<i64>() {
            if let Some(temp_ws) = ws.get_mut(&map_index) {
                temp_ws.cephalopod_values.push(value_int);
            }
        }
    }
}


fn process_values(ws: &mut WorksheetMap, file: &[String]) {
    for row in file {
        for (j, value_string) in row.split_whitespace().enumerate() {
            if let Ok(value_int) = value_string.parse::<i64>() {
                if let Some(temp_ws) = ws.get_mut(&j) {
                    temp_ws.values.push(value_int);
                }
            }
        }
    }
}

fn build_operators(file: &[String]) -> WorksheetMap {
    let mut worksheet_map = WorksheetMap::new();
    if let Some(last_line) = file.last() {
        for (i, operator_string) in last_line.split_whitespace().enumerate() {
            worksheet_map.insert(i, Worksheet::new(operator_string.to_string()));
        }
    }
    worksheet_map
}