<template>
  <button
    type="button"
    class="inline-block float-right px-5 py-3 my-2 mx-7 bg-rose-600 text-white ont-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-rose-700 hover:shadow-lg focus:bg-rose-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-rose-800 active:shadow-lg transition duration-150 ease-in-out"
    data-bs-toggle="modal"
    data-bs-target="#AddProduct"
  >
    <font-awesome-icon icon="fa-solid fa-plus" class="w-4 h-3 mr-1" />Add
  </button>

  <div
    class="modal fade fixed top-0 left-0 hidden w-full h-full outline-none overflow-x-hidden overflow-y-auto"
    id="AddProduct"
    tabindex="-1"
    aria-labelledby="AddProductLabel"
    aria-hidden="true"
  >
    <div class="modal-dialog relative w-auto pointer-events-none">
      <div
        class="modal-content border-none shadow-lg relative flex flex-col w-full pointer-events-auto bg-white bg-clip-padding rounded-md outline-none text-current"
      >
        <div
          class="modal-header flex flex-shrink-0 items-center justify-between p-4 border-b border-gray-200 rounded-t-md"
        >
          <h5
            class="text-xl font-medium leading-normal text-gray-800"
            id="AddProductLabel"
          >
            Add Product
          </h5>
          <button
            type="button"
            class="btn-close box-content w-4 h-4 p-1 text-black border-none rounded-none opacity-50 focus:shadow-none focus:outline-none focus:opacity-100 hover:text-black hover:opacity-75 hover:no-underline"
            data-bs-dismiss="modal"
            aria-label="Close"
          ></button>
        </div>
        <form v-on:submit.prevent @submit="handlerSubmit()">
          <div class="modal-body relative p-4 w-full">
            <div class="mb-3 xl:w-full">
              <label class="form-label inline-block mb-2 text-gray-700">
                Code
              </label>
              <input
                type="number"
                class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
                v-model="product.code"
                placeholder="Code"
                required
              />
            </div>
            <div class="mb-3 xl:w-full">
              <label class="form-label inline-block mb-2 text-gray-700">
                Name
              </label>
              <input
                type="text"
                class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
                v-model="product.name"
                placeholder="Name"
                required
              />
            </div>
            <div class="mb-3 xl:w-full">
              <label class="form-label inline-block mb-2 text-gray-700">
                Description
              </label>
              <textarea
                class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
                rows="3"
                v-model="product.description"
                placeholder="Description"
                required
              ></textarea>
            </div>
            <div class="mb-3 xl:w-full">
              <label class="form-label inline-block mb-2 text-gray-700">
                Price
              </label>
              <input
                type="number"
                class="block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:border-rose-600 focus:outline-none"
                min="0"
                v-model="product.price"
                placeholder="บาท"
              />
            </div>
          </div>
          <div
            class="modal-footer flex flex-shrink-0 flex-wrap items-center justify-end p-4 border-t border-gray-200 rounded-b-md"
          >
            <button
              type="submit"
              class="px-6 py-2.5 bg-rose-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-rose-700 hover:shadow-lg focus:bg-rose-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-rose-800 active:shadow-lg transition duration-150 ease-in-out"
            >
              บันทึก
            </button>
            <button
              type="button"
              class="px-6 py-2.5 ml-1 bg-purple-600 text-white font-medium text-xs leading-tight uppercase rounded shadow-md hover:bg-purple-700 hover:shadow-lg focus:bg-purple-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-purple-800 active:shadow-lg transition duration-150 ease-in-out"
              data-bs-dismiss="modal"
            >
              ยกเลิก
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import ProductModel from "../../models/product";
import ProductServices from "../../services/ProductService";

export default defineComponent({
  name: "AddProductBtn",
  data() {
    return {
      product: {
        code: 0,
        name: "",
        description: "",
        price: 0,
      } as ProductModel,
    };
  },
  methods: {
    handlerSubmit() {
      let data = {
        code: this.product.code,
        name: this.product.name,
        description: this.product.description,
        price: this.product.price,
      };
      ProductServices.AddProduct(data);   
    },
  },
});
</script>
