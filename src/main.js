import "mdb-vue-ui-kit/css/mdb.min.css";
import { createApp } from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import "node-waves/dist/waves.min.css";
import "sweetalert2/src/sweetalert2.scss";

createApp(App).use(router).mount("#app");
