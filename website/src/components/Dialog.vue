<template>
  <div class="dialog" v-if="value">
    <div class="dialog-inner">
      <div class="close" @click="close" v-if="closable">&#10006;</div>
      <div class="dialog-content">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GridComponent from "./Grid.vue";
import { Grid } from "../logic/grid";

export default defineComponent({
  components: {
    Grid: GridComponent,
  },
  props: {
    value: Boolean,
    closable: Boolean,
  },
  methods: {
    close() {
      this.$emit("input", false);
    },
  },
});
</script>

<style lang="scss" scoped>
@import "./src/assets/styles/globals";

.dialog {
  position: absolute;
  left: 0;
  right: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);

  display: flex;
  overflow: auto;
  padding: 1rem;
  box-sizing: border-box;

  .dialog-content {
    padding: 1rem;
  }

  .dialog-inner {
    margin: auto;
    box-sizing: border-box;
    background-color: $color-background-inner;
    border-radius: 1.5vmin;
    box-shadow: 0 3px 10px rgba(0, 0, 0, 0.5);

    .close {
      $size: 3rem;

      float: right;
      width: $size;
      height: $size;
      line-height: $size;
      font-size: $size * 0.5;
      text-align: center;
      vertical-align: middle;
      @include interactive;
    }
  }
}
</style>