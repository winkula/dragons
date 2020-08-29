<template>
  <section class="el-menu">
    <button id="btn-new" type="button" @click="$emit('new')">New</button>
    <button id="btn-difficulty" type="button" @click="$emit('difficulty')">Difficulty</button>
    <button id="btn-help" type="button" @click="$emit('help')">Help</button>
  </section>
</template>

<script lang="ts">
import Vue from "vue";
export default Vue.extend({});
</script>