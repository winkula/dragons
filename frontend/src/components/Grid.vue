<template>
  <div
    :class="['grid', { 'invalid': !isValid }, {'solved': isSolved}, {'grid-interactive': interactive}, {'grid-small': small}]"
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
import Vue from "vue";

import Cell from "./Cell.vue";

export default Vue.extend({
  components: {
    Cell,
  },
  props: {
    grid: Object,
    isValid: { type: Boolean, default: true },
    isSolved: Boolean,
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
  20% {
    top: 4px;
  }
  100% {
    top: 0px;
  }
}

@mixin cell($size) {
  $padding: $size * 0.1;
  $shadow: $size * 0.1;
  $gap: $size * 0.2;

  width: $size;
  height: $size + $shadow;
  margin: (($gap - $shadow) / 2) ($gap / 2);
  padding: $padding;
  position: relative;
  border-bottom: $shadow solid $color-cell-dark; // 1vmin
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
  border: 1px solid $color-cell-dark;
  outline: none;
  @include notSelectable;

  @include cell(10vmin);
  border-color: $color-cell-dark;

  @media (min-aspect-ratio: 7/10) {
    @include cell(7vmin);
  }

  &.dragon {
    @include shine($color-dragon);
  }

  &.fire {
    @include shine($color-fire);
  }

  &.given {
    background-color: $color-static;
    border-color: $color-static-dark;
    @include disabled;
  }

  &.selected {
    background-color: $color-static;
    border-color: $color-static-dark;
    @include disabled;
  }

  // icons
  & > img {
    width: 100%;
    transition: width 1s;
  }
}

.grid-small {
  .grid-cell {
    @include cell(1.5rem);
  }
}

.grid-interactive {
  .grid-cell {
    @include interactive;

    &:not(.given) {
      &:hover,
      &:focus {
        background-color: lighten($color-cell-dark, 15%);
        border-color: darken($color-cell-dark, 15%);
      }

      :active {
        animation-name: bounce;
        animation-duration: 0.2s;
      }
    }
  }
}
</style>