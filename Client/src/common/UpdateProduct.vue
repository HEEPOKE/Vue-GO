<template>
    <HeaderUpdate />
    <div class="flex justify-center mt-5">
        <FormUpdateProduct />
    </div>
</template>

<script lang="ts">
import HeaderUpdate from '../components/UpdateProduct/HeaderUpdate.vue';
import FormUpdateProduct from '../components/UpdateProduct/FormUpdate.vue';

export default {
    name: "UpdateProductCommon",
    components: {
        HeaderUpdate,
        FormUpdateProduct
    }
}
</script>
