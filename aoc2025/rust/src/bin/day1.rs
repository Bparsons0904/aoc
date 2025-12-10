use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let instructions = read_instructions("files/day1.part1").unwrap();

    let mut step1_count = 0;
    let mut step2_count = 0;
    let mut current_value = 50;

    for instruction in instructions {
        match instruction.direction.as_str() {
            "R" => {
                let new_value = current_value + instruction.step;
                if new_value >= 100 {
                    if current_value == 0 {
                        step2_count += instruction.step / 100;
                    } else {
                        step2_count += (instruction.step + current_value) / 100;
                    }
                }
                current_value = new_value % 100;
            }
            "L" => {
                let new_value = current_value - instruction.step;
                if current_value == 0 {
                    step2_count += instruction.step / 100;
                } else if instruction.step >= current_value {
                    step2_count += (instruction.step - current_value) / 100 + 1;
                }
                current_value = ((new_value % 100) + 100) % 100;
            }
            _ => {}
        }
        if current_value == 0 {
            step1_count += 1;
        }
    }

    println!("Day 1, Part 1: {}", step1_count);
    println!("Day 1, Part 2: {}", step2_count);
}

struct Instruction {
    direction: String,
    step: i32,
}

fn read_instructions(filename: &str) -> io::Result<Vec<Instruction>> {
    let path = Path::new(filename);
    let file = File::open(path)?;
    let reader = io::BufReader::new(file);

    let mut instructions = Vec::new();
    for line in reader.lines() {
        let line = line?;
        let direction = line.chars().next().unwrap().to_string();
        let step = line[1..].parse::<i32>().unwrap();
        instructions.push(Instruction { direction, step });
    }

    Ok(instructions)
}