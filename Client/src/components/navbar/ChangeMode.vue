<template>
    <div class="hidden w-full md:flex md:items-center md:w-auto" @click="toggleDarkMode">
        <input type="checkbox" class="checkbox absolute opacity-0" id="checkboxChange" />
        <label class="checkbox-label w-18">
            <div class="toggle-slot">
                <font-awesome-icon icon="fa-solid fa-moon" class="moon-icon w-6 h-4" />
                <font-awesome-icon icon="fa-solid fa-sun" class="sun-icon w-6 h-4" />
                <div class="ball bg-gray-900 left-2px w-28px h-28px rounded-1/2 transition-transform !ease-linear" id="ball"></div>
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
