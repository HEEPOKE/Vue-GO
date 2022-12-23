<template>
    <div>
        <NavbarMenu />
        <ProductList />
    </div>
</template>
<script lang="ts">
import NavbarMenu from '../common/Navbar.vue'
import ProductList from '../components/Product/Index.vue'

export default {
    name: "ProductsListCommon",
    components: {
        NavbarMenu,
        ProductList
    }
}

</script>