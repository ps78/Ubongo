package main

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
	"ubongo/utils"
	"ubongo/utils/array2d"
	"ubongo/utils/array3d"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	p := GetCardFactory().Get(Difficult, 12).Problems[1]
	g := NewGame(p)

	assert.True(t, g.Shape.IsEqual(p.Shape), "Shape is wrong")
	assert.True(t, g.Volume.IsEqual(p.Shape.Extrude(p.Height)), "Wrong volume")
}

func TestGameString(t *testing.T) {
	p := GetCardFactory().Get(Difficult, 12).Problems[1]
	g := NewGame(p)
	s := g.String()
	assert.True(t, len(s) > 10)
}

func TestGameSolution(t *testing.T) {
	f := GetBlockFactory()
	b := []*Block{f.ByNumber(1), f.ByNumber(2)}
	gs := NewGameSolution(b, []int{0, 0}, []utils.Vector{{0, 0, 0}, {0, 0, 0}})
	s := gs.String()
	assert.True(t, len(s) > 10)
}

func TestClear(t *testing.T) {
	p := GetCardFactory().Get(Difficult, 12).Problems[1]
	g := NewGame(p)

	g.Volume.Set(0, 1, 0, OCCUPIED)
	g.Volume.Set(0, 1, 1, OCCUPIED)
	g.Volume.Set(1, 0, 0, OCCUPIED)

	g.Clear()

	assert.True(t, g.Volume.IsEqual(p.Shape.Extrude(p.Height)), "Clear() didn't produce the right result")
}

func TestClone(t *testing.T) {
	p := GetCardFactory().Get(Difficult, 12).Problems[1]
	g := NewGame(p)
	c := g.Clone()

	assert.True(t, g.Shape.IsEqual(c.Shape), "Shape does not match")
	assert.True(t, g.Volume.IsEqual(c.Volume), "Volume does not match")
	assert.True(t, g.Blocks.IsEqual(c.Blocks), "Block arrays do not match")
}

func TestTryAddBlock(t *testing.T) {
	p := GetCardFactory().Get(Difficult, 12).Problems[1]
	p.Shape = array2d.NewFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}})
	g := NewGame(p)
	origVolume := g.Volume.Clone()
	blockShape := NewBlock8().Shapes[0]
	pos := utils.Vector{0, 0, 0}

	// test a case where TryAdd should fail
	nok := g.TryAddBlock(blockShape, utils.Vector{3, 4, 1})
	assert.False(t, nok, "TryAddBlock did not return false where it should")
	assert.True(t, g.Volume.IsEqual(origVolume), "The the volume changed after a failed TryAddBlock() call")

	// test a case where it should succeed
	ok := g.TryAddBlock(blockShape, pos)
	assert.True(t, ok, "TryAddBlock returned no success where it should")
	exp := array3d.NewFromData([][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}})
	assert.True(t, exp.IsEqual(g.Volume), "The resulting volume after TryAddBlock is not as expected")
}

func TestRemoveBlock(t *testing.T) {
	g := new(Game)
	g.Volume = array3d.NewFromData([][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}})
	origVolume := g.Volume.Clone()
	blockShape := NewBlock8().Shapes[0]
	pos := utils.Vector{0, 0, 0}

	// test case where block is outside volume, this should not change the volume
	nok := g.RemoveBlock(blockShape, utils.Vector{3, 4, 1})
	assert.False(t, nok, "RemoveBlock did not return false where it should")
	assert.True(t, g.Volume.IsEqual(origVolume), "The the volume changed after a failed RemoveBlock() call")

	// test case where removal works
	p := GetCardFactory().Get(Difficult, 12).Problems[1]
	p.Shape = array2d.NewFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}})
	exp := NewGame(p)
	ok := g.RemoveBlock(blockShape, pos)
	assert.True(t, ok)
	assert.True(t, exp.Volume.IsEqual(g.Volume))
}

func TestSolveNoSolution(t *testing.T) {
	p := GetCardFactory().Get(Difficult, 12).Problems[1].Clone()
	p.Blocks.RemoveAt(3)
	g := NewGame(p)
	solutions := g.Solve()

	assert.Equal(t, 0, len(solutions), "Expected 0 solutions, but found %d", len(solutions))
}

func TestSolve(t *testing.T) {
	p := GetCardFactory().Get(Difficult, 12).Problems[1]
	g := NewGame(p)

	solutions := g.Solve()

	assert.Equal(t, 6, len(solutions), "Expected 6 solutions, but found %d", len(solutions))
}

func TestCreateSolutionStatistics(t *testing.T) {
	f := GetCardFactory()
	rand.Seed(time.Now().Unix())
	csvFile := "solution_stats_test_" + strconv.Itoa(rand.Int()) + ".csv"
	defer os.Remove(csvFile)

	stats := f.CreateSolutionStatistics(csvFile)

	countAllProblems := len(f.GetAllProblems(Easy)) + len(f.GetAllProblems(Difficult))

	assert.Equal(t, countAllProblems, len(stats))
	_, err := os.Stat(csvFile)
	assert.Nil(t, err)
}

func TestIsPossibleCardSet(t *testing.T) {
	f := GetBlockFactory()
	shape := array2d.New(3, 3)
	okProblemsSet := map[int]*Problem{
		1: NewProblem(shape, 2, NewBlockset(f.Blue_bighook, f.Blue_flash)),
		2: NewProblem(shape, 2, NewBlockset(f.Red_smallhook, f.Blue_flash)),
		3: NewProblem(shape, 2, NewBlockset(f.Yellow_gate, f.Blue_bighook)),
		4: NewProblem(shape, 2, NewBlockset(f.Blue_lighter, f.Yellow_hello)),
	}

	assert.True(t, IsPossibleCardSet(okProblemsSet))

	nokProblemsSet := map[int]*Problem{
		1: NewProblem(shape, 2, NewBlockset(f.Blue_bighook, f.Blue_flash)),
		2: NewProblem(shape, 2, NewBlockset(f.Blue_bighook, f.Blue_flash)),
		3: NewProblem(shape, 2, NewBlockset(f.Blue_bighook, f.Blue_bighook)),
		4: NewProblem(shape, 2, NewBlockset(f.Blue_lighter, f.Blue_bighook)),
	}

	assert.False(t, IsPossibleCardSet(nokProblemsSet))

	emptyProblemSet := map[int]*Problem{}
	assert.False(t, IsPossibleCardSet(emptyProblemSet))

	nilProblemSet := map[int]*Problem{1: nil}
	assert.False(t, IsPossibleCardSet(nilProblemSet))
}

func TestGenerateCardSet(t *testing.T) {
	bf := GetBlockFactory()
	cf := GetCardFactory()

	rand.Seed(time.Now().Unix())
	file := "cardset_test_" + strconv.Itoa(rand.Int()) + ".txt"
	defer os.Remove(file)

	cards := GenerateCardSet(cf, bf, Elephant, Easy, Easy, 2, 3, file)
	assert.Equal(t, 4, len(cards))
	_, err := os.Stat(file)
	assert.Nil(t, err)
}
