import grids from "../assets/data/grids.json";
import { cellCoords } from "./util";
import { Grid } from "./grid";

enum GameStatus {
	Unsolved,
	Invalid,
	Solved
}

class Game {
	readonly width: number;
	readonly height: number;
	readonly puzzle: Grid;
	readonly solution: Grid;

	constructor(size: number, puzzle: string, solution: string) {
		this.width = size;
		this.height = size;
		this.puzzle = new Grid(puzzle);
		this.solution = new Grid(solution);
	}

	private get isValid() {
		return !cellCoords(this.width, this.height)
			.some(({ x, y }) => {
				const puzzleCell = this.puzzle.getCell(x, y);
				const solutionCell = this.solution.getCell(x, y);
				return puzzleCell.isDefined && puzzleCell.value !== solutionCell.value;
			});
	}

	private get isSolved() {
		return cellCoords(this.width, this.height)
			.every(({ x, y }) => this.puzzle.getCell(x, y).value === this.solution.getCell(x, y).value);
	}

	get status() {
		if (!this.isValid) return GameStatus.Invalid;
		if (this.isSolved) return GameStatus.Solved
		return GameStatus.Unsolved;
	}
}

enum Difficulty {
	Easy = "easy",
	Medium = "medium",
	Hard = "hard"
}

function createGame(difficulty: Difficulty = Difficulty.Easy, size: number = 8) {
	const chooseRandom = (arr) => arr[~~(Math.random() * arr.length)];
	const filteredGrids = grids[difficulty].filter(x => x.size === size);
	const chosen = chooseRandom(filteredGrids);
	return new Game(chosen.size, chosen.puzzle, chosen.solution);
}

export {
	Game,
	GameStatus,
	Difficulty,

	createGame
}
