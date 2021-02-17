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

_Dragons_ is a logic-based puzzle inspired by [Battleship](<https://en.wikipedia.org/wiki/Battleship_(puzzle)>), [Minesweeper](https://en.wikipedia.org/wiki/Microsoft_Minesweeper) and [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).

The objective is to fill a grid - some squares are given from the start. A square can either be a dragon (<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/dragon.png" alt="dragon" width="16"/>), fire (<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/fire.png" alt="fire" width="16"/>) or air (<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/air.png" alt="air" width="16"/>).

There are three rules that dictate how the grid must be completed. Only one valid solution exists for a given puzzle.





## Rules

### The territory rule

Every dragon has its own territory (the eight squares surrounding him). **Inside one's territory there can't be other dragons**.
You can mark squares where dragons are impossible with a point (<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/point.png" alt="air" width="16"/>).


<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/rule1.png" alt="illustration of the territory rule" />

### The fire rule

Dragons don't like each other and they spit fire when being provoked. That's why squares of **overlapping territories must always be fire** - but only then.

<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/rule2.png" alt="illustration of the fight rule" />

### The survive rule

Dragons need air to survive. That's why **at least two** of the four **directly adjacent squares** (not diagonal) of a dragon **must be air**.

<img src="https://raw.githubusercontent.com/winkula/dragons/master/assets/rule3.png" alt="illustration of the survive rule" />


## Demo

[Click here](https://dragons-puzzle.netlify.app) to try out the puzzle online. It also works on mobile phones.


## Command line

This repository also contains a command line program that can generate and solve these kind of puzzles.
The command line can be used as follows:

```
git clone https://github.com/winkula/dragons.git

go run .\cmd\dragons\
```
