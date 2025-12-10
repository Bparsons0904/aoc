import time
import math
from utilities import read_file


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


class Interval:
    def __init__(self, start, end):
        self.start = start
        self.end = end

    def __repr__(self):
        return f"Interval({self.start}, {self.end})"


class VerticalEdge:
    def __init__(self, x, min_y, max_y):
        self.x = x
        self.min_y = min_y
        self.max_y = max_y


class HorizontalEdge:
    def __init__(self, y, min_x, max_x):
        self.y = y
        self.min_x = min_x
        self.max_x = max_x


class Candidate:
    def __init__(self, min_x, max_x, min_y, max_y, area):
        self.min_x = min_x
        self.max_x = max_x
        self.min_y = min_y
        self.max_y = max_y
        self.area = area


class SegmentTree:
    def __init__(self, max_width_at_y):
        self.y_coords = sorted(max_width_at_y.keys())
        self.y_index = {y: i for i, y in enumerate(self.y_coords)}
        self.size = len(self.y_coords)
        self.tree = [0] * (4 * self.size)

        if self.size > 0:
            self._build(0, 0, self.size - 1, max_width_at_y)

    def _build(self, arr_index, start, end, max_width_at_y):
        if start == end:
            self.tree[arr_index] = max_width_at_y[self.y_coords[start]]
            return
        mid = (start + end) // 2
        self._build(2 * arr_index + 1, start, mid, max_width_at_y)
        self._build(2 * arr_index + 2, mid + 1, end, max_width_at_y)
        self.tree[arr_index] = min(self.tree[2 * arr_index + 1], self.tree[2 * arr_index + 2])

    def query(self, min_y, max_y):
        if self.size == 0:
            return 0

        if min_y in self.y_index:
            l = self.y_index[min_y]
        else:
            import bisect
            idx = bisect.bisect_left(self.y_coords, min_y)
            if idx >= len(self.y_coords):
                return 0
            l = idx

        if max_y in self.y_index:
            r = self.y_index[max_y]
        else:
            import bisect
            idx = bisect.bisect_left(self.y_coords, max_y)
            if idx == 0:
                return 0
            r = idx - 1

        if l > r:
            if min_y in self.y_index and min_y == max_y:
                r = l
            else:
                return 0

        return self._query(0, 0, self.size - 1, l, r)

    def _query(self, arr_index, start, end, q_start, q_end):
        if q_start > end or q_end < start:
            return float('inf')
        if q_start <= start and q_end >= end:
            return self.tree[arr_index]
        mid = (start + end) // 2
        left_query = self._query(2 * arr_index + 1, start, mid, q_start, q_end)
        right_query = self._query(2 * arr_index + 2, mid + 1, end, q_start, q_end)
        return min(left_query, right_query)


def get_red_tiles(filename):
    lines = read_file(filename)
    red_tiles = []

    for row in lines:
        coordinates = row.split(",")
        x = int(coordinates[0])
        y = int(coordinates[1])
        red_tiles.append(Point(x, y))

    return red_tiles


def get_largest_area(red_tiles):
    result = 0.0

    for point1 in red_tiles:
        for point2 in red_tiles:
            width = abs(point1.x - point2.x) + 1
            height = abs(point1.y - point2.y) + 1
            area = width * height
            if area > result:
                result = area

    return int(result)


def merge_intervals(intervals):
    if len(intervals) == 0:
        return []

    intervals.sort(key=lambda i: i.start)
    merged = [intervals[0]]

    for i in range(1, len(intervals)):
        last = merged[-1]
        curr = intervals[i]

        if curr.start <= last.end + 1:
            if curr.end > last.end:
                last.end = curr.end
        else:
            merged.append(curr)

    return merged


def get_inside_intervals(red_tiles):
    vertical_edges = []
    horizontal_edges = []

    for i in range(len(red_tiles)):
        current = red_tiles[i]
        next_tile = red_tiles[(i + 1) % len(red_tiles)]

        if current.x == next_tile.x:
            min_y = min(current.y, next_tile.y)
            max_y = max(current.y, next_tile.y)
            vertical_edges.append(VerticalEdge(current.x, min_y, max_y))
        else:
            min_x = min(current.x, next_tile.x)
            max_x = max(current.x, next_tile.x)
            horizontal_edges.append(HorizontalEdge(current.y, min_x, max_x))

    min_y = min(p.y for p in red_tiles)
    max_y = max(p.y for p in red_tiles)

    intervals = {}

    for y in range(min_y, max_y + 1):
        crossings = []

        for edge in vertical_edges:
            if y >= edge.min_y and y < edge.max_y:
                crossings.append(edge.x)

        crossings.sort()

        row_intervals = []
        for i in range(0, len(crossings) - 1, 2):
            row_intervals.append(Interval(crossings[i], crossings[i + 1]))

        for edge in horizontal_edges:
            if edge.y == y:
                row_intervals.append(Interval(edge.min_x, edge.max_x))

        if len(row_intervals) > 0:
            intervals[y] = merge_intervals(row_intervals)

    return intervals


def is_rectangle_inside(min_x, max_x, min_y, max_y, intervals):
    for y in range(min_y, max_y + 1):
        if y not in intervals:
            return False

        row_intervals = intervals[y]

        lo, hi = 0, len(row_intervals)
        while lo < hi:
            mid = (lo + hi) // 2
            if row_intervals[mid].start <= min_x:
                lo = mid + 1
            else:
                hi = mid

        if lo == 0 or row_intervals[lo - 1].end < max_x:
            return False

    return True


def get_largest_area_with_intervals(red_tiles, intervals):
    max_width_at_y = {}
    for y, ivs in intervals.items():
        max_w = 0
        for iv in ivs:
            w = iv.end - iv.start + 1
            if w > max_w:
                max_w = w
        max_width_at_y[y] = max_w

    candidates = []
    for i in range(len(red_tiles)):
        for j in range(i + 1, len(red_tiles)):
            p1, p2 = red_tiles[i], red_tiles[j]
            min_x = min(p1.x, p2.x)
            max_x = max(p1.x, p2.x)
            min_y = min(p1.y, p2.y)
            max_y = max(p1.y, p2.y)
            width = max_x - min_x + 1
            height = max_y - min_y + 1
            candidates.append(Candidate(min_x, max_x, min_y, max_y, width * height))

    candidates.sort(key=lambda c: c.area, reverse=True)

    for c in candidates:
        width = c.max_x - c.min_x + 1

        can_fit = True
        for y in range(c.min_y, c.max_y + 1):
            if y not in max_width_at_y or max_width_at_y[y] < width:
                can_fit = False
                break

        if not can_fit:
            continue

        if is_rectangle_inside(c.min_x, c.max_x, c.min_y, c.max_y, intervals):
            return c.area

    return 0


def get_largest_area_with_intervals_optimized(red_tiles, intervals):
    max_width_at_y = {}
    for y, ivs in intervals.items():
        max_w = 0
        for iv in ivs:
            w = iv.end - iv.start + 1
            if w > max_w:
                max_w = w
        max_width_at_y[y] = max_w

    seg_tree = SegmentTree(max_width_at_y)

    candidates = []
    for i in range(len(red_tiles)):
        for j in range(i + 1, len(red_tiles)):
            p1, p2 = red_tiles[i], red_tiles[j]
            min_x = min(p1.x, p2.x)
            max_x = max(p1.x, p2.x)
            min_y = min(p1.y, p2.y)
            max_y = max(p1.y, p2.y)
            width = max_x - min_x + 1

            if p1.y in max_width_at_y and max_width_at_y[p1.y] < width:
                continue
            if p2.y in max_width_at_y and max_width_at_y[p2.y] < width:
                continue

            height = max_y - min_y + 1
            candidates.append(Candidate(min_x, max_x, min_y, max_y, width * height))

    candidates.sort(key=lambda c: c.area, reverse=True)

    for c in candidates:
        width = c.max_x - c.min_x + 1

        min_width_in_range = seg_tree.query(c.min_y, c.max_y)

        if min_width_in_range < width:
            continue

        if is_rectangle_inside(c.min_x, c.max_x, c.min_y, c.max_y, intervals):
            return c.area

    return 0


def day9():
    filename = "day9.part1"
    red_tiles = get_red_tiles(filename)

    start = time.time()
    largest_area_part1 = get_largest_area(red_tiles)
    part1_time = time.time() - start

    start = time.time()
    intervals = get_inside_intervals(red_tiles)
    interval_time = time.time() - start

    start = time.time()
    largest_area_part2 = get_largest_area_with_intervals(red_tiles, intervals)
    part2_time = time.time() - start

    start = time.time()
    largest_area_part2_optimized = get_largest_area_with_intervals_optimized(
        red_tiles, intervals
    )
    part2_optimized_time = time.time() - start

    print(f"Day 9:")
    print(f"  Part 1: {largest_area_part1} ({part1_time:.4f}s)")
    print(f"  Intervals: ({interval_time:.4f}s)")
    print(f"  Part 2: {largest_area_part2} ({part2_time:.4f}s)")
    print(f"  Part 2 Optimized: {largest_area_part2_optimized} ({part2_optimized_time:.4f}s)")


if __name__ == "__main__":
    day9()
