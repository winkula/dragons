<template>
  <footer id="footer">{{ copyright }}</footer>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  components: {},
  props: { copyright: String },
});
</script>
