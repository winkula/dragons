<template>
  <div :class="['grid', { 'invalid': !isValid }, {'solved': isSolved}]">
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
  },
  methods: {
    clicked(cell) {
      this.$emit("filled", cell);
    },
  },
});
</script>
