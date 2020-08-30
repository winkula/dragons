<template>
  <div :class="['help', { 'visible': value } ]">
    <div class="close" @click="close">&#10006;</div>
    <div>
      <h1>Rules</h1>

      <p>
        This is a puzzle about dragons. Try to find out, where the dragons are hiding on the grid. A square can eighter be a dragon
        <Grid class="grid-inline" :grid="dragon" small></Grid>, fire
        <Grid class="grid-inline" :grid="fire" small></Grid>or empty
        <Grid class="grid-inline" :grid="empty" small></Grid>.
      </p>
      <p>There are only 3 rules:</p>

      <h2>The territory rule</h2>
      <p>
        Dragons can not have other dragons in their territory. Territory means the 8 squares that are around itself. Squares where no dragon can be are marked with a point
        <Grid class="grid-inline" :grid="point" small></Grid>. This is a automatic hint of the game.
      </p>
      <Grid :grid="example1" small></Grid>

      <h2>The fight rule</h2>
      <p>Overlapping territories must be fire. Every square that is part of multiple territories must be fire - and only then.</p>
      <Grid :grid="example2" small></Grid>

      <h2>The survive rule</h2>
      <p>At least 2 of the adjacent squares of a dragon must be empty. If a dragon is at the edge or in the corner of the grid, the possible number of adjacent squares is reduced.</p>
      <Grid :grid="example3" small></Grid>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import GridComponent from "./Grid.vue";
import { Grid } from "../logic/grid";

export default Vue.extend({
  components: {
    Grid: GridComponent,
  },
  props: {
    value: Boolean,
  },
  data() {
    return {
      dragon: new Grid("d", true),
      fire: new Grid("f", true),
      empty: new Grid("x", true),
      given: new Grid("d"),
      point: new Grid(".", true),

      example1: new Grid("___,_d_,___", true),
      example2: new Grid("_____,_df__,__fd_,_____", true),
      example3: new Grid("xfx,dfd,xfx", true),
    };
  },
  computed: {
    styles() {
      return {
        display: "block",
      };
    },
  },
  methods: {
    close() {
      this.$emit("input", false);
    },
  },
});
</script>

<style lang="scss">
@import "./src/assets/styles/globals";

.help {
  h1 {
    margin-top: 0;
  }

  p {
    margin-block-start: 0.5rem;
    margin-block-end: 0.5rem;
  }

  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  padding: 1rem;
  box-sizing: border-box;
  display: none;
  background-color: #fff;

  &.visible {
    display: block;
  }

  .close {
    float: right;
    @include interactive;
    width: 7vmin;
    height: 7vmin;
    text-align: center;
    vertical-align: middle;
  }
}
</style>