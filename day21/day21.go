package main

import (
	"advent-2020/utils"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	input := utils.ReadLines("./day21/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Number of ingredients without allergens: %d\n", Part1(input))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Canonical dangerous ingredient list: %s\n", Part2(input))
}

var regex = regexp.MustCompile("([^)]+) \\(contains ([^)]+)\\)")

func Part1(foods []string) int {
	ingredients, definiteAllergens := Parse(foods)

	cnt := 0
	for _, i := range ingredients {
		if _, ok := definiteAllergens[i]; !ok {
			cnt++
		}
	}

	return cnt
}

func Part2(foods []string) string {
	_, definiteAllergens := Parse(foods)

	var allergens []string
	ingredientByAllergen := make(map[string]string)
	for i, a := range definiteAllergens {
		allergens = append(allergens, a)
		ingredientByAllergen[a] = i
	}

	sort.Strings(allergens)

	var ret []string
	for _, a := range allergens {
		ret = append(ret, ingredientByAllergen[a])
	}

	return strings.Join(ret, ",")
}

func Parse(foods []string) (ingredients []string, definiteAllergens map[string]string) {
	// Final results get stored in these two variables:
	ingredients = make([]string, 0)
	definiteAllergens = make(map[string]string)

	// Map of allergens to the ingredients that might be related to them.
	// Throughout this method we whittle that list down to exactly one ingredient per allergen.
	possibleAllergens := make(map[string][]string)

	for _, line := range foods {
		m := regex.FindStringSubmatch(line)
		i, a := strings.Split(m[1], " "), strings.Split(m[2], ", ")
		for _, allergen := range a {
			if _, ok := possibleAllergens[allergen]; !ok {
				// We haven't seen this allergen yet; set its list of possible ingredients to the current ones
				possibleAllergens[allergen] = i
			} else {
				// We've seen this allergen before; intersect that list of ingredients with the current ones
				possibleAllergens[allergen] = utils.IntersectStrings(possibleAllergens[allergen], i)
			}
		}

		// Log each ingredient separately; we'll use this to count the number of ingredients used in Part 1
		for _, ingredient := range i {
			ingredients = append(ingredients, ingredient)
		}
	}

	// Using process of elimination, determine which ingredients are definitely associated to a given allergen
	// until all possible allergens have been identified.
	for len(possibleAllergens) > 0 {
		// Find an allergen which appears in only one ingredient
		for allergen, ingredients := range possibleAllergens {
			if len(ingredients) == 1 {
				// Found one! And we know it must be related to that one particular ingredient
				ingredient := ingredients[0]
				definiteAllergens[ingredient] = allergen

				// Remove this ingredient from the other possibilities
				for otherAllergen, otherAllergenIngredients := range possibleAllergens {
					possibleAllergens[otherAllergen] = utils.RemoveItem(otherAllergenIngredients, ingredient)
				}

				delete(possibleAllergens, allergen)
				break
			}
		}
	}

	return
}
