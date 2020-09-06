<h1 align="center">
<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/logo.png" alt="Logo" width="128"/>
<br/>
Dragons
</h1>

<div align="center">

[![Build Status](https://img.shields.io/github/workflow/status/winkula/dragons/Go)](https://github.com/winkula/dragons/actions)
[![Code coverage](https://img.shields.io/codecov/c/github/winkula/dragons)](https://codecov.io/github/winkula/dragons)
![](https://img.shields.io/github/go-mod/go-version/winkula/dragons)

</div>

## Introduction

Dragons is a logic-based puzzle inspired by [Battleship](<https://en.wikipedia.org/wiki/Battleship_(puzzle)>), [Minesweeper](https://en.wikipedia.org/wiki/Microsoft_Minesweeper) and [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).

The objective is to fill a grid - where some squares are given from the start. A square can either be a dragon (`D`), fire (`F`) or empty (`X`).

There are three rules that dictate how the grid muss be completed. Only one valid solution exists for a given puzzle.

## Demo

[Click here](https://dragons-puzzle.netlify.app) to try out the puzzle online. It also works on mobile phones.

## Rules

### The territory rule

Every dragon has its own territory - the eight squares surrounding him.
**Inside one's territory there can't be other dragons**.

```
┌───────┐
│ . . . │
│ . D . │
│ . . . │
└───────┘
```

### The fight rule

Dragons don't like each other. That's why squares of
**overlapping territories must always be fire** - but only then.

```
┌───────────┐
│ . . F . . │
│ . D F D . │
│ . . F . . │
└───────────┘
```

### The survive rule

Dragons like it hot - but they also need air to survive.
That's why **at least two** of the four **directly adjacent squares** of a dragon **must be empty**.
Squares outside the grid don't count as "empty".

In this example, the survive rule is satisfied - two of the four directly adjacent squares are empty:

```
┌───────┐
│ . X . │
│ X D F │
│ . F . │
└───────┘
```

Here, the survive rule is violated - only one of the two directly adjacent squares are empty:

```
┌───────┐
│ D F . │
│ X . . │
│ . . . │
└───────┘
```

## ToDos

- [ ] Implement text/PDF generation function for the CLI
