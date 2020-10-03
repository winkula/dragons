import soundClick from "../assets/audio/click.mp3";
import soundMusic from "../assets/audio/music.mp3";
import soundError from "../assets/audio/error.mp3";
import soundWin from "../assets/audio/win.mp3";

const music = new Audio(soundMusic);
music.loop = true;
music.volume = 0.4;

function playClick() {
	const sound = new Audio(soundClick);
	sound.volume = 0.4;
	sound.play();
}

function playMusic() {
	music.play();
}

function playError() {
	new Audio(soundError).play();
}

function playWin() {
	new Audio(soundWin).play();
}

export {
	playClick,
	playMusic,
	playError,
	playWin
}
