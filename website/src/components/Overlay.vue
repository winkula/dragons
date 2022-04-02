<template>
  <div
    class="overlay"
    :class="[{ invalid: invalid }, { solved: solved }]"
    @click="dismiss"
  >
    <svg viewBox="0 0 24 24" v-if="solved">
      <polyline points="20 6 9 17 4 12" />
    </svg>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    state: String,
  },
  computed: {
    solved() {
      return this.state === "solved";
    },
    invalid() {
      return this.state === "invalid";
    },
  },
  methods: {
    dismiss() {
      this.$emit("close");
    },
  },
});
</script>

<style lang="scss">
@import "./src/assets/styles/globals";

@keyframes flyin {
  0% {
    transform: translate(0, -300%);
  }
  100% {
    transform: translate(0, 0);
  }
}

@mixin overlay($color) {
  background-color: rgba($color: $color, $alpha: 0.3);
  stroke: $color;
  transition: background 0.5s;
}

.overlay {
  display: flex;
  justify-content: center;
  align-items: center;

  pointer-events: none;
  overflow: hidden;
  background: none;

  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;

  svg {
    display: none;
    position: absolute;
    fill: none;
    stroke-linecap: round;
    stroke-linejoin: round;
    stroke-width: 1.5;

    width: 30vmin;
    height: 30vmin;
  }

  &.invalid {
    @include overlay(#de4741);
  }

  &.solved {
    @include overlay(green);
    pointer-events: auto;
    backdrop-filter: blur(1px);

    svg {
      display: block;
      animation-name: flyin;
      animation-duration: 1s;
    }
  }
}
</style>