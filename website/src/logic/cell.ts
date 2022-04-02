import iconDragon from "../assets/icons/dragon.png";
import iconFire from "../assets/icons/fire.svg";
import iconAir from "../assets/icons/air.svg";
import iconPoint from "../assets/icons/point.svg";

export enum CellType {
	Undefined,
	Air,
	Dragon,
	Fire,

	Point
}

export interface CellDefinition {
	value: number;
	name: CellType;
	desc?: string;
	symbol: string;
	isDefined: boolean;
	icon?: string | null;
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
		name: CellType.Air,
		desc: "air",
		icon: iconAir,
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

export const getCellType = (name: CellType) => cellDefinitions.find(x => x.name === name);
export const getCellTypeByValue = (value: number) => cellDefinitions.find(x => x.value === value);
export const getCellTypeBySymbol = (symbol: string) => cellDefinitions.find(x => x.symbol === symbol);

export class Cell {
	readonly id: number;
	readonly given: boolean;
	value: number;
	neighbours: Cell[] = [];

	constructor(id: number, symbol: string, notGiven: boolean = false) {
		const definition = getCellTypeBySymbol(symbol)!;
		this.id = id;
		this.value = definition.value;
		this.given = (definition.isDefined && !notGiven);
	}

	get type() {
		return getCellTypeByValue(this.value);
	}

	get isDefined() {
		return getCellTypeByValue(this.value)!.isDefined;
	}

	get cantBeDragon() {
		return this.neighbours.some(x => x.type!.name === CellType.Dragon);
	}

	get icon() {
		return getCellTypeByValue(this.value)?.icon;
	}
}
