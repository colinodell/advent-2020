package main

import (
	"advent-2020/utils"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := utils.ReadFile("./day20/input.txt")

	fmt.Println("----- Part 1 -----")
	fmt.Printf("Product of corner tile IDs: %d\n", Part1(input))

	fmt.Println("----- Part 2 -----")
	fmt.Printf("Roughness of the water: %d\n", Part2(input))
}

var SeaMonster map[utils.Vector2]struct{}
func init() {
	buildSeaMonster()
}

func Part1(input string) int {
	tiles := NewTileSet(input)
	allEdges := tiles.CountAllPossibleEdges()

	result := 1

	for id, tile := range tiles {
		// Technically the corner tiles only have 2 unique, unshared edges, but remember that
		// we've also checked for its reverse orientation, which doubles the number of edges,
		// so we need to check for 4 instead of 2
		if tile.CountUniqueEdges(allEdges) == 4 {
			result *= id
		}
	}

	return result
}

func Part2(input string) int {
	tiles := NewTileSet(input)
	stitched := tiles.Stitch()

	return calculateRoughness(stitched)
}

type TileSet map[int]Tile

type Tile struct {
	num int
	grid [][]rune
	size int
}

// Returns a list of all edges in all possible orientations, along with how many tile-orientation combos have that edge
func (ts TileSet) CountAllPossibleEdges() map[string]int {
	// Tracks the number of distinct edges found in all tiles
	ret := make(map[string]int)

	for _, t := range ts {
		for _, edge := range t.AllPossibleEdges() {
			ret[edge]++
		}
	}

	return ret
}

// Parses a tile from text
func NewTile(input string) Tile {
	s := strings.SplitN(input, "\n", 2)
	var id int
	fmt.Sscanf(s[0], "Tile %d:", &id)

	lines := strings.Split(s[1], "\n")
	size := len(lines)

	t := Tile{num: id, grid: make([][]rune, size), size: size}
	for i, line := range lines {
		t.grid[i] = make([]rune, size)
		for j, char := range line {
			t.grid[i][j] = char
		}
	}

	return t
}

// Renders the parsed tile as a string; useful for debugging
func (t *Tile) String() string {
	var sb strings.Builder
	for i, line := range t.grid {
		if i != 0 {
			sb.WriteString("\n")
		}

		sb.WriteString(string(line))
	}

	return sb.String()
}

// Returns the top edge
func (t *Tile) Top() string {
	return string(t.grid[0])
}

// Returns the bottom edge
func (t *Tile) Bottom() string {
	return string(t.grid[t.size-1])
}

// Returns the left edge
func (t *Tile) Left() string {
	var left strings.Builder
	for i := 0; i < len(t.grid[0]); i++ {
		left.WriteRune(t.grid[i][0])
	}

	return left.String()
}

// Returns the right edge
func (t *Tile) Right() string {
	var right strings.Builder
	for i := 0; i < t.size; i++ {
		right.WriteRune(t.grid[i][t.size-1])
	}

	return right.String()
}

// Returns the edges in the current orientation
func (t *Tile) Edges() []string {
	return []string{
		t.Top(),
		t.Bottom(),
		t.Left(),
		t.Right(),
	}
}

// Returns all possible edges regardless of current orientation
// I manually worked out that this would be the current edges plus their mirror images
func (t *Tile) AllPossibleEdges() []string {
	edges := t.Edges()

	return []string{
		edges[0],
		reverseString(edges[0]),
		edges[1],
		reverseString(edges[1]),
		edges[2],
		reverseString(edges[2]),
		edges[3],
		reverseString(edges[3]),
	}
}

func reverseString(input string) string {
	var sb strings.Builder
	runes := []rune(input)
	for i := len(runes) - 1; 0 <= i; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

// Counts how many edges are not shared with other tiles
func (t Tile) CountUniqueEdges(edges map[string]int) int {
	ret := 0
	for _, edge := range t.AllPossibleEdges() {
		if edges[edge] == 2 {
			ret++
		}
	}

	return ret
}

// Rotates the tile 90 degrees clockwise
func (t *Tile) Rotate() Tile {
	newGrid := make([][]rune, t.size)

	for y := range t.grid {
		newGrid[y] = make([]rune, t.size)
	}

	for y := range t.grid {
		for x := range t.grid[y] {
			newGrid[y][t.size-x-1] = t.grid[x][y]
		}
	}

	return Tile{
		num: t.num,
		grid: newGrid,
		size: t.size,
	}
}

// Flips the tile so that the top is on the bottom and vice-versa
func (t *Tile) FlipTopBottom() Tile {
	newGrid := make([][]rune, t.size)

	for y := range t.grid {
		newGrid[y] = t.grid[t.size - y - 1]
	}

	return Tile{
		num: t.num,
		grid: newGrid,
		size: t.size,
	}
}

// Flips the tile so that the left is on the right and vice-versa
func (t *Tile) FlipLeftRight() Tile {
	newGrid := make([][]rune, t.size)

	for y := range t.grid {
		newGrid[y] = make([]rune, t.size)
		for x, char := range t.grid[y] {
			newGrid[y][t.size - x - 1] = char
		}
	}

	return Tile{
		num: t.num,
		grid: newGrid,
		size: t.size,
	}
}

// Removes the border from the tile
func (t *Tile) WithoutBorder() Tile {
	newTile := Tile{num: t.num, grid: make([][]rune, t.size-2), size: t.size - 2}
	for y := range t.grid {
		if y == 0 || y == t.size - 1 {
			continue
		}

		newTile.grid[y-1] = make([]rune, newTile.size)

		for x, char := range t.grid[y] {
			if x == 0 || x == t.size - 1 {
				continue
			}

			newTile.grid[y-1][x-1] = char
		}
	}

	return newTile
}

// Produces a set of all possible orientations for the given tile
func (t *Tile) AllPossibleOrientations() []Tile {
	ret := make([]Tile, 12)

	possible := *t
	for i := 0; i < 4; i++ {
		ret[i*3] = possible
		ret[i*3+1] = possible.FlipLeftRight()
		ret[i*3+2] = possible.FlipTopBottom()

		possible = possible.Rotate()
	}

	return ret
}

// Rotates and/or flips the tile until the left and top edges match what we're looking for
// "left" and "top" can either be:
//   - A string containing a specific edge to match
//   - An empty string, indicating that we don't care what the edge is, so long as it's unique
//     (in other words, it's not shared with any other tile, so it must be along the outside border)
func (t *Tile) Orient(left, top string, edges map[string]int) *Tile {
	for _, orientation := range t.AllPossibleOrientations() {
		matchesLeft := (left == "" && edges[orientation.Left()] == 1) || (left == orientation.Left())
		matchesTop := (top == "" && edges[orientation.Top()] == 1) || (top == orientation.Top())
		if matchesLeft && matchesTop {
			return &orientation
		}
	}

	return nil
}

// Sea Monster identification based on https://gitlab.com/kurisuchan/advent-of-code-2020/-/blob/master/pkg/day20/day20.go
func calculateRoughness(stitched Tile) int {
	var t Tile
	found := false
	for _, t = range stitched.AllPossibleOrientations() {
		for y := range t.grid {
			for x := range t.grid[y] {
				// Does a Sea Monster exist here?
				match := true
				for m := range SeaMonster {
					if y+m.Y >= len(t.grid) {
						match = false
						break
					}

					if x+m.X >= len(t.grid[y+m.Y]) {
						match = false
						break
					}

					if t.grid[y+m.Y][x+m.X] != '#' {
						match = false
						break
					}
				}

				if !match {
					continue
				}

				// It does!
				found = true
				for m := range SeaMonster {
					t.grid[y+m.Y][x+m.X] = 'O'
				}
			}
		}

		if found {
			break
		}
	}

	// Count the waves
	waves := 0
	for y := range t.grid {
		for x := range t.grid[y] {
			if t.grid[y][x] == '#' {
				waves++
			}
		}
	}

	return waves
}

// Parses the puzzle input
func NewTileSet(input string) TileSet {
	tiles := strings.Split(input, "\n\n")
	result := make(TileSet, len(tiles))

	for _, tileData := range tiles {
		t := NewTile(tileData)
		result[t.num] = t
	}

	return result
}

// Takes the input tiles, orients them correctly, and returns the stitched image as a large tile
func (ts *TileSet) Stitch() Tile {
	allPossibleEdges := ts.CountAllPossibleEdges()

	// The edges that the next tile must match
	var left, top string
	// Tracks which tiles have already been used
	usedKeys := make(map[int]struct{})
	// Helps us track the "top" search value we'll need in future rows
	usedTilesByPosition := make(map[utils.Vector2]Tile)
	// Helps us stitch together the final image
	var stitcher ImageStitcher

	pos := utils.Vector2{}

	// Pick any corner to be our top-left; orient it correctly
	firstTileInRow := ts.FindRandomCorner(allPossibleEdges).Orient("", "", allPossibleEdges)

	// Save this corner
	usedKeys[firstTileInRow.num] = struct{}{}
	usedTilesByPosition[pos] = *firstTileInRow
	stitcher = NewImageStitcher(firstTileInRow.size - 2, len(*ts))
	stitcher.AddTile(pos, firstTileInRow.WithoutBorder())

	// Determine the search criteria for the next tile
	pos = pos.Add(utils.Vector2{X: 1})
	left, top = firstTileInRow.Right(), ""

	// Keep looping until all tiles have been used
	for len(usedKeys) < len(*ts) {
		// Track whether we found any matching tiles this time around
		found := false

		// Try to find the next matching tile
		for id, next := range *ts {
			// Have we already used this tile?
			if _, ok := usedKeys[id]; ok {
				continue
			}

			// Does this tile fit?
			if candidate := next.Orient(left, top, allPossibleEdges); candidate != nil {
				// Save this tile
				usedKeys[candidate.num] = struct{}{}
				usedTilesByPosition[pos] = *candidate
				stitcher.AddTile(pos, candidate.WithoutBorder())

				// Determine the search criteria for the next tile
				pos = pos.Add(utils.Vector2{X: 1})
				left = candidate.Right()
				if aboveNext, ok := usedTilesByPosition[pos.Add(utils.Vector2{Y: -1})]; ok {
					top = aboveNext.Bottom()
				}

				found = true
				if firstTileInRow == nil {
					firstTileInRow = candidate
				}
				break
			}
		}

		// Failing to find another tile in this row means we've completed the row and should
		// starting working on the next one
		if !found {
			pos = utils.Vector2{X: 0, Y: pos.Y + 1}
			// Next row should start with a unique border edge on the left, and it's top should match
			// the previous row start's bottom
			left, top = "", firstTileInRow.Bottom()
			// We don't yet know while tile will be the first one in this row
			firstTileInRow = nil
		}
	}

	// All tiles have been arranged correctly! Return the final stitched image as a Tile
	return stitcher.Complete()
}

// Returns a random corner (we don't really care which one)
func (ts *TileSet) FindRandomCorner(allEdges map[string]int) *Tile {
	for _, tile := range *ts {
		if tile.CountUniqueEdges(allEdges) == 4 {
			return &tile
		}
	}

	panic("no corner tile found")
}

type ImageStitcher struct {
	contents [][]rune
	tileSize int
	height, width int
}

func NewImageStitcher(tileSize, tileCount int) ImageStitcher {
	width := int(math.Sqrt(float64(tileCount)))
	height := width

	contents := make([][]rune, tileSize * height)
	for i := 0; i < tileSize * height; i++ {
		contents[i] = make([]rune, tileSize * width)
	}

	return ImageStitcher{
		contents: contents,
		tileSize: tileSize,
		width: width,
		height: height,
	}
}

func (i *ImageStitcher) AddTile(pos utils.Vector2, t Tile) {
	for tileY, row := range t.grid {
		imageY := (i.tileSize * pos.Y) + tileY

		for tileX, char := range row {
			imageX := (i.tileSize * pos.X) + tileX
			i.contents[imageY][imageX] = char
		}
	}
}

func (i *ImageStitcher) Complete() Tile {
	return Tile{
		num: 0,
		grid: i.contents,
		size: i.tileSize * i.height,
	}
}

func buildSeaMonster() {
	SeaMonster = make(map[utils.Vector2]struct{})
	pattern := "                  # \n#    ##    ##    ###\n #  #  #  #  #  #   "
	for y, line := range strings.Split(pattern, "\n") {
		for x, r := range line {
			if r == '#' {
				SeaMonster[utils.Vector2{X: x, Y: y}] = struct{}{}
			}
		}
	}
}
