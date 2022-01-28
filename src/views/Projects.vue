<template>
    <div class="projects">
        <h1>This is a list of projects</h1>

        <button v-on:click="getItems()">/api/list_projects</button>

        <ul id="projects">
            <li v-for="(item, index) in items" :key="index">
                <ProjectRow :data="item" />
            </li>
        </ul>
    </div>
</template>

<script>
import ProjectRow from "@/components/ProjectRow.vue";
import { api } from "../utils/api";

export default {
    name: "Projects",
    components: {
        ProjectRow,
    },
    data() {
        return {
            items: null,
        };
    },
    methods: {
        getItems: async function () {
            const response = await api("/list_projects");
            const json = await response.json();

            // rows includes list of projects
            const rows = JSON.parse(json.list).raw;
            this.items = rows;
        },
    },
    async mounted() {
        const response = await api("/list_projects");
        const json = await response.json();

        // rows includes list of projects
        const rows = JSON.parse(json.list).raw;
        this.items = rows;
    },
};
</script>
