<template>
  <Dialog closable :value="value" @input="$emit('input', $event)">
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
      Every dragon has its own territory - the eight squares surrounding him.
      <strong>Inside one's territory there can't be other dragons</strong>. The game automatically marks territory squares with a point
      <Grid class="grid-inline" :grid="point" small></Grid>.
    </p>
    <Grid :grid="example1" small></Grid>

    <h2>The fight rule</h2>
    <p>
      Dragons don't like each other. That's why squares of
      <strong>overlapping territories must always be fire</strong> - but only then.
    </p>
    <Grid :grid="example2" small></Grid>

    <h2>The survive rule</h2>
    <p>
      Dragons like it hot - but they also need air to survive.
      That's why
      <strong>at least two</strong> of the four
      <strong>directly adjacent squares</strong> of a dragon
      <strong>must be empty</strong>.
      Squares outside the grid don't count as "empty".
    </p>
    <p>In this example, the survive rule is satisfied - two of the four directly adjacent squares are empty:</p>
    <Grid :grid="example3_1" small></Grid>
    <p>Here, the survive rule is violated - only one of the two directly adjacent squares are empty:</p>
    <Grid :grid="example3_2" small></Grid>
    <p class="copyright">© {{ year }} Mathias Winkler</p>
  </Dialog>
</template>

<script lang="ts">
import Vue from "vue";
import GridComponent from "./Grid.vue";
import Dialog from "./Dialog.vue";
import { Grid } from "../logic/grid";

export default Vue.extend({
  components: {
    Grid: GridComponent,
    Dialog,
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
      example3_1: new Grid("_x_,xdf,_f_", true),
      example3_2: new Grid("df_,x__,___", true),
    };
  },
  computed: {
    year() {
      return new Date().getFullYear();
    }
  }
});
</script>
