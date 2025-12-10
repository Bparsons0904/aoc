use std::fs::File;
use std::io::{self, BufRead};
use std::time::Instant;

fn main() {
    let battery_packs = process_day3_file("files/day3.part1").unwrap();

    let part1_result = calculate_12_pack_joltage(&battery_packs, 2);
    println!("Day 3, Part 1: {}", part1_result);

    let now = Instant::now();
    let part2_result = calculate_12_pack_joltage(&battery_packs, 12);
    let elapsed = now.elapsed();
    println!("Day 3, Part 2: {}", part2_result);
    println!("Part 2 took: {:.2?}", elapsed);
}

fn calculate_12_pack_joltage(battery_packs: &[Vec<i32>], battery_size: usize) -> i32 {
    let mut max_joltage = 0;
    for battery_pack in battery_packs {
        max_joltage += get_largest_pack_joltage(battery_pack, battery_size);
    }
    max_joltage
}

fn get_largest_pack_joltage(battery_pack: &[i32], battery_size: usize) -> i32 {
    if battery_pack.len() < battery_size {
        return 0;
    }

    let mut max_joltage_digits = Vec::with_capacity(battery_size);
    let mut current_index = 0;

    for i in 0..battery_size {
        let remaining_needed = battery_size - i;
        let window_end = battery_pack.len() - remaining_needed + 1;
        let window = &battery_pack[current_index..window_end];
        
        if let Some((max_digit_index, &max_digit)) = window.iter().enumerate().max_by_key(|&(_, &val)| val) {
            max_joltage_digits.push(max_digit);
            current_index += max_digit_index + 1;
        }
    }

    let result_str: String = max_joltage_digits.iter().map(|&d| d.to_string()).collect();
    result_str.parse::<i32>().unwrap_or(0)
}

fn process_day3_file(filename: &str) -> io::Result<Vec<Vec<i32>>> {
    let file = File::open(filename)?;
    let reader = io::BufReader::new(file);
    let mut battery_packs = Vec::new();

    for line in reader.lines() {
        let line = line?;
        let battery_pack: Vec<i32> = line
            .chars()
            .map(|c| c.to_digit(10).unwrap_or(0) as i32)
            .collect();
        battery_packs.push(battery_pack);
    }

    Ok(battery_packs)
}