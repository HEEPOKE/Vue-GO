import { createApp } from "vue";
import App from "./App.vue";
import router from "./routes/routes";
import "tw-elements";
import "./index.css";
import "./assets/css/NotFound.css";

createApp(App).use(router).mount("#app");
