declare module "*.svg" {
	const content: string;
	export default content;
}

declare module "*.json" {
	const content: object;
	export default content;
}

declare module "*.wav" {
	const content: string;
	export default content;
}

declare module "*.mp3" {
	const content: string;
	export default content;
}

declare module "*.vue" {
	import Vue from 'vue';
	export default Vue;
}
