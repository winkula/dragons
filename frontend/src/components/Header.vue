<template>
  <header id="header">
    <span>{{title}}</span>
  </header>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  components: {},
  data() {
    return {
      title: "dragons",
    };
  },
});
</script>