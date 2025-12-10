import time
from utilities import read_file


TACHYON = '^'


class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def __hash__(self):
        return hash((self.x, self.y))

    def __repr__(self):
        return f"Point({self.x}, {self.y})"


class Grid:
    def __init__(self, filename):
        lines = read_file(filename)
        self.height = len(lines)
        self.width = len(lines[0]) if lines else 0
        self.map = [list(line) for line in lines]
        self.current = None
        self.start = None

        for i, row in enumerate(self.map):
            for j, char in enumerate(row):
                if char == 'S':
                    self.start = Point(j, i)
                    self.current = Point(j, i)


def locate_tachyon(tachyon_graph, tachyon_grid, point):
    for i in range(point.y, len(tachyon_grid.map)):
        key = Point(point.x, i)
        if key in tachyon_graph:
            return tachyon_graph[key]
    return 1


def process_tachyon_beam_routes_counter(tachyon_grid):
    tachyon_graph = {}

    row_length = len(tachyon_grid.map[0])
    for i in range(len(tachyon_grid.map) - 1, -1, -1):
        for j in range(row_length):
            if tachyon_grid.map[i][j] == TACHYON:
                tachyon_path_count = 0

                if j - 1 >= 0:
                    tachyon_path_count = locate_tachyon(
                        tachyon_graph,
                        tachyon_grid,
                        Point(j - 1, i)
                    )

                if j + 1 < row_length:
                    tachyon_path_count += locate_tachyon(
                        tachyon_graph,
                        tachyon_grid,
                        Point(j + 1, i)
                    )

                tachyon_graph[Point(j, i)] = tachyon_path_count

    total = locate_tachyon(
        tachyon_graph,
        tachyon_grid,
        Point(tachyon_grid.current.x, tachyon_grid.current.y)
    )

    return total


def process_tachyon_beam_split_counter(tachyon_grid):
    tachyon_split_counter = 0
    tachyon_current_lines = {tachyon_grid.current.x: True}

    for row in tachyon_grid.map:
        new_tachyon_current_lines = {}
        for x, space in enumerate(row):
            if space == TACHYON and x in tachyon_current_lines:
                tachyon_split_counter += 1
                new_tachyon_current_lines[x - 1] = True
                new_tachyon_current_lines[x + 1] = True
                if x in tachyon_current_lines:
                    del tachyon_current_lines[x]

        for x in new_tachyon_current_lines:
            if 0 <= x < len(tachyon_grid.map[0]):
                tachyon_current_lines[x] = True

    return tachyon_split_counter


def day7():
    tachyon_grid = Grid("day7.part1")

    start = time.time()
    part1_count = process_tachyon_beam_split_counter(tachyon_grid)
    part1_time = time.time() - start

    tachyon_grid = Grid("day7.part1")
    start = time.time()
    part2_count = process_tachyon_beam_routes_counter(tachyon_grid)
    part2_time = time.time() - start

    print(f"Day 7:")
    print(f"  Part 1: {part1_count} ({part1_time:.4f}s)")
    print(f"  Part 2: {part2_count} ({part2_time:.4f}s)")


if __name__ == "__main__":
    day7()
