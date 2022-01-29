<template>
    <td>{{ data.name }}</td>

    <td>
        <a :href="data.primary_url" target="_blank">
            <code>{{ data.primary_url }}</code>
        </a>
    </td>

    <td v-if="data.docroot"><code>/{{ data.docroot }}</code></td>
    <td v-else><code>/</code></td>

    <td>{{ data.type }}</td>

    <td>{{ data.status }}</td>

    <td>
        <button 
            v-on:click="api('/start_project', {
                'bla': 'HOND'
            })"
            class="btn btn-primary waves-effect waves-light"
        >
            <i class="fas fa-play"></i>
        </button>

        <button 
            v-on:click="api('/start_project', 1)"
            class="btn btn-info waves-effect waves-light"
        >
            <i class="fas fa-pause"></i>
        </button>

        <button 
            v-on:click="deleteProject(1)"
            class="btn btn-danger waves-effect waves-light"
        >
            <i class="fas fa-trash"></i>
        </button>
    </td>
</template>

<script>
import Swal from "sweetalert2";
import { api } from "../utils/api";

export default {
    props: {
        data: Object,
    },
    methods: {
        api: api,

        deleteProject: function() {
            Swal.fire({
                title: "Are you sure to deleting this project?",
                focusCancel: true,
                showCancelButton: true,
                confirmButtonText: "Confirm",
            }).then((result) => {
                if (result.isConfirmed) {
                    Swal.fire("Deleted", "", "success")
                } else if (result.isConfirmed) {
                    Swal.fire("Cancel", "", "info")
                }
            })
        }
    },
};
</script>

<style scoped lang="scss">
button {
    padding-left: 16px;
    padding-right: 16px;
}
</style>
