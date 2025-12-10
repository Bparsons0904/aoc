use std::fs::File;
use std::io::{self, BufRead};
use std::time::Instant;

const PAPER_ROLL: char = '@';
const EMPTY: char = '.';


#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Point {
    x: i32,
    y: i32,
}

struct Grid {
    width: i32,
    height: i32,
    map: Vec<Vec<char>>,
}

impl Grid {
    fn new(filename: &str) -> io::Result<Self> {
        let file = File::open(filename)?;
        let reader = io::BufReader::new(file);
        let map: Vec<Vec<char>> = reader
            .lines()
            .map(|line| line.unwrap().chars().collect())
            .collect();
        
        let height = map.len() as i32;
        let width = if height > 0 { map[0].len() as i32 } else { 0 };

        Ok(Grid { width, height, map })
    }

    fn position_contains_object(&self, point: Point, object: char) -> bool {
        if point.x >= 0 && point.x < self.width && point.y >= 0 && point.y < self.height {
            self.map[point.y as usize][point.x as usize] == object
        } else {
            false
        }
    }

    fn set_object(&mut self, point: Point, object: char) {
        if point.x >= 0 && point.x < self.width && point.y >= 0 && point.y < self.height {
            self.map[point.y as usize][point.x as usize] = object;
        }
    }
}

fn main() {
    let mut grid = Grid::new("files/day4.part1").unwrap();
    let now = Instant::now();
    let (part1_count, part2_count) = calculate_paper_rolls_queue(&mut grid);
    let elapsed = now.elapsed();

    println!("Day 4, Part 1: {}", part1_count);
    println!("Day 4, Part 2: {}", part2_count);
    println!("Part 1 & 2 took: {:.2?}", elapsed);
}

fn calculate_paper_rolls_queue(grid: &mut Grid) -> (i32, i32) {
    let directions = [
        Point { x: -1, y: -1 }, Point { x: 0, y: -1 }, Point { x: 1, y: -1 },
        Point { x: -1, y: 0 },                         Point { x: 1, y: 0 },
        Point { x: -1, y: 1 }, Point { x: 0, y: 1 }, Point { x: 1, y: 1 },
    ];

    let mut part1_count = 0;
    let mut stack = Vec::new();

    for y in 0..grid.height {
        for x in 0..grid.width {
            if grid.map[y as usize][x as usize] == PAPER_ROLL {
                let point = Point { x, y };
                let connected_rolls = count_paper_roll_contacts(grid, point, &directions);
                if connected_rolls < 4 {
                    part1_count += 1;
                    stack.push(point);
                }
            }
        }
    }

    let mut part2_count = 0;
    while let Some(point) = stack.pop() {
        if grid.map[point.y as usize][point.x as usize] != PAPER_ROLL {
            continue;
        }

        grid.set_object(point, EMPTY);
        part2_count += 1;

        for &dir in &directions {
            let neighbor = Point { x: point.x + dir.x, y: point.y + dir.y };
            if grid.position_contains_object(neighbor, PAPER_ROLL) {
                let connected_rolls = count_paper_roll_contacts(grid, neighbor, &directions);
                if connected_rolls < 4 {
                    stack.push(neighbor);
                }
            }
        }
    }

    (part1_count, part2_count)
}

fn count_paper_roll_contacts(grid: &Grid, point: Point, directions: &[Point]) -> i32 {
    let mut count = 0;
    for &direction in directions {
        let neighbor = Point { x: point.x + direction.x, y: point.y + direction.y };
        if grid.position_contains_object(neighbor, PAPER_ROLL) {
            count += 1;
        }
    }
    count
}