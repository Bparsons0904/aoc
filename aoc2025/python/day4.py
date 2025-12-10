import time
from utilities import read_file


EMPTY = '.'
PAPER_ROLL = '@'


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


RIGHT = Point(1, 0)
LEFT = Point(-1, 0)
DOWN = Point(0, 1)
UP = Point(0, -1)
RIGHT_DOWN = Point(1, 1)
RIGHT_UP = Point(1, -1)
LEFT_DOWN = Point(-1, 1)
LEFT_UP = Point(-1, -1)

DIRECTIONS = [LEFT, RIGHT, UP, DOWN, LEFT_DOWN, LEFT_UP, RIGHT_DOWN, RIGHT_UP]


class Grid:
    def __init__(self, filename):
        lines = read_file(filename)
        self.height = len(lines)
        self.width = len(lines[0]) if lines else 0
        self.map = [list(line) for line in lines]

    def set_object(self, point, obj):
        self.map[point.y][point.x] = obj

    def position_contains_object(self, point, obj):
        if not self.point_within_bounds(point):
            return False
        return self.map[point.y][point.x] == obj

    def point_within_bounds(self, point):
        return 0 <= point.x < self.width and 0 <= point.y < self.height


def move_point(point, direction):
    return Point(point.x + direction.x, point.y + direction.y)


def count_paper_roll_contacts(grid, point):
    count = 0
    for direction in DIRECTIONS:
        if grid.position_contains_object(move_point(point, direction), PAPER_ROLL):
            count += 1
    return count


def calculate_paper_rolls_queue(grid):
    part1_count = 0
    part2_count = 0

    stack = []
    for y in range(grid.height):
        for x in range(grid.width):
            if grid.map[y][x] == PAPER_ROLL:
                point = Point(x, y)
                connected_rolls = count_paper_roll_contacts(grid, point)
                if connected_rolls < 4:
                    part1_count += 1
                    stack.append(point)

    while stack:
        point = stack.pop()

        if grid.map[point.y][point.x] != PAPER_ROLL:
            continue

        grid.set_object(point, EMPTY)
        part2_count += 1

        for direction in DIRECTIONS:
            neighbor = move_point(point, direction)
            if grid.position_contains_object(neighbor, PAPER_ROLL):
                connected_rolls = count_paper_roll_contacts(grid, neighbor)
                if connected_rolls < 4:
                    stack.append(neighbor)

    return part1_count, part2_count


def calculate_paper_rolls_optimized(grid):
    part1_count = 0
    part2_count = 0

    mapped_paper_rolls = {}
    for y in range(grid.height):
        for x in range(grid.width):
            if grid.map[y][x] == PAPER_ROLL:
                point = Point(x, y)
                mapped_paper_rolls[point] = True
                connected_rolls = count_paper_roll_contacts(grid, point)
                if connected_rolls < 4:
                    part1_count += 1

    last_pass_count = -1
    while last_pass_count != 0:
        last_pass_count = 0

        points_to_remove = []
        for point in mapped_paper_rolls:
            connected_rolls = count_paper_roll_contacts(grid, point)
            if connected_rolls < 4:
                grid.set_object(point, EMPTY)
                points_to_remove.append(point)
                last_pass_count += 1
                part2_count += 1

        for point in points_to_remove:
            del mapped_paper_rolls[point]

    return part1_count, part2_count


def calculate_paper_rolls(grid):
    part1_count = 0
    part2_count = 0

    first_pass = True
    last_pass_count = -1

    while last_pass_count != 0:
        last_pass_count = 0

        for y in range(grid.height):
            for x in range(grid.width):
                if grid.map[y][x] != PAPER_ROLL:
                    continue

                connected_rolls = count_paper_roll_contacts(grid, Point(x, y))
                if connected_rolls < 4:
                    if first_pass:
                        part1_count += 1
                    else:
                        grid.set_object(Point(x, y), EMPTY)
                        last_pass_count += 1
                        part2_count += 1

        if first_pass:
            first_pass = False
            last_pass_count = -1

    return part1_count, part2_count


def day4():
    grid = Grid("day4.part1")

    start = time.time()
    part1_count_queue, part2_count_queue = calculate_paper_rolls_queue(grid)
    queue_time = time.time() - start

    print(f"Day 4 Queue:")
    print(f"  part1: {part1_count_queue}, part2: {part2_count_queue} ({queue_time:.4f}s)")


if __name__ == "__main__":
    day4()
