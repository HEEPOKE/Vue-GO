<template>
  <nav
    class="relative w-full flex flex-wrap items-center justify-between py-3 bg-rose-600 text-white shadow-lg navbar navbar-expand-lg navbar-light dark:bg-black dark:text-gray-700">
    <div class="container-fluid w-full flex flex-wrap items-center justify-between px-6">
      <button
        class="navbar-toggler text-gray-200 border-0 hover:shadow-none hover:no-underline py-2 px-2.5 bg-transparent focus:outline-none focus:ring-0 focus:shadow-none focus:no-underline"
        type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent1"
        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="bars" class="w-6" role="img"
          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
          <path fill="currentColor"
            d="M16 132h416c8.837 0 16-7.163 16-16V76c0-8.837-7.163-16-16-16H16C7.163 60 0 67.163 0 76v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16zm0 160h416c8.837 0 16-7.163 16-16v-40c0-8.837-7.163-16-16-16H16c-8.837 0-16 7.163-16 16v40c0 8.837 7.163 16 16 16z">
          </path>
        </svg>
      </button>
      <div class="collapse navbar-collapse flex-grow items-center" id="navbarSupportedContent1">
        <Router-link to="/" class="text-xl text-white pr-2 font-semibold dark:text-slate-100 dark:hover:text-white">
          Topic
        </Router-link>
        <ul class="navbar-nav flex flex-col pl-0 list-style-none mr-auto">
          <li class="nav-item p-2">
            <RouterLink to="/" class="nav-link text-white opacity-60 hover:opacity-80 focus:opacity-80 p-0 dark:text-slate-100 dark:hover:text-white">
              Dashboard
            </RouterLink>
          </li>
          <li class="nav-item p-2">
            <RouterLink to="/product" class="nav-link text-white opacity-60 hover:opacity-80 focus:opacity-80 p-0 dark:text-slate-100 dark:hover:text-white">
              Product
            </RouterLink>
          </li>
          <li class="nav-item p-2">
            <RouterLink to="/about" class="nav-link text-white opacity-60 hover:opacity-80 focus:opacity-80 p-0 dark:text-slate-100 dark:hover:text-white">
              About
            </RouterLink>
          </li>
        </ul>
        <div class="flex items-center relative">
          <ul class="navbar-nav flex flex-col pl-0 list-style-none mr-auto">
            <li class="nav-item p-2">
              <RouterLink to="/auth/login" class="nav-link text-white opacity-60 hover:opacity-80 focus:opacity-80 p-0 dark:text-slate-100 dark:hover:text-white">
                Login
              </RouterLink>
            </li>
            <li class="nav-item p-2">
              <RouterLink to="/auth/register"
                class="nav-link text-white opacity-60 hover:opacity-80 focus:opacity-80 p-0 dark:text-slate-100 dark:hover:text-white">
                Register
              </RouterLink>
            </li>
          </ul>
          <ChangeModeButton />
        </div>
      </div>
    </div>
  </nav>
</template>
<script lang="ts">
import ChangeModeButton from './ChangeMode.vue';

export default {
  name: "NavBarVue",
  components: {
    ChangeModeButton,
  }
};
</script>
