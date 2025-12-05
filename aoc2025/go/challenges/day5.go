package challenges

import (
	"log"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"aoc/utilities"

	logger "github.com/Bparsons0904/goLogger"
)

type IngredientRange struct {
	min int
	max int
}

type Ingredients struct {
	FreshIngredientRanges []IngredientRange
	Ingredients           []int
}

func Day5() {
	log := logger.New("Day 5")

	file := utilities.ReadFile("day5.part1")

	var ingredients Ingredients
	ingredients.parseIngredients(file)

	timer := log.Timer("Part 1 Timer")
	part1Count, part2Count := ingredients.countFreshIngredients()
	timer()

	slog.Info("Part 1", "Part 1", part1Count, "Part 2", part2Count)
}

func (ingredients *Ingredients) countFreshIngredients() (int, int) {
	countOfFreshIngredients := 0
	for _, ingredient := range ingredients.Ingredients {
		for _, freshIngredientRange := range ingredients.FreshIngredientRanges {
			if freshIngredientRange.min <= ingredient && freshIngredientRange.max >= ingredient {
				countOfFreshIngredients++
				break
			}
		}
	}

	countOfPotentialFreshIngredients := 0
	for _, freshIngredientRange := range ingredients.FreshIngredientRanges {
		countOfPotentialFreshIngredients += (freshIngredientRange.max - freshIngredientRange.min) + 1
	}
	return countOfFreshIngredients, countOfPotentialFreshIngredients
}

func (ingredients *Ingredients) parseIngredients(file []string) {
	var ingredientRanges []IngredientRange
	isFirstHalf := true
	for _, row := range file {
		if row == "" {
			isFirstHalf = false
			continue
		}

		if isFirstHalf {
			ingredient := strings.Split(row, "-")
			min, err := strconv.Atoi(ingredient[0])
			if err != nil {
				log.Fatal(err)
			}

			max, err := strconv.Atoi(ingredient[1])
			if err != nil {
				log.Fatal(err)
			}
			ingredientRange := IngredientRange{
				min: min,
				max: max,
			}

			ingredientRanges = append(ingredientRanges, ingredientRange)
		} else {
			ingredientID, err := strconv.Atoi(row)
			if err != nil {
				log.Fatal(err)
			}

			ingredients.Ingredients = append(ingredients.Ingredients, ingredientID)
		}
	}

	sort.Slice(ingredientRanges, func(i, j int) bool {
		return ingredientRanges[i].min < ingredientRanges[j].min
	})

	for _, ingredientRange := range ingredientRanges {
		if len(ingredients.FreshIngredientRanges) == 0 {
			ingredients.FreshIngredientRanges = append(
				ingredients.FreshIngredientRanges,
				ingredientRange,
			)
			continue
		}
		found := false
		for i, freshIngredientRange := range ingredients.FreshIngredientRanges {
			if ingredientRange.min >= freshIngredientRange.min &&
				ingredientRange.min <= freshIngredientRange.max || (ingredientRange.max >= freshIngredientRange.min &&
				ingredientRange.max <= freshIngredientRange.max) {
				ingredients.FreshIngredientRanges[i].min = min(
					ingredients.FreshIngredientRanges[i].min,
					ingredientRange.min,
				)
				ingredients.FreshIngredientRanges[i].max = max(
					ingredients.FreshIngredientRanges[i].max,
					ingredientRange.max,
				)
				found = true
			}
		}
		if !found {
			ingredients.FreshIngredientRanges = append(
				ingredients.FreshIngredientRanges,
				ingredientRange,
			)
		}
	}
}
