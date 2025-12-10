use std::fs::File;
use std::io::{self, BufRead};
use std::collections::HashMap;
use std::time::Instant;

const TACHYON: char = '^';
const START: char = 'S';


#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Point {
    x: i32,
    y: i32,
}

struct Grid {
    width: i32,
    height: i32,
    start: Point,
    map: Vec<Vec<char>>,
}

impl Grid {
    fn new(filename: &str) -> io::Result<Self> {
        let file = File::open(filename)?;
        let reader = io::BufReader::new(file);
        let mut start = Point { x: 0, y: 0 };
        let map: Vec<Vec<char>> = reader
            .lines()
            .enumerate()
            .map(|(y, line)| {
                let l = line.unwrap();
                if let Some(x) = l.find(START) {
                    start = Point { x: x as i32, y: y as i32 };
                }
                l.chars().collect()
            })
            .collect();
        
        let height = map.len() as i32;
        let width = if height > 0 { map[0].len() as i32 } else { 0 };

        Ok(Grid { width, height, start, map })
    }
}

fn main() {
    let grid = Grid::new("files/day7.part1").unwrap();
    
    let part1_count = process_tachyon_beam_split_counter(&grid);
    println!("Day 7, Part 1: {}", part1_count);

    let now = Instant::now();
    let part2_count = process_tachyon_beam_routes_counter(&grid);
    let elapsed = now.elapsed();
    println!("Day 7, Part 2: {}", part2_count);
    println!("Part 2 took: {:.2?}", elapsed);
}

fn process_tachyon_beam_split_counter(grid: &Grid) -> i32 {
    let mut tachyon_split_counter = 0;
    let mut tachyon_current_lines: HashMap<i32, bool> = HashMap::new();
    tachyon_current_lines.insert(grid.start.x, true);

    for row in &grid.map {
        let mut new_tachyon_current_lines = HashMap::new();
        for (x, &space) in row.iter().enumerate() {
            let x = x as i32;
            if space == TACHYON && tachyon_current_lines.contains_key(&x) {
                tachyon_split_counter += 1;
                new_tachyon_current_lines.insert(x - 1, true);
                new_tachyon_current_lines.insert(x + 1, true);
                tachyon_current_lines.remove(&x);
            }
        }

        for (x, _) in new_tachyon_current_lines {
            if x >= 0 && x < grid.width {
                tachyon_current_lines.insert(x, true);
            }
        }
    }
    tachyon_split_counter
}

type TachyonGraph = HashMap<Point, usize>;

fn process_tachyon_beam_routes_counter(grid: &Grid) -> usize {
    let mut tachyon_graph = TachyonGraph::new();
    let row_length = grid.width as usize;

    for i in (0..grid.height as usize).rev() {
        for j in 0..row_length {
            if grid.map[i][j] == TACHYON {
                let mut tachyon_path_count = 0;
                if j > 0 {
                    tachyon_path_count += locate_tachyon(&tachyon_graph, grid, Point { x: j as i32 - 1, y: i as i32 });
                }
                if j + 1 < row_length {
                    tachyon_path_count += locate_tachyon(&tachyon_graph, grid, Point { x: j as i32 + 1, y: i as i32 });
                }
                tachyon_graph.insert(Point { x: j as i32, y: i as i32 }, tachyon_path_count);
            }
        }
    }

    locate_tachyon(&tachyon_graph, grid, grid.start)
}

fn locate_tachyon(tachyon_graph: &TachyonGraph, grid: &Grid, point: Point) -> usize {
    for i in point.y as usize + 1..grid.height as usize {
        if let Some(&value) = tachyon_graph.get(&Point { x: point.x, y: i as i32 }) {
            return value;
        }
    }
    1
}