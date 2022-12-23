<template>
    <div>
        <NavbarMenu />
        <div>
            <Home />
        </div>
    </div>
</template>

<script lang="ts">
import NavbarMenu from '../common/Navbar.vue'
import Home from '../common/Home.vue'

export default {
    name: "HomePage",
    components: {
        NavbarMenu,
        Home
    },
}
</script>
