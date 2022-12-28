import { createApp } from "vue";
import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faMoon, faSun, faPlus, faEye, faPenToSquare, faTrashCan } from "@fortawesome/free-solid-svg-icons";
import App from "./App.vue";
import router from "./routes/routes";
import "tw-elements";
import "./index.css";
import "./assets/css/NotFound.css";

library.add(faMoon, faSun, faPlus, faEye, faPenToSquare, faTrashCan);

createApp(App)
  .use(router)
  .component("font-awesome-icon", FontAwesomeIcon)
  .mount("#app");
