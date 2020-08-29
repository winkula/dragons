// const range = (n: number, start = 0) => Array.from(Array(n).keys()).map(x => x + start);
// const minusOneToOne = range(3, -1);
// const cartesian = <T>(a: T[], b: T[]): T[][] => [].concat(...a.map(d => b.map(e => [].concat(d, e))));

// const pairIsNotOrigin = (pair: number[]) => (pair[0] !== 0 || pair[1] !== 0);
// const neighbourDeltas = cartesian(minusOneToOne, minusOneToOne).filter(pairIsNotOrigin);

// console.log(neighbourDeltas);


function* iterateCells(width: number, height: number) {
	for (let x = 0; x < width; x++) {
		for (let y = 0; y < height; y++) {
			yield { x: x, y: y };
		}
	}
}

function cellCoords(width: number, height: number) {
	return Array.from(iterateCells(width, height));
}

export {
	cellCoords
}
