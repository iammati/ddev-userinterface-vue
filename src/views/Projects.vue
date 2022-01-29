<template>
    <div class="projects">
        <h1 >
            Projects – {{ items && items.length }}
        </h1>

        <!-- <button class="btn btn-warning waves-effect" v-on:click="getItems()">/api/list_projects</button> -->

        <div class="table-responsive">
            <table class="table table-striped table-hover table-sm ">
                <thead>
                    <tr>
                        <th scope="col">#</th>
                        <th scope="col">Name</th>
                        <th scope="col">URL</th>
                        <th scope="col">Document Root</th>
                        <th scope="col">Type</th>
                        <th scope="col">Status</th>
                        <th scope="col">Actions</th>
                    </tr>
                </thead>

                <tbody>
                    <tr v-for="(item, index) in items" :key="index">
                        <th scope="row">
                            {{ index + 1 }}
                        </th>

                        <ProjectRow :data="item" />
                    </tr>
                </tbody>

                <caption>
                    List of detected projects configured on your machine.
                </caption>
            </table>
        </div>
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
