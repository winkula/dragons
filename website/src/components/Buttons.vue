<template>
  <div class="buttons">
    <button
      type="button"
      @click="$emit('selected', button.id)"
      v-for="button in buttons"
      :key="button.id"
      v-html="button.icon"
    ></button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  props: {
    buttons: Array,
  },
});
</script>

<style lang="scss">
@import "./src/assets/styles/globals";

.buttons {
  display: flex;
  justify-content: center;

  button {
    margin: 0;
    padding: 0;
    background: none;
    border: 0;
    outline: none;
    line-height: 0;

    &:hover {
      @include interactive;

      svg {
        stroke: $color-button-hover;
      }
    }

    svg {
      fill: none;
      stroke: $color-button;
      stroke-linecap: round;
      stroke-linejoin: round;
      stroke-width: 1.5;

      width: $button-size;
      height: $button-size;
      margin: $button-size * 0.3;

      @media (min-aspect-ratio: 7/10) {
        width: $button-size * 0.7;
        height: $button-size * 0.7;
        margin: $button-size * 0.3 * 0.7;
      }
    }
  }
}
</style>
