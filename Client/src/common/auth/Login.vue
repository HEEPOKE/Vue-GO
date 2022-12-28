<template>
    <section class="h-screen">
        <div class="container px-6 py-12 h-full">
            <div class="flex justify-center items-center flex-wrap h-full g-6 text-gray-800 dark:text-white">
                <div class="md:w-8/12 lg:w-6/12 mb-12 md:mb-0">
                    <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.svg"
                        class="w-full" alt="Phone image" />
                </div>
                <div class="md:w-8/12 lg:w-5/12 lg:ml-20">
                    <HeaderLogin />
                    <form>
                        <div class="mb-6">
                            <label class="form-label inline-block text-lg mb-2 text-gray-700 dark:text-white">Email</label>
                            <input type="text"
                                class="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none dark:text-white"
                                placeholder="Email" required />
                        </div>
                        <div class="mb-6">
                            <label class="form-label inline-block text-lg mb-2 text-gray-700 dark:text-white">Password</label>
                            <input type="password"
                                class="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none dark:text-white"
                                placeholder="Password" required />
                        </div>
                        <div class="flex justify-between items-center mb-6">
                            <div class="form-group form-check">
                                <input type="checkbox"
                                    class="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-blue-600 checked:border-blue-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer dark:text-white"
                                    id="exampleCheck3" checked />
                                <label class="form-check-label inline-block text-gray-800 dark:text-white" for="exampleCheck2">Remember
                                    me</label>
                            </div>
                            <RouterLink to="/auth/forgotPassword">
                                <a
                                    class="text-blue-600 hover:text-blue-700 focus:text-blue-700 active:text-blue-800 duration-200 transition ease-in-out dark:text-white">Forgot
                                    password?</a>
                            </RouterLink>
                        </div>
                        <button type="submit"
                            class="inline-block px-7 py-3 mb-3 bg-blue-600 text-white font-medium text-sm rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out w-full dark:text-white">
                            SIGN IN
                        </button>
                    </form>
                    <button type="button"
                        class="inline-block px-7 py-3 bg-gray-300 text-gray-700 font-medium text-sm rounded shadow-md hover:bg-gray-300 hover:shadow-lg focus:bg-gray-300 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-gray-400 active:shadow-lg transition duration-150 ease-in-out w-full dark:text-white dark:bg-rose-600"
                        @click="goBack">
                        BACK
                    </button>
                    <p class="text-gray-800 mt-6 text-center dark:text-white">Not a member?
                        <RouterLink to="/auth/register">
                            <a
                                class="text-blue-600 hover:text-blue-700 focus:text-blue-700 transition duration-200 ease-in-out dark:text-rose-600">Register</a>
                        </RouterLink>
                    </p>
                    <div
                        class="flex items-center my-4 before:flex-1 before:border-t before:border-gray-300 before:mt-0.5 after:flex-1 after:border-t after:border-gray-300 after:mt-0.5 dark:text-white">
                        <p class="text-center font-semibold mx-4 mb-0 dark:text-white">OR</p>
                    </div>
                    <a class="px-7 py-3 text-white font-medium text-sm leading-snug uppercase rounded shadow-md hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:shadow-lg transition duration-150 ease-in-out w-full flex justify-center items-center mb-3"
                        style="background-color: #3b5998" role="button">
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 320 512" class="w-3.5 h-3.5 mr-2">
                            <path fill="currentColor"
                                d="M279.14 288l14.22-92.66h-88.91v-60.13c0-25.35 12.42-50.06 52.24-50.06h40.42V6.26S260.43 0 225.36 0c-73.22 0-121.08 44.38-121.08 124.72v70.62H22.89V288h81.39v224h100.17V288z" />
                        </svg>Continue with Facebook
                    </a>
                </div>
            </div>
        </div>
    </section>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import HeaderLogin from "../../components/auth/login/HeaderLogin.vue";

export default defineComponent({
    name: "Login",
    components: {
        HeaderLogin,
    },
    methods: {
        goBack() {
            this.$router.go(-1);
        },
    },
});
</script>
