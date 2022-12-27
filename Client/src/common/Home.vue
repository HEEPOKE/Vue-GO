<template>
  <button class="text-white bg-blue-500 hover:bg-blue-700 rounded-full py-2 px-4" :class="{ 'dark-mode': darkMode }"
    @click="toggleDarkMode">
    Toggle Dark Mode
  </button>
</template>
<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: "Home",
  data() {
    return {
      darkMode: false,
    };
  },
  methods: {
    toggleDarkMode() {
      this.darkMode = !this.darkMode;
    },
  },
})
</script>
