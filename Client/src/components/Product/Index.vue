<template>
  <div>
    <HeaderProduct />
    <div class="flex justify-center mt-5">
      <div class="block rounded-lg shadow-lg bg-white mx-auto">
        <div class="py-3 px-6 border-b border-gray-300">
          Product List
        </div>
        <CreateProductBtn />
        <div class="p-6 mt-8">
          <DataTableProduct />
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import HeaderProduct from "./Header.vue";
import DataTableProduct from "./DataTableProduct.vue";
import CreateProductBtn from "./CreateProductBtn.vue";

export default {
  name: "ProductList",
  components: {
    HeaderProduct,
    DataTableProduct,
    CreateProductBtn
  },
};
</script>
