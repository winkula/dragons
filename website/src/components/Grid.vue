<template>
  <div
    :class="[
      'grid',
      { invalid: status === 'invalid' },
      { solved: status === 'solved' },
      { 'grid-interactive': interactive },
      { 'grid-small': small },
    ]"
  >
    <div class="grid-row" v-for="row in grid.rows" :key="row.id">
      <Cell
        v-for="cell in row.cells"
        :key="cell.id"
        :id="cell.id"
        :value="cell.value"
        :given="cell.given"
        :icon="cell.icon"
        @clicked="clicked(cell)"
      ></Cell>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

import Cell from "./Cell.vue";

export default defineComponent({
  components: {
    Cell,
  },
  props: {
    grid: Object,
    status: String,
    interactive: Boolean,
    small: { type: Boolean, default: false },
  },
  methods: {
    clicked(cell) {
      this.$emit("filled", cell);
    },
  },
});
</script>

<style lang="scss">
@import "./src/assets/styles/globals";

@keyframes bounce {
  0% {
    top: 0px;
  }
  30% {
    top: 3px;
  }
  100% {
    top: 0px;
  }
}

@mixin cell($size) {
  $padding: $size * 0; // 0.1
  $shadow: $size * 0.1;
  $gap: $size * 0.2;

  width: $size;
  height: $size + $shadow;
  margin: (calc(($gap - $shadow) / 2)) (calc($gap / 2));
  padding: $padding;
  position: relative;

  //border: 1px solid $color-cell-border;
  border-bottom: $shadow solid $color-cell-border; // 1vmin
  box-shadow: 0 0 10px $color-cell-border;
}

.grid-row {
  display: flex;
  justify-content: center;
}

.grid-inline {
  display: inline-block;
}

.grid-cell {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: $color-cell;
  box-sizing: border-box;
  outline: none;

  @include cell(10vmin);
  border-color: $color-cell-border;

  @media (min-aspect-ratio: 7/10) {
    @include cell(7vmin);
  }

  &.given {
    background-color: $color-cell-given;
    border-color: $color-cell-border-given;
    @include disabled;
  }

  &.selected {
    background-color: $color-cell-given;
    border-color: $color-cell-border-given;
    @include disabled;
  }

  // icons
  & > img {
    width: 100%;
  }
}

.grid-small {
  .grid-cell {
    @include cell(1.5rem);
  }
}

.grid-interactive {
  @include interactive;

  .grid-cell {
    &:not(.given):hover {
      background-color: $color-cell-hover;
      border-color: $color-cell-border-hover;
    }

    &:active {
      animation-name: bounce;
      animation-duration: 0.15s;
    }
  }
}
</style>