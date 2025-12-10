use std::fs::File;
use std::io::{self, BufRead};
use std::collections::HashMap;
use std::time::Instant;

#[derive(Debug, Copy, Clone, Eq, PartialEq, Hash)]
struct JunctionBox {
    x: i64,
    y: i64,
    z: i64,
}


#[derive(Debug, Clone)]
struct ShortestConnection {
    distance: i64,
    from: JunctionBox,
    to: JunctionBox,
}

struct DSU {
    parent: Vec<usize>,
    size: Vec<usize>,
}

impl DSU {
    fn new(n: usize) -> Self {
        DSU {
            parent: (0..n).collect(),
            size: vec![1; n],
        }
    }

    fn find(&mut self, i: usize) -> usize {
        if self.parent[i] == i {
            return i;
        }
        self.parent[i] = self.find(self.parent[i]);
        self.parent[i]
    }

    fn union(&mut self, i: usize, j: usize) -> bool {
        let root_i = self.find(i);
        let root_j = self.find(j);
        if root_i != root_j {
            if self.size[root_i] < self.size[root_j] {
                self.parent[root_i] = root_j;
                self.size[root_j] += self.size[root_i];
            } else {
                self.parent[root_j] = root_i;
                self.size[root_i] += self.size[root_j];
            }
            return true;
        }
        false
    }
}

fn main() {
    let (junction_boxes, box_to_index) = build_junction_boxes("files/day8.part1").unwrap();
    let connections = get_sorted_connections(&junction_boxes);

    let part1_results = build_junction_box_connections_part1(&connections, &box_to_index, 1000);
    println!("Day 8, Part 1: {}", part1_results);

    let now = Instant::now();
    let part2_results = build_junction_box_connections_part2(&connections, &box_to_index);
    let elapsed = now.elapsed();
    println!("Day 8, Part 2: {}", part2_results);
    println!("Part 2 took: {:.2?}", elapsed);
}

fn build_junction_box_connections_part1(
    shortest_connections: &[ShortestConnection],
    box_to_index: &HashMap<JunctionBox, usize>,
    limit: usize,
) -> u64 {
    let mut dsu = DSU::new(box_to_index.len());
    for connection in shortest_connections.iter().take(limit) {
        let from_idx = box_to_index[&connection.from];
        let to_idx = box_to_index[&connection.to];
        dsu.union(from_idx, to_idx);
    }

    let mut component_sizes = HashMap::new();
    for i in 0..box_to_index.len() {
        let root = dsu.find(i);
        *component_sizes.entry(root).or_insert(0) += 1;
    }
    
    let mut sizes: Vec<u64> = component_sizes.values().map(|&v| v as u64).collect();
    sizes.sort_by(|a, b| b.cmp(a));
    
    sizes.iter().take(3).product()
}

fn build_junction_box_connections_part2(
    shortest_connections: &[ShortestConnection],
    box_to_index: &HashMap<JunctionBox, usize>,
) -> i64 {
    let mut dsu = DSU::new(box_to_index.len());
    let mut num_components = box_to_index.len();

    for connection in shortest_connections {
        let from_idx = box_to_index[&connection.from];
        let to_idx = box_to_index[&connection.to];
        if dsu.union(from_idx, to_idx) {
            num_components -= 1;
            if num_components == 1 {
                return connection.from.x * connection.to.x;
            }
        }
    }
    0
}

fn get_sorted_connections(junction_boxes: &[JunctionBox]) -> Vec<ShortestConnection> {
    let mut shortest_connections = Vec::new();
    for i in 0..junction_boxes.len() {
        for k in i + 1..junction_boxes.len() {
            let jb1 = junction_boxes[i];
            let jb2 = junction_boxes[k];
            let dx = jb1.x as f64 - jb2.x as f64;
            let dy = jb1.y as f64 - jb2.y as f64;
            let dz = jb1.z as f64 - jb2.z as f64;
            let distance = (dx * dx + dy * dy + dz * dz).sqrt() as i64;
            shortest_connections.push(ShortestConnection {
                distance,
                from: jb1,
                to: jb2,
            });
        }
    }
    shortest_connections.sort_by_key(|c| c.distance);
    shortest_connections
}

fn build_junction_boxes(filename: &str) -> io::Result<(Vec<JunctionBox>, HashMap<JunctionBox, usize>)> {
    let file = File::open(filename)?;
    let reader = io::BufReader::new(file);
    let mut junction_boxes = Vec::new();
    let mut box_to_index = HashMap::new();

    for (i, line) in reader.lines().enumerate() {
        let line = line?;
        let coords: Vec<i64> = line.split(',').map(|s| s.parse().unwrap()).collect();
        let jb = JunctionBox { x: coords[0], y: coords[1], z: coords[2] };
        junction_boxes.push(jb);
        box_to_index.insert(jb, i);
    }
    Ok((junction_boxes, box_to_index))
}