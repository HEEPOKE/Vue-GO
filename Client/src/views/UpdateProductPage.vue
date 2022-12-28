<template>
    <div>
        <NavbarMenu />
        <UpdateProductCommon />
    </div>
</template>

<script lang="ts">
import NavbarMenu from '../common/Navbar.vue';
import UpdateProductCommon from '../common/UpdateProduct.vue';

export default {
    name: "UpdateProductPage",
    components: {
        NavbarMenu,
        UpdateProductCommon
    }
}
</script>
