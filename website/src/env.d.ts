/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
  const component: DefineComponent<{}, {}, any>
  export default component
}

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
