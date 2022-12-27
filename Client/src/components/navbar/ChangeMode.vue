<template>
    <div class="hidden w-full md:flex md:items-center md:w-auto" @click="toggleDarkMode">
        <input type="checkbox" class="checkboxChange absolute opacity-0" id="checkboxChange" />
        <label for="checkbox"
            class="labell relative bg-white border border-dark-900 dark:border-light-100 flex justify-between items-center scale-150 rounded-4xl w-55px p-5px !cursor-pointer">
            <font-awesome-icon icon="fa-solid fa-moon" />
            <font-awesome-icon icon="fa-solid fa-sun" class="text-dark-900" />
            <div class="ball !absolute bg-gray-900 left-2px w-24px h-24px rounded-1/2 transition-transform !ease-linear"
                id="ball"></div>
        </label>
    </div>
</template>
<script lang="ts">
export default {
    name: "ChangeModeButton",
    methods: {
        toggleDarkMode() {
            const ball = document.querySelector("#ball");
            const html = document.querySelector("html");
            const checkbox = document.querySelector(
                "#checkboxChange"
            ) as HTMLInputElement | null;
            ball?.classList.toggle("translateChange");
            const toggleDarkMode = function () {
                checkbox?.checked;
                html?.classList.toggle("dark");
            };
            toggleDarkMode();
            checkbox?.addEventListener("click", toggleDarkMode);
        },
    },
}
</script>
