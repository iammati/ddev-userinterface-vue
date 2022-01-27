<template>
    <div class="projects">
        <h1>This is a list of projects</h1>

        <button id="list_projects">
            /api/list_projects
        </button>

        <ul id="projects"></ul>
    </div>
</template>

<script lang="ts">
interface ProjectRow {
    approot: string;
    docroot: string;
    htpsurl: string;
    httpurl: string;
    name: string;
    primary_url: string;
    router_disabled: string;
    shortroot: string;
    status: string;
    type: string;
}

export default {
    mounted() {
        const node = document.getElementById("list_projects");

        if (!node) return;

        node.onclick = async e => {
            const response = await fetch("/api/list_projects")
            const text = await response.text();
            const json = JSON.parse(text);

            // rows includes list of projects
            const rows = JSON.parse(json.list).raw as ProjectRow[];
            const ul = document.getElementById("projects");

            if (!ul) return;

            rows.forEach((row: ProjectRow) => {
                ul.insertAdjacentHTML('beforeend', `<li>name: ${row.name}, url: ${row.primary_url}, status: ${row.status}, type: ${row.type}</li>`);
            });
        };
    }
}
</script>
