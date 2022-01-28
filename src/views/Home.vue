<template>
    <div class="home">
        <button v-on:click="api('/router_status')">
            /api/router_status
        </button>
        <button v-on:click="api('/list_projects')">
            /api/list_projects
        </button>
    </div>
</template>

<script>
import { api } from "../utils/api";

export default {
    methods: {
        api: endpoint => api(endpoint),
    },
}
</script>
