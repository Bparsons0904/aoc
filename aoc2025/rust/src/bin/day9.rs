use std::cmp::{min, max};
use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::time::Instant;

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct Point {
    x: i64,
    y: i64,
}

#[derive(Debug, Copy, Clone, Eq, PartialEq, Ord, PartialOrd)]
struct Interval {
    start: i64,
    end: i64,
}

fn main() {
    let red_tiles = get_red_tiles("files/day9.part1").unwrap();
    println!("Day 9, Part 1: {}", get_largest_area(&red_tiles));
    
    let now = Instant::now();
    let intervals = get_inside_intervals(&red_tiles);
    let part2_result = get_largest_area_with_intervals_optimized(&red_tiles, &intervals);
    let elapsed = now.elapsed();
    println!("Day 9, Part 2: {}", part2_result);
    println!("Part 2 took: {:.2?}", elapsed);
}

fn get_largest_area_with_intervals_optimized(red_tiles: &[Point], intervals: &HashMap<i64, Vec<Interval>>) -> i64 {
    let mut max_width_at_y = HashMap::new();
    for (&y, ivs) in intervals {
        let max_w = ivs.iter().map(|iv| iv.end - iv.start + 1).max().unwrap_or(0);
        max_width_at_y.insert(y, max_w);
    }

    let seg_tree = SegmentTree::new(&max_width_at_y);

    let mut candidates = Vec::new();
    for i in 0..red_tiles.len() {
        for j in i + 1..red_tiles.len() {
            let p1 = red_tiles[i];
            let p2 = red_tiles[j];
            let min_x = min(p1.x, p2.x);
            let max_x = max(p1.x, p2.x);
            let min_y = min(p1.y, p2.y);
            let max_y = max(p1.y, p2.y);
            let area = (max_x - min_x + 1) * (max_y - min_y + 1);
            candidates.push((area, min_x, max_x, min_y, max_y));
        }
    }

    candidates.sort_by_key(|k| std::cmp::Reverse(k.0));

    for (area, min_x, max_x, min_y, max_y) in candidates {
        let width = max_x - min_x + 1;
        let min_width_in_range = seg_tree.query(min_y, max_y);

        if min_width_in_range < width {
            continue;
        }

        if is_rectangle_inside(min_x, max_x, min_y, max_y, intervals) {
            return area;
        }
    }

    0
}

fn get_red_tiles(filename: &str) -> io::Result<Vec<Point>> {
    let path = Path::new(filename);
    let file = File::open(path)?;
    let reader = io::BufReader::new(file);

    let mut points = Vec::new();
    for line in reader.lines() {
        let line = line?;
        let coords: Vec<&str> = line.split(',').collect();
        if coords.len() == 2 {
            let x = coords[0].parse::<i64>().unwrap();
            let y = coords[1].parse::<i64>().unwrap();
            points.push(Point { x, y });
        }
    }
    Ok(points)
}


fn get_largest_area(red_tiles: &[Point]) -> i64 {
    let mut result = 0.0;

    for &point1 in red_tiles {
        for &point2 in red_tiles {
            let width = (point1.x - point2.x).abs() as f64 + 1.0;
            let height = (point1.y - point2.y).abs() as f64 + 1.0;
            let area = width * height;
            if area > result {
                result = area;
            }
        }
    }
    result as i64
}

fn merge_intervals(intervals: &mut Vec<Interval>) {
    if intervals.is_empty() {
        return;
    }
    intervals.sort_unstable();
    let mut merged = Vec::new();
    merged.push(intervals[0]);

    for &interval in &intervals[1..] {
        let last = merged.last_mut().unwrap();
        if interval.start <= last.end + 1 {
            last.end = max(last.end, interval.end);
        } else {
            merged.push(interval);
        }
    }
    *intervals = merged;
}

struct SegmentTree {
    tree: Vec<i64>,
    y_coords: Vec<i64>,
    size: usize,
}

impl SegmentTree {
    fn new(max_width_at_y: &HashMap<i64, i64>) -> Self {
        let mut y_coords: Vec<i64> = max_width_at_y.keys().cloned().collect();
        y_coords.sort_unstable();

        let size = y_coords.len();
        let mut tree = vec![0; 4 * size];

        if size > 0 {
            Self::build(&mut tree, &y_coords, max_width_at_y, 0, 0, size - 1);
        }

        SegmentTree { tree, y_coords, size }
    }

    fn build(tree: &mut Vec<i64>, y_coords: &[i64], max_width_at_y: &HashMap<i64, i64>, arr_index: usize, start: usize, end: usize) {
        if start == end {
            tree[arr_index] = *max_width_at_y.get(&y_coords[start]).unwrap_or(&0);
            return;
        }
        let mid = start + (end - start) / 2;
        Self::build(tree, y_coords, max_width_at_y, 2 * arr_index + 1, start, mid);
        Self::build(tree, y_coords, max_width_at_y, 2 * arr_index + 2, mid + 1, end);
        tree[arr_index] = min(tree[2 * arr_index + 1], tree[2 * arr_index + 2]);
    }

    fn query(&self, min_y: i64, max_y: i64) -> i64 {
        if self.size == 0 {
            return 0;
        }

        let l = match self.y_coords.binary_search(&min_y) {
            Ok(i) => i,
            Err(i) => i,
        };
        let r = match self.y_coords.binary_search(&max_y) {
            Ok(i) => i,
            Err(i) => if i > 0 { i - 1 } else { return 0 },
        };
        
        if l > r {
             return 0;
        }

        self.query_internal(0, 0, self.size - 1, l, r)
    }

    fn query_internal(&self, arr_index: usize, start: usize, end: usize, q_start: usize, q_end: usize) -> i64 {
        if q_start > end || q_end < start {
            return i64::MAX;
        }
        if q_start <= start && q_end >= end {
            return self.tree[arr_index];
        }
        let mid = start + (end - start) / 2;
        let left_query = self.query_internal(2 * arr_index + 1, start, mid, q_start, q_end);
        let right_query = self.query_internal(2 * arr_index + 2, mid + 1, end, q_start, q_end);
        min(left_query, right_query)
    }
}

fn get_inside_intervals(red_tiles: &[Point]) -> HashMap<i64, Vec<Interval>> {
    if red_tiles.is_empty() {
        return HashMap::new();
    }


    struct VerticalEdge { x: i64, min_y: i64, max_y: i64 }
    struct HorizontalEdge { y: i64, min_x: i64, max_x: i64 }

    let mut vertical_edges = Vec::new();
    let mut horizontal_edges = Vec::new();

    for i in 0..red_tiles.len() {
        let current = red_tiles[i];
        let next = red_tiles[(i + 1) % red_tiles.len()];

        if current.x == next.x {
            vertical_edges.push(VerticalEdge {
                x: current.x,
                min_y: min(current.y, next.y),
                max_y: max(current.y, next.y),
            });
        } else {
            horizontal_edges.push(HorizontalEdge {
                y: current.y,
                min_x: min(current.x, next.x),
                max_x: max(current.x, next.x),
            });
        }
    }

    let min_y = red_tiles.iter().map(|p| p.y).min().unwrap_or(0);
    let max_y = red_tiles.iter().map(|p| p.y).max().unwrap_or(0);

    let mut intervals = HashMap::new();
    for y in min_y..=max_y {
        let mut crossings = Vec::new();
        for edge in &vertical_edges {
            if y >= edge.min_y && y < edge.max_y {
                crossings.push(edge.x);
            }
        }
        crossings.sort_unstable();

        let mut row_intervals = Vec::new();
        for chunk in crossings.chunks_exact(2) {
            row_intervals.push(Interval { start: chunk[0], end: chunk[1] });
        }

        for edge in &horizontal_edges {
            if edge.y == y {
                row_intervals.push(Interval { start: edge.min_x, end: edge.max_x });
            }
        }

        if !row_intervals.is_empty() {
            merge_intervals(&mut row_intervals);
            intervals.insert(y, row_intervals);
        }
    }

    intervals
}

fn is_rectangle_inside(min_x: i64, max_x: i64, min_y: i64, max_y: i64, intervals: &HashMap<i64, Vec<Interval>>) -> bool {
    for y in min_y..=max_y {
        if let Some(row_intervals) = intervals.get(&y) {
            let mut lo = 0;
            let mut hi = row_intervals.len();
            while lo < hi {
                let mid = lo + (hi - lo) / 2;
                if row_intervals[mid].start <= min_x {
                    lo = mid + 1;
                } else {
                    hi = mid;
                }
            }

            if lo == 0 || row_intervals[lo - 1].end < max_x {
                return false;
            }
        } else {
            return false;
        }
    }
    true
}