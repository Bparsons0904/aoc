import time
from utilities import read_file


class BatteryPack:
    def __init__(self, batteries):
        self.batteries = batteries

    def get_largest_pack_joltage(self, battery_size):
        max_joltage = []

        index = 0
        iteration = 0
        while True:
            if len(max_joltage) == battery_size:
                break
            max_found_joltage, end_index = self.iterate_battery_pack_check(
                index, iteration, battery_size
            )
            max_joltage.append(max_found_joltage)
            index = end_index + 1
            iteration += 1

        result = int(''.join(str(j) for j in max_joltage))
        return result

    def iterate_battery_pack_check(self, starting_index, current_battery_size, battery_size):
        max_found_joltage = 0
        end_index = starting_index
        stop = len(self.batteries) + current_battery_size - (battery_size - 1)

        for i in range(starting_index, stop):
            if self.batteries[i] > max_found_joltage:
                max_found_joltage = self.batteries[i]
                end_index = i

        return max_found_joltage, end_index


def process_day3_file():
    lines = read_file("day3.part1")
    battery_packs = []

    for line in lines:
        batteries = [int(char) for char in line]
        battery_packs.append(BatteryPack(batteries))

    return battery_packs


def calculate_12_pack_joltage(battery_packs, battery_size):
    max_joltage = 0
    for battery_pack in battery_packs:
        max_joltage += battery_pack.get_largest_pack_joltage(battery_size)

    return max_joltage


def day3():
    battery_packs = process_day3_file()

    start = time.time()
    part1_result = calculate_12_pack_joltage(battery_packs, 2)
    part1_time = time.time() - start

    start = time.time()
    part2_result = calculate_12_pack_joltage(battery_packs, 12)
    part2_time = time.time() - start

    print(f"Day 3:")
    print(f"  Part 1: {part1_result} ({part1_time:.4f}s)")
    print(f"  Part 2: {part2_result} ({part2_time:.4f}s)")


if __name__ == "__main__":
    day3()
