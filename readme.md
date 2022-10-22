# Ubongo

## Introduction

This project was created as a way to learn the language Go. It deals with the board game 'Ubongo' - specifically the 3D-edition by company Kosmos (see more here <https://www.kosmos.de/spielware/spiele/ubongo/>). 

**Note**: this is a purely private project, which is in no way affilicated with company Kosmos.

The goal of the game is to fill a volume with a given blueprint and height of 2 levels with 3 or 4 Tetris-like blocks. Unlike in Tetris, some of the blocks have non-flat shapes to make the game more challenging.

## Features

In this repo the original problems of the game are digitally reproduced and the code allows for creating new problems. Specifically:

- A solver finds all solutions to given problems
- New problems can be automatically created, in particular such with a higher difficulty
- Solutions can be rendered using simple 3D graphic

## Notes

Run coverage tests as follows:

```text
go test ./... -coverprofile=coverage
go tool cover -html=coverage
```

Create struct-dependency graph:

```text
embedded-struct-visualizer -out dependencies.dot
```

then visualize it in VS-Code by selecting the file and pressing Ctrl+Shift+V
