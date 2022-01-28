import { createApp } from "vue";
import App from "./App.vue";
import './main.scss'

import "./registerServiceWorker";
import router from "./router";

createApp(App).use(router).mount("#app");
