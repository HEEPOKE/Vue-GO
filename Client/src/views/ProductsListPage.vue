<template>
    <div>
        <NavbarMenu />
        <ProductsListCommon />
    </div>
</template>
<script lang="ts">
import NavbarMenu from '../common/Navbar.vue';
import ProductsListCommon from '../common/Product.vue'

export default {
    name: "ProductsListPage",
    components: {
        NavbarMenu,
        ProductsListCommon
    }
}
</script>
