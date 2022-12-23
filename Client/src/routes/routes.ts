import { createRouter, createWebHistory } from "vue-router";
const HomePage = () => import("../views/HomePage.vue");
const ProductsListPage = () => import("../views/ProductsListPage.vue");
const AboutPage = () => import("../views/AboutPage.vue");

const routes = [
  { path: "/", name: "HomePage", component: HomePage },
  { path: "/product", name: "ProductsListPage", component: ProductsListPage },
  { path: "/about", name: "AboutPage", component: AboutPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
