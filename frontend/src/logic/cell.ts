import iconDragon from "../assets/icons/dragon-dark.png";
import iconFire from "../assets/icons/fire.png";
import iconEmpty from "../assets/icons/empty.svg";
import iconPoint from "../assets/icons/point.svg";

enum CellType {
	Undefined,
	Empty,
	Dragon,
	Fire,

	Point
}

interface CellDefinition {
	value: number;
	name: CellType;
	desc?: string;
	symbol: string;
	isDefined: boolean;
	icon?: string;
}

const cellDefinitions: CellDefinition[] = [
	{
		value: 0,
		name: CellType.Undefined,
		icon: null,
		symbol: "_",
		isDefined: false
	},
	{
		value: 1,
		name: CellType.Empty,
		desc: "empty",
		icon: iconEmpty,
		symbol: "x",
		isDefined: true
	},
	{
		value: 2,
		name: CellType.Dragon,
		desc: "dragon",
		icon: iconDragon,
		symbol: "d",
		isDefined: true
	},
	{
		value: 3,
		name: CellType.Fire,
		desc: "fire",
		icon: iconFire,
		symbol: "f",
		isDefined: true
	},
	{
		value: 4,
		name: CellType.Point,
		desc: "point",
		icon: iconPoint,
		symbol: ".",
		isDefined: false
	}
];

const getCellType = (name: CellType) => cellDefinitions.find(x => x.name === name);
const getCellTypeByValue = (value: number) => cellDefinitions.find(x => x.value === value);
const getCellTypeBySymbol = (symbol: string) => cellDefinitions.find(x => x.symbol === symbol);

class Cell {
	readonly id: number;
	readonly given: boolean;
	value: number;
	neighbours: Cell[] = [];

	constructor(id: number, symbol: string, notGiven: boolean = false) {
		const definition = getCellTypeBySymbol(symbol);
		this.id = id;
		this.value = definition.value;
		this.given = (definition.isDefined && !notGiven);
	}

	get type() {
		return getCellTypeByValue(this.value);
	}

	get isDefined() {
		return getCellTypeByValue(this.value).isDefined;
	}

	get cantBeDragon() {
		return this.neighbours.some(x => x.type.name === CellType.Dragon);
	}

	get icon() {
		if (!this.given && !this.isDefined && this.cantBeDragon) {
			return iconPoint;
		}
		return getCellTypeByValue(this.value)?.icon;
	}
}

export {
	CellDefinition,
	CellType,
	Cell,

	getCellType,
	getCellTypeByValue,
	getCellTypeBySymbol
}
