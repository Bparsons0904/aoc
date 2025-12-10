import time
from utilities import read_file


class IngredientRange:
    def __init__(self, min_val, max_val):
        self.min = min_val
        self.max = max_val


class Ingredients:
    def __init__(self):
        self.fresh_ingredient_ranges = []
        self.ingredients = []

    def count_fresh_ingredients(self):
        count_of_fresh_ingredients = 0
        for ingredient in self.ingredients:
            for fresh_ingredient_range in self.fresh_ingredient_ranges:
                if (fresh_ingredient_range.min <= ingredient <=
                    fresh_ingredient_range.max):
                    count_of_fresh_ingredients += 1
                    break

        count_of_potential_fresh_ingredients = 0
        for fresh_ingredient_range in self.fresh_ingredient_ranges:
            count_of_potential_fresh_ingredients += (
                fresh_ingredient_range.max - fresh_ingredient_range.min + 1
            )

        return count_of_fresh_ingredients, count_of_potential_fresh_ingredients

    def parse_ingredients(self, lines):
        ingredient_ranges = []
        is_first_half = True

        for row in lines:
            if row == "":
                is_first_half = False
                continue

            if is_first_half:
                parts = row.split("-")
                min_val = int(parts[0])
                max_val = int(parts[1])
                ingredient_range = IngredientRange(min_val, max_val)
                ingredient_ranges.append(ingredient_range)
            else:
                ingredient_id = int(row)
                self.ingredients.append(ingredient_id)

        ingredient_ranges.sort(key=lambda r: r.min)

        for ingredient_range in ingredient_ranges:
            if len(self.fresh_ingredient_ranges) == 0:
                self.fresh_ingredient_ranges.append(ingredient_range)
                continue

            found = False
            for i, fresh_ingredient_range in enumerate(self.fresh_ingredient_ranges):
                if ((ingredient_range.min >= fresh_ingredient_range.min and
                     ingredient_range.min <= fresh_ingredient_range.max) or
                    (ingredient_range.max >= fresh_ingredient_range.min and
                     ingredient_range.max <= fresh_ingredient_range.max)):
                    self.fresh_ingredient_ranges[i].min = min(
                        self.fresh_ingredient_ranges[i].min,
                        ingredient_range.min
                    )
                    self.fresh_ingredient_ranges[i].max = max(
                        self.fresh_ingredient_ranges[i].max,
                        ingredient_range.max
                    )
                    found = True

            if not found:
                self.fresh_ingredient_ranges.append(ingredient_range)


def day5():
    lines = read_file("day5.part1")

    ingredients = Ingredients()
    ingredients.parse_ingredients(lines)

    start = time.time()
    part1_count, part2_count = ingredients.count_fresh_ingredients()
    elapsed = time.time() - start

    print(f"Day 5:")
    print(f"  Part 1: {part1_count}, Part 2: {part2_count} ({elapsed:.4f}s)")


if __name__ == "__main__":
    day5()
