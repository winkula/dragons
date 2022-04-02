import soundClick from "../assets/audio/click.mp3";
import soundError from "../assets/audio/error.mp3";
import soundWin from "../assets/audio/win.mp3";

export function playClick() {
	const sound = new Audio(soundClick);
	sound.volume = 0.4;
	sound.play();
}

export function playError() {
	new Audio(soundError).play();
}

export function playWin() {
	new Audio(soundWin).play();
}
