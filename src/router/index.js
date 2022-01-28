import { createRouter, createWebHistory } from "vue-router";
import Progress from "../utils/progress";
import Home from "../views/Home.vue";
import Waves from "node-waves";

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
    },
    {
        path: "/projects",
        name: "Projects",
        // route level code-splitting
        // this generates a separate chunk (projects.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () =>
            import(/* webpackChunkName: "projects" */ "../views/Projects.vue"),
    },
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes,
});

const progress = new Progress();

router.beforeEach(() => progress.start());
router.afterEach(() => {
    Waves.init();
    progress.complete();
});

export default router;
