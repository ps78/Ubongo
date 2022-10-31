# Ubongo

## Introduction

This project was created as a way to learn the language Go. It deals with the board game 'Ubongo' - specifically the 3D-edition by company Kosmos (see more here <https://www.kosmos.de/spielware/spiele/ubongo/>).

**Disclamer**: this is a purely private project, which is in no way affilicated with company Kosmos.

The goal of the game is to fill a volume with a given blueprint and height of 2 levels with 3 or 4 Tetris-like blocks. Unlike in Tetris, some of the blocks have non-flat shapes to make the game more challenging.

## Features

In this repo the original problems of the game are digitally reproduced and the code allows for creating new problems. Specifically:

- A solver finds all solutions to given problems
- New problems can be automatically created, in particular such with a higher difficulty
- Solutions can be rendered using simple 3D graphic

Running the main program starts the command line interface:

```text
go run main.go
```

All results generated will be stored in `./results`:

- `./results/solutions.csv`: counts of solutions for all problems of the original game
- `./images/`: contains the wireframe renders of the 16 blocks of the game
- `./cards/`: these are compelete sets of problems for all 36 cards with difficulty level *insane*, i.e. using the shapes of the easy problems but requiring 5 blocks, building 3 levels high instead of 2.

## Dependencies

The following packages are used by the project (you need to install these first):

- Pinhole <https://github.com/tidwall/pinhole>: Allows drawing simple 3D wireframe graphics
- Fyne <https://fyne.io>: Is a fully fletched GUI framework for Go. We only use it to display a solution rendered with Pinhole on screen

## Known Issues and Limitations

This is a command-line application. The GUI framework Fyne is only used to visualize a solution. Once this has been done and the window was closed, it cannot be opened again without restarting the CLI. A future version of the program might use Fyne as a user interface entirely.

## Notes

The project has a fairly good unit test coverage.
Run test coverage analysis as follows:

```text
go test ./... -coverprofile=coverage
go tool cover -html=coverage
```

Create struct-dependency graph:

```text
embedded-struct-visualizer -out dependencies.dot
```

then visualize it in VS-Code by selecting the file and pressing <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>V</kbd>
