<template>
  <HeaderProduct />
  <div class="flex justify-center mt-5">
    <div class="block rounded-lg shadow-lg bg-white mx-auto">
      <div class="py-3 px-6 border-b border-gray-300">
        Product List
      </div>
      <AddProductBtn />
      <div class="p-6 mt-8">
        <DataTableProduct />
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import HeaderProduct from '../components/Product/Header.vue';
import DataTableProduct from '../components/Product/DataTableProduct.vue';
import AddProductBtn from '../components/Product/AddProductBtn.vue';

export default {
  name: "ProductsListCommon",
  components: {
    HeaderProduct,
    DataTableProduct,
    AddProductBtn
  },
};
</script>
