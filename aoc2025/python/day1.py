import time
from utilities import read_file


class Dial:
    def __init__(self, value):
        self.value = value
        self.next = None
        self.prev = None


class DialList:
    def __init__(self):
        self.head = None
        self.tail = None

    def next(self):
        self.head = self.head.next

    def prev(self):
        self.head = self.head.prev

    def build_linked_list(self):
        prev_node = None
        for i in range(100):
            node = Dial(i)
            node.prev = prev_node

            if self.head is None:
                self.head = node

            if prev_node is not None:
                prev_node.next = node

            self.tail = node
            prev_node = node

        self.tail.next = self.head
        self.head.prev = self.tail


def parse_instructions(lines):
    instructions = []
    for line in lines:
        direction = line[0]
        step = int(line[1:])
        instructions.append((direction, step))
    return instructions


def day1():
    start_time = time.time()

    lines = read_file("day1.part1")
    instructions = parse_instructions(lines)

    dial_list = DialList()
    dial_list.build_linked_list()

    while dial_list.head.value != 50:
        dial_list.next()

    part1_count = 0
    part2_count = 0

    for direction, step in instructions:
        if direction == 'R':
            for _ in range(step):
                dial_list.next()
                if dial_list.head.value == 0:
                    part2_count += 1
        elif direction == 'L':
            for _ in range(step):
                dial_list.prev()
                if dial_list.head.value == 0:
                    part2_count += 1

        if dial_list.head.value == 0:
            part1_count += 1

    elapsed = time.time() - start_time
    print(f"Day 1: part1={part1_count}, part2={part2_count}, time={elapsed:.4f}s")


def day1_optimized():
    start_time = time.time()

    lines = read_file("day1.part1")
    instructions = parse_instructions(lines)

    part1_count = 0
    part2_count = 0
    current_value = 50

    for direction, step in instructions:
        if direction == 'R':
            new_value = current_value + step
            if new_value >= 100:
                if current_value == 0:
                    part2_count += step // 100
                else:
                    part2_count += (step + current_value) // 100
            current_value = new_value % 100
        elif direction == 'L':
            new_value = current_value - step
            if current_value == 0:
                part2_count += step // 100
            elif step >= current_value:
                part2_count += (step - current_value) // 100 + 1
            current_value = ((new_value % 100) + 100) % 100

        if current_value == 0:
            part1_count += 1

    elapsed = time.time() - start_time
    print(f"Day 1 (optimized): part1={part1_count}, part2={part2_count}, time={elapsed:.4f}s")


if __name__ == "__main__":
    day1()
    day1_optimized()
