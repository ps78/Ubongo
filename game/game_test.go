package game_test

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
	"ubongo/base/array2d"
	"ubongo/base/array3d"
	"ubongo/base/vector"
	"ubongo/block"
	"ubongo/blockfactory"
	"ubongo/blockset"
	"ubongo/card"
	"ubongo/cardfactory"
	. "ubongo/game"
	"ubongo/gamesolution"
	"ubongo/problem"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1]
	g := New(p)

	assert.True(t, g.Shape.IsEqual(p.Shape), "Shape is wrong")
	assert.True(t, g.Volume.IsEqual(p.Shape.Extrude(p.Height)), "Wrong volume")
}

func TestGameString(t *testing.T) {
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1]
	g := New(p)
	s := g.String()
	assert.True(t, len(s) > 10)
}

func TestGameSolution(t *testing.T) {
	f := blockfactory.Get()
	b := []*block.B{f.ByNumber(1), f.ByNumber(2)}
	gs := gamesolution.New(b, []int{0, 0}, []vector.V{{0, 0, 0}, {0, 0, 0}})
	s := gs.String()
	assert.True(t, len(s) > 10)
}

func TestClear(t *testing.T) {
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1]
	g := New(p)

	g.Volume.Set(0, 1, 0, 1)
	g.Volume.Set(0, 1, 1, 1)
	g.Volume.Set(1, 0, 0, 1)

	g.Clear()

	assert.True(t, g.Volume.IsEqual(p.Shape.Extrude(p.Height)), "Clear() didn't produce the right result")
}

func TestClone(t *testing.T) {
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1]
	g := New(p)
	c := g.Clone()

	assert.True(t, g.Shape.IsEqual(c.Shape), "Shape does not match")
	assert.True(t, g.Volume.IsEqual(c.Volume), "Volume does not match")
	assert.True(t, g.Blocks.IsEqual(c.Blocks), "Block arrays do not match")
}

func TestTryAddBlock(t *testing.T) {
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1]
	p.Shape = array2d.NewFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}})
	g := New(p)
	origVolume := g.Volume.Clone()
	blockShape := blockfactory.Get().ByNumber(8).Shapes[0]
	pos := vector.V{0, 0, 0}

	// test a case where TryAdd should fail
	nok := g.TryAddBlock(blockShape, vector.V{3, 4, 1})
	assert.False(t, nok, "TryAddBlock did not return false where it should")
	assert.True(t, g.Volume.IsEqual(origVolume), "The the volume changed after a failed TryAddBlock() call")

	// test a case where it should succeed
	ok := g.TryAddBlock(blockShape, pos)
	assert.True(t, ok, "TryAddBlock returned no success where it should")
	exp := array3d.NewFromData([][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}})
	assert.True(t, exp.IsEqual(g.Volume), "The resulting volume after TryAddBlock is not as expected")
}

func TestRemoveBlock(t *testing.T) {
	g := new(G)
	g.Volume = array3d.NewFromData([][][]int8{{{0, 1}, {1, 1}, {-1, -1}, {-1, -1}}, {{-1, -1}, {0, 0}, {0, 0}, {-1, -1}}, {{0, 0}, {0, 0}, {0, 0}, {0, 0}}})
	origVolume := g.Volume.Clone()
	blockShape := blockfactory.Get().ByNumber(8).Shapes[0]
	pos := vector.V{0, 0, 0}

	// test case where block is outside volume, this should not change the volume
	nok := g.RemoveBlock(blockShape, vector.V{3, 4, 1})
	assert.False(t, nok, "RemoveBlock did not return false where it should")
	assert.True(t, g.Volume.IsEqual(origVolume), "The the volume changed after a failed RemoveBlock() call")

	// test case where removal works
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1]
	p.Shape = array2d.NewFromData([][]int8{{0, 0, -1, -1}, {-1, 0, 0, -1}, {0, 0, 0, 0}})
	exp := New(p)
	ok := g.RemoveBlock(blockShape, pos)
	assert.True(t, ok)
	assert.True(t, exp.Volume.IsEqual(g.Volume))
}

func TestSolveNoSolution(t *testing.T) {
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1].Clone()
	p.Blocks.RemoveAt(3)
	g := New(p)
	solutions := g.Solve()

	assert.Equal(t, 0, len(solutions), "Expected 0 solutions, but found %d", len(solutions))
}

func TestSolve(t *testing.T) {
	p := cardfactory.Get().Get(card.Difficult, 12).Problems[1]
	g := New(p)

	solutions := g.Solve()

	assert.Equal(t, 6, len(solutions), "Expected 6 solutions, but found %d", len(solutions))
}

func TestCreateSolutionStatistics(t *testing.T) {
	f := cardfactory.Get()
	rand.Seed(time.Now().Unix())
	csvFile := "solution_stats_test_" + strconv.Itoa(rand.Int()) + ".csv"
	defer os.Remove(csvFile)

	stats := CreateSolutionStatistics(f, csvFile)

	countAllProblems := len(f.GetAllProblems(card.Easy)) + len(f.GetAllProblems(card.Difficult))

	assert.Equal(t, countAllProblems, len(stats))
	_, err := os.Stat(csvFile)
	assert.Nil(t, err)

	assert.Panics(t, func() { CreateSolutionStatistics(f, "><?.txt") })
}

func TestIsPossibleCardSet(t *testing.T) {
	f := blockfactory.Get()
	shape := array2d.New(3, 3)
	okProblemsSet := map[int]*problem.P{
		1: problem.New(shape, 2, blockset.New(f.Blue_bighook, f.Blue_flash)),
		2: problem.New(shape, 2, blockset.New(f.Red_smallhook, f.Blue_flash)),
		3: problem.New(shape, 2, blockset.New(f.Yellow_gate, f.Blue_bighook)),
		4: problem.New(shape, 2, blockset.New(f.Blue_lighter, f.Yellow_hello)),
	}

	assert.True(t, IsPossibleCardSet(okProblemsSet))

	nokProblemsSet := map[int]*problem.P{
		1: problem.New(shape, 2, blockset.New(f.Blue_bighook, f.Blue_flash)),
		2: problem.New(shape, 2, blockset.New(f.Blue_bighook, f.Blue_flash)),
		3: problem.New(shape, 2, blockset.New(f.Blue_bighook, f.Blue_bighook)),
		4: problem.New(shape, 2, blockset.New(f.Blue_lighter, f.Blue_bighook)),
	}

	assert.False(t, IsPossibleCardSet(nokProblemsSet))

	emptyProblemSet := map[int]*problem.P{}
	assert.False(t, IsPossibleCardSet(emptyProblemSet))

	nilProblemSet := map[int]*problem.P{1: nil}
	assert.False(t, IsPossibleCardSet(nilProblemSet))
}

func TestGenerateCardSet(t *testing.T) {
	bf := blockfactory.Get()
	cf := cardfactory.Get()

	rand.Seed(time.Now().Unix())
	file := "cardset_test_" + strconv.Itoa(rand.Int()) + ".txt"
	defer os.Remove(file)

	cards := GenerateCardSet(cf, bf, card.Elephant, card.Easy, card.Easy, 2, 3, file)
	assert.Equal(t, 4, len(cards))
	_, err := os.Stat(file)
	assert.Nil(t, err)
}

func TestGenerateProblems(t *testing.T) {
	fp := cardfactory.Get()
	fb := blockfactory.Get()

	shape := fp.Get(card.Easy, 1).Problems[1].Shape
	problems := GenerateProblems(fb, shape, 3, 5, 10)

	assert.Equal(t, 10, len(problems))

	// check that each problem has a solution
	for _, p := range problems {
		g := New(p)
		solutions := g.Solve()
		assert.Less(t, 0, len(solutions))
	}
}
