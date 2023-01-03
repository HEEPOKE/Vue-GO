<template>
  <div class="block p-6 rounded-lg shadow-lg bg-white w-full ml-20 mr-20">
    <form v-on:submit.prevent @submit="handlerSubmit()">
      <div class="form-group mb-6">
        <label
          for="exampleInputEmail1"
          class="form-label inline-block mb-2 text-gray-700"
          >Code</label
        >
        <input
          type="text"
          class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
        />
      </div>
      <div class="form-group mb-6">
        <label class="form-label inline-block mb-2 text-gray-700">Name</label>
        <input
          type="text"
          class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
        />
      </div>
      <div class="form-group mb-6">
        <label class="form-label inline-block mb-2 text-gray-700"
          >Description</label
        >
        <textarea
          class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
          rows="3"
        >
        </textarea>
      </div>
      <div class="form-group mb-6">
        <label class="form-label inline-block mb-2 text-gray-700">Price</label>
        <input
          type="number"
          class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
          min="0"
        />
      </div>
      <div class="flex justify-center">
        <button
          type="submit"
          class="px-6 py-2.5 mr-3 bg-rose-600 text-white font-xl text-lg leading-tight uppercase rounded shadow-md hover:bg-rose-700 hover:shadow-lg focus:bg-rose-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-rose-800 active:shadow-lg transition duration-150 ease-in-out"
        >
          Submit
        </button>
        <button
          type="button"
          class="px-6 py-2.5 bg-gray-600 text-white font-xl text-lg leading-tight uppercase rounded shadow-md hover:bg-gray-700 hover:shadow-lg focus:bg-gray-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-gray-800 active:shadow-lg transition duration-150 ease-in-out"
          @click="goBack"
        >
          Back
        </button>
      </div>
    </form>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import Http from "../../http/http";
import ProductServices from "../../services/ProductService";
import ProductModel from "../../models/product";

export default defineComponent({
  name: "FormUpdateProduct",
  data() {
    return {
      products: [] as ProductModel[],
    };
  },
  methods: {
    getProduct() {
      const getId = this.$route.params.id;

      if (typeof getId === "string") {
        const id = parseInt(getId);

        ProductServices.GetProductById(id)
          .then((res: any) => {
            console.log(res.data.payload);
          })
          .catch((err: any) => {
            console.log(err);
          });
      } else {
        this.goBack();
      }
    },
    updateProduct(id: number, data: any) {
      Http.put<ProductModel>(`/api/product/update${id}`, data)
        .then((res: any) => {
          console.log(res);
        })
        .catch((err: any) => {
          console.log(err);
        });
    },
    handlerSubmit() {},
    goBack() {
      this.$router.go(-1);
    },
  },
  mounted() {
    this.getProduct();
  },
});
</script>
