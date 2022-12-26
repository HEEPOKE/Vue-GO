import { createRouter, createWebHistory } from "vue-router";
const LoginPage = () => import("../views/auth/LoginPage.vue");
const RegisterPage = () => import("../views/auth/RegisterPage.vue");
const ForgotPasswordPage = () => import("../views/auth/ForgotPasswordPage.vue");
const HomePage = () => import("../views/HomePage.vue");
const ProductsListPage = () => import("../views/ProductsListPage.vue");
const AboutPage = () => import("../views/AboutPage.vue");
const NotFoundPage = () => import("../views/NotFoundPage.vue");

const routes = [
  { path: "/auth/login", name: "LoginPage", component: LoginPage },
  { path: "/auth/register", name: "RegisterPage", component: RegisterPage },
  {
    path: "/auth/forgotPassword",
    name: "ForgotPasswordPage",
    component: ForgotPasswordPage,
  },
  { path: "/", name: "HomePage", component: HomePage },
  { path: "/product", name: "ProductsListPage", component: ProductsListPage },
  { path: "/about", name: "AboutPage", component: AboutPage },
  { path: "/*", name: "NotFoundPage", component: NotFoundPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
