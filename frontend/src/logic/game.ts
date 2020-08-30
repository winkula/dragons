import grids from "../assets/data/grids.json";
import { cellCoords, randomInteger } from "./util";
import { Grid } from "./grid";

enum GameStatus {
	Unsolved = "unsolved",
	Invalid = "invalid",
	Solved = "solved"
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

function createGameBuilder() {
	let index = randomInteger(8);

	return function (difficulty: Difficulty = Difficulty.Easy, size: number = 8) {
		const filteredGrids = grids[difficulty].filter(x => x.size === size);
		index = (index + 1) % filteredGrids.length; // wrap around
		const chosen = filteredGrids[index];
		return new Game(chosen.size, chosen.puzzle, chosen.solution);
	}
}

const builder = createGameBuilder();

function createGame(difficulty: Difficulty = Difficulty.Easy, size: number = 8) {
	return builder(difficulty, size);
}

export {
	Game,
	GameStatus,
	Difficulty,

	createGame
}
