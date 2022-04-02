<template>
  <Dialog
    closable
    :modelValue="modelValue"
    @update:modelValue="$emit('update:modelValue', $event)"
  >
    <h1>Rules</h1>

    <p>
      This is a puzzle about dragons. Try to find out, where the dragons are
      hiding on the grid. A square can either be a dragon
      <Grid class="grid-inline" :grid="dragon" small></Grid>, fire
      <Grid class="grid-inline" :grid="fire" small></Grid>or air
      <Grid class="grid-inline" :grid="air" small></Grid>. The grid must be
      filled completely in order to win the game.
    </p>
    <p>There are only three rules:</p>

    <h2>The territory rule</h2>
    <p>
      Every dragon has its own territory (the eight squares surrounding him).
      <strong>Inside one's territory there can't be other dragons</strong>. You
      can mark squares where dragons are impossible with a point
      <Grid class="grid-inline" :grid="point" small></Grid>.
    </p>
    <Grid :grid="example1" small></Grid>

    <h2>The fire rule</h2>
    <p>
      Dragons don't like each other and they spit fire when being provoked.
      That's why squares of
      <strong>overlapping territories must always be fire</strong> - but only
      then.
    </p>
    <Grid :grid="example2" small></Grid>

    <h2>The survive rule</h2>
    <p>
      Dragons need air to survive. That's why
      <strong>at least two</strong> of the four
      <strong>directly adjacent squares</strong> (not diagonal) of a dragon
      <strong>must be air</strong>.
    </p>
    <Grid :grid="example3" small></Grid>
    <p class="copyright">© 2021 Mathias Winkler</p>
  </Dialog>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import GridComponent from "./Grid.vue";
import Dialog from "./Dialog.vue";
import { Grid } from "../logic/grid";

export default defineComponent({
  components: {
    Grid: GridComponent,
    Dialog,
  },
  props: {
    modelValue: Boolean,
  },
  data() {
    return {
      dragon: new Grid("d", true),
      fire: new Grid("f", true),
      air: new Grid("x", true),
      given: new Grid("d"),
      point: new Grid(".", true),

      example1: new Grid("...,.d.,...", true),
      example2: new Grid("_____,_df__,__fd_", true),
      example3: new Grid("dx_,x__,___", true),
    };
  },
});
</script>
