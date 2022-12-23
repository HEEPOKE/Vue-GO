<template>
  <router-view />
</template>
<script lang="ts">
import { RouterView } from 'vue-router';
import HomePage from './views/HomePage.vue';

export default {
  name: "App",
  components: {
    HomePage,
  }
}
</script>
