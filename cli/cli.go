// Package cli file contains the functionality to run the command line interface of the Ubongo app
package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"ubongo/blockfactory"
	"ubongo/card"
	"ubongo/cardfactory"

	"ubongo/game"
	"ubongo/graphics"
)

// Cli represents the command line interface
type Cli struct {
	Menu []MenuEntry
}

type MenuEntry struct {
	Number int
	Title  string
	Action MenuOptionFunc
}

// MenuOptionFunc is a parameter-less function representing a menu option
type MenuOptionFunc func()

// New creates a new instance of the command line interface
func New() *Cli {
	cli := new(Cli)
	cli.Menu = []MenuEntry{
		{1, "Render all blocks", menuOptionRenderAllBlocks},
		{2, "Calculate solution statistics", menuOptionCalcSolutionStatistics},
		{3, "Generate insane problems", menuOptionGenerateInsaneProblems},
		{4, "Visualize a solution", menuOptionVisualizeSolution},
	}
	return cli
}

// ShowMenu prints the menu options to the console
func (cli *Cli) ShowMenu() {
	fmt.Println("---------------------------------------")
	fmt.Println("Ubongo")
	fmt.Println("---------------------------------------")
	for _, menu := range cli.Menu {
		fmt.Printf(" %d - %s\n", menu.Number, menu.Title)
	}
	fmt.Println("")
}

// GetMenuChoice waits for user input of a menu and returns the value
func (cli *Cli) GetMenuChoice() int {
	reader := bufio.NewReader(os.Stdin)
	replacer := strings.NewReplacer("\n", "", "\r", "")
	for {
		fmt.Print("Choose option: ")

		input, readErr := reader.ReadString('\n')
		option, convErr := strconv.ParseInt(replacer.Replace(input), 10, 32)

		if readErr == nil && convErr == nil && cli.IsValidOption(int(option)) {
			return int(option)
		} else {
			fmt.Println("Invlid option")
		}
	}
}

// Run is the main routine of the command line interface
func (cli *Cli) Run() {
	cli.ShowMenu()

	option := cli.GetMenuChoice()

	for _, entry := range cli.Menu {
		if entry.Number == option {
			start := time.Now()
			entry.Action()
			fmt.Printf("%v elapsed\n", time.Since(start))
		}
	}
}

// IsValidOption returns true if n is a valid selection in the current menu of cli
func (cli *Cli) IsValidOption(n int) bool {
	for _, item := range cli.Menu {
		if item.Number == n {
			return true
		}
	}
	return false
}

func menuOptionRenderAllBlocks() {
	bf := blockfactory.Get()
	blockTargetPath := "./results/images"
	blockRenderResX := 500
	blockRenderResY := 500
	blocks := bf.GetAll()
	graphics.RenderAll(blocks, blockTargetPath, blockRenderResX, blockRenderResY)
	fmt.Printf("Rendered %d blocks to path %s at resolution %dx%d\n", len(blocks), blockTargetPath, blockRenderResX, blockRenderResY)
}

func menuOptionCalcSolutionStatistics() {
	cf := cardfactory.Get()
	game.CreateSolutionStatistics(cf, "./results/solutions.csv")
}

func menuOptionGenerateInsaneProblems() {
	bf := blockfactory.Get()
	cf := cardfactory.Get()
	for _, animal := range card.AllAnimals() {
		game.GenerateCardSet(cf, bf, animal, card.Easy, card.Insane, 3, 5, fmt.Sprintf("./results/cards/%s.txt", animal))
	}
}

func menuOptionVisualizeSolution() {
	cf := cardfactory.Get()
	gs := game.New(cf.Get(card.Difficult, 1).Problems[1]).Solve()[0]
	graphics.Visualize(gs)
}
