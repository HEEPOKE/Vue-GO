<template>
  <div class="flex flex-col">
    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="py-2 inline-block max-w-full sm:px-6 lg:px-8">
        <div class="overflow-hidden">
          <table class="border text-center max-w-full">
            <thead class="border-b bg-rose-600 dark:bg-black">
              <tr>
                <th
                  scope="col"
                  class="text-sm font-medium text-white px-6 py-4 border-r"
                >
                  name
                </th>
                <th
                  scope="col"
                  class="text-sm font-medium text-white px-6 py-4 border-r"
                >
                  description
                </th>
                <th
                  scope="col"
                  class="text-sm font-medium text-white px-6 py-4 border-r"
                >
                  price
                </th>
                <th
                  scope="col"
                  class="text-sm font-medium text-white px-6 py-4"
                >
                  manage
                </th>
              </tr>
            </thead>
            <tbody>
              <tr
                class="bg-white border-b transition duration-300 ease-in-out hover:bg-gray-100"
                v-for="product in products"
                :key="product.ID"
              >
                <td
                  class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap border-r"
                >
                  {{ product.name }}
                </td>
                <td
                  class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap border-r"
                >
                  {{ product.description }}
                </td>
                <td
                  class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap border-r"
                >
                  {{ product.price }} บาท
                </td>
                <td
                  class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap"
                >
                  <div class="justify-center">
                    <button
                      type="button"
                      class="inline-block mx-1 px-4 text-center py-2.5 bg-blue-600 text-white ont-medium text-xs leading-tight rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
                    >
                      <font-awesome-icon
                        icon="fa-solid fa-eye"
                        class="w-4 h-3"
                      />
                    </button>
                    <routerLink
                      :to="{
                        name: 'UpdateProductPage',
                        params: { id: product.ID },
                      }"
                    >
                      <button
                        type="button"
                        class="inline-block mx-1 px-4 text-center py-2.5 bg-yellow-500 text-white font-medium text-xs leading-tight rounded shadow-md hover:bg-yellow-600 hover:shadow-lg focus:bg-yellow-600 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-yellow-700 active:shadow-lg transition duration-150 ease-in-out"
                      >
                        <font-awesome-icon
                          icon="fa-solid fa-pen-to-square"
                          class="w-4 h-3"
                        />
                      </button>
                    </routerLink>
                    <button
                      type="button"
                      class="inline-block mx-1 px-4 text-center py-2.5 bg-red-600 text-white font-medium text-xs leading-tight rounded shadow-md hover:bg-red-700 hover:shadow-lg focus:bg-red-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-red-800 active:shadow-lg transition duration-150 ease-in-out"
                    >
                      <font-awesome-icon
                        icon="fa-solid fa-trash-can"
                        class="w-4 h-3"
                      />
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import Http from "../../http/http";
import ProductModel from "../../models/product";
import ProductServices from "../../services/ProductService";

export default {
  name: "DataTableProduct",
  data() {
    return {
      products: [] as ProductModel[],
    };
  },
  methods: {
    readProduct() {
      ProductServices.ReadProduct()
        .then((res: any) => {
          this.products = res.data.payload;
          console.log(res.data.payload);
        })
        .catch((err: any) => {
          console.log(err);
        });
    },
    deleteProduct() {
      // let id = this.product.findIndex(i => i.id === id)
    },
  },
  mounted() {
    this.readProduct();
  },
};
</script>
