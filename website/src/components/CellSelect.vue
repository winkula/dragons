<template>
  <div class="grid-select grid grid-interactive">
    <div class="grid-row">
      <Cell
        v-for="button in buttons"
        :key="button.id"
        :id="button.id"
        :value="button.value"
        :icon="button.icon"
        :selected="button.value === modelValue"
        @clicked="updateValue(button)"
      ></Cell>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Cell from "./Cell.vue";

import { getCellType, CellType } from "../logic";
import { playClick } from "../sfx";

interface Button {
  id: number;
  value: number;
  icon: number;
}

export default defineComponent({
  components: {
    Cell,
  },
  props: {
    modelValue: Number,
  },
  data() {
    return {
      buttons: [
        {
          id: 0,
          value: getCellType(CellType.Air)!.value,
          icon: getCellType(CellType.Air)!.icon,
        },
        {
          id: 1,
          value: getCellType(CellType.Dragon)!.value,
          icon: getCellType(CellType.Dragon)!.icon,
        },
        {
          id: 2,
          value: getCellType(CellType.Fire)!.value,
          icon: getCellType(CellType.Fire)!.icon,
        },
        {
          id: 3,
          value: getCellType(CellType.Point)!.value,
          icon: getCellType(CellType.Point)!.icon,
        },
      ],
    };
  },
  methods: {
    updateValue(button: Button) {
      playClick();
      this.$emit("update:modelValue", button.value);
    },
  },
});
</script>