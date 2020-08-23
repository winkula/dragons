import iconDragon from "../assets/dragon.svg";
import iconFire from "../assets/fire.svg";
import iconEmpty from "../assets/empty.svg";

const cellMap = [
	{ type: "undefined", icon: null },
	{ type: "dragon", icon: iconDragon },
	{ type: "fire", icon: iconFire },
	{ type: "empty", icon: iconEmpty }
];

function buildCell(value: number, x: number, y: number) {

	function changeValueAndRerender() {
		setCell(x, y, (value + 1) % cellMap.length);
		rerender();
	}

	const isStatic = value > 9;
	const cellInfo = cellMap[value % 10];
	const cell = document.createElement("div");
	cell.id = `cell-${x}-${y}`;
	cell.classList.add(cellInfo.type);
	if (isStatic) {
		cell.classList.add("static");
	} else {
		cell.setAttribute("tabindex", "0");
		cell.setAttribute("role", "button");
		cell.setAttribute("aria-pressed", "true");
		cell.addEventListener("click", () => {
			changeValueAndRerender();
		});
		cell.addEventListener("keypress", (e) => {
			if (e.type === 'keypress' && e.keyCode == 13) {
				changeValueAndRerender();
			}
		});
	}
	if (cellInfo.icon) {
		const img = document.createElement("img");
		img.src = cellInfo.icon;
		img.alt = cellInfo.type;
		cell.appendChild(img);
	}
	return cell;
}

function buildRow(grid: Grid, values: number[], y: number) {
	const row = document.createElement("div");
	let x = 0;
	for (const value of values) {
		const cell = buildCell(value, x, y);
		row.appendChild(cell);
		x++;
	}
	return row;
}

function buildGrid(grid: Grid) {
	const root = document.getElementById("grid");
	root.innerHTML = ""; // clear
	for (let y = 0; y < grid.height; y++) {
		const values = grid.cells.slice(y * grid.width, (y + 1) * grid.width);
		const row = buildRow(grid, values, y);
		root.appendChild(row);
	}
}

interface Grid {
	width: number;
	height: number;
	cells: number[];
}

let grid = {
	width: 8,
	height: 8,
	cells: [
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 12, 0, 0, 0, 0,
		0, 0, 11, 12, 11, 0, 0, 0,
		0, 0, 0, 12, 0, 0, 0, 11,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 11, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 12, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
	]
}

function setCell(x: number, y: number, value: number) {
	grid.cells[grid.width * y + x] = value;
}

function rerender() {
	buildGrid(grid);
}

rerender();
