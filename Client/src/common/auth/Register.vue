<template>
    <section class="h-screen">
        <div class="container px-6 py-12 h-full">
            <div class="flex justify-center items-center flex-wrap h-full g-6 text-gray-800">
                <div class="md:w-8/12 lg:w-6/12 mb-12 md:mb-0">
                    <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/draw2.svg"
                        class="w-full" alt="Phone image" />
                </div>
                <div class="md:w-8/12 lg:w-5/12 lg:ml-20">
                    <HeaderRegister />
                    <form v-on:submit.prevent @submit="handlerSubmit()">
                        <div class="mb-3">
                            <label class="form-label inline-block text-lg mb-2 text-gray-700 dark:text-white">Username</label>
                            <input type="text"
                                class="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none dark:text-gray-700"
                                v-model="register.username" placeholder="Username" required />
                        </div>
                        <div class="mb-3">
                            <label class="form-label inline-block text-lg mb-2 text-gray-700 dark:text-gray-700">Email</label>
                            <input type="email"
                                class="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none dark:text-gray-700"
                                v-model="register.email" placeholder="Email" required />
                        </div>
                        <div class="mb-3">
                            <label class="form-label inline-block text-lg mb-2 text-gray-700 dark:text-gray-700">Password</label>
                            <input type="password"
                                class="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none dark:text-gray-700"
                                v-model="register.password" placeholder="Password" minlength="8" maxlength="20"
                                required />
                        </div>
                        <div class="mb-3">
                            <label class="form-label inline-block text-lg mb-2 text-gray-700 dark:text-gray-700">Confirm Password</label>
                            <input type="password"
                                class="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none dark:text-gray-700"
                                minlength="8" maxlength="20" v-model="register.confirmPassword"
                                placeholder="Confirm Password" required />
                        </div>
                        <div class="mb-3">
                            <label class="form-label inline-block text-lg mb-2 text-gray-700 dark:text-white">Telephone</label>
                            <input type="tel"
                                class="form-control block w-full px-4 py-2 text-xl font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none dark:text-gray-700"
                                minlength="10" maxlength="10" v-model="register.tel" placeholder="xxx-xxx-xxxx"
                                required />
                        </div>
                        <button type="submit"
                            class="inline-block px-7 py-3 mb-3 bg-blue-600 text-white font-medium text-sm rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out w-full dark:text-white">
                            SIGN UP
                        </button>
                    </form>
                    <button type="button"
                        class="inline-block px-7 py-3 bg-gray-300 text-gray-700 font-medium text-sm rounded shadow-md hover:bg-gray-300 hover:shadow-lg focus:bg-gray-300 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-gray-400 active:shadow-lg transition duration-150 ease-in-out w-full dark:text-white dark:bg-rose-600"
                        @click="goBack">
                        BACK
                    </button>
                    <div
                        class="flex items-center my-4 before:flex-1 before:border-t before:border-gray-300 before:mt-0.5 after:flex-1 after:border-t after:border-gray-300 after:mt-0.5 dark:text-white">
                        <p class="text-center font-semibold mx-4 mb-0 dark:text-white">OR</p>
                    </div>
                    <a class="px-7 py-3 text-white font-medium text-sm rounded shadow-md hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:shadow-lg transition duration-150 ease-in-out w-full flex justify-center items-center mb-3"
                        style="background-color: #3b5998" href="#!" role="button" data-mdb-ripple="true"
                        data-mdb-ripple-color="light">
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
import HeaderRegister from '../../components/auth/register/HeaderRegister.vue';
import RegisterModel from '../../models/auth/register';
import AuthService from '../../services/AuthService';
import swalAuth from "../../assets/ts/auth";

export default defineComponent({
    name: "Register",
    data() {
        return {
            register: {
                username: "",
                email: "",
                password: "",
                confirmPassword: "",
                tel: "",
                role: 1,
            } as RegisterModel,
        }
    },
    components: {
        HeaderRegister
    }, methods: {
        handlerSubmit() {
            const password = this.register.password;
            const confirmPassword = this.register.confirmPassword;

            let data = {
                username: this.register.username,
                email: this.register.email,
                password: password,
                tel: this.register.tel,
                role: 1,
            }

            if (password != confirmPassword) {
                swalAuth.checkPassword()
            } else {
                AuthService.RegisterService(data)
            }
        },
        goBack() {
            this.$router.go(-1)
        },
    }
})
</script>
