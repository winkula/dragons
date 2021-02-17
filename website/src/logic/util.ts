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

const randomInteger = (length: number) => Math.floor(Math.random() * length);

export {
	cellCoords,
	randomInteger
}
