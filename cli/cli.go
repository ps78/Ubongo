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
	Menu []*MenuEntry

	// if this flag is set to true, the program will terminate as soon as possible
	doQuitFlag bool
}

// MenuEntry represents an entry in the menu
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

		if readErr == nil {
			m := cli.GetMenuEntry(key)
			if m != nil {
				return m
			} else {
				fmt.Println("Invalid option")
			}
		}

		// abort loop if termination was signalled
		if cli.doQuitFlag {
			return nil
		}
	}
}

// readProblem reads a valid problem id from the command line
// it must be entered as 'cardnumber difficulty dicenumber'
func (cli *Cli) readProblem(cf *cardfactory.F) (int, card.UbongoDifficulty, int) {
	reader := bufio.NewReader(os.Stdin)
	replacer := strings.NewReplacer("\n", "", "\r", "")
	for {
		fmt.Print("Enter problem (CardNumber Difficulty DiceNumber): ")

		input, readErr := reader.ReadString('\n')

		if readErr == nil {
			input = replacer.Replace(input)
			parts := strings.Split(input, " ")
			if len(parts) == 3 {
				cardNumber, err := strconv.Atoi(parts[0])
				if err == nil && cardNumber >= 1 && cardNumber <= 36 {

					difficulty, err := card.ParseDifficulty(parts[1])
					if err == nil {

						diceNumber, err := strconv.Atoi(parts[2])
						if err == nil && diceNumber >= 1 && diceNumber <= 10 {
							card := cf.Get(difficulty, cardNumber)
							if card != nil {
								if _, ok := card.Problems[diceNumber]; ok {
									return cardNumber, difficulty, diceNumber
								} else {
									fmt.Println("Problem does not exist on this card")
								}
							} else {
								fmt.Println("Card does not exist")
							}
						} else {
							fmt.Println("Invalid dice number, must be a number from 1..10")
						}
					} else {
						fmt.Println("Invalid difficulty, must be Easy or Difficult")
					}
				} else {
					fmt.Println("Invalid card number, must be a number from 1..36")
				}
			}
		}

		// abort loop if termination was signalled
		if cli.doQuitFlag {
			break
		}
	}
	return 0, card.Easy, 0
}

// Run is the main routine of the command line interface and runs in a loop until terminated
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

		// check for termination flag
		if cli.doQuitFlag {
			break
		}
	}
}

// Returns the menu entry with the given key, or nil if it not exists
func (cli *Cli) GetMenuEntry(key string) *MenuEntry {
	for _, item := range cli.Menu {
		if item.Key == key {
			return item
		}
	}
	return nil
}

func menuOptionRenderAllBlocks(cli *Cli) {
	blockTargetPath := "./results/images"
	blockRenderResX := 500
	blockRenderResY := 500

	bf := blockfactory.Get()
	blocks := bf.GetAll()
	graphics.RenderBlockset(blocks, blockTargetPath, blockRenderResX, blockRenderResY)

	fmt.Printf("Rendered %d blocks to path %s at resolution %dx%d\n", blocks.Count, blockTargetPath, blockRenderResX, blockRenderResY)
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

	fmt.Printf("Generating problems with height %d and %d blocks based on layouts of %s cards\n", height, blockCount, sourceDifficulty)

	t := time.Now()
	resultFile := fmt.Sprintf("./results/cards/%s_%s-%02d%02d%02d.txt", targetDifficulty, t.Format("20060102"), t.Hour(), t.Minute(), t.Second())

	bf := blockfactory.Get()
	cf := cardfactory.Get()

	// generate card-set for each animal
	cards := make([]*card.C, 0)
	for _, animal := range card.AllAnimals() {
		newCards := game.GenerateCardSet(cf, bf, animal, sourceDifficulty, targetDifficulty, height, blockCount, "")
		cards = append(cards, newCards...)
		newProbCount := 0
		for _, c := range newCards {
			newProbCount += len(c.Problems)
		}
		fmt.Printf("Created %d cards with %d problems for animal %s\n", len(newCards), newProbCount, animal)
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

	fmt.Printf("Generated %d cards with %d problems and saved to %s\n", len(cards), totalProblems, resultFile)
}

func menuOptionVisualizeSolution(cli *Cli) {
	solutionNumber := 0

	cf := cardfactory.Get()
	cardNumber, difficulty, diceNumber := cli.readProblem(cf)

	sols := game.New(cf.Get(difficulty, cardNumber).Problems[diceNumber]).Solve()
	gs := sols[solutionNumber]

	fmt.Printf("Showing solution of %s problem on card %d, dice %d (solution %d out of %d)\n",
		difficulty, cardNumber, diceNumber, solutionNumber+1, len(sols))

	graphics.Visualize(gs, 800, 600)
	cli.doQuitFlag = true // we cannot continue with the cli because the Fyne-App is not reusalbe once closed
}

func menuOptionQuit(cli *Cli) {
	cli.doQuitFlag = true
}
