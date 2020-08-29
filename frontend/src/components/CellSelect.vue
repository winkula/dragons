<template>
  <div class="grid-select grid">
    <div class="grid-row">
      <Cell
        v-for="button in buttons"
        :key="button.id"
        :id="button.id"
        :value="button.value"
        :icon="button.icon"
        :selected="button.value === value"
        @clicked="updateValue(button)"
      ></Cell>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Cell from "./Cell.vue";

import { getCellType, CellType } from "../logic";

export default Vue.extend({
  components: {
    Cell,
  },
  props: {
    value: Number,
  },
  data() {
    return {
      buttons: [
        {
          id: 0,
          value: getCellType(CellType.Empty).value,
          icon: getCellType(CellType.Empty).icon,
        },
        {
          id: 1,
          value: getCellType(CellType.Dragon).value,
          icon: getCellType(CellType.Dragon).icon,
        },
        {
          id: 2,
          value: getCellType(CellType.Fire).value,
          icon: getCellType(CellType.Fire).icon,
        },
      ],
    };
  },
  methods: {
    updateValue(button) {
      this.$emit("input", button.value);
    },
  },
});
</script>