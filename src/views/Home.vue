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
            class="btn btn-info waves-effect"
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
