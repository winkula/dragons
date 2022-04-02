<template>
  <div
    :class="['grid-cell', { given: given || selected }, cellType]"
    :tabindex="tabindex"
    :role="role"
    :aria-pressed="ariaPressed"
    @click="clicked"
  >
    <img :src="icon" alt="grid cell" v-if="icon != null" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { getCellTypeByValue } from "../logic";

export default defineComponent({
  props: {
    id: Number,
    value: Number,
    given: Boolean,
    selected: Boolean,
    icon: String,
  },
  computed: {
    tabindex() {
      return this.given ? null : "0";
    },
    ariaPressed() {
      return this.given ? null : "true";
    },
    role() {
      return this.given ? null : "button";
    },
    cellType() {
      return getCellTypeByValue(this.value!)?.desc;
    },
  },
  methods: {
    clicked() {
      this.$emit("clicked", {
        id: this.id,
        value: this.value,
        given: this.given,
        selected: this.selected,
      });
    },
  },
});
</script>