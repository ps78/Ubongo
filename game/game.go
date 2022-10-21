package game

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"ubongo/base/array2d"
	"ubongo/base/array3d"
	"ubongo/base/vector"
	"ubongo/blockfactory"
	"ubongo/blockset"
	"ubongo/card"
	"ubongo/cardfactory"
	"ubongo/gamesolution"
	"ubongo/problem"
)

type G struct {
	Shape  *array2d.A
	Volume *array3d.A
	Blocks *blockset.S
}

// The number of blocks of each type in the original Ubongo game
// map[BlockNumer]Count
var UbongoBlockSet map[int]int = map[int]int{
	1:  2, // yellow hello
	2:  2, // yellow bighook
	3:  3, // yellow smallhook
	4:  2, // yellow gate
	5:  2, // blue bighook
	6:  2, // blue flash
	7:  3, // blue lighter
	8:  4, // blue v
	9:  3, // red stool
	10: 3, // red smallhook
	11: 2, // red bighook
	12: 2, // red flash
	13: 2, // green flash
	14: 2, // green bighook
	15: 2, // green T
	16: 4, // green L
}

// New creates a new game, initialized with the given shape and height and an empty volume
func New(p *problem.P) *G {
	return &G{
		Shape:  p.Shape.Clone(),
		Volume: p.Shape.Extrude(p.Height),
		Blocks: p.Blocks.Clone()}
}

// Returns a nicely formatted string representation of the game
func (g *G) String() string {
	return fmt.Sprintf("Game (area %d, volume %d, empty %d)",
		g.Shape.Count(0), g.Shape.Count(0)*g.Volume.DimZ, g.Volume.Count(0))
}

// Removes all blocks from a game
func (g *G) Clear() {
	for x := 0; x < g.Volume.DimX; x++ {
		for y := 0; y < g.Volume.DimY; y++ {
			for z := 0; z < g.Volume.DimZ; z++ {
				g.Volume.Set(x, y, z, g.Shape.Get(x, y))
			}
		}
	}
}

// Creates a copy of the game
func (g *G) Clone() *G {
	return &G{
		Shape:  g.Shape.Clone(),
		Volume: g.Volume.Clone(),
		Blocks: g.Blocks.Clone()}
}

// Tries to add the given block to the game volume
// returns true if successful, false if not
func (g *G) TryAddBlock(block *array3d.A, pos vector.V) bool {
	// check overall dimensions
	if pos[0]+block.DimX > g.Volume.DimX ||
		pos[1]+block.DimY > g.Volume.DimY ||
		pos[2]+block.DimZ > g.Volume.DimZ {
		return false
	}

	// step 1: test if it is possible to add block
	for x := 0; x < block.DimX; x++ {
		for y := 0; y < block.DimY; y++ {
			for z := 0; z < block.DimZ; z++ {
				if block.Get(x, y, z) == 1 && g.Volume.Get(x+pos[0], y+pos[1], z+pos[2]) != 0 {
					return false
				}
			}
		}
	}

	// step 2: actually add the block, the expensive step is the cloning of the object
	v := g.Volume.Clone()
	for x := 0; x < block.DimX; x++ {
		for y := 0; y < block.DimY; y++ {
			for z := 0; z < block.DimZ; z++ {
				if block.Get(x, y, z) == 1 {
					v.Set(x+pos[0], y+pos[1], z+pos[2], 1)
				}
			}
		}
	}

	// replace the game's volume with the new one containing the block
	g.Volume = v
	return true
}

// Removes the block at the given position from the volume
// This does not check if the block is actually present and
// simply sets all values from 1 to 0
func (g *G) RemoveBlock(block *array3d.A, pos vector.V) bool {
	// check overall dimensions
	if pos[0]+block.DimX > g.Volume.DimX ||
		pos[1]+block.DimY > g.Volume.DimY ||
		pos[2]+block.DimZ > g.Volume.DimZ {
		return false
	}

	for x := 0; x < block.DimX; x++ {
		for y := 0; y < block.DimY; y++ {
			for z := 0; z < block.DimZ; z++ {
				// only run following test if the block-cube is solid at the current position
				if block.Get(x, y, z) == 1 {
					// if space is occupied by block, set it to 0
					if g.Volume.Get(x+pos[0], y+pos[1], z+pos[2]) == 1 {
						g.Volume.Set(x+pos[0], y+pos[1], z+pos[2], 0) // mark space as empty
					}
				}
			}
		}
	}
	return true
}

// Finds all solutino for a given game using the set of blocks provided
func (g *G) Solve() []*gamesolution.S {
	// check the sum of the block volumes, it must match the empty volume of the game to yield a solution
	if g.Volume.Count(0) != g.Blocks.Volume() {
		return []*gamesolution.S{}
	}

	// working arrays for the recursive solver:
	solutions := make([]*gamesolution.S, 0)
	shapeIdx := make([]int, 0)
	shifts := make([]vector.V, 0)

	g.recursiveSolver(0, &shapeIdx, &shifts, &solutions)

	return solutions
}

// Recursive function called by Solve, don't call directly
func (g *G) recursiveSolver(blockIdx int, shapeIndices *[]int, shifts *[]vector.V, solutions *[]*gamesolution.S) {
	gameBox := g.Volume.GetBoundingBox()

	block := g.Blocks.Get(blockIdx)

	for shapeIdx, shape := range block.Shapes {
		*shapeIndices = append(*shapeIndices, shapeIdx)

		shiftVectors := gameBox.GetShiftVectors(shape.GetBoundingBox())
		for _, shift := range shiftVectors {
			if ok := g.TryAddBlock(shape, shift); ok {

				*shifts = append(*shifts, shift)

				// if this was the last block, stop recursion
				if blockIdx == g.Blocks.Count-1 {
					// check if we have a solution
					if g.Volume.Count(0) == 0 {
						*solutions = append(*solutions, gamesolution.New(g.Blocks.AsSlice(), *shapeIndices, *shifts))
					}
					// if it wasn't the last block, continue recursion
				} else {
					g.recursiveSolver(blockIdx+1, shapeIndices, shifts, solutions)
				}

				*shifts = (*shifts)[:len(*shifts)-1]

				g.RemoveBlock(shape, shift)
			}
		} // end loop over shifts

		*shapeIndices = (*shapeIndices)[:len(*shapeIndices)-1]

	} // end loop over shapes
}

// Represents a single entry of the output of CreateSolutionStatistics()
type SolutionStatisticsRecord struct {
	Difficulty    card.UbongoDifficulty
	Animal        card.UbongoAnimal
	CardNumber    int
	DiceNumber    int
	Area          int
	Height        int
	SolutionCount int
	Blocks        *blockset.S
}

// Solves all Easy & Difficult problems and returns the statistics
// If the csvFile parameter is provided (and not empty), the data is also
// written to a csv file
func CreateSolutionStatistics(f *cardfactory.F, csvFile string) []SolutionStatisticsRecord {

	// create the dataset to return / write
	records := make([]SolutionStatisticsRecord, 0)
	for _, difficulty := range []card.UbongoDifficulty{card.Easy, card.Difficult} {
		for _, c := range f.GetAll(difficulty) {
			for diceNumber, p := range c.Problems {
				g := New(p)
				solutions := g.Solve()
				records = append(records, SolutionStatisticsRecord{
					c.Difficulty, c.Animal, c.CardNumber, diceNumber, p.Area, p.Height,
					len(solutions), p.Blocks})
			}
		}
	}

	// order the problems by difficulty, cardnumber, dicenumber
	sortOrder := func(rec SolutionStatisticsRecord) int {
		return int(rec.Difficulty)*10000000 + rec.CardNumber*1000 + rec.DiceNumber
	}
	sort.Slice(records, func(i, j int) bool {
		return sortOrder(records[i]) < sortOrder(records[j])
	})

	// optionally write csv file
	if csvFile != "" {
		file, err := os.Create(csvFile)
		if err != nil {
			panic(fmt.Sprintf("Failed to open file %s with error %v", csvFile, err))
		}
		defer file.Close()

		w := csv.NewWriter(file)
		defer w.Flush()
		w.Write([]string{"Difficulty", "Animal", "CardNumber", "DiceNumber", "Area", "Height", "SolutionCount", "Blocks"})
		for _, rec := range records {
			err = w.Write([]string{
				rec.Difficulty.String(),
				rec.Animal.String(),
				strconv.Itoa(rec.CardNumber),
				strconv.Itoa(rec.DiceNumber),
				strconv.Itoa(rec.Area),
				strconv.Itoa(rec.Height),
				strconv.Itoa(rec.SolutionCount),
				rec.Blocks.String(),
			})
			if err != nil {
				panic(fmt.Sprintf("Error writing solution statistics file: %v", err))
			}
		}
	}

	return records
}

// verifies if the set of problems given can be used
// as a set in the game. Condition is that the combined
// blocks are available in the game
// The map-key is the card-number
func IsPossibleCardSet(problems map[int]*problem.P) bool {
	if len(problems) == 0 {
		return false
	}

	blockStat := map[int]int{} // counts the blocks per blocknumber
	for _, p := range problems {
		if p == nil {
			return false
		}
		for _, block := range p.Blocks.AsSlice() {
			if _, ok := blockStat[block.Number]; !ok {
				blockStat[block.Number] = 1
			} else {
				blockStat[block.Number] += 1
			}
		}
	}

	for blockNum, blockCount := range blockStat {
		if blockCount > UbongoBlockSet[blockNum] {
			return false
		}
	}

	return true
}

// Creates 10 sets of problems for diceNumber=1..10 where each problem set
// consists of the 4 cards assiciated with the given animal and the blocks are
// chosen such that the game contains enough blocks to play a round with people
// for every possible throw of the dice (and of course every problem has a solution)
// Returns: map[diceNumber][cardNumber]*Problem
// Optionally write the result to the given file, if not empty
func GenerateCardSet(bc *cardfactory.F, bf *blockfactory.F,
	animal card.UbongoAnimal, sourceDifficulty, targetDifficulty card.UbongoDifficulty, height, blockCount int, outputFile string) []*card.C {

	// ** Utility types / functions and constants ** //

	maxTry := 200               // number of tries to build a consistent problem set
	numProblemsPerDiceNum := 20 // number of problems to generate per diceNumber and card

	// this key is used in the 'problems' map
	type key struct {
		animal     card.UbongoAnimal
		diceNumber int
	}
	// this function selects a shape from a given card with a diceNumber
	shapeSelector := func(card *card.C, diceNumber int) *array2d.A {
		if diceNumber <= 5 {
			return card.Problems[1].Shape // top shape
		} else {
			return card.Problems[8].Shape // bottom shape
		}
	}

	// ** Generate problems ** //
	problems := map[key]map[int][]*problem.P{} // value of map: map[cardnumber](problems with with same animal/dice/cardnum)
	sourceCards := bc.GetByAnimal(sourceDifficulty, animal)

	type item struct {
		key        key
		cardNumber int
		problems   []*problem.P
	}
	queueSize := 10 * len(sourceCards)
	queue := make(chan item, queueSize)
	for _, card := range sourceCards {
		for diceNumber := 1; diceNumber <= 10; diceNumber++ {
			shape := shapeSelector(card, diceNumber)
			curKey := key{animal, diceNumber}
			// create new entry in map if necessary
			if _, ok := problems[curKey]; !ok {
				problems[curKey] = map[int][]*problem.P{}
			}
			go func(cardNum int) {
				probs := GenerateProblems(bf, shape, height, blockCount, numProblemsPerDiceNum)
				queue <- item{curKey, cardNum, probs}
			}(card.CardNumber)
		}
	}
	// read results from channel
	for i := 0; i < queueSize; i++ {
		el := <-queue
		problems[el.key][el.cardNumber] = el.problems
	}

	// ** Initialize set of cards to return ** //

	cardSet := make(map[int]*card.C) // key = CardNumber
	for _, crd := range sourceCards {
		cardSet[crd.CardNumber] = card.New(crd.CardNumber, targetDifficulty, animal, make(map[int]*problem.P))
	}

	// ** try to build sets for each diceNumber ** //

	for diceNumber := 1; diceNumber <= 10; diceNumber++ {
		curKey := key{animal, diceNumber}

		for try := 0; try < maxTry; try++ {
			// randomly choose one problem from each card/dicenum
			problemSet := make(map[int]*problem.P) // key=CardNumber
			for cardNum, probs := range problems[curKey] {
				if len(probs) == 0 {
					problemSet = nil
					break
				}
				problemSet[cardNum] = probs[rand.Intn(len(probs))]
			}
			if problemSet == nil {
				break
			}
			if IsPossibleCardSet(problemSet) {
				for cardNum, prob := range problemSet {
					cardSet[cardNum].Problems[diceNumber] = prob
				}
				break
			}
		}
	}

	// flatten and sort map to array
	result := make([]*card.C, 0)
	for _, crd := range cardSet {
		result = append(result, crd)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].CardNumber < result[j].CardNumber
	})

	// write to file
	if outputFile != "" {
		f, _ := os.Create(outputFile)
		defer f.Close()
		for _, c := range result {
			f.WriteString(c.VerbousString())
		}
	}

	return result
}

// GenerateProblems creates numProblems new problems based on the given
// parameters (height, shape, blockCount)
func GenerateProblems(bf *blockfactory.F, shape *array2d.A, height, blockCount, numProblems int) []*problem.P {
	multiplier := 5 // we generate more problems than requested, as some might not have a solution
	results := make([]*problem.P, 0)

	// generate random blocksets, more than we need, as not all might be solvable
	sets := bf.GenerateBlocksets(shape.Count(0)*height, blockCount, multiplier*numProblems)

	for i := range sets {
		p := problem.New(shape, height, sets[i])
		g := New(p)
		solutions := g.Solve()

		if len(solutions) > 0 {
			results = append(results, p)
		}

		if len(results) >= numProblems {
			break
		}
	}
	return results
}
