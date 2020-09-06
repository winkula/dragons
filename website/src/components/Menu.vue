<template>
  <Buttons :buttons="buttons" @click="$emit($event)"></Buttons>
</template>

<script lang="ts">
import Buttons from "./Buttons.vue";
import Vue from "vue";

export default Vue.extend({
  components: {
    Buttons,
  },
  data() {
    return {
      buttons: [
        {
          id: "new",
          icon: `
<svg viewBox="0 0 24 24">
  <polygon points="5 3 19 12 5 21 5 3" />
</svg>
        `,
        },
        {
          id: "help",
          icon: `
<svg viewBox="0 0 24 24">
  <circle cx="12" cy="12" r="10" />
  <path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3" />
  <line x1="12" y1="17" x2="12.01" y2="17" />
</svg>
        `,
        },
      ],
    };
  },
});
</script>
