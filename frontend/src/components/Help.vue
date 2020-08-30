<template>
  <div :class="['help', { 'visible': value } ]">
    <div class="close" @click="close">&#10006;</div>
    <div>
      <h1>Rules</h1>

      <p>
        This is a puzzle about dragons. Try to find out, where the dragons are hiding on the grid. A square can either be a dragon
        <Grid class="grid-inline" :grid="dragon" small></Grid>, fire
        <Grid class="grid-inline" :grid="fire" small></Grid>or empty
        <Grid class="grid-inline" :grid="empty" small></Grid>.
      </p>
      <p>There are only three rules:</p>

      <h2>The territory rule</h2>
      <p>
        Every dragon has its own territory - the eight squares surrounding him. <strong>Inside ones territory there can't be other dragons</strong>. The game automatically marks territory squares with a point
        <Grid class="grid-inline" :grid="point" small></Grid>.
      </p>
      <Grid :grid="example1" small></Grid>

      <h2>The fight rule</h2>
      <p>Dragons don't like each other. That's why squares of <strong>overlapping territories must always be fire</strong> - but only then.</p>
      <Grid :grid="example2" small></Grid>

      <h2>The survive rule</h2>
      <p>
        Dragons like it hot - but they also need air to survive.
        That's why <strong>at least two</strong> of the four <strong>directly adjacent squares</strong> of a dragon <strong>must be empty</strong>.
        Squares outside the grid don't count as "empty".
      </p>
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

  display: none;
  opacity: 0;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  padding: 1rem;
  box-sizing: border-box;
  background-color: #fff;

  &.visible {
    display: block;
    transition: all 5s;
    opacity: 1;
  }

  .close {
    float: right;
    width: 7vmin;
    height: 7vmin;
    text-align: center;
    vertical-align: middle;
    @include interactive;
  }
}
</style>