<template>
  <div :class="['app-help', { 'visible': visible } ]">
    <div>
      <h1>Rules </h1>

      <p>...</p>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  components: {},
  props: {
    visible: Boolean
  },
  computed: {
    styles() {
      return {
        "display": 'block'
      }
    }
  },
  data() {
    return {};
  },
});
</script>