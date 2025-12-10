import time
import math
from utilities import read_file


class ProductIDRange:
    def __init__(self, min_val, max_val):
        self.min = min_val
        self.max = max_val


def process_day2():
    row = read_file("day2.part1")[0]
    ranges = row.split(",")
    product_id_ranges = []

    for value in ranges:
        id_range = value.split("-")
        min_val = int(id_range[0])
        max_val = int(id_range[1])
        product_id_ranges.append(ProductIDRange(min_val, max_val))

    return product_id_ranges


def calculate_part1(product_id_ranges):
    result = 0
    for product_id_range in product_id_ranges:
        for i in range(product_id_range.min, product_id_range.max + 1):
            i_string = str(i)
            if len(i_string) % 2 != 0:
                continue

            half = len(i_string) // 2
            if i_string[:half] == i_string[half:]:
                result += i

    return result


def calculate_part2(product_id_ranges):
    result = 0

    for product_id_range in product_id_ranges:
        for i in range(product_id_range.min, product_id_range.max + 1):
            i_string = str(i)
            for j in range(len(i_string) - 1):
                to_check = i_string[:j + 1]
                expected_count = math.ceil(len(i_string) / len(to_check))
                count = i_string.count(to_check)
                if count >= 2 and count == expected_count:
                    result += i
                    break

    return result


def calculate_part2_1(product_id_ranges):
    result = 0

    for product_id_range in product_id_ranges:
        for value in range(product_id_range.min, product_id_range.max + 1):
            i_string = str(value)
            for j in range(1, len(i_string) // 2 + 1):
                pattern = i_string[:j]
                if i_string[j:].replace(pattern, "") == "":
                    result += value
                    break

    return result


def calculate_part2_2(product_id_ranges):
    result = 0

    for product_id_range in product_id_ranges:
        for value in range(product_id_range.min, product_id_range.max + 1):
            i_string = str(value)
            str_len = len(i_string)

            for pattern_len in range(1, str_len // 2 + 1):
                if str_len % pattern_len != 0:
                    continue

                pattern = i_string[:pattern_len]
                is_repeating = True

                for i in range(pattern_len, str_len, pattern_len):
                    if i_string[i:i + pattern_len] != pattern:
                        is_repeating = False
                        break

                if is_repeating:
                    result += value
                    break

    return result


def calculate_part2_3(product_id_ranges):
    result = 0

    for product_id_range in product_id_ranges:
        for value in range(product_id_range.min, product_id_range.max + 1):
            s = str(value)
            n = len(s)
            if n < 2:
                continue

            period = (s + s)[1:].find(s) + 1
            if period > 0 and period < n:
                result += value

    return result


def day2():
    product_id_ranges = process_day2()

    start = time.time()
    part1_count = calculate_part1(product_id_ranges)
    part1_time = time.time() - start

    start = time.time()
    part2_count = calculate_part2(product_id_ranges)
    part2_time = time.time() - start

    start = time.time()
    part2_1_count = calculate_part2_1(product_id_ranges)
    part2_1_time = time.time() - start

    start = time.time()
    part2_2_count = calculate_part2_2(product_id_ranges)
    part2_2_time = time.time() - start

    start = time.time()
    part2_3_count = calculate_part2_3(product_id_ranges)
    part2_3_time = time.time() - start

    print(f"Day 2:")
    print(f"  part1={part1_count} ({part1_time:.4f}s)")
    print(f"  part2={part2_count} ({part2_time:.4f}s)")
    print(f"  part2_1={part2_1_count} ({part2_1_time:.4f}s)")
    print(f"  part2_2={part2_2_count} ({part2_2_time:.4f}s)")
    print(f"  part2_3={part2_3_count} ({part2_3_time:.4f}s)")


if __name__ == "__main__":
    day2()
