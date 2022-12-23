import { createRouter, createWebHistory } from "vue-router";
// import HomePage from "../views/HomePage.vue";
import AboutPage from "../views/AboutPage.vue";
const HomePage = () => import("../views/HomePage.vue");

const routes = [
  { path: "/", name: "HomePage", component: HomePage },
  { path: "/about", name: "AboutPage", component: AboutPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
