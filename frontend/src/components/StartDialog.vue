<template>
  <Dialog v-model="value">
    <p class="difficulty-text">Select the difficulty:</p>
    <Buttons :buttons="buttons" @click="$emit('start', $event)"></Buttons>
  </Dialog>
</template>

<script lang="ts">
import Vue from "vue";
import Dialog from "./Dialog.vue";
import Buttons from "./Buttons.vue";

export default Vue.extend({
  components: {
    Dialog,
    Buttons,
  },
  props: {
    value: Boolean,
  },
  data() {
    return {
      buttons: [
        {
          id: "easy",
          icon: `<circle cx="12" cy="12" r="10"></circle><path d="M8 14s1.5 2 4 2 4-2 4-2"></path><line x1="9" y1="9" x2="9.01" y2="9"></line><line x1="15" y1="9" x2="15.01" y2="9"></line>`,
        },
        {
          id: "medium",
          icon: `<circle cx="12" cy="12" r="10"></circle><line x1="8" y1="15" x2="16" y2="15"></line><line x1="9" y1="9" x2="9.01" y2="9"></line><line x1="15" y1="9" x2="15.01" y2="9"></line>`,
        },
        {
          id: "hard",
          icon: `<circle cx="12" cy="12" r="10"></circle><path d="M16 16s-1.5-2-4-2-4 2-4 2"></path><line x1="9" y1="9" x2="9.01" y2="9"></line><line x1="15" y1="9" x2="15.01" y2="9"></line>`,
        },
      ],
    };
  },
});
</script>

<style lang="scss">
.difficulty-text {
  text-align: center;
}
</style>
