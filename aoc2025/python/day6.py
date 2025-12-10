import time
from utilities import read_file


class Worksheet:
    def __init__(self, operator):
        self.values = []
        self.cephalopod_values = []
        self.operator = operator


def build_operators(lines):
    worksheet_map = {}
    operators = lines[-1].split()

    for i, operator_string in enumerate(operators):
        worksheet_map[i] = Worksheet(operator_string)

    return worksheet_map


def process_values(worksheet_map, lines):
    for row in lines:
        values = row.split()
        for j, value_string in enumerate(values):
            value_int = int(value_string)
            worksheet_map[j].values.append(value_int)


def process_cephalopod_values(worksheet_map, lines):
    temp_array = []
    for row in lines:
        temp_array.append(list(row))

    length = len(temp_array[0])
    map_index = len(worksheet_map) - 1

    for j in range(length - 1, -1, -1):
        column_chars = []
        for i in range(len(temp_array)):
            column_chars.append(temp_array[i][j])

        value = ''.join(column_chars).strip()
        if value == "":
            map_index -= 1
            continue

        value_int = int(value)
        worksheet_map[map_index].cephalopod_values.append(value_int)


def process_day6_file():
    lines = read_file("day6.part1")
    worksheet_map = build_operators(lines)
    lines = lines[:-1]
    process_values(worksheet_map, lines)
    process_cephalopod_values(worksheet_map, lines)

    return worksheet_map


def calculate_worksheets(worksheet_map):
    total = 0
    for ws in worksheet_map.values():
        subtotal = 0
        for value in ws.values:
            if ws.operator == "+":
                subtotal += value
            elif ws.operator == "*":
                if subtotal == 0:
                    subtotal = 1
                subtotal *= value
        total += subtotal

    cephalopod_total = 0
    for ws in worksheet_map.values():
        cephalopod_subtotal = 0
        for cephalopod_value in ws.cephalopod_values:
            if ws.operator == "+":
                cephalopod_subtotal += cephalopod_value
            elif ws.operator == "*":
                if cephalopod_subtotal == 0:
                    cephalopod_subtotal = 1
                cephalopod_subtotal *= cephalopod_value
        cephalopod_total += cephalopod_subtotal

    return total, cephalopod_total


def day6():
    worksheet = process_day6_file()

    start = time.time()
    part1_total, part2_total = calculate_worksheets(worksheet)
    elapsed = time.time() - start

    print(f"Day 6:")
    print(f"  Part 1: {part1_total}, Part 2: {part2_total} ({elapsed:.4f}s)")


if __name__ == "__main__":
    day6()
