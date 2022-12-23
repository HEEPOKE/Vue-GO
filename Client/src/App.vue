<template>
  <router-view />
</template>
<script lang="ts">
import HomePage from './views/HomePage.vue';

export default {
  name: "App",
  components: {
    HomePage,
  }
}
</script>
