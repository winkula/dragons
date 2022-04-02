<template>
  <Dialog v-model="value">
    <p class="difficulty-text">Select the difficulty:</p>
    <Buttons :buttons="buttons" @click="$emit('start', $event)"></Buttons>
  </Dialog>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Dialog from "./Dialog.vue";
import Buttons from "./Buttons.vue";

export default defineComponent({
  components: {
    Dialog,
    Buttons,
  },
  props: {
    value: Boolean,
  },
  data() {
    return {
      buttons: [
        {
          id: "easy",
          icon: `
<svg viewBox="0 0 24 24">
  <circle cx="12" cy="12" r="10" />
  <path d="M8 14s1.5 2 4 2 4-2 4-2" />
  <line x1="9" y1="9" x2="9.01" y2="9" />
  <line x1="15" y1="9" x2="15.01" y2="9" />
</svg>`,
        },
        {
          id: "medium",
          icon: `
<svg viewBox="0 0 24 24">
  <circle cx="12" cy="12" r="10" />
  <line x1="8" y1="15" x2="16" y2="15" />
  <line x1="9" y1="9" x2="9.01" y2="9" />
  <line x1="15" y1="9" x2="15.01" y2="9" />
</svg>
        `,
        },
        {
          id: "hard",
          icon: `
<svg viewBox="0 0 24 24">
  <circle cx="12" cy="12" r="10" />
  <path d="M16 16s-1.5-2-4-2-4 2-4 2" />
  <line x1="9" y1="9" x2="9.01" y2="9" />
  <line x1="15" y1="9" x2="15.01" y2="9" />
</svg>`,
        },
      ],
    };
  },
});
</script>

<style lang="scss">
.difficulty-text {
  text-align: center;
}
</style>
