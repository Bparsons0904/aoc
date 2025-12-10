use std::fs;
use std::io;
use std::time::Instant;
use itoa;

#[derive(Debug)]
struct ProductIdRange {
    min: i64,
    max: i64,
}

fn main() {
    let product_id_ranges = process_day2("files/day2.part1").unwrap();

    let part1_count = calculate_part1(&product_id_ranges);
    println!("Day 2, Part 1: {}", part1_count);

    let now = Instant::now();
    let part2_count = calculate_part2(&product_id_ranges);
    let elapsed = now.elapsed();
    println!("Day 2, Part 2: {}", part2_count);
    println!("Part 2 took: {:.2?}", elapsed);
}

fn calculate_part1(product_id_ranges: &[ProductIdRange]) -> i64 {
    let mut result = 0;
    for product_id_range in product_id_ranges {
        for i in product_id_range.min..=product_id_range.max {
            let s = i.to_string();
            if s.len() % 2 != 0 {
                continue;
            }
            let half = s.len() / 2;
            if &s[..half] == &s[half..] {
                result += i;
            }
        }
    }
    result
}

fn calculate_part2(product_id_ranges: &[ProductIdRange]) -> i64 {
    let mut result = 0;
    let mut buffer = itoa::Buffer::new();
    for product_id_range in product_id_ranges {
        for value in product_id_range.min..=product_id_range.max {
            let s = buffer.format(value);
            let n = s.len();
            if n < 2 {
                continue;
            }
            
            let mut concatenated_bytes = [0u8; 40];
            concatenated_bytes[..n].copy_from_slice(s.as_bytes());
            concatenated_bytes[n..2*n].copy_from_slice(s.as_bytes());
            let concatenated = std::str::from_utf8(&concatenated_bytes[..2*n]).unwrap();

            if let Some(period) = concatenated[1..].find(s) {
                if period + 1 < n {
                    result += value;
                }
            }
        }
    }
    result
}

fn process_day2(filename: &str) -> io::Result<Vec<ProductIdRange>> {
    let row = fs::read_to_string(filename)?;
    let ranges: Vec<&str> = row.trim().split(',').collect();
    let mut product_id_ranges = Vec::new();
    for value in ranges {
        let id_range: Vec<&str> = value.split('-').collect();
        if id_range.len() == 2 {
            let min = id_range[0].parse::<i64>().unwrap();
            let max = id_range[1].parse::<i64>().unwrap();
            product_id_ranges.push(ProductIdRange { min, max });
        }
    }
    Ok(product_id_ranges)
}