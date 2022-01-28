<template>
    <div class="home">
              <span class="relative inline-flex">
      <button type="button" class="inline-flex items-center px-4 py-2 font-semibold leading-6 text-sm shadow rounded-md text-sky-500 bg-white dark:bg-slate-800 transition ease-in-out duration-150 cursor-not-allowed ring-1 ring-slate-900/10 dark:ring-slate-200/20" disabled="">
        Transactions
      </button>
      <span class="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
        <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-sky-400 opacity-75"></span>
        <span class="relative inline-flex rounded-full h-3 w-3 bg-sky-500"></span>
      </span>
    </span>


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
