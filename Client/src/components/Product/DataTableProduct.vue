<template>
    <div class="flex flex-col">
        <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
            <div class="py-2 inline-block min-w-full sm:px-6 lg:px-8">
                <div class="overflow-hidden">
                    <table class="min-w-full border text-center">
                        <thead class="border-b bg-gray-800">
                            <tr>
                                <th scope="col" class="text-sm font-medium text-white px-6 py-4 border-r">
                                    #
                                </th>
                                <th scope="col" class="text-sm font-medium text-white px-6 py-4 border-r">
                                    First
                                </th>
                                <th scope="col" class="text-sm font-medium text-white px-6 py-4 border-r">
                                    Last
                                </th>
                                <th scope="col" class="text-sm font-medium text-white px-6 py-4">
                                    Handle
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr class="bg-white border-b transition duration-300 ease-in-out hover:bg-gray-100">
                                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 border-r">1
                                </td>
                                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap border-r">
                                    Mark
                                </td>
                                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap border-r">
                                    Otto
                                </td>
                                <td class="text-sm text-gray-900 font-light px-6 py-4 whitespace-nowrap">
                                    @mdo
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</template>
<script lang="ts" >

export default {
    name: "DataTableProduct",

}
</script>
