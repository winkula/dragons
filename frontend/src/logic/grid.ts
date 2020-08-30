import { Cell } from "./cell";
import { cellCoords } from "./util";

interface Row {
	id: number;
	cells: Cell[];
}

class Grid {
	rows: Row[];

	get width() { return this.rows[0]?.cells.length ?? 0; }
	get height() { return this.rows.length; }

	constructor(serialized: string, notGiven = false) {
		this.rows = this.parse(serialized, notGiven);
		this.setNeighbourCells();
	}

	private parse(serialized: string, notGiven: boolean) {
		const rowSeparator = ',';
		const rows = serialized.split(rowSeparator);
		return rows.map((row, rowIndex) => <Row>{
			id: rowIndex,
			cells: row.split('').map((cell, cellIndex) => new Cell(cellIndex, cell, notGiven))
		});
	}

	private setNeighbourCells() {
		const neighbours = [
			{ x: -1, y: -1 },
			{ x: 0, y: -1 },
			{ x: 1, y: -1 },

			{ x: -1, y: 0 },
			{ x: 1, y: 0 },

			{ x: -1, y: 1 },
			{ x: 0, y: 1 },
			{ x: 1, y: 1 },
		];
		for (const { x, y } of cellCoords(this.width, this.height)) {
			const cell = this.getCell(x, y);
			for (const n of neighbours) {
				const neighbourCell = this.getCell(x + n.x, y + n.y);
				if (neighbourCell != null) {
					cell.neighbours.push(neighbourCell);
				}
			}
		}
	}

	getCell(x: number, y: number) {
		return this.rows[y]?.cells[x];
	}
}

export {
	Grid
}
