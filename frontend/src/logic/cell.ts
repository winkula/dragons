import iconDragon from "../assets/icons/dragon.svg";
import iconFire from "../assets/icons/fire.svg";
import iconEmpty from "../assets/icons/empty.svg";
import iconPoint from "../assets/icons/point.svg";

enum CellType {
	Undefined,
	Empty,
	Dragon,
	Fire
}

interface CellDefinition {
	value: number;
	name: CellType;
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
		icon: iconEmpty,
		symbol: "x",
		isDefined: true
	},
	{
		value: 2,
		name: CellType.Dragon,
		icon: iconDragon,
		symbol: "d",
		isDefined: true
	},
	{
		value: 3,
		name: CellType.Fire,
		icon: iconFire,
		symbol: "f",
		isDefined: true
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

	constructor(id: number, symbol: string) {
		const definition = getCellTypeBySymbol(symbol);
		this.id = id;
		this.value = definition.value;
		this.given = definition.isDefined;
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
