// Package cli file contains the functionality to run the command line interface of the Ubongo app
package cli

import (
	"bufio"
	"fmt"
	"os"
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
	Menu []*MenuEntry

	// if this flag is set to true, the program will terminate as soon as possible
	doQuitFlag bool
}

type MenuEntry struct {
	Key    string
	Title  string
	Action MenuOptionFunc
}

// MenuOptionFunc is a parameter-less function representing a menu option
type MenuOptionFunc func(cli *Cli)

// New creates a new instance of the command line interface
func New() *Cli {
	cli := new(Cli)
	cli.Menu = []*MenuEntry{
		{"1", "Render all blocks", menuOptionRenderAllBlocks},
		{"2", "Calculate solution statistics", menuOptionCalcSolutionStatistics},
		{"3", "Generate insane problems", menuOptionGenerateInsaneProblems},
		{"4", "Visualize a solution", menuOptionVisualizeSolution},
		{"0", "Quit", menuOptionQuit},
	}
	return cli
}

// ShowMenu prints the menu options to the console
func (cli *Cli) ShowMenu() {
	fmt.Println("---------------------------------------")
	fmt.Println("Ubongo")
	fmt.Println("---------------------------------------")
	for _, menu := range cli.Menu {
		fmt.Printf(" %s - %s\n", menu.Key, menu.Title)
	}
	fmt.Println("")
}

// GetMenuChoice waits for user input of a menu and returns the value
func (cli *Cli) GetMenuChoice() *MenuEntry {
	reader := bufio.NewReader(os.Stdin)
	replacer := strings.NewReplacer("\n", "", "\r", "")
	for {
		fmt.Print("Choose option: ")

		input, readErr := reader.ReadString('\n')
		key := replacer.Replace(input)

		if readErr == nil && cli.IsValidOption(key) {
			for _, m := range cli.Menu {
				if m.Key == key {
					return m
				}
			}
			break
		} else {
			fmt.Println("Invalid option")
		}
	}
	return nil
}

// Run is the main routine of the command line interface
func (cli *Cli) Run() {

	for {
		cli.ShowMenu()

		entry := cli.GetMenuChoice()
		if entry != nil {
			fmt.Printf("%s\n", entry.Title)

			start := time.Now()
			entry.Action(cli)

			fmt.Printf("%v elapsed\n\n", time.Since(start))
		}

		if cli.doQuitFlag {
			break
		}
	}

}

// IsValidOption returns true if n is a valid selection in the current menu of cli
func (cli *Cli) IsValidOption(key string) bool {
	for _, item := range cli.Menu {
		if item.Key == key {
			return true
		}
	}
	return false
}

func menuOptionRenderAllBlocks(cli *Cli) {
	blockTargetPath := "./results/images"
	blockRenderResX := 500
	blockRenderResY := 500

	bf := blockfactory.Get()
	blocks := bf.GetAll()
	graphics.RenderAll(blocks, blockTargetPath, blockRenderResX, blockRenderResY)

	fmt.Printf("Rendered %d blocks to path %s at resolution %dx%d\n", len(blocks), blockTargetPath, blockRenderResX, blockRenderResY)
}

func menuOptionCalcSolutionStatistics(cli *Cli) {
	statsFile := "./results/solutions.csv"

	cf := cardfactory.Get()
	game.CreateSolutionStatistics(cf, statsFile)

	fmt.Printf("Calculated solution statics and stored these in file %s\n", statsFile)
}

func menuOptionGenerateInsaneProblems(cli *Cli) {
	targetDifficulty := card.Insane
	sourceDifficulty := card.Easy
	height := 3
	blockCount := 5

	t := time.Now()
	resultFile := fmt.Sprintf("./results/cards/%s_%s-%02d%02d%02d.txt", targetDifficulty, t.Format("20060102"), t.Hour(), t.Minute(), t.Second())

	bf := blockfactory.Get()
	cf := cardfactory.Get()

	// generate card-set for each animal
	cards := make([]*card.C, 0)
	for _, animal := range card.AllAnimals() {
		cards = append(cards, game.GenerateCardSet(cf, bf, animal, sourceDifficulty, targetDifficulty, height, blockCount, "")...)
	}

	// write all to one file
	totalProblems := 0
	f, err := os.Create(resultFile)
	if err != nil {
		fmt.Printf("Error opening file %s for writing, aborted\n", resultFile)
	} else {
		defer f.Close()
		for _, c := range cards {
			f.WriteString(c.VerbousString())
			totalProblems += len(c.Problems)
		}
	}

	fmt.Printf("Generated %d cards with %d problems and saved to %s", len(cards), totalProblems, resultFile)
}

func menuOptionVisualizeSolution(cli *Cli) {
	cf := cardfactory.Get()
	gs := game.New(cf.Get(card.Difficult, 1).Problems[1]).Solve()[0]
	graphics.Visualize(gs)
}

func menuOptionQuit(cli *Cli) {
	cli.doQuitFlag = true
}
