<template>
  <div>
    <HeaderProduct />
    <div class="flex justify-center">
      <div class="block rounded-lg shadow-lg bg-white mx-auto">
        <div class="py-3 px-6 border-b border-gray-300">
          Product List
        </div>
        <div class="p-6">
          <DataTableProduct />
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import HeaderProduct from "./Header.vue";
import DataTableProduct from "./DataTableProduct.vue";

export default {
  name: "ProductList",
  components: {
    HeaderProduct,
    DataTableProduct
  },
};
</script>
