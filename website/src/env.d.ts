/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
  const component: DefineComponent<{}, {}, any>
  export default component
}

/*
declare module "*.json" {
  const json: object;
  export default json;
}

declare module "*.svg" {
  const svg: string;
  export default svg;
}

declare module "*.wav" {
  const wav: string;
  export default wav;
}

declare module "*.mp3" {
  const mp3: string;
  export default mp3;
}
*/