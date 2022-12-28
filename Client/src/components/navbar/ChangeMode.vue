<template>
    <div class="hidden w-full md:flex md:items-center md:w-auto" @click="toggleDarkMode">
        <input type="checkbox" class="checkbox absolute opacity-0" id="checkboxChange" />
        <label class="checkbox-label flex relative p-1.5 items-center justify-between bg-rose-400 rounded-full cursor-pointer border border-white dark:bg-black">
            <div class="toggle-slot">
                <font-awesome-icon icon="fa-solid fa-moon" class="text-amber-300 w-6 h-4" />
                <font-awesome-icon icon="fa-solid fa-sun" class="text-yellow-400 w-6 h-4" />
                <div class="ball bg-white absolute left-2 top-1 w-6 h-6 rounded-full transition-transform !ease-linear" id="ball"></div>
            </div>
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
