<template>
    <div class="home">
        <button
            v-on:click="api('/router_status')"
            class="btn btn-primary waves-effect"
        >
            /api/router_status
        </button>

        <button
            v-on:click="api('/list_projects')"
            class="px-4 py-2 font-semibold text-sm bg-cyan-500 text-white rounded-full shadow"
        >
            /api/list_projects
        </button>
    </div>
</template>

<script>
import { api } from "../utils/api";

export default {
    methods: {
        api: (endpoint) => api(endpoint),
    },
};
</script>
